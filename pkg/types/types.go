package types

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// ------------------------------------------
// common interfaces
// ------------------------------------------

// Stringer a type which can be serializes as a string
type Stringer interface {
	Stringify() string
}

// WithAttributes base interface for types on which attributes can be substituted
type WithAttributes interface {
	GetAttributes() Attributes
	AddAttributes(Attributes)
	SetAttributes(Attributes)
}

type WithElementAddition interface {
	AddElement(interface{}) error
}

type WithElements interface {
	WithAttributes
	GetElements() []interface{}
	SetElements([]interface{}) error
}

type Filterable interface {
	IsEmpty() bool
}

type WithTitle interface {
	WithAttributes
	GetTitle() []interface{}
	SetTitle([]interface{}) error
}

type WithLocation interface {
	WithAttributes
	GetLocation() *Location
}

type Referencable interface {
	Reference(refs ElementReferences)
}

// ------------------------------------------
// Substitution support
// ------------------------------------------

// DocumentFragment a single fragment of document
type DocumentFragment struct {
	Position Position
	Elements []interface{}
	Error    error
}

type Position struct {
	Start int
	End   int
}

func NewDocumentFragment(p Position, elements ...interface{}) DocumentFragment {
	return DocumentFragment{
		Position: p,
		Elements: elements,
	}
}

func NewErrorFragment(p Position, err error) DocumentFragment {
	return DocumentFragment{
		Position: p,
		Error:    err,
	}
}

// ------------------------------------------
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	Elements          []interface{}
	ElementReferences ElementReferences
	Footnotes         []*Footnote
	TableOfContents   *TableOfContents
}

// Header returns the header, i.e., the section with level 0 if it found as the first element of the document
// For manpage documents, this also includes the first section (`Name` along with its first paragraph)
func (d *Document) Header() *DocumentHeader {
	if len(d.Elements) == 0 {
		log.Debug("no header for empty doc")
		return nil
	}
	// expect header (if it exists) to be in first position of the doc
	if h, ok := d.Elements[0].(*DocumentHeader); ok {
		return h
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("no header in document: %T", d.Elements[0])
	}
	return nil
}

// BodyElements returns the elements to render in the body
func (d *Document) BodyElements() []interface{} {
	if len(d.Elements) == 0 {
		return nil
	}
	elements := make([]interface{}, 0, len(d.Elements))
	for i, e := range d.Elements {
		if _, ok := e.(*DocumentHeader); ok {
			elements = append(elements, d.Elements[i+1:]...)
			return elements
		}
		elements = append(elements, e)
	}
	return elements
}

var _ WithElementAddition = &Document{}

func (d *Document) AddElement(element interface{}) error {
	d.Elements = append(d.Elements, element)
	return nil
}

// ------------------------------------------
// Document Metadata
// ------------------------------------------

// Metadata the document metadata returned after the rendering
type Metadata struct {
	Title           string
	LastUpdated     string
	TableOfContents *TableOfContents
	Authors         []*DocumentAuthor
	Revision        DocumentRevision
}

func NewTableOfContents(maxDepth int) *TableOfContents {
	log.Debugf("new TableOfContents with depth=%d", maxDepth)
	return &TableOfContents{
		MaxDepth: maxDepth,
	}
}

// TableOfContents the table of contents
type TableOfContents struct {
	MaxDepth int
	Sections []*ToCSection
}

// ToCSection a section in the table of contents
type ToCSection struct {
	ID       string
	Level    int
	Title    string // the title as it was rendered in HTML
	Number   string // the number assigned during rendering, if the `sectnums` attribute was set
	Children []*ToCSection
}

// Add adds a ToCSection associated with the given Section
func (t *TableOfContents) Add(s *Section) {
	if s.Level > t.MaxDepth {
		log.Debugf("skipping section with level %d (> %d)", s.Level, t.MaxDepth)
		// skip for elements with a too low level in the hierarchy
		return
	}
	ts := &ToCSection{
		ID:    s.GetAttributes().GetAsStringWithDefault(AttrID, ""),
		Level: s.Level,
	}
	// lookup the last child at the given section's level
	if len(t.Sections) == 0 {
		t.Sections = []*ToCSection{ts}
		return
	} else if s.Level == t.Sections[0].Level {
		// add at top level
		t.Sections = append(t.Sections, ts)
		return
	}
	// look-up parent for the ts, starting with the last section at root level
	parent := t.Sections[len(t.Sections)-1]
	for {
		if len(parent.Children) == 0 ||
			parent.Children[0].Level == s.Level {
			// (first) child level matches section level
			// or no child beneath current parent
			break
		}
		// move to last child of current parent
		parent = parent.Children[len(parent.Children)-1]
	}
	parent.Children = append(parent.Children, ts)
}

// ------------------------------------------
// Document Element
// ------------------------------------------

// DocumentElement a document element can have attributes
type DocumentElement interface {
	GetAttributes() Attributes
}

// ------------------------------------------
// Document Header
// ------------------------------------------

type DocumentHeader struct {
	Title      []interface{}
	Attributes Attributes
	Elements   []interface{}
}

func NewDocumentHeader(info interface{}, extraAttrs []interface{}) (*DocumentHeader, error) {
	header := &DocumentHeader{}
	elements := make([]interface{}, 0, 2+len(extraAttrs)) // estimated max capacity
	if info, ok := info.(*DocumentInformation); ok {
		header.Title = info.Title
		if len(info.Authors) > 0 {
			// header.Authors = info.Authors
			elements = append(elements, &AttributeDeclaration{
				Name:  AttrAuthors,
				Value: info.Authors,
			})
		}
		if info.Revision != nil {
			// header.Elements = append(header.Elements, info.Revision.Expand()...)
			elements = append(elements, &AttributeDeclaration{
				Name:  AttrRevision,
				Value: info.Revision,
			})
		}

	}
	elements = append(elements, extraAttrs...)
	if len(elements) > 0 {
		header.Elements = elements
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new doc header: %s", spew.Sdump(header))
	}
	return header, nil
}

func (h *DocumentHeader) Authors() DocumentAuthors {
	for _, e := range h.Elements {
		if e, ok := e.(*AttributeDeclaration); ok && e.Name == AttrAuthors {
			if authors, ok := e.Value.(DocumentAuthors); ok {
				return authors
			}
		}
	}
	return nil
}

func (h *DocumentHeader) Revision() *DocumentRevision {
	for _, e := range h.Elements {
		if e, ok := e.(*AttributeDeclaration); ok && e.Name == AttrRevision {
			if revision, ok := e.Value.(*DocumentRevision); ok {
				return revision
			}
		}
	}
	return nil
}

var _ Filterable = &DocumentHeader{}

func (h *DocumentHeader) IsEmpty() bool {
	return len(h.Title) == 0 &&
		len(h.Attributes) == 0 &&
		len(h.Elements) == 0
}

var _ WithTitle = &DocumentHeader{}

func (h *DocumentHeader) GetTitle() []interface{} {
	return h.Title
}

func (h *DocumentHeader) SetTitle(title []interface{}) error {
	h.Title = title
	return nil
}

var _ WithElements = &DocumentHeader{}

func (h *DocumentHeader) GetElements() []interface{} {
	return h.Elements
}

func (h *DocumentHeader) SetElements(elements []interface{}) error {
	h.Elements = elements
	return nil
}

func (h *DocumentHeader) GetAttributes() Attributes {
	return h.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (h *DocumentHeader) AddAttributes(attributes Attributes) {
	h.Attributes = h.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (h *DocumentHeader) SetAttributes(attributes Attributes) {
	h.Attributes = attributes
	if _, exists := h.Attributes[AttrID]; exists {
		// needed to track custom ID during rendering
		h.Attributes[AttrCustomID] = true
	}
}

type DocumentInformation struct {
	Title []interface{}
	DocumentAuthorsAndRevision
}

func NewDocumentInformation(title []interface{}, ar interface{}) (*DocumentInformation, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("new document info")
		log.Debugf("authors and revision: %s", spew.Sdump(ar))
	}
	info := &DocumentInformation{
		Title: title,
	}
	if ar, ok := ar.(*DocumentAuthorsAndRevision); ok {
		info.DocumentAuthorsAndRevision = *ar
	}
	return info, nil
}

type DocumentAuthorsAndRevision struct {
	Authors  DocumentAuthors
	Revision *DocumentRevision
}

func NewDocumentAuthorsAndRevision(authors DocumentAuthors, revision interface{}) (*DocumentAuthorsAndRevision, error) {
	ar := &DocumentAuthorsAndRevision{
		Authors: authors,
	}
	if r, ok := revision.(*DocumentRevision); ok {
		ar.Revision = r
	}
	return ar, nil
}

// ------------------------------------------
// Document Author
// ------------------------------------------

type DocumentAuthors []*DocumentAuthor

// NewDocumentAuthors converts the given authors into an array of `DocumentAuthor`
func NewDocumentAuthors(authors ...interface{}) (DocumentAuthors, error) {
	// log.Debugf("new array of document authors from `%+v`", authors)
	result := make([]*DocumentAuthor, len(authors))
	for i, author := range authors {
		switch author := author.(type) {
		case *DocumentAuthor:
			result[i] = author
		default:
			return nil, errors.Errorf("unexpected type of author: %T", author)
		}
	}
	return result, nil
}

// Expand returns a map of attributes for the given authors.
// those attributes can be used in attribute substitutions in the document
func (authors DocumentAuthors) Expand() Attributes {
	result := Attributes{}
	result[AttrAuthors] = []*DocumentAuthor(authors)
	for i, author := range authors {
		result[key("author", i)] = author.FullName()
		result[key("authorinitials", i)] = author.Initials()
		result[key("firstname", i)] = author.FirstName
		if author.MiddleName != "" {
			result[key("middlename", i)] = author.MiddleName
		}
		if author.LastName != "" {
			result[key("lastname", i)] = author.LastName
		}
		if author.Email != "" {
			result[key("email", i)] = author.Email
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("authors: %s", spew.Sdump(result))
	}
	return result
}

func key(k string, i int) string {
	if i == 0 {
		return k
	}
	return k + "_" + strconv.Itoa(i+1)
}

func initial(s string) string {
	if len(s) > 0 {
		return s[0:1]
	}
	return ""
}

// DocumentAuthor a document author
type DocumentAuthor struct {
	*DocumentAuthorFullName
	Email string
}

// NewDocumentAuthor initializes a new DocumentAuthor
func NewDocumentAuthor(fullName, email interface{}) (*DocumentAuthor, error) {
	author := &DocumentAuthor{}
	if fullName, ok := fullName.(*DocumentAuthorFullName); ok {
		author.DocumentAuthorFullName = fullName
	}
	if email, ok := email.(string); ok {
		author.Email = strings.TrimSpace(email)
	}
	return author, nil
}

type DocumentAuthorFullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func NewDocumentAuthorFullName(part1 string, part2, part3 interface{}) (*DocumentAuthorFullName, error) {
	result := &DocumentAuthorFullName{
		FirstName: strings.ReplaceAll(part1, "_", " "),
	}
	// if part 3 is defined, then it's the last name
	if lastName, ok := part3.(string); ok {
		result.LastName = strings.TrimSpace(strings.ReplaceAll(lastName, "_", " "))
		if middleName, ok := part2.(string); ok {
			result.MiddleName = strings.ReplaceAll(middleName, "_", " ")
		}

	} else if lastName, ok := part2.(string); ok {
		// otherwise, let's use part2 as the last name (and skip the middle name)
		result.LastName = strings.ReplaceAll(lastName, "_", " ")
	}
	return result, nil
}

func (n *DocumentAuthorFullName) FullName() string {
	result := &strings.Builder{}
	result.WriteString(n.FirstName)
	if n.MiddleName != "" {
		result.WriteString(" ")
		result.WriteString(n.MiddleName)
	}
	if n.LastName != "" {
		result.WriteString(" ")
		result.WriteString(n.LastName)
	}
	return result.String()
}

func (n *DocumentAuthorFullName) Initials() string {
	return strings.Join([]string{
		initial(n.FirstName),
		initial(n.MiddleName),
		initial(n.LastName),
	}, "")
}

// ------------------------------------------
// Document Revision
// ------------------------------------------

// DocumentRevision a document revision
type DocumentRevision struct {
	Revnumber string
	Revdate   string
	Revremark string
}

// NewDocumentRevision intializes a new DocumentRevision
func NewDocumentRevision(revnumber, revdate, revremark interface{}) (*DocumentRevision, error) {
	// log.Debugf("initializing document revision with revnumber=%v, revdate=%v, revremark=%v", revnumber, revdate, revremark)
	// remove the "v" prefix and trim spaces
	var number, date, remark string
	if revnumber, ok := revnumber.(string); ok {
		number = Apply(revnumber,
			func(s string) string {
				return strings.TrimPrefix(s, "v")
			}, func(s string) string {
				return strings.TrimPrefix(s, "V")
			}, func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	if revdate, ok := revdate.(string); ok {
		// trim spaces
		date = Apply(revdate,
			func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	if revremark, ok := revremark.(string); ok {
		// then we need to strip the leading ":" and spaces
		remark = Apply(revremark,
			func(s string) string {
				return strings.TrimPrefix(s, ":")
			}, func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	return &DocumentRevision{
		Revnumber: number,
		Revdate:   date,
		Revremark: remark,
	}, nil
}

// Expand returns a map of attributes for the given revision.
// those attributes can be used in attribute substitutions in the document
func (r *DocumentRevision) Expand() Attributes {
	result := Attributes{}
	result[AttrRevision] = r
	if r.Revnumber != "" {
		result["revnumber"] = r.Revnumber
	}
	if r.Revdate != "" {
		result["revdate"] = r.Revdate
	}
	if r.Revremark != "" {
		result["revremark"] = r.Revremark
	}
	// log.Debugf("revision: %v", result)
	return result
}

// ------------------------------------------
// Document Attributes
// ------------------------------------------

// AttributeDeclaration the type for Document Attribute Declarations
type AttributeDeclaration struct {
	Name    string
	Value   interface{}
	rawText string
}

// NewAttributeDeclaration initializes a new AttributeDeclaration with the given name and optional value
func NewAttributeDeclaration(name string, value interface{}, rawText string) (*AttributeDeclaration, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new attribute declaration: '%s'", spew.Sdump(rawText))
	}
	return &AttributeDeclaration{
		Name:    name,
		Value:   Reduce(value, strings.TrimSpace),
		rawText: rawText,
	}, nil
}

func (a *AttributeDeclaration) RawText() string {
	return a.rawText
}

// AttributeReset the type for AttributeReset
type AttributeReset struct {
	Name    string
	rawText string
}

// NewAttributeReset initializes a new Document Attribute Resets.
func NewAttributeReset(attrName string, rawText string) (*AttributeReset, error) {
	// log.Debugf("new AttributeReset: '%s'", attrName)
	return &AttributeReset{
		Name:    attrName,
		rawText: rawText,
	}, nil
}

func (a *AttributeReset) RawText() string {
	return a.rawText
}

// AttributeReference the type for AttributeReference
type AttributeReference struct {
	Name    string
	rawText string
}

// NewAttributeSubstitution initializes a new Attribute Substitutions
func NewAttributeSubstitution(name, rawText string) (interface{}, error) {
	if isPrefedinedAttribute(name) {
		return &PredefinedAttribute{
				Name:    name,
				rawText: rawText},
			nil
	}
	return &AttributeReference{
			Name:    name,
			rawText: rawText},
		nil
}

// PredefinedAttribute a special kind of attribute substitution, which
// uses a predefined attribute
type PredefinedAttribute AttributeReference

// CounterSubstitution is a counter, that may increment when it is substituted.
// If Increment is set, then it will increment before being expanded.
type CounterSubstitution struct {
	Name    string
	Hidden  bool
	Value   interface{} // may be a byte for character
	rawText string
}

// NewCounterSubstitution returns a counter substitution.
func NewCounterSubstitution(name string, hidden bool, val interface{}, rawText string) (*CounterSubstitution, error) {
	if v, ok := val.(string); ok {
		val = rune(v[0])
	}
	return &CounterSubstitution{
		Name:    name,
		Hidden:  hidden,
		Value:   val,
		rawText: rawText,
	}, nil
}

// ------------------------------------------
// Preamble
// ------------------------------------------

// Preamble the structure for document Preamble
type Preamble struct {
	Elements        []interface{}
	TableOfContents *TableOfContents
}

// HasContent returns `true` if this Preamble has at least one element which is neither a
// BlankLine nor a AttributeDeclaration
func (p *Preamble) HasContent() bool {
	for _, pe := range p.Elements {
		switch pe.(type) {
		case *BlankLine, *AttributeDeclaration, *AttributeReset:
			continue
		default:
			return true
		}
	}
	return false
}

// ------------------------------------------
// Front Matter
// ------------------------------------------

// FrontMatter the structure for document front-matter
type FrontMatter struct {
	Attributes map[string]interface{}
}

// NewYamlFrontMatter initializes a new FrontMatter from the given `content`
func NewYamlFrontMatter(content string) (*FrontMatter, error) {
	attributes := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(content), &attributes); err != nil {
		return nil, errors.Wrapf(err, "failed to parse the yaml content in the front-matter block")
	}
	if len(attributes) == 0 {
		attributes = nil
	}
	// log.Debugf("new FrontMatter with attributes: %+v", attributes)
	return &FrontMatter{
		Attributes: attributes,
	}, nil
}

// ------------------------------------------
// Lists
// ------------------------------------------

// ListElement a list item
type ListElement interface { // TODO: convert to struct and use as composant in OrderedListElement, etc.
	WithElements
	WithElementAddition
	LastElement() interface{}
	ListKind() ListKind
	AdjustStyle(*List)
	matchesStyle(ListElement) bool
}

func addToListElement(e ListElement, element interface{}) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("adding element of type '%T' to '%T'", element, e)
	}
	switch element := element.(type) {
	case RawLine:
		// append to last element of this OrderedListElement if it's a Paragraph,
		// otherwise, append a new Paragraph with this RawLine
		if p, ok := e.LastElement().(*Paragraph); ok {
			return p.AddElement(element)
		}
		return e.SetElements(append(e.GetElements(), &Paragraph{
			Elements: []interface{}{
				element,
			},
		}))

	default:
		return e.SetElements(append(e.GetElements(), element))
	}
}

type List struct {
	Kind       ListKind
	Attributes Attributes
	Elements   []ListElement
}

// CanAddElement checks if the given element can be added
func (l *List) CanAddElement(element interface{}) bool {
	switch e := element.(type) {
	case ListElement:
		// any listelement can be added if there was no blankline before
		// otherwise, only accept list element with attribute if there is no blankline before
		return e.ListKind() == l.Kind && e.matchesStyle(l.LastElement()) // TODO: compare to `FirstElement` is enough and faster
	case *ListContinuation:
		return true
	default:
		return false
	}
}

// AddElement adds the given element `e` in the target list or sublist (depending on its type)
func (l *List) AddElement(element interface{}) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("adding element of type '%T' to list of kind '%s'", element, l.Kind)
	}
	switch e := element.(type) {
	case ListElement:
		if e.ListKind() == l.Kind {
			l.Elements = append(l.Elements, e)
			// interactive unordered list elements
			if u, ok := element.(*UnorderedListElement); ok && l.Attributes.HasOption(AttrInteractive) {
				u.interactive()
			}
			return nil
		}
	case *ListContinuation:
		return l.LastElement().AddElement(e.Element)
	}

	return errors.Errorf("cannot add element of type '%T' to list of kind '%s'", element, l.Kind)
}

func (l *List) Reference(refs ElementReferences) {
	id := l.Attributes.GetAsStringWithDefault(AttrID, "")
	title := l.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit elements
	for _, e := range l.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

func (l *List) LastElement() ListElement {
	if len(l.Elements) == 0 {
		return nil
	}
	return l.Elements[len(l.Elements)-1]
}

type ListElements struct {
	Elements []interface{}
}

func NewListElements(elements []interface{}) (*ListElements, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("initializing new ListElements with \n%s", spew.Sdump(elements...))
	}
	elmts := make([]interface{}, 0, len(elements))
	var attrs Attributes
	// attributes must be attached to their following element
	for _, e := range elements {
		switch e := e.(type) {
		case Attributes:
			attrs = attrs.AddAll(e)
		case WithAttributes:
			if attrs != nil {
				e.SetAttributes(attrs)
				attrs = nil
			}
			elmts = append(elmts, e)
		case RawLine, *SinglelineComment:
			// append to last element
			if len(elmts) > 0 {
				switch elmt := elmts[len(elmts)-1].(type) {
				case WithElementAddition:
					if err := elmt.AddElement(e); err != nil {
						return nil, errors.Errorf("unable to attach element of type '%T' in a list", e)
					}
				default:
					return nil, errors.Errorf("unable to attach element of type '%T' in a list", e)
				}
			}
		case *ListContinuation:
			attrs = nil           // couldn't attach attribute to element, so discard it
			if e.Element != nil { // only retain list element continuations with actual content
				elmts = append(elmts, e)
			}
		default:
			attrs = nil // couldn't attach attribute to element, so discard it
			elmts = append(elmts, e)
		}
	}
	result := &ListElements{
		Elements: elmts,
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("initialized new ListElements: %s", spew.Sdump(result))
	// }
	return result, nil
}

var _ WithElements = &ListElements{}

func (l *ListElements) GetAttributes() Attributes {
	return nil // unused
}

func (l *ListElements) AddAttributes(attrs Attributes) {
	// set attribute on first element
	if len(l.Elements) > 0 {
		if e, ok := l.Elements[0].(ListElement); ok {
			e.AddAttributes(attrs)
		}
	}
}

func (l *ListElements) SetAttributes(Attributes) {
	// unused
}

// GetElements returns the elements
func (l *ListElements) GetElements() []interface{} {
	return l.Elements
}

// SetElements set the elements
func (l *ListElements) SetElements(elements []interface{}) error {
	l.Elements = elements
	return nil
}

var _ WithFootnotes = &ListElements{}

// SubstituteFootnotes replaces the footnotes in the list element
// with footnote references. The footnotes are stored in the given 'notes' param
func (l *ListElements) SubstituteFootnotes(notes *Footnotes) {
	for _, e := range l.Elements {
		if e, ok := e.(WithFootnotes); ok {
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("collecting footnotes in element of type '%T'", e)
			}
			e.SubstituteFootnotes(notes)
		}
	}
}

type ListKind string

const (
	LabeledListKind   ListKind = "labeled_list"
	OrderedListKind   ListKind = "ordered_list"
	UnorderedListKind ListKind = "unordered_list"
	CalloutListKind   ListKind = "callout_list"
)

func NewList(element ListElement) (*List, error) {
	// also, move the element attributes to the List
	attrs := element.GetAttributes()
	element.SetAttributes(nil)
	list := &List{
		Kind:       element.ListKind(),
		Attributes: attrs,
		Elements: []ListElement{
			element,
		},
	}
	// interactive unordered list elements (applies to all elements of the list)
	// TODO: move into `types.GenericList.AddElement()`` ?
	if u, ok := element.(*UnorderedListElement); ok && list.Attributes.HasOption(AttrInteractive) {
		u.interactive()
	}

	return list, nil
}

type ListContinuation struct {
	Offset  int
	Element interface{}
}

func NewListContinuation(offset int, Element interface{}) (*ListContinuation, error) {
	return &ListContinuation{
		Offset:  offset,
		Element: Element,
	}, nil
}

var _ WithElementAddition = &ListContinuation{}

func (c *ListContinuation) AddElement(element interface{}) error {
	if e, ok := c.Element.(WithElementAddition); ok {
		return e.AddElement(element)
	}
	return errors.Errorf("cannot add element of type '%T' to list element continuation", c.Element)
}

// ------------------------------------------
// Callouts
// ------------------------------------------

// Callout a reference at the end of a line in a delimited block with verbatim content (eg: listing, source code)
type Callout struct {
	Ref int
}

// NewCallout returns a new Callout with the given reference
func NewCallout(ref int) (*Callout, error) {
	return &Callout{
		Ref: ref,
	}, nil
}

// CalloutListElement the description of a call out which will appear as an ordered list item after the delimited block
type CalloutListElement struct {
	Attributes Attributes
	Ref        int
	Elements   []interface{}
}

var _ ListElement = &CalloutListElement{}

var _ DocumentElement = &CalloutListElement{}

// NewCalloutListElement returns a new CalloutListElement
func NewCalloutListElement(ref int, content *Paragraph) (*CalloutListElement, error) {
	return &CalloutListElement{
		Attributes: nil,
		Ref:        ref,
		Elements: []interface{}{
			content,
		},
	}, nil
}

var _ Referencable = &CalloutListElement{}

func (e *CalloutListElement) Reference(refs ElementReferences) {
	id := e.Attributes.GetAsStringWithDefault(AttrID, "")
	title := e.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit elements
	for _, e := range e.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

// checks if the given list element matches the level of this element
func (e *CalloutListElement) matchesStyle(other ListElement) bool {
	_, ok := other.(*CalloutListElement)
	return ok // no level in Callout lists
}

func (e *CalloutListElement) AdjustStyle(l *List) {
	// do nothing, there's a single level in callout lists
}

// ListKind returns the kind of list to which this element shall be attached
func (e *CalloutListElement) ListKind() ListKind {
	return CalloutListKind
}

// GetAttributes returns the attributes of this element
func (e *CalloutListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this element
func (e *CalloutListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes of this element
func (e *CalloutListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
}

var _ WithElementAddition = &CalloutListElement{}

// AddElement add an element to this UnorderedListElement
func (e *CalloutListElement) AddElement(element interface{}) error {
	return addToListElement(e, element)
}

// GetElements returns this CalloutListElement's elements
func (e *CalloutListElement) GetElements() []interface{} {
	return e.Elements
}

func (e *CalloutListElement) LastElement() interface{} {
	if len(e.Elements) == 0 {
		return nil
	}
	return e.Elements[len(e.Elements)-1]
}

// SetElements sets this CalloutListElement's elements
func (e *CalloutListElement) SetElements(elements []interface{}) error {
	e.Elements = elements
	return nil
}

var _ WithFootnotes = &CalloutListElement{}

// SubstituteFootnotes replaces the footnotes in the list element
// with footnote references. The footnotes are stored in the given 'notes' param
func (e *CalloutListElement) SubstituteFootnotes(notes *Footnotes) {
	for _, e := range e.Elements {
		if e, ok := e.(WithFootnotes); ok {
			e.SubstituteFootnotes(notes)
		}
	}
}

const (
	// TODO: define a `NumberingStyle` type
	// Arabic the arabic numbering (1, 2, 3, etc.)
	Arabic = "arabic"
	// LowerAlpha the lower-alpha numbering (a, b, c, etc.)
	LowerAlpha = "loweralpha"
	// UpperAlpha the upper-alpha numbering (A, B, C, etc.)
	UpperAlpha = "upperalpha"
	// LowerRoman the lower-roman numbering (i, ii, iii, etc.)
	LowerRoman = "lowerroman"
	// UpperRoman the upper-roman numbering (I, II, III, etc.)
	UpperRoman = "upperroman"

	// Other styles are possible, but "uppergreek", "lowergreek", but aren't
	// generated automatically.
)

// OrderedListElement the structure for the ordered list items
type OrderedListElement struct {
	Attributes Attributes
	Style      string // TODO: rename to `OrderedListElementNumberingStyle`? TODO: define as an attribute instead?
	Elements   []interface{}
}

// making sure that the `ListElement` interface is implemented by `OrderedListElement`
var _ ListElement = &OrderedListElement{}

// NewOrderedListElement initializes a new `orderedListItem` from the given content
func NewOrderedListElement(prefix OrderedListElementPrefix, content interface{}) (*OrderedListElement, error) {
	// log.Debugf("new OrderedListElement")
	return &OrderedListElement{
		Style: prefix.Style,
		Elements: []interface{}{
			content,
		},
	}, nil
}

var _ Referencable = &OrderedListElement{}

func (e *OrderedListElement) Reference(refs ElementReferences) {
	id := e.Attributes.GetAsStringWithDefault(AttrID, "")
	title := e.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit elements
	for _, e := range e.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

// checks if the given list element matches the level of this element
func (e *OrderedListElement) matchesStyle(other ListElement) bool {
	if element, ok := other.(*OrderedListElement); ok {
		return e.Style == element.Style
	}
	return false
}

func (e *OrderedListElement) AdjustStyle(l *List) {
	// do nothing
}

// ListKind returns the kind of list to which this element shall be attached
func (e *OrderedListElement) ListKind() ListKind {
	return OrderedListKind
}

// GetElements returns this item's elements
func (e *OrderedListElement) GetElements() []interface{} {
	return e.Elements
}

func (e *OrderedListElement) LastElement() interface{} {
	if len(e.Elements) == 0 {
		return nil
	}
	return e.Elements[len(e.Elements)-1]
}

// SetElements sets this OrderedListElement's elements
func (e *OrderedListElement) SetElements(elements []interface{}) error {
	e.Elements = elements
	return nil
}

var _ WithElementAddition = &OrderedListElement{}

// AddElement add an element to this UnorderedListElement
func (e *OrderedListElement) AddElement(element interface{}) error {
	return addToListElement(e, element)
}

var _ WithAttributes = &OrderedListElement{}

// GetAttributes returns this element's attributes
func (e *OrderedListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this element
func (e *OrderedListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// SetAttributes sets the attributes in this element
func (e *OrderedListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *OrderedListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
}

var _ WithFootnotes = &OrderedListElement{}

// SubstituteFootnotes replaces the footnotes in the list element
// with footnote references. The footnotes are stored in the given 'notes' param
func (e *OrderedListElement) SubstituteFootnotes(notes *Footnotes) {
	for _, e := range e.Elements {
		if e, ok := e.(WithFootnotes); ok {
			e.SubstituteFootnotes(notes)
		}
	}
}

// OrderedListElementPrefix the prefix used to construct an OrderedListElement
type OrderedListElementPrefix struct {
	Style string
}

// NewOrderedListElementPrefix initializes a new OrderedListElementPrefix
func NewOrderedListElementPrefix(s string) (OrderedListElementPrefix, error) {
	return OrderedListElementPrefix{
		Style: s,
	}, nil
}

// ------------------------------------------
// Unordered Lists
// ------------------------------------------

// UnorderedListElement the structure for the unordered list items
type UnorderedListElement struct {
	BulletStyle UnorderedListElementBulletStyle
	CheckStyle  UnorderedListElementCheckStyle
	Attributes  Attributes
	Elements    []interface{}
}

var _ ListElement = &UnorderedListElement{}

// NewUnorderedListElement initializes a new `UnorderedListElement` from the given content
func NewUnorderedListElement(prefix UnorderedListElementPrefix, checkstyle interface{}, content interface{}) (*UnorderedListElement, error) {
	// log.Debugf("new UnorderedListElement with %d elements", len(elements))
	cs := toCheckStyle(checkstyle)
	if cs != NoCheck {
		if p, ok := content.(*Paragraph); ok {
			if p.Attributes == nil {
				p.Attributes = Attributes{}
			}
			p.Attributes[AttrCheckStyle] = cs
		}
	}
	return &UnorderedListElement{
		BulletStyle: prefix.BulletStyle,
		CheckStyle:  cs,
		Elements: []interface{}{
			content,
		},
	}, nil
}

var _ Referencable = &UnorderedListElement{}

func (e *UnorderedListElement) Reference(refs ElementReferences) {
	id := e.Attributes.GetAsStringWithDefault(AttrID, "")
	title := e.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit elements
	for _, e := range e.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

// checks if the given list element matches the level of this element
func (e *UnorderedListElement) matchesStyle(other ListElement) bool {
	if other, ok := other.(*UnorderedListElement); ok {
		log.Debugf("checking if list elements match: %v/%v", e.BulletStyle, other.BulletStyle)
		return e.BulletStyle == other.BulletStyle
	}
	return false
}

func (e *UnorderedListElement) AdjustStyle(l *List) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("attempt to adjust bulletstyle '%v' compared to\n%s", e.BulletStyle, spew.Sdump(l))
	}
	if l != nil && len(l.Elements) > 0 {
		if o, ok := l.Elements[0].(*UnorderedListElement); ok {
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("adjusting bulletstyle from '%v' to '%v'", e.BulletStyle, o.BulletStyle.next())
			}
			e.BulletStyle = o.BulletStyle.next()
		}
	}
}

// ListKind returns the kind of list to which this element shall be attached
func (e *UnorderedListElement) ListKind() ListKind {
	return UnorderedListKind
}

func (e *UnorderedListElement) interactive() {
	e.CheckStyle = e.CheckStyle.interactive()
	if e.CheckStyle != NoCheck && len(e.Elements) > 0 {
		if p, ok := e.Elements[0].(*Paragraph); ok {
			if p.Attributes == nil {
				p.Attributes = Attributes{}
				// e.Elements[0] = p // need to update the element in the slice
			}
			p.Attributes[AttrCheckStyle] = e.CheckStyle
		}
	}
}

var _ WithElementAddition = &UnorderedListElement{}

// AddElement add an element to this UnorderedListElement
func (e *UnorderedListElement) AddElement(element interface{}) error {
	return addToListElement(e, element)
}

// GetElements returns this UnorderedListElement's elements
func (e *UnorderedListElement) GetElements() []interface{} {
	return e.Elements
}

func (e *UnorderedListElement) LastElement() interface{} {
	if len(e.Elements) == 0 {
		return nil
	}
	return e.Elements[len(e.Elements)-1]
}

// SetElements sets this UnorderedListElement's elements
func (e *UnorderedListElement) SetElements(elements []interface{}) error {
	e.Elements = elements
	return nil
}

var _ WithAttributes = &UnorderedListElement{}

// GetAttributes returns this element's attributes
func (e *UnorderedListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this element
func (e *UnorderedListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// SetAttributes replaces the attributes in this element
func (e *UnorderedListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *UnorderedListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
}

var _ WithFootnotes = &UnorderedListElement{}

// SubstituteFootnotes replaces the footnotes in the list element
// with footnote references. The footnotes are stored in the given 'notes' param
func (e *UnorderedListElement) SubstituteFootnotes(notes *Footnotes) {
	for _, e := range e.Elements {
		if e, ok := e.(WithFootnotes); ok {
			e.SubstituteFootnotes(notes)
		}
	}
}

// UnorderedListElementCheckStyle the check style that applies on an unordered list item
type UnorderedListElementCheckStyle string

const (
	// Checked when the unordered list item is checked
	Checked UnorderedListElementCheckStyle = "checked"
	// CheckedInteractive when the unordered list item is checked (with an interactive checkbox)
	CheckedInteractive UnorderedListElementCheckStyle = "checked-interactive"
	// Unchecked when the unordered list item is not checked
	Unchecked UnorderedListElementCheckStyle = "unchecked"
	// UncheckedInteractive when the unordered list item is not checked (with an interactive checkbox)
	UncheckedInteractive UnorderedListElementCheckStyle = "unchecked-interactive"
	// NoCheck when the unodered list item has no specific check annotation
	NoCheck UnorderedListElementCheckStyle = "nocheck"
)

func (s UnorderedListElementCheckStyle) interactive() UnorderedListElementCheckStyle {
	switch s {
	case Checked, CheckedInteractive:
		return CheckedInteractive
	case Unchecked, UncheckedInteractive:
		return UncheckedInteractive
	default:
		return NoCheck
	}
}

func toCheckStyle(checkstyle interface{}) UnorderedListElementCheckStyle {
	if cs, ok := checkstyle.(UnorderedListElementCheckStyle); ok {
		return cs
	}
	return NoCheck
}

// UnorderedListElementBulletStyle the type of bullet for items in an unordered list
type UnorderedListElementBulletStyle string

func (s UnorderedListElementBulletStyle) next() UnorderedListElementBulletStyle {
	switch s {
	case Dash:
		return OneAsterisk
	case OneAsterisk:
		return TwoAsterisks
	case TwoAsterisks:
		return ThreeAsterisks
	case ThreeAsterisks:
		return FourAsterisks
	default:
		return FiveAsterisks
	}
}

const (
	// Dash an unordered item can begin with a single dash
	Dash UnorderedListElementBulletStyle = "dash"
	// OneAsterisk an unordered item marked with a single asterisk
	OneAsterisk UnorderedListElementBulletStyle = "1asterisk"
	// TwoAsterisks an unordered item marked with two asterisks
	TwoAsterisks UnorderedListElementBulletStyle = "2asterisks"
	// ThreeAsterisks an unordered item marked with three asterisks
	ThreeAsterisks UnorderedListElementBulletStyle = "3asterisks"
	// FourAsterisks an unordered item marked with four asterisks
	FourAsterisks UnorderedListElementBulletStyle = "4asterisks"
	// FiveAsterisks an unordered item marked with five asterisks
	FiveAsterisks UnorderedListElementBulletStyle = "5asterisks"
)

func NewUnorderedListElementBulletStyle(style string) (UnorderedListElementBulletStyle, error) {
	switch style {
	case "-":
		return Dash, nil
	case "*":
		return OneAsterisk, nil
	case "**":
		return TwoAsterisks, nil
	case "***":
		return ThreeAsterisks, nil
	case "****":
		return FourAsterisks, nil
	case "*****":
		return FiveAsterisks, nil
	default:
		return "", fmt.Errorf("unexpected unordered list element bullet style: '%s'", style)
	}
}

// UnorderedListElementPrefix the prefix used to construct an UnorderedListElement
type UnorderedListElementPrefix struct {
	BulletStyle UnorderedListElementBulletStyle
}

// NewUnorderedListElementPrefix initializes a new UnorderedListElementPrefix
func NewUnorderedListElementPrefix(style string) (UnorderedListElementPrefix, error) {
	s, err := NewUnorderedListElementBulletStyle(style)
	if err != nil {
		return UnorderedListElementPrefix{}, err
	}
	return UnorderedListElementPrefix{
		BulletStyle: s,
	}, nil
}

// ------------------------------------------
// Labeled List
// ------------------------------------------

type LabeledListElementStyle string

const (
	DoubleColons    LabeledListElementStyle = "::"
	TripleColons    LabeledListElementStyle = ":::"
	QuadrupleColons LabeledListElementStyle = "::::"
)

func toLabeledListElementStyle(level int) (LabeledListElementStyle, error) {
	switch level {
	case 1:
		return DoubleColons, nil
	case 2:
		return TripleColons, nil
	case 3:
		return QuadrupleColons, nil
	default:
		return LabeledListElementStyle(""), fmt.Errorf("unsupported level of labeled list element: %d", level)
	}
}

// LabeledListElement an item in a labeled
type LabeledListElement struct {
	Term       []interface{}
	Attributes Attributes
	Style      LabeledListElementStyle
	Elements   []interface{}
}

// NewLabeledListElement initializes a new LabeledListElement
func NewLabeledListElement(level int, term, description interface{}) (*LabeledListElement, error) {
	// log.Debugf("new LabeledListElement")
	t := []interface{}{
		term,
	}
	style, err := toLabeledListElementStyle(level)
	if err != nil {
		return nil, err
	}
	var elements []interface{}
	if desc, ok := description.(*Paragraph); ok {
		elements = []interface{}{
			desc,
		}
	}
	return &LabeledListElement{
		Style:    style,
		Term:     t,
		Elements: elements,
	}, nil
}

var _ Referencable = &LabeledListElement{}

func (e *LabeledListElement) Reference(refs ElementReferences) {
	id := e.Attributes.GetAsStringWithDefault(AttrID, "")
	title := e.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit the term
	if len(e.Term) > 0 {
		if anchor, ok := e.Term[0].(*InlineLink); ok {
			if id := anchor.Attributes.GetAsStringWithDefault(AttrID, ""); id != "" {
				refs[id] = e.Term[1:]
			}
		}
	}
	// also, visit elements
	for _, e := range e.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

// making sure that the `ListElement` interface is implemented by `LabeledListElement`
var _ ListElement = &LabeledListElement{}

// checks if the given list element matches the style of this element
func (e *LabeledListElement) matchesStyle(other ListElement) bool {
	if element, ok := other.(*LabeledListElement); ok {
		return e.Style == element.Style
	}
	return false
}

func (e *LabeledListElement) AdjustStyle(l *List) {
	// do nothing
}

// ListKind returns the kind of list to which this element shall be attached
func (e *LabeledListElement) ListKind() ListKind {
	return LabeledListKind
}

var _ WithElementAddition = &LabeledListElement{}

// AddElement add an element to this LabeledListElement
func (e *LabeledListElement) AddElement(element interface{}) error {
	return addToListElement(e, element)
}

// GetElements returns this LabeledListElement's elements
func (e *LabeledListElement) GetElements() []interface{} {
	return e.Elements
}

func (e *LabeledListElement) LastElement() interface{} {
	if len(e.Elements) == 0 {
		return nil
	}
	return e.Elements[len(e.Elements)-1]
}

// SetElements sets this LabeledListElement's elements
func (e *LabeledListElement) SetElements(elements []interface{}) error {
	e.Elements = elements
	return nil
}

var _ WithAttributes = &LabeledListElement{}

// GetAttributes returns this element's attributes
func (e *LabeledListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this element
func (e *LabeledListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// SetAttributes sets the attributes in this element
func (e *LabeledListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *LabeledListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
}

var _ WithFootnotes = &LabeledListElement{}

// SubstituteFootnotes replaces the footnotes in the list element
// with footnote references. The footnotes are stored in the given 'notes' param
func (e *LabeledListElement) SubstituteFootnotes(notes *Footnotes) {
	for i, element := range e.Term {
		if note, ok := element.(*Footnote); ok {
			e.Term[i] = notes.Reference(note)
		}
	}
	for _, element := range e.Elements {
		if e, ok := element.(WithFootnotes); ok {
			e.SubstituteFootnotes(notes)
		}
	}
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraphs
type Paragraph struct {
	Attributes Attributes
	Elements   []interface{}
}

// NewParagraph initializes a new `Paragraph`
func NewParagraph(style interface{}, elements ...interface{}) (*Paragraph, error) {
	// log.Debugf("new paragraph with attributes: '%v'", attributes)
	for i, l := range elements {
		if l, ok := l.(RawLine); ok {
			// add `\n` unless the we're on the last element
			if i < len(elements)-1 {
				elements[i] = RawLine(l + "\n") // TODO: add `NewRawlines()` func which takes care of appending with "\n"
			}
		}
	}
	p := &Paragraph{
		Elements: elements,
	}
	if style != nil {
		p.AddAttributes(Attributes{
			AttrStyle: style,
		})
	}
	return p, nil
}

var _ WithElements = &Paragraph{}

// GetElements returns this paragraph's elements (or lines)
func (p *Paragraph) GetElements() []interface{} {
	return p.Elements
}

// SetElements sets this paragraph's elements
func (p *Paragraph) SetElements(elements []interface{}) error {
	p.Elements = elements
	return nil
}

var _ WithElementAddition = &Paragraph{}

func (p *Paragraph) AddElement(e interface{}) error {
	if r, ok := p.Elements[len(p.Elements)-1].(RawLine); ok {
		p.Elements[len(p.Elements)-1] = RawLine(r + "\n")
	}
	p.Elements = append(p.Elements, e)
	return nil
}

var _ WithAttributes = &Paragraph{}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (p *Paragraph) GetAttributes() Attributes {
	return p.Attributes
}

// AddAttributes adds the attributes of this element
func (p *Paragraph) AddAttributes(attributes Attributes) {
	p.Attributes = p.Attributes.AddAll(attributes)
	p.mapAttributes()
}

// SetAttributes sets the attributes in this element
func (p *Paragraph) SetAttributes(attributes Attributes) {
	p.Attributes = attributes
	p.mapAttributes()
}

func (p *Paragraph) mapAttributes() {
	p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	switch p.Attributes[AttrStyle] {
	case Source:
		p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
			AttrPositional2: AttrLanguage,
		})
	case Quote, Verse:
		p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
			AttrPositional2: AttrQuoteAuthor,
			AttrPositional3: AttrQuoteTitle,
		})
	}
}

var _ Referencable = &Paragraph{}

func (p *Paragraph) Reference(refs ElementReferences) {
	id := p.Attributes.GetAsStringWithDefault(AttrID, "")
	title := p.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
}

var _ WithFootnotes = &Paragraph{}

// SubstituteFootnotes replaces the footnotes in the paragraph lines
// with footnote references. The footnotes are stored in the given 'notes' param
func (p *Paragraph) SubstituteFootnotes(notes *Footnotes) {
	for i, element := range p.Elements {
		if note, ok := element.(*Footnote); ok {
			p.Elements[i] = notes.Reference(note)
		}
	}
}

// ------------------------------------------
// Admonitions
// ------------------------------------------

// TODO: use custom `AdmonitionKind` type
const (
	// Tip the 'TIP' type of admonition
	Tip string = "TIP"
	// Note the 'NOTE' type of admonition
	Note string = "NOTE"
	// Important the 'IMPORTANT' type of admonition
	Important string = "IMPORTANT"
	// Warning the 'WARNING' type of admonition
	Warning string = "WARNING"
	// Caution the 'CAUTION' type of admonition
	Caution string = "CAUTION"
)

// ------------------------------------------
// Inline Elements
// ------------------------------------------

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements ...interface{}) ([]interface{}, error) {
	elements = merge(elements...)
	// due to grammar optimization (and a bit hack-ish): remove space suffix from `*StringElement`` when followed by `*LineBreak`
	for i, e := range elements {
		if _, ok := e.(*LineBreak); ok && i > 0 {
			if s, ok := elements[i-1].(*StringElement); ok {
				s.Content = strings.TrimSuffix(s.Content, " ")
			}
		}
	}
	return elements, nil
}

// ------------------------------------------
// Cross References
// ------------------------------------------

// InternalCrossReference the struct for Cross References
type InternalCrossReference struct {
	ID    interface{}
	Label interface{}
}

// NewInternalCrossReference initializes a new `InternalCrossReference` from the given ID
func NewInternalCrossReference(id, label interface{}) (*InternalCrossReference, error) {
	// log.Debugf("new InternalCrossReference with ID=%s", id)
	return &InternalCrossReference{
		ID:    Reduce(id),
		Label: Reduce(label),
	}, nil
}

// ExternalCrossReference the struct for Cross References
type ExternalCrossReference struct {
	Location   *Location
	Attributes Attributes
}

// NewExternalCrossReference initializes a new `InternalCrossReference` from the given ID
func NewExternalCrossReference(location *Location, attributes interface{}) (*ExternalCrossReference, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrXRefLabel,
	})
	// log.Debugf("new ExternalCrossReference with Location=%v and label='%s' (attrs=%v / %T)", location, label, attributes, attrs[AttrInlineLinkText])
	return &ExternalCrossReference{
		Location:   location,
		Attributes: attrs,
	}, nil
}

var _ WithLocation = &ExternalCrossReference{}

func (x *ExternalCrossReference) GetLocation() *Location {
	return x.Location
}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (x *ExternalCrossReference) GetAttributes() Attributes {
	return x.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (x *ExternalCrossReference) AddAttributes(attributes Attributes) {
	x.Attributes = x.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (x *ExternalCrossReference) SetAttributes(attributes Attributes) {
	x.Attributes = x.Attributes.SetAll(attributes)
}

// ------------------------------------------
// Inline Button
// ------------------------------------------

// InlineButton a button (requires `experimental` doc attribute to be set)
type InlineButton struct {
	Attributes Attributes
}

// NewInlineButton initializes a new `InlineButton`
func NewInlineButton(attrs Attributes) (*InlineButton, error) {
	return &InlineButton{
		Attributes: toAttributesWithMapping(
			attrs, map[string]string{
				AttrPositional1: AttrButtonLabel,
			},
		),
	}, nil
}

// ------------------------------------------
// Inline Menu
// ------------------------------------------

// InlineMenu a menu with optional subpaths defined in its attributes (requires `experimental` doc attribute to be set)
type InlineMenu struct {
	Path []string
}

// NewInlineMenu initializes a new `InlineMenu`
func NewInlineMenu(id string, attrs Attributes) (*InlineMenu, error) {
	path := []string{id}
	if s, ok := attrs[AttrPositional1].(string); ok {
		subpaths := strings.Split(s, ">")
		for _, s := range subpaths {
			path = append(path, strings.TrimSpace(s))
		}
	}
	return &InlineMenu{
		Path: path,
	}, nil
}

// ------------------------------------------
// Images
// ------------------------------------------

// ImageBlock the structure for the block images
type ImageBlock struct {
	Location   *Location
	Attributes Attributes
}

// NewImageBlock initializes a new `ImageBlock`
func NewImageBlock(location *Location, inlineAttributes Attributes) (*ImageBlock, error) {
	// inline attributes trump block attributes
	attrs := Attributes{}
	attrs.SetAll(inlineAttributes)
	attrs = toAttributesWithMapping(attrs, map[string]string{
		AttrPositional1: AttrImageAlt,
		AttrPositional2: AttrWidth,
		AttrPositional3: AttrHeight,
	})
	return &ImageBlock{
		Location:   location,
		Attributes: attrs,
	}, nil
}

var _ WithAttributes = &ImageBlock{}

// GetAttributes returns this element's attributes
func (i *ImageBlock) GetAttributes() Attributes {
	return i.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (i *ImageBlock) AddAttributes(attributes Attributes) {
	i.Attributes = i.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (i *ImageBlock) SetAttributes(attributes Attributes) {
	i.Attributes = attributes
}

var _ WithLocation = &ImageBlock{}

func (i *ImageBlock) GetLocation() *Location {
	return i.Location
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Location   *Location
	Attributes Attributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(location *Location, attributes interface{}) (*InlineImage, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrImageAlt,
		AttrPositional2: AttrWidth,
		AttrPositional3: AttrHeight,
	})
	return &InlineImage{
		Attributes: attrs,
		Location:   location,
	}, nil
}

var _ WithAttributes = &InlineImage{}

// GetAttributes returns this inline image's attributes
func (i *InlineImage) GetAttributes() Attributes {
	return i.Attributes
}

// AddAttributes adds the attributes of this element
func (i *InlineImage) AddAttributes(attributes Attributes) {
	i.Attributes = i.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (i *InlineImage) SetAttributes(attributes Attributes) {
	i.Attributes = attributes
}

var _ WithLocation = &InlineImage{}

func (i *InlineImage) GetLocation() *Location {
	return i.Location
}

// ------------------------------------------
// Icons
// ------------------------------------------

// Icon an icon
type Icon struct {
	Class      string
	Attributes Attributes
}

// NewIcon initializes a new `Icon`
func NewIcon(class string, attributes interface{}) (*Icon, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrIconSize,
	})
	return &Icon{
		Class:      class,
		Attributes: attrs,
	}, nil
}

// ------------------------------------------
// Footnotes
// ------------------------------------------

// type Footnotes []*Footnote

// Footnote a foot note, without or without explicit reference (an explicit reference is used to refer
// multiple times to the same footnote across the document)
type Footnote struct {
	// ID is only set during document processing
	ID int
	// Ref the optional reference
	Ref string
	// the footnote content (can be "rich")
	Elements []interface{}
}

// NewFootnote returns a new Footnote with the given content
func NewFootnote(ref interface{}, elements []interface{}) (*Footnote, error) {
	log.Debugf("new footnote with elements: '%s'", spew.Sdump(elements))
	var r string
	if ref, ok := ref.(string); ok {
		r = ref
	}
	return &Footnote{
		// ID is only set during document processing
		Ref:      r,
		Elements: elements,
	}, nil

}

// FootnoteReference a footnote reference. Substitutes the actual footnote in the document,
// and only contains a generated, sequential ID (which will be displayed)
type FootnoteReference struct {
	ID        int
	Ref       string // the user-specified reference (optional)
	Duplicate bool   // indicates if this reference targets an already-existing footnote // TODO: find a better name?
}

// WithFootnotes interface for all types which may contain footnotes
type WithFootnotes interface {
	SubstituteFootnotes(existing *Footnotes)
}

// Footnotes the footnotes of a document. Footnotes are "collected"
// during the parsing phase and displayed at the bottom of the document
// during the rendering.
type Footnotes struct {
	sequence *sequence
	Notes    []*Footnote
}

// NewFootnotes initializes a new Footnotes
func NewFootnotes() *Footnotes {
	return &Footnotes{
		sequence: &sequence{},
		Notes:    []*Footnote{},
	}
}

// IndexOf returns the index of the given note in the footnotes.
func (f *Footnotes) indexOf(actual *Footnote) (int, bool) {
	for _, note := range f.Notes {
		if note.Ref == actual.Ref {
			return note.ID, true
		}
	}
	return -1, false
}

const (
	// InvalidFootnoteReference a constant to mark the footnote reference as invalid
	InvalidFootnoteReference int = -1
)

// Reference adds the given footnote and returns a FootnoteReference in replacement
func (f *Footnotes) Reference(note *Footnote) *FootnoteReference {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("referencing footnote")
	}
	r := &FootnoteReference{}
	if len(note.Elements) > 0 {
		note.ID = f.sequence.nextVal()
		f.Notes = append(f.Notes, note)
		r.ID = note.ID
	} else if id, found := f.indexOf(note); found {
		r.ID = id
		r.Duplicate = true
	} else {
		r.ID = InvalidFootnoteReference
		log.Warnf("no footnote with reference '%s'", note.Ref)
	}
	r.Ref = note.Ref
	return r
}

type sequence struct {
	counter int
}

func (s *sequence) nextVal() int {
	s.counter++
	return s.counter
}

// ------------------------------------------
// Delimited blocks
// ------------------------------------------

// TODO: use custom `BlockKind` type
const (
	// Fenced a fenced block
	Fenced string = "fenced"
	// Listing a listing block
	Listing string = "listing"
	// Example an example block
	Example string = "example"
	// Comment a comment block
	Comment string = "comment"
	// Quote a quote block
	Quote string = "quote"
	// MarkdownCode a code block in the Markdown style
	MarkdownCode string = "markdown-code"
	// MarkdownQuote a quote block in the Markdown style
	MarkdownQuote string = "markdown-quote"
	// Open an Open block
	Open string = "open"
	// Verse a verse block
	Verse string = "verse"
	// Sidebar a sidebar block
	Sidebar string = "sidebar"
	// Literal a literal block
	Literal string = "literal"
	// LiteralParagraph a literal parsgraph
	LiteralParagraph = "literal_paragraph"
	// Source a source block
	Source string = "source"
	// Passthrough a passthrough block
	Passthrough string = "pass"
)

// LiteralParagraph custom type to retain the number of spaces on the first line (needed during rendering)

type BlockDelimiter struct { // TODO: use string directly?
	Kind       string
	Length     int
	Attributes Attributes
	rawText    string
}

func NewBlockDelimiter(kind string, length int, rawText string) (*BlockDelimiter, error) {
	return &BlockDelimiter{
		Kind:    kind,
		Length:  length,
		rawText: rawText,
	}, nil
}

func NewMarkdownCodeBlockDelimiter(language, rawText string) (*BlockDelimiter, error) {
	return &BlockDelimiter{
		Kind:   Fenced,
		Length: 3,
		Attributes: Attributes{
			AttrStyle:    Source,
			AttrLanguage: language,
		},
		rawText: rawText,
	}, nil
}

func (b *BlockDelimiter) RawText() string {
	return b.rawText
}

// DelimitedBlock the structure for the Listing blocks
type DelimitedBlock struct {
	Kind       string
	Attributes Attributes
	Elements   []interface{}
}

func NewDelimitedBlock(kind string, elements []interface{}) (*DelimitedBlock, error) {
	for i, l := range elements {
		if l, ok := l.(RawLine); ok {
			// add `\n` unless the we're on the last element
			if i < len(elements)-1 {
				elements[i] = RawLine(l + "\n")
			}
		}
	}
	return &DelimitedBlock{
		Kind:     kind,
		Elements: elements,
	}, nil
}

var _ WithElements = &DelimitedBlock{}

// GetElements returns this paragraph's elements (or lines)
func (b *DelimitedBlock) GetElements() []interface{} {
	return b.Elements
}

// SetElements sets this paragraph's elements
func (b *DelimitedBlock) SetElements(elements []interface{}) error {
	if len(elements) > 0 {
		switch b.Kind {
		case Listing, Literal:
			// preserve space but discard empty lines
			// log.Debugf("discarding leading crlf on elements in block of kind '%s'", b.Kind)
			// discard leading spaces and CR/LF
			if s, ok := elements[0].(*StringElement); ok {
				s.Content = strings.TrimLeft(s.Content, "\r\n")
			}
		default:
			// log.Debugf("discarding leading spaces+crlf on elements in block of kind '%s'", b.Kind)
			// discard leading spaces and CR/LF
			if s, ok := elements[0].(*StringElement); ok {
				s.Content = strings.TrimLeft(s.Content, " \t\r\n")
			}
		}
		// discard trailing spaces and CR/LF
		// log.Debugf("discarding trailing spaces+crlf on elements in block of kind '%s'", b.Kind)
		if s, ok := elements[len(elements)-1].(*StringElement); ok {
			s.Content = strings.TrimRight(s.Content, " \t\r\n")
		}
	}
	b.Elements = elements
	return nil
}

var _ WithAttributes = &DelimitedBlock{}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (b *DelimitedBlock) GetAttributes() Attributes {
	return b.Attributes
}

// AddAttributes adds the attributes of this element
func (b *DelimitedBlock) AddAttributes(attributes Attributes) {
	b.Attributes = b.Attributes.AddAll(attributes)
	b.mapAttributes()
}

// SetAttributes sets the attributes in this element
func (b *DelimitedBlock) SetAttributes(attributes Attributes) {
	b.Attributes = attributes
	b.mapAttributes()
}

func (b *DelimitedBlock) mapAttributes() {
	switch b.Kind {
	case Quote:
		b.Attributes = toAttributesWithMapping(b.Attributes, map[string]string{
			AttrPositional1: AttrStyle,
			AttrPositional2: AttrQuoteAuthor,
			AttrPositional3: AttrQuoteTitle,
		})
		// override the `kind` with `style` attribute (if exists)
		if style, exists := b.Attributes[AttrStyle].(string); exists {
			b.Kind = style
		}
	case Example:
		b.Attributes = toAttributesWithMapping(b.Attributes, map[string]string{
			AttrPositional1: AttrStyle,
		})
	case Listing:
		b.Attributes = toAttributesWithMapping(b.Attributes, map[string]string{
			AttrPositional1: AttrStyle,
			AttrPositional2: AttrLanguage,
			AttrPositional3: AttrLineNums,
		})
	}
}

var _ Referencable = &DelimitedBlock{}

func (b *DelimitedBlock) Reference(refs ElementReferences) {
	id := b.Attributes.GetAsStringWithDefault(AttrID, "")
	title := b.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
	// also, visit elements
	for _, e := range b.Elements {
		if e, ok := e.(Referencable); ok {
			e.Reference(refs)
		}
	}
}

// ------------------------------------------
// Sections
// ------------------------------------------

// RawSection the structure for a rawText section, using during preparsing (needed to support level offsets)
type RawSection struct {
	Level   int
	RawText string
}

func NewRawSection(level int, rawText string) (*RawSection, error) {
	return &RawSection{
		Level:   level,
		RawText: rawText,
	}, nil
}

func (s *RawSection) OffsetLevel(offset int) {
	s.RawText = strings.Replace(s.RawText, strings.Repeat("=", s.Level+1)+" ", strings.Repeat("=", s.Level+1+offset)+" ", 1)
	s.Level = s.Level + offset
}

func (s *RawSection) Stringify() string {
	return s.RawText
}

// Section the structure for a section
type Section struct {
	Level      int
	Attributes Attributes
	Title      []interface{}
	Elements   []interface{}
}

// NewSection returns a new Section
func NewSection(level int, title []interface{}) (*Section, error) {
	// log.Debugf("new rawsection: '%s' (%d)", title, level)
	return &Section{
		Level: level,
		Title: title,
	}, nil
}

func (s *Section) GetID() (string, error) {
	id, _, err := s.Attributes.GetAsString(AttrID)
	return id, err
}

var _ WithElements = &Section{}

// GetElements returns this Section's elements
func (s *Section) GetElements() []interface{} {
	return s.Elements
}

// SetElements sets this Sections's elements
func (s *Section) SetElements(elements []interface{}) error {
	s.Elements = elements
	return nil
}

var _ WithTitle = &Section{}

// GetTitle returns this section's title
func (s *Section) GetTitle() []interface{} {
	return s.Title
}

// SetTitle sets this section's title
func (s *Section) SetTitle(title []interface{}) error {
	// inline ID attribute foud at the end is *moved* at the attributes level of the section
	if id, ok := title[len(title)-1].(*Attribute); ok {
		sectionID := stringify(id.Value)
		s.AddAttributes(Attributes{
			AttrID:       sectionID,
			AttrCustomID: true,
		})
		title = title[:len(title)-1]
	}
	s.Title = title
	return nil
}

// GetAttributes returns this section's attributes
func (s *Section) GetAttributes() Attributes {
	return s.Attributes
}

// AddAttributes adds the attributes of this element
func (s *Section) AddAttributes(attributes Attributes) {
	s.Attributes = s.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (s *Section) SetAttributes(attributes Attributes) {
	s.Attributes = attributes
	if _, exists := s.Attributes[AttrID]; exists {
		// needed to track custom ID during rendering
		s.Attributes[AttrCustomID] = true
	}
}

// ResolveID resolves/updates the "ID" attribute in the section (in case the title changed after some document attr substitution)
func (s *Section) ResolveID(attrs map[string]interface{}, refs ElementReferences) error {
	base, err := s.resolveID(attrs)
	if err != nil {
		return err
	}

	for i := 1; ; i++ {
		var id string
		if i == 1 {
			id = base
		} else {
			id = base + "_" + strconv.Itoa(i)
			log.Debugf("updated section id to '%s' (to avoid duplicate refs)", s.Attributes[AttrID])
		}
		if _, exists := refs[id]; !exists {
			s.Attributes[AttrID] = id
			break
		}
	}
	return nil
}

// resolveID resolves/updates the "ID" attribute in the section (in case the title changed after some document attr substitution)
func (s *Section) resolveID(attrs Attributes) (string, error) {
	if s.Attributes == nil {
		s.Attributes = Attributes{}
	}
	// block attribute
	if id := s.Attributes.GetAsStringWithDefault(AttrID, ""); id != "" {
		return id, nil
	}
	log.Debugf("resolving section id")
	prefix := attrs.GetAsStringWithDefault(AttrIDPrefix, DefaultIDPrefix)
	separator := attrs.GetAsStringWithDefault(AttrIDSeparator, DefaultIDSeparator)
	id, err := ReplaceNonAlphanumerics(s.Title, prefix, separator)
	if err != nil {
		return "", errors.Wrapf(err, "failed to generate default ID on Section element")
	}
	s.Attributes[AttrID] = id
	log.Debugf("updated section id to '%s'", s.Attributes[AttrID])
	return id, nil
}

var _ Referencable = &Section{}

func (s *Section) Reference(refs ElementReferences) {
	id := s.Attributes.GetAsStringWithDefault(AttrID, "")
	if id != "" && s.Title != nil {
		refs[id] = s.Title
	}
}

var _ WithElementAddition = &Section{}

// AddElement adds the given child element to this section
func (s *Section) AddElement(e interface{}) error {
	s.Elements = append(s.Elements, e)
	return nil
}

var _ WithFootnotes = &Section{}

// SubstituteFootnotes replaces the footnotes in the section title
// with footnote references. The footnotes are stored in the given 'notes' param
func (s *Section) SubstituteFootnotes(notes *Footnotes) {
	for i, element := range s.Title {
		if note, ok := element.(*Footnote); ok {
			s.Title[i] = notes.Reference(note)
		}
	}
}

// ------------------------------------------
// Table of Contents
// ------------------------------------------

// TableOfContentsPlaceHolder a place holder for Table of Contents, so
// the renderer knows when to render it.
type TableOfContentsPlaceHolder struct {
}

// ------------------------------------------
// Thematic breaks
// ------------------------------------------

// ThematicBreak a thematic break
type ThematicBreak struct{}

// NewThematicBreak returns a new ThematicBreak
func NewThematicBreak() (*ThematicBreak, error) {
	return &ThematicBreak{}, nil
}

// ------------------------------------------
// User Macro
// ------------------------------------------

const (
	// InlineMacro a inline user macro
	InlineMacro MacroKind = "inline"
	// BlockMacro a block user macro
	BlockMacro MacroKind = "block"
)

// MacroKind the type of user macro
type MacroKind string

// UserMacro the structure for User Macro
type UserMacro struct {
	Kind       MacroKind
	Name       string
	Value      string
	Attributes Attributes
	RawText    string
}

// NewUserMacroBlock returns an UserMacro
func NewUserMacroBlock(name string, value string, attributes Attributes, rawText string) (*UserMacro, error) {
	return &UserMacro{
		Name:       name,
		Kind:       BlockMacro,
		Value:      value,
		Attributes: attributes,
		RawText:    rawText,
	}, nil
}

// NewInlineUserMacro returns an UserMacro
func NewInlineUserMacro(name, value string, attributes Attributes, rawText string) (*UserMacro, error) {
	return &UserMacro{
		Name:       name,
		Kind:       InlineMacro,
		Value:      value,
		Attributes: attributes,
		RawText:    rawText,
	}, nil
}

// ------------------------------------------
// BlankLine
// ------------------------------------------

// BlankLine the structure for the empty lines, which are used to separate logical blocks
type BlankLine struct {
}

// NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (*BlankLine, error) {
	return &BlankLine{}, nil
}

// ------------------------------------------
// Comments
// ------------------------------------------

// SinglelineComment a single line comment
type SinglelineComment struct {
	Content string
}

// NewSinglelineComment initializes a new single line content
func NewSinglelineComment(content string) (*SinglelineComment, error) {
	// log.Debugf("initializing a single line comment with content: '%s'", content)
	return &SinglelineComment{
		Content: content,
	}, nil
}

// ------------------------------------------
// StringElement
// ------------------------------------------

// StringElement the structure for strings
type StringElement struct {
	Content string
}

// NewStringElement initializes a new `StringElement` from the given content
func NewStringElement(content string) (*StringElement, error) {
	return &StringElement{
		Content: content,
	}, nil
}

// RawText returns the rawText text representation of this element as it was (supposedly) written in the source document
func (s StringElement) RawText() (string, error) {
	return s.Content, nil
}

// ------------------------------------------
// Explicit line breaks
// ------------------------------------------

// LineBreak an explicit line break in a paragraph
type LineBreak struct{}

// NewLineBreak returns a new line break, that's all.
func NewLineBreak() (*LineBreak, error) {
	return &LineBreak{}, nil
}

// ------------------------------------------
// Quoted text
// ------------------------------------------

// QuotedText the structure for quoted text
type QuotedText struct {
	Kind       QuotedTextKind
	Elements   []interface{}
	Attributes Attributes
}

// QuotedTextKind the type for
type QuotedTextKind string

const (
	// SingleQuoteBold bold quoted text (wrapped with '*')
	SingleQuoteBold = QuotedTextKind("*")
	// DoubleQuoteBold bold quoted text (wrapped with '**')
	DoubleQuoteBold = QuotedTextKind("**")
	// SingleQuoteItalic italic quoted text (wrapped with '_')
	SingleQuoteItalic = QuotedTextKind("_")
	// DoubleQuoteItalic italic quoted text (wrapped with '__')
	DoubleQuoteItalic = QuotedTextKind("__")
	// SingleQuoteMarked text highlighter (wrapped with '#')
	SingleQuoteMarked = QuotedTextKind("#")
	// DoubleQuoteMarked text highlighter (wrapped '##')
	DoubleQuoteMarked = QuotedTextKind("##")
	// SingleQuoteMonospace monospace quoted text (wrapped with '`')
	SingleQuoteMonospace = QuotedTextKind("`")
	// DoubleQuoteMonospace monospace quoted text (wrapped with '``')
	DoubleQuoteMonospace = QuotedTextKind("``")
	// SingleQuoteSubscript subscript quoted text (wrapped with '~')
	SingleQuoteSubscript = QuotedTextKind("~")
	// SingleQuoteSuperscript superscript quoted text (wrapped with '^')
	SingleQuoteSuperscript = QuotedTextKind("^")
)

// NewQuotedText initializes a new `QuotedText` from the given kind and content
func NewQuotedText(kind QuotedTextKind, elements ...interface{}) (*QuotedText, error) {
	return &QuotedText{
		Kind:     kind,
		Elements: merge(elements),
	}, nil
}

var _ WithElements = &QuotedText{}

// GetElements returns this QuotedText's elements
func (t *QuotedText) GetElements() []interface{} {
	return t.Elements
}

// SetElements sets this QuotedText's elements
func (t *QuotedText) SetElements(elements []interface{}) error {
	t.Elements = elements
	return nil
}

var _ WithAttributes = &QuotedText{}

// GetAttributes returns the attributes of this element
func (t *QuotedText) GetAttributes() Attributes {
	return t.Attributes
}

// AddAttributes adds the attributes of this element
func (t *QuotedText) AddAttributes(attributes Attributes) {
	t.Attributes = t.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (t *QuotedText) SetAttributes(attributes Attributes) {
	t.Attributes = attributes
}

// WithAttributes returns a _new_ QuotedText with the given attributes (with some mapping)
func (t *QuotedText) WithAttributes(attributes interface{}) (*QuotedText, error) {
	// log.Debugf("adding attributes on quoted text: %v", attributes)
	t.Attributes = toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrRoles,
	})
	return t, nil
}

// -------------------------------------------------------
// Escaped Quoted Text (i.e., with substitution preserved)
// -------------------------------------------------------

// NewEscapedQuotedText returns a new []interface{} where the nested elements are preserved (ie, substituted as expected)
func NewEscapedQuotedText(backslashes string, marker string, content interface{}) ([]interface{}, error) {
	// log.Debugf("new escaped quoted text: %s %s %v", backslashes, punctuation, content)
	backslashesStr := Apply(backslashes,
		func(s string) string {
			// remove the number of back-slashes that match the length of the punctuation. Eg: `\*` or `\\**`, but keep extra back-slashes
			if len(s) > len(marker) {
				return s[len(marker):]
			}
			return ""
		})
	return []interface{}{
		&StringElement{
			Content: backslashesStr + marker,
		},
		content,
		&StringElement{
			Content: marker,
		},
	}, nil
}

// ------------------------------------------
// InlinePassthrough
// ------------------------------------------

// InlinePassthrough the structure for Passthroughs
type InlinePassthrough struct {
	Kind     PassthroughKind
	Elements []interface{} // TODO: refactor to `Content string` ?
}

// PassthroughKind the kind of passthrough
type PassthroughKind string

const (
	// SinglePlusPassthrough a passthrough with a single `+` punctuation
	SinglePlusPassthrough = PassthroughKind("+")
	// TriplePlusPassthrough a passthrough with a triple `+++` punctuation
	TriplePlusPassthrough = PassthroughKind("+++")
	// PassthroughMacro a passthrough with the `pass:[]` macro
	PassthroughMacro = PassthroughKind("pass:[]")
)

// NewInlinePassthrough returns a new passthrough
func NewInlinePassthrough(kind PassthroughKind, elements []interface{}) (*InlinePassthrough, error) {
	return &InlinePassthrough{
		Kind:     kind,
		Elements: merge(elements...),
	}, nil
}

// ------------------------------------------
// Inline Links
// ------------------------------------------

// InlineLink the structure for the external links
type InlineLink struct {
	Attributes Attributes
	Location   *Location
}

// NewInlineLink initializes a new inline `InlineLink`
func NewInlineLink(url *Location, attributes interface{}) (*InlineLink, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrInlineLinkText,
	})
	// also, look for `^` suffix in `AttrInlineLinkText` attribute for a `_blank` target
	if text, found := attrs[AttrInlineLinkText].(string); found && strings.HasSuffix(text, "^") {
		attrs[AttrInlineLinkText] = strings.TrimSuffix(text, "^")
		attrs[AttrInlineLinkTarget] = "_blank"
	}
	return &InlineLink{
		Location:   url,
		Attributes: attrs,
	}, nil
}

// NewInlineAnchor initializes a new InlineLink map with a single entry for the ID using the given value
func NewInlineAnchor(id string) (*InlineLink, error) {
	return &InlineLink{
		Attributes: Attributes{
			AttrID: id,
		},
	}, nil
}

var _ WithAttributes = &InlineLink{}

// GetAttributes returns this link's attributes
func (l *InlineLink) GetAttributes() Attributes {
	return l.Attributes
}

// AddAttributes adds the attributes of this element
func (l *InlineLink) AddAttributes(attributes Attributes) {
	l.Attributes = l.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (l *InlineLink) SetAttributes(attributes Attributes) {
	l.Attributes = attributes
}

var _ WithLocation = &InlineLink{}

func (l *InlineLink) GetLocation() *Location {
	return l.Location
}

// ------------------------------------------
// Conditionals
// ------------------------------------------

type ConditionalInclusion interface {
	Eval(attributes map[string]interface{}) bool
	SingleLineContent() (string, bool)
}

type IfdefCondition struct {
	Name         string
	Substitution string
}

func NewIfdefCondition(name string, attr interface{}) (*IfdefCondition, error) {
	log.Debugf("new Ifdef::%s conditional inclusion", name)
	c := &IfdefCondition{
		Name: name,
	}
	if subs, ok := attr.(string); ok {
		c.Substitution = subs
	}
	return c, nil
}

var _ ConditionalInclusion = &IfdefCondition{}

func (c *IfdefCondition) Eval(attributes map[string]interface{}) bool {
	_, found := attributes[c.Name]
	return found
}

func (c *IfdefCondition) SingleLineContent() (string, bool) {
	return c.Substitution, c.Substitution != ""
}

type IfndefCondition struct {
	Name         string
	Substitution string
}

func NewIfndefCondition(name string, attr interface{}) (*IfndefCondition, error) {
	log.Debugf("new Ifndef::%s conditional inclusion", name)
	c := &IfndefCondition{
		Name: name,
	}
	if subs, ok := attr.(string); ok {
		c.Substitution = subs
	}
	return c, nil
}

var _ ConditionalInclusion = &IfndefCondition{}

func (c *IfndefCondition) Eval(attributes map[string]interface{}) bool {
	_, found := attributes[c.Name]
	return !found
}

func (c *IfndefCondition) SingleLineContent() (string, bool) {
	return c.Substitution, c.Substitution != ""
}

type IfevalCondition struct {
	Left    interface{}
	Right   interface{}
	Operand IfevalOperand
}

func NewIfevalCondition(left, right interface{}, operand IfevalOperand) (*IfevalCondition, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new Ifeval conditional inclusion")
	}
	return &IfevalCondition{
		Left:    left,
		Right:   right,
		Operand: operand,
	}, nil
}

var _ ConditionalInclusion = &IfevalCondition{}

func (c *IfevalCondition) Eval(attributes map[string]interface{}) bool {
	return c.Operand(c.left(attributes), c.right(attributes))
}

func (c *IfevalCondition) left(attributes map[string]interface{}) interface{} {
	if s, ok := c.Left.(*AttributeReference); ok {
		if v, found := attributes[s.Name]; found {
			return v
		}
	}
	return c.Left
}

func (c *IfevalCondition) right(attributes map[string]interface{}) interface{} {
	if s, ok := c.Right.(*AttributeReference); ok {
		if v, found := attributes[s.Name]; found {
			return v
		}
	}
	return c.Right
}

func (c *IfevalCondition) SingleLineContent() (string, bool) {
	return "", false
}

type IfevalOperand func(left, right interface{}) bool

var EqualOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v==%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left == right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left == right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left == right
		}
	}
	return false
}

func NewEqualOperand() (IfevalOperand, error) {
	return EqualOperand, nil
}

var NotEqualOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v!=%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left != right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left != right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left != right
		}
	}
	return false
}

func NewNotEqualOperand() (IfevalOperand, error) {
	return NotEqualOperand, nil
}

var LessThanOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v<%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left < right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left < right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left < right
		}
	}
	return false
}

func NewLessThanOperand() (IfevalOperand, error) {
	return LessThanOperand, nil
}

var LessOrEqualOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v<=%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left <= right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left <= right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left <= right
		}
	}
	return false
}

func NewLessOrEqualOperand() (IfevalOperand, error) {
	return LessOrEqualOperand, nil
}

var GreaterThanOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v>%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left > right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left > right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left > right
		}
	}
	return false
}

func NewGreaterThanOperand() (IfevalOperand, error) {
	return GreaterThanOperand, nil
}

var GreaterOrEqualOperand = func(left, right interface{}) bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("comparing %v>=%v", left, right)
	}
	// comparing strings
	if left, ok := left.(string); ok {
		if right, ok := right.(string); ok {
			return left >= right
		}
	}
	// comparing floats
	if left, ok := left.(float64); ok {
		if right, ok := right.(float64); ok {
			return left >= right
		}
	}
	// comparing ints
	if left, ok := left.(int); ok {
		if right, ok := right.(int); ok {
			return left >= right
		}
	}
	return false
}

func NewGreaterOrEqualOperand() (IfevalOperand, error) {
	return GreaterOrEqualOperand, nil
}

type EndOfCondition struct{}

func NewEndOfCondition() (*EndOfCondition, error) {
	log.Debug("new end of conditional inclusion")
	return &EndOfCondition{}, nil
}

// ------------------------------------------
// File Inclusions
// ------------------------------------------

// FileInclusion the structure for the file inclusions
type FileInclusion struct {
	Attributes Attributes
	Location   *Location
	RawText    string
}

// NewFileInclusion initializes a new inline `FileInclusion`
func NewFileInclusion(location *Location, attributes interface{}, rawText string) (*FileInclusion, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		"tag": "tags", // convert `tag` to `tags`
	})
	return &FileInclusion{
		Attributes: attrs,
		Location:   location,
		RawText:    rawText,
	}, nil
}

var _ WithLocation = &FileInclusion{}

func (f *FileInclusion) GetLocation() *Location {
	return f.Location
}

// GetAttributes returns this elements's attributes
func (f *FileInclusion) GetAttributes() Attributes {
	return f.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (f *FileInclusion) AddAttributes(attributes Attributes) {
	f.Attributes = f.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this element
func (f *FileInclusion) SetAttributes(attributes Attributes) {
	f.Attributes = attributes
}

// -------------------------------------------------------------------------------------
// Raw Line
// -------------------------------------------------------------------------------------
type RawLine string // TODO: convert to struct with `Content` field, or alias to StringElement

// NewRawLine returns a new RawLine wrapper for the given string
func NewRawLine(content string) (RawLine, error) {
	return RawLine(strings.TrimRight(content, " \t")), nil
}

// -------------------------------------------------------------------------------------
// LineRanges: one or more ranges of lines to limit the content of a file to include
// -------------------------------------------------------------------------------------

// NewLineRanges returns a slice of line ranges attribute for a file inclusion.
func NewLineRanges(ranges interface{}) LineRanges {
	switch ranges := ranges.(type) {
	case []interface{}:
		result := LineRanges{}
		for _, r := range ranges {
			if lr, ok := r.(LineRange); ok {
				result = append(result, lr)
			}
		}
		// sort the range by `start` line
		sort.Sort(result)
		return result
	case LineRange:
		return LineRanges{ranges}
	default:
		log.Warnf("invalid type of line range: '%T'", ranges)
		return LineRanges{}
	}
}

// LineRange the range of lines of the child doc to include in the master doc
// `Start` and `End` are the included limits of the child document
// - if there's a single line to include, then `End = Start`
// - if there is all remaining content after a given line (included), then `End = -1`
type LineRange struct {
	StartLine int
	EndLine   int
}

// NewLineRange returns a new line range
func NewLineRange(start, end int) (LineRange, error) {
	// log.Debugf("new multiline range: %d..%d", start, end)
	return LineRange{
		StartLine: start,
		EndLine:   end,
	}, nil
}

// LineRanges the ranges of lines of the child doc to include in the master doc
type LineRanges []LineRange

// Match checks if the given line number matches one of the line ranges
func (r LineRanges) Match(line int) bool {
	for _, lr := range r {
		if lr.StartLine <= line && (lr.EndLine >= line || lr.EndLine == -1) {
			return true
		}
		if lr.StartLine > line {
			// no need to carry on with the ranges
			return false
		}
	}
	return false
}

// make sure that the LineRanges type implements the `sort.Interface
var _ sort.Interface = LineRanges{}

func (r LineRanges) Len() int           { return len(r) }
func (r LineRanges) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r LineRanges) Less(i, j int) bool { return r[i].StartLine < r[j].StartLine }

// -------------------------------------------------------------------------------------
// TagRanges: one or more ranges of tags to limit the content of a file to include
// -------------------------------------------------------------------------------------

// TagRanges the ranges of tags of the child doc to include in the master doc
type TagRanges []TagRange

// NewTagRanges returns a slice of tag ranges attribute for a file inclusion.
func NewTagRanges(ranges interface{}) TagRanges {
	switch ranges := ranges.(type) {
	case []interface{}:
		result := TagRanges{}
		for _, r := range ranges {
			if lr, ok := r.(TagRange); ok {
				result = append(result, lr)
			}
		}
		return result
	case TagRange:
		return TagRanges{ranges}
	default:
		log.Warnf("invalid type of tag range: '%T'", ranges)
		return TagRanges{}
	}
}

// Match checks if the given tag matches one of the range
func (tr TagRanges) Match(line int, currentRanges CurrentRanges) bool {
	match := false
	// log.Debugf("checking line %d", line)

	// compare with expected tag ranges
	for _, t := range tr {
		if t.Name == "**" {
			match = true
			continue
		}
		for n, r := range currentRanges {
			// log.Debugf("checking if range %s (%v) matches one of %v", n, r, tr)
			if r.EndLine != -1 {
				// tag range is closed, skip
				continue
			} else if t.Name == "*" {
				match = t.Included
			} else if t.Name == n { //TODO: all accept '*', '**' snd '!'
				match = t.Included
			}
		}
	}

	return match
}

// TagRange the range to include or exclude from the file inclusion.
// The range is excluded if it is prefixed with '!'
// Also, '*' and '**' have a special meaning:
// - '*' means that all tag ranges are included (except the lines having the start and end ranges)
// - '**' means that all content is included, regardless of whether it is in a tag or not (except the lines having the start and end ranges)
type TagRange struct {
	Name     string
	Included bool
}

// NewTagRange returns a new TagRange
func NewTagRange(name string, included bool) (TagRange, error) {
	return TagRange{
		Name:     name,
		Included: included,
	}, nil
}

// CurrentRanges the current ranges, ie, as they are "discovered"
// while processing one line at a time in the file to include
type CurrentRanges map[string]*CurrentTagRange

// CurrentTagRange a tag range found while processing a document. When the 'start' tag is found,
// the `EndLine` is still unknown and thus its value is set to `-1`.
type CurrentTagRange struct {
	StartLine int
	EndLine   int
}

// -------------------------------------------------------------------------------------
// IncludedFileLine a line of a file that is being included
// -------------------------------------------------------------------------------------

// IncludedFileLine a line, containing rawText text and inclusion tags
type IncludedFileLine []interface{}

// NewIncludedFileLine returns a new IncludedFileLine
func NewIncludedFileLine(content []interface{}) (IncludedFileLine, error) {
	return IncludedFileLine(merge(content)), nil
}

// HasTag returns true if the line has at least one inclusion tag (start or end), false otherwise
func (l IncludedFileLine) HasTag() bool {
	for _, e := range l {
		if _, ok := e.(IncludedFileStartTag); ok {
			return true
		}
		if _, ok := e.(IncludedFileEndTag); ok {
			return true
		}
	}
	return false
}

// GetStartTag returns the first IncludedFileStartTag found in the line // TODO: support multiple tags on the same line ?
func (l IncludedFileLine) GetStartTag() (IncludedFileStartTag, bool) {
	for _, e := range l {
		if s, ok := e.(IncludedFileStartTag); ok {
			return s, true
		}
	}
	return IncludedFileStartTag{}, false
}

// GetEndTag returns the first IncludedFileEndTag found in the line // TODO: support multiple tags on the same line ?
func (l IncludedFileLine) GetEndTag() (IncludedFileEndTag, bool) {
	for _, e := range l {
		if s, ok := e.(IncludedFileEndTag); ok {
			return s, true
		}
	}
	return IncludedFileEndTag{}, false
}

// IncludedFileStartTag the type for the `tag::` macro
type IncludedFileStartTag struct {
	Value string
}

// NewIncludedFileStartTag returns a new IncludedFileStartTag
func NewIncludedFileStartTag(tag string) (IncludedFileStartTag, error) {
	return IncludedFileStartTag{Value: tag}, nil
}

// IncludedFileEndTag the type for the `end::` macro
type IncludedFileEndTag struct {
	Value string
}

// NewIncludedFileEndTag returns a new IncludedFileEndTag
func NewIncludedFileEndTag(tag string) (IncludedFileEndTag, error) {
	return IncludedFileEndTag{Value: tag}, nil
}

// -------------------------------------------------------------------------------------
// Location: a Location (ie, with a scheme) or a path to a file (can be absolute or relative)
// -------------------------------------------------------------------------------------

// Location a Location contains characters and optionaly, document attributes
type Location struct {
	Scheme string
	Path   interface{}
}

// NewLocation return a new location with the given elements
func NewLocation(scheme interface{}, path []interface{}) (*Location, error) {
	// log.Debugf("new location: scheme='%v' path='%+v", scheme, path)
	s := ""
	if scheme, ok := scheme.([]byte); ok {
		s = string(scheme)
	}
	return &Location{
		Scheme: s,
		Path:   Reduce(path),
	}, nil
}

func (l *Location) SetPath(path interface{}) {
	p := Reduce(path)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("setting path in location: %v", p)
	}
	l.Path = p
}

// SetPathPrefix adds the given prefix to the path if this latter is NOT an absolute
// path and if there is no defined scheme
func (l *Location) SetPathPrefix(p interface{}) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("setting path with prefix: '%s' + '%s'", p, spew.Sdump(l.Path))
	}
	if p, ok := p.(string); ok && p != "" {
		if !strings.HasSuffix(p, "/") {
			p = p + "/"
		}
		if l.Scheme == "" && !strings.HasPrefix(l.Stringify(), "/") {
			if u, err := url.Parse(l.Stringify()); err == nil {
				if !u.IsAbs() {
					l.SetPath(merge(p, l.Path))
				}
			}
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("set path with prefix: '%s'", spew.Sdump(l.Path))
	}
}

func (l *Location) TrimAngleBracketSuffix() (bool, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("trimming angle bracket suffix in %s", spew.Sdump(l.Path))
	}
	if p, ok := l.Path.([]interface{}); ok {
		if c, ok := p[len(p)-1].(*SpecialCharacter); ok && c.Name == ">" {
			l.Path = Reduce(p[:len(p)-1]) // trim last element
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("trimmed angle bracket suffix in location: %s", spew.Sdump(l))
			}
			return true, nil
		}
	}
	log.Debug("no angle brack suffix to trim in location")
	return false, nil
}

// Stringify returns a string representation of the location
// or empty string if the location is nil
func (l *Location) Stringify() string {
	if l == nil {
		return ""
	}
	result := &strings.Builder{}
	result.WriteString(l.Scheme)
	result.WriteString(stringify(l.Path))
	return result.String()
}

// -------------------------------------------------------------------------------------
// Index terms
// -------------------------------------------------------------------------------------

// IndexTerm a index term, with a single term
type IndexTerm struct {
	Term []interface{}
}

// NewIndexTerm returns a new IndexTerm
func NewIndexTerm(term []interface{}) (*IndexTerm, error) {
	return &IndexTerm{
		Term: term,
	}, nil
}

// ConcealedIndexTerm a concealed index term, with 1 required and 2 optional terms
type ConcealedIndexTerm struct {
	Term1 interface{}
	Term2 interface{}
	Term3 interface{}
}

// NewConcealedIndexTerm returns a new ConcealedIndexTerm
func NewConcealedIndexTerm(term1, term2, term3 interface{}) (*ConcealedIndexTerm, error) {
	return &ConcealedIndexTerm{
		Term1: term1,
		Term2: term2,
		Term3: term3,
	}, nil
}

// ------------------------------------------------------------------------------------
// Special Characters
// They need to be identified as they may have a special treatment during the rendering
// ------------------------------------------------------------------------------------

// SpecialCharacter a special character, which may get a special treatment later during rendering
// Eg `<`, `>`, `&`
type SpecialCharacter struct {
	Name string
}

// NewSpecialCharacter return a new SpecialCharacter
func NewSpecialCharacter(name string) (*SpecialCharacter, error) {
	return &SpecialCharacter{
		Name: name,
	}, nil
}

// Symbol a sequence of characters, which may get a special treatment later during rendering
// Eg: `(C)`, `(TM)`, `...`, etc.
type Symbol struct {
	Prefix string // optional
	Name   string
}

// NewSymbol return a new Symbol
func NewSymbol(name string) (*Symbol, error) {
	return &Symbol{
		Name: name,
	}, nil
}

// NewSymbolWithForeword return a new Symbol prefixed with a foreword
func NewSymbolWithForeword(name, foreword string) (*Symbol, error) {
	return &Symbol{
		Name:   name,
		Prefix: foreword,
	}, nil
}

// ------------------------------------------------------------------------------------
// ElementPlaceHolder
// They need to be identified as they may have a special treatment during the rendering
// ------------------------------------------------------------------------------------

// ElementPlaceHolder a placeholder for elements which may have been parsed
// during previous substitution, and which are substituted with a placeholder while
// serializing the content to parse with the "macros" substitution
type ElementPlaceHolder struct {
	Ref string
}

// NewElementPlaceHolder returns a new ElementPlaceHolder with the given reference.
func NewElementPlaceHolder(ref string) (*ElementPlaceHolder, error) {
	return &ElementPlaceHolder{
		Ref: ref,
	}, nil
}

func (p *ElementPlaceHolder) String() string {
	return "\uFFFD" + p.Ref + "\uFFFD"
}

// ------------------------------------------
// Tables
// ------------------------------------------

// Table the structure for the tables
type Table struct {
	Attributes Attributes
	Header     *TableRow
	Footer     *TableRow
	Rows       []*TableRow
}

func NewTable(lines []interface{}) (*Table, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new table from %s", spew.Sdump(lines))
	}
	if len(lines) == 0 {
		return &Table{}, nil
	}
	header, rows := scanTableElements(lines)
	return &Table{
		Header: header,
		Rows:   rows,
	}, nil
}

func scanTableElements(rows []interface{}) (*TableRow, []*TableRow) {
	// check first 2 elements, expecting a row followed by a blankline as the optional header row
	if len(rows) > 2 {
		if header, ok := rows[0].(*TableRow); ok {
			if _, ok := rows[1].(*BlankLine); ok {
				rows := organizeTableCells(rows[2:], len(header.Cells))
				return header, rows
			}
		}
	}
	rowLength := 1
	for _, e := range rows {
		if r, ok := e.(*TableRow); ok {
			rowLength = len(r.Cells)
			break
		}
	}
	body := organizeTableCells(rows, rowLength)
	return nil, body
}

func organizeTableCells(elements []interface{}, rowLength int) []*TableRow {
	// add all cells in a single slice, then group by rows
	cells := make([]*TableCell, 0, len(elements))
	for _, e := range elements {
		if e, ok := e.(*TableRow); ok { // silently ignore 'BlankLines'
			cells = append(cells, e.Cells...)
		}
	}
	log.Debugf("dispatching %d cells in rows of %d cells", len(cells), rowLength)
	rows := make([]*TableRow, 0, int(len(cells)/rowLength)+1)
	for len(cells) > 0 {
		r := &TableRow{}
		rows = append(rows, r)
		l := int(math.Min(float64(len(cells)), float64(rowLength)))
		r.Cells = make([]*TableCell, l)
		copy(r.Cells, cells[:l])
		cells = cells[l:]
	}
	return rows
}

var _ WithElements = &Table{}

// return the optional header line and the cell lines
func (t *Table) GetElements() []interface{} {
	rows := make([]interface{}, len(t.Rows))
	for i, l := range t.Rows {
		rows[i] = l
	}
	return rows
}

func (t *Table) SetElements(elements []interface{}) error {
	if len(elements) == 0 {
		t.Rows = nil
		return nil
	}
	rows := make([]*TableRow, len(elements))
	for i, e := range elements {
		switch e := e.(type) {
		case *TableRow:
			rows[i] = e
		default:
			return errors.Errorf("unexpected type of table row: '%T'", e)
		}
	}
	t.Rows = rows
	return nil
}

var _ WithAttributes = &Table{}

func (t *Table) GetAttributes() Attributes {
	return t.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (t *Table) AddAttributes(attributes Attributes) {
	t.Attributes = t.Attributes.AddAll(attributes)
	t.reorganizeRows()
}

// SetAttributes sets the attributes in this element
func (t *Table) SetAttributes(attributes Attributes) {
	t.Attributes = attributes
	t.reorganizeRows()
}

func (t *Table) reorganizeRows() {
	// if `header` option, then make sure that the first row is the header
	if t.Header == nil && len(t.Rows) > 0 && t.Attributes.HasOption("header") {
		t.Header = t.Rows[0]
		t.Rows = t.Rows[1:]
	}
	// if `footer` option, then make sure that the last row is the header
	if t.Footer == nil && len(t.Rows) > 0 && t.Attributes.HasOption("footer") {
		t.Footer = t.Rows[len(t.Rows)-1]
		t.Rows = t.Rows[:len(t.Rows)-1]
	}
}

func (t *Table) SetColumnDefinitions(cols interface{}) error {
	switch cols := cols.(type) {
	case []interface{}:
		t.Attributes[AttrCols] = cols
		size := 0
		for _, c := range cols {
			if c, ok := c.(*TableColumn); ok {
				size += c.Multiplier
			}
		}
		log.Debugf("re-organizing table in rows of %d cells", size)
		// reorganize rows/columns

		rows, header, footer := t.rows()
		t.Rows = organizeTableCells(rows, size)
		// restore header and footer
		if header || t.Attributes.HasOption("header") {
			t.Header = t.Rows[0]
			t.Rows = t.Rows[1:]
		}
		if footer || t.Attributes.HasOption("footer") {
			t.Footer = t.Rows[len(t.Rows)-1]
			t.Rows = t.Rows[:len(t.Rows)-1]
		}
		return nil
	default:
		return fmt.Errorf("unexpected type of column definitions: '%T'", cols)
	}
}

func (t *Table) rows() ([]interface{}, bool, bool) {
	rows := make([]interface{}, 0, len(t.Rows)+2)
	var header, footer bool
	if t.Header != nil {
		header = true
		rows = append(rows, t.Header)
	}
	for _, r := range t.Rows {
		rows = append(rows, r)
	}
	if t.Footer != nil {
		footer = true
		rows = append(rows, t.Footer)
	}
	return rows, header, footer
}

var _ Referencable = &Table{}

func (t *Table) Reference(refs ElementReferences) {
	id := t.Attributes.GetAsStringWithDefault(AttrID, "")
	title := t.Attributes[AttrTitle]
	if id != "" && title != nil {
		refs[id] = title
	}
}

type HAlign string

const (
	HAlignLeft   HAlign = "<"
	HAlignRight  HAlign = ">"
	HAlignCenter HAlign = "^"
)

type VAlign string

const (
	VAlignTop    VAlign = "<"
	VAlignBottom VAlign = ">"
	VAlignMiddle VAlign = "^"
)

type ContentStyle string

const (
	AsciidocStyle  ContentStyle = "a"
	DefaultStyle   ContentStyle = "d"
	EmphasisStyle  ContentStyle = "e"
	HeaderStyle    ContentStyle = "h"
	LiteralStyle   ContentStyle = "l"
	MonospaceStyle ContentStyle = "m"
	StrongStyle    ContentStyle = "s"
)

type TableColumn struct {
	Multiplier int
	HAlign     HAlign
	VAlign     VAlign
	Weight     int
	Width      string // computed value
	Style      ContentStyle
	Autowidth  bool
}

func NewTableColumn(multiplier, halign, valign, weight, style interface{}) (*TableColumn, error) {
	col := newDefaultTableColumn()
	if multiplier, ok := multiplier.(int); ok {
		col.Multiplier = multiplier
	}
	if halign, ok := halign.(HAlign); ok {
		col.HAlign = halign
	}
	if valign, ok := valign.(VAlign); ok {
		col.VAlign = valign
	}
	if weight == "~" {
		col.Autowidth = true
		col.Weight = 0
	} else if weight, ok := weight.(int); ok {
		col.Weight = weight
	}
	if style, ok := style.(string); ok {
		col.Style = ContentStyle(style)
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new TableColumnDef: multiplier=%v halign=%v valign=%v weight=%v", multiplier, halign, valign, weight)
		log.Debug(spew.Sdump(col))
	}
	return col, nil
}

func newDefaultTableColumn() *TableColumn {
	return &TableColumn{
		Multiplier: 1,
		HAlign:     HAlignLeft,
		VAlign:     VAlignTop,
		Weight:     1,
	}
}

func (t *Table) Columns() ([]*TableColumn, error) {
	result := []*TableColumn{}
	if cols, ok := t.Attributes[AttrCols].([]interface{}); ok {
		for _, col := range cols {
			switch col := col.(type) {
			case *TableColumn:
				for i := 0; i < col.Multiplier; i++ {
					// affect in new variable so we have a *copy*!
					c := *col
					result = append(result, &c)
				}
			default:
				return nil, fmt.Errorf("invalid type of column definition: '%T'", col)
			}

		}
	}
	// add empty entries if the first row has more cells than the number of column than what's specified in the `cols` attribute
	// (also works when the `cols` attribute is missing/empty)
	if len(t.Rows) > 0 && len(t.Rows[0].Cells) > len(result) {
		m := len(t.Rows[0].Cells) - len(result)
		for i := 0; i < m; i++ {
			col := newDefaultTableColumn()
			result = append(result, col)
		}
	}
	// unless table is set with "full autowidth"
	if !t.Attributes.HasOption(AttrAutoWidth) {
		sumWeight := 0
		colsWithAutowidth := false
		for _, col := range result {
			sumWeight += col.Weight
			colsWithAutowidth = colsWithAutowidth || col.Autowidth
		}
		if colsWithAutowidth {
			// some cols have `autowidth` (`~`) enabled, so the `weight` becomes the `width`
			// TODO: check that sum of `width < 100`
			for _, col := range result {
				if !col.Autowidth {
					col.Width = strconv.Itoa(col.Weight)
				}
			}
		} else {
			// now, compute the Width for each column, based on each one's relative weight
			sumWidth := 0 // sum by a factor 1e5 to retain precision
			for i, col := range result {
				if i < len(result)-1 {
					width := float64(col.Weight*100) / float64(sumWeight)
					col.Width = strconv.FormatFloat(width, 'g', 6, 64)
					sumWidth += int(width * 1e4)
					log.Debugf("col %d width: %s", i, col.Width)
				} else {
					// rounding on the last column, to make sure that the sum reaches 100
					width := (float64(1e6-sumWidth) / 1e4)
					col.Width = strconv.FormatFloat(width, 'g', 6, 64)
					log.Debugf("col %d width (last): %s", i, col.Width)
				}
			}
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("cols: %s", spew.Sdump(result))
	}
	return result, nil
}

// TableRow a table line is made of columns, each column being a group of []interface{} (to support quoted text, etc.)
type TableRow struct {
	Cells []*TableCell
}

func NewTableRow(elements []interface{}) (*TableRow, error) {
	cells := make([]*TableCell, len(elements))
	for i, e := range elements {
		switch e := e.(type) {
		case *TableCell:
			cells[i] = e
		default:
			return nil, fmt.Errorf("unexpected type of table cell: '%T'", e)
		}
	}
	return &TableRow{
		Cells: cells,
	}, nil
}

var _ WithElements = &TableRow{}

func (r *TableRow) GetAttributes() Attributes {
	return nil
}

// AddAttributes adds the attributes of this CalloutListElement
func (r *TableRow) AddAttributes(_ Attributes) {
	// e.Attributes = e.Attributes.AddAll(attributes)
}

func (r *TableRow) SetAttributes(_ Attributes) {
}

func (r *TableRow) GetElements() []interface{} {
	elements := make([]interface{}, len(r.Cells))
	for i, c := range r.Cells {
		elements[i] = c
	}
	return elements
}

func (r *TableRow) SetElements(elements []interface{}) error {
	cells := make([]*TableCell, len(elements))
	for i, e := range elements {
		c, ok := e.(*TableCell)
		if !ok {
			return errors.Errorf("unexpected type of cell: '%T'", e)
		}
		cells[i] = c
	}
	r.Cells = cells
	return nil
}

type TableCell struct {
	Format   string
	Elements []interface{}
}

func NewInlineTableCell(content RawLine) (*TableCell, error) {
	return &TableCell{
		Elements: []interface{}{
			content,
		},
	}, nil
}

func NewMultilineTableCell(elements []interface{}, format interface{}) (*TableCell, error) {
	for i, l := range elements {
		if l, ok := l.(RawLine); ok {
			// add `\n` unless the we're on the last element
			if i < len(elements)-1 {
				elements[i] = RawLine(l + "\n")
			}
		}
	}
	c := &TableCell{
		Elements: elements,
	}
	if format, ok := format.(string); ok {
		c.Format = format
	}
	return c, nil
}

var _ WithElements = &TableCell{}

func (c *TableCell) GetAttributes() Attributes {
	return nil
}

func (c *TableCell) AddAttributes(_ Attributes) {
}

func (c *TableCell) SetAttributes(_ Attributes) {
}

func (c *TableCell) GetElements() []interface{} {
	return c.Elements
}

func (c *TableCell) SetElements(elements []interface{}) error {
	c.Elements = elements
	return nil
}

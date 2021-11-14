package types

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// // RawDocument document with a front-matter and raw blocks (will be refined in subsequent processing phases)
// type RawDocument struct {
// 	FrontMatter FrontMatter
// 	Elements    []interface{}
// }

// // NewRawDocument initializes a new `RawDocument` from the given lines
// func NewRawDocument(frontMatter interface{}, elements []interface{}) (RawDocument, error) {
// 	// log.Debugf("new RawDocument with %d block element(s)", len(elements))
// 	result := RawDocument{
// 		Elements: elements,
// 	}
// 	if fm, ok := frontMatter.(FrontMatter); ok {
// 		result.FrontMatter = fm
// 	}
// 	return result, nil
// }

// Attributes returns the document attributes on the top-level section
// and all the document attribute declarations at the top of the document only.
// func (d RawDocument) Attributes() Attributes {
// 	result := Attributes{}
// elements:
// 	for _, b := range d.Elements {
// 		switch b := b.(type) {
// 		case Section:
// 			if b.Level == 0 {
// 				// also, expand document authors and revision
// 				if authors, ok := b.Attributes[AttrAuthors].([]DocumentAuthor); ok {
// 					// move to the Document attributes
// 					result.SetAll(expandAuthors(authors))
// 					delete(b.Attributes, AttrAuthors)
// 				}
// 				// also, expand document authors and revision
// 				if revision, ok := b.Attributes[AttrRevision].(DocumentRevision); ok {
// 					// move to the Document attributes
// 					result.SetAll(expandRevision(revision))
// 					delete(b.Attributes, AttrRevision)
// 				}
// 				continue // allow to continue if the section is level 0
// 			}
// 			break elements // otherwise, just stop
// 		case AttributeDeclaration:
// 			result.Set(b.Name, b.Value)
// 		default:
// 			break elements
// 		}
// 	}
// 	// log.Debugf("document attributes: %+v", result)
// 	return result
// }

// RawSection a document section (when processing file inclusions)
// We only care about the level here
// type RawSection struct {
// 	Level   int
// 	Title   string
// 	RawText string
// }

// NewRawSection returns a new RawSection
func NewRawSection(level int, title []interface{}) (*Section, error) {
	// log.Debugf("new rawsection: '%s' (%d)", title, level)
	return &Section{
		Level: level,
		Title: title,
	}, nil
}

// var _ Stringer = RawSection{}

// // Stringify returns the string representation of this section, as it existed in the source document
// func (s RawSection) Stringify() string {
// 	return strings.Repeat("=", s.Level+1) + " " + s.Title
// }

// ------------------------------------------
// common interfaces
// ------------------------------------------

// Stringer a type which can be serializes as a string
type Stringer interface {
	Stringify() string
}

// // WithPlaceholdersInElements interface for all blocks in which elements can
// // be substituted with placeholders while applying the substitutions
// type WithPlaceholdersInElements interface {
// 	RestoreElements(placeholders map[string]interface{}) interface{}
// }

// // WithPlaceholdersInAttributes interface for all blocks in which attribute content can
// // be substituted with placeholders while applying the substitutions
// type WithPlaceholdersInAttributes interface {
// 	RestoreAttributes(placeholders map[string]interface{}) interface{}
// }

// // WithPlaceholdersInLocation interface for all blocks in which location elements can
// // be substituted with placeholders while applying the substitutions
// type WithPlaceholdersInLocation interface {
// 	RestoreLocation(placeholders map[string]interface{}) interface{}
// }

// RawText interface for the elements that can provide the raw text representation of this element
// as it was (supposedly) written in the source document
type RawText interface {
	RawText() (string, error)
}

// BlockWithAttributes base interface for types on which attributes can be substituted
type BlockWithAttributes interface {
	GetAttributes() Attributes
	AddAttributes(Attributes)
	SetAttributes(Attributes)
}

type WithElementAddition interface {
	// WithElements
	CanAddElement(interface{}) bool
	AddElement(interface{}) error
}
type BlockWithElements interface {
	BlockWithAttributes
	GetElements() []interface{}
	SetElements([]interface{}) error
	// WithElementAddition
}

type BlockWithLocation interface {
	BlockWithAttributes
	GetLocation() *Location
	SetLocation(*Location) // TODO: unused?
}

// ------------------------------------------
// Substitution support
// ------------------------------------------

// DocumentFragment a single fragment of document
type DocumentFragment struct {
	Elements []interface{}
	Error    error
}

func NewDocumentFragment(elements ...interface{}) DocumentFragment {
	return DocumentFragment{
		Elements: elements,
	}
}

func NewErrorFragment(err error) DocumentFragment {
	return DocumentFragment{
		Error: err,
	}
}

// type RawBlock interface {
// 	AddLine(l RawLine)
// }

// type RawParagraph struct {
// 	Attributes Attributes
// 	Lines      []interface{}
// }

// func NewRawParagraph(attributes Attributes) *RawParagraph {
// 	return &RawParagraph{
// 		Attributes: attributes,
// 		Lines:      []interface{}{},
// 	}
// }

// var _ RawBlock = &RawParagraph{}

// func (p *RawParagraph) AddLine(l RawLine) {
// 	p.Lines = append(p.Lines, l)
// }

// type RawDelimitedBlock struct {
// 	Attributes Attributes
// 	Kind       string
// 	Lines      []RawLine
// }

// func NewRawDelimitedBlock(kind string, attributes Attributes) *RawDelimitedBlock {
// 	return &RawDelimitedBlock{
// 		Attributes: attributes,
// 		Kind:       kind,
// 		Lines:      []RawLine{},
// 	}
// }

// // var _ RawBlock = &RawDelimitedBlock{}

// func (b *RawDelimitedBlock) AddLine(l RawLine) {
// 	b.Lines = append(b.Lines, l)
// }

// ------------------------------------------
// Draft Document: document in which
// all substitutions have been applied
// DEPRECATED
// ------------------------------------------

// DraftDocument the linear-level structure for a document
type DraftDocument struct {
	Attributes  Attributes
	FrontMatter FrontMatter
	Elements    []interface{}
}

// ------------------------------------------
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	// Attributes        Attributes    // DEPRECATED
	Elements          []interface{} // TODO: rename to `Blocks`?
	ElementReferences ElementReferences
	Footnotes         []*Footnote
	TableOfContents   *TableOfContents
}

// Header returns the header, i.e., the section with level 0 if it found as the first element of the document
// For manpage documents, this also includes the first section (`Name` along with its first paragraph)
func (d *Document) Header() (*DocumentHeader, []interface{}, bool) { // TODO: remove `bool` return value? (consistency with other funcs)
	if len(d.Elements) == 0 {
		log.Debug("no header for empty doc")
		return nil, nil, false
	}
	elements := make([]interface{}, 0, len(d.Elements))
	for i, e := range d.Elements {
		if h, ok := e.(*DocumentHeader); ok {
			elements = append(elements, d.Elements[i+1:]...)
			return h, elements, true
		}
		elements = append(elements, e)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("no header in document: %T", d.Elements[0])
	}
	return nil, elements, false
}

var _ WithElementAddition = &Document{}

func (d *Document) CanAddElement(element interface{}) bool {
	return true
}

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
	TableOfContents TableOfContents
	Authors         []*DocumentAuthor
	Revision        DocumentRevision
}

func NewTableOfContents(maxDepth int) *TableOfContents {
	return &TableOfContents{
		maxDepth: maxDepth,
	}
}

// TableOfContents the table of contents
type TableOfContents struct {
	maxDepth int
	Sections []*ToCSection
}

// ToCSection a section in the table of contents
type ToCSection struct {
	ID       string
	Level    int
	Title    string // the title as it was rendered in HTML
	Children []*ToCSection
}

// Add adds a ToCSection associated with the given Section
func (t *TableOfContents) Add(s *Section) {
	if s.Level > t.maxDepth {
		log.Debugf("skipping section with level %d (> %d)", s.Level, t.maxDepth)
		// skip for elements with a too low level in the hierarchy
		return
	}
	ts := &ToCSection{
		ID:    s.GetAttributes().GetAsStringWithDefault(AttrID, ""),
		Level: s.Level,
		Title: stringify(s.Title),
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

func NewDocumentHeader(title []interface{}, info interface{}, extraAttrs []interface{}) (*DocumentHeader, error) {
	header := &DocumentHeader{
		Title: title,
	}
	elements := make([]interface{}, 0, 2+len(extraAttrs)) // estimated max capacity
	if info, ok := info.(*DocumentInformation); ok {
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
	for _, attr := range extraAttrs {
		switch attr := attr.(type) {
		case *AttributeDeclaration, *AttributeReset:
			elements = append(elements, attr)
		default:
			return nil, fmt.Errorf("unexpected type of attribute declaration in the document header: '%T'", attr)
		}
	}
	if len(elements) > 0 {
		header.Elements = elements
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

var _ BlockWithAttributes = &DocumentHeader{}

func (h *DocumentHeader) GetAttributes() Attributes {
	return h.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (h *DocumentHeader) AddAttributes(attributes Attributes) {
	h.Attributes = h.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this section
func (h *DocumentHeader) SetAttributes(attributes Attributes) {
	h.Attributes = attributes
	if _, exists := h.Attributes[AttrID]; exists {
		// needed to track custom ID during rendering
		h.Attributes[AttrCustomID] = true
	}
}

type DocumentInformation struct {
	Authors  DocumentAuthors
	Revision *DocumentRevision
}

func NewDocumentInformation(authors DocumentAuthors, revision interface{}) (*DocumentInformation, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("new document info")
		log.Debugf("authors: %s", spew.Sdump(authors))
		log.Debugf("revision: %s", spew.Sdump(revision))
	}
	info := &DocumentInformation{
		Authors: authors,
	}
	if revision, ok := revision.(*DocumentRevision); ok {
		info.Revision = revision
	}
	return info, nil
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
	// result = append(result, NewAttributeDeclaration(AttrAuthors, authors))
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
		// then we need to strip the heading ":" and spaces
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
	Name  string
	Value interface{}
}

// NewAttributeDeclaration initializes a new AttributeDeclaration with the given name and optional value
func NewAttributeDeclaration(name string, value interface{}) *AttributeDeclaration {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	// log.Debugf("new AttributeDeclaration: '%s'", name)
	// 	spew.Fdump(log.StandardLogger().Out, value)
	// }
	return &AttributeDeclaration{
		Name:  name,
		Value: value,
	}
}

var _ Stringer = &AttributeDeclaration{}

// Stringify returns the string representation of this attribute declaration, as it existed in the source document
func (a *AttributeDeclaration) Stringify() string {
	result := strings.Builder{}
	result.WriteString(":" + a.Name + ":")
	result.WriteString(stringify(a.Value))
	return result.String()
}

// ReplaceElements replaces the elements in this section
func (a *AttributeDeclaration) ReplaceElements(value []interface{}) interface{} {
	a.Value = Reduce(value, strings.TrimSpace)
	return a
}

// AttributeReset the type for AttributeReset
type AttributeReset struct {
	Name string
}

// NewAttributeReset initializes a new Document Attribute Resets.
func NewAttributeReset(attrName string) (*AttributeReset, error) {
	log.Debugf("new AttributeReset: '%s'", attrName)
	return &AttributeReset{Name: attrName}, nil
}

// AttributeSubstitution the type for AttributeSubstitution
type AttributeSubstitution struct {
	Name string
}

// NewAttributeSubstitution initializes a new Attribute Substitutions
func NewAttributeSubstitution(name string) (interface{}, error) {
	if isPrefedinedAttribute(name) {
		return &PredefinedAttribute{Name: name}, nil
	}
	// log.Debugf("new AttributeSubstitution: '%s'", name)
	return &AttributeSubstitution{Name: name}, nil
}

var _ RawText = &AttributeSubstitution{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (s *AttributeSubstitution) RawText() (string, error) {
	return "{" + s.Name + "}", nil
}

// PredefinedAttribute a special kind of attribute substitution, which
// uses a predefined attribute
type PredefinedAttribute AttributeSubstitution

// CounterSubstitution is a counter, that may increment when it is substituted.
// If Increment is set, then it will increment before being expanded.
type CounterSubstitution struct {
	Name   string
	Hidden bool
	Value  interface{} // may be a byte for character
}

// NewCounterSubstitution returns a counter substitution.
func NewCounterSubstitution(name string, hidden bool, val interface{}) (CounterSubstitution, error) {
	if v, ok := val.(string); ok {
		val = rune(v[0])
	}
	return CounterSubstitution{
		Name:   name,
		Hidden: hidden,
		Value:  val,
	}, nil
}

// StandaloneAttributes are attributes at the end of
// a delimited block or at the end of the doc, ie, not
// associated with any block. They shall be ignored/discarded
// in the final document
type StandaloneAttributes Attributes

// NewStandaloneAttributes returns a new StandaloneAttributes element
func NewStandaloneAttributes(attributes interface{}) (StandaloneAttributes, error) {
	log.Debug("new standalone attributes")
	return StandaloneAttributes(toAttributes(attributes)), nil
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
	err := yaml.Unmarshal([]byte(content), &attributes)
	if err != nil {
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
	BlockWithElements
	WithElementAddition
	LastElement() interface{}
	ListKind() ListKind
	AdjustStyle(*List)
	matchesStyle(ListElement) bool
}

func canAddToListElement(element interface{}) bool {
	switch element.(type) {
	case RawLine, *SingleLineComment:
		return true
	default:
		return false
	}
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
			return p.AddElement(element.trim())
		}
		return e.SetElements(append(e.GetElements(), &Paragraph{
			Elements: []interface{}{
				element.trim(),
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

var _ BlockWithElements = &List{}

// CanAddElement checks if the given element can be added
func (l *List) CanAddElement(element interface{}) bool {
	switch e := element.(type) {
	case ListElement:
		// any listelement can be added if there was no blankline before
		// otherwise, only accept list element with attribute if there is no blankline before
		return e.ListKind() == l.Kind && e.matchesStyle(l.LastElement()) // TODO: compare to `FirstElement` is enough and faster
	case *ListElementContinuation:
		return true
	// case *Paragraph, *DelimitedBlock:
	// 	lastList := l.lastLists[len(l.lastLists)-1]
	// 	_, ok := lastList.lastElement().LastElement().(*ListElementContinuation)
	// 	return ok
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
	case *ListElementContinuation:
		return l.LastElement().AddElement(e.Element)
	}

	return errors.Errorf("cannot add element of type '%T' to list of kind '%s'", element, l.Kind)
}

func (l *List) LastElement() ListElement {
	if len(l.Elements) == 0 {
		return nil
	}
	return l.Elements[len(l.Elements)-1]
}

// GetElements returns this paragraph's elements (or lines)
func (l *List) GetElements() []interface{} {
	elements := make([]interface{}, len(l.Elements))
	for i, e := range l.Elements {
		elements[i] = e
	}
	return elements
}

// SetElements sets this paragraph's elements
func (l *List) SetElements(elements []interface{}) error {
	// ensure that all elements are `ListElement`
	l.Elements = make([]ListElement, len(elements))
	for i, e := range elements {
		if e, ok := e.(ListElement); ok {
			l.Elements[i] = e
			continue
		}
		return fmt.Errorf("unexpected kind of element to set in a list: '%T'", e)
	}
	return nil
}

// GetAttributes returns this first item's attributes (if applicable)
func (l *List) GetAttributes() Attributes {
	if len(l.Elements) > 0 {
		if f, ok := l.Elements[0].(BlockWithAttributes); ok {
			return f.GetAttributes()
		}
	}
	return Attributes(nil)
}

// AddAttributes adds the attributes in this fist list item (if applicable)
func (l *List) AddAttributes(attributes Attributes) {
	if len(l.Elements) > 0 {
		if f, ok := l.Elements[0].(BlockWithAttributes); ok {
			f.AddAttributes(attributes)
		}
	}
}

// SetAttributes replaces the attributes in this fist list item (if applicable)
func (l *List) SetAttributes(attributes Attributes) {
	if len(l.Elements) > 0 {
		if f, ok := l.Elements[0].(BlockWithAttributes); ok {
			f.SetAttributes(attributes)
		}
	}
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
		case BlockWithAttributes:
			if attrs != nil {
				e.SetAttributes(attrs)
				attrs = nil
			}
			elmts = append(elmts, e)
		case RawLine, *SingleLineComment:
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

var _ BlockWithElements = &ListElements{}

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

// CanAddElement checks if the given element can be added
func (l *ListElements) CanAddElement(element interface{}) bool {
	return true
}

// AddElement adds the given element
func (l *ListElements) AddElement(element interface{}) error {
	l.Elements = append(l.Elements, element)
	return nil
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

type ListElementContinuation struct {
	Offset  int
	Element interface{}
}

func NewListElementContinuation(offset int, Element interface{}) (*ListElementContinuation, error) {
	return &ListElementContinuation{
		Offset:  offset,
		Element: Element,
	}, nil
}

var _ WithElementAddition = &ListElementContinuation{}

func (c *ListElementContinuation) CanAddElement(element interface{}) bool {
	if e, ok := c.Element.(WithElementAddition); ok {
		return e.CanAddElement(element)
	}
	return false
}

func (c *ListElementContinuation) AddElement(element interface{}) error {
	if e, ok := c.Element.(WithElementAddition); ok {
		return e.AddElement(element)
	}
	return errors.Errorf("cannot add element of type '%T' to list element continuation", c.Element)
}

// ContinuedListItemElement a wrapper for an element which should be attached to a list item (same level or an ancestor)
type ContinuedListItemElement struct {
	Offset  int // the relative ancestor. Should be a negative number
	Element interface{}
}

// NewContinuedListItemElement returns a wrapper for an element which should be attached to a list item (same level or an ancestor)
func NewContinuedListItemElement(element interface{}) (ContinuedListItemElement, error) {
	// log.Debugf("new continued list element for element of type %T", element)
	return ContinuedListItemElement{
		Offset:  0,
		Element: element,
	}, nil
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
func NewCalloutListElement(ref int, content RawLine) (*CalloutListElement, error) {
	return &CalloutListElement{
		Attributes: nil,
		Ref:        ref,
		Elements: []interface{}{
			&Paragraph{
				Elements: []interface{}{
					content,
				},
			},
		},
	}, nil
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

// GetAttributes returns the attributes of this CalloutListElement
func (e *CalloutListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (e *CalloutListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes of this CalloutListElement
func (e *CalloutListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
}

// CanAddElement checks if the given element can be added
func (e *CalloutListElement) CanAddElement(element interface{}) bool {
	return canAddToListElement(element)
}

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
	Style      string        // TODO: rename to `OrderedListElementNumberingStyle`? TODO: define as an attribute instead?
	Elements   []interface{} // TODO: rename to `Blocks`?
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

// CanAddElement checks if the given element can be added
func (e *OrderedListElement) CanAddElement(element interface{}) bool {
	return canAddToListElement(element)
}

// AddElement add an element to this UnorderedListElement
func (e *OrderedListElement) AddElement(element interface{}) error {
	return addToListElement(e, element)
}

var _ BlockWithAttributes = &OrderedListElement{}

// GetAttributes returns this list item's attributes
func (e *OrderedListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (e *OrderedListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// SetAttributes sets the attributes in this list item
func (e *OrderedListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *OrderedListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
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
	Elements    []interface{} // TODO: rename to `Blocks`?
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

// // GetLevel returns the level of this element
// func (e *UnorderedListElement) GetLevel() int {
// 	return e.Level
// }

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

// CanAddElement checks if the given element can be added
func (e *UnorderedListElement) CanAddElement(element interface{}) bool {
	return canAddToListElement(element)
}

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

var _ BlockWithAttributes = &UnorderedListElement{}

// GetAttributes returns this list item's attributes
func (e *UnorderedListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (e *UnorderedListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// SetAttributes replaces the attributes in this list item
func (e *UnorderedListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *UnorderedListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
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

// UnorderedListElementPrefix the prefix used to construct an UnorderedListElement
type UnorderedListElementPrefix struct {
	BulletStyle UnorderedListElementBulletStyle
}

// NewUnorderedListElementPrefix initializes a new UnorderedListElementPrefix
func NewUnorderedListElementPrefix(s UnorderedListElementBulletStyle) (UnorderedListElementPrefix, error) {
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
	Elements   []interface{} // TODO: rename to `Blocks`?
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

// CanAddElement checks if the given element can be added
func (e *LabeledListElement) CanAddElement(element interface{}) bool {
	return canAddToListElement(element)
}

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

var _ BlockWithAttributes = &LabeledListElement{}

// GetAttributes returns this list item's attributes
func (e *LabeledListElement) GetAttributes() Attributes {
	return e.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (e *LabeledListElement) AddAttributes(attributes Attributes) {
	e.Attributes = e.Attributes.AddAll(attributes)
	e.mapAttributes()
}

// ReplaceAttributes replaces the attributes in this list item
func (e *LabeledListElement) SetAttributes(attributes Attributes) {
	e.Attributes = attributes
	e.mapAttributes()
}

func (e *LabeledListElement) mapAttributes() {
	e.Attributes = toAttributesWithMapping(e.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraphs
type Paragraph struct {
	Attributes Attributes
	Elements   []interface{}
}

// AttrHardBreaks the attribute to set on a paragraph to render with hard breaks on each line
// TODO: remove?
const AttrHardBreaks = "hardbreaks"

// DocumentAttrHardBreaks the attribute to set at the document level to render with hard breaks on each line of all paragraphs
const DocumentAttrHardBreaks = "hardbreaks"

// NewParagraph initializes a new `Paragraph`
func NewParagraph(elements ...interface{}) (*Paragraph, error) {
	// log.Debugf("new paragraph with attributes: '%v'", attributes)
	return &Paragraph{
		Elements: elements,
	}, nil
}

func NewAdmonitionParagraph(kind string, elements []interface{}) (*Paragraph, error) {
	return &Paragraph{
		Attributes: Attributes{
			AttrStyle: kind,
		},
		Elements: elements,
	}, nil
}

func NewLiteralParagraph(kind string, elements []interface{}) (*Paragraph, error) {
	return &Paragraph{
		Attributes: Attributes{
			AttrStyle:            Literal,
			AttrLiteralBlockType: kind,
		},
		Elements: elements,
	}, nil
}

var _ BlockWithElements = &Paragraph{}

// GetElements returns this paragraph's elements (or lines)
func (p *Paragraph) GetElements() []interface{} {
	return p.Elements
}

// SetElements sets this paragraph's elements
func (p *Paragraph) SetElements(elements []interface{}) error {
	p.Elements = elements
	return nil
}

// CanAddElement checks if the given element can be added
func (p *Paragraph) CanAddElement(element interface{}) bool {
	switch element.(type) {
	case RawLine, *SingleLineComment:
		return true
	default:
		return false
	}
}

func (p *Paragraph) AddElement(e interface{}) error {
	p.Elements = append(p.Elements, e)
	return nil
}

var _ BlockWithAttributes = &Paragraph{}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (p *Paragraph) GetAttributes() Attributes {
	return p.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (p *Paragraph) AddAttributes(attributes Attributes) {
	p.Attributes = p.Attributes.AddAll(attributes)
	p.mapAttributes()
}

// ReplaceAttributes replaces the attributes in this paragraph
func (p *Paragraph) SetAttributes(attributes Attributes) {
	p.Attributes = attributes
	p.mapAttributes()
}

func (p *Paragraph) mapAttributes() {
	// add an extra `literalBlockType: literalBlockWithAttribute` attribute
	// before mapping, so we know that the `style=Literal` came from
	// block attributes
	if p.Attributes.GetAsStringWithDefault(AttrPositional1, "") == Literal {
		p.Attributes.Set(AttrLiteralBlockType, LiteralBlockWithAttribute)
	}
	p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	switch p.Attributes.GetAsStringWithDefault(AttrStyle, "") {
	case string(Source):
		p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
			AttrPositional2: AttrLanguage,
		})
	case string(Quote), string(Verse):
		p.Attributes = toAttributesWithMapping(p.Attributes, map[string]string{
			AttrPositional2: AttrQuoteAuthor,
			AttrPositional3: AttrQuoteTitle,
		})
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

const (
	// Tip the 'TIP' type of admonition
	Tip = "TIP"
	// Note the 'NOTE' type of admonition
	Note = "NOTE"
	// Important the 'IMPORTANT' type of admonition
	Important = "IMPORTANT"
	// Warning the 'WARNING' type of admonition
	Warning = "WARNING"
	// Caution the 'CAUTION' type of admonition
	Caution = "CAUTION"
	// Unknown is the zero value for admonition kind
	Unknown = ""
)

type AdmonitionLine struct {
	Kind    string
	Content RawLine
}

// NewAdmonitionLine returns a new AdmonitionLine with the given kind and content
func NewAdmonitionLine(kind string, content string) (*AdmonitionLine, error) {
	log.Debugf("new admonition paragraph")
	return &AdmonitionLine{
		Kind:    kind,
		Content: RawLine(content),
	}, nil
}

// ------------------------------------------
// Inline Elements
// ------------------------------------------

type InlineElements []interface{} // TODO: unnecessary alias?

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements ...interface{}) (InlineElements, error) {
	return Merge(elements...), nil
}

// // HasAttributeSubstitutions returns `true` if at least one of the element is an `AttributeSubstitution`
// func (e InlineElements) HasAttributeSubstitutions() bool {
// 	for _, elmt := range e {
// 		if _, match := elmt.(*AttributeSubstitution); match {
// 			return true
// 		}
// 	}
// 	return false
// }

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

// var _ WithPlaceholdersInElements = ExternalCrossReference{}

var _ BlockWithLocation = &ExternalCrossReference{}

func (x *ExternalCrossReference) GetLocation() *Location {
	return x.Location
}

func (x *ExternalCrossReference) SetLocation(l *Location) {
	x.Location = l
}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (x *ExternalCrossReference) GetAttributes() Attributes {
	return x.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (x *ExternalCrossReference) AddAttributes(attributes Attributes) {
	x.Attributes = x.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this paragraph
func (x *ExternalCrossReference) SetAttributes(attributes Attributes) {
	x.Attributes = x.Attributes.SetAll(attributes)
}

// var _ WithElementsToSubstitute = ExternalCrossReference{}

// // ElementsToSubstitute returns this corss reference location path so that substitutions can be applied onto it
// func (r ExternalCrossReference) ElementsToSubstitute() []interface{} {
// 	return r.Location.Path
// }

// // ReplaceElements replaces the elements in this example block
// func (r ExternalCrossReference) ReplaceElements(path []interface{}) interface{} {
// 	r.Location.Path = path
// 	return r
// }

// ------------------------------------------
// Images
// ------------------------------------------

// ImageBlock the structure for the block images
type ImageBlock struct {
	Location   *Location
	Attributes Attributes
}

// NewImageBlock initializes a new `ImageBlock`
func NewImageBlock(location *Location, inlineAttributes Attributes, attributes interface{}) (*ImageBlock, error) {
	// inline attributes trump block attributes
	attrs := toAttributes(inlineAttributes)
	attrs.SetAll(attributes)
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

// var _ WithPlaceholdersInAttributes = &ImageBlock{}

// // RestoreAttributes restores the attributes which had been substituted by placeholders
// func (i *ImageBlock) RestoreAttributes(placeholders map[string]interface{}) interface{} {
// 	i.Attributes = restoreAttributes(i.Attributes, placeholders)
// 	return i
// }

// var _ WithPlaceholdersInLocation = &ImageBlock{}

// // RestoreLocation restores the location elements which had been substituted by placeholders
// func (i ImageBlock) RestoreLocation(placeholders map[string]interface{}) interface{} {
// 	i.Location.Path = restoreElements(i.Location.Path, placeholders)
// 	return i
// }

var _ BlockWithAttributes = &ImageBlock{}

// GetAttributes returns this list item's attributes
func (i *ImageBlock) GetAttributes() Attributes {
	return i.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (i *ImageBlock) AddAttributes(attributes Attributes) {
	i.Attributes = i.Attributes.AddAll(attributes)
}

// SetAttributes sets the attributes in this image
func (i *ImageBlock) SetAttributes(attributes Attributes) {
	i.Attributes = attributes
}

var _ BlockWithLocation = &ImageBlock{}

func (i *ImageBlock) GetLocation() *Location {
	return i.Location
}

func (i *ImageBlock) SetLocation(value *Location) {
	i.Location = value
}

// var _ WithElementsToSubstitute = ImageBlock{}

// // ElementsToSubstitute returns this image's location path so that substitutions can be applied onto it
// func (i ImageBlock) ElementsToSubstitute() []interface{} {
// 	return i.Location.Path
// }

// // ReplaceElements replaces the elements in this example block
// func (i ImageBlock) ReplaceElements(path []interface{}) interface{} {
// 	i.Location.Path = path
// 	return i
// }

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Location   *Location
	Attributes Attributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(location *Location, attributes interface{}, imagesdir interface{}) (*InlineImage, error) {
	location.SetPathPrefix(imagesdir)
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

// var _ WithPlaceholdersInAttributes = &InlineImage{}

// // RestoreAttributes restores the attributes which had been substituted by placeholders
// func (i *InlineImage) RestoreAttributes(placeholders map[string]interface{}) interface{} {
// 	i.Attributes = restoreAttributes(i.Attributes, placeholders)
// 	return i
// }

// var _ WithPlaceholdersInLocation = &InlineImage{}

// // RestoreLocation restores the location elements which had been substituted by placeholders
// func (i *InlineImage) RestoreLocation(placeholders map[string]interface{}) interface{} {
// 	i.Location.Path = restoreElements(i.Location.Path, placeholders)
// 	return i
// }

var _ BlockWithAttributes = &InlineImage{}

// GetAttributes returns this inline image's attributes
func (i *InlineImage) GetAttributes() Attributes {
	return i.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (i *InlineImage) AddAttributes(attributes Attributes) {
	i.Attributes = i.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this inline image
func (i *InlineImage) SetAttributes(attributes Attributes) {
	i.Attributes = attributes
}

var _ BlockWithLocation = &InlineImage{}

func (i *InlineImage) GetLocation() *Location {
	return i.Location
}

func (i *InlineImage) SetLocation(value *Location) {
	i.Location = value
}

// var _ WithElementsToSubstitute = InlineImage{}

// // ElementsToSubstitute returns this inline image location path so that substitutions can be applied onto its elements
// func (i InlineImage) ElementsToSubstitute() []interface{} {
// 	return i.Location.Path // TODO: should return the location so substitution can also take place on the scheme
// }

// // ReplaceElements replaces the elements in this inline image
// func (i InlineImage) ReplaceElements(path []interface{}) interface{} {
// 	i.Location.Path = path
// 	return i
// }

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
func NewFootnote(ref string, elements interface{}) (*Footnote, error) {
	log.Debugf("new footnote with elements: '%s'", spew.Sdump(elements))
	// footnote with content get an ID
	if elements, ok := elements.(InlineElements); ok {
		return &Footnote{
			// ID is only set during document processing
			Ref:      ref,
			Elements: elements,
		}, nil
	} // footnote which are just references don't get an ID, so we don't increment the sequence
	return &Footnote{
		Ref:      ref,
		Elements: []interface{}{},
	}, nil
}

// var _ WithPlaceholdersInElements = Footnote{}

// // RestoreElements restores the elements which had been substituted by placeholders
// func (n Footnote) RestoreElements(placeholders map[string]interface{}) interface{} {
// 	n.Elements = restoreElements(n.Elements, placeholders)
// 	return n
// }

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

type BlockDelimiter struct {
	Kind string
}

type BlockDelimiterKind string // TODO: use it

func NewBlockDelimiter(kind string) (*BlockDelimiter, error) {
	return &BlockDelimiter{
		Kind: kind,
	}, nil
}

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
	// MarkdownQuote a quote block in the Markdown style
	MarkdownQuote string = "markdown-quote"
	// Verse a verse block
	Verse string = "verse"
	// Sidebar a sidebar block
	Sidebar string = "sidebar"
	// Literal a literal block
	Literal string = "literal"
	// Source a source block
	Source string = "source"
	// Passthrough a passthrough block
	Passthrough string = "pass"

	// AttrSourceBlockOption the option set on a source block, using the `source%<option>` attribute
	AttrSourceBlockOption = "source-option" // DEPRECATED
)

// DelimitedBlock the structure for the Listing blocks
type DelimitedBlock struct {
	Kind       string
	Attributes Attributes
	Elements   []interface{}
}

func NewDelimitedBlock(kind string, elements []interface{}) (*DelimitedBlock, error) {
	return &DelimitedBlock{
		Kind:     kind,
		Elements: elements,
	}, nil
}

var _ BlockWithElements = &DelimitedBlock{}

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
			log.Debugf("discarding heading crlf on elements in block of kind '%s'", b.Kind)
			// discard heading spaces and CR/LF
			if s, ok := elements[0].(*StringElement); ok {
				s.Content = strings.TrimLeft(s.Content, "\r\n")
			}
		default:
			log.Debugf("discarding heading spaces+crlf on elements in block of kind '%s'", b.Kind)
			// discard heading spaces and CR/LF
			if s, ok := elements[0].(*StringElement); ok {
				s.Content = strings.TrimLeft(s.Content, " \t\r\n")
			}
		}
		// discard trailing spaces and CR/LF
		log.Debugf("discarding trailing spaces+crlf on elements in block of kind '%s'", b.Kind)
		if s, ok := elements[len(elements)-1].(*StringElement); ok {
			s.Content = strings.TrimRight(s.Content, " \t\r\n")
		}
	}
	b.Elements = elements
	return nil
}

// CanAddElement checks if the given element can be added
func (b *DelimitedBlock) CanAddElement(element interface{}) bool {
	switch element.(type) {
	case *BlockDelimiter, RawLine:
		return true
	default:
		switch b.Kind {
		// // Normal blocks can have more kinds of elements
		// case Example, Quote:
		// 	return true
		default:
			return false
		}
	}
}

func (b *DelimitedBlock) AddElement(element interface{}) error {
	log.Debugf("adding element of type '%T' to delimited block of kind '%s'", element, b.Kind)
	b.Elements = append(b.Elements, element)
	return nil
}

var _ BlockWithAttributes = &DelimitedBlock{}

// GetAttributes returns the attributes of this paragraph so that substitutions can be applied onto them
func (b *DelimitedBlock) GetAttributes() Attributes {
	return b.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (b *DelimitedBlock) AddAttributes(attributes Attributes) {
	b.Attributes = b.Attributes.AddAll(attributes)
	b.mapAttributes()
}

// ReplaceAttributes replaces the attributes in this paragraph
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

// TODO: not needed?
const (
	// AttrLiteralBlockType the type of literal block, ie, how it was parsed
	AttrLiteralBlockType = "literalBlockType"
	// LiteralBlockWithDelimiter a literal block parsed with a delimiter
	LiteralBlockWithDelimiter = "literalBlockWithDelimiter"
	// LiteralBlockWithSpacesOnFirstLine a literal block parsed with one or more spaces on the first line
	LiteralBlockWithSpacesOnFirstLine = "literalBlockWithSpacesOnFirstLine"
	// LiteralBlockWithAttribute a literal block parsed with a `[literal]` attribute`
	LiteralBlockWithAttribute = "literalBlockWithAttribute"
)

// ------------------------------------------
// Sections
// ------------------------------------------

// Section the structure for a section
type Section struct {
	Level      int
	Attributes Attributes
	Title      []interface{}
	Elements   []interface{}
}

// NewSection initializes a new `Section` from the given section title and elements
func NewSection(level int, title []interface{}, ids []interface{}) (*Section, error) {
	// attrs := toAttributes(attributes)
	// // multiple IDs can be defined (by mistake), but only the last one is used
	// attrs = attrs.SetAll(ids)
	// // also, set the `AttrCustomID` flag if an ID was set
	// if _, exists := attrs[AttrID]; exists {
	// 	attrs[AttrCustomID] = true
	// }
	// attrs := Attributes{}
	// set the default ID

	return &Section{
		Level: level,
		// Attributes: attrs,
		Title: title,
		// Elements: []interface{}{},
	}, nil
}

var _ BlockWithElements = &Section{}

// GetElements returns this section's title
func (s *Section) GetElements() []interface{} {
	return s.Title
}

// SetElements sets this section's title
func (s *Section) SetElements(title []interface{}) error {
	s.Title = title
	return nil
}

// GetAttributes returns this section's attributes
func (s *Section) GetAttributes() Attributes {
	return s.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (s *Section) AddAttributes(attributes Attributes) {
	s.Attributes = s.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this section
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
			refs[id] = s.Title
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
	// inline attribute
	if id, ok := s.Title[len(s.Title)-1].(*Attribute); ok {
		sectionID := stringify(id.Value)
		s.Attributes[AttrID] = sectionID
		s.Attributes[AttrCustomID] = true
		return sectionID, nil
	}
	log.Debugf("resolving section id")
	separator := attrs.GetAsStringWithDefault(AttrIDSeparator, DefaultIDSeparator)
	replacement, err := ReplaceNonAlphanumerics(s.Title, separator)
	if err != nil {
		return "", errors.Wrapf(err, "failed to generate default ID on Section element")
	}
	idPrefix := attrs.GetAsStringWithDefault(AttrIDPrefix, DefaultIDPrefix)
	id := idPrefix + replacement
	s.Attributes[AttrID] = id
	log.Debugf("updated section id to '%s'", s.Attributes[AttrID])
	return id, nil
}

// CanAddElement checks if the given element can be added
func (s *Section) CanAddElement(_ interface{}) bool {
	return true
}

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
func NewUserMacroBlock(name string, value string, attributes interface{}, raw string) (*UserMacro, error) {
	return &UserMacro{
		Name:       name,
		Kind:       BlockMacro,
		Value:      value,
		Attributes: toAttributes(attributes),
		RawText:    raw,
	}, nil
}

// NewInlineUserMacro returns an UserMacro
func NewInlineUserMacro(name, value string, attributes interface{}, raw string) (*UserMacro, error) {
	return &UserMacro{
		Name:       name,
		Kind:       InlineMacro,
		Value:      value,
		Attributes: toAttributes(attributes),
		RawText:    raw,
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
	// log.Debug("new BlankLine")
	return &BlankLine{}, nil
}

// ------------------------------------------
// Comments
// ------------------------------------------

// SingleLineComment a single line comment
type SingleLineComment struct {
	Content string
}

// NewSingleLineComment initializes a new single line content
func NewSingleLineComment(content string) (*SingleLineComment, error) {
	// log.Debugf("initializing a single line comment with content: '%s'", content)
	return &SingleLineComment{
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

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (s StringElement) RawText() (string, error) {
	return s.Content, nil
}

// ------------------------------------------
// VerbatimLine
// ------------------------------------------

// VerbatimLine the structure for verbatim line, ie, read "as-is" from a given text document.
//TODO: remove
type VerbatimLine struct {
	Elements []interface{}
	Callouts []Callout
}

// NewVerbatimLine initializes a new `VerbatimLine` from the given content
func NewVerbatimLine(elements []interface{}, callouts []interface{}) (VerbatimLine, error) {
	var cos []Callout
	for _, c := range callouts {
		cos = append(cos, c.(Callout))
	}
	return VerbatimLine{
		Elements: elements,
		Callouts: cos,
	}, nil
}

// IsEmpty return `true` if the line contains only whitespaces and tabs
func (s VerbatimLine) IsEmpty() bool {
	return len(s.Elements) == 0 // || emptyStringRE.MatchString(s.Content)
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
		Elements: Merge(elements),
	}, nil
}

var _ RawText = &QuotedText{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (t *QuotedText) RawText() (string, error) {
	result := strings.Builder{}
	result.WriteString(string(t.Kind)) // opening delimiter
	s, err := toRawText(t.Elements)
	if err != nil {
		return "", err
	}
	result.WriteString(s)
	result.WriteString(string(t.Kind)) // closing delimiter
	return result.String(), nil
}

func toRawText(elements []interface{}) (string, error) {
	result := strings.Builder{}
	for _, e := range elements {
		r, ok := e.(RawText)
		if !ok {
			return "", fmt.Errorf("element of type '%T' cannot be converted to string", e)
		}
		s, err := r.RawText()
		if err != nil {
			return "", err
		}
		result.WriteString(s)
	}
	return result.String(), nil
}

var _ BlockWithElements = &QuotedText{}

// GetElements returns this QuotedText's elements
func (t *QuotedText) GetElements() []interface{} {
	return t.Elements
}

// SetElements sets this QuotedText's elements
func (t *QuotedText) SetElements(elements []interface{}) error {
	t.Elements = elements
	return nil
}

// CanAddElement checks if the given element can be added
func (t *QuotedText) CanAddElement(_ interface{}) bool {
	return true
}

func (t *QuotedText) AddElement(e interface{}) error {
	t.Elements = append(t.Elements, e)
	return nil
}

var _ BlockWithAttributes = &QuotedText{}

// GetAttributes returns the attributes of this QuotedText
func (t *QuotedText) GetAttributes() Attributes {
	return t.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (t *QuotedText) AddAttributes(attributes Attributes) {
	t.Attributes = t.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this QuotedText
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
func NewEscapedQuotedText(backslashes string, punctuation string, content interface{}) ([]interface{}, error) {
	// log.Debugf("new escaped quoted text: %s %s %v", backslashes, punctuation, content)
	backslashesStr := Apply(backslashes,
		func(s string) string {
			// remove the number of back-slashes that match the length of the punctuation. Eg: `\*` or `\\**`, but keep extra back-slashes
			if len(s) > len(punctuation) {
				return s[len(punctuation):]
			}
			return ""
		})
	return []interface{}{
		&StringElement{
			Content: backslashesStr,
		},
		&StringElement{
			Content: punctuation,
		},
		content,
		&StringElement{
			Content: punctuation,
		},
	}, nil
}

// -------------------------------------------------------
// Quoted Strings
// -------------------------------------------------------

// QuotedStringKind indicates whether this is 'single' or "double" quoted.
type QuotedStringKind string

const (
	// SingleQuote means single quotes (')
	SingleQuote = QuotedStringKind("'")
	// DoubleQuote means double quotes (")
	DoubleQuote = QuotedStringKind("\"")
)

// QuotedString a quoted string
type QuotedString struct {
	Kind     QuotedStringKind
	Elements []interface{}
}

// NewQuotedString returns a new QuotedString
func NewQuotedString(kind QuotedStringKind, elements []interface{}) (*QuotedString, error) {
	return &QuotedString{Kind: kind, Elements: elements}, nil
}

var _ RawText = QuotedString{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (s QuotedString) RawText() (string, error) {
	result := strings.Builder{}
	result.WriteString("`")            // opening delimiter
	result.WriteString(string(s.Kind)) // opening delimiter
	e, err := toRawText(s.Elements)
	if err != nil {
		return "", err
	}
	result.WriteString(e)
	result.WriteString(string(s.Kind)) // closing delimiter
	result.WriteString("`")            // closing delimiter
	return result.String(), nil
}

// var _ WithPlaceholdersInElements = QuotedString{}

// // RestoreElements restores the elements which had been substituted by placeholders
// func (s QuotedString) RestoreElements(placeholders map[string]interface{}) interface{} {
// 	s.Elements = restoreElements(s.Elements, placeholders)
// 	return s
// }

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
		Elements: Merge(elements...),
	}, nil
}

var _ RawText = &InlinePassthrough{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (p *InlinePassthrough) RawText() (string, error) {
	result := strings.Builder{}
	switch p.Kind {
	case PassthroughMacro:
		result.WriteString("pass:[") // opening delimiter
	default:
		result.WriteString(string(p.Kind)) // opening delimiter
	}
	e, err := toRawText(p.Elements)
	if err != nil {
		return "", err
	}
	result.WriteString(e)
	switch p.Kind {
	case PassthroughMacro:
		result.WriteString("]") // closing delimiter
	default:
		result.WriteString(string(p.Kind)) // closing delimiter
	}
	return result.String(), nil
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
	return &InlineLink{
		Location:   url,
		Attributes: attrs,
	}, nil
}

var _ BlockWithAttributes = &InlineLink{}

// GetAttributes returns this link's attributes
func (l *InlineLink) GetAttributes() Attributes {
	return l.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (l *InlineLink) AddAttributes(attributes Attributes) {
	l.Attributes = l.Attributes.AddAll(attributes)
}

func (l *InlineLink) SetAttributes(attributes Attributes) {
	l.Attributes = attributes
}

var _ BlockWithLocation = &InlineLink{}

func (l *InlineLink) GetLocation() *Location {
	return l.Location
}

func (l *InlineLink) SetLocation(value *Location) {
	l.Location = value
}

// NewInlineLinkAttributes returns a map of link attributes
func NewInlineLinkAttributes(attributes []interface{}) (Attributes, error) {
	// log.Debugf("new inline link attributes: %v", attributes)
	if len(attributes) == 0 {
		return nil, nil
	}
	result := Attributes{}
	for i, attr := range attributes {
		// log.Debugf("new inline link attribute: '%[1]v' (%[1]T)", attr)
		switch attr := attr.(type) {
		case *Attribute:
			result[attr.Key] = attr.Value
		case Attributes:
			for k, v := range attr {
				result[k] = v
			}
		case []interface{}:
			result["positional-"+strconv.Itoa(i+1)] = attr
		}
	}
	// log.Debugf("new inline link attributes: %v", result)
	return result, nil
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
func NewFileInclusion(location *Location, attributes interface{}, rawtext string) (*FileInclusion, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		"tag": "tags", // convert `tag` to `tags`
	})
	return &FileInclusion{
		Attributes: attrs,
		Location:   location,
		RawText:    rawtext,
	}, nil
}

var _ BlockWithLocation = &FileInclusion{}

func (f *FileInclusion) GetLocation() *Location {
	return f.Location
}

func (f *FileInclusion) SetLocation(value *Location) {
	f.Location = value
}

// GetAttributes returns this elements's attributes
func (f *FileInclusion) GetAttributes() Attributes {
	return f.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (f *FileInclusion) AddAttributes(attributes Attributes) {
	f.Attributes = f.Attributes.AddAll(attributes)
}

// ReplaceAttributes replaces the attributes in this element
func (f *FileInclusion) SetAttributes(attributes Attributes) {
	f.Attributes = attributes
}

// LineRanges returns the line ranges of the file to include.
func (f *FileInclusion) LineRanges() (LineRanges, bool) {
	if lr, ok := f.Attributes[AttrLineRanges].(LineRanges); ok {
		return lr, true
	}
	return LineRanges{ // default line ranges: include all content
		{
			StartLine: 1,
			EndLine:   -1,
		},
	}, false
}

// TagRanges returns the tag ranges of the file to include.
func (f *FileInclusion) TagRanges() (TagRanges, bool) {
	if lr, ok := f.Attributes[AttrTagRanges].(TagRanges); ok {
		return lr, true
	}
	return TagRanges{}, false // default tag ranges: include all content
}

// -------------------------------------------------------------------------------------
// Raw Line
// -------------------------------------------------------------------------------------
type RawLine string

// NewRawLine returns a new RawLine wrapper for the given string
func NewRawLine(content string) (RawLine, error) {
	// log.Debugf("new line: '%v'", content)
	return RawLine(strings.TrimRight(content, " \t")), nil
	// return RawLine(strings.Trim(content, " \t")), nil
}

func (l RawLine) trim() RawLine {
	return RawLine(strings.TrimSpace(string(l)))
}

// -------------------------------------------------------------------------------------
// Raw Content
// -------------------------------------------------------------------------------------
type RawContent string

// NewRawContent returns a new RawContent wrapper for the given string
func NewRawContent(content string) (RawContent, error) {
	// log.Debugf("new line: '%v'", content)
	return RawContent(content), nil
}

// -------------------------------------------------------------------------------------
// Raw Line
// -------------------------------------------------------------------------------------
type MarkdownQuoteRawLine string

// NewMarkdownQuoteRawLine returns a new slice containing a single StringElement with the given content
func NewMarkdownQuoteRawLine(content string) (MarkdownQuoteRawLine, error) {
	// log.Debugf("new line: '%v'", content)
	return MarkdownQuoteRawLine(content), nil
}

// -------------------------------------------------------------------------------------
// LineRanges: one or more ranges of lines to limit the content of a file to include
// -------------------------------------------------------------------------------------

// NewLineRangesAttribute returns an element attribute with a slice of line ranges attribute for a file inclusion.
// TODO: DEPRECATED
func NewLineRangesAttribute(ranges interface{}) (Attributes, error) {
	switch ranges := ranges.(type) {
	case []interface{}:
		return Attributes{
			AttrLineRanges: NewLineRanges(ranges),
		}, nil
	case LineRange:
		return Attributes{
			AttrLineRanges: NewLineRanges(ranges),
		}, nil
	default:
		return Attributes{
			AttrLineRanges: ranges,
		}, nil
	}
}

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

// NewTagRangesAttribute returns an element attribute with a slice of tag ranges attribute for a file inclusion.
// TODO: DEPRECATED
func NewTagRangesAttribute(ranges interface{}) (Attributes, error) {
	switch ranges := ranges.(type) {
	case []interface{}:
		return Attributes{
			AttrTagRanges: NewTagRanges(ranges),
		}, nil
	case LineRange:
		return Attributes{
			AttrTagRanges: NewTagRanges(ranges),
		}, nil
	default:
		return Attributes{
			AttrTagRanges: ranges,
		}, nil
	}
}

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

// IncludedFileLine a line, containing raw text and inclusion tags
type IncludedFileLine []interface{}

// NewIncludedFileLine returns a new IncludedFileLine
func NewIncludedFileLine(content []interface{}) (IncludedFileLine, error) {
	return IncludedFileLine(Merge(content)), nil
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
	Path   []interface{}
}

// NewLocation return a new location with the given elements
func NewLocation(scheme interface{}, path []interface{}) (*Location, error) {
	path = Merge(path)
	// log.Debugf("new location: scheme='%v' path='%+v", scheme, path)
	s := ""
	if scheme, ok := scheme.([]byte); ok {
		s = string(scheme)
	}
	return &Location{
		Scheme: s,
		Path:   path,
	}, nil
}

// var _ WithElements = &Location{}

// // GetElements returns this section's title
// func (l *Location) GetElements() []interface{} {
// 	return l.Path
// }

// // SetElements sets this section's title
// func (l *Location) SetElements(path []interface{}) {
// 	l.Path = path
// }
func (l *Location) SetPath(elements []interface{}) {
	l.Path = Merge(elements)
}

// SetPathPrefix adds the given prefix to the path if this latter is NOT an absolute
// path and if there is no defined scheme
func (l *Location) SetPathPrefix(p interface{}) {
	if p, ok := p.(string); ok && p != "" {
		if !strings.HasSuffix(p, "/") {
			p = p + "/"
		}
		if l.Scheme == "" && !strings.HasPrefix(l.Stringify(), "/") {
			if u, err := url.Parse(l.Stringify()); err == nil {
				if !u.IsAbs() {
					l.Path = Merge(p, l.Path)
				}
			}
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("set path with prefix: '%s'", spew.Sdump(l.Path...))
	// }
}

// Stringify returns a string representation of the location
func (l Location) Stringify() string {
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

// var _ WithPlaceholdersInElements = &IndexTerm{}

// // RestoreElements restores the elements which had been substituted by placeholders
// func (t *IndexTerm) RestoreElements(placeholders map[string]interface{}) interface{} {
// 	t.Term = restoreElements(t.Term, placeholders)
// 	return t
// }

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

// NewString takes either a single string, or an array of interfaces or strings, and makes
// a single concatenated string.  Used by the parser when simply collecting all characters that
// match would not be desired.
func NewString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case []interface{}:
		res := strings.Builder{}
		for _, item := range v {
			s, e := NewString(item)
			if e != nil {
				return "", e
			}
			res.WriteString(s)
		}
		return res.String(), nil
	default:
		return "", fmt.Errorf("bad string type (%T)", v)
	}
}

// NewInlineAttribute returns a new InlineAttribute if the value is a string (or an error otherwise)
func NewInlineAttribute(name string, value interface{}) (interface{}, error) {
	// log.Debugf("new inline attribute: '%s':'%v'", name, value)
	if value == nil {
		return nil, nil
	}
	value = Reduce(value)
	return Attributes{name: value}, nil
}

// ------------------------------------------------------------------------------------
// Special Characters
// They need to be identified as they may have a special treatment during the rendering
// ------------------------------------------------------------------------------------

// SpecialCharacter a special character, which may get a special treatment later during rendering
type SpecialCharacter struct {
	Name string
}

// NewSpecialCharacter return a new SpecialCharacter
func NewSpecialCharacter(name string) (*SpecialCharacter, error) {
	return &SpecialCharacter{
		Name: name,
	}, nil
}

var _ RawText = SpecialCharacter{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (c SpecialCharacter) RawText() (string, error) {
	return c.Name, nil
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
	// Columns    []*TableColumn
	Rows []*TableRow
}

// func NewTable(attributes interface{}) *Table {
// 	return &Table{
// 		Attributes: toAttributes(attributes),
// 	}
// }
func NewTable(header interface{}, elements []interface{}) (*Table, error) {
	rows := make([]*TableRow, len(elements))
	for i, row := range elements {
		r, ok := row.(*TableRow)
		if !ok {
			return nil, fmt.Errorf("unexpected type of table row: '%T'", r)
		}
		rows[i] = r
	}
	t := &Table{
		Rows: rows,
	}
	if header, ok := header.(*TableRow); ok {
		t.Header = header
	}
	return t, nil
}

// var _ WithElements = &Table{}

// func (t *Table) CanAddElement(element interface{}) bool {
// 	return true
// }

// func (t *Table) AddElement(element interface{}) error {
// 	switch e := element.(type) {
// 	case *TableRow:
// 		if t.Rows == nil {
// 			t.Rows = []*TableRow{}
// 		}
// 		if len(t.Rows) == 0 {
// 			t.Rows = append(t.Rows, &TableRow{})
// 		}
// 		// add all cells of the RawTableRow to the last row of the table
// 		row := t.Rows[len(t.Rows)-1]
// 		for _, c := range e.Cells { // TODO: add a `AddElement()` method to *TableRow?
// 			row.Cells = append(row.Cells, &TableCell{
// 				Elements: []interface{}{
// 					&StringElement{
// 						Content: string(c.Content),
// 					},
// 				},
// 			})
// 		}
// 		return nil
// 	default:
// 		return errors.Errorf("unexpected kind of element to add to a table: '%T'", e)
// 	}
// }

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

var _ BlockWithAttributes = &Table{}

func (t *Table) GetAttributes() Attributes {
	return t.Attributes
}

// AddAttributes adds the attributes of this CalloutListElement
func (t *Table) AddAttributes(attributes Attributes) {
	t.Attributes = t.Attributes.AddAll(attributes)
}

func (t *Table) SetAttributes(attributes Attributes) {
	t.Attributes = attributes
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
					result = append(result, col)
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
	if !t.Attributes.HasOption(AttrAutowidth) {
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
				} else {
					// rounding on the last column, to make sure that the sum reaches 100
					width := (float64(1e6-sumWidth) / 1e4)
					col.Width = strconv.FormatFloat(width, 'g', 6, 64)
				}
			}
		}
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

var _ BlockWithElements = &TableRow{}

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
	Elements []interface{}
}

func NewTableCell(content RawContent) (*TableCell, error) {
	return &TableCell{
		Elements: []interface{}{
			content,
		},
	}, nil
}

var _ BlockWithElements = &TableCell{}

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

// // NewTableLine initializes a new TableLine with the given columns
// func NewTableLine(columns []interface{}) (*TableLine, error) {
// 	c := make([][]interface{}, 0, len(columns))
// 	for _, column := range columns {
// 		if e, ok := column.([]interface{}); ok {
// 			c = append(c, e)
// 		} else {
// 			return nil, errors.Errorf("unsupported element of type %T", column)
// 		}
// 	}
// 	// log.Debugf("initialized a new table line with %d columns", len(c))
// 	return &TableLine{
// 		Cells: c,
// 	}, nil
// }

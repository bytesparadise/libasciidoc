package types

import (
	"bytes"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// interface{} (and other interfaces)
// ------------------------------------------

// Visitable the interface for visitable elements
type Visitable interface {
	AcceptVisitor(Visitor) error
}

// Substituable the interface for substituable elements, ie, which can
// be replaced by another element, for example if they include a FileInclusion
type Substituable interface {
	AcceptSubstitutor(Substitutor) (interface{}, error)
}

// Visitor a visitor that can visit/traverse the interface{} and its children (if applicable)
type Visitor interface {
	Visit(Visitable) error
}

// Substitutor a substitutor that can visit/traverse the interface{} and its children (if applicable)
// and return a new element (or slice of elements) in replacement of the visited element
type Substitutor interface {
	Visit(Substituable) (interface{}, error)
}

// ------------------------------------------
// Draft Document: document parsed in a linear fashion, and which needs further
// processing before rendering
// ------------------------------------------

// DraftDocument the linear-level structure for a document
type DraftDocument struct {
	FrontMatter FrontMatter
	Blocks      []interface{}
}

// NewDraftDocument initializes a new Draft`Document` from the given lines
func NewDraftDocument(frontMatter interface{}, blocks []interface{}) (DraftDocument, error) {
	log.Debugf("initializing a new DraftDocument with %d block element(s)", len(blocks))
	result := DraftDocument{
		Blocks: NilSafe(blocks),
	}
	if fm, ok := frontMatter.(FrontMatter); ok {
		result.FrontMatter = fm
	}
	return result, nil
}

// DocumentAttributes returns
func (d DraftDocument) DocumentAttributes() DocumentAttributes {
	result := DocumentAttributes{}
blocks:
	for _, b := range d.Blocks {
		switch b := b.(type) {
		case Section:
			if b.Level == 0 {
				continue // allow to continue if the section is level 0
			}
			break blocks // otherwise, just stop
		case DocumentAttributeDeclaration:
			result[b.Name] = b.Value
		default:
			break blocks
		}
	}
	return result
}

// ------------------------------------------
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	Attributes         DocumentAttributes
	Elements           []interface{} // TODO: rename to Blocks?
	ElementReferences  ElementReferences
	Footnotes          Footnotes
	FootnoteReferences FootnoteReferences
}

// Title retrieves the document title in its metadata, or empty section title if the title was not specified
func (d Document) Title() (InlineElements, bool) {
	if header, ok := d.Header(); ok {
		return header.Title, true
	}
	return InlineElements{}, false
}

// Authors retrieves the document authors from the document header, or empty array if no author was found
func (d Document) Authors() ([]DocumentAuthor, bool) {
	if header, ok := d.Header(); ok {
		if authors, ok := header.Attributes[AttrAuthors].([]DocumentAuthor); ok {
			return authors, true
		}
	}
	return []DocumentAuthor{}, false
}

// Revision retrieves the document revision from the document header, or empty array if no revision was found
func (d Document) Revision() (DocumentRevision, bool) {
	if header, ok := d.Header(); ok {
		if rev, ok := header.Attributes[AttrRevision].(DocumentRevision); ok {
			return rev, true
		}
	}
	return DocumentRevision{}, false
}

// Header returns the header, i.e., the section with level 0 if it exists as the first element of the document
func (d Document) Header() (Section, bool) {
	if len(d.Elements) > 0 {
		if section, ok := d.Elements[0].(Section); ok && section.Level == 0 {
			return section, true
		}
	}
	return Section{}, false
}

// ------------------------------------------
// Document Element
// ------------------------------------------

// DocumentElement a document element can have attributes
type DocumentElement interface {
	GetAttributes() ElementAttributes
}

// ------------------------------------------
// Document Author
// ------------------------------------------

// DocumentAuthor a document author
type DocumentAuthor struct {
	FullName string
	Email    string
}

// NewDocumentAuthors converts the given authors into an array of `DocumentAuthor`
func NewDocumentAuthors(authors []interface{}) ([]DocumentAuthor, error) {
	log.Debugf("initializing a new array of document authors from `%+v`", authors)
	result := make([]DocumentAuthor, len(authors))
	for i, author := range authors {
		switch author := author.(type) {
		case DocumentAuthor:
			result[i] = author
		default:
			return nil, errors.Errorf("unexpected type of author: %T", author)
		}
	}
	return result, nil
}

// NewDocumentAuthor initializes a new DocumentAuthor
func NewDocumentAuthor(fullName, email interface{}) (DocumentAuthor, error) {
	author := DocumentAuthor{}
	if fullName, ok := fullName.(string); ok {
		author.FullName = fullName
	}
	if email, ok := email.(string); ok {
		author.Email = email
	}
	return author, nil
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
func NewDocumentRevision(revnumber, revdate, revremark interface{}) (DocumentRevision, error) {
	log.Debugf("initializing document revision with revnumber=%v, revdate=%v, revremark=%v", revnumber, revdate, revremark)
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
	result := DocumentRevision{
		Revnumber: number,
		Revdate:   date,
		Revremark: remark,
	}
	return result, nil
}

// ------------------------------------------
// Document Attributes
// ------------------------------------------

// DocumentAttributeDeclaration the type for Document Attribute Declarations
type DocumentAttributeDeclaration struct {
	Name  string
	Value string
}

// NewDocumentAttributeDeclaration initializes a new DocumentAttributeDeclaration with the given name and optional value
func NewDocumentAttributeDeclaration(name string, value interface{}) (DocumentAttributeDeclaration, error) {
	var attrName, attrValue string
	attrName = Apply(name,
		func(s string) string {
			return strings.TrimSpace(s)
		})
	if value, ok := value.(string); ok {
		attrValue = Apply(value,
			func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	log.Debugf("initialized a new DocumentAttributeDeclaration: '%s' -> '%s'", attrName, attrValue)
	return DocumentAttributeDeclaration{
		Name:  attrName,
		Value: attrValue,
	}, nil
}

// DocumentAttributeReset the type for DocumentAttributeReset
type DocumentAttributeReset struct {
	Name string
}

// NewDocumentAttributeReset initializes a new Document Attribute Resets.
func NewDocumentAttributeReset(attrName string) (DocumentAttributeReset, error) {
	log.Debugf("initialized a new DocumentAttributeReset: '%s'", attrName)
	return DocumentAttributeReset{Name: attrName}, nil
}

// DocumentAttributeSubstitution the type for DocumentAttributeSubstitution
type DocumentAttributeSubstitution struct {
	Name string
}

// NewDocumentAttributeSubstitution initializes a new Document Attribute Substitutions
func NewDocumentAttributeSubstitution(attrName string) (DocumentAttributeSubstitution, error) {
	log.Debugf("initialized a new DocumentAttributeSubstitution: '%s'", attrName)
	return DocumentAttributeSubstitution{Name: attrName}, nil
}

// ------------------------------------------
// Element kinds
// ------------------------------------------

// BlockKind the kind of block
type BlockKind string

const (
	// AttrKind the key for the kind of block
	AttrKind string = "kind"
	// Fenced a fenced block
	Fenced BlockKind = "fenced"
	// Listing a listing block
	Listing BlockKind = "listing"
	// Example an example block
	Example BlockKind = "example"
	// Comment a comment block
	Comment BlockKind = "comment"
	// Quote a quote block
	Quote BlockKind = "quote"
	// Verse a verse block
	Verse BlockKind = "verse"
	// Sidebar a sidebar block
	Sidebar BlockKind = "sidebar"
	// Literal a literal block
	Literal BlockKind = "literal"
	// Source a source block
	Source BlockKind = "source"
)

// ------------------------------------------
// Table of Contents
// ------------------------------------------

// TableOfContentsMacro the structure for Table of Contents
type TableOfContentsMacro struct {
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
	Attributes ElementAttributes
	RawText    string
}

// NewUserMacroBlock returns an UserMacro
func NewUserMacroBlock(name string, value string, attrs ElementAttributes, raw string) (UserMacro, error) {
	return UserMacro{
		Name:       name,
		Kind:       BlockMacro,
		Value:      value,
		Attributes: attrs,
		RawText:    raw,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (d UserMacro) AddAttributes(attributes ElementAttributes) {
	d.Attributes.AddAll(attributes)

}

// NewInlineUserMacro returns an UserMacro
func NewInlineUserMacro(name, value string, attrs ElementAttributes, raw string) (UserMacro, error) {
	return UserMacro{Name: name, Kind: InlineMacro, Value: value, Attributes: attrs, RawText: raw}, nil
}

// ------------------------------------------
// Preamble
// ------------------------------------------

// Preamble the structure for document Preamble
type Preamble struct {
	Elements []interface{}
}

// HasContent returns `true` if this Preamble has at least one element which is neither a
// BlankLine nor a DocumentAttributeDeclaration
func (p Preamble) HasContent() bool {
	for _, pe := range p.Elements {
		switch pe.(type) {
		case DocumentAttributeDeclaration, BlankLine:
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
	Content map[string]interface{}
}

// NewYamlFrontMatter initializes a new FrontMatter from the given `content`
func NewYamlFrontMatter(content string) (FrontMatter, error) {
	attributes := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(content), &attributes)
	if err != nil {
		return FrontMatter{}, errors.Wrapf(err, "unable to parse yaml content in front-matter of document")
	}
	log.Debugf("initialized a new FrontMatter with attributes: %+v", attributes)
	return FrontMatter{Content: attributes}, nil
}

// ------------------------------------------
// Sections
// ------------------------------------------

// Section the structure for a section
type Section struct {
	Level      int
	Attributes ElementAttributes
	Title      InlineElements
	Elements   []interface{}
}

// NewSection initializes a new `Section` from the given section title and elements
func NewSection(level int, title InlineElements, ids []interface{}, attributes interface{}) (Section, error) {
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	// log.Debugf("initialized a new Section level %d", level)
	// multiple IDs can be defined (by mistake), and the last one is used
	for _, id := range ids {
		if id, ok := id.(ElementAttributes); ok {
			attrs.AddAll(id)
		}
	}
	attrs[AttrCustomID] = true
	// make a default id from the sectionTitle's inline content
	if _, found := attrs[AttrID]; !found {
		replacement, err := ReplaceNonAlphanumerics(title, "_")
		if err != nil {
			return Section{}, errors.Wrapf(err, "unable to generate default ID while instanciating a new Section element")
		}
		attrs[AttrID] = replacement
		attrs[AttrCustomID] = false
	}
	return Section{
		Level:      level,
		Attributes: attrs,
		Title:      title,
		Elements:   []interface{}{},
	}, nil
}

// UpdateAttrID updates the "ID" attribute in the section (in case the title changed after some document attr substitution)
func (s Section) UpdateAttrID() error {
	if !s.Attributes.GetAsBool(AttrCustomID) {
		replacement, err := ReplaceNonAlphanumerics(s.Title, "_")
		if err != nil {
			return errors.Wrapf(err, "unable to generate default ID while instanciating a new Section element")
		}
		s.Attributes[AttrID] = replacement
	}
	return nil
}

// AddElement adds the given child element to this section
func (s *Section) AddElement(e interface{}) {
	s.Elements = append(s.Elements, e)
}

// Footnotes returns the footnotes found in all the lines of this paragraph
func (s Section) Footnotes() (Footnotes, FootnoteReferences, error) {
	v := NewFootnotesCollector()
	err := s.Title.AcceptVisitor(v)
	return v.Footnotes, v.FootnoteReferences, err
}

// NewDocumentHeader initializes a new Section with level 0 which can have authors and a revision, among other attributes
func NewDocumentHeader(title InlineElements, authors interface{}, revision interface{}) (Section, error) {
	// log.Debugf("initializing a new Section level 0 with authors '%v' and revision '%v'", authors, revision)
	section, err := NewSection(0, title, nil, nil)
	if err != nil {
		return Section{}, err
	}
	if _, ok := authors.([]DocumentAuthor); ok {
		section.Attributes[AttrAuthors] = authors
	}
	if _, ok := revision.(DocumentRevision); ok {
		section.Attributes[AttrRevision] = revision
	}
	return section, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (s *Section) AddAttributes(attributes ElementAttributes) {
	// log.Debugf("adding attributes to section: %v", attributes)
	s.Attributes.AddAll(attributes)
}

// AcceptSubstitutor implements Substituable#AcceptSubstitutor(Substitutor)
// in a section, the substitutor only cares about the elements for now.
func (s *Section) AcceptSubstitutor(v Substitutor) (interface{}, error) {
	substitute := Section{
		Level: s.Level,
		Title: s.Title,
	}
	elements := []interface{}{}
	for _, element := range s.Elements {
		if e, ok := element.(Substituable); ok {
			e, err := e.AcceptSubstitutor(v)
			if err != nil {
				return nil, errors.Wrapf(err, "error while visiting section element for substitution")
			}
			elements = append(elements, e)
		}
	}

	substitute.Elements = elements
	return substitute, nil
}

// ------------------------------------------
// Lists
// ------------------------------------------

// List a list of items
type List interface {
	LastItem() ListItem
}

// ListItem a list item
type ListItem interface {
	AddElement(interface{})
}

// ContinuedListItemElement a wrapper for an element which should be attached to a list item (same level or an ancestor)
type ContinuedListItemElement struct {
	Offset  int // the relative ancestor. Should be a negative number
	Element interface{}
}

// NewContinuedListItemElement returns a wrapper for an element which should be attached to a list item (same level or an ancestor)
func NewContinuedListItemElement(offset int, element interface{}) (ContinuedListItemElement, error) {
	// log.Debugf("initializing a new continued list element for element of type %T", element)
	return ContinuedListItemElement{
		Offset:  offset,
		Element: element,
	}, nil
}

// ------------------------------------------
// List Item Continuation
// ------------------------------------------

// ListItemContinuation the special "+" character to specify that an element belongs to a list item
type ListItemContinuation struct {
}

// NewListItemContinuation returns a new ListItemContinuation
func NewListItemContinuation() (ListItemContinuation, error) {
	return ListItemContinuation{}, nil
}

// ------------------------------------------
// Ordered Lists
// ------------------------------------------

// OrderedList the structure for the Ordered Lists
type OrderedList struct {
	Attributes ElementAttributes
	Items      []OrderedListItem
}

var _ List = &OrderedList{}

// NumberingStyle the type of numbering for items in an ordered list
type NumberingStyle string

const (
	// UnknownNumberingStyle the default, unknown type
	UnknownNumberingStyle NumberingStyle = "unknown"
	// Arabic the arabic numbering (1, 2, 3, etc.)
	Arabic NumberingStyle = "arabic"
	// Decimal the decimal numbering (01, 02, 03, etc.)
	Decimal NumberingStyle = "decimal"
	// LowerAlpha the lower-alpha numbering (a, b, c, etc.)
	LowerAlpha NumberingStyle = "loweralpha"
	// UpperAlpha the upper-alpha numbering (A, B, C, etc.)
	UpperAlpha NumberingStyle = "upperalpha"
	// LowerRoman the lower-roman numbering (i, ii, iii, etc.)
	LowerRoman NumberingStyle = "lowerroman"
	// UpperRoman the upper-roman numbering (I, II, III, etc.)
	UpperRoman NumberingStyle = "upperroman"
	// LowerGreek the lower-greek numbering (alpha, beta, etc.)
	LowerGreek NumberingStyle = "lowergreek"
	// UpperGreek the upper-roman numbering (Alpha, Beta, etc.)
	UpperGreek NumberingStyle = "uppergreek"
)

// NewOrderedList initializes a new ordered list with the given item
func NewOrderedList(item *OrderedListItem) *OrderedList {
	attrs := rearrangeListAttributes(item.Attributes)
	item.Attributes = ElementAttributes{}
	list := &OrderedList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []OrderedListItem{
			*item,
		},
	}
	return list
}

// moves the "upperroman", etc. attributes as values of the `AttrNumberingStyle` key
func rearrangeListAttributes(attributes ElementAttributes) ElementAttributes {
	for k := range attributes {
		switch k {
		case "upperalpha":
			attributes[AttrNumberingStyle] = "upperalpha"
			delete(attributes, k)
		case "upperroman":
			attributes[AttrNumberingStyle] = "upperroman"
			delete(attributes, k)
		case "lowerroman":
			attributes[AttrNumberingStyle] = "lowerroman"
			delete(attributes, k)
		case "loweralpha":
			attributes[AttrNumberingStyle] = "loweralpha"
			delete(attributes, k)
		case "arabic":
			attributes[AttrNumberingStyle] = "arabic"
			delete(attributes, k)
		}

	}
	return attributes
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l *OrderedList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
}

// AddItem adds the given item
func (l *OrderedList) AddItem(item OrderedListItem) {
	l.Items = append(l.Items, item)
}

// FirstItem returns the first item in this list
func (l *OrderedList) FirstItem() ListItem {
	i := l.Items[0]
	return &i
}

// LastItem returns the last item in this list
func (l *OrderedList) LastItem() ListItem {
	return &(l.Items[len(l.Items)-1])
}

// OrderedListItem the structure for the ordered list items
type OrderedListItem struct {
	Attributes     ElementAttributes
	Level          int
	NumberingStyle NumberingStyle
	Elements       []interface{}
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}, attributes interface{}) (OrderedListItem, error) {
	log.Debugf("initializing a new OrderedListItem")
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	return OrderedListItem{
		Attributes:     attrs,
		NumberingStyle: prefix.NumberingStyle,
		Level:          prefix.Level,
		Elements:       elements,
	}, nil
}

// GetAttributes returns the elements of this UnorderedListItem
func (i OrderedListItem) GetAttributes() ElementAttributes {
	return i.Attributes
}

// AddElement add an element to this OrderedListItem
func (i *OrderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i *OrderedListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// OrderedListItemPrefix the prefix used to construct an OrderedListItem
type OrderedListItemPrefix struct {
	NumberingStyle NumberingStyle
	Level          int
}

// NewOrderedListItemPrefix initializes a new OrderedListItemPrefix
func NewOrderedListItemPrefix(s NumberingStyle, l int) (OrderedListItemPrefix, error) {
	return OrderedListItemPrefix{
		NumberingStyle: s,
		Level:          l,
	}, nil
}

// ------------------------------------------
// Unordered Lists
// ------------------------------------------

// UnorderedList the structure for the Unordered Lists
type UnorderedList struct {
	Attributes ElementAttributes
	Items      []UnorderedListItem
}

var _ List = &UnorderedList{}

// NewUnorderedList returns a new UnorderedList with 1 item
// The attributes of the given item are moved to the resulting list
func NewUnorderedList(item *UnorderedListItem) *UnorderedList {
	attrs := item.Attributes
	item.Attributes = ElementAttributes{}
	list := &UnorderedList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []UnorderedListItem{
			*item,
		},
	}
	return list
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l *UnorderedList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
}

// AddItem adds the given item
func (l *UnorderedList) AddItem(item UnorderedListItem) {
	l.Items = append(l.Items, item)
}

// LastItem returns the last item in this list
func (l *UnorderedList) LastItem() ListItem {
	return &(l.Items[len(l.Items)-1])
}

// UnorderedListItem the structure for the unordered list items
type UnorderedListItem struct {
	Level       int
	BulletStyle BulletStyle
	CheckStyle  UnorderedListItemCheckStyle
	Attributes  ElementAttributes
	Elements    []interface{}
}

// NewUnorderedListItem initializes a new `UnorderedListItem` from the given content
func NewUnorderedListItem(prefix UnorderedListItemPrefix, checkstyle interface{}, elements []interface{}, attributes interface{}) (UnorderedListItem, error) {
	log.Debugf("initializing a new UnorderedListItem with %d elements", len(elements))
	// log.Debugf("initializing a new UnorderedListItem with '%d' lines (%T) and input level '%d'", len(elements), elements, lvl.Len())
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	cs := toCheckStyle(checkstyle)
	if cs != NoCheck && len(elements) > 0 {
		if e, ok := elements[0].(ElementWithAttributes); ok {
			e.AddAttributes(ElementAttributes{
				AttrCheckStyle: cs,
			})
		}
	}
	return UnorderedListItem{
		Level:       prefix.Level,
		Attributes:  attrs,
		BulletStyle: prefix.BulletStyle,
		CheckStyle:  cs,
		Elements:    elements,
	}, nil
}

// GetAttributes returns the elements of this UnorderedListItem
func (i UnorderedListItem) GetAttributes() ElementAttributes {
	return i.Attributes
}

// AddElement add an element to this UnorderedListItem
func (i *UnorderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// LastElement returns the last element of this OrderedListItem
func (i *UnorderedListItem) LastElement() interface{} {
	return i.Elements[len(i.Elements)-1]
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i *UnorderedListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// UnorderedListItemCheckStyle the check style that applies on an unordered list item
type UnorderedListItemCheckStyle string

const (
	// Checked when the unordered list item is checked
	Checked UnorderedListItemCheckStyle = "checked"
	// Unchecked when the unordered list item is not checked
	Unchecked UnorderedListItemCheckStyle = "unchecked"
	// NoCheck when the unodered list item has no specific check annotation
	NoCheck UnorderedListItemCheckStyle = "nocheck"
)

func toCheckStyle(checkstyle interface{}) UnorderedListItemCheckStyle {
	if cs, ok := checkstyle.(UnorderedListItemCheckStyle); ok {
		return cs
	}
	return NoCheck
}

// BulletStyle the type of bullet for items in an unordered list
type BulletStyle string

const (
	// UnknownBulletStyle the default, unknown type
	UnknownBulletStyle BulletStyle = "unkwown"
	// Dash an unordered item can begin with a single dash
	Dash BulletStyle = "dash"
	// OneAsterisk an unordered item marked with a single asterisk
	OneAsterisk BulletStyle = "1asterisk"
	// TwoAsterisks an unordered item marked with two asterisks
	TwoAsterisks BulletStyle = "2asterisks"
	// ThreeAsterisks an unordered item marked with three asterisks
	ThreeAsterisks BulletStyle = "3asterisks"
	// FourAsterisks an unordered item marked with four asterisks
	FourAsterisks BulletStyle = "4asterisks"
	// FiveAsterisks an unordered item marked with five asterisks
	FiveAsterisks BulletStyle = "5asterisks"
)

// NextLevel returns the BulletStyle for the next level:
// `-` -> `*`
// `*` -> `**`
// `**` -> `***`
// `***` -> `****`
// `****` -> `*****`
// `*****` -> `-`
func (b BulletStyle) NextLevel(p BulletStyle) BulletStyle {
	switch p {
	case Dash:
		return OneAsterisk
	case OneAsterisk:
		return TwoAsterisks
	case TwoAsterisks:
		return ThreeAsterisks
	case ThreeAsterisks:
		return FourAsterisks
	case FourAsterisks:
		return FiveAsterisks
	case FiveAsterisks:
		return Dash
	}
	// default, return the level itself
	return b
}

// UnorderedListItemPrefix the prefix used to construct an UnorderedListItem
type UnorderedListItemPrefix struct {
	BulletStyle BulletStyle
	Level       int
}

// NewUnorderedListItemPrefix initializes a new UnorderedListItemPrefix
func NewUnorderedListItemPrefix(s BulletStyle, l int) (UnorderedListItemPrefix, error) {
	return UnorderedListItemPrefix{
		BulletStyle: s,
		Level:       l,
	}, nil
}

// NewListItemContent initializes a new `UnorderedListItemContent`
func NewListItemContent(content []interface{}) ([]interface{}, error) {
	// log.Debugf("initializing a new ListItemContent with %d line(s)", len(content))
	elements := make([]interface{}, 0)
	for _, element := range content {
		// log.Debugf("Processing line element of type %T", element)
		switch element := element.(type) {
		case []interface{}:
			elements = append(elements, element...)
		case interface{}:
			elements = append(elements, element)
		}
	}
	// log.Debugf("initialized a new ListItemContent with %d elements(s)", len(elements))
	// no need to return an empty ListItemContent
	if len(elements) == 0 {
		return nil, nil
	}
	return elements, nil
}

// // ListItemContinuation a list item continuation
// type ListItemContinuation struct {
// }

// // NewListItemContinuation returns a new ListItemContinuation
// func NewListItemContinuation() (ListItemContinuation, error) {
// 	return ListItemContinuation{}, nil
// }

// ------------------------------------------
// Labeled List
// ------------------------------------------

// LabeledList the structure for the Labeled Lists
type LabeledList struct {
	Attributes ElementAttributes
	Items      []LabeledListItem
}

var _ List = &LabeledList{}

// NewLabeledList returns a new LabeledList with 1 item
// The attributes of the given item are moved to the resulting list
func NewLabeledList(item LabeledListItem) *LabeledList {
	attrs := rearrangeListAttributes(item.Attributes)
	item.Attributes = ElementAttributes{}
	result := LabeledList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []LabeledListItem{
			item,
		},
	}
	item.Attributes = ElementAttributes{}
	return &result
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l *LabeledList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
}

// AddItem adds the given item
func (l *LabeledList) AddItem(item LabeledListItem) {
	l.Items = append(l.Items, item)
}

// LastItem returns the last item in this list
func (l *LabeledList) LastItem() ListItem {
	return &(l.Items[len(l.Items)-1])
}

// LabeledListItem an item in a labeled
type LabeledListItem struct {
	Term       string
	Level      int
	Attributes ElementAttributes
	Elements   []interface{}
}

// making sure that the `ListItem` interface is implemented by `LabeledListItem`
var _ ListItem = &LabeledListItem{}

// NewLabeledListItem initializes a new LabeledListItem
func NewLabeledListItem(level int, term string, description interface{}, attributes interface{}) (LabeledListItem, error) {
	log.Debugf("initializing a new LabeledListItem")
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	var elements []interface{}
	if description, ok := description.([]interface{}); ok {
		elements = description
	} else {
		elements = []interface{}{}
	}
	return LabeledListItem{
		Attributes: attrs,
		Term:       strings.TrimSpace(term),
		Level:      level,
		Elements:   elements,
	}, nil
}

// GetAttributes returns the elements of this UnorderedListItem
func (i LabeledListItem) GetAttributes() ElementAttributes {
	return i.Attributes
}

// AddElement add an element to this LabeledListItem
func (i *LabeledListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// LastElement returns the last element of this OrderedListItem
func (i LabeledListItem) LastElement() interface{} {
	return i.Elements[len(i.Elements)-1]
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i LabeledListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraphs
type Paragraph struct {
	Attributes ElementAttributes
	Lines      []InlineElements
}

// AttrHardBreaks the attribute to set on a paragraph to render with hard breaks on each line
const AttrHardBreaks = "%hardbreaks"

// DocumentAttrHardBreaks the attribute to set at the document level to render with hard breaks on each line of all paragraphs
const DocumentAttrHardBreaks = "hardbreaks"

// NewParagraph initializes a new `Paragraph`
func NewParagraph(lines []interface{}, attributes interface{}) (Paragraph, error) {
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	// log.Debugf("initializing a new paragraph with %d line(s) and %d attribute(s)", len(lines), len(attrs))
	elements := make([]InlineElements, 0)
	for _, line := range lines {
		if l, ok := line.(InlineElements); ok {
			// log.Debugf("processing paragraph line of type %T", line)
			// if len(l) > 0 {
			elements = append(elements, l)
			// }
		} else {
			return Paragraph{}, errors.Errorf("unsupported paragraph line of type %[1]T: %[1]v", line)
		}

	}
	// log.Debugf("generated a paragraph with %d line(s): %v", len(elements), elements)
	return Paragraph{
		Attributes: attrs,
		Lines:      elements,
	}, nil
}

// Footnotes returns the footnotes found in all the lines of this paragraph
func (p Paragraph) Footnotes() (Footnotes, FootnoteReferences, error) {
	v := NewFootnotesCollector()
	err := p.AcceptVisitor(v)
	if err != nil {
		return nil, nil, err
	}
	return v.Footnotes, v.FootnoteReferences, nil
}

// NewAdmonitionParagraph returns a new Paragraph with an extra admonition attribute
func NewAdmonitionParagraph(lines []interface{}, admonitionKind AdmonitionKind, attributes interface{}) (Paragraph, error) {
	log.Debugf("new admonition paragraph")
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	p, err := NewParagraph(lines, attrs)
	if err != nil {
		return Paragraph{}, err
	}
	p.Attributes[AttrAdmonitionKind] = admonitionKind
	return p, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (p Paragraph) AddAttributes(attributes ElementAttributes) {
	p.Attributes.AddAll(attributes)
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (p *Paragraph) AcceptVisitor(v Visitor) error {
	err := v.Visit(p)
	if err != nil {
		return errors.Wrapf(err, "error while visiting paragraph")
	}
	for _, line := range p.Lines {
		err = line.AcceptVisitor(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting paragraph line")
		}
	}
	return nil
}

// ------------------------------------------
// Admonitions
// ------------------------------------------

// AdmonitionKind the type of admonition
type AdmonitionKind string

const (
	// Tip the 'TIP' type of admonition
	Tip AdmonitionKind = "tip"
	// Note the 'NOTE' type of admonition
	Note AdmonitionKind = "note"
	// Important the 'IMPORTANT' type of admonition
	Important AdmonitionKind = "important"
	// Warning the 'WARNING' type of admonition
	Warning AdmonitionKind = "warning"
	// Caution the 'CAUTION' type of admonition
	Caution AdmonitionKind = "caution"
	// Unknown is the zero value for admonition kind
	Unknown AdmonitionKind = ""
)

// ------------------------------------------
// InlineElements
// ------------------------------------------

// InlineElements the structure for the lines in paragraphs
type InlineElements []interface{}

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements ...interface{}) (InlineElements, error) {
	result := MergeStringElements(elements...)
	return result, nil
}

var _ Visitable = InlineElements{}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (e InlineElements) AcceptVisitor(v Visitor) error {
	err := v.Visit(e)
	if err != nil {
		return errors.Wrapf(err, "error while visiting inline content")
	}
	for _, element := range e {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.AcceptVisitor(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting inline content element")
			}
		}
	}
	return nil
}

// ------------------------------------------
// Cross References
// ------------------------------------------

// CrossReference the struct for Cross References
type CrossReference struct {
	ID    string
	Label string
}

// NewCrossReference initializes a new `CrossReference` from the given ID
func NewCrossReference(id string, label interface{}) (CrossReference, error) {
	log.Debugf("initializing a new CrossReference with ID=%s", id)
	var l string
	if label, ok := label.(string); ok {
		l = Apply(label, strings.TrimSpace)
	}
	return CrossReference{
		ID:    id,
		Label: l,
	}, nil
}

// ------------------------------------------
// Images
// ------------------------------------------

const (
	// AttrImageAlt the image `alt` attribute
	AttrImageAlt string = "alt"
	// AttrImageWidth the image `width` attribute
	AttrImageWidth string = "width"
	// AttrImageHeight the image `height` attribute
	AttrImageHeight string = "height"
	// AttrImageTitle the image `title` attribute
	AttrImageTitle string = "title"
)

// ImageBlock the structure for the block images
type ImageBlock struct {
	Location   Location
	Attributes ElementAttributes
}

// NewImageBlock initializes a new `ImageBlock`
func NewImageBlock(path Location, inlineAttributes ElementAttributes, attributes interface{}) (ImageBlock, error) {
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	attrs.AddAll(inlineAttributes)
	return ImageBlock{
		Location:   path,
		Attributes: attrs,
	}, nil
}

// ResolveLocation resolves the image path using the given document attributes
// also, updates the `alt` attribute based on the resolved path of the image
func (b ImageBlock) ResolveLocation(attrs map[string]string) (interface{}, bool) {
	b.Location = b.Location.Resolve(attrs)
	if _, found := b.Attributes[AttrImageAlt]; !found {
		b.Attributes[AttrImageAlt] = resolveAlt(b.Location)
	}
	return b, true
}

func resolveAlt(path Location) string {
	_, filename := filepath.Split(path.String())
	ext := filepath.Ext(filename)
	if ext != "" {
		return strings.TrimSuffix(filename, ext)
	}
	return filename
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (b ImageBlock) AddAttributes(attributes ElementAttributes) {
	b.Attributes.AddAll(attributes)
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Location   Location
	Attributes ElementAttributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(path Location, attributes ElementAttributes) (InlineImage, error) {
	log.Debugf("new inline image with attrs: %+v", attributes)
	return InlineImage{
		Location:   path,
		Attributes: attributes,
	}, nil
}

// ResolveLocation resolves the image path using the given document attributes
// also, updates the `alt` attribute based on the resolved path of the image
func (i InlineImage) ResolveLocation(attrs map[string]string) (interface{}, bool) {
	i.Location = i.Location.Resolve(attrs)
	if _, found := i.Attributes[AttrImageAlt]; !found {
		i.Attributes[AttrImageAlt] = resolveAlt(i.Location)
	}
	return i, true
}

// // ResolveLocation resolves the image path using the given document attributes
// // also, updates the `alt` attribute based on the resolved path of the image
// func (b InlineImage) ResolveLocation(attrs map[string]string) interface{} {
// 	path := b.Path.Resolve(attrs)
// 	if _, found := attrs[AttrImageAlt]; !found {
// 		_, filename := filepath.Split(path)
// 		ext := filepath.Ext(filename)
// 		log.Debugf("adding alt based on filename '%s' (ext=%s)", filename, ext)
// 		if ext != "" {
// 			b.Attributes[AttrImageAlt] = strings.TrimSuffix(filename, ext)
// 		} else {
// 			b.Attributes[AttrImageAlt] = filename
// 		}
// 	}
// 	return b
// }

// NewImageAttributes returns a map of image attributes, some of which have implicit keys (`alt`, `width` and `height`)
func NewImageAttributes(alt, width, height interface{}, otherattrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	var altStr, widthStr, heightStr string
	if alt, ok := alt.(string); ok {
		if altStr = Apply(alt, strings.TrimSpace); altStr != "" {
			result[AttrImageAlt] = altStr
		}
	}
	if width, ok := width.(string); ok {
		if widthStr = Apply(width, strings.TrimSpace); widthStr != "" {
			result[AttrImageWidth] = widthStr
		}
	}
	if height, ok := height.(string); ok {
		if heightStr = Apply(height, strings.TrimSpace); heightStr != "" {
			result[AttrImageHeight] = heightStr
		}
	}
	for _, otherAttr := range otherattrs {
		if otherAttr, ok := otherAttr.(ElementAttributes); ok {
			for k, v := range otherAttr {
				result[k] = v
				if k == AttrID {
					// mark custom_id flag to `true`
					result[AttrCustomID] = true
				}
			}
		}
	}
	return result, nil
}

// ------------------------------------------
// Footnotes
// ------------------------------------------

var footnoteSequence int

// ResetFootnoteSequence resets the footnote sequence (for test purpose only)
func ResetFootnoteSequence() {
	footnoteSequence = 0
}

// Footnote a foot note, without or without explicit reference (an explicit reference is used to refer
// multiple times to the same footnote across the document)
type Footnote struct {
	ID int
	// Ref the optional reference
	Ref string
	// the footnote content (can be "rich")
	Elements InlineElements
}

// NewFootnote returns a new Footnote with the given content
func NewFootnote(ref string, elements InlineElements) (Footnote, error) {
	defer func() {
		footnoteSequence++
	}()
	footnote := Footnote{
		ID:       footnoteSequence,
		Ref:      ref,
		Elements: elements,
	}
	return footnote, nil
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (f Footnote) AcceptVisitor(v Visitor) error {
	err := v.Visit(f)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	return nil
}

// ------------------------------------------
// Delimited blocks
// ------------------------------------------

// DelimitedBlock the structure for the delimited blocks
type DelimitedBlock struct {
	Kind       BlockKind
	Attributes ElementAttributes
	Elements   []interface{}
}

// Substitution the substitution group to apply when initializing a delimited block
type Substitution func([]interface{}) ([]interface{}, error)

// None returns the content as-is, but nil-safe
func None(content []interface{}) ([]interface{}, error) {
	return NilSafe(content), nil
}

// Verbatim the verbatim substitution: the given content is converted into an array of strings.
func Verbatim(content []interface{}) ([]interface{}, error) {
	result := make([]interface{}, len(content))
	for i, c := range content {
		if c, ok := c.(string); ok {
			c = Apply(c, func(s string) string {
				return strings.TrimRight(c, "\n\r")
			})
			result[i], _ = NewStringElement(c)
		}
	}
	return result, nil
}

// NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given content
func NewDelimitedBlock(kind BlockKind, content []interface{}, substitution Substitution, attributes interface{}) (DelimitedBlock, error) {
	// log.Debugf("initializing a new DelimitedBlock of kind '%v' with %d elements", kind, len(content))
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	elements, err := substitution(content)
	if err != nil {
		return DelimitedBlock{}, errors.Wrapf(err, "failed to initialize a new delimited block")
	}
	if k := attrs.GetAsString(AttrKind); k != "" { // override default kind
		// log.Debugf("overriding kind '%s' to '%s'", b.Kind, attributes[AttrKind])
		kind = BlockKind(k)
	}
	return DelimitedBlock{
		Attributes: attrs,
		Kind:       kind,
		Elements:   elements,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (b *DelimitedBlock) AddAttributes(attributes ElementAttributes) {
	b.Attributes.AddAll(attributes)
	if _, found := attributes[AttrKind]; found { // override default kind
		log.Debugf("overriding kind '%s' to '%s'", b.Kind, attributes[AttrKind])
		b.Kind = BlockKind(attributes.GetAsString(AttrKind))
	}
}

// ------------------------------------------
// Tables
// ------------------------------------------

// Table the structure for the tables
type Table struct {
	Attributes ElementAttributes
	Header     TableLine
	Lines      []TableLine
}

// NewTable initializes a new table with the given lines and attributes
func NewTable(header interface{}, lines []interface{}, attributes interface{}) (Table, error) {
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	t := Table{
		Attributes: attrs,
	}
	columnsPerLine := -1 // unknown until first "line" is processed
	if header, ok := header.(TableLine); ok {
		t.Header = header
		columnsPerLine = len(header.Cells)
	}
	// need to regroup columns of all lines, they dispatch on lines
	cells := make([]InlineElements, 0)
	for _, l := range lines {
		if l, ok := l.(TableLine); ok {
			// if no header line was set, inspect the first line to determine the number of columns per line
			if columnsPerLine == -1 {
				columnsPerLine = len(l.Cells)
			}
			cells = append(cells, l.Cells...)
		}
	}
	t.Lines = make([]TableLine, 0, len(cells))
	if len(lines) > 0 {
		log.Debugf("buffered %d columns for the table", len(cells))
		l := TableLine{
			Cells: make([]InlineElements, columnsPerLine),
		}
		for i, c := range cells {
			log.Debugf("adding cell with content '%v' in table line at offset %d", c, (i % columnsPerLine))
			l.Cells[i%columnsPerLine] = c
			if (i+1)%columnsPerLine == 0 { // switch to next line
				log.Debugf("adding line with content '%v' in table", l)
				t.Lines = append(t.Lines, l)
				l = TableLine{
					Cells: make([]InlineElements, columnsPerLine),
				}
			}
		}
	}
	// log.Debugf("initialized a new table with %d line(s)", len(lines))
	return t, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (t Table) AddAttributes(attributes ElementAttributes) {
	t.Attributes.AddAll(attributes)
}

// TableLine a table line is made of columns, each column being a group of InlineElements (to support quoted text, etc.)
type TableLine struct {
	Cells []InlineElements
}

// NewTableLine initializes a new TableLine with the given columns
func NewTableLine(columns []interface{}) (TableLine, error) {
	c := make([]InlineElements, 0)
	for _, column := range columns {
		if e, ok := column.(InlineElements); ok {
			c = append(c, e)
		} else {
			return TableLine{}, errors.Errorf("unsupported element of type %T", column)
		}
	}
	// log.Debugf("initialized a new table line with %d columns", len(c))
	return TableLine{
		Cells: c,
	}, nil
}

// ------------------------------------------
// Literal blocks
// ------------------------------------------

// LiteralBlock the structure for the literal blocks
type LiteralBlock struct {
	Attributes ElementAttributes
	Lines      []string
}

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

// NewLiteralBlock initializes a new `DelimitedBlock` of the given kind with the given content,
// along with the given sectionTitle spaces
func NewLiteralBlock(origin string, lines []interface{}, attributes interface{}) (LiteralBlock, error) {
	l, err := toString(lines)
	if err != nil {
		return LiteralBlock{}, errors.Wrapf(err, "unable to initialize a new LiteralBlock")
	}
	// log.Debugf("initialized a new LiteralBlock with %d lines", len(lines))
	attrs := ElementAttributes{}
	if attributes, ok := attributes.(ElementAttributes); ok {
		attrs.AddAll(attributes)
	}
	attrs.AddAll(ElementAttributes{
		AttrKind:             Literal,
		AttrLiteralBlockType: origin,
	})
	return LiteralBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (b LiteralBlock) AddAttributes(attributes ElementAttributes) {
	b.Attributes.AddAll(attributes)
}

// ------------------------------------------
// BlankLine
// ------------------------------------------

// BlankLine the structure for the empty lines, which are used to separate logical blocks
type BlankLine struct {
}

// NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (BlankLine, error) {
	// log.Debug("initializing a new BlankLine")
	return BlankLine{}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l BlankLine) AddAttributes(attributes ElementAttributes) {
	// nothing to do
	// TODO: raise a warning?
}

// ------------------------------------------
// Comments
// ------------------------------------------

// SingleLineComment a single line comment
type SingleLineComment struct {
	Content string
}

// NewSingleLineComment initializes a new single line content
func NewSingleLineComment(content string) (SingleLineComment, error) {
	// log.Debugf("initializing a single line comment with content: '%s'", content)
	return SingleLineComment{
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
func NewStringElement(content string) (StringElement, error) {
	return StringElement{Content: content}, nil
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (s StringElement) AcceptVisitor(v Visitor) error {
	err := v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting string element")
	}
	return nil
}

func (s StringElement) String() string {
	return s.Content
}

// ------------------------------------------
// Explicit line breaks
// ------------------------------------------

// LineBreak an explicit line break in a paragraph
type LineBreak struct{}

// NewLineBreak returns a new line break, that's all.
func NewLineBreak() (LineBreak, error) {
	return LineBreak{}, nil
}

// ------------------------------------------
// Quoted text
// ------------------------------------------

// QuotedText the structure for quoted text
type QuotedText struct {
	Kind     QuotedTextKind
	Elements InlineElements
}

// QuotedTextKind the type for
type QuotedTextKind int

const (
	// Bold bold quoted text (wrapped with '*' or '**')
	Bold QuotedTextKind = iota
	// Italic italic quoted text (wrapped with '_' or '__')
	Italic
	// Monospace monospace quoted text (wrapped with '`' or '``')
	Monospace
	// Subscript subscript quoted text (wrapped with '~' or '~~')
	Subscript
	// Superscript superscript quoted text (wrapped with '^' or '^^')
	Superscript
)

// NewQuotedText initializes a new `QuotedText` from the given kind and content
func NewQuotedText(kind QuotedTextKind, content ...interface{}) (QuotedText, error) {
	elements := MergeStringElements(content...)
	return QuotedText{
		Kind:     kind,
		Elements: elements,
	}, nil
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (t QuotedText) AcceptVisitor(v Visitor) error {
	err := v.Visit(t)
	if err != nil {
		return errors.Wrapf(err, "error while visiting quoted text")
	}
	for _, element := range t.Elements {
		if visitable, ok := element.(Visitable); ok {
			err := visitable.AcceptVisitor(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting quoted text element")
			}
		}
	}
	return nil
}

// -------------------------------------------------------
// Escaped Quoted Text (i.e., with substitution preserved)
// -------------------------------------------------------

// NewEscapedQuotedText returns a new InlineElements where the nested elements are preserved (ie, substituted as expected)
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
		StringElement{
			Content: backslashesStr,
		},
		StringElement{
			Content: punctuation,
		},
		content,
		StringElement{
			Content: punctuation,
		},
	}, nil
}

// ------------------------------------------
// Passthrough
// ------------------------------------------

// Passthrough the structure for Passthroughs
type Passthrough struct {
	Kind     PassthroughKind
	Elements InlineElements
}

// PassthroughKind the kind of passthrough
type PassthroughKind int

const (
	// SinglePlusPassthrough a passthrough with a single `+` punctuation
	SinglePlusPassthrough PassthroughKind = iota
	// TriplePlusPassthrough a passthrough with a triple `+++` punctuation
	TriplePlusPassthrough
	// PassthroughMacro a passthrough with the `pass:[]` macro
	PassthroughMacro
)

// NewPassthrough returns a new passthrough
func NewPassthrough(kind PassthroughKind, elements []interface{}) (Passthrough, error) {
	return Passthrough{
		Kind:     kind,
		Elements: MergeStringElements(elements...),
	}, nil

}

// ------------------------------------------
// Inline Links
// ------------------------------------------

// InlineLink the structure for the external links
type InlineLink struct {
	Location   Location
	Attributes ElementAttributes
}

// NewInlineLink initializes a new inline `InlineLink`
func NewInlineLink(url Location, attrs interface{}) (InlineLink, error) {
	result := InlineLink{
		Location: url,
	}
	if attrs, ok := attrs.(ElementAttributes); ok {
		result.Attributes = attrs
	} else {
		result.Attributes = ElementAttributes{}
	}
	return result, nil
}

// // ResolveLocation resolves the link path using the given document attributes
// func (b InlineLink) ResolveLocation() interface{} {
// 	b.Path.Resolve(attrs)
// 	return b
// }

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (l InlineLink) AcceptVisitor(v Visitor) error {
	err := v.Visit(l)
	if err != nil {
		return errors.Wrapf(err, "error while visiting inline link")
	}
	return nil
}

// AttrInlineLinkText the link `text` attribute
const AttrInlineLinkText string = "text"

// NewInlineLinkAttributes returns a map of link attributes, some of which have implicit keys (`text`)
func NewInlineLinkAttributes(text interface{}, otherattrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	if text, ok := text.(InlineElements); ok {
		result[AttrInlineLinkText] = text
	}
	for _, otherAttr := range otherattrs {
		if otherAttr, ok := otherAttr.(ElementAttributes); ok {
			for k, v := range otherAttr {
				result[k] = v
			}
		}
	}
	return result, nil
}

// ------------------------------------------
// File Inclusions
// ------------------------------------------

// FileInclusion the structure for the file inclusions
type FileInclusion struct {
	Attributes ElementAttributes
	Location   Location
	RawText    string
}

var _ ElementWithAttributes = FileInclusion{}

// NewFileInclusion initializes a new inline `InlineLink`
func NewFileInclusion(location Location, attributes interface{}, rawtext string) (FileInclusion, error) {
	attrs, ok := attributes.(ElementAttributes)
	// init attributes with empty 'text' attribute
	if !ok {
		attrs = ElementAttributes{}
	}
	return FileInclusion{
		Attributes: attrs,
		Location:   location,
		RawText:    rawtext,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (f FileInclusion) AddAttributes(attributes ElementAttributes) {
	f.Attributes.AddAll(attributes)
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
// LineRanges: one or more ranges of lines to limit the content of a file to include
// -------------------------------------------------------------------------------------

// NewLineRangesAttribute returns an element attribute with a slice of line ranges attribute for a file inclusion.
func NewLineRangesAttribute(ranges interface{}) (ElementAttributes, error) {
	switch ranges := ranges.(type) {
	case []interface{}:
		return ElementAttributes{
			AttrLineRanges: NewLineRanges(ranges...),
		}, nil
	case LineRange:
		return ElementAttributes{
			AttrLineRanges: NewLineRanges(ranges),
		}, nil
	default:
		return ElementAttributes{
			AttrLineRanges: ranges,
		}, nil
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
	// log.Debugf("initializing a new multiline range: %d..%d", start, end)
	return LineRange{
		StartLine: start,
		EndLine:   end,
	}, nil
}

// LineRanges the ranges of lines of the child doc to include in the master doc
type LineRanges []LineRange

// NewLineRanges returns a slice of line ranges attribute for a file inclusion.
func NewLineRanges(ranges ...interface{}) LineRanges {
	result := LineRanges{}
	for _, r := range ranges {
		if lr, ok := r.(LineRange); ok {
			result = append(result, lr)
		}
	}
	// sort the range by `start` line
	sort.Sort(result)
	return result
}

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
func NewTagRangesAttribute(ranges []interface{}) (ElementAttributes, error) {
	r, err := NewTagRanges(ranges...)
	if err != nil {
		return nil, err
	}
	log.Debugf("initialized a new TagRanges attribute with values: %v", r)
	return ElementAttributes{
		AttrTagRanges: r,
	}, nil
}

// TagRanges the ranges of tags of the child doc to include in the master doc
type TagRanges []TagRange

// NewTagRanges returns a slice of tag ranges attribute for a file inclusion.
func NewTagRanges(ranges ...interface{}) (TagRanges, error) {
	result := TagRanges{}
	for _, r := range ranges {
		if tr, ok := r.(TagRange); ok {
			result = append(result, tr)
		} else {
			return nil, errors.Errorf("unexpected type of tag range: %T", r)
		}
	}
	return result, nil
}

// Match checks if the given tag matches one of the range
func (tr TagRanges) Match(line int, currentRanges CurrentRanges) bool {
	match := false
	log.Debugf("checking line %d", line)

	// compare with expected tag ranges
	for _, t := range tr {
		if t.Name == "**" {
			match = true
			continue
		}
		for n, r := range currentRanges {
			log.Debugf("checking if range %s (%v) matches one of %v", n, r, tr)
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
	return IncludedFileLine(MergeStringElements(content)), nil
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
	Elements []interface{}
}

// NewLocation return a new location with the given elements
func NewLocation(elements []interface{}) (Location, error) {
	return Location{
		Elements: MergeStringElements(elements),
	}, nil
}

// Resolve resolves the Location by replacing all document attribute substitutions
// with their associated values, or their corresponding raw text if
// no attribute matched
// returns the resolved attribute
func (l Location) String() string { // (attrs map[string]string) string {
	result := bytes.NewBuffer(nil)
	for _, e := range l.Elements {
		result.WriteString(fmt.Sprintf("%s", e))
	}
	return result.String()
}

// Resolve resolves the Location by replacing all document attribute substitutions
// with their associated values, or their corresponding raw text if
// no attribute matched
// returns `true` if some document attribute substitution occurred
func (l *Location) Resolve(attrs map[string]string) Location {
	result := bytes.NewBuffer(nil)
	for _, e := range l.Elements {
		switch e := e.(type) {
		case DocumentAttributeSubstitution:
			if value, found := attrs[e.Name]; found {
				result.WriteString(value)
			} else {
				result.WriteRune('{')
				result.WriteString(e.Name)
				result.WriteRune('}')
			}
		default:
			result.WriteString(fmt.Sprintf("%s", e))
		}
	}
	return Location{
		Elements: []interface{}{
			StringElement{
				Content: result.String(),
			},
		},
	}
}

// LocationHolder interface for the types that embed and can resolve a location
type LocationHolder interface {
	ResolveLocation(attrs map[string]string) (interface{}, bool)
}

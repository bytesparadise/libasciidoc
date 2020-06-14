package types

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

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
		Blocks: blocks,
	}
	if fm, ok := frontMatter.(FrontMatter); ok {
		result.FrontMatter = fm
	}
	return result, nil
}

// Attributes returns the document attributes on the top-level section
// and all the document attribute declarations at the top of the document only.
func (d DraftDocument) Attributes() Attributes {
	result := Attributes{}
blocks:
	for _, b := range d.Blocks {
		switch b := b.(type) {
		case Section:
			if b.Level == 0 {
				// also, expand document authors and revision
				if authors, ok := b.Attributes[AttrAuthors].([]DocumentAuthor); ok {
					// move to the Document attributes
					result.Add(expandAuthors(authors))
					delete(b.Attributes, AttrAuthors)
				}
				// also, expand document authors and revision
				if revision, ok := b.Attributes[AttrRevision].(DocumentRevision); ok {
					// move to the Document attributes
					result.Add(expandRevision(revision))
					delete(b.Attributes, AttrRevision)
				}
				continue // allow to continue if the section is level 0
			}
			break blocks // otherwise, just stop
		case AttributeDeclaration:
			result.Set(b.Name, b.Value)
		default:
			break blocks
		}
	}
	log.Debugf("document attributes: %+v", result)
	return result
}

// ------------------------------------------
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	Attributes        Attributes
	Elements          []interface{} // TODO: rename to `Blocks`?
	ElementReferences ElementReferences
	Footnotes         []Footnote
}

// Authors retrieves the document authors from the document header, or empty array if no author was found
func (d Document) Authors() ([]DocumentAuthor, bool) {
	if authors, ok := d.Attributes[AttrAuthors].([]DocumentAuthor); ok {
		return authors, true
	}
	return []DocumentAuthor{}, false
}

// Revision retrieves the document revision from the document header, or empty array if no revision was found
func (d Document) Revision() (DocumentRevision, bool) {
	if rev, ok := d.Attributes[AttrRevision].(DocumentRevision); ok {
		return rev, true
	}
	return DocumentRevision{}, false
}

// Header returns the header, i.e., the section with level 0 if it found as the first element of the document
// For manpage documents, this also includes the first section (`Name` along with its first paragraph)
func (d Document) Header() (Section, bool) {
	if len(d.Elements) == 0 {
		return Section{}, false
	}
	if section, ok := d.Elements[0].(Section); ok && section.Level == 0 {
		return section, true
	}
	return Section{}, false
}

// ------------------------------------------
// Document Metadata
// ------------------------------------------

// Metadata the document metadata returned after the rendering
type Metadata struct {
	Title           string
	LastUpdated     string
	TableOfContents TableOfContents
	Authors         []DocumentAuthor
	Revision        DocumentRevision
}

// TableOfContents the table of contents
type TableOfContents struct {
	Sections []ToCSection
}

// ToCSection a section in the table of contents
type ToCSection struct {
	ID       string
	Level    int
	Title    string // the title as it was rendered in HTML
	Children []ToCSection
}

// ------------------------------------------
// Document Element
// ------------------------------------------

// DocumentElement a document element can have attributes
type DocumentElement interface {
	GetAttributes() Attributes
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

// AttributeDeclaration the type for Document Attribute Declarations
type AttributeDeclaration struct {
	Name  string
	Value string
}

// NewAttributeDeclaration initializes a new AttributeDeclaration with the given name and optional value
func NewAttributeDeclaration(name string, value interface{}) (AttributeDeclaration, error) {
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
	log.Debugf("initialized a new AttributeDeclaration: '%s' -> '%s'", attrName, attrValue)
	return AttributeDeclaration{
		Name:  attrName,
		Value: attrValue,
	}, nil
}

// AttributeReset the type for AttributeReset
type AttributeReset struct {
	Name string
}

// NewAttributeReset initializes a new Document Attribute Resets.
func NewAttributeReset(attrName string) (AttributeReset, error) {
	log.Debugf("initialized a new AttributeReset: '%s'", attrName)
	return AttributeReset{Name: attrName}, nil
}

// AttributeSubstitution the type for AttributeSubstitution
type AttributeSubstitution struct {
	Name string
}

// NewAttributeSubstitution initializes a new Document Attribute Substitutions
func NewAttributeSubstitution(attrName string) (AttributeSubstitution, error) {
	log.Debugf("initialized a new AttributeSubstitution: '%s'", attrName)
	return AttributeSubstitution{Name: attrName}, nil
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
	// MarkdownQuote a quote block in the Markdown style
	MarkdownQuote BlockKind = "markdown-quote"
	// Verse a verse block
	Verse BlockKind = "verse"
	// Sidebar a sidebar block
	Sidebar BlockKind = "sidebar"
	// Literal a literal block
	Literal BlockKind = "literal"
	// Source a source block
	Source BlockKind = "source"
	// Passthrough a passthrough block
	Passthrough BlockKind = "passthrough"
)

// ------------------------------------------
// Table of Contents
// ------------------------------------------

// TableOfContentsPlaceHolder a place holder for Table of Contents, so
// the renderer knows when to render it.
type TableOfContentsPlaceHolder struct {
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
func NewUserMacroBlock(name string, value string, attributes interface{}, raw string) (UserMacro, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return UserMacro{}, errors.Wrap(err, "failed to initialize a UserMacro element")
	}
	return UserMacro{
		Name:       name,
		Kind:       BlockMacro,
		Value:      value,
		Attributes: attrs,
		RawText:    raw,
	}, nil
}

// NewInlineUserMacro returns an UserMacro
func NewInlineUserMacro(name, value string, attributes interface{}, raw string) (UserMacro, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return UserMacro{}, errors.Wrap(err, "failed to initialize a UserMacro element")
	}
	return UserMacro{
		Name:       name,
		Kind:       InlineMacro,
		Value:      value,
		Attributes: attrs,
		RawText:    raw,
	}, nil
}

// ------------------------------------------
// Preamble
// ------------------------------------------

// Preamble the structure for document Preamble
type Preamble struct {
	Elements []interface{}
}

// HasContent returns `true` if this Preamble has at least one element which is neither a
// BlankLine nor a AttributeDeclaration
func (p Preamble) HasContent() bool {
	for _, pe := range p.Elements {
		switch pe.(type) {
		case BlankLine:
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
		return FrontMatter{}, errors.Wrapf(err, "failed to parse yaml content in front-matter of document")
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
	Attributes Attributes
	Title      []interface{}
	Elements   []interface{}
}

// NewSection initializes a new `Section` from the given section title and elements
func NewSection(level int, title []interface{}, ids []interface{}, attributes interface{}) (Section, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return Section{}, errors.Wrapf(err, "failed to initialize a Section element")
	}
	// multiple IDs can be defined (by mistake), but only the last one is used
	for _, id := range ids {
		attrs = attrs.Add(id)

	}
	return Section{
		Level:      level,
		Attributes: attrs,
		Title:      title,
		Elements:   []interface{}{},
	}, nil
}

// ResolveID resolves/updates the "ID" attribute in the section (in case the title changed after some document attr substitution)
func (s Section) ResolveID(docAttributes AttributesWithOverrides) (Section, error) {
	if !s.Attributes.GetAsBool(AttrCustomID) {
		replacement, err := ReplaceNonAlphanumerics(s.Title, "_")
		if err != nil {
			return s, errors.Wrapf(err, "failed to generate default ID on Section element")
		}
		idPrefix := docAttributes.GetAsStringWithDefault(AttrIDPrefix, DefaultIDPrefix)
		s.Attributes = s.Attributes.Set(AttrID, idPrefix+replacement)
		log.Debugf("updated section id to '%s'", s.Attributes[AttrID])
	}
	return s, nil
}

// AddElement adds the given child element to this section
func (s *Section) AddElement(e interface{}) {
	s.Elements = append(s.Elements, e)
}

var _ FootnotesContainer = Section{}

// ReplaceFootnotes replaces the footnotes in the section title
// with footnote references. The footnotes are stored in the given 'notes' param
func (s Section) ReplaceFootnotes(notes *Footnotes) interface{} {
	for i, element := range s.Title {
		if note, ok := element.(Footnote); ok {
			s.Title[i] = notes.Reference(note)
		}
	}
	return s
}

// NewDocumentHeader initializes a new Section with level 0 which can have authors and a revision, among other attributes
func NewDocumentHeader(title []interface{}, authors interface{}, revision interface{}) (Section, error) {
	// log.Debugf("initializing a new Section level 0 with authors '%v' and revision '%v'", authors, revision)
	section, err := NewSection(0, title, nil, nil)
	if err != nil {
		return Section{}, err
	}
	if authors, ok := authors.([]DocumentAuthor); ok {
		section.Attributes = section.Attributes.Set(AttrAuthors, authors)
	}
	if revision, ok := revision.(DocumentRevision); ok {
		section.Attributes = section.Attributes.Set(AttrRevision, revision)
	}
	return section, nil
}

// expandAuthors returns a map of attributes for the given authors.
// those attributes can be used in attribute substitutions in the document
func expandAuthors(authors []DocumentAuthor) Attributes {
	result := make(map[string]interface{}, 1+6*len(authors)) // each author may add up to 6 fields in the result map
	sanitized := make([]DocumentAuthor, 0, len(authors))
	for i, author := range authors {
		var part1, part2, part3, email string
		author.FullName = strings.ReplaceAll(author.FullName, "  ", " ")
		parts := strings.Split(author.FullName, " ")
		if len(parts) > 0 {
			part1 = Apply(parts[0],
				func(s string) string {
					return strings.TrimSpace(s)
				},
				func(s string) string {
					return strings.Replace(s, "_", " ", -1)
				},
			)
		}
		if len(parts) > 1 {
			part2 = Apply(parts[1],
				func(s string) string {
					return strings.TrimSpace(s)
				},
				func(s string) string {
					return strings.Replace(s, "_", " ", -1)
				},
			)
		}
		if len(parts) > 2 {
			part3 = Apply(strings.Join(parts[2:], " "),
				func(s string) string {
					return strings.TrimSpace(s)
				},
				func(s string) string {
					return strings.Replace(s, "_", " ", -1)
				},
			)
		}
		if author.Email != "" {
			email = strings.TrimSpace(author.Email)
		}
		if part2 != "" && part3 != "" {
			result[key("firstname", i)] = strings.TrimSpace(part1)
			result[key("middlename", i)] = strings.TrimSpace(part2)
			result[key("lastname", i)] = strings.TrimSpace(part3)
			result[key("author", i)] = strings.Join([]string{part1, part2, part3}, " ")
			result[key("authorinitials", i)] = strings.Join([]string{initial(part1), initial(part2), initial(part3)}, "")
		} else if part2 != "" {
			result[key("firstname", i)] = strings.TrimSpace(part1)
			result[key("lastname", i)] = strings.TrimSpace(part2)
			result[key("author", i)] = strings.Join([]string{part1, part2}, " ")
			result[key("authorinitials", i)] = strings.Join([]string{initial(part1), initial(part2)}, "")
		} else {
			result[key("firstname", i)] = strings.TrimSpace(part1)
			result[key("author", i)] = strings.TrimSpace(part1)
			result[key("authorinitials", i)] = initial(part1)
		}
		if email != "" {
			result[key("email", i)] = email
		}
		// also include a "sanitized" version of the given author
		sanitized = append(sanitized, DocumentAuthor{
			FullName: result[key("author", i)].(string),
			Email:    email,
		})
	}
	result[AttrAuthors] = sanitized
	log.Debugf("authors: %v", result)
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

// expandRevision returns a map of attributes for the given revision.
// those attributes can be used in attribute substitutions in the document
func expandRevision(revision DocumentRevision) Attributes {
	result := make(Attributes, 3)
	result.AddNonEmpty("revnumber", revision.Revnumber)
	result.AddNonEmpty("revdate", revision.Revdate)
	result.AddNonEmpty("revremark", revision.Revremark)
	// also add the revision itself
	result.AddNonEmpty(AttrRevision, revision)
	log.Debugf("revision: %v", result)
	return result
}

// ------------------------------------------
// Lists
// ------------------------------------------

// List a list of items
type List interface {
	LastItem() ListItem
}

// ListItem a list item
type ListItem interface { // TODO: convert to struct and use as composant in OrderedListItem, etc.
	AddElement(interface{})
}

// ContinuedListItemElement a wrapper for an element which should be attached to a list item (same level or an ancestor)
type ContinuedListItemElement struct {
	Offset  int // the relative ancestor. Should be a negative number
	Element interface{}
}

// NewContinuedListItemElement returns a wrapper for an element which should be attached to a list item (same level or an ancestor)
func NewContinuedListItemElement(element interface{}) (ContinuedListItemElement, error) {
	// log.Debugf("initializing a new continued list element for element of type %T", element)
	return ContinuedListItemElement{
		Offset:  0,
		Element: element,
	}, nil
}

// ------------------------------------------
// Ordered Lists
// ------------------------------------------

// OrderedList the structure for the Ordered Lists
type OrderedList struct {
	Attributes Attributes
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
	item.Attributes = nil
	return &OrderedList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []OrderedListItem{
			*item,
		},
	}
}

// moves the "upperroman", etc. attributes as values of the `AttrNumberingStyle` key
func rearrangeListAttributes(attributes Attributes) Attributes {
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

// AddItem adds the given item
func (l *OrderedList) AddItem(item OrderedListItem) {
	l.Items = append(l.Items, item)
}

// LastItem returns the last item in this list
func (l *OrderedList) LastItem() ListItem {
	return &(l.Items[len(l.Items)-1])
}

// OrderedListItem the structure for the ordered list items
type OrderedListItem struct {
	Attributes     Attributes
	Level          int
	NumberingStyle NumberingStyle
	Elements       []interface{} // TODO: rename to `Blocks`?
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}, attributes interface{}) (OrderedListItem, error) {
	log.Debugf("initializing a new OrderedListItem")
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return OrderedListItem{}, errors.Wrapf(err, "failed to initialize an OrderedListItem element")
	}
	return OrderedListItem{
		Attributes:     attrs,
		NumberingStyle: prefix.NumberingStyle,
		Level:          prefix.Level,
		Elements:       elements,
	}, nil
}

// GetAttributes returns the elements of this OrderedListItem
func (i OrderedListItem) GetAttributes() Attributes {
	return i.Attributes
}

// AddElement add an element to this OrderedListItem
func (i *OrderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
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
	Attributes Attributes
	Items      []UnorderedListItem
}

var _ List = &UnorderedList{}

// NewUnorderedList returns a new UnorderedList with 1 item
// The attributes of the given item are moved to the resulting list
func NewUnorderedList(item *UnorderedListItem) *UnorderedList {
	attrs := item.Attributes
	item.Attributes = nil
	list := &UnorderedList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []UnorderedListItem{
			*item,
		},
	}
	return list
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
	Attributes  Attributes
	Elements    []interface{} // TODO: rename to `Blocks`?
}

// NewUnorderedListItem initializes a new `UnorderedListItem` from the given content
func NewUnorderedListItem(prefix UnorderedListItemPrefix, checkstyle interface{}, elements []interface{}, attributes interface{}) (UnorderedListItem, error) {
	log.Debugf("initializing a new UnorderedListItem with %d elements", len(elements))
	// log.Debugf("initializing a new UnorderedListItem with '%d' lines (%T) and input level '%d'", len(elements), elements, lvl.Len())
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return UnorderedListItem{}, errors.Wrapf(err, "failed to initialize an UnorderedListItem element")
	}
	cs := toCheckStyle(checkstyle)
	if cs != NoCheck && len(elements) > 0 {
		if p, ok := elements[0].(Paragraph); ok {
			if p.Attributes == nil {
				p.Attributes = Attributes{}
				elements[0] = p // need to update the element in the slice
			}
			p.Attributes[AttrCheckStyle] = cs
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
func (i UnorderedListItem) GetAttributes() Attributes {
	return i.Attributes
}

// AddElement add an element to this UnorderedListItem
func (i *UnorderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
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

// ------------------------------------------
// Labeled List
// ------------------------------------------

// LabeledList the structure for the Labeled Lists
type LabeledList struct {
	Attributes Attributes
	Items      []LabeledListItem
}

var _ List = &LabeledList{}

// NewLabeledList returns a new LabeledList with 1 item
// The attributes of the given item are moved to the resulting list
func NewLabeledList(item LabeledListItem) *LabeledList {
	attrs := rearrangeListAttributes(item.Attributes)
	item.Attributes = nil
	result := LabeledList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []LabeledListItem{
			item,
		},
	}
	return &result
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
	Term       []interface{}
	Level      int
	Attributes Attributes
	Elements   []interface{} // TODO: rename to `Blocks`?
}

// making sure that the `ListItem` interface is implemented by `LabeledListItem`
var _ ListItem = &LabeledListItem{}

// NewLabeledListItem initializes a new LabeledListItem
func NewLabeledListItem(level int, term []interface{}, description interface{}, attributes interface{}) (LabeledListItem, error) {
	log.Debugf("initializing a new LabeledListItem")
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return LabeledListItem{}, errors.Wrapf(err, "failed to initialize a LabeledListItem element")
	}
	var elements []interface{}
	if description, ok := description.([]interface{}); ok {
		elements = description
	} else {
		elements = []interface{}{}
	}
	return LabeledListItem{
		Attributes: attrs,
		Term:       term,
		Level:      level,
		Elements:   elements,
	}, nil
}

// GetAttributes returns the elements of this LabeledListItem
func (i LabeledListItem) GetAttributes() Attributes {
	return i.Attributes
}

// AddElement add an element to this LabeledListItem
func (i *LabeledListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraphs
type Paragraph struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// AttrHardBreaks the attribute to set on a paragraph to render with hard breaks on each line
const AttrHardBreaks = "%hardbreaks"

// DocumentAttrHardBreaks the attribute to set at the document level to render with hard breaks on each line of all paragraphs
const DocumentAttrHardBreaks = "hardbreaks"

// NewParagraph initializes a new `Paragraph`
func NewParagraph(lines []interface{}, attributes interface{}) (Paragraph, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return Paragraph{}, errors.Wrapf(err, "failed to initialize a Paragraph element")
	}
	// log.Debugf("initializing a new paragraph with %d line(s) and %d attribute(s)", len(lines), len(attrs))
	elements := make([][]interface{}, 0)
	for _, line := range lines {
		if l, ok := line.([]interface{}); ok {
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

var _ FootnotesContainer = Paragraph{}

// ReplaceFootnotes replaces the footnotes in the paragraph lines
// with footnote references. The footnotes are stored in the given 'notes' param
func (p Paragraph) ReplaceFootnotes(notes *Footnotes) interface{} {
	for i, line := range p.Lines {
		for j, element := range line {
			if note, ok := element.(Footnote); ok {
				p.Lines[i][j] = notes.Reference(note)
			}
		}
	}
	return p
}

// NewAdmonitionParagraph returns a new Paragraph with an extra admonition attribute
func NewAdmonitionParagraph(lines []interface{}, admonitionKind AdmonitionKind, attributes interface{}) (Paragraph, error) {
	log.Debugf("new admonition paragraph")
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return Paragraph{}, errors.Wrapf(err, "failed to initialize an Admonition Paragraph element")
	}
	p, err := NewParagraph(lines, attrs)
	if err != nil {
		return Paragraph{}, err
	}
	p.Attributes = p.Attributes.Set(AttrAdmonitionKind, admonitionKind)
	return p, nil
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

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements ...interface{}) ([]interface{}, error) {
	return Merge(elements...), nil
}

// ------------------------------------------
// Cross References
// ------------------------------------------

// InternalCrossReference the struct for Cross References
type InternalCrossReference struct {
	ID    string
	Label string
}

// NewInternalCrossReference initializes a new `InternalCrossReference` from the given ID
func NewInternalCrossReference(id string, label interface{}) (InternalCrossReference, error) {
	log.Debugf("initializing a new InternalCrossReference with ID=%s", id)
	var l string
	if label, ok := label.(string); ok {
		l = Apply(label, strings.TrimSpace)
	}
	return InternalCrossReference{
		ID:    id,
		Label: l,
	}, nil
}

// ExternalCrossReference the struct for Cross References
type ExternalCrossReference struct {
	Location Location
	Label    []interface{}
}

// NewExternalCrossReference initializes a new `InternalCrossReference` from the given ID
func NewExternalCrossReference(location Location, attributes Attributes) (ExternalCrossReference, error) {
	var label []interface{}
	if l, ok := attributes["positional-1"].([]interface{}); ok {
		label = l
	}
	log.Debugf("initializing a new ExternalCrossReference with Location=%v and label='%s' (attrs=%v / %T)", location, label, attributes, attributes[AttrInlineLinkText])
	return ExternalCrossReference{
		Location: location,
		Label:    label,
	}, nil
}

// ResolveLocation resolves the image path using the given document attributes
// also, updates the `alt` attribute based on the resolved path of the image
func (r ExternalCrossReference) ResolveLocation(attrs AttributesWithOverrides) ExternalCrossReference {
	r.Location = r.Location.Resolve(attrs)
	return r
}

// ------------------------------------------
// Images
// ------------------------------------------

// ImageBlock the structure for the block images
type ImageBlock struct {
	Location   Location
	Attributes Attributes
}

// NewImageBlock initializes a new `ImageBlock`
func NewImageBlock(path Location, inlineAttributes Attributes, attributes interface{}) (ImageBlock, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return ImageBlock{}, errors.Wrapf(err, "failed to initialize an ImageBlock element")
	}
	if attrs == nil && len(inlineAttributes) > 0 {
		attrs = inlineAttributes
	} else if len(inlineAttributes) > 0 {
		attrs = attrs.Add(inlineAttributes)
	}
	return ImageBlock{
		Location:   path,
		Attributes: attrs,
	}, nil
}

// ResolveLocation resolves the image path using the given document attributes
// also, updates the `alt` attribute based on the resolved path of the image
func (b ImageBlock) ResolveLocation(attrs AttributesWithOverrides) ImageBlock {
	b.Location = b.Location.Resolve(attrs)
	if _, found := b.Attributes[AttrImageAlt]; !found {
		b.Attributes = b.Attributes.Set(AttrImageAlt, resolveAlt(b.Location))
	}
	return b
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Location   Location
	Attributes Attributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(path Location, attributes Attributes) (InlineImage, error) {
	return InlineImage{
		Location:   path,
		Attributes: attributes,
	}, nil
}

// ResolveLocation resolves the image path using the given document attributes
// also, updates the `alt` attribute based on the resolved path of the image
func (i InlineImage) ResolveLocation(attrs AttributesWithOverrides) InlineImage {
	i.Location = i.Location.Resolve(attrs)
	if _, found := i.Attributes[AttrImageAlt]; !found {
		i.Attributes = i.Attributes.Set(AttrImageAlt, resolveAlt(i.Location))
	}
	return i
}

// NewImageAttributes returns a map of image attributes, some of which have implicit keys (`alt`, `width` and `height`)
func NewImageAttributes(alt, width, height interface{}, otherattrs []interface{}) (Attributes, error) {
	var result Attributes

	if alt, ok := alt.(string); ok {
		if altStr := Apply(alt, strings.TrimSpace); altStr != "" {
			result = result.Set(AttrImageAlt, altStr)
		}
	}
	if width, ok := width.(string); ok {
		if widthStr := Apply(width, strings.TrimSpace); widthStr != "" {
			result = result.Set(AttrImageWidth, widthStr)
		}
	}
	if height, ok := height.(string); ok {
		if heightStr := Apply(height, strings.TrimSpace); heightStr != "" {
			result = result.Set(AttrImageHeight, heightStr)
		}
	}
	for _, otherAttr := range otherattrs {
		if otherAttr, ok := otherAttr.(Attributes); ok {
			for k, v := range otherAttr {
				result = result.Set(k, v)
				if k == AttrID {
					// mark custom_id flag to `true`
					result = result.Set(AttrCustomID, true)
				}
			}
		}
	}
	return result, nil
}

// ------------------------------------------
// Footnotes
// ------------------------------------------

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
func NewFootnote(ref string, elements interface{}) (Footnote, error) {
	// footnote with content get an ID
	if elements, ok := elements.([]interface{}); ok {
		return Footnote{
			// ID is only set during document processing
			Ref:      ref,
			Elements: elements,
		}, nil
	} // footnote which are just references don't get an ID, so we don't increment the sequence
	return Footnote{
		Ref:      ref,
		Elements: []interface{}{},
	}, nil
}

// FootnoteReference a footnote reference. Replaces the actual footnote in the document,
// and only contains a generated, sequential ID (which will be displayed)
type FootnoteReference struct {
	ID        int
	Ref       string // the user-specified reference (optional)
	Duplicate bool   // indicates if this reference targets an already-existing footnote // TODO: find a better name?
}

// FootnotesContainer interface for all types which may contain footnotes
type FootnotesContainer interface {
	ReplaceFootnotes(existing *Footnotes) interface{}
}

// Footnotes the footnotes of a document. Footnotes are "collected"
// during the parsing phase and displayed at the bottom of the document
// during the rendering.
type Footnotes struct {
	sequence *sequence
	notes    []Footnote
}

// NewFootnotes initializes a new Footnotes
func NewFootnotes() *Footnotes {
	return &Footnotes{
		sequence: &sequence{},
		notes:    []Footnote{},
	}
}

// IndexOf returns the index of the given note in the footnotes.
func (f *Footnotes) indexOf(actual Footnote) (int, bool) {
	for _, note := range f.notes {
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
func (f *Footnotes) Reference(note Footnote) FootnoteReference {
	ref := FootnoteReference{}
	if len(note.Elements) > 0 {
		note.ID = f.sequence.nextVal()
		f.notes = append(f.notes, note)
		ref.ID = note.ID
	} else if id, found := f.indexOf(note); found {
		ref.ID = id
		ref.Duplicate = true
	} else {
		ref.ID = InvalidFootnoteReference
		logrus.Errorf("no footnote with reference '%s'", note.Ref)
	}
	ref.Ref = note.Ref
	return ref
}

// Notes returns all footnotes
func (f *Footnotes) Notes() []Footnote {
	if len(f.notes) == 0 {
		return nil
	}
	return f.notes
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

// DelimitedBlock the structure for the delimited blocks
type DelimitedBlock struct {
	Kind       BlockKind
	Attributes Attributes
	Elements   []interface{} // TODO: rename to `Blocks`?
}

// NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given elements
func NewDelimitedBlock(kind BlockKind, elements []interface{}, attributes interface{}) (DelimitedBlock, error) {
	log.Debugf("initializing a new DelimitedBlock of kind '%v' with %d elements", kind, len(elements))
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return DelimitedBlock{}, errors.Wrap(err, "failed to initialize a delimited block")
	}
	if k, found := attrs[AttrKind].(BlockKind); found { // override default kind
		log.Debugf("overriding kind '%s' to '%s'", kind, k)
		kind = k
	}
	return DelimitedBlock{
		Attributes: attrs,
		Kind:       kind,
		Elements:   elements,
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
func NewCallout(ref int) (Callout, error) {
	return Callout{
		Ref: ref,
	}, nil
}

// CalloutListItem the description of a call out which will appear as an ordered list item after the delimited block
type CalloutListItem struct {
	Attributes Attributes
	Ref        int
	Elements   []interface{}
}

var _ ListItem = &CalloutListItem{}

var _ DocumentElement = &CalloutListItem{}

// GetAttributes returns the elements of this CalloutListItem
func (i CalloutListItem) GetAttributes() Attributes {
	return i.Attributes
}

// AddElement add an element to this CalloutListItem
func (i *CalloutListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// NewCalloutListItem returns a new CalloutListItem
func NewCalloutListItem(ref int, description []interface{}) (CalloutListItem, error) {
	return CalloutListItem{
		Attributes: nil,
		Ref:        ref,
		Elements:   description,
	}, nil
}

// CalloutList the structure for the Callout Lists
type CalloutList struct {
	Attributes Attributes
	Items      []CalloutListItem
}

var _ List = &CalloutList{}

// NewCalloutList initializes a new CalloutList and uses the given item's attributes as the list attributes
func NewCalloutList(item CalloutListItem) *CalloutList {
	attrs := item.Attributes
	item.Attributes = nil
	return &CalloutList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []CalloutListItem{
			item,
		},
	}
}

// AddItem adds the given item to the list
func (l *CalloutList) AddItem(item CalloutListItem) {
	l.Items = append(l.Items, item)
}

// LastItem returns the last item in the list
func (l *CalloutList) LastItem() ListItem {
	return &(l.Items[len(l.Items)-1])
}

// ------------------------------------------
// Tables
// ------------------------------------------

// Table the structure for the tables
type Table struct {
	Attributes Attributes
	Header     TableLine
	Lines      []TableLine
}

// NewTable initializes a new table with the given lines and attributes
func NewTable(header interface{}, lines []interface{}, attributes interface{}) (Table, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return Table{}, errors.Wrapf(err, "failed to initialize a Table element")
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
	cells := make([][]interface{}, 0)
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
			Cells: make([][]interface{}, columnsPerLine),
		}
		for i, c := range cells {
			log.Debugf("adding cell with content '%v' in table line at offset %d", c, (i % columnsPerLine))
			l.Cells[i%columnsPerLine] = c
			if (i+1)%columnsPerLine == 0 { // switch to next line
				log.Debugf("adding line with content '%v' in table", l)
				t.Lines = append(t.Lines, l)
				l = TableLine{
					Cells: make([][]interface{}, columnsPerLine),
				}
			}
		}
	}
	// log.Debugf("initialized a new table with %d line(s)", len(lines))
	return t, nil
}

// TableLine a table line is made of columns, each column being a group of []interface{} (to support quoted text, etc.)
type TableLine struct {
	Cells [][]interface{}
}

// NewTableLine initializes a new TableLine with the given columns
func NewTableLine(columns []interface{}) (TableLine, error) {
	c := make([][]interface{}, 0)
	for _, column := range columns {
		if e, ok := column.([]interface{}); ok {
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
	Attributes Attributes
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
		return LiteralBlock{}, errors.Wrapf(err, "failed to initialize a new LiteralBlock")
	}
	// log.Debugf("initialized a new LiteralBlock with %d lines", len(lines))
	attrs, err := NewAttributes(Attributes{
		AttrKind:             Literal,
		AttrLiteralBlockType: origin,
	})
	attrs.Add(attributes)
	if err != nil {
		return LiteralBlock{}, errors.Wrapf(err, "failed to initialize a Literal block")
	}
	return LiteralBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
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

// ------------------------------------------
// Comments
// ------------------------------------------

// SingleLineComment a single line comment
type SingleLineComment struct {
	Content string
}

// NewSingleLineComment initializes a new single line content
func NewSingleLineComment(content string) (SingleLineComment, error) {
	log.Debugf("initializing a single line comment with content: '%s'", content)
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

// String return the content of this StringElement
func (s StringElement) String() string {
	return s.Content
}

// ------------------------------------------
// VerbatimLine
// ------------------------------------------

// VerbatimLine the structure for verbatim line, ie, read "as-is" from a given text document.
type VerbatimLine struct {
	Content  string
	Callouts []Callout
}

// NewVerbatimLine initializes a new `VerbatimLine` from the given content
func NewVerbatimLine(content string, callouts interface{}) (VerbatimLine, error) {
	var cos []Callout
	if callouts, ok := callouts.([]interface{}); ok {
		for _, c := range callouts {
			cos = append(cos, c.(Callout))
		}
	}
	return VerbatimLine{
		Content:  content,
		Callouts: cos,
	}, nil
}

// NewVerbatimFileLine initializes a new `VerbatimLine` from the given content in a file
func NewVerbatimFileLine(content string) (VerbatimLine, error) {
	return VerbatimLine{
		Content: content,
	}, nil
}

var emptyStringRE = regexp.MustCompile(` \t`)

// IsEmpty return `true` if the line contains only whitespaces and tabs
func (s VerbatimLine) IsEmpty() bool {
	return len(s.Content) == 0 || emptyStringRE.MatchString(s.Content)
}

// String return the content of this VerbatimLine
func (s VerbatimLine) String() string {
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
	Kind       QuotedTextKind
	Elements   []interface{}
	Attributes Attributes
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
func NewQuotedText(kind QuotedTextKind, attributes interface{}, elements ...interface{}) (QuotedText, error) {
	attrs, err := NewQuotedTextAttributes(attributes)
	if err != nil {
		return QuotedText{}, errors.Wrap(err, "failed to initialize a QuotedText element")
	}
	return QuotedText{
		Kind:       kind,
		Elements:   Merge(elements),
		Attributes: attrs,
	}, nil
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
// InlinePassthrough
// ------------------------------------------

// InlinePassthrough the structure for Passthroughs
type InlinePassthrough struct {
	Kind     PassthroughKind
	Elements []interface{}
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

// NewInlinePassthrough returns a new passthrough
func NewInlinePassthrough(kind PassthroughKind, elements []interface{}) (InlinePassthrough, error) {
	return InlinePassthrough{
		Kind:     kind,
		Elements: Merge(elements...),
	}, nil

}

// ------------------------------------------
// Inline Links
// ------------------------------------------

// InlineLink the structure for the external links
type InlineLink struct {
	Location   Location
	Attributes Attributes
}

// NewInlineLink initializes a new inline `InlineLink`
func NewInlineLink(url Location, attrs interface{}) (InlineLink, error) {
	result := InlineLink{
		Location: url,
	}
	if attrs, ok := attrs.(Attributes); ok && len(attrs) > 0 {
		result.Attributes = attrs
	}
	return result, nil
}

// AttrInlineLinkText the link `text` attribute
const AttrInlineLinkText string = "positional-1"

// NewInlineLinkAttributes returns a map of link attributes
func NewInlineLinkAttributes(attributes []interface{}) (Attributes, error) {
	log.Debugf("new inline link attributes: %v", attributes)
	if len(attributes) == 0 {
		return nil, nil
	}
	result := Attributes{}
	for i, attr := range attributes {
		log.Debugf("new inline link attribute: '%[1]v' (%[1]T)", attr)
		switch attr := attr.(type) {
		case Attributes:
			for k, v := range attr {
				result[k] = v
			}
		case []interface{}:
			result["positional-"+strconv.Itoa(i+1)] = attr
		}
	}
	log.Debugf("new inline link attributes: %v", result)
	return result, nil
}

// ------------------------------------------
// File Inclusions
// ------------------------------------------

// FileInclusion the structure for the file inclusions
type FileInclusion struct {
	Attributes Attributes
	Location   Location
	RawText    string
}

// NewFileInclusion initializes a new inline `InlineLink`
func NewFileInclusion(location Location, attributes interface{}, rawtext string) (FileInclusion, error) {
	attrs, err := NewAttributes(attributes)
	if err != nil {
		return FileInclusion{}, errors.Wrap(err, "failed to initialize a FileInclusion element")
	}
	return FileInclusion{
		Attributes: attrs,
		Location:   location,
		RawText:    rawtext,
	}, nil
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
func NewLineRangesAttribute(ranges interface{}) (Attributes, error) {
	switch ranges := ranges.(type) {
	case []interface{}:
		return Attributes{
			AttrLineRanges: NewLineRanges(ranges...),
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
func NewTagRangesAttribute(ranges []interface{}) (Attributes, error) {
	r, err := NewTagRanges(ranges...)
	if err != nil {
		return nil, err
	}
	log.Debugf("initialized a new TagRanges attribute with values: %v", r)
	return Attributes{
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
func NewLocation(scheme interface{}, path []interface{}) (Location, error) {
	path = Merge(path)
	// log.Debugf("new location: '%[1]s' (%[1]T) '%+[2]v", scheme, path)
	s := ""
	if scheme, ok := scheme.([]byte); ok {
		s = string(scheme)
	}
	return Location{
		Scheme: s,
		Path:   path,
	}, nil
}

// Resolve resolves the Location by replacing all document attribute substitutions
// with their associated values, or their corresponding raw text if
// no attribute matched
// returns the resolved attribute
func (l Location) String() string { // (attrs map[string]string) string {
	result := bytes.NewBuffer(nil)
	result.WriteString(l.Scheme)
	for _, e := range l.Path {
		if s, ok := e.(string); ok {
			result.WriteString(s) // no need to use `fmt.Sprintf` for elements of type 'string'
		} else {
			result.WriteString(fmt.Sprintf("%s", e))
		}
	}
	return result.String()
}

const imagesdir = "imagesdir"

// Resolve resolves the Location by replacing all document attribute substitutions
// with their associated values, or their corresponding raw text if
// no attribute matched
// returns `true` if some document attribute substitution occurred
func (l *Location) Resolve(attrs AttributesWithOverrides) Location {
	content := bytes.NewBuffer(nil)
	for _, e := range l.Path {
		switch e := e.(type) {
		case AttributeSubstitution:
			if value, found := attrs.GetAsString(e.Name); found {
				content.WriteString(value)
			} else {
				content.WriteRune('{')
				content.WriteString(e.Name)
				content.WriteRune('}')
			}
		default:
			if s, ok := e.(string); ok {
				content.WriteString(s) // no need to use `fmt.Sprintf` for elements of type 'string'
			} else {
				content.WriteString(fmt.Sprintf("%s", e))
			}
		}
	}
	location := content.String()
	if l.Scheme == "" && !strings.HasPrefix(location, "/") {
		if u, err := url.Parse(location); err == nil {
			if !u.IsAbs() {
				if imagesdir, ok := attrs.GetAsString(imagesdir); ok {
					location = imagesdir + "/" + location
				}
			}
		}
	}
	return Location{
		Scheme: l.Scheme,
		Path: []interface{}{
			StringElement{
				Content: location,
			},
		},
	}
}

// -------------------------------------------------------------------------------------
// Index terms
// -------------------------------------------------------------------------------------

// IndexTerm a index term, with a single term
type IndexTerm struct {
	Term []interface{}
}

// NewIndexTerm returns a new IndexTerm
func NewIndexTerm(term []interface{}) (IndexTerm, error) {
	return IndexTerm{
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
func NewConcealedIndexTerm(term1, term2, term3 interface{}) (ConcealedIndexTerm, error) {
	return ConcealedIndexTerm{
		Term1: term1,
		Term2: term2,
		Term3: term3,
	}, nil
}

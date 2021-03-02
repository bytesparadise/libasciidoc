package types

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// RawDocument document with a front-matter and raw blocks (will be refined in subsequent processing phases)
type RawDocument struct {
	FrontMatter FrontMatter
	Elements    []interface{}
}

// NewRawDocument initializes a new `RawDocument` from the given lines
func NewRawDocument(frontMatter interface{}, elements []interface{}) (RawDocument, error) {
	// log.Debugf("new RawDocument with %d block element(s)", len(elements))
	result := RawDocument{
		Elements: elements,
	}
	if fm, ok := frontMatter.(FrontMatter); ok {
		result.FrontMatter = fm
	}
	return result, nil
}

// Attributes returns the document attributes on the top-level section
// and all the document attribute declarations at the top of the document only.
func (d RawDocument) Attributes() Attributes {
	result := Attributes{}
elements:
	for _, b := range d.Elements {
		switch b := b.(type) {
		case Section:
			if b.Level == 0 {
				// also, expand document authors and revision
				if authors, ok := b.Attributes[AttrAuthors].([]DocumentAuthor); ok {
					// move to the Document attributes
					result.SetAll(expandAuthors(authors))
					delete(b.Attributes, AttrAuthors)
				}
				// also, expand document authors and revision
				if revision, ok := b.Attributes[AttrRevision].(DocumentRevision); ok {
					// move to the Document attributes
					result.SetAll(expandRevision(revision))
					delete(b.Attributes, AttrRevision)
				}
				continue // allow to continue if the section is level 0
			}
			break elements // otherwise, just stop
		case AttributeDeclaration:
			result.Set(b.Name, b.Value)
		default:
			break elements
		}
	}
	// log.Debugf("document attributes: %+v", result)
	return result
}

// RawSection a document section (when processing file inclusions)
// We only care about the level here
type RawSection struct {
	Level   int
	Title   string
	RawText string
}

// NewRawSection returns a new RawSection
func NewRawSection(level int, title string) (RawSection, error) {
	// log.Debugf("new rawsection: '%s' (%d)", title, level)
	return RawSection{
		Level: level,
		Title: title,
	}, nil
}

var _ Stringer = RawSection{}

// Stringify returns the string representation of this section, as it existed in the source document
func (s RawSection) Stringify() string {
	return strings.Repeat("=", s.Level+1) + " " + s.Title
}

// ------------------------------------------
// Lines
// ------------------------------------------

// NewRawLine returns a new slice containing a single StringElement with the given content
func NewRawLine(content string) ([]interface{}, error) {
	// log.Debugf("new line: '%v'", content)
	return []interface{}{
		StringElement{
			Content: content,
		},
	}, nil
}

// ------------------------------------------
// common interfaces
// ------------------------------------------

// Stringer a type which can be serializes as a string
type Stringer interface {
	Stringify() string
}

// WithPlaceholdersInElements interface for all blocks in which elements can
// be substituted with placeholders while applying the substitutions
type WithPlaceholdersInElements interface {
	RestoreElements(placeholders map[string]interface{}) interface{}
}

// WithPlaceholdersInAttributes interface for all blocks in which attribute content can
// be substituted with placeholders while applying the substitutions
type WithPlaceholdersInAttributes interface {
	RestoreAttributes(placeholders map[string]interface{}) interface{}
}

// WithPlaceholdersInLocation interface for all blocks in which location elements can
// be substituted with placeholders while applying the substitutions
type WithPlaceholdersInLocation interface {
	RestoreLocation(placeholders map[string]interface{}) interface{}
}

// RawText interface for the elements that can provide the raw text representation of this element
// as it was (supposedly) written in the source document
type RawText interface {
	RawText() (string, error)
}

// ------------------------------------------
// Substitution support
// ------------------------------------------

// WithCustomSubstitutions base interface for types on which custom substitutions apply
type WithCustomSubstitutions interface {
	SubstitutionsToApply() []string
	DefaultSubstitutions() []string
}

// WithAttributesToSubstitute base interface for types on which attributes can be substituted
type WithAttributesToSubstitute interface {
	AttributesToSubstitute() Attributes
	ReplaceAttributes(Attributes) interface{}
}

// WithElementsToSubstitute interface for types on which elements can be substituted
type WithElementsToSubstitute interface {
	ElementsToSubstitute() []interface{}
	ReplaceElements([]interface{}) interface{}
}

// WithLineSubstitution interface for types on which lines can be substituted
type WithLineSubstitution interface {
	WithCustomSubstitutions
	LinesToSubstitute() [][]interface{}
	SubstituteLines([][]interface{}) interface{}
}

// WithNestedElementSubstitution a block in which nested elements can be substituted
type WithNestedElementSubstitution interface {
	WithCustomSubstitutions
	WithElementsToSubstitute
}

var defaultSubstitutionsForBlockElements = []string{
	"inline_passthrough",
	"specialcharacters",
	"quotes",
	"attributes",
	"replacements",
	"macros",
	"post_replacements",
}
var defaultExampleBlockSubstitutions = defaultSubstitutionsForBlockElements
var defaultQuoteBlockSubstitutions = defaultSubstitutionsForBlockElements
var defaultSidebarBlockSubstitutions = defaultSubstitutionsForBlockElements
var defaultVerseBlockSubstitutions = defaultSubstitutionsForBlockElements // even though it's a block of lines, not a block of blocks
var defaultParagraphSubstitutions = defaultSubstitutionsForBlockElements  // even though it's a block of lines, not a block of blocks

// blocks of lines
var defaultSubstitutionsForBlockLines = []string{
	"callouts", // must be executed before "specialcharacters"
	"specialcharacters",
}
var defaultFencedBlockSubstitutions = defaultSubstitutionsForBlockLines
var defaultListingBlockSubstitutions = defaultSubstitutionsForBlockLines
var defaultLiteralBlockSubstitutions = defaultSubstitutionsForBlockLines

// other blocks
var defaultPassthroughBlockSubstitutions = []string{}

// ------------------------------------------
// Draft Document: document in which
// all substitutions have been applied
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
	// log.Debugf("new array of document authors from `%+v`", authors)
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
	Value interface{}
}

// NewAttributeDeclaration initializes a new AttributeDeclaration with the given name and optional value
func NewAttributeDeclaration(name string, value interface{}) (AttributeDeclaration, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		// log.Debugf("new AttributeDeclaration: '%s'", name)
		spew.Fdump(log.StandardLogger().Out, value)
	}
	return AttributeDeclaration{
		Name:  name,
		Value: value,
	}, nil
}

var _ Stringer = AttributeDeclaration{}

// Stringify returns the string representation of this attribute declaration, as it existed in the source document
func (a AttributeDeclaration) Stringify() string {
	result := strings.Builder{}
	result.WriteString(":" + a.Name + ":")
	switch v := a.Value.(type) {
	case string:
		result.WriteString(" " + v)
	case []interface{}:
		result.WriteString(" ")
		for _, e := range v {
			switch e := e.(type) {
			case StringElement:
				result.WriteString(e.Content)
			case AttributeSubstitution:
				result.WriteString("{" + e.Name + "}")
			}
		}
	}
	return result.String()
}

var _ WithElementsToSubstitute = AttributeDeclaration{}

// ElementsToSubstitute returns this section's title so that substitutions can be applied onto its elements
func (a AttributeDeclaration) ElementsToSubstitute() []interface{} {
	switch v := a.Value.(type) {
	case []interface{}:
		return v
	default:
		return []interface{}{v}
	}
}

// ReplaceElements replaces the elements in this section
func (a AttributeDeclaration) ReplaceElements(value []interface{}) interface{} {
	a.Value = Reduce(value, strings.TrimSpace)
	return a
}

// AttributeReset the type for AttributeReset
type AttributeReset struct {
	Name string
}

// NewAttributeReset initializes a new Document Attribute Resets.
func NewAttributeReset(attrName string) (AttributeReset, error) {
	// log.Debugf("new AttributeReset: '%s'", attrName)
	return AttributeReset{Name: attrName}, nil
}

// AttributeSubstitution the type for AttributeSubstitution
type AttributeSubstitution struct {
	Name string
}

var _ RawText = AttributeSubstitution{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (s AttributeSubstitution) RawText() (string, error) {
	return "{" + s.Name + "}", nil
}

// PredefinedAttribute a special kind of attribute substitution, which
// uses a predefined attribute
type PredefinedAttribute AttributeSubstitution

// NewAttributeSubstitution initializes a new Attribute Substitutions
func NewAttributeSubstitution(name string) (interface{}, error) {
	if isPrefedinedAttribute(name) {
		return PredefinedAttribute{Name: name}, nil
	}
	// log.Debugf("new AttributeSubstitution: '%s'", name)
	return AttributeSubstitution{Name: name}, nil
}

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
// Element kinds
// ------------------------------------------

const (
	// Fenced a fenced block
	Fenced = "fenced"
	// Listing a listing block
	Listing = "listing"
	// Example an example block
	Example = "example"
	// Comment a comment block
	Comment = "comment"
	// Quote a quote block
	Quote = "quote"
	// MarkdownQuote a quote block in the Markdown style
	MarkdownQuote = "markdown-quote"
	// Verse a verse block
	Verse = "verse"
	// Sidebar a sidebar block
	Sidebar = "sidebar"
	// Literal a literal block
	Literal = "literal"
	// Source a source block
	Source = "source"
	// Passthrough a passthrough block
	Passthrough = "pass"

	// AttrSourceBlockOption the option set on a source block, using the `source%<option>` attribute
	AttrSourceBlockOption = "source-option" // DEPRECATED
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
	return UserMacro{
		Name:       name,
		Kind:       BlockMacro,
		Value:      value,
		Attributes: toAttributes(attributes),
		RawText:    raw,
	}, nil
}

// NewInlineUserMacro returns an UserMacro
func NewInlineUserMacro(name, value string, attributes interface{}, raw string) (UserMacro, error) {
	return UserMacro{
		Name:       name,
		Kind:       InlineMacro,
		Value:      value,
		Attributes: toAttributes(attributes),
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
	// log.Debugf("new FrontMatter with attributes: %+v", attributes)
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
	attrs := toAttributes(attributes)
	// multiple IDs can be defined (by mistake), but only the last one is used
	attrs = attrs.SetAll(ids)
	// also, set the `AttrCustomID` flag if an ID was set
	if _, exists := attrs[AttrID]; exists {
		attrs[AttrCustomID] = true
	}
	return Section{
		Level:      level,
		Attributes: attrs,
		Title:      title,
		Elements:   []interface{}{},
	}, nil
}

var _ WithElementsToSubstitute = Section{}

// ElementsToSubstitute returns this section's title so that substitutions can be applied onto its elements
func (s Section) ElementsToSubstitute() []interface{} {
	return s.Title
}

// ReplaceElements replaces the elements in this section
func (s Section) ReplaceElements(title []interface{}) interface{} {
	s.Title = title
	return s
}

var _ WithAttributesToSubstitute = Section{}

// AttributesToSubstitute returns this section's attributes
func (s Section) AttributesToSubstitute() Attributes {
	return s.Attributes
}

// ReplaceAttributes replaces the attributes in this section
func (s Section) ReplaceAttributes(attributes Attributes) interface{} {
	s.Attributes = attributes
	return s
}

// ResolveID resolves/updates the "ID" attribute in the section (in case the title changed after some document attr substitution)
func (s Section) ResolveID(docAttributes AttributesWithOverrides) (Section, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		// log.Debugf("section attributes:")
		spew.Fdump(log.StandardLogger().Out, s.Attributes)
	}

	if !s.Attributes.Has(AttrID) {
		// log.Debugf("resolving section id")
		separator := docAttributes.GetAsStringWithDefault(AttrIDSeparator, DefaultIDSeparator)
		replacement, err := ReplaceNonAlphanumerics(s.Title, separator)
		if err != nil {
			return s, errors.Wrapf(err, "failed to generate default ID on Section element")
		}
		idPrefix := docAttributes.GetAsStringWithDefault(AttrIDPrefix, DefaultIDPrefix)
		s.Attributes = s.Attributes.Set(AttrID, idPrefix+replacement)
		// log.Debugf("updated section id to '%s'", s.Attributes[AttrID])
	}
	return s, nil
}

// AddElement adds the given child element to this section
func (s *Section) AddElement(e interface{}) {
	s.Elements = append(s.Elements, e)
}

var _ FootnotesContainer = Section{}

// SubstituteFootnotes replaces the footnotes in the section title
// with footnote references. The footnotes are stored in the given 'notes' param
func (s Section) SubstituteFootnotes(notes *Footnotes) interface{} {
	for i, element := range s.Title {
		if note, ok := element.(Footnote); ok {
			s.Title[i] = notes.Reference(note)
		}
	}
	return s
}

// NewDocumentHeader initializes a new Section with level 0 which can have authors and a revision, among other attributes
func NewDocumentHeader(title []interface{}, authors interface{}, revision interface{}) (Section, error) {
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
	s := make([]DocumentAuthor, len(authors))
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
		// also include a "string" version of the given author
		s[i] = DocumentAuthor{
			FullName: result[key("author", i)].(string),
			Email:    email,
		}
	}
	result[AttrAuthors] = s
	// log.Debugf("authors: %v", result)
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
	// log.Debugf("revision: %v", result)
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
	// log.Debugf("new continued list element for element of type %T", element)
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

const (
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
			attributes[AttrStyle] = "upperalpha"
			delete(attributes, k)
		case "upperroman":
			attributes[AttrStyle] = "upperroman"
			delete(attributes, k)
		case "lowerroman":
			attributes[AttrStyle] = "lowerroman"
			delete(attributes, k)
		case "loweralpha":
			attributes[AttrStyle] = "loweralpha"
			delete(attributes, k)
		case "arabic":
			attributes[AttrStyle] = "arabic"
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
	Attributes Attributes
	Level      int
	Style      string
	Elements   []interface{} // TODO: rename to `Blocks`?
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}, attributes interface{}) (OrderedListItem, error) {
	// log.Debugf("new OrderedListItem")
	return OrderedListItem{
		Attributes: toAttributesWithMapping(attributes, map[string]string{AttrPositional1: AttrStyle}),
		Style:      prefix.Style,
		Level:      prefix.Level,
		Elements:   elements,
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

var _ WithElementsToSubstitute = OrderedListItem{}

// ElementsToSubstitute returns this item's elements so that substitutions can be applied onto them
func (i OrderedListItem) ElementsToSubstitute() []interface{} {
	return i.Elements
}

// ReplaceElements replaces the elements in this example block
func (i OrderedListItem) ReplaceElements(elements []interface{}) interface{} {
	i.Elements = elements
	return i
}

var _ WithAttributesToSubstitute = OrderedListItem{}

// AttributesToSubstitute returns this list item's attributes
func (i OrderedListItem) AttributesToSubstitute() Attributes {
	return i.Attributes
}

// ReplaceAttributes replaces the attributes in this list item
func (i OrderedListItem) ReplaceAttributes(attributes Attributes) interface{} {
	i.Attributes = attributes
	return i
}

// OrderedListItemPrefix the prefix used to construct an OrderedListItem
type OrderedListItemPrefix struct {
	Style string
	Level int
}

// NewOrderedListItemPrefix initializes a new OrderedListItemPrefix
func NewOrderedListItemPrefix(s string, l int) (OrderedListItemPrefix, error) {
	return OrderedListItemPrefix{
		Style: s,
		Level: l,
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
func NewUnorderedList(item UnorderedListItem) *UnorderedList {
	attrs := item.Attributes
	item.Attributes = nil // move the item's attributes to the list level
	// convert the checkstyle attribute if the list is interactive
	// log.Debugf("interactive list: %t", attrs.HasOption(AttrInteractive))
	if attrs.HasOption(AttrInteractive) {
		item = item.toInteractiveListItem()
	}
	list := &UnorderedList{
		Attributes: attrs, // move the item's attributes to the list level
		Items: []UnorderedListItem{
			item,
		},
	}
	return list
}

// AddItem adds the given item
func (l *UnorderedList) AddItem(item UnorderedListItem) {
	if l.Attributes.HasOption(AttrInteractive) {
		item = item.toInteractiveListItem()
	}
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
	// log.Debugf("new UnorderedListItem with %d elements", len(elements))
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
		Attributes:  toAttributesWithMapping(attributes, map[string]string{AttrPositional1: AttrStyle}),
		BulletStyle: prefix.BulletStyle,
		CheckStyle:  cs,
		Elements:    elements,
	}, nil
}

func (i UnorderedListItem) toInteractiveListItem() UnorderedListItem {
	i.CheckStyle = i.CheckStyle.toInteractive()
	if i.CheckStyle != NoCheck && len(i.Elements) > 0 {
		if p, ok := i.Elements[0].(Paragraph); ok {
			if p.Attributes == nil {
				p.Attributes = Attributes{}
				i.Elements[0] = p // need to update the element in the slice
			}
			p.Attributes[AttrCheckStyle] = i.CheckStyle
		}
	}
	return i
}

// GetAttributes returns the elements of this UnorderedListItem
func (i UnorderedListItem) GetAttributes() Attributes {
	return i.Attributes
}

// AddElement add an element to this UnorderedListItem
func (i *UnorderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

var _ WithElementsToSubstitute = UnorderedListItem{}

// ElementsToSubstitute returns this item's elements so that substitutions can be applied onto them
func (i UnorderedListItem) ElementsToSubstitute() []interface{} {
	return i.Elements
}

// ReplaceElements replaces the elements in this example block
func (i UnorderedListItem) ReplaceElements(elements []interface{}) interface{} {
	i.Elements = elements
	return i
}

var _ WithAttributesToSubstitute = UnorderedListItem{}

// AttributesToSubstitute returns this list item's attributes
func (i UnorderedListItem) AttributesToSubstitute() Attributes {
	return i.Attributes
}

// ReplaceAttributes replaces the attributes in this list item
func (i UnorderedListItem) ReplaceAttributes(attributes Attributes) interface{} {
	i.Attributes = attributes
	return i
}

// UnorderedListItemCheckStyle the check style that applies on an unordered list item
type UnorderedListItemCheckStyle string

const (
	// Checked when the unordered list item is checked
	Checked UnorderedListItemCheckStyle = "checked"
	// CheckedInteractive when the unordered list item is checked (with an interactive checkbox)
	CheckedInteractive UnorderedListItemCheckStyle = "checked-interactive"
	// Unchecked when the unordered list item is not checked
	Unchecked UnorderedListItemCheckStyle = "unchecked"
	// UncheckedInteractive when the unordered list item is not checked (with an interactive checkbox)
	UncheckedInteractive UnorderedListItemCheckStyle = "unchecked-interactive"
	// NoCheck when the unodered list item has no specific check annotation
	NoCheck UnorderedListItemCheckStyle = "nocheck"
)

func (s UnorderedListItemCheckStyle) toInteractive() UnorderedListItemCheckStyle {
	switch s {
	case Checked, CheckedInteractive:
		return CheckedInteractive
	case Unchecked, UncheckedInteractive:
		return UncheckedInteractive
	default:
		return NoCheck
	}
}

func toCheckStyle(checkstyle interface{}) UnorderedListItemCheckStyle {
	if cs, ok := checkstyle.(UnorderedListItemCheckStyle); ok {
		return cs
	}
	return NoCheck
}

// BulletStyle the type of bullet for items in an unordered list
type BulletStyle string

const (
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
	// log.Debugf("new ListItemContent with %d line(s)", len(content))
	elements := make([]interface{}, 0, len(content))
	for _, element := range content {
		// log.Debugf("Processing line element of type %T", element)
		switch element := element.(type) {
		case []interface{}:
			elements = append(elements, element...)
		case interface{}:
			elements = append(elements, element)
		}
	}
	// log.Debugf("new ListItemContent with %d elements(s)", len(elements))
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
	attrs := item.Attributes
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
	// log.Debugf("new LabeledListItem")
	var elements []interface{}
	if description, ok := description.([]interface{}); ok {
		elements = description
	} else {
		elements = []interface{}{}
	}
	return LabeledListItem{
		Attributes: toAttributesWithMapping(attributes, map[string]string{AttrPositional1: AttrStyle}),
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

var _ WithElementsToSubstitute = LabeledListItem{}

// ElementsToSubstitute returns this item's elements so that substitutions can be applied onto them
func (i LabeledListItem) ElementsToSubstitute() []interface{} {
	return i.Elements
}

// ReplaceElements replaces the elements in this example block
func (i LabeledListItem) ReplaceElements(elements []interface{}) interface{} {
	i.Elements = elements
	return i
}

var _ WithAttributesToSubstitute = LabeledListItem{}

// AttributesToSubstitute returns this list item's attributes
func (i LabeledListItem) AttributesToSubstitute() Attributes {
	return i.Attributes
}

// ReplaceAttributes replaces the attributes in this list item
func (i LabeledListItem) ReplaceAttributes(attributes Attributes) interface{} {
	i.Attributes = attributes
	return i
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
// TODO: remove?
const AttrHardBreaks = "hardbreaks"

// DocumentAttrHardBreaks the attribute to set at the document level to render with hard breaks on each line of all paragraphs
const DocumentAttrHardBreaks = "hardbreaks"

// NewParagraph initializes a new `Paragraph`
func NewParagraph(lines []interface{}, attributes interface{}) (Paragraph, error) {
	// log.Debugf("new paragraph with attributes: '%v'", attributes)
	l, err := toLines(lines)
	if err != nil {
		return Paragraph{}, errors.Wrapf(err, "failed to initialize a Paragraph")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	switch attrs.GetAsStringWithDefault(AttrStyle, "") {
	case string(Source):
		attrs = toAttributesWithMapping(attrs, map[string]string{
			AttrPositional2: AttrLanguage,
		})
	case string(Quote), string(Verse):
		attrs = toAttributesWithMapping(attrs, map[string]string{
			AttrPositional2: AttrQuoteAuthor,
			AttrPositional3: AttrQuoteTitle,
		})
	}
	return Paragraph{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

func toLines(lines []interface{}) ([][]interface{}, error) {
	result := make([][]interface{}, len(lines))
	for i, line := range lines {
		switch line := line.(type) {
		case []interface{}:
			result[i] = line
		case SingleLineComment:
			result[i] = []interface{}{line}
		default:
			return nil, fmt.Errorf("unexpected type of line: '%T'", line)
		}
	}
	return result, nil
}

var _ WithAttributesToSubstitute = Paragraph{}

// AttributesToSubstitute returns the attributes of this paragraph so that substitutions can be applied onto them
func (p Paragraph) AttributesToSubstitute() Attributes {
	return p.Attributes
}

// ReplaceAttributes replaces the attributes in this paragraph
func (p Paragraph) ReplaceAttributes(attributes Attributes) interface{} {
	p.Attributes = attributes
	return p
}

var _ WithLineSubstitution = Paragraph{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (p Paragraph) SubstitutionsToApply() []string {
	if subs, found := p.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return p.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (p Paragraph) DefaultSubstitutions() []string {
	// support for masquerading
	// treat 'Listing' paragraphs as verbatim blocks
	if p.Attributes.GetAsStringWithDefault(AttrStyle, "") == string(Listing) {
		return defaultListingBlockSubstitutions
	}
	return defaultParagraphSubstitutions
}

// LinesToSubstitute returns the lines of this paragraph so that substitutions can be applied onto them
func (p Paragraph) LinesToSubstitute() [][]interface{} {
	return p.Lines
}

// SubstituteLines replaces the elements in this paragraph
func (p Paragraph) SubstituteLines(lines [][]interface{}) interface{} {
	p.Lines = lines
	return p
}

var _ WithPlaceholdersInAttributes = Paragraph{}

// RestoreAttributes restores the attributes which had been substituted by placeholders
func (p Paragraph) RestoreAttributes(placeholders map[string]interface{}) interface{} {
	p.Attributes = restoreAttributes(p.Attributes, placeholders)
	return p
}

var _ WithPlaceholdersInElements = Paragraph{}

// RestoreElements restores the elements which had been substituted by placeholders
func (p Paragraph) RestoreElements(placeholders map[string]interface{}) interface{} {
	for i, line := range p.Lines {
		p.Lines[i] = restoreElements(line, placeholders)
	}
	return p
}

var _ FootnotesContainer = Paragraph{}

// SubstituteFootnotes replaces the footnotes in the paragraph lines
// with footnote references. The footnotes are stored in the given 'notes' param
func (p Paragraph) SubstituteFootnotes(notes *Footnotes) interface{} {
	for i, line := range p.Lines {
		for j, element := range line {
			if note, ok := element.(Footnote); ok {
				line[j] = notes.Reference(note)
			}
		}
		p.Lines[i] = line
	}
	return p
}

// NewAdmonitionParagraph returns a new Paragraph with an extra admonition attribute
func NewAdmonitionParagraph(lines []interface{}, admonitionKind string, attributes interface{}) (Paragraph, error) {
	// log.Debugf("new admonition paragraph")
	p, err := NewParagraph(lines, attributes)
	if err != nil {
		return Paragraph{}, err
	}
	p.Attributes = p.Attributes.Set(AttrStyle, admonitionKind)
	return p, nil
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

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements ...interface{}) ([]interface{}, error) {
	return Merge(elements...), nil
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
func NewInternalCrossReference(id, label interface{}) (InternalCrossReference, error) {
	// log.Debugf("new InternalCrossReference with ID=%s", id)
	return InternalCrossReference{
		ID:    Reduce(id),
		Label: Reduce(label),
	}, nil
}

// ExternalCrossReference the struct for Cross References
type ExternalCrossReference struct {
	Location Location
	Label    interface{}
}

// NewExternalCrossReference initializes a new `InternalCrossReference` from the given ID
func NewExternalCrossReference(location Location, attributes interface{}) (ExternalCrossReference, error) {
	var label interface{}
	attrs := toAttributes(attributes)
	if l, ok := attrs[AttrPositional1]; ok {
		label = l
	}
	// log.Debugf("new ExternalCrossReference with Location=%v and label='%s' (attrs=%v / %T)", location, label, attributes, attrs[AttrInlineLinkText])
	return ExternalCrossReference{
		Location: location,
		Label:    label,
	}, nil
}

var _ WithPlaceholdersInElements = ExternalCrossReference{}

// RestoreElements restores the elements which had been substituted by placeholders
func (r ExternalCrossReference) RestoreElements(placeholders map[string]interface{}) interface{} {
	if l, ok := r.Label.([]interface{}); ok {
		r.Label = restoreElements(l, placeholders)
	}
	return r
}

var _ WithElementsToSubstitute = ExternalCrossReference{}

// ElementsToSubstitute returns this corss reference location path so that substitutions can be applied onto it
func (r ExternalCrossReference) ElementsToSubstitute() []interface{} {
	return r.Location.Path
}

// ReplaceElements replaces the elements in this example block
func (r ExternalCrossReference) ReplaceElements(path []interface{}) interface{} {
	r.Location.Path = path
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
func NewImageBlock(location Location, inlineAttributes Attributes, attributes interface{}) (ImageBlock, error) {
	// inline attributes trump block attributes
	attrs := toAttributes(inlineAttributes)
	attrs.SetAll(attributes)
	attrs = toAttributesWithMapping(attrs, map[string]string{
		AttrPositional1: AttrImageAlt,
		AttrPositional2: AttrWidth,
		AttrPositional3: AttrHeight,
	})
	return ImageBlock{
		Location:   location,
		Attributes: attrs,
	}, nil
}

var _ WithPlaceholdersInAttributes = ImageBlock{}

// RestoreAttributes restores the attributes which had been substituted by placeholders
func (i ImageBlock) RestoreAttributes(placeholders map[string]interface{}) interface{} {
	i.Attributes = restoreAttributes(i.Attributes, placeholders)
	return i
}

var _ WithPlaceholdersInLocation = ImageBlock{}

// RestoreLocation restores the location elements which had been substituted by placeholders
func (i ImageBlock) RestoreLocation(placeholders map[string]interface{}) interface{} {
	i.Location.Path = restoreElements(i.Location.Path, placeholders)
	return i
}

var _ WithAttributesToSubstitute = ImageBlock{}

// AttributesToSubstitute returns this list item's attributes
func (i ImageBlock) AttributesToSubstitute() Attributes {
	return i.Attributes
}

// ReplaceAttributes replaces the attributes in this list item
func (i ImageBlock) ReplaceAttributes(attributes Attributes) interface{} {
	i.Attributes = attributes
	return i
}

var _ WithElementsToSubstitute = ImageBlock{}

// ElementsToSubstitute returns this image's location path so that substitutions can be applied onto it
func (i ImageBlock) ElementsToSubstitute() []interface{} {
	return i.Location.Path
}

// ReplaceElements replaces the elements in this example block
func (i ImageBlock) ReplaceElements(path []interface{}) interface{} {
	i.Location.Path = path
	return i
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Location   Location
	Attributes Attributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(location Location, attributes interface{}, imagesdir interface{}) (InlineImage, error) {
	location = location.WithPathPrefix(imagesdir)
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrImageAlt,
		AttrPositional2: AttrWidth,
		AttrPositional3: AttrHeight,
	})
	return InlineImage{
		Attributes: attrs,
		Location:   location,
	}, nil
}

var _ WithPlaceholdersInAttributes = InlineImage{}

// RestoreAttributes restores the attributes which had been substituted by placeholders
func (i InlineImage) RestoreAttributes(placeholders map[string]interface{}) interface{} {
	i.Attributes = restoreAttributes(i.Attributes, placeholders)
	return i
}

var _ WithPlaceholdersInLocation = InlineImage{}

// RestoreLocation restores the location elements which had been substituted by placeholders
func (i InlineImage) RestoreLocation(placeholders map[string]interface{}) interface{} {
	i.Location.Path = restoreElements(i.Location.Path, placeholders)
	return i
}

var _ WithAttributesToSubstitute = InlineImage{}

// AttributesToSubstitute returns this inline image's attributes
func (i InlineImage) AttributesToSubstitute() Attributes {
	return i.Attributes
}

// ReplaceAttributes replaces the attributes in this inline image
func (i InlineImage) ReplaceAttributes(attributes Attributes) interface{} {
	i.Attributes = attributes
	return i
}

var _ WithElementsToSubstitute = InlineImage{}

// ElementsToSubstitute returns this inline image location path so that substitutions can be applied onto its elements
func (i InlineImage) ElementsToSubstitute() []interface{} {
	return i.Location.Path // TODO: should return the location so substitution can also take place on the scheme
}

// ReplaceElements replaces the elements in this inline image
func (i InlineImage) ReplaceElements(path []interface{}) interface{} {
	i.Location.Path = path
	return i
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
func NewIcon(class string, attributes interface{}) (Icon, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrIconSize,
	})
	return Icon{
		Class:      class,
		Attributes: attrs,
	}, nil
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

var _ WithPlaceholdersInElements = Footnote{}

// RestoreElements restores the elements which had been substituted by placeholders
func (n Footnote) RestoreElements(placeholders map[string]interface{}) interface{} {
	n.Elements = restoreElements(n.Elements, placeholders)
	return n
}

// FootnoteReference a footnote reference. Substitutes the actual footnote in the document,
// and only contains a generated, sequential ID (which will be displayed)
type FootnoteReference struct {
	ID        int
	Ref       string // the user-specified reference (optional)
	Duplicate bool   // indicates if this reference targets an already-existing footnote // TODO: find a better name?
}

// FootnotesContainer interface for all types which may contain footnotes
type FootnotesContainer interface {
	SubstituteFootnotes(existing *Footnotes) interface{}
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
		logrus.Warnf("no footnote with reference '%s'", note.Ref)
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

// ExampleBlock the structure for the example blocks
type ExampleBlock struct {
	Attributes Attributes
	Elements   []interface{}
}

// NewExampleBlock initializes a new `ExampleBlock` with the given elements
func NewExampleBlock(elements []interface{}, attributes interface{}) (ExampleBlock, error) {
	// log.Debugf("new ExampleBlock with %d blocks", len(elements))
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return ExampleBlock{
		Attributes: attrs,
		Elements:   elements,
	}, nil
}

var _ WithNestedElementSubstitution = ExampleBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b ExampleBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b ExampleBlock) DefaultSubstitutions() []string {
	return defaultExampleBlockSubstitutions
}

// ElementsToSubstitute returns the elements of this example block so that substitutions can be applied onto them
func (b ExampleBlock) ElementsToSubstitute() []interface{} {
	return b.Elements
}

// ReplaceElements replaces the elements in this example block
func (b ExampleBlock) ReplaceElements(elements []interface{}) interface{} {
	b.Elements = elements
	return b
}

var _ WithAttributesToSubstitute = ExampleBlock{}

// AttributesToSubstitute returns this example block's attributes
func (b ExampleBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this example block
func (b ExampleBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// QuoteBlock the structure for the quote blocks
type QuoteBlock struct {
	Attributes Attributes
	Elements   []interface{}
}

// NewQuoteBlock initializes a new `QuoteBlock` with the given elements
func NewQuoteBlock(elements []interface{}, attributes interface{}) (QuoteBlock, error) {
	// log.Debugf("new QuoteBlock with %d blocks", len(elements))
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
		AttrPositional2: AttrQuoteAuthor,
		AttrPositional3: AttrQuoteTitle,
	})
	return QuoteBlock{
		Attributes: attrs,
		Elements:   elements,
	}, nil
}

var _ WithNestedElementSubstitution = QuoteBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b QuoteBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b QuoteBlock) DefaultSubstitutions() []string {
	return defaultQuoteBlockSubstitutions
}

// ElementsToSubstitute returns the elements of this quote block so that substitutions can be applied onto them
func (b QuoteBlock) ElementsToSubstitute() []interface{} {
	return b.Elements
}

// ReplaceElements replaces the elements in this quote block
func (b QuoteBlock) ReplaceElements(elements []interface{}) interface{} {
	b.Elements = elements
	return b
}

var _ WithAttributesToSubstitute = QuoteBlock{}

// AttributesToSubstitute returns this quote block's attributes
func (b QuoteBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this quote block
func (b QuoteBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// SidebarBlock the structure for the example blocks
type SidebarBlock struct {
	Attributes Attributes
	Elements   []interface{}
}

// NewSidebarBlock initializes a new `SidebarBlock` with the given elements
func NewSidebarBlock(elements []interface{}, attributes interface{}) (SidebarBlock, error) {
	// log.Debugf("new SidebarBlock with %d blocks", len(elements))
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return SidebarBlock{
		Attributes: attrs,
		Elements:   elements,
	}, nil
}

var _ WithNestedElementSubstitution = SidebarBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b SidebarBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b SidebarBlock) DefaultSubstitutions() []string {
	return defaultSidebarBlockSubstitutions
}

// ElementsToSubstitute returns the elements of this sidebar block so that substitutions can be applied onto them
func (b SidebarBlock) ElementsToSubstitute() []interface{} {
	return b.Elements
}

// ReplaceElements replaces the elements in this sidebar block
func (b SidebarBlock) ReplaceElements(elements []interface{}) interface{} {
	b.Elements = elements
	return b
}

var _ WithAttributesToSubstitute = SidebarBlock{}

// AttributesToSubstitute returns this sidebar block's attributes
func (b SidebarBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this sidebar block
func (b SidebarBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// FencedBlock the structure for the fenced blocks
type FencedBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewFencedBlock initializes a new `FencedBlock` with the given lines
func NewFencedBlock(lines []interface{}, attributes interface{}) (FencedBlock, error) {
	// log.Debugf("new FencedBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return FencedBlock{}, errors.Wrapf(err, "failed to initialize a new fenced block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return FencedBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

var _ WithLineSubstitution = FencedBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b FencedBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b FencedBlock) DefaultSubstitutions() []string {
	return defaultFencedBlockSubstitutions
}

// LinesToSubstitute returns the lines of this fenced block so that substitutions can be applied onto them
func (b FencedBlock) LinesToSubstitute() [][]interface{} {
	return b.Lines
}

// SubstituteLines replaces the elements in this fenced block
func (b FencedBlock) SubstituteLines(lines [][]interface{}) interface{} {
	b.Lines = lines
	return b
}

var _ WithAttributesToSubstitute = FencedBlock{}

// AttributesToSubstitute returns this fenced block's attributes
func (b FencedBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this fenced block
func (b FencedBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// ListingBlock the structure for the Listing blocks
type ListingBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewListingBlock initializes a new `ListingBlock` with the given lines
func NewListingBlock(lines []interface{}, attributes interface{}) (ListingBlock, error) {
	// log.Debugf("new ListingBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return ListingBlock{}, errors.Wrapf(err, "failed to initialize a new listing block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	if style, ok := attrs.GetAsString(AttrStyle); ok && style == AttrSource {
		attrs = toAttributesWithMapping(attributes, map[string]string{
			AttrPositional1: AttrStyle,
			AttrPositional2: AttrLanguage,
			AttrPositional3: AttrLineNums,
		})
	}
	return ListingBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

var _ WithLineSubstitution = ListingBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b ListingBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b ListingBlock) DefaultSubstitutions() []string {
	return defaultListingBlockSubstitutions
}

// LinesToSubstitute returns the lines of this listing block so that substitutions can be applied onto them
func (b ListingBlock) LinesToSubstitute() [][]interface{} {
	return b.Lines
}

// SubstituteLines replaces the elements in this listing block
func (b ListingBlock) SubstituteLines(lines [][]interface{}) interface{} {
	b.Lines = lines
	return b
}

var _ WithAttributesToSubstitute = ListingBlock{}

// AttributesToSubstitute returns this listing block's attributes
func (b ListingBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this listing block
func (b ListingBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// VerseBlock the structure for the Listing blocks
type VerseBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewVerseBlock initializes a new `VerseBlock` with the given lines
func NewVerseBlock(lines []interface{}, attributes interface{}) (VerseBlock, error) {
	// log.Debugf("new VerseBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return VerseBlock{}, errors.Wrapf(err, "failed to initialize a new verse block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
		AttrPositional2: AttrQuoteAuthor,
		AttrPositional3: AttrQuoteTitle,
	})
	return VerseBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

var _ WithLineSubstitution = VerseBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b VerseBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b VerseBlock) DefaultSubstitutions() []string {
	return defaultVerseBlockSubstitutions
}

// LinesToSubstitute returns the lines of this verse block so that substitutions can be applied onto them
func (b VerseBlock) LinesToSubstitute() [][]interface{} {
	return b.Lines
}

// SubstituteLines replaces the elements in this verse block
func (b VerseBlock) SubstituteLines(lines [][]interface{}) interface{} {
	b.Lines = lines
	return b
}

var _ WithAttributesToSubstitute = VerseBlock{}

// AttributesToSubstitute returns this verse block's attributes
func (b VerseBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this verse block
func (b VerseBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// MarkdownQuoteBlock the structure for the markdown quote blocks
type MarkdownQuoteBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewMarkdownQuoteBlock initializes a new `MarkdownQuoteBlock` with the given lines
func NewMarkdownQuoteBlock(lines []interface{}, attributes interface{}) (MarkdownQuoteBlock, error) {
	// log.Debugf("new MarkdownQuoteBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return MarkdownQuoteBlock{}, errors.Wrapf(err, "failed to initialize a new markdown quote block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return MarkdownQuoteBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

// PassthroughBlock the structure for the comment blocks
type PassthroughBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewPassthroughBlock initializes a new `PassthroughBlock` with the given lines
func NewPassthroughBlock(lines []interface{}, attributes interface{}) (PassthroughBlock, error) {
	// log.Debugf("new PassthroughBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return PassthroughBlock{}, errors.Wrapf(err, "failed to initialize a new passthrough block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return PassthroughBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

var _ WithLineSubstitution = PassthroughBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b PassthroughBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b PassthroughBlock) DefaultSubstitutions() []string {
	return defaultPassthroughBlockSubstitutions
}

// LinesToSubstitute returns the lines of this passthrough block so that substitutions can be applied onto them
func (b PassthroughBlock) LinesToSubstitute() [][]interface{} {
	return b.Lines
}

// SubstituteLines replaces the elements in this passthrough block
func (b PassthroughBlock) SubstituteLines(lines [][]interface{}) interface{} {
	b.Lines = lines
	return b
}

var _ WithAttributesToSubstitute = PassthroughBlock{}

// AttributesToSubstitute returns this passthrough block's attributes
func (b PassthroughBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this passthrough block
func (b PassthroughBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// CommentBlock the structure for the comment blocks
type CommentBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
}

// NewCommentBlock initializes a new `CommentBlock` with the given lines
func NewCommentBlock(lines []interface{}, attributes interface{}) (CommentBlock, error) {
	// log.Debugf("new CommentBlock with %d lines", len(lines))
	l, err := toLines(lines)
	if err != nil {
		return CommentBlock{}, errors.Wrapf(err, "failed to initialize a new comment block")
	}
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	return CommentBlock{
		Attributes: attrs, // TODO: should we expect attributes on comment blocks??
		Lines:      l,
	}, nil
}

// LiteralBlock the structure for the literal blocks
type LiteralBlock struct {
	Attributes Attributes
	Lines      [][]interface{}
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

// NewLiteralBlock initializes a new `LiteralBlock` of the given kind with the given content,
// along with the given sectionTitle spaces
func NewLiteralBlock(origin string, lines []interface{}, attributes interface{}) (LiteralBlock, error) {
	// log.Debugf("new LiteralBlock with %d lines", len(lines))
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrStyle,
	})
	// make sure these 2 attributes are set, even if there was no other attrs (above)
	attrs = attrs.Set(AttrStyle, Literal)
	attrs = attrs.Set(AttrLiteralBlockType, origin)
	l, err := toLines(lines)
	if err != nil {
		return LiteralBlock{}, errors.Wrapf(err, "failed to initialize a new literal block")
	}
	return LiteralBlock{
		Attributes: attrs,
		Lines:      l,
	}, nil
}

var _ WithLineSubstitution = LiteralBlock{}

// SubstitutionsToApply returns the name of the substitutions to apply
func (b LiteralBlock) SubstitutionsToApply() []string {
	if subs, found := b.Attributes.GetAsString(AttrSubstitutions); found {
		return strings.Split(subs, ",")
	}
	return b.DefaultSubstitutions()
}

// DefaultSubstitutions the default substitutions for the paragraph
func (b LiteralBlock) DefaultSubstitutions() []string {
	return defaultLiteralBlockSubstitutions
}

// AttributesToSubstitute returns this literal block's attributes
func (b LiteralBlock) AttributesToSubstitute() Attributes {
	return b.Attributes
}

// ReplaceAttributes replaces the attributes in this literal block
func (b LiteralBlock) ReplaceAttributes(attributes Attributes) interface{} {
	b.Attributes = attributes
	return b
}

// LinesToSubstitute returns the lines of this literal block so that substitutions can be applied onto them
func (b LiteralBlock) LinesToSubstitute() [][]interface{} {
	return b.Lines
}

// SubstituteLines replaces the elements in this literal block
func (b LiteralBlock) SubstituteLines(lines [][]interface{}) interface{} {
	b.Lines = lines
	return b
}

// ------------------------------------------
// Thematic breaks
// ------------------------------------------

// ThematicBreak a thematic break
type ThematicBreak struct{}

// NewThematicBreak returns a new ThematicBreak
func NewThematicBreak() (ThematicBreak, error) {
	return ThematicBreak{}, nil
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
// BlankLine
// ------------------------------------------

// BlankLine the structure for the empty lines, which are used to separate logical blocks
type BlankLine struct {
}

// NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (BlankLine, error) {
	// log.Debug("new BlankLine")
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
func NewQuotedText(kind QuotedTextKind, elements ...interface{}) (QuotedText, error) {
	return QuotedText{
		Kind:     kind,
		Elements: Merge(elements),
	}, nil
}

var _ RawText = QuotedText{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (t QuotedText) RawText() (string, error) {
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
			return "", fmt.Errorf("element of type '%T' cannot be converted to raw text", e)
		}
		s, err := r.RawText()
		if err != nil {
			return "", err
		}
		result.WriteString(s)
	}
	return result.String(), nil
}

// WithAttributes returns a _new_ QuotedText with the given attributes (with some mapping)
func (t QuotedText) WithAttributes(attributes interface{}) (QuotedText, error) {
	log.Debugf("adding attributes on quoted text")
	t.Attributes = toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrRoles,
	})
	return t, nil
}

var _ WithPlaceholdersInAttributes = QuotedText{}

// RestoreAttributes restores the attributes which had been substituted by placeholders
func (t QuotedText) RestoreAttributes(placeholders map[string]interface{}) interface{} {
	t.Attributes = restoreAttributes(t.Attributes, placeholders)
	return t
}

var _ WithPlaceholdersInElements = QuotedText{}

// RestoreElements restores the elements which had been substituted by placeholders
func (t QuotedText) RestoreElements(placeholders map[string]interface{}) interface{} {
	t.Elements = restoreElements(t.Elements, placeholders)
	return t
}

var _ WithElementsToSubstitute = OrderedListItem{}

// ElementsToSubstitute returns this quoted text elements so that substitutions can be applied onto then
func (t QuotedText) ElementsToSubstitute() []interface{} {
	return t.Elements
}

// ReplaceElements replaces the elements in this example block
func (t QuotedText) ReplaceElements(elements []interface{}) interface{} {
	t.Elements = elements
	return t
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
func NewQuotedString(kind QuotedStringKind, elements []interface{}) (QuotedString, error) {
	return QuotedString{Kind: kind, Elements: elements}, nil
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

var _ WithPlaceholdersInElements = QuotedString{}

// RestoreElements restores the elements which had been substituted by placeholders
func (s QuotedString) RestoreElements(placeholders map[string]interface{}) interface{} {
	s.Elements = restoreElements(s.Elements, placeholders)
	return s
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
func NewInlinePassthrough(kind PassthroughKind, elements []interface{}) (InlinePassthrough, error) {
	return InlinePassthrough{
		Kind:     kind,
		Elements: Merge(elements...),
	}, nil
}

var _ RawText = InlinePassthrough{}

// RawText returns the raw text representation of this element as it was (supposedly) written in the source document
func (p InlinePassthrough) RawText() (string, error) {
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
	Location   Location
	Attributes Attributes
}

// NewInlineLink initializes a new inline `InlineLink`
func NewInlineLink(url Location, attributes interface{}) (InlineLink, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		AttrPositional1: AttrInlineLinkText,
	})
	return InlineLink{
		Location:   url,
		Attributes: attrs,
	}, nil
}

var _ WithPlaceholdersInAttributes = InlineLink{}

// RestoreAttributes restores the attributes which had been substituted by placeholders
func (l InlineLink) RestoreAttributes(placeholders map[string]interface{}) interface{} {
	l.Attributes = restoreAttributes(l.Attributes, placeholders)
	return l
}

var _ WithPlaceholdersInLocation = InlineLink{}

// RestoreLocation restores the location elements which had been substituted by placeholders
func (l InlineLink) RestoreLocation(placeholders map[string]interface{}) interface{} {
	l.Location.Path = restoreElements(l.Location.Path, placeholders)
	return l
}

var _ WithElementsToSubstitute = InlineLink{}

// ElementsToSubstitute returns this inline link's location path so that substitutions can be applied onto its elements
func (l InlineLink) ElementsToSubstitute() []interface{} {
	return l.Location.Path
}

// ReplaceElements replaces the elements in this example block
func (l InlineLink) ReplaceElements(path []interface{}) interface{} {
	l.Location.Path = path
	return l
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
		case Attribute:
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
	Location   Location
	RawText    string
}

// NewFileInclusion initializes a new inline `InlineLink`
func NewFileInclusion(location Location, attributes interface{}, rawtext string) (FileInclusion, error) {
	attrs := toAttributesWithMapping(attributes, map[string]string{
		"tag": "tags", // convert `tag` to `tags`
	})
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
func NewLocation(scheme interface{}, path []interface{}) (Location, error) {
	path = Merge(path)
	// log.Debugf("new location: scheme='%v' path='%+v", scheme, path)
	s := ""
	if scheme, ok := scheme.([]byte); ok {
		s = string(scheme)
	}
	return Location{
		Scheme: s,
		Path:   path,
	}, nil
}

// WithPathPrefix adds the given prefix to the path if this latter is NOT an absolute
// path and if there is no defined scheme
func (l Location) WithPathPrefix(p interface{}) Location {
	if p, ok := p.(string); ok && p != "" {
		if !strings.HasSuffix(p, "/") {
			p = p + "/"
		}
		if l.Scheme == "" && !strings.HasPrefix(l.Stringify(), "/") {
			if u, err := url.Parse(l.Stringify()); err == nil {
				if !u.IsAbs() {
					l.Path = append([]interface{}{
						StringElement{
							Content: p,
						},
					}, l.Path...)
				}
			}
		}
	}
	return l
}

// Stringify returns a string representation of the location
func (l Location) Stringify() string {
	result := &strings.Builder{}
	result.WriteString(l.Scheme)
	for _, e := range l.Path {
		if s, ok := e.(StringElement); ok {
			result.WriteString(s.Content) // no need to use `fmt.Sprintf` for elements of type 'string'
		} else if s, ok := e.(SpecialCharacter); ok {
			result.WriteString(s.Name) // no need to use `fmt.Sprintf` for elements of type 'string'
		} else {
			result.WriteString(fmt.Sprintf("%s", e))
		}
	}
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
func NewIndexTerm(term []interface{}) (IndexTerm, error) {
	return IndexTerm{
		Term: term,
	}, nil
}

var _ WithPlaceholdersInElements = IndexTerm{}

// RestoreElements restores the elements which had been substituted by placeholders
func (t IndexTerm) RestoreElements(placeholders map[string]interface{}) interface{} {
	t.Term = restoreElements(t.Term, placeholders)
	return t
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
func NewSpecialCharacter(name string) (SpecialCharacter, error) {
	return SpecialCharacter{
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
func NewElementPlaceHolder(ref string) (ElementPlaceHolder, error) {
	return ElementPlaceHolder{
		Ref: ref,
	}, nil
}

func (p ElementPlaceHolder) String() string {
	return "\uFFFD" + p.Ref + "\uFFFD"
}

// replace the placeholders with their original element in the given elements
func restoreElements(elements []interface{}, placeholders map[string]interface{}) []interface{} {
	for i, e := range elements {
		switch e := e.(type) {
		case ElementPlaceHolder:
			elements[i] = placeholders[e.Ref]
		case []interface{}:
			elements[i] = restoreElements(e, placeholders)
		}
	}
	return elements
}

// replace the placeholders with their original element in the given attributes
func restoreAttributes(attrs Attributes, placeholders map[string]interface{}) Attributes {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("restoring placeholders in")
		spew.Fdump(log.StandardLogger().Out, attrs)
	}
	for key, value := range attrs {
		switch value := value.(type) {
		case ElementPlaceHolder:
			attrs[key] = placeholders[value.Ref]
		// case ElementRoles:
		// 	attrs[key] = ElementRoles(restoreElements(value, placeholders))
		case []interface{}:
			attrs[key] = restoreElements(value, placeholders)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("restored placeholders in")
		spew.Fdump(log.StandardLogger().Out, attrs)
	}
	return attrs
}

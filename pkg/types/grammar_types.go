package types

import (
	"bytes"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"reflect"

	"github.com/davecgh/go-spew/spew"
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
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	Attributes         DocumentAttributes
	Elements           []interface{}
	ElementReferences  ElementReferences
	Footnotes          Footnotes
	FootnoteReferences FootnoteReferences
}

// NewDocument initializes a new `Document` from the given lines
func NewDocument(frontmatter interface{}, elements []interface{}) (Document, error) {
	log.Debugf("initializing a new Document with %d block element(s)", len(elements))
	attributes := DocumentAttributes{}
	if frontmatter != nil {
		for attrName, attrValue := range frontmatter.(FrontMatter).Content {
			attributes[attrName] = attrValue
		}
	}
	//TODO: those collectors could be called at the beginning of rendering, and in concurrent routines
	// visit AST and collect element references
	xrefsCollector := NewElementReferencesCollector()
	for _, e := range elements {
		if v, ok := e.(Visitable); ok {
			err := v.AcceptVisitor(xrefsCollector)
			if err != nil {
				return Document{}, errors.Wrapf(err, "unable to create document")
			}
		}
	}

	// visit AST and collect footnotes
	footnotesCollector := NewFootnotesCollector()
	for _, e := range elements {
		log.Debugf("collecting footnotes in element of type %T", e)
		if v, ok := e.(Visitable); ok {
			err := v.AcceptVisitor(footnotesCollector)
			if err != nil {
				return Document{}, errors.Wrapf(err, "unable to create document")
			}
		}
	}
	document := Document{
		Attributes:         attributes,
		Elements:           NilSafe(elements),
		ElementReferences:  xrefsCollector.ElementReferences,
		Footnotes:          footnotesCollector.Footnotes,
		FootnoteReferences: footnotesCollector.FootnoteReferences,
	}
	// visit all elements in the `AST` to retrieve their reference (ie, their ElementID if they have any)
	return document, nil
}

// Title retrieves the document title in its metadata, or empty section title if the title was not specified
func (d Document) Title() (SectionTitle, bool) {
	if header, ok := d.Header(); ok {
		return header.Title, true
	}
	return SectionTitle{}, false
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
		switch author.(type) {
		case DocumentAuthor:
			result[i] = author.(DocumentAuthor)
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
	// log.Debugf("initialized a new document revision: `%s`", result.String())
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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (d DocumentAttributeDeclaration) AddAttributes(attributes ElementAttributes) {
	// nothing to do
	// TODO: raise a warning?
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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (r DocumentAttributeReset) AddAttributes(attributes ElementAttributes) {
	// nothing to do
	// TODO: raise a warning?
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
// PreparsedDocument (plus related types)
// ------------------------------------------

// PreparsedDocument a preprocessed document, aimed for file inclusions,
// beofre being fully parsed to obtain a Document
type PreparsedDocument struct {
	Elements []interface{}
}

// NewPreparsedDocument initializes a new PreparsedDocument with the given elements
func NewPreparsedDocument(elements []interface{}) (PreparsedDocument, error) {
	return PreparsedDocument{
		Elements: elements,
	}, nil
}

// RawSectionTitle a section. Just need to have the prefix which can be changed
// if there is a file inclusion with a level offset defined.
type RawSectionTitle struct {
	Prefix RawSectionTitlePrefix
	Title  RawSectionTitleContent
}

// NewRawSectionTitle return a new RawSectionTitle
func NewRawSectionTitle(prefix RawSectionTitlePrefix, title RawSectionTitleContent) (RawSectionTitle, error) {
	return RawSectionTitle{
		Prefix: prefix,
		Title:  title,
	}, nil
}

// Bytes returns the content of the preparsed section as an array of byte
func (s RawSectionTitle) Bytes(levelOffset string) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	b, err := s.Prefix.Bytes(levelOffset)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert section title")
	}
	result.Write(b)
	result.Write(s.Title)
	return result.Bytes(), nil
}

// RawSectionTitlePrefix a raw section prefix, with a distinction between the sequence of `=` characters
// and the following spaces
type RawSectionTitlePrefix struct {
	Level  []byte
	Spaces []byte
}

// NewRawSectionTitlePrefix returns a new RawSectionTitlePrefix with the given `=` sequences following by spaces.
func NewRawSectionTitlePrefix(level, spaces []byte) (RawSectionTitlePrefix, error) {
	return RawSectionTitlePrefix{
		Level:  level,
		Spaces: spaces,
	}, nil
}

// Bytes return the representation of this raw title prefix as an array of bytes,
// while applying an optional offset (if non-empty)
func (p RawSectionTitlePrefix) Bytes(levelOffset string) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if levelOffset != "" {
		log.Debugf("applying level offset '%s'", levelOffset)
		offset, err := strconv.Atoi(levelOffset)
		if err != nil {
			return nil, errors.Wrapf(err, "fail to apply level offset '%s' to document to include", levelOffset)
		}
		if offset > 0 {
			result.Write(p.Level)
			for i := 0; i < offset; i++ {
				result.WriteRune('=')
			}
		}
	} else {
		result.Write(p.Level)
	}
	result.Write(p.Spaces)
	return result.Bytes(), nil
}

// RawSectionTitleContent a raw title prefix
type RawSectionTitleContent []byte

// NewRawSectionTitleContent returns a new raw section title
func NewRawSectionTitleContent(content []byte) (RawSectionTitleContent, error) {
	return RawSectionTitleContent(content), nil
}

// Bytes return the content of the title as an array of bytes
func (t RawSectionTitleContent) Bytes() []byte {
	return []byte(t)
}

// RawText a line of raw text without the trailing `EOL`
type RawText []byte

// NewRawText initializes a RawText with the given content
func NewRawText(content []byte) (RawText, error) {
	return RawText(content), nil
}

// Bytes return the content of the text as an array of bytes
func (t RawText) Bytes() []byte {
	return []byte(t)
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
// Preamble
// ------------------------------------------

// Preamble the structure for document Preamble
type Preamble struct {
	Elements []interface{}
}

// NewEmptyPreamble return an empty Preamble
func NewEmptyPreamble() Preamble {
	return Preamble{
		Elements: make([]interface{}, 0),
	}
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
	Title      SectionTitle
	Attributes ElementAttributes
	Elements   []interface{}
}

// NewSection initializes a new `Section` from the given section title and elements
func NewSection(level int, title SectionTitle, elements []interface{}) (Section, error) {
	log.Debugf("initialized a new Section level %d with %d block(s)", level, len(elements))
	return Section{
		Level:      level,
		Title:      title,
		Attributes: ElementAttributes{},
		Elements:   NilSafe(elements),
	}, nil
}

// NewSection0WithMetadata initializes a new Section with level 0 which can have authors and a revision, among other attributes
func NewSection0WithMetadata(title SectionTitle, authors interface{}, revision interface{}, elements []interface{}) (Section, error) {
	log.Debugf("initializing a new Section0 with authors '%v' and revision '%v'", authors, revision)
	section := Section{
		Level:      0,
		Title:      title,
		Attributes: ElementAttributes{},
		Elements:   NilSafe(elements),
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
func (s Section) AddAttributes(attributes ElementAttributes) {
	log.Debugf("adding attributes to section: %v", attributes)
	s.Title.AddAttributes(attributes)
}

// GetElements returns the elements
func (s *Section) GetElements() []interface{} {
	return s.Elements
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (s Section) AcceptVisitor(v Visitor) error {
	err := v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	err = s.Title.AcceptVisitor(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section element")
	}
	for _, element := range s.Elements {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.AcceptVisitor(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting section element")
			}
		}

	}
	return nil
}

// AcceptSubstitutor implements Substituable#AcceptSubstitutor(Substitutor)
// in a section, the substitutor only cares about the elements for now.
func (s Section) AcceptSubstitutor(v Substitutor) (interface{}, error) {
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
// SectionTitle
// ------------------------------------------

// SectionTitle the structure for the section titles
type SectionTitle struct {
	Attributes ElementAttributes
	Elements   InlineElements
}

// NewSectionTitle initializes a new `SectionTitle`` from the given level and content, with the optional attributes.
// In the attributes, only the ElementID is retained
func NewSectionTitle(elements InlineElements, ids []interface{}) (SectionTitle, error) {
	attributes := ElementAttributes{}
	// multiple IDs can be defined (by mistake), and the last one is used
	for _, id := range ids {
		if id, ok := id.(ElementAttributes); ok {
			attributes.AddAll(id)
		}
	}
	attributes[AttrCustomID] = true
	// make a default id from the sectionTitle's inline content
	if _, found := attributes[AttrID]; !found {
		log.Debugf("did not find ID attribute for section with elements %v", elements)
		replacement, err := replaceNonAlphanumerics(elements, "_")
		if err != nil {
			return SectionTitle{}, errors.Wrapf(err, "unable to generate default ID while instanciating a new SectionTitle element")
		}
		attributes[AttrID] = replacement
		attributes[AttrCustomID] = false
	}
	sectionTitle := SectionTitle{
		Attributes: attributes,
		Elements:   elements,
	}
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("initialized a new SectionTitle with %d element(s)", len(elements))
	}
	return sectionTitle, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (st SectionTitle) AddAttributes(attributes ElementAttributes) {
	st.Attributes.AddAll(attributes)
	// look for custom ID
	for attr := range attributes {
		if attr == AttrID {
			// mark custom_id flag to `true`
			st.Attributes[AttrCustomID] = true
		}
	}
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (st SectionTitle) AcceptVisitor(v Visitor) error {
	err := v.Visit(st)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	for _, element := range st.Elements {
		visitable, ok := element.(Visitable)
		if ok {
			err = visitable.AcceptVisitor(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting section element")
			}
		}
	}
	return nil
}

// ------------------------------------------
// Lists
// ------------------------------------------

// List a list
type List interface {
	processContinuations(ancestors []ListItem) List
}

// ListItem a list item
type ListItem interface {
	GetElements() []interface{}
	AddElement(interface{})
	processContinuations([]ListItem) ListItem
}

// NewList initializes a new `List` from the given content
func NewList(items []interface{}) (List, error) {
	log.Debugf("initializing a new List with %d items(s)", len(items))
	builder := newListBuilder()
	for _, item := range items {
		listItem, ok := toPtr(item).(ListItem)
		if !ok {
			return nil, errors.Errorf("item of type '%T' is not a valid list item", item)
		}
		err := builder.process(listItem)
		if err != nil {
			return nil, errors.Wrap(err, "failed to initialize a list")
		}
	}
	// finally, process the first level of the builder's stack
	return builder.build()
}

type listBuilder struct {
	stacks        [][]ListItem
	previousDepth int
}

func newListBuilder() *listBuilder {
	return &listBuilder{
		stacks:        make([][]ListItem, 0),
		previousDepth: 0,
	}
}

// process:
// - checks if the given item's type is already known and at which level it is in the list
// - stores the item in its stack, at the detemined level
func (builder *listBuilder) process(item ListItem) error {
	log.Debugf("processing item of type %T", item)
	depth := builder.depth(item)
	// if moving up in the tree, then a new list needs to be build
	if depth < builder.previousDepth {
		log.Debugf("moving up in the stack, need to build %d list(s)", (builder.previousDepth - depth))
		for i := builder.previousDepth; i > depth; i-- {
			subitems := builder.stacks[i]
			sublist, err := newList(subitems)
			if err != nil {
				return errors.Wrap(err, "failed to initialize a new sublist")
			}
			// attach the new sublist to the last item of the parent level
			parentItem, err := builder.parentItem(i)
			if err != nil {
				return errors.Wrap(err, "failed to attach a new sublist to its parent item")
			}
			parentItem.AddElement(sublist)
			// clear the stack (ie, remove the last level)
			builder.stacks = builder.stacks[:len(builder.stacks)-1]
		}
	}
	builder.previousDepth = depth
	items := builder.stacks[depth]
	items = append(items, item)
	builder.stacks[depth] = items // 'items' was changed, needs to be put in the stack again
	return nil
}

// ends: builds a new list of each layer in the stack, starting by the end, and attach to the parent item
func (builder *listBuilder) build() (List, error) {
	for i := len(builder.stacks) - 1; i > 0; i-- {
		// if len(builder.stacks[i]) == 0 {
		// 	// ignore empty layer
		// 	continue
		// }
		sublist, err := newList(builder.stacks[i])
		if err != nil {
			return nil, errors.Wrapf(err, "failed to initialize a new list")
		}
		// look-up parent layer at the previous (ie, upper) level in the stack
		parentItems := builder.stacks[i-1]
		// look-up parent in the layer
		parentItem := parentItems[len(parentItems)-1]
		// build a new list from the remaining items at the current level of the stack
		// log.Debugf("building a new list from the remaining items of type '%T' and parent of type '%T' at the current level of the stack", buffer[stacks[i]][0], parentItem)
		// add this list to the parent
		parentItem.AddElement(sublist)
	}
	// finish with to "root" list
	result, err := newList(builder.stacks[0])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new list")
	}
	return result.processContinuations([]ListItem{}), nil
}

// depth finds at which depth of the stack the given item belongs, based on its type
func (builder *listBuilder) depth(item ListItem) int {
	itemType := reflect.TypeOf(item)
	log.Debugf("checking depth of item of type %T in a stack of size: %d", item, len(builder.stacks))
	for idx, items := range builder.stacks {
		// if layer of the stack is empty ior if first item has the same type
		if len(items) == 0 || reflect.TypeOf(items[0]) == itemType {
			log.Debugf("found matching layer in the stack for item of type %T: %d", item, idx)
			return idx
		}
	}
	// if there's no match, then add a new depth in the stack for this
	// type of item
	log.Debugf("adding a new layer in the stack for item of type %T", item)
	items := make([]ListItem, 0)
	builder.stacks = append(builder.stacks, items)
	return len(builder.stacks) - 1
}

func (builder *listBuilder) parentItem(childDepth int) (ListItem, error) {
	if childDepth == 0 {
		return nil, errors.New("unable to lookup parent for a root item (depth=0)")
	}
	parentItems := builder.stacks[childDepth-1]
	if len(parentItems) == 0 {
		return nil, errors.New("unable to lookup parent (none found at this level)")
	}
	return parentItems[len(parentItems)-1], nil
}

func newList(items []ListItem) (List, error) {
	// log.Debugf("initializing a new list with %d items", len(items))
	if len(items) == 0 {
		return nil, errors.Errorf("cannot build a list from an empty slice")
	}
	switch items[0].(type) {
	case *OrderedListItem:
		return newOrderedList(items)
	case *UnorderedListItem:
		return newUnorderedList(items)
	case *LabeledListItem:
		return newLabeledList(items)
	default:
		return nil, errors.Errorf("unsupported type of element as the root list: '%T'", items[0])
	}
}

// ContinuedListElement a wrapper for an element which should be attached to a list item (same level or an ancestor)
type ContinuedListElement struct {
	Offset  int // the relative ancestor. Should be a negative number
	Element interface{}
}

// NewContinuedListElement returns a wrapper for an element which should be attached to a list item (same level or an ancestor)
func NewContinuedListElement(offset int, element interface{}) (ContinuedListElement, error) {
	log.Debugf("imitializing a new continued list element for element of type %T", element)
	return ContinuedListElement{
		Offset:  offset,
		Element: element,
	}, nil
}

// ------------------------------------------
// Ordered Lists
// ------------------------------------------

// OrderedList the structure for the Ordered Lists
type OrderedList struct {
	Attributes ElementAttributes
	Items      []OrderedListItem
}

var _ List = OrderedList{}

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

var numberingStyles []NumberingStyle

func init() {
	numberingStyles = []NumberingStyle{Arabic, Decimal, LowerAlpha, UpperAlpha, LowerRoman, UpperRoman, LowerGreek, UpperGreek}
}

// newOrderedList initializes a new `OrderedList` from the given content
func newOrderedList(elements []ListItem) (OrderedList, error) {
	log.Debugf("initializing a new ordered list from %d element(s)...", len(elements))
	result := make([]OrderedListItem, 0)
	bufferedItemsPerLevel := make(map[int][]*OrderedListItem) // buffered items for the current level
	levelPerStyle := make(map[NumberingStyle]int)
	previousLevel := 0
	previousNumberingStyle := UnknownNumberingStyle
	for _, element := range elements {
		item, ok := element.(*OrderedListItem)
		if !ok {
			return OrderedList{}, errors.Errorf("element of type '%T' is not a valid orderedlist item", element)
		}
		if item.Level > previousLevel {
			// force the current item level to (last seen level + 1)
			item.Level = previousLevel + 1
			// log.Debugf("setting item level to %d (#1 - new level)", item.Level)
			levelPerStyle[item.NumberingStyle] = item.Level
		} else if item.NumberingStyle != previousNumberingStyle {
			// check if this numbering type was already found previously
			if level, found := levelPerStyle[item.NumberingStyle]; found {
				item.Level = level // 0-based level in the bufferedItemsPerLevel
				// log.Debugf("setting item level to %d / %v (#2 - existing style)", item.Level, item.NumberingStyle)
			} else {
				item.Level = previousLevel + 1
				// log.Debugf("setting item level to %d (#3 - new level for numbering style %v)", item.Level, item.NumberingStyle)
				levelPerStyle[item.NumberingStyle] = item.Level
			}
		} else if item.NumberingStyle == previousNumberingStyle {
			item.Level = previousLevel
			// log.Debugf("setting item level to %d (#4 - same as previous item)", item.Level)
		}
		// log.Debugf("list item %v -> level= %d", item.Elements[0], item.Level)
		// join item *values* in the parent item when the level decreased
		if item.Level < previousLevel {
			for i := previousLevel; i > item.Level; i-- {
				parentLayer := bufferedItemsPerLevel[i-2]
				parentItem := parentLayer[len(parentLayer)-1]
				log.Debugf("moving buffered items at level %d (%v) in parent (%v) ", i, bufferedItemsPerLevel[i-1][0].NumberingStyle, parentItem.NumberingStyle)
				childList, err := toOrderedList(bufferedItemsPerLevel[i-1])
				if err != nil {
					return OrderedList{}, err
				}
				parentItem.Elements = append(parentItem.Elements, childList)
				// clear the previously buffered items at level 'previousLevel'
				delete(bufferedItemsPerLevel, i-1)
			}
		}
		// new level of element: put it in the buffer
		if item.Level > len(bufferedItemsPerLevel) {
			// log.Debugf("initializing a new level of list items: %d", item.Level)
			bufferedItemsPerLevel[item.Level-1] = make([]*OrderedListItem, 0)
		}
		// append item to buffer of its level
		log.Debugf("adding list item %v in the current buffer at level %d", item.Elements[0], item.Level)
		bufferedItemsPerLevel[item.Level-1] = append(bufferedItemsPerLevel[item.Level-1], item)
		previousLevel = item.Level
		previousNumberingStyle = item.NumberingStyle
	}
	log.Debugf("processing the rest of the buffer...")
	// clear the remaining buffer and get the result in the reverse order of levels
	for level := len(bufferedItemsPerLevel) - 1; level >= 0; level-- {
		items := bufferedItemsPerLevel[level]
		// top-level items
		if level == 0 {
			for idx, item := range items {
				// set the position
				// log.Debugf("setting item #%d position to %d+%d", (idx + 1), items[0].Position, idx)
				item.Position = items[0].Position + idx
				result = append(result, *item)
			}
		} else {
			childList, err := toOrderedList(items)
			if err != nil {
				return OrderedList{}, err
			}
			parentLayer := bufferedItemsPerLevel[level-1]
			parentItem := parentLayer[len(parentLayer)-1]
			parentItem.Elements = append(parentItem.Elements, childList)
		}
	}

	return OrderedList{
		Attributes: ElementAttributes{},
		Items:      result,
	}, nil
}

func toOrderedList(items []*OrderedListItem) (OrderedList, error) {
	result := OrderedList{
		Attributes: ElementAttributes{}, // avoid nil `attributes`
	}
	// set the position and numbering style based on the optional attributes of the first item
	if len(items) == 0 {
		return result, nil
	}
	err := items[0].applyAttributes()
	if err != nil {
		return result, errors.Wrapf(err, "failed to convert items into an ordered list")
	}
	for idx, item := range items {
		// log.Debugf("setting item #%d position to %d+%d", (idx + 1), bufferedItemsPerLevel[previousLevel-1][0].Position, idx)
		item.Position = items[0].Position + idx
		item.NumberingStyle = items[0].NumberingStyle
		result.Items = append(result.Items, *item)
	}
	return result, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l OrderedList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
	// override the numbering style if applicable
	for attr := range attributes {
		switch attr {
		case string(Arabic):
			setNumberingStyle(l.Items, Arabic)
		case string(Decimal):
			setNumberingStyle(l.Items, Decimal)
		case string(LowerAlpha):
			setNumberingStyle(l.Items, LowerAlpha)
		case string(UpperAlpha):
			setNumberingStyle(l.Items, UpperAlpha)
		case string(LowerRoman):
			setNumberingStyle(l.Items, LowerRoman)
		case string(UpperRoman):
			setNumberingStyle(l.Items, UpperRoman)
		case string(LowerGreek):
			setNumberingStyle(l.Items, LowerGreek)
		case string(UpperGreek):
			setNumberingStyle(l.Items, UpperGreek)
		}
	}
}

func setNumberingStyle(items []OrderedListItem, n NumberingStyle) {
	log.Debugf("setting numbering style to %v on %d items", n, len(items))
	for i, item := range items {
		item.NumberingStyle = n
		items[i] = item // copy back in the list since this is not a list of pointers :/
	}
}

func (l OrderedList) processContinuations(ancestors []ListItem) List {
	log.Debugf("processing continuations on OrderedList with %d item(s)", len(l.Items))
	items := make([]OrderedListItem, len(l.Items))
	for i, item := range l.Items {
		pi := item.processContinuations(ancestors).(*OrderedListItem)
		items[i] = *pi
	}
	return OrderedList{
		Attributes: l.Attributes,
		Items:      items,
	}
}

// OrderedListItem the structure for the ordered list items
type OrderedListItem struct {
	Attributes     ElementAttributes
	Level          int
	Position       int
	NumberingStyle NumberingStyle
	Elements       []interface{}
}

// GetElements returns the elements of this OrderedListItem
func (i OrderedListItem) GetElements() []interface{} {
	return i.Elements
}

// AddElement add an element to this OrderedListItem
func (i *OrderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i *OrderedListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}) (OrderedListItem, error) {
	log.Debugf("initializing a new OrderedListItem")
	p := 1 // default position
	return OrderedListItem{
		Attributes:     ElementAttributes{},
		NumberingStyle: prefix.NumberingStyle,
		Level:          prefix.Level,
		Position:       p,
		Elements:       elements,
	}, nil
}

func (i *OrderedListItem) applyAttributes() error {
	log.Debugf("applying attributes on %[1]v: %[2]v (%[2]T)", i.Elements[0], i.Attributes)
	// numbering type override
	for _, style := range numberingStyles {
		if _, ok := i.Attributes[string(style)]; ok {
			i.NumberingStyle = style
			break
		}
	}
	// numbering offset
	if start, ok := i.Attributes["start"]; ok {
		if start, ok := start.(string); ok {
			s, err := strconv.ParseInt(start, 10, 64)
			if err != nil {
				return errors.Wrapf(err, "unable to parse 'start' value %v", start)
			}
			i.Position = int(s)
		}
	}
	log.Debugf("applied attributes on %v: position=%d, numbering=%v", i.Elements[0], i.Position, i.NumberingStyle)
	return nil
}

func (i *OrderedListItem) processContinuations(ancestors []ListItem) ListItem {
	return &OrderedListItem{
		Attributes:     i.Attributes,
		Level:          i.Level,
		Position:       i.Position,
		NumberingStyle: i.NumberingStyle,
		Elements:       processContinuations(ancestors, i),
	}
}

func processContinuations(ancestors []ListItem, i ListItem) []interface{} {
	log.Debugf("processing continuations on item of type '%T' with %d elements - hierarchy: %d ancestor(s)", i, len(i.GetElements()), len(ancestors))
	elements := []interface{}{}
	log.Debugf("starting to process %d elements of %T", len(i.GetElements()), i)
	s := len(i.GetElements())
	for _, element := range i.GetElements() {
		switch e := element.(type) {
		case List:
			// replace the list with a new list whose continued elements have been processed
			elements = append(elements, e.processContinuations(append(ancestors, i)))
		case ContinuedListElement:
			// move the element wrapped in the `ContinuedListElement` to the target ancestor
			idx := len(ancestors) + e.Offset
			if idx < 0 {
				idx = 0
			}
			// if no ancestor is available at all or no sibling at the expected level, then just add the wrapped element
			if len(ancestors) == 0 || len(ancestors) <= idx {
				elements = append(elements, e.Element)
				continue
			}
			log.WithField("ancestor_offset", e.Offset).WithField("ancestor_index", idx).WithField("ancestors_count", len(ancestors)).Debugf("moving element of type '%T' to last item of ancestor level", e.Element)
			ancestorItem := ancestors[idx]
			ancestorItem.AddElement(e.Element) // but we need to add the wrapped element to the copy of the ancestor
		default:
			// any other kind of element is kept
			elements = append(elements, element)
		}
	}
	// add all elements that were appended to this item's elements while continuation in sublists were processed
	return append(elements, i.GetElements()[s:]...)
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

var _ List = UnorderedList{}

// newUnorderedList initializes a new `UnorderedList` from the given content
func newUnorderedList(elements []ListItem) (UnorderedList, error) {
	log.Debugf("initializing a new unordered list from %d element(s)...", len(elements))
	result := make([]UnorderedListItem, 0)
	bufferedItemsPerLevel := make(map[int][]*UnorderedListItem) // buffered items for the current level
	levelPerStyle := make(map[BulletStyle]int)
	previousLevel := 0
	previousBulletStyle := UnknownBulletStyle
	for _, element := range elements {
		item, ok := element.(*UnorderedListItem)
		if !ok {
			return UnorderedList{}, errors.Errorf("element of type '%T' is not a valid unordered list item", element)
		}
		if item.Level > previousLevel {
			// force the current item level to (last seen level + 1)
			item.adjustBulletStyle(previousBulletStyle)
			item.Level = previousLevel + 1
			levelPerStyle[item.BulletStyle] = item.Level
		} else if item.BulletStyle != previousBulletStyle {
			if level, found := levelPerStyle[item.BulletStyle]; found {
				item.Level = level
			} else {
				item.Level = previousLevel + 1
				levelPerStyle[item.BulletStyle] = item.Level
			}
		} else if item.BulletStyle == previousBulletStyle {
			// adjust level on previous item of same style (in case the level
			// of the latter has been adjusted before)
			item.Level = previousLevel
		}
		log.Debugf("processing list item of level %d: %v", item.Level, item.Elements[0])
		// join item *values* in the parent item when the level decreased
		if item.Level < previousLevel {
			// merge previous levels in parents.
			// eg: when reaching `list item 2`, the level 3 items must be merged into the level 2 item, which must
			// be itself merged in the level 1 item:
			// * list item 1
			// ** nested list item
			// *** nested nested list item 1
			// *** nested nested list item 2
			// * list item 2
			for l := previousLevel; l > item.Level; l-- {
				log.Debugf("merging previously buffered items at level '%d' in parent", l)
				parentLayer := bufferedItemsPerLevel[l-2]
				parentItem := parentLayer[len(parentLayer)-1]
				childList := UnorderedList{
					Attributes: ElementAttributes{}, // avoid nil `attributes`
				}
				for _, i := range bufferedItemsPerLevel[l-1] {
					childList.Items = append(childList.Items, *i)
				}
				parentItem.Elements = append(parentItem.Elements, childList)
				// clear the previously buffered items at level 'previousLevel'
				delete(bufferedItemsPerLevel, l-1)
			}
		}
		// new level of element: put it in the buffer
		if item.Level > len(bufferedItemsPerLevel) {
			log.Debugf("initializing a new level of list items: %d", item.Level)
			bufferedItemsPerLevel[item.Level-1] = make([]*UnorderedListItem, 0)
		}
		// append item to buffer of its level
		log.Debugf("adding list item %v in the current buffer", item.Elements[0])
		bufferedItemsPerLevel[item.Level-1] = append(bufferedItemsPerLevel[item.Level-1], item)
		previousLevel = item.Level
		previousBulletStyle = item.BulletStyle
	}
	log.Debugf("processing the rest of the buffer: %v", bufferedItemsPerLevel)
	// clear the remaining buffer and get the result in the reverse order of levels
	for level := len(bufferedItemsPerLevel) - 1; level >= 0; level-- {
		items := bufferedItemsPerLevel[level]
		// top-level items
		if level == 0 {
			for _, item := range items {
				result = append(result, *item)
			}
		} else {
			childList := UnorderedList{
				Attributes: ElementAttributes{}, // avoid nil `attributes`
			}
			for _, item := range items {
				childList.Items = append(childList.Items, *item)
			}
			parentLayer := bufferedItemsPerLevel[level-1]
			parentItem := parentLayer[len(parentLayer)-1]
			parentItem.Elements = append(parentItem.Elements, childList)
		}
	}
	return UnorderedList{
		Attributes: ElementAttributes{},
		Items:      result,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l UnorderedList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
}

func (l UnorderedList) processContinuations(ancestors []ListItem) List {
	items := make([]UnorderedListItem, len(l.Items))
	for i, item := range l.Items {
		pi := item.processContinuations(ancestors).(*UnorderedListItem)
		items[i] = *pi
	}
	return UnorderedList{
		Attributes: l.Attributes,
		Items:      items,
	}
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
func NewUnorderedListItem(prefix UnorderedListItemPrefix, checkstyle interface{}, elements []interface{}) (UnorderedListItem, error) {
	log.Debugf("initializing a new UnorderedListItem with %d elements", len(elements))
	// log.Debugf("initializing a new UnorderedListItem with '%d' lines (%T) and input level '%d'", len(elements), elements, lvl.Len())
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
		Attributes:  ElementAttributes{},
		BulletStyle: prefix.BulletStyle,
		CheckStyle:  cs,
		Elements:    elements,
	}, nil
}

// GetElements returns the elements of this UnorderedListItem
func (i UnorderedListItem) GetElements() []interface{} {
	return i.Elements
}

// AddElement add an element to this UnorderedListItem
func (i *UnorderedListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i *UnorderedListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

func (i *UnorderedListItem) processContinuations(ancestors []ListItem) ListItem {
	return &UnorderedListItem{
		Attributes:  i.Attributes,
		Level:       i.Level,
		BulletStyle: i.BulletStyle,
		CheckStyle:  i.CheckStyle,
		Elements:    processContinuations(ancestors, i),
	}
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

// adjustBulletStyle
func (i *UnorderedListItem) adjustBulletStyle(p BulletStyle) {
	n := i.BulletStyle.nextLevelStyle(p)
	log.Debugf("adjusting bullet style for item with level '%v' to '%v' (previously processed/parent level: '%v')", i.BulletStyle, p, n)
	i.BulletStyle = n
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

// nextLevelStyle returns the BulletStyle for the next level:
// `-` -> `*`
// `*` -> `**`
// `**` -> `***`
// `***` -> `****`
// `****` -> `*****`
// `*****` -> `-`

func (b BulletStyle) nextLevelStyle(p BulletStyle) BulletStyle {
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

var _ List = LabeledList{}

// newLabeledList initializes a new `LabeledList` from the given content
func newLabeledList(elements []ListItem) (LabeledList, error) {
	log.Debugf("initializing a new labeled list from %d element(s)...", len(elements))
	result := make([]LabeledListItem, 0)
	bufferedItemsPerLevel := make(map[int][]*LabeledListItem) // buffered items for the current level
	previousLevel := 0
	for _, element := range elements {
		item, ok := element.(*LabeledListItem)
		if !ok {
			return LabeledList{}, errors.Errorf("element of type '%T' is not a valid labeled list item", element)
		}
		if item.Level > previousLevel {
			// force the current item level to (last seen level + 1)
			item.Level = previousLevel + 1
		}
		log.Debugf("list item %v -> level= %d", item.Elements, item.Level)
		// join item *values* in the parent item when the level decreased
		for l := previousLevel; l > item.Level; l-- {
			log.Debugf("merging previously buffered items at level '%d' in parent", l)
			parentLayer := bufferedItemsPerLevel[l-2]
			parentItem := parentLayer[len(parentLayer)-1]
			childList := LabeledList{
				Attributes: ElementAttributes{}, // avoid nil `attributes`
			}
			for _, i := range bufferedItemsPerLevel[l-1] {
				childList.Items = append(childList.Items, *i)
			}
			parentItem.Elements = append(parentItem.Elements, childList)
			// clear the previously buffered items at level 'previousLevel'
			delete(bufferedItemsPerLevel, l-1)
		}
		// new level of element: put it in the buffer
		if item.Level > len(bufferedItemsPerLevel) {
			log.Debugf("initializing a new level of list items: %d", item.Level)
			bufferedItemsPerLevel[item.Level-1] = make([]*LabeledListItem, 0)
		}
		// append item to buffer of its level
		log.Debugf("adding list item %v in the current buffer at level %d", item, item.Level)
		bufferedItemsPerLevel[item.Level-1] = append(bufferedItemsPerLevel[item.Level-1], item)
		previousLevel = item.Level
	}
	log.Debugf("processing the rest of the buffer: %v", bufferedItemsPerLevel)
	// clear the remaining buffer and get the result in the reverse order of levels
	for level := len(bufferedItemsPerLevel) - 1; level >= 0; level-- {
		items := bufferedItemsPerLevel[level]
		// top-level items
		if level == 0 {
			for _, item := range items {
				result = append(result, *item)
			}
		} else {
			childList := LabeledList{
				Attributes: ElementAttributes{}, // avoid nil `attributes`
			}
			for _, item := range items {
				childList.Items = append(childList.Items, *item)
			}
			parentLayer := bufferedItemsPerLevel[level-1]
			parentItem := parentLayer[len(parentLayer)-1]
			parentItem.Elements = append(parentItem.Elements, childList)
		}
	}
	return LabeledList{
		Attributes: ElementAttributes{},
		Items:      result,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l LabeledList) AddAttributes(attributes ElementAttributes) {
	l.Attributes.AddAll(attributes)
}

func (l LabeledList) processContinuations(ancestors []ListItem) List {
	log.Debugf("processing continuations on LabeledList with %d item(s)", len(l.Items))
	items := make([]LabeledListItem, len(l.Items))
	for i, item := range l.Items {
		pi := item.processContinuations(ancestors).(*LabeledListItem)
		items[i] = *pi
	}
	return LabeledList{
		Attributes: l.Attributes,
		Items:      items,
	}
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
func NewLabeledListItem(level int, term string, elements []interface{}) (LabeledListItem, error) {
	log.Debugf("initializing a new LabeledListItem")
	return LabeledListItem{
		Term:       strings.TrimSpace(term),
		Level:      level,
		Attributes: ElementAttributes{},
		Elements:   elements,
	}, nil
}

// GetElements returns the elements of this LabeledListItem
func (i LabeledListItem) GetElements() []interface{} {
	return i.Elements
}

// AddElement add an element to this LabeledListItem
func (i *LabeledListItem) AddElement(element interface{}) {
	i.Elements = append(i.Elements, element)
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i *LabeledListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

func (i *LabeledListItem) processContinuations(ancestors []ListItem) ListItem {
	return &LabeledListItem{
		Attributes: i.Attributes,
		Level:      i.Level,
		Term:       i.Term,
		Elements:   processContinuations(ancestors, i),
	}
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
func NewParagraph(lines []interface{}, attributes ...interface{}) (Paragraph, error) {
	log.Debugf("initializing a new paragraph with %d line(s) and %d attribute(s)", len(lines), len(attributes))
	elements := make([]InlineElements, 0)
	for _, line := range lines {
		if l, ok := line.(InlineElements); ok {
			log.Debugf("processing paragraph line of type %T", line)
			// if len(l) > 0 {
			elements = append(elements, l)
			// }
		} else {
			return Paragraph{}, errors.Errorf("unsupported paragraph line of type %[1]T: %[1]v", line)
		}

	}
	log.Debugf("generated a paragraph with %d line(s): %v", len(elements), elements)
	return Paragraph{
		Attributes: NewElementAttributes(attributes),
		Lines:      elements,
	}, nil
}

// NewAdmonitionParagraph returns a new Paragraph with an extra admonition attribute
func NewAdmonitionParagraph(lines []interface{}, admonitionKind AdmonitionKind, attributes ...interface{}) (Paragraph, error) {
	log.Debugf("new admonition paragraph")
	p, err := NewParagraph(lines, attributes)
	if err != nil {
		return p, err
	}
	p.Attributes[AttrAdmonitionKind] = admonitionKind
	return p, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (p Paragraph) AddAttributes(attributes ElementAttributes) {
	p.Attributes.AddAll(attributes)
}

// AcceptVisitor implements Visitable#AcceptVisitor(Visitor)
func (p Paragraph) AcceptVisitor(v Visitor) error {
	err := v.Visit(p)
	if err != nil {
		return errors.Wrapf(err, "error while visiting paragraph")
	}
	for _, line := range p.Lines {
		for _, element := range line {
			if visitable, ok := element.(Visitable); ok {
				err = visitable.AcceptVisitor(v)
				if err != nil {
					return errors.Wrapf(err, "error while visiting paragraph line")
				}
			}
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
	result := mergeElements(elements...)
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
	Path       string
	Attributes ElementAttributes
}

// NewImageBlock initializes a new `ImageBlock`
func NewImageBlock(path string, inlineAttributes ElementAttributes) (ImageBlock, error) {
	allAttributes := ElementAttributes{}
	for k, v := range inlineAttributes {
		allAttributes[k] = v
	}
	if alt, found := allAttributes[AttrImageAlt]; !found || alt == "" {
		_, filename := filepath.Split(path)
		ext := filepath.Ext(filename)
		log.Debugf("adding alt based on filename '%s' (ext=%s)", filename, ext)
		if ext != "" {
			allAttributes[AttrImageAlt] = strings.TrimSuffix(filename, ext)
		} else {
			allAttributes[AttrImageAlt] = filename
		}
	}
	return ImageBlock{
		Path:       path,
		Attributes: allAttributes,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i ImageBlock) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Path       string
	Attributes ElementAttributes
}

// NewInlineImage initializes a new `InlineImage` (similar to ImageBlock, but without attributes)
func NewInlineImage(path string, attributes ElementAttributes) (InlineImage, error) {
	if alt, found := attributes[AttrImageAlt]; !found || alt == "" {
		_, filename := filepath.Split(path)
		log.Debugf("adding alt based on filename '%s'", filename)
		ext := filepath.Ext(filename)
		if ext != "" {
			attributes[AttrImageAlt] = strings.TrimSuffix(filename, ext)
		} else {
			attributes[AttrImageAlt] = filename
		}
	}
	return InlineImage{
		Path:       path,
		Attributes: attributes,
	}, nil
}

// NewImageAttributes returns a map of image attributes, some of which have implicit keys (`alt`, `width` and `height`)
func NewImageAttributes(alt, width, height interface{}, otherattrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	var altStr, widthStr, heightStr string
	if alt, ok := alt.(string); ok {
		altStr = Apply(alt, strings.TrimSpace)
	}
	if width, ok := width.(string); ok {
		widthStr = Apply(width, strings.TrimSpace)
	}
	if height, ok := height.(string); ok {
		heightStr = Apply(height, strings.TrimSpace)
	}
	result[AttrImageAlt] = altStr
	result[AttrImageWidth] = widthStr
	result[AttrImageHeight] = heightStr
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
func NewDelimitedBlock(kind BlockKind, content []interface{}, substitution Substitution) (DelimitedBlock, error) {
	log.Debugf("initializing a new DelimitedBlock of kind '%v' with %d elements", kind, len(content))
	elements, err := substitution(content)
	if err != nil {
		return DelimitedBlock{}, errors.Wrapf(err, "failed to initialize a new delimited block")
	}
	return DelimitedBlock{
		Kind:       kind,
		Attributes: ElementAttributes{},
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
func NewTable(header interface{}, lines []interface{}) (Table, error) {
	t := Table{
		Attributes: ElementAttributes{},
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
	t.Lines = make([]TableLine, 0)
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
	log.Debugf("initialized a new table with %d line(s)", len(lines))
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
			log.Debugf("unsupported element of type %T", column)
		}
	}
	log.Debugf("initialized a new table line with %d columns", len(c))
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
func NewLiteralBlock(origin string, lines []interface{}, attributes ...interface{}) (LiteralBlock, error) {
	l, err := toString(lines)
	if err != nil {
		return LiteralBlock{}, errors.Wrapf(err, "unable to initialize a new LiteralBlock")
	}
	log.Debugf("initialized a new LiteralBlock with %d lines", len(lines))
	return LiteralBlock{
		Attributes: NewElementAttributes(
			attributes,
			ElementAttributes{
				AttrKind:             Literal,
				AttrLiteralBlockType: origin,
			},
		),
		Lines: l,
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

// ------------------------------------------------------------------------------------------------------------------------------
// Inline elements
// ------------------------------------------------------------------------------------------------------------------------------

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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (l SingleLineComment) AddAttributes(attributes ElementAttributes) {
	// nothing to do
	// TODO: raise a warning?
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
func NewQuotedText(kind QuotedTextKind, content []interface{}) (QuotedText, error) {
	elements := mergeElements(content...)
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("initialized a new QuotedText with %d elements: %v", len(elements), spew.Sdump(elements))
	}
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
func NewEscapedQuotedText(backslashes string, punctuation string, content []interface{}) ([]interface{}, error) {
	backslashesStr := Apply(backslashes,
		func(s string) string {
			// remove the number of back-slashes that match the length of the punctuation. Eg: `\*` or `\\**`, but keep extra back-slashes
			if len(s) > len(punctuation) {
				return s[len(punctuation):]
			}
			return ""
		})
	return []interface{}{backslashesStr, punctuation, content, punctuation}, nil
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
		Elements: mergeElements(elements...),
	}, nil

}

// ------------------------------------------
// Inline Links
// ------------------------------------------

// InlineLink the structure for the external links
type InlineLink struct {
	URL        string
	Attributes ElementAttributes
}

// NewInlineLink initializes a new inline `InlineLink`
func NewInlineLink(url string, attributes interface{}) (InlineLink, error) {
	attrs, ok := attributes.(ElementAttributes)
	// init attributes with empty 'text' attribute
	if !ok {
		attrs = ElementAttributes{
			AttrInlineLinkText: "",
		}
	}
	return InlineLink{
		URL:        url,
		Attributes: attrs,
	}, nil
}

// Text returns the `text` value for the InlineLink,
func (l InlineLink) Text() string {
	if text, ok := l.Attributes[AttrInlineLinkText].(string); ok {
		return text
	}
	return ""
}

// AttrInlineLinkText the link `text` attribute
const AttrInlineLinkText string = "text"

// NewInlineLinkAttributes returns a map of link attributes, some of which have implicit keys (`text`)
func NewInlineLinkAttributes(text interface{}, otherattrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	var textStr string
	if text, ok := text.(string); ok {
		textStr = Apply(text, strings.TrimSpace)
	}
	result[AttrInlineLinkText] = textStr
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
	Path       string
}

var _ ElementWithAttributes = FileInclusion{}

// NewFileInclusion initializes a new inline `InlineLink`
func NewFileInclusion(path string, attributes interface{}) (FileInclusion, error) {
	attrs, ok := attributes.(ElementAttributes)
	// init attributes with empty 'text' attribute
	if !ok {
		attrs = ElementAttributes{}
	}
	return FileInclusion{
		Attributes: attrs,
		Path:       path,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (f FileInclusion) AddAttributes(attributes ElementAttributes) {
	f.Attributes.AddAll(attributes)
}

// IsAsciidoc returns true if the file to include is an asciidoc file (based on the file path extension)
func (f FileInclusion) IsAsciidoc() bool {
	ext := filepath.Ext(f.Path)
	return ext == ".asciidoc" || ext == ".adoc" || ext == ".ad" || ext == ".asc" || ext == ".txt"
}

// LineRanges the ranges of lines of the child doc to include in the master doc
type LineRanges []LineRange

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

// NewLineRanges returns a slice of line ranges attribute for a file inclusion.
func NewLineRanges(ranges ...interface{}) LineRanges {
	result := LineRanges{}
	for _, r := range ranges {
		if r, ok := r.(LineRange); ok {
			result = append(result, r)
		}
	}
	// sort the range by `start` line
	sort.Sort(result)
	return result
}

// Match checks if the given line number matches one of the line ranges
func (r LineRanges) Match(line int) bool {
	for _, lr := range r {
		if lr.Start <= line && (lr.End >= line || lr.End == -1) {
			return true
		}
		if lr.Start > line {
			// no need to carry on with the ranges
			return false
		}
	}
	return false
}

// make sure that the LineRanges type implemnents the `sort.Interface
var _ sort.Interface = LineRanges{}

func (r LineRanges) Len() int           { return len(r) }
func (r LineRanges) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r LineRanges) Less(i, j int) bool { return r[i].Start < r[j].Start }

// LineRange the range of lines of the child doc to include in the master doc
// `Start` and `End` are the included limits of the child document
// - if there's a single line to include, then `End = Start`
// - if there is all remaining content after a given line (included), then `End = -1`
type LineRange struct {
	Start int
	End   int
}

// NewLineRangeAttribute returns a line range attribute for a file inclusion.
// The attribute value can be a single line range, a slice of line ranges
// or a string if the specified value could not be parsed.
func NewLineRangeAttribute(lines interface{}) (ElementAttributes, error) {
	return ElementAttributes{
		AttrLineRanges: lines,
	}, nil
}

// NewSingleLineRange returns a new single line range
func NewSingleLineRange(line int) (LineRange, error) {
	return LineRange{
		Start: line,
		End:   line,
	}, nil
}

// NewMultilineRange returns a new multi-line range
func NewMultilineRange(start, end int) (LineRange, error) {
	return LineRange{
		Start: start,
		End:   end,
	}, nil
}

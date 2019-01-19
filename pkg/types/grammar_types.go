package types

import (
	"fmt"
	"path/filepath"
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
	Accept(Visitor) error
}

// Visitor a visitor that can visit/traverse the interface{} and its children (if applicable)
type Visitor interface {
	Visit(Visitable) error
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
func NewDocument(frontmatter, header interface{}, blocks []interface{}) (Document, error) {
	log.Debugf("initializing a new Document with %d blocks(s)", len(blocks))
	// elements := convertBlocksTointerface{}s(blocks)
	// elements := filterEmptyElements(blocks, filterBlankLine(), filterEmptyPreamble())
	elements := insertPreamble(blocks)
	attributes := make(DocumentAttributes)
	if frontmatter != nil {
		for attrName, attrValue := range frontmatter.(FrontMatter).Content {
			attributes[attrName] = attrValue
		}
	}
	if header != nil {
		for attrName, attrValue := range header.(DocumentHeader).Content {
			attributes[attrName] = attrValue
			if attrName == "toc" {
				// insert a TableOfContentsMacro element if `toc` value is:
				// - "auto" (or empty)
				// - "preamble"
				switch attrValue {
				case "", "auto":
					// insert TableOfContentsMacro at first position
					elements = append([]interface{}{TableOfContentsMacro{}}, elements...)
				case "preamble":
					// lookup preamble in elements (should be first)
					preambleIndex := 0
					for i, e := range elements {
						if _, ok := e.(Preamble); ok {
							preambleIndex = i
							break
						}
					}
					// insert TableOfContentsMacro just after preamble
					remainingElements := make([]interface{}, len(elements)-(preambleIndex+1))
					copy(remainingElements, elements[preambleIndex+1:])
					elements = append(elements[0:preambleIndex+1], TableOfContentsMacro{})
					elements = append(elements, remainingElements...)
				case "macro":
				default:
					log.Warnf("invalid value for 'toc' attribute: '%s'", attrValue)

				}
			}
		}
	}
	//TODO: those collectors could be called at the beginning of rendering, and in concurrent routines
	// visit AST and collect element references
	xrefsCollector := NewElementReferencesCollector()
	for _, e := range elements {
		if v, ok := e.(Visitable); ok {
			err := v.Accept(xrefsCollector)
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
			err := v.Accept(footnotesCollector)
			if err != nil {
				return Document{}, errors.Wrapf(err, "unable to create document")
			}
		}
	}

	document := Document{
		Attributes:         attributes,
		Elements:           elements,
		ElementReferences:  xrefsCollector.ElementReferences,
		Footnotes:          footnotesCollector.Footnotes,
		FootnoteReferences: footnotesCollector.FootnoteReferences,
	}

	// visit all elements in the `AST` to retrieve their reference (ie, their ElementID if they have any)
	return document, nil
}

func insertPreamble(blocks []interface{}) []interface{} {
	// log.Debugf("generating preamble from %d blocks", len(blocks))
	preamble := NewEmptyPreamble()
	for _, block := range blocks {
		switch block.(type) {
		case Section:
			break
		default:
			preamble.Elements = append(preamble.Elements, block)
		}
	}
	// no element in the preamble, or no section in the document, so no preamble to generate
	if len(preamble.Elements) == 0 || len(preamble.Elements) == len(blocks) {
		log.Debugf("skipping preamble (%d vs %d)", len(preamble.Elements), len(blocks))
		return nilSafe(blocks)
	}
	// now, insert the preamble instead of the 'n' blocks that belong to the preamble
	// and copy the other items
	result := make([]interface{}, len(blocks)-len(preamble.Elements)+1)
	result[0] = preamble
	copy(result[1:], blocks[len(preamble.Elements):])
	log.Debugf("generated preamble with %d blocks", len(preamble.Elements))
	return result
}

// ------------------------------------------
// Document Header
// ------------------------------------------

// DocumentHeader the document header
type DocumentHeader struct {
	Content DocumentAttributes
}

// NewDocumentHeader initializes a new DocumentHeader
func NewDocumentHeader(header, authors, revision interface{}, otherAttributes []interface{}) (DocumentHeader, error) {
	content := DocumentAttributes{}
	if header != nil {
		content["doctitle"] = header.(SectionTitle)
	}
	log.Debugf("initializing a new DocumentHeader with content '%v', authors '%+v' and revision '%+v'", content, authors, revision)
	if authors != nil {
		for i, author := range authors.([]DocumentAuthor) {
			if i == 0 {
				content.AddNonEmpty("firstname", author.FirstName)
				content.AddNonEmpty("middlename", author.MiddleName)
				content.AddNonEmpty("lastname", author.LastName)
				content.AddNonEmpty("author", author.FullName)
				content.AddNonEmpty("authorinitials", author.Initials)
				content.AddNonEmpty("email", author.Email)
			} else {
				content.AddNonEmpty(fmt.Sprintf("firstname_%d", i+1), author.FirstName)
				content.AddNonEmpty(fmt.Sprintf("middlename_%d", i+1), author.MiddleName)
				content.AddNonEmpty(fmt.Sprintf("lastname_%d", i+1), author.LastName)
				content.AddNonEmpty(fmt.Sprintf("author_%d", i+1), author.FullName)
				content.AddNonEmpty(fmt.Sprintf("authorinitials_%d", i+1), author.Initials)
				content.AddNonEmpty(fmt.Sprintf("email_%d", i+1), author.Email)
			}
		}
	}
	if revision != nil {
		rev := revision.(DocumentRevision)
		content.AddNonEmpty("revnumber", rev.Revnumber)
		content.AddNonEmpty("revdate", rev.Revdate)
		content.AddNonEmpty("revremark", rev.Revremark)
	}
	for _, attr := range otherAttributes {
		if attr, ok := attr.(DocumentAttributeDeclaration); ok {
			content.AddAttribute(attr)
		}
	}
	return DocumentHeader{
		Content: content,
	}, nil
}

// ------------------------------------------
// Document Author
// ------------------------------------------

// DocumentAuthor a document author
type DocumentAuthor struct {
	FullName   string
	Initials   string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
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

//NewDocumentAuthor initializes a new DocumentAuthor
func NewDocumentAuthor(namePart1, namePart2, namePart3, emailAddress interface{}) (DocumentAuthor, error) {
	var part1, part2, part3, email string
	if namePart1, ok := namePart1.(string); ok {
		part1 = apply(namePart1,
			func(s string) string {
				return strings.TrimSpace(s)
			},
			func(s string) string {
				return strings.Replace(s, "_", " ", -1)
			},
		)
	}
	if namePart2, ok := namePart2.(string); ok {
		part2 = apply(namePart2,
			func(s string) string {
				return strings.TrimSpace(s)
			},
			func(s string) string {
				return strings.Replace(s, "_", " ", -1)
			},
		)
	}
	if namePart3, ok := namePart3.(string); ok {
		part3 = apply(namePart3,
			func(s string) string {
				return strings.TrimSpace(s)
			},
			func(s string) string {
				return strings.Replace(s, "_", " ", -1)
			},
		)
	}
	if emailAddress, ok := emailAddress.(string); ok {
		email = apply(emailAddress,
			func(s string) string {
				return strings.TrimPrefix(s, "<")
			}, func(s string) string {
				return strings.TrimSuffix(s, ">")
			}, func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	result := DocumentAuthor{}
	if part2 != "" && part3 != "" {
		result.FirstName = part1
		result.MiddleName = part2
		result.LastName = part3
		result.FullName = fmt.Sprintf("%s %s %s", part1, part2, part3)
		result.Initials = initials(result.FirstName, result.MiddleName, result.LastName)
	} else if part2 != "" {
		result.FirstName = part1
		result.LastName = part2
		result.FullName = fmt.Sprintf("%s %s", part1, part2)
		result.Initials = initials(result.FirstName, result.LastName)
	} else {
		result.FirstName = part1
		result.FullName = part1
		result.Initials = initials(result.FirstName)
	}
	result.Email = email
	// log.Debugf("initialized a new document author: `%v`", result.String())
	return result, nil
}

func initials(firstPart string, otherParts ...string) string {
	result := firstPart[0:1]
	for _, otherPart := range otherParts {
		result = result + otherPart[0:1]
	}
	return result
}

func (a *DocumentAuthor) String() string {
	email := ""
	if a.Email != "" {
		email = a.Email
	}
	return fmt.Sprintf("%s (%s)", a.FullName, email)
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
		number = apply(revnumber,
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
		date = apply(revdate,
			func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	if revremark, ok := revremark.(string); ok {
		// then we need to strip the heading ":" and spaces
		remark = apply(revremark,
			func(s string) string {
				return strings.TrimPrefix(s, ":")
			}, func(s string) string {
				return strings.TrimSpace(s)
			})
	}
	// log.Debugf("initializing a new DocumentRevision with revnumber='%v', revdate='%v' and revremark='%v'", *n, *d, *r)
	result := DocumentRevision{
		Revnumber: number,
		Revdate:   date,
		Revremark: remark,
	}
	// log.Debugf("initialized a new document revision: `%s`", result.String())
	return result, nil
}

func (r DocumentRevision) String() string {
	// return fmt.Sprintf("%v, %v: %v", number, date, remark)
	return fmt.Sprintf("%v, %v: %v", r.Revnumber, r.Revdate, r.Revremark)
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
	attrName = apply(name,
		func(s string) string {
			return strings.TrimSpace(s)
		})
	if value, ok := value.(string); ok {
		attrValue = apply(value,
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
// Element kinds
// ------------------------------------------

// BlockKind the kind of block
type BlockKind int

const (
	// AttrKind the key for the kind of block
	AttrKind string = "kind"
	// Fenced a fenced block
	Fenced BlockKind = iota // 1
	// Listing a listing block
	Listing
	// Example an example block
	Example
	// Comment a comment block
	Comment
	// Quote a quote block
	Quote
	// Verse a verse block
	Verse
	// Sidebar a sidebar block
	Sidebar
	// Literal a literal block
	Literal
	// Source a source block
	Source
)

// ------------------------------------------
// Table of Contents
// ------------------------------------------

// TableOfContentsMacro the structure for Table of Contents
type TableOfContentsMacro struct {
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (m TableOfContentsMacro) AddAttributes(attributes ElementAttributes) {
	// nothing to do
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

// Accept implements Visitable#Accept(Visitor)
func (p Preamble) Accept(v Visitor) error {
	err := v.Visit(p)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	for _, element := range p.Elements {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.Accept(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting section element")
			}
		}
	}
	return nil
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
	Level    int
	Title    SectionTitle
	Elements []interface{}
}

// NewSection initializes a new `Section` from the given section title and elements
func NewSection(level int, sectionTitle SectionTitle, blocks []interface{}) (Section, error) {
	log.Debugf("initialized a new Section level %d with %d block(s)", level, len(blocks))
	return Section{
		Level:    level,
		Title:    sectionTitle,
		Elements: nilSafe(blocks),
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (s Section) AddAttributes(attributes ElementAttributes) {
	s.Title.AddAttributes(attributes)
}

// Accept implements Visitable#Accept(Visitor)
func (s Section) Accept(v Visitor) error {
	err := v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	err = s.Title.Accept(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section element")
	}
	for _, element := range s.Elements {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.Accept(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting section element")
			}
		}

	}
	return nil
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
	// make a default id from the sectionTitle's inline content
	if _, found := attributes[AttrID]; !found {
		replacement, err := replaceNonAlphanumerics(elements, "_")
		if err != nil {
			return SectionTitle{}, errors.Wrapf(err, "unable to generate default ID while instanciating a new SectionTitle element")
		}
		attributes[AttrID] = replacement
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
}

// Accept implements Visitable#Accept(Visitor)
func (st SectionTitle) Accept(v Visitor) error {
	err := v.Visit(st)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	for _, element := range st.Elements {
		visitable, ok := element.(Visitable)
		if ok {
			err = visitable.Accept(v)
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
	// AddItems() []interface{}
}

// ListItem a list item
type ListItem interface {
	AddChild(interface{})
}

// NewList initializes a new `List` from the given content
func NewList(items []interface{}) (List, error) {
	log.Debugf("initializing a new List with %d items(s)", len(items))
	monitor := newListMonitor()
	for _, item := range items {
		listItem, ok := toPtr(item).(ListItem)
		if !ok {
			return nil, errors.Errorf("item of type '%T' is not a valid list item", item)
		}
		err := monitor.process(listItem)
		if err != nil {
			return nil, errors.Wrap(err, "failed to initialize a list")
		}
	}
	// finally, process the first level of the monitor's stack
	return monitor.end()
}

type listMonitor struct {
	stack                       [][]ListItem
	currentDepth, previousDepth int
}

func newListMonitor() *listMonitor {
	return &listMonitor{
		stack:         make([][]ListItem, 0),
		currentDepth:  0,
		previousDepth: 0,
	}
}

// process:
// - checks if the given item's type is already known and at which level it is in the list
// - stores the item in the inner stack, at the detemined level
// (ie, if the list is a mixed list)
// return the level (0-based offset) and `true` if the type of the item was already know, false otherwise
func (l *listMonitor) process(item ListItem) error {
	log.Debugf("processing item of type %T", item)
	depth := l.depth(item)
	// if moving up in the tree, then a new list needs to be build
	if depth < l.previousDepth {
		log.Debugf("moving up in the stack, need to build %d list(s)", (l.previousDepth - depth))
		for i := l.previousDepth; i > depth; i-- {
			subitems := l.stack[i]
			sublist, err := newList(subitems)
			if err != nil {
				return errors.Wrap(err, "failed to initialize a new sublist")
			}
			// attach the new sublist to the last item of the parent level
			parentItem, err := l.parentItem(i)
			if err != nil {
				return errors.Wrap(err, "failed to attach a new sublist to its parent item")
			}
			parentItem.AddChild(sublist)
			// clear the stack
			l.stack = l.stack[:len(l.stack)-1]
		}
	}
	l.previousDepth = depth
	// process the given item
	items := l.stack[depth]
	items = append(items, item)
	l.stack[depth] = items // 'items' was changed, needs to be put in the stack again
	return nil
}

// ends: builds a new list of each layer in the stack, starting by the end, and attach to the parent item
func (l *listMonitor) end() (List, error) {
	for i := len(l.stack) - 1; i > 0; i-- {
		// if len(l.stack[i]) == 0 {
		// 	// ignore empty layer
		// 	continue
		// }
		sublist, err := newList(l.stack[i])
		if err != nil {
			return nil, errors.Wrapf(err, "failed to initialize a new sublist")
		}
		// look-up parent layer at the previous (ie, upper) level in the stack
		parentItems := l.stack[i-1]
		// look-up parent in the layer
		parentItem := parentItems[len(parentItems)-1]
		// build a new list from the remaining items at the current level of the stack
		// log.Debugf("building a new list from the remaining items of type '%T' and parent of type '%T' at the current level of the stack", buffer[stack[i]][0], parentItem)
		// add this list to the parent
		parentItem.AddChild(sublist)
	}
	// finish with sublist
	return newList(l.stack[0])
}

// depth finds at which depth of the stack the given item belongs
func (l *listMonitor) depth(item ListItem) int {
	itemType := reflect.TypeOf(item)
	log.Debugf("checking depth of item of type %T in a stack of size: %d", item, len(l.stack))
	for idx, items := range l.stack {
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
	l.stack = append(l.stack, items)
	return len(l.stack) - 1
}

func (l *listMonitor) parentItem(childDepth int) (ListItem, error) {
	if childDepth == 0 {
		return nil, errors.New("unable to lookup parent for a root item (depth=0)")
	}
	parentItems := l.stack[childDepth-1]
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

// ------------------------------------------
// Ordered Lists
// ------------------------------------------

// OrderedList the structure for the Ordered Lists
type OrderedList struct {
	Attributes ElementAttributes
	Items      []OrderedListItem
}

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
	log.Debugf(" initializing a new ordered list from %d element(s)...", len(elements))
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
				item.Level = level // 0-based offset in the bufferedItemsPerLevel
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
			parentLayer := bufferedItemsPerLevel[previousLevel-2]
			parentItem := parentLayer[len(parentLayer)-1]
			log.Debugf(" moving buffered items at level %d (%v) in parent (%v) ", previousLevel, bufferedItemsPerLevel[previousLevel-1][0].NumberingStyle, parentItem.NumberingStyle)
			childList, err := toOrderedList(bufferedItemsPerLevel[previousLevel-1])
			if err != nil {
				return OrderedList{}, err
			}
			parentItem.Elements = append(parentItem.Elements, childList)
			// clear the previously buffered items at level 'previousLevel'
			delete(bufferedItemsPerLevel, previousLevel-1)
		}
		// new level of element: put it in the buffer
		if item.Level > len(bufferedItemsPerLevel) {
			// log.Debugf("initializing a new level of list items: %d", item.Level)
			bufferedItemsPerLevel[item.Level-1] = make([]*OrderedListItem, 0)
		}
		// append item to buffer of its level
		log.Debugf(" adding list item %v in the current buffer at level %d", item.Elements[0], item.Level)
		bufferedItemsPerLevel[item.Level-1] = append(bufferedItemsPerLevel[item.Level-1], item)
		previousLevel = item.Level
		previousNumberingStyle = item.NumberingStyle
	}
	log.Debugf(" processing the rest of the buffer...")
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
func (o OrderedList) AddAttributes(attributes ElementAttributes) {
	o.Attributes.AddAll(attributes)
	// override the numbering style if applicable
	for attr := range attributes {
		switch attr {
		case string(Arabic):
			setNumberingStyle(o.Items, Arabic)
		case string(Decimal):
			setNumberingStyle(o.Items, Decimal)
		case string(LowerAlpha):
			setNumberingStyle(o.Items, LowerAlpha)
		case string(UpperAlpha):
			setNumberingStyle(o.Items, UpperAlpha)
		case string(LowerRoman):
			setNumberingStyle(o.Items, LowerRoman)
		case string(UpperRoman):
			setNumberingStyle(o.Items, UpperRoman)
		case string(LowerGreek):
			setNumberingStyle(o.Items, LowerGreek)
		case string(UpperGreek):
			setNumberingStyle(o.Items, UpperGreek)
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

// OrderedListItem the structure for the ordered list items
type OrderedListItem struct {
	Level          int
	Position       int
	NumberingStyle NumberingStyle
	Attributes     ElementAttributes
	Elements       []interface{}
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}) (OrderedListItem, error) {
	log.Debugf("initializing a new OrderedListItem")
	p := 1 // default position
	return OrderedListItem{
		NumberingStyle: prefix.NumberingStyle,
		Level:          prefix.Level,
		Position:       p,
		Elements:       elements,
		Attributes:     ElementAttributes{},
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i OrderedListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)

}

// AddChild appends the given item to the content of this OrderedListItem
func (i *OrderedListItem) AddChild(item interface{}) {
	log.Debugf("adding item of type %T to list item of type %T (%v)", item, i, i.Elements)
	i.Elements = append(i.Elements, item)
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
		log.Debugf("Processing list item of level %d: %v", item.Level, item.Elements[0])
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
	log.Debugf("initializing a new UnorderedListItem...")
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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i UnorderedListItem) AddAttributes(attributes ElementAttributes) {
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

// AddChild appends the given item to the content of this UnorderedListItem
func (i *UnorderedListItem) AddChild(item interface{}) {
	log.Debugf("adding item of type %T to list item of type %T (%v)", item, i, i.Elements)
	i.Elements = append(i.Elements, item)
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

// ListItemContinuation a list item continuation
type ListItemContinuation struct {
}

// NewListItemContinuation returns a new ListItemContinuation
func NewListItemContinuation() (ListItemContinuation, error) {
	return ListItemContinuation{}, nil
}

// ------------------------------------------
// Labeled List
// ------------------------------------------

// LabeledList the structure for the Labeled Lists
type LabeledList struct {
	Attributes ElementAttributes
	Items      []LabeledListItem
}

// newLabeledList initializes a new `LabeledList` from the given content
func newLabeledList(elements []ListItem) (LabeledList, error) {
	log.Debugf(" initializing a new labeled list from %d element(s)...", len(elements))
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
		log.Debugf(" adding list item %v in the current buffer at level %d", item, item.Level)
		bufferedItemsPerLevel[item.Level-1] = append(bufferedItemsPerLevel[item.Level-1], item)
		previousLevel = item.Level
	}
	log.Debugf(" processing the rest of the buffer: %v", bufferedItemsPerLevel)
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

// LabeledListItem an item in a labeled
type LabeledListItem struct {
	Term       string
	Level      int
	Attributes ElementAttributes
	Elements   []interface{}
}

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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i LabeledListItem) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// AddChild appends the given item to the content of this LabeledListItem
func (i *LabeledListItem) AddChild(item interface{}) {
	log.Debugf("adding item of type %T to list item of type %T (%v)", item, i, i.Elements)
	i.Elements = append(i.Elements, item)
}

// making sure that the `ListItem` interface is implemented by `LabeledListItem`
var _ ListItem = &LabeledListItem{}

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

// Accept implements Visitable#Accept(Visitor)
func (p Paragraph) Accept(v Visitor) error {
	err := v.Visit(p)
	if err != nil {
		return errors.Wrapf(err, "error while visiting paragraph")
	}
	for _, line := range p.Lines {
		for _, element := range line {
			if visitable, ok := element.(Visitable); ok {
				err = visitable.Accept(v)
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

// Accept implements Visitable#Accept(Visitor)
func (e InlineElements) Accept(v Visitor) error {
	err := v.Visit(e)
	if err != nil {
		return errors.Wrapf(err, "error while visiting inline content")
	}
	for _, element := range e {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.Accept(v)
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
		l = apply(label, strings.TrimSpace)
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
	// allAttributes := mergeAttributes(attributes)
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

// AddAttributes adds all given attributes to the current set of attribute of the element
func (i InlineImage) AddAttributes(attributes ElementAttributes) {
	i.Attributes.AddAll(attributes)
}

// ImageMacro the structure for the block image macros
type ImageMacro struct {
	Path       string
	Attributes ElementAttributes
}

// NewImageMacro initializes a new `ImageMacro`
func NewImageMacro(path string, attributes ElementAttributes) (ImageMacro, error) {
	// use the image filename without the extension as the default `alt` attribute
	log.Debugf("processing alt: '%s'", attributes[AttrImageAlt])
	if attributes[AttrImageAlt] == "" {
		_, filename := filepath.Split(path)
		log.Debugf("adding alt based on filename '%s'", filename)
		ext := filepath.Ext(filename)
		if ext != "" {
			attributes[AttrImageAlt] = strings.TrimSuffix(filename, ext)
		} else {
			attributes[AttrImageAlt] = filename
		}
	}
	return ImageMacro{
		Path:       path,
		Attributes: attributes,
	}, nil
}

// NewImageAttributes returns a map of image attributes, some of which have implicit keys (`alt`, `width` and `height`)
func NewImageAttributes(alt, width, height interface{}, otherAttrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	var altStr, widthStr, heightStr string
	if alt, ok := alt.(string); ok {
		altStr = apply(alt, strings.TrimSpace)
	}
	if width, ok := width.(string); ok {
		widthStr = apply(width, strings.TrimSpace)
	}
	if height, ok := height.(string); ok {
		heightStr = apply(height, strings.TrimSpace)
	}
	result[AttrImageAlt] = altStr
	result[AttrImageWidth] = widthStr
	result[AttrImageHeight] = heightStr
	for _, otherAttr := range otherAttrs {
		if otherAttr, ok := otherAttr.(ElementAttributes); ok {
			for k, v := range otherAttr {
				result[k] = v
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

// Accept implements Visitable#Accept(Visitor)
func (f Footnote) Accept(v Visitor) error {
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
	Attributes ElementAttributes
	Elements   []interface{}
}

// Substitution the substitution group to apply when initializing a delimited block
type Substitution func([]interface{}) ([]interface{}, error)

// None returns the content as-is, but nil-safe
func None(content []interface{}) ([]interface{}, error) {
	return nilSafe(content), nil
}

// Verbatim the verbatim substitution: the given content is converted into an array of strings.
func Verbatim(content []interface{}) ([]interface{}, error) {
	result := make([]interface{}, len(content))
	for i, c := range content {
		if c, ok := c.(string); ok {
			c = apply(c, func(s string) string {
				return strings.TrimRight(c, "\n\r")
			})
			result[i] = NewStringElement(c)
		}
	}
	return result, nil
}

// NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given content
func NewDelimitedBlock(kind BlockKind, content []interface{}, substitution Substitution, attributes ...interface{}) (DelimitedBlock, error) {
	log.Debugf("initializing a new DelimitedBlock of kind '%v' with %d elements", kind, len(content))
	attrbs := NewElementAttributes(attributes)
	if _, found := attrbs[AttrKind]; !found { // add if missing
		attrbs[AttrKind] = kind
	}
	elements, err := substitution(content)
	if err != nil {
		return DelimitedBlock{}, errors.Wrapf(err, "failed to initialize a new delimited block")
	}
	return DelimitedBlock{
		Attributes: attrbs,
		Elements:   elements,
	}, nil
}

// AddAttributes adds all given attributes to the current set of attribute of the element
func (b DelimitedBlock) AddAttributes(attributes ElementAttributes) {
	b.Attributes.AddAll(attributes)
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
func NewStringElement(content string) StringElement {
	return StringElement{Content: content}
}

// Accept implements Visitable#Accept(Visitor)
func (s StringElement) Accept(v Visitor) error {
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
	Attributes ElementAttributes
	Elements   InlineElements
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
		Attributes: map[string]interface{}{AttrKind: kind},
		Elements:   elements,
	}, nil
}

// Accept implements Visitable#Accept(Visitor)
func (t QuotedText) Accept(v Visitor) error {
	err := v.Visit(t)
	if err != nil {
		return errors.Wrapf(err, "error while visiting quoted text")
	}
	for _, element := range t.Elements {
		if visitable, ok := element.(Visitable); ok {
			err := visitable.Accept(v)
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
	backslashesStr := apply(backslashes,
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

// NewInlineLinkAttributes returns a map of image attributes, some of which have implicit keys (`text`)
func NewInlineLinkAttributes(text interface{}, otherAttrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	var textStr string
	if text, ok := text.(string); ok {
		textStr = apply(text, strings.TrimSpace)
	}
	result[AttrInlineLinkText] = textStr
	for _, otherAttr := range otherAttrs {
		if otherAttr, ok := otherAttr.(ElementAttributes); ok {
			for k, v := range otherAttr {
				result[k] = v
			}
		}
	}
	return result, nil
}

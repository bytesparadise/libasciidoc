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
	BeforeVisit(Visitable) error
	Visit(Visitable) error
	AfterVisit(Visitable) error
}

// ------------------------------------------
// Document
// ------------------------------------------

// Document the top-level structure for a document
type Document struct {
	Attributes        DocumentAttributes
	Elements          []interface{}
	ElementReferences ElementReferences
}

// NewDocument initializes a new `Document` from the given lines
func NewDocument(frontmatter, header interface{}, blocks []interface{}) (Document, error) {
	log.Debugf("Initializing a new Document with %d blocks(s)", len(blocks))
	for i, block := range blocks {
		log.Debugf("Line #%d: %T", i, block)
	}
	// elements := convertBlocksTointerface{}s(blocks)
	elements := filterEmptyElements(blocks, filterBlankLine(), filterEmptyPreamble())
	attributes := make(map[string]interface{})

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

	c := NewElementReferencesCollector()
	for _, e := range elements {
		if v, ok := e.(Visitable); ok {
			v.Accept(c)
		}
	}
	document := Document{
		Attributes:        attributes,
		Elements:          elements,
		ElementReferences: c.ElementReferences,
	}

	// visit all elements in the `AST` to retrieve their reference (ie, their ElementID if they have any)
	return document, nil
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
	log.Debugf("Initializing a new DocumentHeader with content '%v', authors '%+v' and revision '%+v'", content, authors, revision)
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
	log.Debugf("Initializing a new array of document authors from `%+v`", authors)
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
	var err error
	if namePart1 != nil {
		part1, err = stringify(namePart1.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			},
			func(s string) (string, error) {
				return strings.Replace(s, "_", " ", -1), nil
			},
		)
		if err != nil {
			return DocumentAuthor{}, errors.Wrapf(err, "error while initializing a DocumentAuthor")
		}
	}
	if namePart2 != nil {
		part2, err = stringify(namePart2.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			},
			func(s string) (string, error) {
				return strings.Replace(s, "_", " ", -1), nil
			},
		)
		if err != nil {
			return DocumentAuthor{}, errors.Wrapf(err, "error while initializing a DocumentAuthor")
		}
	}
	if namePart3 != nil {
		part3, err = stringify(namePart3.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			},
			func(s string) (string, error) {
				return strings.Replace(s, "_", " ", -1), nil
			},
		)
		if err != nil {
			return DocumentAuthor{}, errors.Wrapf(err, "error while initializing a DocumentAuthor")
		}
	}
	if emailAddress != nil {
		email, err = stringify(emailAddress.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimPrefix(s, "<"), nil
			}, func(s string) (string, error) {
				return strings.TrimSuffix(s, ">"), nil
			}, func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			})
		if err != nil {
			return DocumentAuthor{}, errors.Wrapf(err, "error while initializing a DocumentAuthor")
		}
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
	log.Debugf("Initialized a new document author: `%v`", result.String())
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
	// log.Debugf("Initializing document revision with revnumber=%v, revdate=%v, revremark=%v", revnumber, revdate, revremark)
	// stringify, then remove the "v" prefix and trim spaces
	var number, date, remark string
	var err error
	if revnumber != nil {
		number, err = stringify(revnumber.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimPrefix(s, "v"), nil
			}, func(s string) (string, error) {
				return strings.TrimPrefix(s, "V"), nil
			}, func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			})
		if err != nil {
			return DocumentRevision{}, errors.Wrapf(err, "error while initializing a DocumentRevision")
		}
	}
	if revdate != nil {
		// stringify, then remove the "," prefix and trim spaces
		date, err = stringify(revdate.([]interface{}), func(s string) (string, error) {
			return strings.TrimSpace(s), nil
		})
		if err != nil {
			return DocumentRevision{}, errors.Wrapf(err, "error while initializing a DocumentRevision")
		}
		// do not keep empty values
		// if date == "" {
		// 	date = nil
		// }
	}
	if revremark != nil {
		// then we need to strip the heading "," and spaces
		remark, err = stringify(revremark.([]interface{}),
			func(s string) (string, error) {
				return strings.TrimPrefix(s, ":"), nil
			}, func(s string) (string, error) {
				return strings.TrimSpace(s), nil
			})
		if err != nil {
			return DocumentRevision{}, errors.Wrapf(err, "error while initializing a DocumentRevision")
		}
		// do not keep empty values
		// if *remark == "" {
		// 	remark = nil
		// }
	}
	// log.Debugf("Initializing a new DocumentRevision with revnumber='%v', revdate='%v' and revremark='%v'", *n, *d, *r)
	result := DocumentRevision{
		Revnumber: number,
		Revdate:   date,
		Revremark: remark,
	}
	log.Debugf("Initialized a new document revision: `%s`", result.String())
	return result, nil
}

func (r DocumentRevision) String() string {
	// number := ""
	// if r.Revnumber != nil {
	// 	number = *r.Revnumber
	// }
	// date := ""
	// if r.Revdate != nil {
	// 	date = *r.Revdate
	// }
	// remark := ""
	// if r.Revremark != nil {
	// 	remark = *r.Revremark
	// }
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

// NewDocumentAttributeDeclaration initializes a new DocumentAttributeDeclaration
func NewDocumentAttributeDeclaration(name []interface{}, value []interface{}) (DocumentAttributeDeclaration, error) {
	attrName, err := stringify(name,
		func(s string) (string, error) {
			return strings.TrimSpace(s), nil
		})
	if err != nil {
		return DocumentAttributeDeclaration{}, errors.Wrapf(err, "error while initializing a DocumentAttributeDeclaration")
	}
	attrValue, err := stringify(value,
		func(s string) (string, error) {
			return strings.TrimSpace(s), nil
		})
	if err != nil {
		return DocumentAttributeDeclaration{}, errors.Wrapf(err, "error while initializing a DocumentAttributeDeclaration")
	}
	log.Debugf("Initialized a new DocumentAttributeDeclaration: '%s' -> '%s'", attrName, attrValue)
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
func NewDocumentAttributeReset(name []interface{}) (DocumentAttributeReset, error) {
	attrName, err := stringify(name)
	if err != nil {
		return DocumentAttributeReset{}, errors.Wrapf(err, "error while initializing a DocumentAttributeReset")
	}
	log.Debugf("Initialized a new DocumentAttributeReset: '%s'", attrName)
	return DocumentAttributeReset{Name: attrName}, nil
}

// DocumentAttributeSubstitution the type for DocumentAttributeSubstitution
type DocumentAttributeSubstitution struct {
	Name string
}

// NewDocumentAttributeSubstitution initializes a new Document Attribute Substitutions
func NewDocumentAttributeSubstitution(name []interface{}) (DocumentAttributeSubstitution, error) {
	attrName, err := stringify(name)
	if err != nil {
		return DocumentAttributeSubstitution{}, errors.Wrapf(err, "error while initializing a DocumentAttributeSubstitution")
	}
	log.Debugf("Initialized a new DocumentAttributeSubstitution: '%s'", attrName)
	return DocumentAttributeSubstitution{Name: attrName}, nil
}

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

// NewPreamble initializes a new Preamble from the given elements
func NewPreamble(elements []interface{}) (Preamble, error) {
	log.Debugf("Initialiazing new Preamble with %d elements", len(elements))
	return Preamble{Elements: filterEmptyElements(elements, filterBlankLine())}, nil
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
	log.Debugf("Initialized a new FrontMatter with attributes: %+v", attributes)
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
	log.Debugf("Initializing a new Section with %d block(s)", len(blocks))
	elements := filterEmptyElements(blocks, filterBlankLine())
	log.Debugf("Initialized a new Section of level %d with %d block(s)", level, len(blocks))
	return Section{
		Level:    level,
		Title:    sectionTitle,
		Elements: elements,
	}, nil
}

// Accept implements Visitable#Accept(Visitor)
func (s Section) Accept(v Visitor) error {
	err := v.BeforeVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting section")
	}
	err = v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	for _, element := range s.Elements {
		if visitable, ok := element.(Visitable); ok {
			err = visitable.Accept(v)
			if err != nil {
				return errors.Wrapf(err, "error while visiting section element")
			}
		}

	}
	err = v.AfterVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting section")
	}
	return nil
}

// ------------------------------------------
// SectionTitle
// ------------------------------------------

// SectionTitle the structure for the section titles
type SectionTitle struct {
	Attributes map[string]interface{}
	Content    InlineElements
}

// NewSectionTitle initializes a new `SectionTitle`` from the given level and content, with the optional attributes.
// In the attributes, only the ElementID is retained
func NewSectionTitle(inlineContent InlineElements, attributes []interface{}) (SectionTitle, error) {
	// counting the lenght of the 'level' value (ie, the number of `=` chars)
	attrbs := NewElementAttributes(attributes)
	// make a default id from the sectionTitle's inline content
	if _, found := attrbs[AttrID]; !found {
		replacement, err := ReplaceNonAlphanumerics(inlineContent, "_")
		if err != nil {
			return SectionTitle{}, errors.Wrapf(err, "unable to generate default ID while instanciating a new SectionTitle element")
		}
		id, err := NewElementID(replacement)
		if err != nil {
			return SectionTitle{}, errors.Wrapf(err, "unable to generate default ID while instanciating a new SectionTitle element")
		}
		attrbs[AttrID] = id.Value
	}
	sectionTitle := SectionTitle{
		Attributes: attrbs,
		Content:    inlineContent,
	}
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("Initialized a new SectionTitle:")
		spew.Dump(sectionTitle)
	}
	return sectionTitle, nil
}

// ------------------------------------------
// Lists
// ------------------------------------------

// List a List
type List interface {
	// AddItems() []interface{}
}

// ListItem a list item
type ListItem interface {
	AddChild(interface{})
}

// NewList initializes a new `List` from the given content
func NewList(elements []interface{}, attributes []interface{}) (List, error) {
	log.Debugf("Initializing a new List with %d elements", len(elements))
	buffer := make(map[reflect.Type][]ListItem)
	rootType := reflect.TypeOf(toPtr(elements[0])) // elements types will be pointers
	previousType := rootType
	stack := make([]reflect.Type, 0)
	stack = append(stack, rootType)
	for _, element := range elements {
		log.Debugf("processing list item of type %T", element)
		// val := reflect.ValueOf(element).Elem().Addr().Interface()
		item, ok := toPtr(element).(ListItem)
		if !ok {
			return nil, errors.Errorf("element of type '%T' is not a valid list item", element)
		}
		// collect all elements of the same kind and make a sub list from them
		// each time a change of type is detected, except for the root type
		currentType := reflect.TypeOf(item)
		if currentType != previousType && previousType != rootType {
			log.Debugf(" detected a switch of type when processing item of type %T: currentType=%v != previousType=%v", item, currentType, previousType)
			// change of type: make a list from the buffer[t], reset and keep iterating
			sublist, err := newList(buffer[previousType], nil)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to initialize a sublist")
			}
			// look-up the previous item of the same type as the current type
			parentItems := buffer[currentType]
			parentItem := parentItems[len(parentItems)-1]
			parentItem.AddChild(sublist)
			buffer[previousType] = make([]ListItem, 0)
			// add element type to stack if not already found
			found := false
			for _, t := range stack {
				log.Debugf("comparing stack type %v to %v: %t", t, previousType, (t == previousType))
				if t == previousType {
					found = true
					break
				}
			}
			if !found {
				log.Debugf("adding element of type %v to stack", previousType)
				stack = append(stack, previousType)
			}
		}
		previousType = currentType
		// add item to buffer
		buffer[currentType] = append(buffer[currentType], item)
	}
	// end of processing: take into account the remainings in the buffer, by stack
	log.Debugf("end of list init: stack=%v, buffer= %v", stack, buffer)
	// process all sub lists
	for i := len(stack) - 1; i > 0; i-- {
		// skip if no item at this layer/level
		if len(buffer[stack[i]]) == 0 {
			continue
		}
		// look-up parent layer at the previous (ie, upper) level in the stack
		parentItems := buffer[stack[i-1]]
		// look-up parent in the layer
		parentItem := parentItems[len(parentItems)-1]
		// build a new list from the remaining items at the current level of the stack
		sublist, err := newList(buffer[stack[i]], nil)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to initialize a sublist")
		}
		// add this list to the parent
		parentItem.AddChild(sublist)
	}

	// log.Debugf("end of list init: current type=%v / previous type=%v / buffer= %v", currentType, previousType, buffer)
	// finally, the top-level list
	return newList(buffer[rootType], attributes)
}

func newList(items []ListItem, attributes []interface{}) (List, error) {
	// log.Debugf("initializing a new list with %d items", len(items))
	if len(items) == 0 {
		return nil, errors.Errorf("cannot build a list from an empty slice")
	}
	switch items[0].(type) {
	case *OrderedListItem:
		return NewOrderedList(items, attributes)
	case *UnorderedListItem:
		return NewUnorderedList(items, attributes)
	case *LabeledListItem:
		return NewLabeledList(items, attributes)
	default:
		return nil, errors.Errorf("unsupported type of element as the root list: '%T'", items[0])
	}
}

// ------------------------------------------
// Ordered Lists
// ------------------------------------------

// OrderedList the structure for the Ordered Lists
type OrderedList struct {
	Attributes map[string]interface{}
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

// NewOrderedList initializes a new `OrderedList` from the given content
func NewOrderedList(elements []ListItem, attributes []interface{}) (OrderedList, error) {
	log.Debugf("Initializing a new OrderedList from %d element(s)...", len(elements))
	result := make([]OrderedListItem, 0)
	bufferedItemsPerLevel := make(map[int][]*OrderedListItem, 0) // buffered items for the current level
	levelPerStyle := make(map[NumberingStyle]int, 0)
	previousLevel := 0
	previousNumberingStyle := UnknownNumberingStyle
	for _, element := range elements {
		item, ok := element.(*OrderedListItem)
		if !ok {
			return OrderedList{}, errors.Errorf("element of type '%T' is not a valid unorderedlist item", element)
		}
		log.Debugf("processing list item: %v", item.Elements[0])
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
			log.Debugf("moving buffered items at level %d (%v) in parent (%v) ", previousLevel, bufferedItemsPerLevel[previousLevel-1][0].NumberingStyle, parentItem.NumberingStyle)
			childList := toOrderedList(bufferedItemsPerLevel[previousLevel-1])
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
			childList := toOrderedList(items)
			parentLayer := bufferedItemsPerLevel[level-1]
			parentItem := parentLayer[len(parentLayer)-1]
			parentItem.Elements = append(parentItem.Elements, childList)
		}
	}

	return OrderedList{
		Attributes: mergeAttributes(attributes),
		Items:      result,
	}, nil
}

func toOrderedList(items []*OrderedListItem) OrderedList {
	result := OrderedList{
		Attributes: map[string]interface{}{}, // avoid nil `attributes`
	}
	// set the position and numbering style based on the optional attributes of the first item
	if len(items) == 0 {
		return result
	}
	items[0].applyAttributes()
	for idx, item := range items {
		// log.Debugf("setting item #%d position to %d+%d", (idx + 1), bufferedItemsPerLevel[previousLevel-1][0].Position, idx)
		item.Position = items[0].Position + idx
		item.NumberingStyle = items[0].NumberingStyle
		result.Items = append(result.Items, *item)
	}
	return result
}

// OrderedListItem the structure for the ordered list items
type OrderedListItem struct {
	Level          int
	Position       int
	NumberingStyle NumberingStyle
	Elements       []interface{}
	Attributes     map[string]interface{}
}

// making sure that the `ListItem` interface is implemented by `OrderedListItem`
var _ ListItem = &OrderedListItem{}

// NewOrderedListItem initializes a new `orderedListItem` from the given content
func NewOrderedListItem(prefix OrderedListItemPrefix, elements []interface{}, attributes []interface{}) (OrderedListItem, error) {
	log.Debugf("Initializing a new OrderedListItem with attributes %v", attributes)
	p := 1 // default position
	return OrderedListItem{
		NumberingStyle: prefix.NumberingStyle,
		Level:          prefix.Level,
		Position:       p,
		Elements:       elements,
		Attributes:     mergeAttributes(attributes),
	}, nil
}

// AddChild appends the given item to the content of this OrderedListItem
func (i *OrderedListItem) AddChild(item interface{}) {
	log.Debugf("Adding item %v to %v", item, i.Elements)
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
	Attributes map[string]interface{}
	Items      []UnorderedListItem
}

// NewUnorderedList initializes a new `UnorderedList` from the given content
func NewUnorderedList(elements []ListItem, attributes []interface{}) (UnorderedList, error) {
	log.Debugf("Initializing a new UnorderedList from %d element(s)...", len(elements))
	result := make([]UnorderedListItem, 0)
	bufferedItemsPerLevel := make(map[int][]*UnorderedListItem, 0) // buffered items for the current level
	levelPerStyle := make(map[BulletStyle]int, 0)
	previousLevel := 0
	previousBulletStyle := UnknownBulletStyle
	for _, element := range elements {
		item, ok := element.(*UnorderedListItem)
		if !ok {
			return UnorderedList{}, errors.Errorf("element of type '%T' is not a valid unorderedlist item", element)
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
			log.Debugf("merging previously buffered items at level %d in parent", previousLevel)
			parentLayer := bufferedItemsPerLevel[previousLevel-2]
			parentItem := parentLayer[len(parentLayer)-1]
			childList := UnorderedList{
				Attributes: map[string]interface{}{}, // avoid nil `attributes`
			}
			for _, i := range bufferedItemsPerLevel[previousLevel-1] {
				childList.Items = append(childList.Items, *i)
			}
			parentItem.Elements = append(parentItem.Elements, childList)
			// clear the previously buffered items at level 'previousLevel'
			delete(bufferedItemsPerLevel, previousLevel-1)
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
				Attributes: map[string]interface{}{}, // avoid nil `attributes`
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
		Attributes: mergeAttributes(attributes),
		Items:      result,
	}, nil
}

// UnorderedListItem the structure for the unordered list items
type UnorderedListItem struct {
	Level       int
	BulletStyle BulletStyle
	Elements    []interface{}
}

// NewUnorderedListItem initializes a new `UnorderedListItem` from the given content
func NewUnorderedListItem(prefix UnorderedListItemPrefix, elements []interface{}) (UnorderedListItem, error) {
	log.Debugf("Initializing a new UnorderedListItem...")
	// log.Debugf("Initializing a new UnorderedListItem with '%d' lines (%T) and input level '%d'", len(elements), elements, lvl.Len())
	return UnorderedListItem{
		Level:       prefix.Level,
		BulletStyle: prefix.BulletStyle,
		Elements:    elements,
	}, nil
}

// AddChild appends the given item to the content of this UnorderedListItem
func (i *UnorderedListItem) AddChild(item interface{}) {
	i.Elements = append(i.Elements, item)
}

// adjustBulletStyle
func (i *UnorderedListItem) adjustBulletStyle(p BulletStyle) {
	n := i.BulletStyle.nextLevelStyle(p)
	log.Debugf("adjusting bullet style for item with level %v (previous=%v) to %v", i.BulletStyle, p, n)
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
	// log.Debugf("Initializing a new ListItemContent with %d line(s)", len(content))
	elements := make([]interface{}, 0)
	for _, element := range content {
		// log.Debugf("Processing line element of type %T", element)
		switch element := element.(type) {
		case []interface{}:
			for _, e := range element {
				// if e, ok := e.(interface{}); ok {
				elements = append(elements, e)
				// }
			}
		case interface{}:
			elements = append(elements, element)
		}
	}
	// log.Debugf("Initialized a new ListItemContent with %d elements(s)", len(elements))
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
	Attributes map[string]interface{}
	Items      []LabeledListItem
}

// NewLabeledList initializes a new `LabeledList` from the given content
func NewLabeledList(elements []ListItem, attributes []interface{}) (LabeledList, error) {
	log.Debugf("Initializing a new LabeledList from %d elements", len(elements))
	items := make([]LabeledListItem, 0)
	for _, element := range elements {
		if item, ok := element.(*LabeledListItem); ok {
			items = append(items, *item)
		}
	}
	log.Debugf("Initialized a new LabeledList with %d root item(s)", len(items))
	return LabeledList{
		Attributes: mergeAttributes(attributes),
		Items:      items,
	}, nil
}

// LabeledListItem an item in a labeled
type LabeledListItem struct {
	Term     string
	Elements []interface{}
}

// NewLabeledListItem initializes a new LabeledListItem
func NewLabeledListItem(term []interface{}, elements []interface{}) (LabeledListItem, error) {
	log.Debugf("Initializing a new LabeledListItem with %d elements (%T)", len(elements), elements)
	t, err := stringify(term)
	if err != nil {
		return LabeledListItem{}, errors.Wrapf(err, "unable to get term while instanciating a new LabeledListItem element")
	}
	return LabeledListItem{
		Term:     t,
		Elements: elements,
	}, nil
}

// AddChild appends the given item to the content of this LabeledListItem
func (i *LabeledListItem) AddChild(item interface{}) {
	log.Debugf("Adding item %v to %v", item, i.Elements)
	i.Elements = append(i.Elements, item)
}

// making sure that the `ListItem` interface is implemented by `LabeledListItem`
var _ ListItem = &LabeledListItem{}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraphs
type Paragraph struct {
	Attributes map[string]interface{}
	Lines      []InlineElements
}

// NewParagraph initializes a new `Paragraph`
func NewParagraph(lines []interface{}, attributes []interface{}) (Paragraph, error) {
	log.Debugf("Initializing a new Paragraph with %d line(s)", len(lines))
	attrbs := NewElementAttributes(attributes)
	elements := make([]InlineElements, 0)
	for _, line := range lines {
		if lineElements, ok := line.([]interface{}); ok {
			for _, lineElement := range lineElements {
				if lineElement, ok := lineElement.(InlineElements); ok {
					// log.Debugf(" processing paragraph line of type %T", lineElement)
					elements = append(elements, lineElement)
				}
			}
		}
	}
	return Paragraph{
		Attributes: attrbs,
		Lines:      elements,
	}, nil
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

func (e InlineElements) lastElement() interface{} {
	if len(e) > 0 {
		return e[len(e)-1]
	}
	return nil
}

// NewInlineElements initializes a new `InlineElements` from the given values
func NewInlineElements(elements []interface{}) (InlineElements, error) {
	result := mergeElements(elements)
	// // trim righ spaces on content of last element if it is a StringElement
	// lastElement := result.lastElement()
	// if lastElement, ok := lastElement.(StringElement); ok {

	// }
	return result, nil
}

// Accept implements Visitable#Accept(Visitor)
func (e InlineElements) Accept(v Visitor) error {
	err := v.BeforeVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting inline content")
	}
	err = v.Visit(e)
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
	err = v.AfterVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting sectionTitle")
	}
	return nil
}

// ------------------------------------------
// Cross References
// ------------------------------------------

// CrossReference the struct for Cross References
type CrossReference struct {
	ID string
}

// NewCrossReference initializes a new `CrossReference` from the given ID
func NewCrossReference(id string) (CrossReference, error) {
	log.Debugf("Initializing a new CrossReference with ID=%s", id)
	return CrossReference{ID: id}, nil
}

// ------------------------------------------
// Images
// ------------------------------------------

// BlockImage the structure for the block images
type BlockImage struct {
	Macro      ImageMacro
	Attributes map[string]interface{}
}

// NewBlockImage initializes a new `BlockImage`
func NewBlockImage(imageMacro ImageMacro, attributes []interface{}) (BlockImage, error) {
	return BlockImage{
		Macro:      imageMacro,
		Attributes: NewElementAttributes(attributes),
	}, nil
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Macro ImageMacro
}

// NewInlineImage initializes a new `InlineImage` (similar to BlockImage, but without attributes)
func NewInlineImage(imageMacro ImageMacro) (InlineImage, error) {
	return InlineImage{
		Macro: imageMacro,
	}, nil
}

// ImageMacro the structure for the block image macros
type ImageMacro struct {
	Path   string
	Alt    string
	Width  *string
	Height *string
}

// NewImageMacro initializes a new `ImageMacro`
func NewImageMacro(path string, attributes interface{}) (ImageMacro, error) {
	var alt string
	var width, height *string
	if attributes != nil {
		// optionally, the width and height can be specified in the alt text, using `,` as a separator
		// eg: `image::foo.png[a title,200,100]`
		splittedAttributes := strings.Split(attributes.(string), ",")
		// naively assume that if the splitted 'alt' contains more than 3 elements, the 2 last ones are for the width and height
		splitCount := len(splittedAttributes)
		alt = splittedAttributes[0]
		if splitCount > 1 {
			w := strings.Trim(splittedAttributes[1], " ")
			width = &w
		}
		if splitCount > 2 {
			h := strings.Trim(splittedAttributes[2], " ")
			height = &h
		}
	} else {
		dir := filepath.Dir(path)
		extension := filepath.Ext(path)
		var offset int
		if dir == "." {
			offset = 0
		} else {
			offset = len(dir) + 1
		}
		alt = path[offset : len(path)-len(extension)]
	}
	return ImageMacro{
		Path:   path,
		Alt:    alt,
		Width:  width,
		Height: height,
	}, nil
}

// ------------------------------------------
// Delimited blocks
// ------------------------------------------

// DelimitedBlockKind the type for delimited blocks
type DelimitedBlockKind int

const (
	// FencedBlock a fenced block
	FencedBlock DelimitedBlockKind = iota
	// ListingBlock a listing block
	ListingBlock
	// ExampleBlock an example block
	ExampleBlock
)

// DelimitedBlock the structure for the delimited blocks
type DelimitedBlock struct {
	Kind       DelimitedBlockKind
	Attributes map[string]interface{}
	Elements   []interface{}
}

// NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given content
func NewDelimitedBlock(kind DelimitedBlockKind, content []interface{}, attributes []interface{}) (DelimitedBlock, error) {
	attrbs := NewElementAttributes(attributes)
	elements := filterEmptyElements(content)
	log.Debugf("Initialized a new DelimitedBlock with content=`%s`", elements)
	return DelimitedBlock{
		Kind:       kind,
		Attributes: attrbs,
		Elements:   elements,
	}, nil
}

// ------------------------------------------
// Literal blocks
// ------------------------------------------

// LiteralBlock the structure for the literal blocks
type LiteralBlock struct {
	Content string
}

// NewLiteralBlock initializes a new `DelimitedBlock` of the given kind with the given content,
// along with the given sectionTitle spaces
func NewLiteralBlock(spaces, content []interface{}) (LiteralBlock, error) {
	// concatenates the spaces with the actual content in a single 'stringified' value
	// log.Debugf("Initializing a new LiteralBlock with spaces='%v' and content=`%v`", spaces, content)
	c, err := stringify(append(spaces, content...))
	if err != nil {
		return LiteralBlock{}, errors.Wrapf(err, "unable to initialize a new literal block")
	}
	// remove "\n" or "\r\n", depending on the OS.
	blockContent := strings.TrimRight(strings.TrimRight(c, "\n"), "\r")
	log.Debugf("Initialized a new LiteralBlock with content=`%s`", blockContent)
	return LiteralBlock{
		Content: blockContent,
	}, nil
}

// ------------------------------------------
// Elements attributes
// ------------------------------------------

const (
	// AttrID the key to retrieve the ID in the element attributes
	AttrID string = "elementID"
	// AttrTitle the key to retrieve the title in the element attributes
	AttrTitle string = "title"
	// AttrLink the key to retrieve the link in the element attributes
	AttrLink string = "link"
	// AttrAdmonitionKind the key to retrieve the kind of admonition in the element attributes, if a "masquerade" is used
	AttrAdmonitionKind string = "admonitionKind"
)

// NewElementAttributes retrieves the ElementID, ElementTitle and ElementLink from the given slice of attributes
func NewElementAttributes(attributes []interface{}) map[string]interface{} {
	attrbs := make(map[string]interface{})
	for _, attrb := range attributes {
		log.Debugf("processing attribute %[1]v (%[1]T)", attrb)
		switch attrb := attrb.(type) {
		case ElementID:
			attrbs[AttrID] = attrb.Value
		case ElementTitle:
			attrbs[AttrTitle] = attrb.Value
		case AdmonitionKind:
			attrbs[AttrAdmonitionKind] = attrb
		case map[string]interface{}:
			for k, v := range attrb {
				attrbs[k] = v
			}
		case nil:
			// ignore
		default:
			log.Warnf("Unexpected attributes: %T", attrb)
		}
	}
	return attrbs
}

// ElementID the structure for element IDs
type ElementID struct {
	Value string
}

// NewElementID initializes a new `ElementID` from the given ID
func NewElementID(id string) (ElementID, error) {
	log.Debugf("Initializing a new ElementID with ID=%s", id)
	return ElementID{Value: id}, nil
}

// ElementTitle the structure for element Titles
type ElementTitle struct {
	Value string
}

// NewElementTitle initializes a new `ElementTitle` from the given value
func NewElementTitle(value []interface{}) (ElementTitle, error) {
	v, err := stringify(value)
	if err != nil {
		return ElementTitle{}, errors.Wrapf(err, "failed to initialize a new ElementTitle")
	}
	log.Debugf("Initializing a new ElementTitle with content=%s", v)
	return ElementTitle{Value: v}, nil
}

// NewAttributeGroup initializes a group of attributes from the given generic attributes.
func NewAttributeGroup(attributes []interface{}) (map[string]interface{}, error) {
	// log.Debugf("Initializing a new AttributeGroup with %v", attributes)
	result := make(map[string]interface{}, 0)
	for _, a := range attributes {
		// log.Debugf("processing attribute group element of type %T", a)
		if a, ok := a.(GenericAttribute); ok {
			for k, v := range a {
				result[k] = v
			}
		}
	}
	// log.Debugf("Initialized a new AttributeGroup: %v", result)
	return result, nil
}

// GenericAttribute the structure for single, generic attribute.
// If the attribute was specified in the form of [foo], then its key is 'foo' and its value is 'nil'.
type GenericAttribute map[string]interface{}

// NewGenericAttribute initializes a new GenericAttribute from the given key and optional value
func NewGenericAttribute(key []interface{}, value []interface{}) (GenericAttribute, error) {
	result := make(map[string]interface{})
	k, err := stringify(key,
		// remove surrounding quotes
		func(s string) (string, error) {
			return strings.Trim(s, "\""), nil
		})
	if err != nil {
		return GenericAttribute{}, errors.Wrapf(err, "failed to initialize a new generic attribute")
	}
	if value != nil {
		v, err := stringify(value,
			// remove surrounding quotes
			func(s string) (string, error) {
				return strings.Trim(s, "\""), nil
			})
		if err != nil {
			return GenericAttribute{}, errors.Wrapf(err, "failed to initialize a new generic attribute")
		}
		result[k] = v
	} else {
		result[k] = nil
	}
	// log.Debugf("Initialized a new GenericAttribute: %v", result)
	return result, nil

}

// InvalidElementAttribute the struct for invalid element attributes
type InvalidElementAttribute struct {
	Value string
}

// NewInvalidElementAttribute initializes a new `InvalidElementAttribute` from the given text
func NewInvalidElementAttribute(text []byte) (InvalidElementAttribute, error) {
	value := string(text)
	log.Debugf("Initializing a new InvalidElementAttribute with text=%s", value)
	return InvalidElementAttribute{Value: value}, nil
}

// ------------------------------------------
// StringElement
// ------------------------------------------

// StringElement the structure for strings
type StringElement struct {
	Content string
}

// NewStringElement initializes a new `StringElement` from the given content
func NewStringElement(content interface{}) StringElement {
	return StringElement{Content: content.(string)}
}

// Accept implements Visitable#Accept(Visitor)
func (s StringElement) Accept(v Visitor) error {
	err := v.BeforeVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting string element")
	}
	err = v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting string element")
	}
	err = v.AfterVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting string element")
	}
	return nil
}

// ------------------------------------------
// Quoted text
// ------------------------------------------

// QuotedText the structure for quoted text
type QuotedText struct {
	Kind     QuotedTextKind
	Elements []interface{}
}

// QuotedTextKind the type for
type QuotedTextKind int

const (
	// Bold bold quoted text
	Bold QuotedTextKind = iota
	// Italic italic quoted text
	Italic
	// Monospace monospace quoted text
	Monospace
)

// NewQuotedText initializes a new `QuotedText` from the given kind and content
func NewQuotedText(kind QuotedTextKind, content []interface{}) (QuotedText, error) {
	elements := mergeElements(content)
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("Initialized a new QuotedText with %d elements:", len(elements))
		spew.Dump(elements)
	}
	return QuotedText{Kind: kind, Elements: elements}, nil
}

// Accept implements Visitable#Accept(Visitor)
func (t QuotedText) Accept(v Visitor) error {
	err := v.BeforeVisit(t)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting quoted text")
	}
	err = v.Visit(t)
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
	err = v.AfterVisit(t)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting quoted text")
	}
	return nil
}

// ------------------------------------------------------
// Escaped Quoted Text (i.e., with substitution prevention)
// ------------------------------------------------------

// NewEscapedQuotedText returns a new InlineElements where the nested elements are preserved (ie, substituted as expected)
func NewEscapedQuotedText(backslashes []interface{}, punctuation string, content []interface{}) ([]interface{}, error) {
	backslashesStr, err := stringify(backslashes,
		func(s string) (string, error) {
			// remove the number of back-slashes that match the length of the punctuation. Eg: `\*` or `\\**`, but keep extra back-slashes
			if len(s) > len(punctuation) {
				return s[len(punctuation):], nil
			}
			return "", nil
		})
	if err != nil {
		return []interface{}{}, errors.Wrapf(err, "error while initializing quoted text with substitution prevention")
	}
	return []interface{}{backslashesStr, punctuation, content, punctuation}, nil
}

// ------------------------------------------
// Passthrough
// ------------------------------------------

// Passthrough the structure for Passthroughs
type Passthrough struct {
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

// NewPassthrough returns a new passthrough
func NewPassthrough(kind PassthroughKind, elements []interface{}) (Passthrough, error) {
	return Passthrough{
		Kind:     kind,
		Elements: mergeElements(elements),
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
	// log.Debug("Initializing a new BlankLine")
	return BlankLine{}, nil
}

// ------------------------------------------
// Links
// ------------------------------------------

// Link the structure for the external links
type Link struct {
	URL  string
	Text string
}

// NewLink initializes a new `Link`
func NewLink(url, text []interface{}) (Link, error) {
	urlStr, err := stringify(url)
	if err != nil {
		return Link{}, errors.Wrapf(err, "failed to initialize a new Link element")
	}
	textStr, err := stringify(text, // remove "\n" or "\r\n", depending on the OS.
		// remove heading "[" and traingin "]"
		func(s string) (string, error) {
			return strings.TrimPrefix(s, "["), nil
		},
		func(s string) (string, error) {
			return strings.TrimSuffix(s, "]"), nil
		})
	if err != nil {
		return Link{}, errors.Wrapf(err, "failed to initialize a new Link element")
	}
	return Link{URL: urlStr, Text: textStr}, nil
}

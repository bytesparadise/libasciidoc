package types

import (
	"fmt"
	"path/filepath"
	"strings"

	"flag"

	"reflect"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	debugMode := flag.Bool("debug", false, "when set, enables debug log messages")
	flag.Parse()
	if *debugMode {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

// *****************************
// DocElement
// *****************************

//DocElement the interface for all document elements
type DocElement interface {
	String() string
}

// *****************************
// Document
// *****************************

//Document the top-level structure for a a document
type Document struct {
	Elements []DocElement
}

//NewDocument initializes a new `Document` from the given lines
func NewDocument(lines []interface{}) (*Document, error) {
	elements := make([]DocElement, len(lines))
	for i := range lines {
		elements[i] = lines[i].(DocElement)
	}
	return &Document{Elements: elements}, nil
}

func (d *Document) String() string {
	// todo : use a bufferwriter
	result := ""
	for i := range d.Elements {
		result = result + "\n" + d.Elements[i].String()
	}
	return result
}

// *****************************
// Heading
// *****************************

// Heading the structure for the headings
type Heading struct {
	Level   int
	Content *InlineContent
}

//NewHeading initializes a new `Heading from the given level and content
func NewHeading(level interface{}, inlineContent *InlineContent) (*Heading, error) {
	// counting the lenght of the 'level' value (ie, the number of `=` chars)
	actualLevel := len(level.([]interface{}))
	heading := Heading{Level: actualLevel, Content: inlineContent}
	log.Debugf("New heading: %v", heading)
	return &heading, nil
}

func (h Heading) String() string {
	return fmt.Sprintf("<Heading %d> '%s'", h.Level, h.Content.String())
}

// *****************************
// ListItem
// *****************************

// ListItem the structure for the list items
type ListItem struct {
	Content *InlineContent
}

//NewListItem initializes a new `ListItem` from the given content
func NewListItem(content *InlineContent) (*ListItem, error) {
	log.Debugf("New list item based on %v", content)
	// items := strings.Split(value, " ")
	// number, err := strconv.Atoi(items[0])
	// if err != nil {
	// 	return nil, errs.Wrapf(err, "failed to create new ListItem")
	// }
	// content := items[1]
	return &ListItem{
		// Number:  number,
		Content: content,
	}, nil
}

func (l ListItem) String() string {
	return fmt.Sprintf("<List item> %v", *l.Content)
}

// *****************************
// Paragraph
// *****************************

// Paragraph the structure for the paragraph
type Paragraph struct {
	// Input    []byte
	Lines []*InlineContent
}

//NewParagraph initializes a new `Paragraph`
func NewParagraph(text []byte, lines []interface{}) (*Paragraph, error) {
	log.Debugf("New paragraph with %d lines: ", len(lines))
	typedLines := make([]*InlineContent, 0)
	for _, line := range lines {
		typedLines = append(typedLines, line.(*InlineContent))
	}
	// return &InlineContent{Input: text, Elements: mergedElements}, nil
	return &Paragraph{Lines: typedLines}, nil
}

func (p Paragraph) String() string {
	return fmt.Sprintf("<Paragraph> %[1]v", p.Lines)
}

// InlineContent the structure for the lines in paragraphs
type InlineContent struct {
	// Input    []byte
	Elements []DocElement
}

//NewInlineContent initializes a new `InlineContent` from the given values
func NewInlineContent(text []byte, elements []interface{}) (*InlineContent, error) {
	log.Debugf("New paragraph lines based on %d values: ", len(elements))
	mergedElements := make([]DocElement, 0)
	for _, e := range merge(elements) {
		mergedElements = append(mergedElements, e.(DocElement))
	}
	log.Debugf("New InlineContent: %v (%d)", mergedElements, len(mergedElements))
	// return &InlineContent{Input: text, Elements: mergedElements}, nil
	return &InlineContent{Elements: mergedElements}, nil
}

func (c InlineContent) String() string {
	return fmt.Sprintf("<InlineContent (l=%[2]d)> %[1]v", c.Elements, len(c.Elements))
}

// *****************************
// Images
// *****************************

// BlockImage the structure for the block images
type BlockImage struct {
	Macro BlockImageMacro
	ID    *ElementID
	Title *ElementTitle
	Link  *ElementLink
}

//NewBlockImage initializes a new `BlockImage`
func NewBlockImage(input []byte, imageMacro BlockImageMacro, metadata []interface{}) (*BlockImage, error) {
	log.Debugf("Initializing a new BlockImage from '%s'", input)
	var id *ElementID
	var title *ElementTitle
	var link *ElementLink
	for _, item := range metadata {
		switch item := item.(type) {
		case *ElementID:
			id = item
		case *ElementLink:
			link = item
		case *ElementTitle:
			title = item
		default:
			log.Warnf("Unexpected metadata: %s", reflect.TypeOf(item))
		}
	}
	return &BlockImage{
		Macro: imageMacro,
		ID:    id,
		Title: title,
		Link:  link,
	}, nil
}

func (img BlockImage) String() string {
	return "<BlockImage>" + img.Macro.String()
}

// BlockImageMacro the structure for the block image macros
type BlockImageMacro struct {
	Path   string
	Alt    string
	Width  *string
	Height *string
}

//NewBlockImageMacro initializes a new `BlockImageMacro`
func NewBlockImageMacro(input []byte, path string, attributes *string) (*BlockImageMacro, error) {
	var alt string
	var width, height *string
	if attributes != nil {
		// optionally, the width and height can be specified in the alt text, using `,` as a separator
		// eg: `image::foo.png[a title,200,100]`
		splittedAttributes := strings.Split(*attributes, ",")
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
	log.Debugf("Initializing a new BlockImageMacro from '%s'", input)
	return &BlockImageMacro{
		Path:   path,
		Alt:    alt,
		Width:  width,
		Height: height}, nil
}

func (m BlockImageMacro) String() string {
	var width, height string
	if m.Width != nil {
		width = *m.Width
	}
	if m.Height != nil {
		height = *m.Height
	}
	return fmt.Sprintf("<BlockImageMacroMacro> %s[%s,w=%s h=%s]", m.Path, m.Alt, width, height)
}

// *****************************
// Meta Elements
// *****************************

// ElementLink the structure for element links
type ElementLink struct {
	Path string
}

//NewElementLink initializes a new `ElementLink` from the given path
func NewElementLink(path string) (*ElementLink, error) {
	log.Debugf("New ElementLink with path=%s", path)
	return &ElementLink{Path: path}, nil
}

func (e ElementLink) String() string {
	return fmt.Sprintf("<ElementLink> %s", e.Path)
}

// ElementID the structure for element IDs
type ElementID struct {
	Value string
}

//NewElementID initializes a new `ElementID` from the given path
func NewElementID(id string) (*ElementID, error) {
	log.Debugf("New ElementID with ID=%s", id)
	return &ElementID{Value: id}, nil
}

func (e ElementID) String() string {
	return fmt.Sprintf("<ElementID> %s", e.Value)
}

// ElementTitle the structure for element IDs
type ElementTitle struct {
	Content string
}

//NewElementTitle initializes a new `ElementTitle` from the given content
func NewElementTitle(content []interface{}) (*ElementTitle, error) {
	c, err := stringify(content)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ElementTitle")
	}
	log.Debugf("New ElementTitle with content=%s", c)
	return &ElementTitle{Content: *c}, nil
}

func (e ElementTitle) String() string {
	return fmt.Sprintf("<ElementTitle> %s", e.Content)
}

// *****************************
// StringElement
// *****************************

// StringElement the structure for strings
type StringElement struct {
	Content string
}

//NewStringElement initializes a new `StringElement` from the given content
func NewStringElement(content interface{}) *StringElement {
	return &StringElement{Content: content.(string)}
}

func (e StringElement) String() string {
	return fmt.Sprintf("<String> '%s' (%d)", e.Content, len(e.Content))
}

// *****************************
// Quoted text
// *****************************

// QuotedText the structure for quoted text
type QuotedText struct {
	Kind     QuotedTextKind
	Elements []DocElement
}

// QuotedTextKind the type for
type QuotedTextKind int

const (
	// Bold bold quoted text
	Bold = iota
	// Italic italic quoted text
	Italic
	// Monospace monospace quoted text
	Monospace
)

//NewQuotedText initializes a new `QuotedText` from the given kind and content
func NewQuotedText(kind QuotedTextKind, content []interface{}) (*QuotedText, error) {
	elements, err := toDocElements(merge(content))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initialize a new QuotedText")
	}
	log.Debugf("Initializing a new QuotedText with %d elements:", len(elements))
	for _, element := range elements {
		log.Debugf("- %v (%v)", element, reflect.TypeOf(element))
	}
	return &QuotedText{Kind: kind, Elements: elements}, nil
}

func (t QuotedText) String() string {
	return fmt.Sprintf("<QuotedText (%d)> %v", t.Kind, t.Elements)
}

// *****************************
// BlankLine
// *****************************

// BlankLine the structure for the empty lines, which are used to separate logical blocks
type BlankLine struct {
}

//NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (*BlankLine, error) {
	return &BlankLine{}, nil
}

func (e BlankLine) String() string {
	return "<BlankLine>"
}

// *****************************
// Links
// *****************************

// ExternalLink the structure for the external links
type ExternalLink struct {
	URL  string
	Text string
}

//NewExternalLink initializes a new `ExternalLink`
func NewExternalLink(url, text []interface{}) (*ExternalLink, error) {
	urlStr, err := stringify(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ExternalLink element")
	}
	textStr, err := stringify(text)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ExternalLink element")
	}
	// the text includes the surrounding '[' and ']' which should be removed
	trimmedText := strings.TrimPrefix(strings.TrimSuffix(*textStr, "]"), "[")
	return &ExternalLink{URL: *urlStr, Text: trimmedText}, nil
}

func (e ExternalLink) String() string {
	return fmt.Sprintf("<ExternalLink> %s[%s]", e.URL, e.Text)
}

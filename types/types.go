package types

import (
	"fmt"
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
	log.Info(fmt.Sprintf("New heading: %v", heading))
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
	log.Debug(fmt.Sprintf("New list item based on %v", content))
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
// InlineContent
// *****************************

// InlineContent the structure for the inline content
type InlineContent struct {
	Elements []DocElement
}

//NewInlineContent initializes a new `InlineContent` from the given values
func NewInlineContent(elements []interface{}) (*InlineContent, error) {
	log.Debug(fmt.Sprintf("New inline content based on %d values: ", len(elements)))
	mergedElements := make([]DocElement, 0)
	for _, e := range merge(elements) {
		mergedElements = append(mergedElements, e.(DocElement))
	}
	log.Debug(fmt.Sprintf("New InlineContent: %v (%d)", mergedElements, len(mergedElements)))
	return &InlineContent{Elements: mergedElements}, nil
}

func (c InlineContent) String() string {
	return fmt.Sprintf("<InlineContent (l=%[2]d)> %[1]v", c.Elements, len(c.Elements))
}

// -----------------------------
// Meta Elements
// -----------------------------

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
	ID string
}

//NewElementID initializes a new `ElementID` from the given path
func NewElementID(id string) (*ElementID, error) {
	log.Debugf("New ElementID with ID=%s", id)
	return &ElementID{ID: id}, nil
}

func (e ElementID) String() string {
	return fmt.Sprintf("<ElementID> %s", e.ID)
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

// -----------------------------
// StringElement
// -----------------------------

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

// -----------------------------
// Quoted text
// -----------------------------

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

// -----------------------------
// EmptyLine
// -----------------------------

// EmptyLine the structure for the empty lines, which are used to separate logical blocks
type EmptyLine struct {
}

//NewEmptyLine initializes a new `EmptyLine`
func NewEmptyLine() (*EmptyLine, error) {
	return &EmptyLine{}, nil
}

func (e EmptyLine) String() string {
	return "<EmptyLine>"
}

// -----------------------------
// Links
// -----------------------------

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

// -----------------------------
// Images
// -----------------------------

// BlockImage the structure for the block image macross
type BlockImage struct {
	Path    string
	AltText *string
	Width   *string
	Height  *string
}

//NewBlockImage initializes a new `BlockImageMacro`
func NewBlockImage(path string) (*BlockImage, error) {
	return &BlockImage{
		Path: path,
	}, nil
}

//NewBlockImageWithAltText initializes a new `BlockImageMacro`
func NewBlockImageWithAltText(path string, altText string) (*BlockImage, error) {
	var alt, width, height *string
	// optionally, the width and height can be specified in the alt text, using `,` as a separator
	// eg: `image::foo.png[a title,200,100]`
	splittedAltText := strings.Split(altText, ",")
	// naively assume that if the splitted 'alt' contains more than 3 elements, the 2 last ones are for the width and height
	splitCount := len(splittedAltText)
	if splitCount >= 3 {
		actualAltText := strings.Join(splittedAltText[0:splitCount-2], ",")
		alt = &actualAltText
		w := splittedAltText[splitCount-2]
		width = &w
		h := splittedAltText[splitCount-1]
		height = &h
	} else {
		alt = &altText
	}
	return &BlockImage{
		Path:    path,
		AltText: alt,
		Width:   width,
		Height:  height}, nil
}

// func NewBlockImage(path string, altText string) (*BlockImage, error) {
// 	var alt, width, height *string
// 	alt, err := stringify(altText)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "failed to create block image")
// 	}
// 	// the text includes the surrounding '[' and ']' which should be removed
// 	trimmedAltText := strings.TrimSuffix(strings.TrimPrefix(*alt, "["), "]")
// 	// optionally, the width and height can be specified in the alt text, using `,` as a separator
// 	// eg: `image::foo.png[a title,200,100]`
// 	splittedAltText := strings.Split(trimmedAltText, ",")
// 	// naively assume that if the splitted 'alt' contains more than 3 elements, the 2 last ones are for the width and height
// 	splitCount := len(splittedAltText)
// 	if splitCount >= 3 {
// 		actualAltText := strings.Join(splittedAltText[0:splitCount-2], ",")
// 		alt = &actualAltText
// 		w := splittedAltText[splitCount-2]
// 		width = &w
// 		h := splittedAltText[splitCount-1]
// 		height = &h
// 	}

// 	return &BlockImage{
// 		Path:    path,
// 		AltText: alt,
// 		Width:   width,
// 		Height:  height}, nil
// }

func (m BlockImage) String() string {
	var altText, width, height string
	if m.AltText != nil {
		altText = *m.AltText
	}
	if m.Width != nil {
		width = *m.Width
	}
	if m.Height != nil {
		height = *m.Height
	}
	return fmt.Sprintf("<BlockImageMacro> %s[%s,w=%s h=%s]", m.Path, altText, width, height)
}

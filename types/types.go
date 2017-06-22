package types

import (
	"fmt"
	"strings"

	"flag"

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
	return fmt.Sprintf("<String> %s (%d)", e.Content, len(e.Content))
}

// *****************************
// BoldQuote
// *****************************

// BoldQuote the structure for the bold quotes
type BoldQuote struct {
	Content string
}

//NewBoldQuote initializes a new `BoldQuote` from the given content
func NewBoldQuote(content interface{}) (*BoldQuote, error) {
	return &BoldQuote{Content: content.(string)}, nil
}

func (b BoldQuote) String() string {
	return fmt.Sprintf("<BoldQuote> %v", b.Content)
}

// *****************************
// EmptyLine
// *****************************

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

// *****************************
// ExternalLink
// *****************************

// ExternalLink the structure for the external links
type ExternalLink struct {
	URL  string
	Text string
}

//NewExternalLink initializes a new `ExternalLink`
func NewExternalLink(url, text []interface{}) (*ExternalLink, error) {
	u := stringify(merge(url))
	t := stringify(merge(text))
	// the text includes the surrounding '[' and ']' which should be removed
	t = strings.TrimPrefix(t, "[")
	t = strings.TrimSuffix(t, "]")
	return &ExternalLink{URL: u, Text: t}, nil
}

func (e ExternalLink) String() string {
	return fmt.Sprintf("<ExternalLink> %s[%s]", e.URL, e.Text)
}

// *****************************
// BlockImage
// *****************************

// BlockImage the structure for the block images
type BlockImage struct {
	Path    string
	AltText *string
	Width   *string
	Height  *string
}

//NewBlockImage initializes a new `BlockImage`
func NewBlockImage(path string, altText []interface{}) (*BlockImage, error) {
	var width, height *string
	alt := stringify(merge(altText))
	// the text includes the surrounding '[' and ']' which should be removed
	alt = strings.TrimPrefix(alt, "[")
	alt = strings.TrimSuffix(alt, "]")
	// if no alttext was provided
	if len(alt) == 0 {
		return &BlockImage{
			Path: path,
		}, nil
	}
	// optionally, the width and height can be specified in the alt text, using `,` as a separator
	// eg: `image::foo.png[a title,200,100]`
	splittedAlt := strings.Split(alt, ",")
	// naively assume that if the splitted 'alt' contains more than 3 elements, the 2 last ones are for the width and height
	splitCount := len(splittedAlt)
	if splitCount >= 3 {
		alt = strings.Join(splittedAlt[0:splitCount-2], ",")
		w := splittedAlt[splitCount-2]
		width = &w
		h := splittedAlt[splitCount-1]
		height = &h
	}

	return &BlockImage{
		Path:    path,
		AltText: &alt,
		Width:   width,
		Height:  height}, nil
}

func (i BlockImage) String() string {
	var altText, width, height string
	if i.AltText != nil {
		altText = *i.AltText
	}
	if i.Width != nil {
		width = *i.Width
	}
	if i.Height != nil {
		height = *i.Height
	}
	return fmt.Sprintf("<BlockImage> %s[%s,w=%s h=%s]", i.Path, altText, width, height)
}

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

// ------------------------------------------
// DocElement (and other interfaces)
// ------------------------------------------

//DocElement the interface for all document elements
// TODO: 'String()' remove this method ? no real value here (we could use a visitor to print/debug the elements), by having a `Visit(Visitor)` method instead
type DocElement interface {
	Visitable
	String() string
}

type Visitable interface {
	Accept(Visitor) error
}

// Visitor a visitor that can visit/traverse the DocElement and its children (if applicable)
type Visitor interface {
	BeforeVisit(interface{}) error
	Visit(interface{}) error
	AfterVisit(interface{}) error
}

// ------------------------------------------
// Document
// ------------------------------------------

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

//String implements the DocElement#String() method
func (d *Document) String() string {
	// todo : use a bufferwriter
	result := ""
	for i := range d.Elements {
		result = result + "\n" + d.Elements[i].String()
	}
	return result
}

// ------------------------------------------
// Heading
// ------------------------------------------

// Heading the structure for the headings
type Heading struct {
	ID      *ElementID
	Level   int
	Content *InlineContent
}

//NewHeading initializes a new `Heading from the given level and content, with the optional metadata.
// In the metadata, only the ElementID is retained
func NewHeading(level interface{}, inlineContent *InlineContent, metadata []interface{}) (*Heading, error) {
	// counting the lenght of the 'level' value (ie, the number of `=` chars)
	actualLevel := len(level.([]interface{}))
	id, _, _ := newMetaElements(metadata)
	if id == nil {
		v := NewReplaceNonAlphanumericsVisitor()
		err := inlineContent.Accept(v)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate default ID while instanciating a new Heading element")
		}

		id, _ = NewElementID(v.NormalizedContent())
	}
	heading := Heading{Level: actualLevel, Content: inlineContent, ID: id}
	log.Debugf("New Heading: %v", heading)
	return &heading, nil
}

//String implements the DocElement#String() method
func (h Heading) String() string {
	return fmt.Sprintf("<Heading %d> '%s'", h.Level, h.Content.String())
}

//Accept implements DocElement#Accept(Visitor)
func (h Heading) Accept(v Visitor) error {
	err := v.BeforeVisit(h)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting heading")
	}
	err = h.Content.Accept(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting heading")
	}
	err = v.AfterVisit(h)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting heading")
	}
	return nil
}

// ------------------------------------------
// ListItem
// ------------------------------------------

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

//String implements the DocElement#String() method
func (i ListItem) String() string {
	return fmt.Sprintf("<List item> %v", *i.Content)
}

//Accept implements DocElement#Accept(Visitor)
func (i ListItem) Accept(v Visitor) error {
	err := v.BeforeVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting list item")
	}
	err = v.Visit(i)
	if err != nil {
		return errors.Wrapf(err, "error while visiting list item")
	}
	err = i.Content.Accept(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting list item content")
	}
	err = v.AfterVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting list item")
	}
	return nil
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

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

//String implements the DocElement#String() method
func (p Paragraph) String() string {
	return fmt.Sprintf("<Paragraph> %[1]v", p.Lines)
}

//Accept implements DocElement#Accept(Visitor)
func (p Paragraph) Accept(v Visitor) error {
	err := v.BeforeVisit(p)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting pararaph")
	}
	err = v.Visit(p)
	if err != nil {
		return errors.Wrapf(err, "error while visiting pararaph")
	}
	for _, line := range p.Lines {
		err := line.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting paragraph line")
		}

	}
	err = v.AfterVisit(p)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting pararaph")
	}
	return nil
}

// ------------------------------------------
// InlineContent
// ------------------------------------------

// InlineContent the structure for the lines in paragraphs
type InlineContent struct {
	// Input    []byte
	Elements []DocElement
}

//NewInlineContent initializes a new `InlineContent` from the given values
func NewInlineContent(text []byte, elements []interface{}) (*InlineContent, error) {
	mergedElements := make([]DocElement, 0)
	for _, e := range merge(elements) {
		mergedElements = append(mergedElements, e.(DocElement))
	}
	log.Debugf("New InlineContent: %v (%d)", mergedElements, len(mergedElements))
	return &InlineContent{Elements: mergedElements}, nil
}

//String implements the DocElement#String() method
func (c InlineContent) String() string {
	return fmt.Sprintf("<InlineContent (l=%[2]d)> %[1]v", c.Elements, len(c.Elements))
}

//Accept implements DocElement#Accept(Visitor)
func (c InlineContent) Accept(v Visitor) error {
	err := v.BeforeVisit(c)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting inline content")
	}
	err = v.Visit(c)
	if err != nil {
		return errors.Wrapf(err, "error while visiting inline content")
	}
	for _, element := range c.Elements {
		err = element.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting inline content element")
		}

	}
	err = v.AfterVisit(c)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting heading")
	}
	return nil
}

// ------------------------------------------
// Images
// ------------------------------------------

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
	id, title, link := newMetaElements(metadata)
	return &BlockImage{
		Macro: imageMacro,
		ID:    id,
		Title: title,
		Link:  link,
	}, nil
}

//String implements the DocElement#String() method
func (img BlockImage) String() string {
	return "<BlockImage>" + img.Macro.String()
}

func (img BlockImage) elements() []Visitable {
	return []Visitable{img.ID, img.Link, img.Macro, img.Title}
}

//Accept implements DocElement#Accept(Visitor)
func (img BlockImage) Accept(v Visitor) error {
	err := v.BeforeVisit(img)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting block image")
	}
	err = v.Visit(img)
	if err != nil {
		return errors.Wrapf(err, "error while visiting block image")
	}
	for _, element := range img.elements() {
		err := element.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting block image element")
		}

	}
	err = v.AfterVisit(img)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting block image")
	}
	return nil
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

//String implements the DocElement#String() method
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

//Accept implements DocElement#Accept(Visitor)
func (m BlockImageMacro) Accept(v Visitor) error {
	err := v.BeforeVisit(m)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting block image macro")
	}
	err = v.Visit(m)
	if err != nil {
		return errors.Wrapf(err, "error while visiting block image macro")
	}
	err = v.AfterVisit(m)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting block image macro")
	}
	return nil
}

// ------------------------------------------
// Delimited blocks
// ------------------------------------------

// DelimitedBlockKind the type for delimited blocks
type DelimitedBlockKind int

const (
	// SourceBlock a source block
	SourceBlock DelimitedBlockKind = iota
)

//DelimitedBlock the structure for the delimited blocks
type DelimitedBlock struct {
	Kind    DelimitedBlockKind
	Content string
}

//NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given content
func NewDelimitedBlock(kind DelimitedBlockKind, content []interface{}) (*DelimitedBlock, error) {
	c, err := stringify(content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initialize a new delimited block")
	}
	return &DelimitedBlock{
		Kind:    kind,
		Content: strings.TrimSuffix(strings.TrimSuffix(*c, "\n"), "\r"), // remove "\n" or "\r\n", depending on the OS.
	}, nil
}

//String implements the DocElement#String() method
func (b DelimitedBlock) String() string {
	switch b.Kind {
	case SourceBlock:
		return fmt.Sprintf("<SourceBlock> %v", b.Content)
	default:
		return fmt.Sprintf("<Unknown type of block> %v", b.Content)
	}
}

//Accept implements DocElement#Accept(Visitor)
func (b DelimitedBlock) Accept(v Visitor) error {
	err := v.BeforeVisit(b)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting delimited block")
	}
	err = v.Visit(b)
	if err != nil {
		return errors.Wrapf(err, "error while visiting delimited block")
	}
	err = v.AfterVisit(b)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting delimited block")
	}
	return nil
}

// ------------------------------------------
// Meta Elements
// ------------------------------------------

func newMetaElements(metadata []interface{}) (*ElementID, *ElementTitle, *ElementLink) {
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
	return id, title, link
}

// ElementLink the structure for element links
type ElementLink struct {
	Path string
}

//NewElementLink initializes a new `ElementLink` from the given path
func NewElementLink(path string) (*ElementLink, error) {
	log.Debugf("New ElementLink with path=%s", path)
	return &ElementLink{Path: path}, nil
}

//String implements the DocElement#String() method
func (e ElementLink) String() string {
	return fmt.Sprintf("<ElementLink> %s", e.Path)
}

//Accept implements DocElement#Accept(Visitor)
func (e ElementLink) Accept(v Visitor) error {
	err := v.BeforeVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting element link")
	}
	err = v.Visit(e)
	if err != nil {
		return errors.Wrapf(err, "error while visiting element link")
	}
	err = v.AfterVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error whie post-visiting element link")
	}
	return nil
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

//String implements the DocElement#String() method
func (e ElementID) String() string {
	return fmt.Sprintf("<ElementID> %s", e.Value)
}

//Accept implements DocElement#Accept(Visitor)
func (e ElementID) Accept(v Visitor) error {
	err := v.BeforeVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting element ID")
	}
	err = v.Visit(e)
	if err != nil {
		return errors.Wrapf(err, "error while visiting element ID")
	}
	err = v.AfterVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting element ID")
	}

	return nil
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

//String implements the DocElement#String() method
func (e ElementTitle) String() string {
	return fmt.Sprintf("<ElementTitle> %s", e.Content)
}

//Accept implements DocElement#Accept(Visitor)
func (e ElementTitle) Accept(v Visitor) error {
	err := v.BeforeVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting element link")
	}
	err = v.Visit(e)
	if err != nil {
		return errors.Wrapf(err, "error while visiting element title")
	}
	err = v.AfterVisit(e)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting element link")
	}
	return nil
}

// ------------------------------------------
// StringElement
// ------------------------------------------

// StringElement the structure for strings
type StringElement struct {
	Content string
}

//NewStringElement initializes a new `StringElement` from the given content
func NewStringElement(content interface{}) *StringElement {
	return &StringElement{Content: content.(string)}
}

//String implements the DocElement#String() method
func (s StringElement) String() string {
	return fmt.Sprintf("<String> '%s' (%d)", s.Content, len(s.Content))
}

//Accept implements DocElement#Accept(Visitor)
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
	Elements []DocElement
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

//String implements the DocElement#String() method
func (t QuotedText) String() string {
	return fmt.Sprintf("<QuotedText (%d)> %v", t.Kind, t.Elements)
}

//Accept implements DocElement#Accept(Visitor)
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
		err := element.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting quoted text element")
		}

	}
	err = v.AfterVisit(t)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting quoted text")
	}
	return nil
}

// ------------------------------------------
// BlankLine
// ------------------------------------------

// BlankLine the structure for the empty lines, which are used to separate logical blocks
type BlankLine struct {
}

//NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (*BlankLine, error) {
	return &BlankLine{}, nil
}

//String implements the DocElement#String() method
func (l BlankLine) String() string {
	return "<BlankLine>"
}

//Accept implements DocElement#Accept(Visitor)
func (l BlankLine) Accept(v Visitor) error {
	err := v.BeforeVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting blank line")
	}
	err = v.Visit(l)
	if err != nil {
		return errors.Wrapf(err, "error while visiting blank line")
	}
	err = v.AfterVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting blank line")
	}
	return nil
}

// ------------------------------------------
// Links
// ------------------------------------------

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

//String implements the DocElement#String() method
func (l ExternalLink) String() string {
	return fmt.Sprintf("<ExternalLink> %s[%s]", l.URL, l.Text)
}

//Accept implements DocElement#Accept(Visitor)
func (l ExternalLink) Accept(v Visitor) error {
	err := v.BeforeVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting external link")
	}
	err = v.Visit(l)
	if err != nil {
		return errors.Wrapf(err, "error while visiting external link")
	}
	err = v.AfterVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting external link")
	}
	return nil
}

package types

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"reflect"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// DocElement (and other interfaces)
// ------------------------------------------

//DocElement the interface for all document elements
// TODO: 'String()' remove this method ? no real value here (we could use a visitor to print/debug the elements), by having a `Visit(Visitor)` method instead
type DocElement interface {
	Visitable
	String(int) string
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

//Document the top-level structure for a document
type Document struct {
	Elements []DocElement
}

//NewDocument initializes a new `Document` from the given lines
func NewDocument(blocks []interface{}) (*Document, error) {
	log.Debugf("Initializing a new Document with %d blocks(s)", len(blocks))
	for i, block := range blocks {
		log.Debugf("Line #%d: %v", i, reflect.TypeOf(block))
	}
	elements := make([]DocElement, len(blocks))
	for i := range blocks {
		elements[i] = blocks[i].(DocElement)
	}
	return &Document{Elements: elements}, nil
}

//String implements the DocElement#String() method
func (d *Document) String(indentLevel int) string {
	log.Debugf("Printing document...")
	result := bytes.NewBuffer(make([]byte, 0))
	for i := range d.Elements {
		result.WriteString(fmt.Sprintf("\n%s", d.Elements[i].String(0)))
	}
	return result.String()
}

// ------------------------------------------
// Section
// ------------------------------------------

//Section the structure for a section
type Section struct {
	Heading  Heading
	Elements []DocElement
}

//NewSection initializes a new `Section` from the given heading and elements
func NewSection(heading *Heading, blocks []interface{}) (*Section, error) {
	// log.Debugf("Initializing a new Section with %d block(s)", len(blocks))
	elements := make([]DocElement, len(blocks))
	for i := range blocks {
		elements[i] = blocks[i].(DocElement)
	}
	log.Debugf("Initialized a new Section of level %d with %d block(s)", heading.Level, len(blocks))
	return &Section{
		Heading:  *heading,
		Elements: elements,
	}, nil
}

//String implements the DocElement#String() method
func (s *Section) String(indentLevel int) string {
	result := bytes.NewBuffer(make([]byte, 0))
	result.WriteString(fmt.Sprintf("%s<Section %d> '%s'\n", indent(indentLevel), s.Heading.Level, s.Heading.Content.String(0)))
	for _, element := range s.Elements {
		result.WriteString(fmt.Sprintf("%s", element.String(indentLevel+1)))
	}
	return result.String()
}

//Accept implements DocElement#Accept(Visitor)
func (s *Section) Accept(v Visitor) error {
	err := v.BeforeVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting section")
	}
	err = v.Visit(s)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section")
	}
	err = s.Heading.Accept(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting section heading")
	}
	for _, element := range s.Elements {
		err := element.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting section element")
		}

	}
	err = v.AfterVisit(s)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting section")
	}
	return nil
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
func NewHeading(level int, inlineContent *InlineContent, metadata []interface{}) (*Heading, error) {
	// counting the lenght of the 'level' value (ie, the number of `=` chars)
	id, _, _ := newMetaElements(metadata)
	// make a default id from the heading's inline content
	if id == nil {
		replacement, err := ReplaceNonAlphanumerics(inlineContent, "_")
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate default ID while instanciating a new Heading element")
		}
		id, _ = NewElementID(*replacement)
	}
	heading := Heading{Level: level, Content: inlineContent, ID: id}
	log.Debugf("Initialized a new Heading: %s", heading.String(0))
	// GlobalStore
	return &heading, nil
}

//String implements the DocElement#String() method
func (h *Heading) String(indentLevel int) string {
	return fmt.Sprintf("%s<Heading %d> %s", indent(indentLevel), h.Level, h.Content.String(0))
}

//Accept implements DocElement#Accept(Visitor)
func (h *Heading) Accept(v Visitor) error {
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
// Lists
// ------------------------------------------

// List the structure for the lists
type List struct {
	ID    *ElementID
	Items []*ListItem
}

//NewList initializes a new `ListItem` from the given content
func NewList(elements []interface{}, metadata []interface{}) (*List, error) {
	id, _, _ := newMetaElements(metadata)
	items := make([]*ListItem, 0)
	log.Debugf("Initializing a new List from %d elements", len(elements))
	currentLevel := 1
	lastItems := make([]*ListItem, 10)
	for _, element := range elements {
		// each "list item" can be a "list item" element followed by an optional blank line (ignored during the processing below)
		//  also, a list item may need to be divided when it contains lines starting with a caret or a group of stars...

		if itemElements, ok := element.([]interface{}); ok {
			if item, ok := itemElements[0].(*ListItem); ok {
				//log.Debugf("  processing element of type '%v' with current level=%d...", reflect.TypeOf(itemElements[0]), item.Level)
				if item.Level == 1 {
					items = append(items, item)
				} else if item.Level > currentLevel {
					// force the current item level to (last seen level + 1)
					item.Level = currentLevel + 1
				}

				if item.Level > 1 {
					// now join the item to its parent
					parentItem := lastItems[item.Level-2]
					if parentItem.Children == nil {
						parentItem.Children = &List{}
					}
					parentItem.Children.Items = append(parentItem.Children.Items, item)
				}
				// memorizes the current item for further processing
				if item.Level > cap(lastItems) { // increase capacity
					newCap := 2 * item.Level
					newSlice := make([]*ListItem, newCap)
					copy(newSlice, lastItems)
					lastItems = newSlice
				}
				if item.Level < currentLevel { // remove some items
					for i := item.Level; i < currentLevel; i++ {
						lastItems[i] = nil

					}
				}
				currentLevel = item.Level
				lastItems[item.Level-1] = item
			}
		}
	}
	log.Debugf("Initialized a new List with %d item(s)", len(items))
	return &List{
		ID:    id,
		Items: items,
	}, nil
}

//String implements the DocElement#String() method
func (l *List) String(indentLevel int) string {
	result := fmt.Sprintf("%s<%v|size=%d>", indent(indentLevel), reflect.TypeOf(l), len(l.Items))
	for _, item := range l.Items {
		result = result + "\n" + item.String(indentLevel+1)
	}
	return result
}

//Accept implements DocElement#Accept(Visitor)
func (l *List) Accept(v Visitor) error {
	err := v.BeforeVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting list")
	}
	err = v.Visit(l)
	if err != nil {
		return errors.Wrapf(err, "error while visiting list")
	}
	for _, item := range l.Items {
		err := item.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting list item")
		}
	}
	err = v.AfterVisit(l)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting list")
	}
	return nil
}

// ListItem the structure for the list items
type ListItem struct {
	Level    int
	Content  *ListItemContent
	Children *List
}

//NewListItem initializes a new `ListItem` from the given content
func NewListItem(level interface{}, content *ListItemContent, children *List) (*ListItem, error) {
	switch vals := reflect.ValueOf(level); vals.Kind() {
	case reflect.Slice:
		log.Debugf("Initializing a new ListItem with content '%s' lines and input level '%d'", content, vals.Len())
		return &ListItem{
			Level:    vals.Len(),
			Content:  content,
			Children: children,
		}, nil
	default:
		return nil, errors.Errorf("Unable to initialize a ListItem with level '%v", level)
	}
}

//String implements the DocElement#String() method
func (i *ListItem) String(indentLevel int) string {
	return i.StringWithIndent(indentLevel)
}

// StringWithIndent same as String() but with a specified number of spaces at the beginning of the line, to produce a given level of indentation
func (i *ListItem) StringWithIndent(indentLevel int) string {
	result := fmt.Sprintf("%s<%v|level=%d> %s", indent(indentLevel), reflect.TypeOf(i), i.Level, i.Content)
	if i.Children != nil {
		for _, c := range i.Children.Items {
			result = result + "\n\t" + c.StringWithIndent(indentLevel+1)
		}
	}
	return result
}

//Accept implements DocElement#Accept(Visitor)
func (i *ListItem) Accept(v Visitor) error {
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
	for _, child := range i.Children.Items {
		err := child.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting list item child")
		}
	}
	err = v.AfterVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting list item")
	}
	return nil
}

// ListItemContent the structure for the list item content
type ListItemContent struct {
	Lines []*InlineContent
}

//NewListItemContent initializes a new `ListItemContent`
func NewListItemContent(text []byte, lines []interface{}) (*ListItemContent, error) {
	log.Debugf("Initializing a new ListItemContent with %d line(s)", len(lines))
	typedLines := make([]*InlineContent, 0)
	for _, line := range lines {
		// here, `line` is an []interface{} in which we need to locate the relevant `*InlineContent` fragment
		if lineFragments, ok := line.([]interface{}); ok {
			for i := range lineFragments {
				if fragment, ok := lineFragments[i].(*InlineContent); ok {
					typedLines = append(typedLines, fragment)
				}
			}
		}
	}
	return &ListItemContent{Lines: typedLines}, nil
}

//String implements the DocElement#String() method
func (c *ListItemContent) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %v", indent(indentLevel), reflect.TypeOf(c), c.Lines)
}

//Accept implements DocElement#Accept(Visitor)
func (c *ListItemContent) Accept(v Visitor) error {
	err := v.BeforeVisit(c)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting ListItemContent")
	}
	err = v.Visit(c)
	if err != nil {
		return errors.Wrapf(err, "error while visiting ListItemContent")
	}
	for _, line := range c.Lines {
		err := line.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting ListItemContent line")
		}

	}
	err = v.AfterVisit(c)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting ListItemContent")
	}
	return nil
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraph
type Paragraph struct {
	Lines []*InlineContent
}

//NewParagraph initializes a new `Paragraph`
func NewParagraph(text []byte, lines []interface{}) (*Paragraph, error) {
	log.Debugf("Initializing a new Paragraph with %d line(s)", len(lines))
	typedLines := make([]*InlineContent, 0)
	for _, line := range lines {
		typedLines = append(typedLines, line.(*InlineContent))
	}
	return &Paragraph{Lines: typedLines}, nil
}

//String implements the DocElement#String() method
func (p *Paragraph) String(indentLevel int) string {
	result := bytes.NewBuffer(make([]byte, 0))
	result.WriteString(fmt.Sprintf("%s<p>", indent(indentLevel)))
	for _, line := range p.Lines {
		result.WriteString(fmt.Sprintf("%s\n", line.String(0)))
	}
	return result.String()
}

//Accept implements DocElement#Accept(Visitor)
func (p *Paragraph) Accept(v Visitor) error {
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
	result := &InlineContent{Elements: mergedElements}
	log.Debugf("Initialized new InlineContent: '%s'", result.String(0))
	return &InlineContent{Elements: mergedElements}, nil
}

//String implements the DocElement#String() method
func (c *InlineContent) String(indentLevel int) string {
	result := bytes.NewBuffer(make([]byte, 0))
	result.WriteString(indent(indentLevel))
	for i, element := range c.Elements {
		result.WriteString(fmt.Sprintf("%s", element.String(0)))
		if i < len(c.Elements)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

//Accept implements DocElement#Accept(Visitor)
func (c *InlineContent) Accept(v Visitor) error {
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
func (i *BlockImage) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %v", indent(indentLevel), reflect.TypeOf(i), i.Macro)
}

func (i *BlockImage) elements() []Visitable {
	return []Visitable{i.ID, i.Link, &i.Macro, i.Title}
}

//Accept implements DocElement#Accept(Visitor)
func (i *BlockImage) Accept(v Visitor) error {
	err := v.BeforeVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting block image")
	}
	err = v.Visit(i)
	if err != nil {
		return errors.Wrapf(err, "error while visiting block image")
	}
	for _, element := range i.elements() {
		err := element.Accept(v)
		if err != nil {
			return errors.Wrapf(err, "error while visiting block image element")
		}

	}
	err = v.AfterVisit(i)
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
	log.Debugf("Initializing a new BlockImageMacro from '%s'", input)
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
	return &BlockImageMacro{
		Path:   path,
		Alt:    alt,
		Width:  width,
		Height: height}, nil
}

//String implements the DocElement#String() method
func (m *BlockImageMacro) String(indentLevel int) string {
	var width, height string
	if m.Width != nil {
		width = *m.Width
	}
	if m.Height != nil {
		height = *m.Height
	}
	return fmt.Sprintf("%s<%v> %s[%s,w=%s h=%s]", indent(indentLevel), reflect.TypeOf(m), m.Path, m.Alt, width, height)
}

//Accept implements DocElement#Accept(Visitor)
func (m *BlockImageMacro) Accept(v Visitor) error {
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
func (b *DelimitedBlock) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %v", indent(indentLevel), reflect.TypeOf(b), b.Content)
}

//Accept implements DocElement#Accept(Visitor)
func (b *DelimitedBlock) Accept(v Visitor) error {
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
	log.Debugf("Initializing a new ElementLink with path=%s", path)
	return &ElementLink{Path: path}, nil
}

//String implements the DocElement#String() method
func (e *ElementLink) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %s", indent(indentLevel), reflect.TypeOf(e), e.Path)
}

//Accept implements DocElement#Accept(Visitor)
func (e *ElementLink) Accept(v Visitor) error {
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
	log.Debugf("Initializing a new ElementID with ID=%s", id)
	return &ElementID{Value: id}, nil
}

//String implements the DocElement#String() method
func (e *ElementID) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %s", indent(indentLevel), reflect.TypeOf(e), e.Value)
}

//Accept implements DocElement#Accept(Visitor)
func (e *ElementID) Accept(v Visitor) error {
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
	log.Debugf("Initializing a new ElementTitle with content=%s", c)
	return &ElementTitle{Content: *c}, nil
}

//String implements the DocElement#String() method
func (e *ElementTitle) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %s", indent(indentLevel), reflect.TypeOf(e), e.Content)
}

//Accept implements DocElement#Accept(Visitor)
func (e *ElementTitle) Accept(v Visitor) error {
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
func (s *StringElement) String(indentLevel int) string {
	return fmt.Sprintf("%s%s", indent(indentLevel), s.Content)
}

//Accept implements DocElement#Accept(Visitor)
func (s *StringElement) Accept(v Visitor) error {
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
func (t *QuotedText) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v (%d)> %v", indent(indentLevel), reflect.TypeOf(t), t.Kind, t.Elements)
}

//Accept implements DocElement#Accept(Visitor)
func (t *QuotedText) Accept(v Visitor) error {
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
	log.Debug("Initializing a new BlankLine")
	return &BlankLine{}, nil
}

//String implements the DocElement#String() method
func (l *BlankLine) String(indentLevel int) string {
	return fmt.Sprintf("%s<Blankline>\n", indent(indentLevel))
}

//Accept implements DocElement#Accept(Visitor)
func (l *BlankLine) Accept(v Visitor) error {
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
func (l *ExternalLink) String(indentLevel int) string {
	return fmt.Sprintf("%s<%v> %s[%s]", indent(indentLevel), reflect.TypeOf(l), l.URL, l.Text)
}

//Accept implements DocElement#Accept(Visitor)
func (l *ExternalLink) Accept(v Visitor) error {
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

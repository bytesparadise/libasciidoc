package types

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"reflect"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// DocElement (and other interfaces)
// ------------------------------------------

// DocElement the interface for all document elements
// TODO: 'String()' remove this method ? no real value here (we could use a visitor to print/debug the elements), by having a `Visit(Visitor)` method instead
type DocElement interface {
	String(int) string
}

// InlineElement the interface for inline elements
type InlineElement interface {
	DocElement
	Visitable
	PlainString() string
}

// Visitable the interface for visitable elements
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

// Document the top-level structure for a document
type Document struct {
	FrontMatter *FrontMatter
	Attributes  *DocumentAttributes
	Elements    []DocElement
}

// NewDocument initializes a new `Document` from the given lines
func NewDocument(frontmatter *FrontMatter, blocks []interface{}) (*Document, error) {
	log.Debugf("Initializing a new Document with %d blocks(s)", len(blocks))
	for i, block := range blocks {
		log.Debugf("Line #%d: %T", i, block)
	}
	elements := convertBlocksToDocElements(blocks)
	document := &Document{Elements: elements}
	document.initAttributes()
	if frontmatter != nil {
		document.Attributes.AddAll(frontmatter.Content)
	}
	return document, nil
}

// initAttributes initializes the Document's attributes
func (d *Document) initAttributes() {
	d.Attributes = &DocumentAttributes{}
	// look-up the document title in the (first) section of level 1
	var headSection *Section
	for _, element := range d.Elements {
		if section, ok := element.(*Section); ok {
			if section.Heading.Level == 1 {
				headSection = section
			}
		}
	}
	if headSection != nil {
		d.Attributes.SetTitle(headSection.Heading.PlainString())
	}

}

// String implements the DocElement#String() method
func (d *Document) String(indentLevel int) string {
	result := bytes.NewBuffer(nil)
	for i := range d.Elements {
		result.WriteString(fmt.Sprintf("\n%s", d.Elements[i].String(0)))
	}
	// log.Debug(fmt.Sprintf("Printing document:\n%s", result.String()))
	return result.String()
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
func NewDocumentAttributeDeclaration(name []interface{}, value []interface{}) (*DocumentAttributeDeclaration, error) {
	attrName, err := Stringify(name)
	if err != nil {
		return nil, errors.Wrapf(err, "error while initializing a DocumentAttributeDeclaration")
	}
	attrValue, err := Stringify(value)
	if err != nil {
		return nil, errors.Wrapf(err, "error while initializing a DocumentAttributeDeclaration")
	}
	log.Debugf("Initialized a new DocumentAttributeDeclaration: '%s' -> '%s'", *attrName, *attrValue)
	return &DocumentAttributeDeclaration{
		Name:  *attrName,
		Value: *attrValue,
	}, nil
}

// String implements the DocElement#String() method
func (a *DocumentAttributeDeclaration) String(indentLevel int) string {
	return fmt.Sprintf("%s<DocumentAttributeDeclaration> '%s' -> '%s'\n", indent(indentLevel), a.Name, a.Value)
}

// DocumentAttributeReset the type for DocumentAttributeReset
type DocumentAttributeReset struct {
	Name string
}

// NewDocumentAttributeReset initializes a new Document Attribute Resets.
func NewDocumentAttributeReset(name []interface{}) (*DocumentAttributeReset, error) {
	attrName, err := Stringify(name)
	if err != nil {
		return nil, errors.Wrapf(err, "error while initializing a DocumentAttributeReset")
	}
	log.Debugf("Initialized a new DocumentAttributeReset: '%s'", *attrName)
	return &DocumentAttributeReset{Name: *attrName}, nil
}

// String implements the DocElement#String() method
func (a *DocumentAttributeReset) String(indentLevel int) string {
	return fmt.Sprintf("%s<DocumentAttributeReset> '%s'\n", indent(indentLevel), a.Name)
}

// PlainString implements the InlineElement#PlainString() method
func (a *DocumentAttributeReset) PlainString() string {
	return fmt.Sprintf("{%s}'\n", a.Name)
}

// DocumentAttributeSubstitution the type for DocumentAttributeSubstitution
type DocumentAttributeSubstitution struct {
	Name string
}

// NewDocumentAttributeSubstitution initializes a new Document Attribute Substitutions
func NewDocumentAttributeSubstitution(name []interface{}) (*DocumentAttributeSubstitution, error) {
	attrName, err := Stringify(name)
	if err != nil {
		return nil, errors.Wrapf(err, "error while initializing a DocumentAttributeSubstitution")
	}
	log.Debugf("Initialized a new DocumentAttributeSubstitution: '%s'", *attrName)
	return &DocumentAttributeSubstitution{Name: *attrName}, nil
}

// String implements the DocElement#String() method
func (a *DocumentAttributeSubstitution) String(indentLevel int) string {
	return fmt.Sprintf("%s<DocumentAttributeSubstitution> '%s'\n", indent(indentLevel), a.Name)
}

// PlainString implements the InlineElement#PlainString() method
func (a *DocumentAttributeSubstitution) PlainString() string {
	return fmt.Sprintf("{%s}'\n", a.Name)
}

// Accept implements Visitable#Accept(Visitor)
func (a *DocumentAttributeSubstitution) Accept(v Visitor) error {
	return v.Visit(a)
}

// ------------------------------------------
// Front Matter
// ------------------------------------------

// FrontMatter the structure for document front-matter
type FrontMatter struct {
	Content map[string]interface{}
}

// NewYamlFrontMatter initializes a new FrontMatter from the given `content`
func NewYamlFrontMatter(content []interface{}) (*FrontMatter, error) {
	c, err := Stringify(content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to 'Stringify' content in front-matter of document")
	}
	attributes := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(*c), &attributes)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse yaml content in front-matter of document")
	}
	log.Debugf("Initialized a new FrontMatter with attributes: %+v", attributes)
	return &FrontMatter{Content: attributes}, nil
}

// ------------------------------------------
// Section
// ------------------------------------------

// Section the structure for a section
type Section struct {
	Heading  Heading
	Elements []DocElement
}

// NewSection initializes a new `Section` from the given heading and elements
func NewSection(heading *Heading, blocks []interface{}) (*Section, error) {
	// log.Debugf("Initializing a new Section with %d block(s)", len(blocks))
	elements := convertBlocksToDocElements(blocks)
	log.Debugf("Initialized a new Section of level %d with %d block(s)", heading.Level, len(blocks))
	return &Section{
		Heading:  *heading,
		Elements: elements,
	}, nil
}

// String implements the DocElement#String() method
func (s *Section) String(indentLevel int) string {
	result := bytes.NewBuffer(nil)
	result.WriteString(fmt.Sprintf("%s<Section %d> '%s'\n", indent(indentLevel), s.Heading.Level, s.Heading.Content.String(0)))
	for _, element := range s.Elements {
		result.WriteString(fmt.Sprintf("%s", element.String(indentLevel+1)))
	}
	return result.String()
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

// NewHeading initializes a new `Heading from the given level and content, with the optional attributes.
// In the attributes, only the ElementID is retained
func NewHeading(level int, inlineContent *InlineContent, attributes []interface{}) (*Heading, error) {
	// counting the lenght of the 'level' value (ie, the number of `=` chars)
	id, _, _ := newElementAttributes(attributes)
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
	return &heading, nil
}

// String implements the DocElement#String() method
func (h *Heading) String(indentLevel int) string {
	return fmt.Sprintf("%s<Heading %d> %s", indent(indentLevel), h.Level, h.Content.String(0))
}

// PlainString returns a plain string version of all elements in this Heading's Content, without any rendering
func (h *Heading) PlainString() string {
	result := bytes.NewBuffer(nil)
	for i, element := range h.Content.Elements {
		result.WriteString(element.PlainString())
		if i < len(h.Content.Elements)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// ------------------------------------------
// Lists
// ------------------------------------------

// List the structure for the lists
type List struct {
	ID    *ElementID
	Items []*ListItem
}

// NewList initializes a new `ListItem` from the given content
func NewList(elements []interface{}, attributes []interface{}) (*List, error) {
	id, _, _ := newElementAttributes(attributes)
	items := make([]*ListItem, 0)
	log.Debugf("Initializing a new List from %d elements", len(elements))
	currentLevel := 1
	lastItems := make([]*ListItem, 10)
	for _, element := range elements {
		// each "list item" can be a "list item" element followed by an optional blank line (ignored during the processing below)
		//  also, a list item may need to be divided when it contains lines starting with a caret or a group of stars...
		if itemElements, ok := element.([]interface{}); ok {
			if item, ok := itemElements[0].(*ListItem); ok {
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

// String implements the DocElement#String() method
func (l *List) String(indentLevel int) string {
	result := fmt.Sprintf("%s<%T|size=%d>", indent(indentLevel), l, len(l.Items))
	for _, item := range l.Items {
		result = result + "\n" + item.String(indentLevel+1)
	}
	return result
}

// ListItem the structure for the list items
type ListItem struct {
	Level    int
	Content  *ListItemContent
	Children *List
}

// NewListItem initializes a new `ListItem` from the given content
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

// String implements the DocElement#String() method
func (i *ListItem) String(indentLevel int) string {
	return i.StringWithIndent(indentLevel)
}

// StringWithIndent same as String() but with a specified number of spaces at the beginning of the line, to produce a given level of indentation
func (i *ListItem) StringWithIndent(indentLevel int) string {
	result := fmt.Sprintf("%s<%T|level=%d> %s", indent(indentLevel), i, i.Level, i.Content)
	if i.Children != nil {
		for _, c := range i.Children.Items {
			result = result + "\n\t" + c.StringWithIndent(indentLevel+1)
		}
	}
	return result
}

// ListItemContent the structure for the list item content
type ListItemContent struct {
	Lines []*InlineContent
}

// NewListItemContent initializes a new `ListItemContent`
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

// String implements the DocElement#String() method
func (c *ListItemContent) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %v", indent(indentLevel), c, c.Lines)
}

// ------------------------------------------
// Paragraph
// ------------------------------------------

// Paragraph the structure for the paragraph
type Paragraph struct {
	Lines []*InlineContent
	ID    *ElementID
	Title *ElementTitle
}

// NewParagraph initializes a new `Paragraph`
func NewParagraph(text []byte, lines []interface{}, attributes []interface{}) (*Paragraph, error) {
	log.Debugf("Initializing a new Paragraph with %d line(s)", len(lines))
	id, title, _ := newElementAttributes(attributes)

	typedLines := make([]*InlineContent, 0)
	for _, line := range lines {
		typedLines = append(typedLines, line.(*InlineContent))
	}
	return &Paragraph{
		Lines: typedLines,
		ID:    id,
		Title: title,
	}, nil
}

// String implements the DocElement#String() method
func (p *Paragraph) String(indentLevel int) string {
	result := bytes.NewBuffer(nil)
	result.WriteString(fmt.Sprintf("%s<p>", indent(indentLevel)))
	for _, line := range p.Lines {
		result.WriteString(fmt.Sprintf("%s\n", line.String(0)))
	}
	return result.String()
}

// ------------------------------------------
// InlineContent
// ------------------------------------------

// InlineContent the structure for the lines in paragraphs
type InlineContent struct {
	// Input    []byte
	Elements []InlineElement
}

// NewInlineContent initializes a new `InlineContent` from the given values
func NewInlineContent(text []byte, elements []interface{}) (*InlineContent, error) {
	mergedElements := merge(elements)
	mergedInlineElements := make([]InlineElement, len(mergedElements))
	for i, element := range mergedElements {
		mergedInlineElements[i] = element.(InlineElement)
	}
	result := &InlineContent{Elements: mergedInlineElements}
	log.Debugf("Initialized new InlineContent with %d element(s): %s", len(result.Elements), result.String(0))
	return result, nil
}

// String implements the DocElement#String() method
func (c *InlineContent) String(indentLevel int) string {
	result := bytes.NewBuffer(nil)
	result.WriteString(indent(indentLevel))
	for i, element := range c.Elements {
		result.WriteString(fmt.Sprintf("%s", element.String(0)))
		if i < len(c.Elements)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// Accept implements Visitable#Accept(Visitor)
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
	Macro ImageMacro
	ID    *ElementID
	Title *ElementTitle
	Link  *ElementLink
}

// NewBlockImage initializes a new `BlockImage`
func NewBlockImage(input []byte, imageMacro ImageMacro, attributes []interface{}) (*BlockImage, error) {
	log.Debugf("Initializing a new BlockImage from %s", input)
	id, title, link := newElementAttributes(attributes)
	return &BlockImage{
		Macro: imageMacro,
		ID:    id,
		Title: title,
		Link:  link,
	}, nil
}

// String implements the DocElement#String() method
func (i *BlockImage) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s", indent(indentLevel), i, i.Macro.String(0))
}

// InlineImage the structure for the inline image macros
type InlineImage struct {
	Macro ImageMacro
}

// NewInlineImage initializes a new `InlineImage` (similar to BlockImage, but without attributes)
func NewInlineImage(input []byte, imageMacro ImageMacro) (*InlineImage, error) {
	log.Debugf("Initializing a new InlineImage from %s", input)
	return &InlineImage{
		Macro: imageMacro,
	}, nil
}

// String implements the DocElement#String() method
func (i *InlineImage) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s", indent(indentLevel), i, i.Macro.String(0))
}

// PlainString implements the InlineElement#PlainString() method
func (i *InlineImage) PlainString() string {
	return i.Macro.Alt
}

// Accept implements Visitable#Accept(Visitor)
func (i *InlineImage) Accept(v Visitor) error {
	err := v.BeforeVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while pre-visiting inline image")
	}
	err = v.Visit(i)
	if err != nil {
		return errors.Wrapf(err, "error while visiting inline image")
	}
	err = i.Macro.Accept(v)
	if err != nil {
		return errors.Wrapf(err, "error while visiting block image element")
	}
	err = v.AfterVisit(i)
	if err != nil {
		return errors.Wrapf(err, "error while post-visiting block image")
	}
	return nil
}

// ImageMacro the structure for the block image macros
type ImageMacro struct {
	Path   string
	Alt    string
	Width  *string
	Height *string
}

// NewImageMacro initializes a new `ImageMacro`
func NewImageMacro(input []byte, path string, attributes interface{}) (*ImageMacro, error) {
	log.Debugf("Initializing a new ImageMacro from %s", input)
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
	return &ImageMacro{
		Path:   path,
		Alt:    alt,
		Width:  width,
		Height: height}, nil
}

// String implements the DocElement#String() method
func (m *ImageMacro) String(indentLevel int) string {
	var width, height string
	if m.Width != nil {
		width = *m.Width
	}
	if m.Height != nil {
		height = *m.Height
	}
	return fmt.Sprintf("%s<%T> %s[%s, w=%s h=%s]", indent(indentLevel), m, m.Path, m.Alt, width, height)
}

// Accept implements Visitable#Accept(Visitor)
func (m *ImageMacro) Accept(v Visitor) error {
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
	// FencedBlock a fenced block
	FencedBlock DelimitedBlockKind = iota
)

// DelimitedBlock the structure for the delimited blocks
type DelimitedBlock struct {
	Kind    DelimitedBlockKind
	Content string
}

// NewDelimitedBlock initializes a new `DelimitedBlock` of the given kind with the given content
func NewDelimitedBlock(kind DelimitedBlockKind, content []interface{}) (*DelimitedBlock, error) {
	c, err := Stringify(content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initialize a new delimited block")
	}
	// remove "\n" or "\r\n", depending on the OS.
	blockContent := strings.TrimSuffix(strings.TrimSuffix(*c, "\n"), "\r")
	log.Debugf("Initialized a new DelimitedBlock with content=`%s`", blockContent)
	return &DelimitedBlock{
		Kind:    kind,
		Content: blockContent,
	}, nil
}

// String implements the DocElement#String() method
func (b *DelimitedBlock) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T|%v> %v", indent(indentLevel), b, b.Kind, b.Content)
}

// Accept implements Visitable#Accept(Visitor)
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
// Literal blocks
// ------------------------------------------

// LiteralBlock the structure for the literal blocks
type LiteralBlock struct {
	Content string
}

// NewLiteralBlock initializes a new `DelimitedBlock` of the given kind with the given content,
// along with the given heading spaces
func NewLiteralBlock(spaces, content []interface{}) (*LiteralBlock, error) {
	// concatenates the spaces with the actual content in a single 'stringified' value
	// log.Debugf("Initializing a new LiteralBlock with spaces='%v' and content=`%v`", spaces, content)
	c, err := Stringify(append(spaces, content...))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initialize a new literal block")
	}
	// remove "\n" or "\r\n", depending on the OS.
	blockContent := strings.TrimSuffix(strings.TrimSuffix(*c, "\n"), "\r")
	log.Debugf("Initialized a new LiteralBlock with content=`%s`", blockContent)
	return &LiteralBlock{
		Content: blockContent,
	}, nil
}

// String implements the DocElement#String() method
func (b *LiteralBlock) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %v", indent(indentLevel), b, b.Content)
}

// Accept implements Visitable#Accept(Visitor)
func (b *LiteralBlock) Accept(v Visitor) error {
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

func newElementAttributes(attributes []interface{}) (*ElementID, *ElementTitle, *ElementLink) {
	var id *ElementID
	var title *ElementTitle
	var link *ElementLink
	for _, item := range attributes {
		switch item := item.(type) {
		case *ElementID:
			id = item
		case *ElementLink:
			link = item
		case *ElementTitle:
			title = item
		default:
			log.Warnf("Unexpected attributes: %T", item)
		}
	}
	return id, title, link
}

// ElementLink the structure for element links
type ElementLink struct {
	Path string
}

// NewElementLink initializes a new `ElementLink` from the given path
func NewElementLink(path string) (*ElementLink, error) {
	log.Debugf("Initializing a new ElementLink with path=%s", path)
	return &ElementLink{Path: path}, nil
}

// String implements the DocElement#String() method
func (e *ElementLink) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s", indent(indentLevel), e, e.Path)
}

// ElementID the structure for element IDs
type ElementID struct {
	Value string
}

// NewElementID initializes a new `ElementID` from the given path
func NewElementID(id string) (*ElementID, error) {
	log.Debugf("Initializing a new ElementID with value=%s", id)
	return &ElementID{Value: id}, nil
}

// String implements the DocElement#String() method
func (e *ElementID) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s", indent(indentLevel), e, e.Value)
}

// ElementTitle the structure for element IDs
type ElementTitle struct {
	Value string
}

// NewElementTitle initializes a new `ElementTitle` from the given content
func NewElementTitle(content []interface{}) (*ElementTitle, error) {
	c, err := Stringify(content)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ElementTitle")
	}
	log.Debugf("Initializing a new ElementTitle with content=%s", *c)
	return &ElementTitle{Value: *c}, nil
}

// String implements the DocElement#String() method
func (e *ElementTitle) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s", indent(indentLevel), e, e.Value)
}

// ------------------------------------------
// StringElement
// ------------------------------------------

// StringElement the structure for strings
type StringElement struct {
	Content string
}

// NewStringElement initializes a new `StringElement` from the given content
func NewStringElement(content interface{}) *StringElement {
	return &StringElement{Content: content.(string)}
}

// String implements the DocElement#String() method
func (s *StringElement) String(indentLevel int) string {
	return fmt.Sprintf("%s%s", indent(indentLevel), s.Content)
}

// PlainString implements the InlineElement#PlainString() method
func (s *StringElement) PlainString() string {
	return s.Content
}

// Accept implements Visitable#Accept(Visitor)
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
	Elements []InlineElement
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
func NewQuotedText(kind QuotedTextKind, content []interface{}) (*QuotedText, error) {
	elements, err := toInlineElements(merge(content))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initialize a new QuotedText")
	}
	log.Debugf("Initializing a new QuotedText with %d elements:", len(elements))
	for _, element := range elements {
		log.Debugf("- %v (%T)", element, element)
	}
	return &QuotedText{Kind: kind, Elements: elements}, nil
}

// String implements the DocElement#String() method
func (t *QuotedText) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T (%d)> %v", indent(indentLevel), t, t.Kind, t.Elements)
}

// PlainString implements the InlineElement#PlainString() method
func (t *QuotedText) PlainString() string {
	result := bytes.NewBuffer(nil)
	for i, element := range t.Elements {
		result.WriteString(element.PlainString())
		if i < len(t.Elements)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// Accept implements Visitable#Accept(Visitor)
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

// NewBlankLine initializes a new `BlankLine`
func NewBlankLine() (*BlankLine, error) {
	log.Debug("Initializing a new BlankLine")
	return &BlankLine{}, nil
}

// String implements the DocElement#String() method
func (l *BlankLine) String(indentLevel int) string {
	return fmt.Sprintf("%s<Blankline>\n", indent(indentLevel))
}

// ------------------------------------------
// Links
// ------------------------------------------

// ExternalLink the structure for the external links
type ExternalLink struct {
	URL  string
	Text string
}

// NewExternalLink initializes a new `ExternalLink`
func NewExternalLink(url, text []interface{}) (*ExternalLink, error) {
	urlStr, err := Stringify(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ExternalLink element")
	}
	textStr, err := Stringify(text)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize a new ExternalLink element")
	}
	// the text includes the surrounding '[' and ']' which should be removed
	trimmedText := strings.TrimPrefix(strings.TrimSuffix(*textStr, "]"), "[")
	return &ExternalLink{URL: *urlStr, Text: trimmedText}, nil
}

// String implements the DocElement#String() method
func (l *ExternalLink) String(indentLevel int) string {
	return fmt.Sprintf("%s<%T> %s[%s]", indent(indentLevel), l, l.URL, l.Text)
}

// PlainString implements the InlineElement#PlainString() method
func (l *ExternalLink) PlainString() string {
	return l.Text
}

// Accept implements Visitable#Accept(Visitor)
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

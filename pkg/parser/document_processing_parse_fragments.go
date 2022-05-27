package parser

import (
	"io"
	"io/ioutil"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func ParseFragments(ctx *ParseContext, source io.Reader, done <-chan interface{}) <-chan types.DocumentFragment {
	resultStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(resultStream)
		b, err := ioutil.ReadAll(source)
		if err != nil {
			resultStream <- types.NewErrorFragment(types.Position{}, err)
			return
		}
		p := newParser(ctx.filename, b, ctx.opts...) // we want to parse block attributes to detect AttributeReferences
		if err := p.setup(g); err != nil {
			resultStream <- types.NewErrorFragment(types.Position{}, err)
			return
		}
		log.WithField("pipeline_task", "document_parsing").Debug("start of document parsing")
	parsing:
		for {
			// if log.IsLevelEnabled(log.DebugLevel) {
			// 	log.Debugf("starting new fragment at line %d", p.pt.line)
			// }
			// line := p.pt.line
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("parsing fragment starting at p.pt.line:%d / p.cur.pos.line:%d", p.pt.line, p.cur.pos.line)
			}
			startOffset := p.pt.offset
			element, err := p.next()
			endOffset := p.pt.offset
			p := types.Position{
				Start: startOffset,
				End:   endOffset,
			}
			if err != nil {
				log.WithError(err).Error("error while parsing")
				resultStream <- types.NewErrorFragment(p, err)
				break parsing
			}
			if element == nil {
				break parsing
			}
			f := types.DocumentFragment{
				Position: p,
			}
			if elements, ok := element.([]interface{}); ok {
				f.Elements = elements
			} else {
				f.Elements = []interface{}{element}
			}
			for _, e := range f.Elements { // TODO: change the grammar rules of these delimited blocks to avoid 2nd parsing
				if err := reparseElement(ctx, e); err != nil {
					log.WithError(err).Errorf("unable to parse content of element of type '%T'", e)
					f.Error = err
					f.Elements = nil
					break
				}
			}
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("parsed fragment:\n%s", spew.Sdump(f))
			}
			select {
			case <-done:
				log.Debug("exiting the document parsing routine")
				break parsing // stops/exits the go routine
			case resultStream <- f:
			}
		}
		log.WithField("pipeline_task", "document_parsing").Debug("end of document parsing")
	}()
	return resultStream
}

func reparseElement(ctx *ParseContext, element interface{}) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("reparsing element of type '%T'", element)
	}
	switch e := element.(type) {
	case *types.ListElements:
		for _, e := range e.Elements {
			if err := reparseElement(ctx, e); err != nil {
				return err
			}
		}
	case *types.ListContinuation:
		if err := reparseElement(ctx, e.Element); err != nil {
			return err
		}
	case *types.Table:
		if err := reparseTable(ctx, e); err != nil {
			return err
		}
	case *types.DelimitedBlock:
		if err := reparseDelimitedBlock(ctx, e); err != nil {
			return err
		}
	}
	return nil
}

func reparseTable(ctx *ParseContext, t *types.Table) error {
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			if err := reparseTableCell(ctx, c); err != nil {
				return err
			}
		}
	}
	if t.Rows != nil {
		for _, r := range t.Rows {
			for _, c := range r.Cells {
				if err := reparseTableCell(ctx, c); err != nil {
					return err
				}
			}
		}
	}
	if t.Footer != nil {
		for _, c := range t.Footer.Cells {
			if err := reparseTableCell(ctx, c); err != nil {
				return err
			}
		}
	}
	return nil
}

func reparseTableCell(ctx *ParseContext, c *types.TableCell) error {
	log.Debugf("reparsing content of table cell")
	switch c.Format {
	case "a":
		opts := append(ctx.opts, Entrypoint("DelimitedBlockElements"))
		elements, err := reparseElements(c.Elements, opts...)
		if err != nil {
			return err
		}
		c.Elements = elements
	default:
		// wrap in a paragraph
		c.Elements = []interface{}{
			&types.Paragraph{
				Elements: c.Elements,
			},
		}
	}

	return nil
}

func reparseDelimitedBlock(ctx *ParseContext, b *types.DelimitedBlock) error {
	switch b.Kind {
	case types.Example, types.Quote, types.Sidebar, types.Open:
		log.Debugf("parsing elements of delimited block of kind '%s'", b.Kind)
		opts := append(ctx.opts, Entrypoint("DelimitedBlockElements"), withinDelimitedBlock(true))
		elements, err := reparseElements(b.Elements, opts...)
		if err != nil {
			return err
		}
		b.Elements = elements
		for _, e := range b.Elements { // TODO: change the grammar rules of these delimited blocks to avoid 2nd parsing
			if err := reparseElement(ctx, e); err != nil {
				return err
			}
		}
	}
	return nil
}

func reparseElements(elements []interface{}, opts ...Option) ([]interface{}, error) {
	content, placeholders := serialize(elements)
	elmts, err := Parse("", content, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse elements") // ignore error (malformed content)
	}
	switch elmts := elmts.(type) {
	case []interface{}:
		// case where last element is `nil` because the parser found a standlone attribute
		for {
			if len(elmts) > 0 && elmts[len(elmts)-1] == nil {
				elmts = elmts[:len(elmts)-1]
			} else {
				break
			}
		}
		return placeholders.restore(elmts), nil
	default:
		return nil, errors.Errorf("unexpected type of result after parsing elements: '%T'", elmts)
	}
}

const documentHeaderKey = "document_header"
const frontMatterKey = "front_matter"

func (c *current) isFrontMatterAllowed() bool {
	allowed, found := c.globalStore[frontMatterKey].(bool)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("checking if FrontMatter is allowed: %t", found && allowed && !c.isWithinDelimitedBlock())
	// }
	return found && allowed && !c.isWithinDelimitedBlock()
}

func (c *current) disableFrontMatterRule() {
	c.globalStore[frontMatterKey] = false
}

func (c *current) isDocumentHeaderAllowed() bool {
	allowed, found := c.globalStore[documentHeaderKey].(bool)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("checking if DocumentHeader is allowed: %t", found && allowed && !c.isWithinDelimitedBlock())
	// }
	return found && allowed && !c.isWithinDelimitedBlock()
}

func (c *current) disableDocumentHeaderRule() {
	c.globalStore[documentHeaderKey] = false
}

const blockAttributesKey = "block_attributes"

func (c *current) storeBlockAttributes(attributes types.Attributes) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("storing block attributes in global store: %s", spew.Sdump(attributes))
	}
	c.globalStore[blockAttributesKey] = attributes
}

func (c *current) isWithinLiteralParagraph() bool {
	if attrs, ok := c.globalStore[blockAttributesKey].(types.Attributes); ok {
		log.Debugf("within literal paragraph: %t", attrs[types.AttrStyle] == types.Literal)
		return attrs[types.AttrPositional1] == types.Literal || attrs[types.AttrStyle] == types.Literal
	}
	log.Debug("not within literal paragraph")
	return false
}

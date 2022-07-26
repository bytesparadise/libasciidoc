package parser

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
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
			start := time.Now()
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
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("parsed fragment:\n%s", spew.Sdump(f))
			}
			log.Debugf("time to parse fragment at %d: %d microseconds", f.Position.Start, time.Since(start).Microseconds())
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

// disables the `DocumentHeader` grammar rule if the element is anything but a BlankLine or a FrontMatter
func (c *current) disableDocumentHeaderRule(element interface{}) {
	switch element.(type) {
	case *types.BlankLine, *types.FrontMatter, *types.AttributeDeclaration:
		return
	default:
		c.globalStore[documentHeaderKey] = false
	}
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
	// log.Debug("not within literal paragraph")
	return false
}

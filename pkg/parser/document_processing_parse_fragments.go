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
		p := newParser(ctx.filename, b, ctx.Opts...) // we want to parse block attributes to detect AttributeReferences
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
				switch e := e.(type) {
				case *types.DelimitedBlock:
					switch e.Kind {
					case types.Example, types.Quote, types.Sidebar:
						if e.Elements, err = parseDelimitedBlockElements(ctx, e); err != nil {
							// log the error, but keep the delimited block empty so we can carry on with the whole processing
							log.WithError(err).Error("unable to parse content of delimited block")
							e.Elements = nil
							break
						}
					}
				}
			}
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("parsed fragment:\n%s", spew.Sdump(f))
			}
			select {
			case <-done:
				log.Info("exiting the document parsing routine")
				break parsing // stops/exits the go routine
			case resultStream <- f:
			}
		}
		log.WithField("pipeline_task", "document_parsing").Debug("end of document parsing")
	}()
	return resultStream
}

func parseDelimitedBlockElements(ctx *ParseContext, b *types.DelimitedBlock) ([]interface{}, error) {
	log.Debugf("parsing content of delimited block of kind '%s'", b.Kind)
	// TODO: use real Substitution?
	content, placeholders := serialize(b.Elements)
	opts := append(ctx.Opts, Entrypoint("DelimitedBlockElements"), withinDelimitedBlock(true))
	elements, err := Parse("", content, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse content") // ignore error (malformed content)
	}
	log.Debugf("elements: '%s'", spew.Sdump(elements))
	switch e := elements.(type) {
	case []interface{}:
		return placeholders.restore(e)
	default:
		return nil, errors.Errorf("unexpected type of result after parsing elements of delimited block: '%T'", e)
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

func (c *current) disableDocumentHeaderRule(element interface{}) {
	// if already disabled, then skip
	if enabled, found := c.globalStore[documentHeaderKey].(bool); found && !enabled {
		return
	}
	// disable based on type of element
	switch element.(type) {
	case *types.AttributeDeclaration, *types.AttributeReset, *types.BlankLine:
		// do not disable yet
		return
	default:
		// if log.IsLevelEnabled(log.DebugLevel) {
		// 	log.Debugf("disabling DocumentHeader after parsing element of type '%T'", element)
		// }
		c.globalStore[documentHeaderKey] = false

	}
}

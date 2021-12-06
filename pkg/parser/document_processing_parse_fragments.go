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
		p := newParser(ctx.filename, b, ctx.Opts...)
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
	opts := append(ctx.Opts, Entrypoint("DelimitedBlockElements"), GlobalStore(delimitedBlockScopeKey, b.Kind))
	elements, err := Parse("", content, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse content") // ignore error (malformed content)
	}
	result, ok := elements.([]interface{})
	if !ok {
		return nil, errors.Errorf("unexpected type of result after parsing elements of delimited block: '%T'", result)
	}
	return placeholders.restore(result)
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

func (c *current) setFrontMatterAllowed(a bool) {
	c.globalStore[frontMatterKey] = a
}

func (c *current) isDocumentHeaderAllowed() bool {
	allowed, found := c.globalStore[documentHeaderKey].(bool)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("checking if DocumentHeader is allowed: %t", found && allowed && !c.isWithinDelimitedBlock())
	// }
	return found && allowed && !c.isWithinDelimitedBlock()
}

func (c *current) setDocumentHeaderAllowed(a bool) {
	c.globalStore[documentHeaderKey] = a
}

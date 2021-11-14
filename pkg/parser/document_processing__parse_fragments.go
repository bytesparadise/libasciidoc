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
	resultStream := make(chan types.DocumentFragment, 1)
	go func() {
		defer close(resultStream)
		b, err := ioutil.ReadAll(source)
		if err != nil {
			resultStream <- types.NewErrorFragment(err)
			return
		}
		p := newParser(ctx.filename, b, ctx.Opts...)
		if err := p.setup(g); err != nil {
			resultStream <- types.NewErrorFragment(err)
			return
		}
		log.WithField("pipeline_task", "document_parsing").Debug("start of document parsing")
	parsing:
		for {
			element, err := p.next()
			if err != nil {
				log.WithError(err).Error("error while parsing")
				resultStream <- types.NewErrorFragment(err)
				break parsing
			}
			if element == nil {
				break parsing
			}
			f := types.DocumentFragment{}
			if elements, ok := element.([]interface{}); ok {
				f.Elements = elements
			} else {
				f.Elements = []interface{}{element}
			}
			// look-up delimited blocks with normal content (will need 2nd pass to parse their content)
			for _, e := range f.Elements {
				if b, ok := e.(*types.DelimitedBlock); ok &&
					(b.Kind == types.Example || b.Kind == types.Quote || b.Kind == types.Sidebar) {
					// if parsing failed, delimited block will be empty
					if b.Elements, err = parseDelimitedBlockElements(ctx, b); err != nil {
						f.Error = err
						break
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
	content, placeholders := serialize(b.Elements) // don't expect placeholders
	opts := append(ctx.Opts, Entrypoint("DelimitedBlockElements"), GlobalStore(delimitedBlockScopeKey, b.Kind))
	f, err := Parse("", content, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse content") // ignore error (malformed content)
	}
	result, ok := f.([]interface{})
	if !ok {
		return nil, errors.Errorf("unexpected type of result after parsing fragments: '%T'", result)
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

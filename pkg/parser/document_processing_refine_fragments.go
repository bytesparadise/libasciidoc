package parser

import (
	"io"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Parses the content of the conplex elements in the incoming fragments
// (for example, some delimited blocks may contain paragraphs, etc.)
func RefineFragments(ctx *ParseContext, source io.Reader, done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) chan types.DocumentFragment {
	resultStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(resultStream)
		for fragment := range fragmentStream {
			select {
			case resultStream <- refineFragment(ctx, fragment):
			case <-done:
				log.WithField("pipeline_task", "refine_fragments").Debug("received 'done' signal")
				return
			}
		}
		log.WithField("pipeline_task", "refine_fragments").Debug("done")
	}()
	return resultStream
}

func refineFragment(ctx *ParseContext, f types.DocumentFragment) types.DocumentFragment {
	if f.Error != nil {
		log.Debugf("skipping fragment refine because of fragment with error: %v", f.Error)
		return f
	}
	start := time.Now()
	for _, e := range f.Elements {
		if err := refineElement(ctx, e); err != nil {
			return types.NewErrorFragment(f.Position, err)
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("refined fragment:\n%s", spew.Sdump(f))
	// }
	log.Debugf("time to refine fragment at %d: %d microseconds", f.Position.Start, time.Since(start).Microseconds())
	return f
}

func refineElement(ctx *ParseContext, element interface{}) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("reparsing element of type '%T'", element)
	}
	switch e := element.(type) {
	case *types.ListElements:
		for _, e := range e.Elements {
			if err := refineElement(ctx, e); err != nil {
				return err
			}
		}
	case *types.ListContinuation:
		if err := refineElement(ctx, e.Element); err != nil {
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
		opts := append(ctx.opts, Entrypoint("DelimitedBlockElements"))
		elements, err := reparseElements(b.Elements, opts...)
		if err != nil {
			return err
		}
		b.Elements = elements
		for _, e := range b.Elements { // TODO: change the grammar rules of these delimited blocks to avoid 2nd parsing
			if err := refineElement(ctx, e); err != nil {
				return err
			}
		}
	}
	return nil
}

func reparseElements(elements []interface{}, opts ...Option) ([]interface{}, error) {
	content, placeholders, err := serialize(elements)
	if err != nil {
		return nil, err
	}
	elmts, err := Parse("", content, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse elements") // ignore error (malformed content)
	}
	switch elmts := elmts.(type) {
	case []interface{}:
		// case where last element is `nil` because the parser found a standlone attribute
		// TODO: still needed?
		for {
			if len(elmts) > 0 && elmts[len(elmts)-1] == nil {
				elmts = elmts[:len(elmts)-1]
			} else {
				break
			}
		}
		return placeholders.restore(elmts)
	default:
		return nil, errors.Errorf("unexpected type of result after parsing elements: '%T'", elmts)
	}
}

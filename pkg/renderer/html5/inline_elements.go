package html5

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// and includes an `\n` character in-between, until the last one.
// Trailing spaces are removed for each line.
func renderLines(ctx *renderer.Context, elements []types.InlineElements) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for i, e := range elements {
		renderedElement, err := renderElement(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render lines")
		}
		if len(renderedElement) > 0 {
			_, err := buff.Write(renderedElement)
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}

		if i < len(elements)-1 && (len(renderedElement) > 0 || ctx.WithinDelimitedBlock()) {
			log.Debugf("rendered line is not the last one in the slice", e)
			_, err := buff.WriteString("\n")
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}
	}
	log.Debugf("rendered line(s): '%s'", buff.String())
	return buff.Bytes(), nil
}

func renderLine(ctx *renderer.Context, elements types.InlineElements, renderElementFunc rendererFunc) ([]byte, error) {
	log.Debugf("rendered %d inline element(s)...", len(elements))
	buff := bytes.NewBuffer(nil)
	for i, element := range elements {
		renderedElement, err := renderElementFunc(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		if i == len(elements)-1 {
			if _, ok := element.(types.StringElement); ok {
				// trim trailing spaces before returning the line
				buff.WriteString(strings.TrimRight(string(renderedElement), " "))
				log.Debugf("trimmed spaces on '%v'", string(renderedElement))
			} else {
				buff.Write(renderedElement)
			}
		} else {
			buff.Write(renderedElement)
		}
	}
	log.Debugf("rendered elements: '%s'", buff.String())
	return buff.Bytes(), nil
}

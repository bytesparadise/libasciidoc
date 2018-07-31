package html5

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type rendererFunc func(*renderer.Context, interface{}) ([]byte, error)

func renderInlineElements(ctx *renderer.Context, e []interface{}, r rendererFunc) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for i, element := range e {
		renderedElement, err := r(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		if i == len(e)-1 {
			if _, ok := element.(types.StringElement); ok {
				// trim trailing spaces before returning the line
				buff.WriteString(strings.TrimRight(string(renderedElement), " "))
				log.Debugf("trimmed spaces on '%v'\n", string(renderedElement))
			} else {
				buff.Write(renderedElement)
			}
		} else {
			buff.Write(renderedElement)
		}
	}
	return buff.Bytes(), nil
}

// renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// and includes an `\n` character in-between, until the last one.
// Trailing spaces are removed for each line.
func renderLines(ctx *renderer.Context, elements []types.InlineElements) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for i, e := range elements {
		renderedElement, err := renderElement(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render element")
		}
		if len(renderedElement) > 0 {
			// if ctx.TrimTrailingSpaces() {
			// 	// trim trailing spaces before returning the line
			// 	buff.WriteString(strings.TrimRight(string(renderedElement), " "))
			// 	log.Debugf("trimmed spaces on '%v'\n", string(renderedElement))
			// } else {
			buff.Write(renderedElement)
			// }
		}

		if i < len(elements)-1 && (len(renderedElement) > 0 || ctx.WithinDelimitedBlock()) {
			log.Debugf("rendered line is not the last one", e)
			buff.WriteString("\n")
		}
	}
	return buff.Bytes(), nil
}

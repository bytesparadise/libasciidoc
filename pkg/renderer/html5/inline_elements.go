package html5

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func renderInlineElements(ctx *renderer.Context, c types.InlineElements) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(nil)
	for _, element := range c {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	return renderedElementsBuff.Bytes(), nil
}

// renderAllInlineElements renders all given InlineElements and includes an `\n` character in-between, until the last one
func renderAllInlineElements(ctx *renderer.Context, elements []types.InlineElements) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for i, e := range elements {
		renderedElement, err := renderElement(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render element")
		}
		if len(renderedElement) > 0 {
			buff.Write(renderedElement)
			if len(renderedElement) > 0 && i < len(elements)-1 {
				log.Debugf("rendered element of type %T is not the last one", e)
				buff.WriteString("\n")
			}
		}
	}
	return buff.Bytes(), nil
}

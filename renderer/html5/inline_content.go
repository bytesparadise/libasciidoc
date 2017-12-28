package html5

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

func renderInlineContent(ctx *renderer.Context, c *types.InlineContent) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(nil)
	for _, element := range c.Elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	return renderedElementsBuff.Bytes(), nil
}

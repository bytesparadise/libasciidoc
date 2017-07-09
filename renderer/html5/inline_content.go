package html5

import (
	"bytes"

	"context"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

func renderInlineContent(ctx context.Context, content types.InlineContent) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for _, element := range content.Elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	return renderedElementsBuff.Bytes(), nil
}

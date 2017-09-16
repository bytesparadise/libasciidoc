package html5

import (
	"bytes"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

func renderInlineContent(ctx asciidoc.Context, content types.InlineContent) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(nil)
	for _, element := range content.Elements {
		renderedElement, err := processElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	return renderedElementsBuff.Bytes(), nil
}

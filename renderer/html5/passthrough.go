package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

func renderPassthrough(ctx *renderer.Context, p types.Passthrough) ([]byte, error) {
	renderedContent, err := renderPassthroughContent(ctx, p)
	if err != nil {
		return nil, errors.Wrap(err, "unable to render passthrough")
	}
	switch p.Kind {
	case types.SinglePlusPassthrough:
		// rendered passthrough content is in an HTML-escaped form
		buf := bytes.NewBuffer(nil)
		template.HTMLEscape(buf, renderedContent)
		return buf.Bytes(), nil
	default:
		return renderedContent, nil
	}
}

// renderPassthroughMacro renders the passthrough content in ist raw from
func renderPassthroughContent(ctx *renderer.Context, p types.Passthrough) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	for _, element := range p.Elements {
		switch element := element.(type) {
		case *types.StringElement:
			// "string" elements must be rendered as-is, ie, without any HTML escaping.
			_, err := buf.WriteString(element.Content)
			if err != nil {
				return nil, err
			}
		default:
			renderedElement, err := renderElement(ctx, element)
			if err != nil {
				return nil, err
			}
			_, err = buf.Write(renderedElement)
			if err != nil {
				return nil, err
			}

		}
	}
	return buf.Bytes(), nil
}

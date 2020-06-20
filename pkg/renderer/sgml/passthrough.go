package sgml

import (
	"html/template"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderInlinePassthrough(ctx *renderer.Context, p types.InlinePassthrough) (string, error) {
	renderedContent, err := r.renderPassthroughContent(ctx, p)
	if err != nil {
		return "", errors.Wrap(err, "unable to render passthrough")
	}
	switch p.Kind {
	case types.SinglePlusPassthrough:
		// rendered passthrough content is in an HTML-escaped form
		buf := &strings.Builder{}
		template.HTMLEscape(buf, []byte(renderedContent))
		return buf.String(), nil
	default:
		return renderedContent, nil
	}
}

// renderPassthroughMacro renders the passthrough content in its raw from
func (r *sgmlRenderer) renderPassthroughContent(ctx *renderer.Context, p types.InlinePassthrough) (string, error) {
	buf := &strings.Builder{}
	for _, element := range p.Elements {
		switch element := element.(type) {
		case types.StringElement:
			// "string" elements must be rendered as-is, ie, without any HTML escaping.
			_, err := buf.WriteString(element.Content)
			if err != nil {
				return "", err
			}
		default:
			renderedElement, err := r.renderElement(ctx, element)
			if err != nil {
				return "", err
			}
			_, err = buf.WriteString(renderedElement)
			if err != nil {
				return "", err
			}

		}
	}
	return buf.String(), nil
}

package sgml

import (
	"html/template"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderPassthroughParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	content, err := r.renderPassthroughContent(ctx, p.Elements)
	if err != nil {
		return "", err
	}
	return content + "\n", nil
}

func (r *sgmlRenderer) renderInlinePassthrough(ctx *renderer.Context, p *types.InlinePassthrough) (string, error) {
	renderedContent, err := r.renderPassthroughContent(ctx, p.Elements)
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
func (r *sgmlRenderer) renderPassthroughContent(ctx *renderer.Context, elements []interface{}) (string, error) {
	result := &strings.Builder{}
	for _, element := range elements {
		switch element := element.(type) {
		case *types.StringElement:
			// "string" elements must be rendered as-is, ie, without any HTML escaping.
			_, err := result.WriteString(element.Content)
			if err != nil {
				return "", err
			}
		default:
			renderedElement, err := r.renderElement(ctx, element)
			if err != nil {
				return "", err
			}
			_, err = result.WriteString(renderedElement)
			if err != nil {
				return "", err
			}

		}
	}
	return strings.Trim(result.String(), "\n"), nil // remove leading and trailing empty lines
}

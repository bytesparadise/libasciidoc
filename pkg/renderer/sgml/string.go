package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var quotes = map[types.QuotedStringKind]struct {
	Open  string
	Close string
	Plain string
}{
	types.SingleQuote: {
		Open:  "&#8216;",
		Close: "&#8217;",
		Plain: "'",
	},
	types.DoubleQuote: {
		Open:  "&#8220;",
		Close: "&#8221;",
		Plain: `"`,
	},
}

func (r *sgmlRenderer) renderQuotedStringPlain(ctx *renderer.Context, s types.QuotedString) (string, error) {
	buf := &strings.Builder{}
	b, err := r.renderPlainText(ctx, s.Elements)
	if err != nil {
		return "", err
	}
	buf.WriteString(quotes[s.Kind].Plain)
	buf.WriteString(b)
	buf.WriteString(quotes[s.Kind].Plain)
	return buf.String(), nil
}

func (r *sgmlRenderer) renderQuotedString(ctx *renderer.Context, s types.QuotedString) (string, error) {
	elements := append([]interface{}{
		types.StringElement{Content: quotes[s.Kind].Open},
	}, s.Elements...)
	elements = append(elements, types.StringElement{Content: quotes[s.Kind].Close})
	return r.renderInlineElements(ctx, elements)
}

func (r *sgmlRenderer) renderStringElement(_ *renderer.Context, str types.StringElement) (string, error) {
	buf := &strings.Builder{}
	err := r.stringElement.Execute(buf, str.Content)
	if err != nil {
		return "", errors.Wrap(err, "unable to render string")
	}

	// NB: For all SGML flavors we are aware of, the numeric entities from
	// Unicode are supported.  We generally avoid named entities.
	result := convert(buf.String(), ellipsis, copyright, trademark, registered)
	return result, nil
}

func ellipsis(source string) string {
	return strings.Replace(source, "...", "&#8230;&#8203;", -1)
}

func copyright(source string) string {
	return strings.Replace(source, "(C)", "&#169;", -1)
}

func trademark(source string) string {
	return strings.Replace(source, "(TM)", "&#8482;", -1)
}

func registered(source string) string {
	return strings.Replace(source, "(R)", "&#174;", -1)
}

type converter func(string) string

func convert(source string, converters ...converter) string {
	result := source
	for _, convert := range converters {
		result = convert(result)
	}
	return result
}

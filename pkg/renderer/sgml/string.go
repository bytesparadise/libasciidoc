package sgml

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var quotes = map[types.QuotedStringKind]struct {
	Open  string
	Close string
}{
	types.SingleQuote: {
		Open:  "\u2018",
		Close: "\u2019",
	},
	types.DoubleQuote: {
		Open:  "\u201c",
		Close: "\u201d",
	},
}

func (r *sgmlRenderer) renderQuotedString(ctx *renderer.Context, s *types.QuotedString) (string, error) {
	elements := append([]interface{}{
		&types.StringElement{
			Content: quotes[s.Kind].Open,
		},
	}, s.Elements...)
	elements = append(elements, &types.StringElement{
		Content: quotes[s.Kind].Close,
	})
	return r.renderInlineElements(ctx, elements)
}

func (r *sgmlRenderer) renderStringElement(ctx *renderer.Context, str *types.StringElement) (string, error) {
	buf := &strings.Builder{}
	err := r.stringElement.Execute(buf, str.Content)
	if err != nil {
		return "", errors.Wrap(err, "unable to render string")
	}

	// NB: For all SGML flavors we are aware of, the numeric entities from
	// Unicode are supported.  We generally avoid named entities.
	result := buf.String()
	if !ctx.UseUnicode() {
		// convert to entities
		result = asciiEntify(result)
	}
	return result, nil
}

func asciiEntify(source string) string {
	out := &strings.Builder{}
	out.Grow(len(source))
	for _, r := range source {
		// This will certain characters that should be escaped alone.  Run them through
		// escape first if that is a concern.
		if r < 128 && (unicode.IsPrint(r) || unicode.IsSpace(r)) {
			out.WriteRune(r)
			continue
		}
		// take care that the entity is unsigned (should always be)
		fmt.Fprintf(out, "&#%d;", uint32(r)) // TODO: avoid `fmt.Fprintf`, use `fmt.Fprint` instead?
	}
	return out.String()
}

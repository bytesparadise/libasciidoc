package sgml

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderStringElement(ctx *context, str *types.StringElement) (string, error) {
	// NB: For all SGML flavors we are aware of, the numeric entities from
	// Unicode are supported.  We generally avoid named entities.
	result := str.Content
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

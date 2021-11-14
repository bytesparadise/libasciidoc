package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderUserMacro(ctx *renderer.Context, m *types.UserMacro) (string, error) {
	buf := &strings.Builder{}
	macro, ok := ctx.Config.Macros[m.Name]
	if !ok {
		if m.Kind == types.BlockMacro {
			// fallback to paragraph
			p, _ := types.NewParagraph(
				&types.StringElement{Content: m.RawText},
			)
			return r.renderParagraph(ctx, p)
		}
		// fallback to render raw text
		return m.RawText, nil
	}
	if err := macro.Execute(buf, m); err != nil {
		return "", err
	}
	return buf.String(), nil

}

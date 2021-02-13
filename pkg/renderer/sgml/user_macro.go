package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderUserMacro(ctx *renderer.Context, um types.UserMacro) (string, error) {
	buf := &strings.Builder{}
	macro, ok := ctx.Config.Macros[um.Name]
	if !ok {
		if um.Kind == types.BlockMacro {
			// fallback to paragraph
			p, _ := types.NewParagraph([]interface{}{
				[]interface{}{
					types.StringElement{Content: um.RawText},
				},
			}, nil)
			return r.renderParagraph(ctx, p)
		}
		// fallback to render raw text
		return um.RawText, nil
	}
	if err := macro.Execute(buf, um); err != nil {
		return "", err
	}
	return buf.String(), nil

}

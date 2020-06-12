package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderUserMacro(ctx *renderer.Context, um types.UserMacro) ([]byte, error) {
	buf := &bytes.Buffer{}
	macro, err := ctx.Config.MacroTemplate(um.Name)
	if err != nil {
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
		_, err = buf.WriteString(um.RawText)
	} else {
		err = macro.Execute(buf, um)
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

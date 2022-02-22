package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderInlineButton(b *types.InlineButton) (string, error) {
	buf := &strings.Builder{}
	if err := r.inlineButton.Execute(buf, b.Attributes[types.AttrButtonLabel]); err != nil {
		return "", err
	}
	return buf.String(), nil
}

package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderInlineMenu(m *types.InlineMenu) (string, error) {
	return r.execute(r.inlineMenu, struct {
		Path []string
	}{
		Path: m.Path,
	})
}

package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderUnorderedList(ctx *renderer.Context, l types.UnorderedList) (string, error) {
	// make sure nested elements are aware of that their rendering occurs within a list
	checkList := false
	if len(l.Items) > 0 {
		if l.Items[0].CheckStyle != types.NoCheck {
			checkList = true
		}
	}
	result := &strings.Builder{}
	// here we must preserve the HTML tags
	err := r.unorderedList.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID        string
			Title     string
			Role      string
			Checklist bool
			Items     []types.UnorderedListItem
		}{
			ID:        r.renderElementID(l.Attributes),
			Title:     r.renderElementTitle(l.Attributes),
			Role:      l.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
			Checklist: checkList,
			Items:     l.Items,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render unordered list")
	}
	return result.String(), nil
}

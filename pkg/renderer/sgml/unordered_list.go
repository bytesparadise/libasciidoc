package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderUnorderedList(ctx *renderer.Context, l types.UnorderedList) ([]byte, error) {
	// make sure nested elements are aware of that their rendering occurs within a list
	checkList := false
	if len(l.Items) > 0 {
		if l.Items[0].CheckStyle != types.NoCheck {
			checkList = true
		}
	}
	result := &bytes.Buffer{}
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
		return nil, errors.Wrapf(err, "unable to render unordered list")
	}
	return result.Bytes(), nil
}

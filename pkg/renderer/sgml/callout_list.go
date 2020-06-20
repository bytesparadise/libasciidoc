package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderCalloutList(ctx *renderer.Context, l types.CalloutList) (string, error) {
	result := &strings.Builder{}
	err := r.calloutList.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Items []types.CalloutListItem
		}{
			ID:    r.renderElementID(l.Attributes),
			Title: l.Attributes.GetAsStringWithDefault(types.AttrTitle, ""),
			Role:  l.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
			Items: l.Items,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list")
	}
	return result.String(), nil
}

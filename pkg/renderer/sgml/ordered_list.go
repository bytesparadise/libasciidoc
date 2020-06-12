package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderOrderedList(ctx *renderer.Context, l types.OrderedList) ([]byte, error) {
	result := &bytes.Buffer{}
	err := r.orderedList.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID             string
			Title          string
			Role           string
			NumberingStyle string
			ListStyle      string
			Start          string
			Items          []types.OrderedListItem
		}{
			ID:             r.renderElementID(l.Attributes),
			Title:          l.Attributes.GetAsStringWithDefault(types.AttrTitle, ""),
			Role:           l.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
			NumberingStyle: getNumberingStyle(l),
			ListStyle:      r.numberingType(getNumberingStyle(l)),
			Start:          l.Attributes.GetAsStringWithDefault(types.AttrStart, ""),
			Items:          l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render ordered list")
	}
	return result.Bytes(), nil
}

func getNumberingStyle(l types.OrderedList) string {
	if s, found := l.Attributes.GetAsString(types.AttrNumberingStyle); found {
		return s
	}
	return string(l.Items[0].NumberingStyle)
}

// TODO: Move this to the HTML output perhaps.
// this numbering style is only really relevant to HTML
func (r *sgmlRenderer) numberingType(style string) string {
	switch style {
	case string(types.LowerAlpha):
		return ` type="a"`
	case string(types.UpperAlpha):
		return ` type="A"`
	case string(types.LowerRoman):
		return ` type="i"`
	case string(types.UpperRoman):
		return ` type="I"`
	default:
		return ""
	}
}

package sgml

import (
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderOrderedList(ctx *renderer.Context, l types.OrderedList) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}

	for _, item := range l.Items {
		if err := r.renderOrderedListItem(ctx, content, item); err != nil {
			return "", errors.Wrap(err, "unable to render unordered list")
		}
	}

	err := r.orderedList.Execute(result, struct {
		Context        *renderer.Context
		ID             sanitized
		Title          sanitized
		Roles          sanitized
		NumberingStyle string
		ListStyle      string
		Start          string
		Content        sanitized
		Items          []types.OrderedListItem
	}{
		ID:             r.renderElementID(l.Attributes),
		Title:          r.renderElementTitle(l.Attributes),
		Roles:          r.renderElementRoles(l.Attributes),
		NumberingStyle: getNumberingStyle(l),
		ListStyle:      r.numberingType(getNumberingStyle(l)),
		Start:          l.Attributes.GetAsStringWithDefault(types.AttrStart, ""),
		Content:        sanitized(content.String()),
		Items:          l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list")
	}
	return result.String(), nil
}

func getNumberingStyle(l types.OrderedList) string {
	if s, found := l.Attributes.GetAsString(types.AttrNumberingStyle); found {
		return s
	}
	return string(l.Items[0].NumberingStyle)
}

// this numbering style is only really relevant to HTML
func (r *sgmlRenderer) numberingType(style string) string {
	switch style {
	case string(types.LowerAlpha):
		return `a`
	case string(types.UpperAlpha):
		return `A`
	case string(types.LowerRoman):
		return `i`
	case string(types.UpperRoman):
		return `I`
	default:
		return ""
	}
}

func (r *sgmlRenderer) renderOrderedListItem(ctx *renderer.Context, w io.Writer, item types.OrderedListItem) error {

	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return errors.Wrap(err, "unable to render unordered list item content")
	}
	return r.orderedListItem.Execute(w, struct {
		Context *renderer.Context
		Content sanitized
	}{
		Context: ctx,
		Content: sanitized(content),
	})
}

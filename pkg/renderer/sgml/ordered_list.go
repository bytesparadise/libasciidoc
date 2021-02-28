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
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list roles")
	}
	style, err := getNumberingStyle(l)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list roles")
	}
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	err = r.orderedList.Execute(result, struct {
		Context   *renderer.Context
		ID        string
		Title     string
		Roles     string
		Style     string
		ListStyle string
		Start     string
		Content   string
		Reversed  bool
		Items     []types.OrderedListItem
	}{
		ID:        r.renderElementID(l.Attributes),
		Title:     title,
		Roles:     roles,
		Style:     style,
		ListStyle: r.numberingType(style),
		Start:     l.Attributes.GetAsStringWithDefault(types.AttrStart, ""),
		Content:   string(content.String()),
		Reversed:  l.Attributes.HasOption("reversed"),
		Items:     l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list")
	}
	return result.String(), nil
}

func getNumberingStyle(l types.OrderedList) (string, error) {
	if s, found, err := l.Attributes.GetAsString(types.AttrStyle); err != nil {
		return "", err
	} else if found {
		return s, nil
	}
	return l.Items[0].Style, nil
}

// this numbering style is only really relevant to HTML
func (r *sgmlRenderer) numberingType(style string) string {
	switch style {
	case types.LowerAlpha:
		return `a`
	case types.UpperAlpha:
		return `A`
	case types.LowerRoman:
		return `i`
	case types.UpperRoman:
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
		Content string
	}{
		Context: ctx,
		Content: string(content),
	})
}

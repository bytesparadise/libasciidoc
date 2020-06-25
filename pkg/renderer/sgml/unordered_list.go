package sgml

import (
	"io"
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
	content := &strings.Builder{}

	for _, item := range l.Items {
		if err := r.renderUnorderedListItem(ctx, content, item); err != nil {
			return "", errors.Wrap(err, "unable to render unordered list")
		}
	}
	// here we must preserve the HTML tags
	err := r.unorderedList.Execute(result, struct {
		Context   *renderer.Context
		ID        sanitized
		Title     sanitized
		Roles     sanitized
		Checklist bool
		Items     []types.UnorderedListItem
		Content   sanitized
	}{
		Context:   ctx,
		ID:        r.renderElementID(l.Attributes),
		Title:     r.renderElementTitle(l.Attributes),
		Checklist: checkList,
		Items:     l.Items,
		Content:   sanitized(content.String()),
		Roles:     r.renderElementRoles(l.Attributes),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render unordered list")
	}
	return result.String(), nil
}
func (r *sgmlRenderer) renderUnorderedListItem(ctx *renderer.Context, w io.Writer, item types.UnorderedListItem) error {

	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return errors.Wrap(err, "unable to render unordered list item content")
	}
	return r.unorderedListItem.Execute(w, struct {
		Context *renderer.Context
		Content sanitized
	}{
		Context: ctx,
		Content: sanitized(content),
	})
}

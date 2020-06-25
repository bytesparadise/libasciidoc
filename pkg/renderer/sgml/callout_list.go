package sgml

import (
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderCalloutList(ctx *renderer.Context, l types.CalloutList) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}

	for _, item := range l.Items {

		err := r.renderCalloutListItem(ctx, content, item)
		if err != nil {
			return "", errors.Wrap(err, "unable to render callout list item")
		}
	}
	err := r.calloutList.Execute(result, struct {
		Context *renderer.Context
		ID      sanitized
		Title   sanitized
		Roles   sanitized
		Content sanitized
		Items   []types.CalloutListItem
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   r.renderElementTitle(l.Attributes),
		Roles:   r.renderElementRoles(l.Attributes),
		Content: sanitized(content.String()),
		Items:   l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderCalloutListItem(ctx *renderer.Context, w io.Writer, item types.CalloutListItem) error {

	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return errors.Wrap(err, "unable to render callout list item content")
	}
	err = r.calloutListItem.Execute(w, struct {
		Context *renderer.Context
		Ref     int
		Content sanitized
	}{
		Context: ctx,
		Ref:     item.Ref,
		Content: sanitized(content),
	})
	if err != nil {
		return errors.Wrap(err, "unable to render callout list")
	}
	return nil
}

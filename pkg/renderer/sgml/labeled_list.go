package sgml

import (
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLabeledList(ctx *renderer.Context, l types.LabeledList) (string, error) {
	tmpl, itemTmpl, err := r.getLabeledListTmpl(l)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}

	content := &strings.Builder{}
	cont := false
	for _, item := range l.Items {
		if cont, err = r.renderLabeledListItem(ctx, itemTmpl, content, cont, item); err != nil {
			return "", errors.Wrap(err, "unable to render unordered list")
		}
	}

	result := &strings.Builder{}
	// here we must preserve the HTML tags
	err = tmpl.Execute(result, struct {
		Context *renderer.Context
		ID      sanitized
		Title   sanitized
		Roles   sanitized
		Content sanitized
		Items   []types.LabeledListItem
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   r.renderElementTitle(l.Attributes),
		Roles:   r.renderElementRoles(l.Attributes),
		Content: sanitized(content.String()),
		Items:   l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}
	// log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) getLabeledListTmpl(l types.LabeledList) (*textTemplate, *textTemplate, error) {
	if layout, ok := l.Attributes["layout"]; ok {
		switch layout {
		case "horizontal":
			return r.labeledListHorizontal, r.labeledListHorizontalItem, nil
		default:
			return nil, nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	if l.Attributes.Has(types.AttrQandA) {
		return r.qAndAList, r.qAndAListItem, nil
	}
	return r.labeledList, r.labeledListItem, nil
}

func (r *sgmlRenderer) renderLabeledListItem(ctx *renderer.Context, tmpl *textTemplate, w io.Writer, continuation bool, item types.LabeledListItem) (bool, error) {

	term, err := r.renderInlineElements(ctx, item.Term)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list term")
	}
	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list content")
	}
	err = tmpl.Execute(w, struct {
		Context      *renderer.Context
		Term         sanitized
		Content      sanitized
		Continuation bool
	}{
		Context:      ctx,
		Term:         sanitized(term),
		Continuation: continuation,
		Content:      sanitized(content),
	})
	return content == "", err
}

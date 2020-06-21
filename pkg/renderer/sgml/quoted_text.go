package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	"strings"
)

// TODO: The bold, italic, and monospace items should be refactored to support semantic tags instead.

func (r *sgmlRenderer) renderQuotedText(ctx *renderer.Context, t types.QuotedText) (string, error) {
	elementsBuffer := &strings.Builder{}
	for _, element := range t.Elements {
		b, err := r.renderElement(ctx, element)
		if err != nil {
			return "", errors.Wrap(err, "unable to render text quote")
		}
		_, err = elementsBuffer.WriteString(b)
		if err != nil {
			return "", errors.Wrapf(err, "unable to render text quote")
		}
	}
	var tmpl *textTemplate
	switch t.Kind {
	case types.Bold:
		tmpl = r.boldText
	case types.Italic:
		tmpl = r.italicText
	case types.Marked:
		tmpl = r.markedText
	case types.Monospace:
		tmpl = r.monospaceText
	case types.Subscript:
		tmpl = r.subscriptText
	case types.Superscript:
		tmpl = r.superscriptText
	default:
		return "", errors.Errorf("unsupported quoted text kind: '%v'", t.Kind)
	}

	result := &strings.Builder{}
	err := tmpl.Execute(result, struct {
		ID         sanitized
		Roles      sanitized
		Attributes types.Attributes
		Content    sanitized
	}{
		Attributes: t.Attributes,
		ID:         r.renderElementID(t.Attributes),
		Roles:      r.renderElementRoles(t.Attributes),
		Content:    sanitized(elementsBuffer.String()),
	}) //nolint: gosec
	if err != nil {
		return "", errors.Wrap(err, "unable to render monospaced quote")
	}
	// log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.String(), nil
}

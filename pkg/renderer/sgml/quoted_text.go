package sgml

import (
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

// TODO: The bold, italic, and monospace items should be refactored to support semantic tags instead.

func (r *sgmlRenderer) renderQuotedText(ctx *context, t *types.QuotedText) (string, error) {
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
	roles, err := r.renderElementRoles(ctx, t.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quoted text roles")
	}
	var tmpl *texttemplate.Template
	switch t.Kind {
	case types.SingleQuoteBold, types.DoubleQuoteBold:
		tmpl, err = r.boldText()
	case types.SingleQuoteItalic, types.DoubleQuoteItalic:
		tmpl, err = r.italicText()
	case types.SingleQuoteMarked, types.DoubleQuoteMarked:
		tmpl, err = r.markedText()
	case types.SingleQuoteMonospace, types.DoubleQuoteMonospace:
		tmpl, err = r.monospaceText()
	case types.SingleQuoteSubscript:
		tmpl, err = r.subscriptText()
	case types.SingleQuoteSuperscript:
		tmpl, err = r.superscriptText()
	default:
		return "", errors.Errorf("unsupported quoted text kind: '%v'", t.Kind)
	}
	if err != nil {
		return "", errors.Wrap(err, "unable to load quoted text template")
	}
	result := &strings.Builder{}
	if err := tmpl.Execute(result, struct {
		ID         string
		Roles      string
		Attributes types.Attributes
		Content    string
	}{
		Attributes: t.Attributes,
		ID:         r.renderElementID(t.Attributes),
		Roles:      roles,
		Content:    string(elementsBuffer.String()),
	}); err != nil {
		return "", errors.Wrap(err, "unable to render quoted text")
	}
	return result.String(), nil
}

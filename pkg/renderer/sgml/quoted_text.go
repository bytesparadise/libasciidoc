package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

// TODO: The bold, italic, and monospace items should be refactored to support semantic tags instead.

func (r *sgmlRenderer) renderQuotedText(ctx *renderer.Context, t *types.QuotedText) (string, error) {
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
	case types.SingleQuoteBold, types.DoubleQuoteBold:
		tmpl = r.boldText
	case types.SingleQuoteItalic, types.DoubleQuoteItalic:
		tmpl = r.italicText
	case types.SingleQuoteMarked, types.DoubleQuoteMarked:
		tmpl = r.markedText
	case types.SingleQuoteMonospace, types.DoubleQuoteMonospace:
		tmpl = r.monospaceText
	case types.SingleQuoteSubscript:
		tmpl = r.subscriptText
	case types.SingleQuoteSuperscript:
		tmpl = r.superscriptText
	default:
		return "", errors.Errorf("unsupported quoted text kind: '%v'", t.Kind)
	}
	roles, err := r.renderElementRoles(ctx, t.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quoted text roles")
	}
	result := &strings.Builder{}
	err = tmpl.Execute(result, struct {
		ID         string
		Roles      string
		Attributes types.Attributes
		Content    string
	}{
		Attributes: t.Attributes,
		ID:         r.renderElementID(t.Attributes),
		Roles:      roles,
		Content:    string(elementsBuffer.String()),
	}) // nolint:gosec
	if err != nil {
		return "", errors.Wrap(err, "unable to render monospaced quote")
	}
	// log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.String(), nil
}

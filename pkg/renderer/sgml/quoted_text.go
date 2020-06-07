package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

// TODO: The bold, italic, and monospace items should be refactored to support semantic tags instead.

func (r *sgmlRenderer) renderQuotedText(ctx *renderer.Context, t types.QuotedText) ([]byte, error) {
	elementsBuffer := &bytes.Buffer{}
	for _, element := range t.Elements {
		b, err := r.renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
		_, err = elementsBuffer.Write(b)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
	}
	result := &bytes.Buffer{}
	var tmpl *textTemplate
	switch t.Kind {
	case types.Bold:
		tmpl = r.boldText
	case types.Italic:
		tmpl = r.italicText
	case types.Monospace:
		tmpl = r.monospaceText
	case types.Subscript:
		tmpl = r.subscriptText
	case types.Superscript:
		tmpl = r.superscriptText
	default:
		return nil, errors.Errorf("unsupported quoted text kind: '%v'", t.Kind)
	}
	err := tmpl.Execute(result, sanitized(elementsBuffer.String())) //nolint: gosec
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render monospaced quote")
	}
	// log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.Bytes(), nil
}

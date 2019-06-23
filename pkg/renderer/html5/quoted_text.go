package html5

import (
	"bytes"
	"html/template"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var boldTextTmpl texttemplate.Template
var italicTextTmpl texttemplate.Template
var monospaceTextTmpl texttemplate.Template
var subscriptTextTmpl texttemplate.Template
var superscriptTextTmpl texttemplate.Template

// initializes the templates
func init() {
	boldTextTmpl = newTextTemplate("bold text", "<strong>{{ . }}</strong>")
	italicTextTmpl = newTextTemplate("italic text", "<em>{{ . }}</em>")
	monospaceTextTmpl = newTextTemplate("monospace text", "<code>{{ . }}</code>")
	subscriptTextTmpl = newTextTemplate("subscript text", "<sub>{{ . }}</sub>")
	superscriptTextTmpl = newTextTemplate("superscript text", "<sup>{{ . }}</sup>")
}

func renderQuotedText(ctx *renderer.Context, t *types.QuotedText) ([]byte, error) {
	elementsBuffer := bytes.NewBuffer(nil)
	for _, element := range t.Elements {
		b, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
		_, err = elementsBuffer.Write(b)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
	}
	result := bytes.NewBuffer(nil)
	var tmpl texttemplate.Template
	switch t.Kind {
	case types.Bold:
		tmpl = boldTextTmpl
	case types.Italic:
		tmpl = italicTextTmpl
	case types.Monospace:
		tmpl = monospaceTextTmpl
	case types.Subscript:
		tmpl = subscriptTextTmpl
	case types.Superscript:
		tmpl = superscriptTextTmpl
	default:
		return nil, errors.Errorf("unsupported quoted text kind: '%v'", t.Kind)
	}
	err := tmpl.Execute(result, template.HTML(elementsBuffer.String())) //nolint: gosec
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render monospaced quote")
	}
	log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.Bytes(), nil
}

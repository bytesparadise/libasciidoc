package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var boldTextTmpl *template.Template
var italicTextTmpl *template.Template
var monospaceTextTmpl *template.Template

// initializes the templates
func init() {
	boldTextTmpl = newHTMLTemplate("bold text", "<strong>{{.}}</strong>")
	italicTextTmpl = newHTMLTemplate("italic text", "<em>{{.}}</em>")
	monospaceTextTmpl = newHTMLTemplate("monospace text", "<code>{{.}}</code>")
}

func renderQuotedText(ctx *renderer.Context, t types.QuotedText) ([]byte, error) {
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
	var tmpl *template.Template
	switch t.Kind {
	case types.Bold:
		tmpl = boldTextTmpl
	case types.Italic:
		tmpl = italicTextTmpl
	case types.Monospace:
		tmpl = monospaceTextTmpl
	default:
		return nil, errors.Errorf("unsupported quoted text kind: %v", t.Kind)
	}
	err := tmpl.Execute(result, template.HTML(elementsBuffer.String()))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render monospaced quote")
	}
	log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.Bytes(), nil
}

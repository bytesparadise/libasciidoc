package html5

import (
	"context"
	"html/template"
	"io"

	"reflect"

	"bytes"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var inlineContentTmpl *template.Template
var stringElementTmpl *template.Template
var boldTextTmpl *template.Template
var italicTextTmpl *template.Template
var monospaceTextTmpl *template.Template

// initializes the templates
func init() {

	inlineContentTmpl = newTemplate("inline content", "<div class=\"paragraph\">\n<p>{{.}}</p>\n</div>")
	stringElementTmpl = newTemplate("string element", "{{.}}")
	boldTextTmpl = newTemplate("bold text", "<strong>{{.}}</strong>")
	italicTextTmpl = newTemplate("italic text", "<em>{{.}}</em>")
	monospaceTextTmpl = newTemplate("monospace text", "<code>{{.}}</code>")
}

func newTemplate(name, src string) *template.Template {
	t, err := template.New(name).Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

// RenderToString renders the givem `document` in HTML and returns the result as a `string`
func RenderToString(ctx context.Context, document types.Document) (*string, error) {
	buff := bytes.NewBuffer(make([]byte, 0))
	err := RenderToWriter(ctx, document, buff)
	if err != nil {
		return nil, err
	}
	result := string(buff.Bytes())
	return &result, nil
}

// RenderToWriter renders the givem `document` in HTML and writes the result in the given `writer`
func RenderToWriter(ctx context.Context, document types.Document, writer io.Writer) error {
	for _, element := range document.Elements {
		renderedElement, err := renderDocElement(element)
		if err != nil {
			return errors.Wrapf(err, "failed to render document")
		}
		_, err = writer.Write(renderedElement)
		if err != nil {
			return errors.Wrapf(err, "failed to render document")
		}
	}
	return nil
}

func renderDocElement(docElement types.DocElement) ([]byte, error) {
	switch docElement.(type) {
	case *types.InlineContent:
		return renderInlineContent(*docElement.(*types.InlineContent))
	case *types.QuotedText:
		return renderQuotedText(*docElement.(*types.QuotedText))
	case *types.StringElement:
		return renderStringElement(*docElement.(*types.StringElement))
	default:
		return nil, errors.Errorf("unsupported element type: %v", reflect.TypeOf(docElement))
	}

}

func renderInlineContent(inlineContent types.InlineContent) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for _, element := range inlineContent.Elements {
		renderedElement, err := renderDocElement(element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render inline content element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	result := bytes.NewBuffer(make([]byte, 0))
	// here we must preserve the HTML tags
	err := inlineContentTmpl.Execute(result, template.HTML(renderedElementsBuff.String()))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline content")
	}
	log.Debugf("rendered inline content: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderQuotedText(t types.QuotedText) ([]byte, error) {
	elementsBuffer := bytes.NewBuffer(make([]byte, 0))
	for _, element := range t.Elements {
		b, err := renderDocElement(element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
		_, err = elementsBuffer.Write(b)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render text quote")
		}
	}
	result := bytes.NewBuffer(make([]byte, 0))
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

func renderStringElement(s types.StringElement) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	err := stringElementTmpl.Execute(result, s.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render string element")
	}
	log.Debugf("rendered string: %s", result.Bytes())
	return result.Bytes(), nil
}

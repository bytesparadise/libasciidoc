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

var stringTemplate *template.Template
var inlineContentTemplate *template.Template
var boldContentTemplate *template.Template

// initializes the templates
func init() {
	var err error
	inlineContentTemplate, err = template.New("inlineContent").Parse("<div class=\"paragraph\">\n<p>{{.}}</p>\n</div>")
	if err != nil {
		log.Fatalf("failed to initialize HTML template: %s", err.Error())
	}
	stringTemplate, err = template.New("string").Parse("{{.}}")
	if err != nil {
		log.Fatalf("failed to initialize HTML template: %s", err.Error())
	}
	boldContentTemplate, err = template.New("boldQuote").Parse("<strong>{{.}}</strong>")
	if err != nil {
		log.Fatalf("failed to initialize HTML template: %s", err.Error())
	}
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
	case *types.BoldQuote:
		return renderBoldQuote(*docElement.(*types.BoldQuote))
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
	err := inlineContentTemplate.Execute(result, template.HTML(renderedElementsBuff.String()))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline content")
	}
	log.Debugf("rendered inline content: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderBoldQuote(q types.BoldQuote) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	err := boldContentTemplate.Execute(result, q.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render bold quote")
	}
	log.Debugf("rendered bold quote: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderStringElement(s types.StringElement) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	err := stringTemplate.Execute(result, s.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render string element")
	}
	log.Debugf("rendered string: %s", result.Bytes())
	return result.Bytes(), nil
}

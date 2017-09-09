package html5

import (
	"bytes"
	"context"
	"html/template"
	"io"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var documentTmpl *texttemplate.Template

func init() {
	documentTmpl = newTextTemplate("root document",
		`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
{{ if .Generator }}<meta name="generator" content="{{.Generator}}">{{ end }}
<title>{{.Title}}</title>
<body class="article">
<div id="header">
<h1>{{.Title}}</h1>
</div>
<div id="content">
{{.Content}}
</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`)
}

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx context.Context, document types.Document, output io.Writer, options renderer.Options) error {
	includeHeaderFooter, err := options.IncludeHeaderFooter()
	if err != nil {
		return errors.Wrap(err, "error while rendering the HTML document")
	}

	lastUpdated, err := options.LastUpdated()
	if err != nil {
		return errors.Wrap(err, "error while rendering the HTML document")
	}

	if *includeHeaderFooter {
		// use a temporary writer for the document's content
		renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
		renderElements(ctx, document.Elements, renderedElementsBuff)
		renderedHTMLElements := template.HTML(renderedElementsBuff.String())

		title := "undefined"
		if document.Attributes.GetTitle() != nil {
			title = *document.Attributes.GetTitle()
		}
		err := documentTmpl.Execute(output, struct {
			Generator   string
			Title       string
			Content     template.HTML
			LastUpdated string
		}{
			Generator:   "libasciidoc", // TODO: externalize this value and include the lib version ?
			Title:       title,
			Content:     renderedHTMLElements,
			LastUpdated: *lastUpdated,
		})
		if err != nil {
			return errors.Wrap(err, "error while rendering the HTML document")
		}
		return nil
	}
	return renderElements(ctx, document.Elements, output)

}

func renderElements(ctx context.Context, elements []types.DocElement, output io.Writer) error {
	hasContent := false
	for _, element := range elements {
		content, err := renderElement(ctx, element)
		if err != nil {
			return errors.Wrapf(err, "failed to render the document")
		}
		// if there's already some content, we need to insert a `\n` before writing
		// the rendering output of the current element (if application, ie, not empty)
		if hasContent && len(content) > 0 {
			log.Debugf("Appending '\\n' after element of type '%T'", element)
			output.Write([]byte("\n"))
		}
		// if the element was rendering into 'something' (ie, not enpty result)
		if len(content) > 0 {
			output.Write(content)
			hasContent = true
		}
	}
	return nil
}

func renderElement(ctx context.Context, element types.DocElement) ([]byte, error) {
	switch element.(type) {
	case *types.Section:
		return renderSection(ctx, *element.(*types.Section))
	case *types.List:
		return renderList(ctx, *element.(*types.List))
	case *types.Paragraph:
		return renderParagraph(ctx, *element.(*types.Paragraph))
	case *types.QuotedText:
		return renderQuotedText(ctx, *element.(*types.QuotedText))
	case *types.BlockImage:
		return renderBlockImage(ctx, *element.(*types.BlockImage))
	case *types.InlineImage:
		return renderInlineImage(ctx, *element.(*types.InlineImage))
	case *types.DelimitedBlock:
		return renderDelimitedBlock(ctx, *element.(*types.DelimitedBlock))
	case *types.InlineContent:
		return renderInlineContent(ctx, *element.(*types.InlineContent))
	case *types.StringElement:
		return renderStringElement(ctx, *element.(*types.StringElement))
	case *types.DocumentAttribute:
		// for now, silently ignored in the output
		return make([]byte, 0), nil
	default:
		return nil, errors.Errorf("unsupported type of element: %T", element)
	}

}

package html5

import (
	"bytes"
	htmltemplate "html/template"
	"io"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var documentTmpl texttemplate.Template

func init() {
	documentTmpl = newTextTemplate("root document",
		`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">{{ if .Generator }}
<meta name="generator" content="{{.Generator}}">{{ end }}
<title>{{.Title}}</title>
<body class="article">
<div id="header">
<h1>{{.Title}}</h1>{{ if .Details }}
{{ .Details }}{{ end }}
</div>
<div id="content">
{{.Content}}
</div>
<div id="footer">
<div id="footer-text">{{ if .RevNumber }}
Version {{.RevNumber}}<br>{{ end }}
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`)

}

func renderDocument(ctx *renderer.Context, output io.Writer) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	renderedTitle, err := renderDocumentTitle(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render full document")
	}

	if ctx.IncludeHeaderFooter() {
		log.Debugf("Rendering full document...")
		// use a temporary writer for the document's content
		renderedElements, err := renderElements(ctx, ctx.Document.Elements)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		documentDetails, err := renderDocumentDetails(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		err = documentTmpl.Execute(output, struct {
			Generator   string
			Title       string
			Content     htmltemplate.HTML
			RevNumber   *string
			LastUpdated string
			Details     *htmltemplate.HTML
		}{
			Generator:   "libasciidoc", // TODO: externalize this value and include the lib version ?
			Title:       string(renderedTitle),
			Content:     htmltemplate.HTML(string(renderedElements)),
			RevNumber:   ctx.Document.Attributes.GetAsString("revnumber"),
			LastUpdated: ctx.LastUpdated(),
			Details:     documentDetails,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		renderedElements, err := renderElements(ctx, ctx.Document.Elements)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		output.Write(renderedElements)
	}
	// copy all document attributes, and override the title with its rendered value instead of the `types.Section` struct
	for k, v := range ctx.Document.Attributes {
		switch k {
		case "doctitle":
			metadata[k] = string(renderedTitle)
		default:
			metadata[k] = v
		}
	}
	return metadata, nil
}

func renderElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(nil)
	hasContent := false
	for _, element := range elements {
		content, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to render the elements")
		}
		// if there's already some content, we need to insert a `\n` before writing
		// the rendering output of the current element (if output is not empty)
		if hasContent && len(content) > 0 {
			log.Debugf("rendered element of type %T (%d)", element, len(content))
			renderedElementsBuff.WriteString("\n")
		}
		// if the element was rendering into 'something' (ie, not enpty result)
		if len(content) > 0 {
			renderedElementsBuff.Write(content)
			hasContent = true
		}
	}
	return renderedElementsBuff.Bytes(), nil
}

// renderDocumentTitle renders the document title
func renderDocumentTitle(ctx *renderer.Context) ([]byte, error) {
	documentTitle, err := ctx.Document.Attributes.GetTitle()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render document title")
	}
	if _, found := documentTitle.Attributes[types.AttrID]; found { // ignore if no ID was set, ie, title is not defined
		title, err := renderPlainString(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document title")
		}
		return title, nil
	}
	return nil, nil
}

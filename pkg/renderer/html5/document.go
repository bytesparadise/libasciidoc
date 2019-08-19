package html5

import (
	"bytes"
	"html"
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
<meta name="generator" content="{{ .Generator }}">{{ end }}
<title>{{ escape .Title }}</title>
</head>
<body class="article">
<div id="header">
<h1>{{ .Header }}</h1>{{ if .Details }}
{{ .Details }}{{ end }}
</div>
<div id="content">
{{ .Content }}
</div>
<div id="footer">
<div id="footer-text">{{ if .RevNumber }}
Version {{ .RevNumber }}<br>{{ end }}
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>`,
		texttemplate.FuncMap{
			"escape": html.EscapeString,
		})

}

// renderDocument renders the whole document, including the HEAD and BODY containers if needed
func renderDocument(ctx *renderer.Context, output io.Writer) (map[string]interface{}, error) {
	renderedTitle, err := renderDocumentTitle(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render full document")
	}
	// log.Debugf("rendered title: '%s'\n", string(renderedTitle))
	renderedHeader, err := renderDocumentHeader(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render full document")
	}
	if ctx.IncludeHeaderFooter() {
		log.Debugf("Rendering full document...")
		// use a temporary writer for the document's content
		renderedElements, err := renderDocumentElements(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		documentDetails, err := renderDocumentDetails(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		revNumber, _ := ctx.Document.Attributes.GetAsString("revnumber")
		err = documentTmpl.Execute(output, struct {
			Generator   string
			Title       string
			Header      string
			Content     htmltemplate.HTML
			RevNumber   string
			LastUpdated string
			Details     *htmltemplate.HTML
		}{
			Generator:   "libasciidoc", // TODO: externalize this value and include the lib version ?
			Title:       string(renderedTitle),
			Header:      string(renderedHeader),
			Content:     htmltemplate.HTML(string(renderedElements)), //nolint: gosec
			RevNumber:   revNumber,
			LastUpdated: ctx.LastUpdated(),
			Details:     documentDetails,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		renderedElements, err := renderDocumentElements(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		_, err = output.Write(renderedElements)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
	}

	// copy all document attributes, and override the title with its rendered value instead of the `types.Section` struct
	metadata := ctx.Document.Attributes
	if len(renderedTitle) > 0 {
		metadata[types.AttrTitle] = string(renderedTitle)
	}
	return metadata, nil
}

// renderDocumentElements renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func renderDocumentElements(ctx *renderer.Context) ([]byte, error) {
	elements := []interface{}{}
	if len(ctx.Document.Elements) > 0 {
		// retrieve elements of the first section 0 (if available), plus remaining elements
		if s, ok := ctx.Document.Elements[0].(types.Section); ok && s.Level == 0 {
			elements = append(elements, s.Elements)
			if len(ctx.Document.Elements) > 1 {
				elements = append(elements, ctx.Document.Elements[1:]...)
			}
		} else {
			elements = ctx.Document.Elements
		}
	}
	// log.Debugf("rendered document with %d element(s)...", len(elements))
	buff := bytes.NewBuffer(nil)
	renderedElements, err := renderElements(ctx, elements)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to render document elements")
	}
	buff.Write(renderedElements)
	renderedFootnotes, err := renderFootnotes(ctx, ctx.Document.Footnotes)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to render document elements")
	}
	buff.Write(renderedFootnotes)

	return buff.Bytes(), nil
}

func renderDocumentTitle(ctx *renderer.Context) ([]byte, error) {
	if documentTitle, hasTitle := ctx.Document.Title(); hasTitle {
		title, err := renderPlainString(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document title")
		}
		return title, nil
	}
	return nil, nil
}

func renderDocumentHeader(ctx *renderer.Context) ([]byte, error) {
	if documentTitle, hasTitle := ctx.Document.Title(); hasTitle {
		title, err := renderElement(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document header")
		}
		return title, nil
	}
	return nil, nil
}

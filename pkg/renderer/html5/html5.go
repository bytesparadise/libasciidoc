package html5

import (
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
<meta name="generator" content="{{ .Generator }}">{{ end }}{{ if .CSS}}
<link type="text/css" rel="stylesheet" href="{{ .CSS }}">{{ end }}
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
			"escape": EscapeString,
		})
}

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx renderer.Context, output io.Writer) (types.Metadata, error) {
	renderedTitle, err := renderDocumentTitle(ctx)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}
	// log.Debugf("rendered title: '%s'\n", string(renderedTitle))
	renderedHeader, err := renderDocumentHeader(ctx)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}
	ctx.TableOfContents, err = NewTableOfContents(ctx)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}
	if ctx.IncludeHeaderFooter() {
		log.Debugf("Rendering full document...")
		// use a temporary writer for the document's content
		renderedElements, err := renderDocumentElements(ctx)
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
		documentDetails, err := renderDocumentDetails(ctx)
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
		revNumber, _ := ctx.Document.Attributes.GetAsString("revnumber")
		err = documentTmpl.Execute(output, struct {
			Generator   string
			Title       string
			Header      string
			Content     htmltemplate.HTML
			RevNumber   string
			LastUpdated string
			CSS         string
			Details     *htmltemplate.HTML
		}{
			Generator:   "libasciidoc", // TODO: externalize this value and include the lib version ?
			Title:       string(renderedTitle),
			Header:      string(renderedHeader),
			Content:     htmltemplate.HTML(string(renderedElements)), //nolint: gosec
			RevNumber:   revNumber,
			LastUpdated: ctx.LastUpdated(),
			CSS:         ctx.CSS(),
			Details:     documentDetails,
		})
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		renderedElements, err := renderDocumentElements(ctx)
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
		_, err = output.Write(renderedElements)
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
	}

	// generate the metadata to be returned to the caller
	metadata := types.Metadata{}
	metadata.Title = string(renderedTitle)
	metadata.LastUpdated = ctx.LastUpdated()
	// include a version of the table of contents
	metadata.TableOfContents = ctx.TableOfContents
	return metadata, err
}

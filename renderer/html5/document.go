package html5

import (
	"bytes"
	"fmt"
	htmltemplate "html/template"
	"io"
	"strconv"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var documentTmpl *texttemplate.Template
var documentDetailsTmpl *htmltemplate.Template
var documentAuthorDetailsTmpl *htmltemplate.Template

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

	documentDetailsTmpl = newHTMLTemplate("document details", `<div class="details">{{ if .Authors }}
{{.Authors}}{{ end }}{{ if .RevNumber }}
<span id="revnumber">version {{.RevNumber}},</span>{{ end }}{{ if .RevDate }}
<span id="revdate">{{.RevDate}}</span>{{ end }}{{ if .RevRemark }}
<br><span id="revremark">{{.RevRemark}}</span>{{ end }}
</div>`)

	documentAuthorDetailsTmpl = newHTMLTemplate("author details", `{{ if .Name }}<span id="author{{.Index}}" class="author">{{.Name}}</span><br>{{ end }}{{ if .Email }}
<span id="email{{.Index}}" class="email"><a href="mailto:{{.Email}}">{{.Email}}</a></span><br>{{ end }}`)
}

func renderDocument(ctx *renderer.Context, output io.Writer) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	documentTitle, err := ctx.Document.Attributes.GetTitle()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render full document")
	}
	renderedTitle, err := renderDocumentTitle(ctx, documentTitle)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render full document")
	}
	if ctx.IncludeHeaderFooter() {
		log.Debugf("Rendering full document...")
		// use a temporary writer for the document's content
		renderedElementsBuff := bytes.NewBuffer(nil)
		renderElements(ctx, ctx.Document.Elements, renderedElementsBuff)
		renderedHTMLElements := htmltemplate.HTML(renderedElementsBuff.String())
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
			Content:     renderedHTMLElements,
			RevNumber:   ctx.Document.Attributes.GetAsString("revnumber"),
			LastUpdated: ctx.LastUpdated(),
			Details:     documentDetails,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		err := renderElements(ctx, ctx.Document.Elements, output)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
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

func renderElements(ctx *renderer.Context, elements []types.DocElement, output io.Writer) error {
	hasContent := false
	for _, element := range elements {
		content, err := renderElement(ctx, element)
		if err != nil {
			return errors.Wrapf(err, "failed to render the document")
		}
		// if there's already some content, we need to insert a `\n` before writing
		// the rendering output of the current element (if application, ie, not empty)
		if hasContent && len(content) > 0 {
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

// renderDocumentTitle renders the document title
func renderDocumentTitle(ctx *renderer.Context, documentTitle *types.SectionTitle) ([]byte, error) {
	if documentTitle != nil {
		title, err := renderPlainString(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document title")
		}
		return title, nil
	}
	return nil, nil
}

func renderDocumentDetails(ctx *renderer.Context) (*htmltemplate.HTML, error) {
	if ctx.Document.Attributes.HasAuthors() {
		authors, err := renderDocumentAuthorsDetails(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error while rendering the document details")
		}
		documentDetailsBuff := bytes.NewBuffer(nil)
		err = documentDetailsTmpl.Execute(documentDetailsBuff, struct {
			Authors   htmltemplate.HTML
			RevNumber *string
			RevDate   *string
			RevRemark *string
		}{
			Authors:   *authors,
			RevNumber: ctx.Document.Attributes.GetAsString("revnumber"),
			RevDate:   ctx.Document.Attributes.GetAsString("revdate"),
			RevRemark: ctx.Document.Attributes.GetAsString("revremark"),
		})
		if err != nil {
			return nil, errors.Wrap(err, "error while rendering the document details")
		}
		documentDetails := htmltemplate.HTML(documentDetailsBuff.String())
		return &documentDetails, nil
	}
	return nil, nil
}

func renderDocumentAuthorsDetails(ctx *renderer.Context) (*htmltemplate.HTML, error) {
	authorsDetailsBuff := bytes.NewBuffer(nil)
	i := 1
	for {
		var authorKey string
		var emailKey string
		var index string
		if i == 1 {
			authorKey = "author"
			emailKey = "email"
			index = ""
		} else {
			authorKey = fmt.Sprintf("author_%d", i)
			emailKey = fmt.Sprintf("email_%d", i)
			index = strconv.Itoa(i)
		}
		// having at least one author is the minimal requirement for document details
		if author := ctx.Document.Attributes.GetAsString(authorKey); author != nil {
			authorDetailsBuff := bytes.NewBuffer(nil)
			err := documentAuthorDetailsTmpl.Execute(authorDetailsBuff, struct {
				Index string
				Name  *string
				Email *string
			}{
				Index: index,
				Name:  author,
				Email: ctx.Document.Attributes.GetAsString(emailKey),
			})
			if err != nil {
				return nil, errors.Wrap(err, "error while rendering the document author")
			}
			// if there were authors before, need to insert a `\n`
			if i > 1 {
				authorsDetailsBuff.WriteString("\n")
			}
			authorsDetailsBuff.Write(authorDetailsBuff.Bytes())
			i++
		} else {
			log.Debugf("No match found for '%s'", authorKey)
			break
		}
	}
	result := htmltemplate.HTML(authorsDetailsBuff.String())
	return &result, nil
}

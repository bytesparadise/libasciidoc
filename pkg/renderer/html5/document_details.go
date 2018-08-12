package html5

import (
	"bytes"
	"fmt"
	htmltemplate "html/template"
	"strconv"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var documentDetailsTmpl texttemplate.Template
var documentAuthorDetailsTmpl texttemplate.Template

func init() {
	documentDetailsTmpl = newTextTemplate("document details", `<div class="details">{{ if .Authors }}
{{.Authors}}{{ end }}{{ if .RevNumber }}
<span id="revnumber">version {{.RevNumber}},</span>{{ end }}{{ if .RevDate }}
<span id="revdate">{{.RevDate}}</span>{{ end }}{{ if .RevRemark }}
<br><span id="revremark">{{.RevRemark}}</span>{{ end }}
</div>`)

	documentAuthorDetailsTmpl = newTextTemplate("author details", `{{ if .Name }}<span id="author{{.Index}}" class="author">{{.Name}}</span><br>{{ end }}{{ if .Email }}
<span id="email{{.Index}}" class="email"><a href="mailto:{{.Email}}">{{.Email}}</a></span><br>{{ end }}`)
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

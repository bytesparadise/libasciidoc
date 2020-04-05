package html5

import (
	"bytes"
	htmltemplate "html/template"
	"io"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var articleTmpl texttemplate.Template
var articleHeaderTmpl texttemplate.Template
var manpageHeaderTmpl texttemplate.Template

func init() {
	articleTmpl = newTextTemplate("article",
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
<body class="{{ .Doctype }}">{{ if .IncludeHeader }}
{{ .Header }}{{ end }}
<div id="content">
{{ .Content }}
</div>{{ if .IncludeFooter }}
<div id="footer">
<div id="footer-text">{{ if .RevNumber }}
Version {{ .RevNumber }}<br>{{ end }}
Last updated {{ .LastUpdated }}
</div>
</div>{{ end }}
</body>
</html>`,
		texttemplate.FuncMap{
			"escape": EscapeString,
		})

	articleHeaderTmpl = newTextTemplate("article header", `<div id="header">
<h1>{{ .Header }}</h1>{{ if .Details }}
{{ .Details }}{{ end }}
</div>`)

	manpageHeaderTmpl = newTextTemplate("manpage header", `{{ if .IncludeH1 }}<div id="header">
<h1>{{ .Header }} Manual Page</h1>
{{ end }}<h2 id="_name">{{ .Name }}</h2>
<div class="sectionbody">
{{ .Content }}
</div>{{ if .IncludeH1 }}
</div>{{ end }}`)
}

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx renderer.Context, output io.Writer) (types.Metadata, error) {
	renderedTitle, err := renderDocumentTitle(ctx)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}
	// needs to be set before rendering the content elements
	ctx.TableOfContents, err = NewTableOfContents(ctx)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}
	renderedHeader, renderedContent, err := splitAndRender(ctx, ctx.Document)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
	}

	if ctx.Config.IncludeHeaderFooter {
		log.Debugf("Rendering full document...")
		revNumber, _ := ctx.Document.Attributes.GetAsString("revnumber")
		err = articleTmpl.Execute(output, struct {
			Generator     string
			Doctype       string
			Title         string
			Header        string
			Content       htmltemplate.HTML
			RevNumber     string
			LastUpdated   string
			CSS           string
			IncludeHeader bool
			IncludeFooter bool
		}{
			Generator:     "libasciidoc", // TODO: externalize this value and include the lib version ?
			Doctype:       ctx.Document.Attributes.GetAsStringWithDefault(types.AttrDocType, "article"),
			Title:         string(renderedTitle),
			Header:        string(renderedHeader),
			Content:       htmltemplate.HTML(string(renderedContent)), //nolint: gosec
			RevNumber:     revNumber,
			LastUpdated:   ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat),
			CSS:           ctx.Config.CSS,
			IncludeHeader: !ctx.Document.Attributes.Has(types.AttrNoHeader),
			IncludeFooter: !ctx.Document.Attributes.Has(types.AttrNoFooter),
		})
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		_, err = output.Write(renderedContent)
		if err != nil {
			return types.Metadata{}, errors.Wrapf(err, "unable to render full document")
		}
	}
	// generate the metadata to be returned to the caller
	metadata := types.Metadata{
		Title:           string(renderedTitle),
		LastUpdated:     ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat),
		TableOfContents: ctx.TableOfContents,
	}
	return metadata, err
}

// splitAndRender the document with the header elements on one side
// and all other elements (table of contents, with preamble, content) on the other side,
// then renders the header and other elements
func splitAndRender(ctx renderer.Context, doc types.Document) ([]byte, []byte, error) {
	switch doc.Attributes.GetAsStringWithDefault(types.AttrDocType, "article") {
	case "manpage":
		return splitAndRenderForManpage(ctx, doc)
	default:
		return splitAndRenderForArticle(ctx, doc)
	}
}

// splits the document with the title of the section 0 (if available) on one side
// and all other elements (table of contents, with preamble, content) on the other side
func splitAndRenderForArticle(ctx renderer.Context, doc types.Document) ([]byte, []byte, error) {
	if ctx.Config.IncludeHeaderFooter {
		if header, exists := doc.Header(); exists {
			renderedHeader, err := renderArticleHeader(ctx, header)
			if err != nil {
				return nil, nil, err
			}
			renderedContent, err := renderDocumentElements(ctx, header.Elements)
			if err != nil {
				return nil, nil, err
			}
			return renderedHeader, renderedContent, nil
		}
	}
	renderedContent, err := renderDocumentElements(ctx, doc.Elements)
	if err != nil {
		return nil, nil, err
	}
	return []byte{}, renderedContent, nil
}

// splits the document with the header elements on one side
// and the other elements (table of contents, with preamble, content) on the other side
func splitAndRenderForManpage(ctx renderer.Context, doc types.Document) ([]byte, []byte, error) {
	header, _ := doc.Header()
	nameSection := header.Elements[0].(types.Section)

	if ctx.Config.IncludeHeaderFooter {
		renderedHeader, err := renderManpageHeader(ctx, header, nameSection)
		if err != nil {
			return nil, nil, err
		}
		renderedContent, err := renderDocumentElements(ctx, header.Elements[1:])
		if err != nil {
			return nil, nil, err
		}
		return renderedHeader, renderedContent, nil
	}
	// in that case, we still want to display the name section
	renderedHeader, err := renderManpageHeader(ctx, types.Section{}, nameSection)
	if err != nil {
		return nil, nil, err
	}
	renderedContent, err := renderDocumentElements(ctx, header.Elements[1:])
	if err != nil {
		return nil, nil, err
	}
	result := bytes.NewBuffer(nil)
	result.Write(renderedHeader)
	result.WriteString("\n")
	result.Write(renderedContent)
	return []byte{}, result.Bytes(), nil
}

func renderDocumentTitle(ctx renderer.Context) ([]byte, error) {
	if header, exists := ctx.Document.Header(); exists {
		title, err := renderPlainText(ctx, header.Title)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document title")
		}
		return title, nil
	}
	return nil, nil
}

func renderArticleHeader(ctx renderer.Context, header types.Section) ([]byte, error) {
	renderedHeader, err := renderInlineElements(ctx, header.Title)
	if err != nil {
		return nil, err
	}
	documentDetails, err := renderDocumentDetails(ctx)
	if err != nil {
		return nil, err
	}

	output := bytes.NewBuffer(nil)
	err = articleHeaderTmpl.Execute(output, struct {
		Header  string
		Details *htmltemplate.HTML
	}{
		Header:  string(renderedHeader),
		Details: documentDetails,
	})
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

func renderManpageHeader(ctx renderer.Context, header types.Section, nameSection types.Section) ([]byte, error) {
	renderedHeader, err := renderInlineElements(ctx, header.Title)
	if err != nil {
		return nil, err
	}
	renderedName, err := renderInlineElements(ctx, nameSection.Title)
	if err != nil {
		return nil, err
	}
	description := nameSection.Elements[0].(types.Paragraph) // TODO: type check
	description.Attributes.AddNonEmpty(types.AttrKind, "manpage")
	renderedContent, err := renderParagraph(ctx, description)
	if err != nil {
		return nil, err
	}
	output := bytes.NewBuffer(nil)
	err = manpageHeaderTmpl.Execute(output, struct {
		Header    string
		Name      string
		Content   htmltemplate.HTML
		IncludeH1 bool
	}{
		Header:    string(renderedHeader),
		Name:      string(renderedName),
		Content:   htmltemplate.HTML(string(renderedContent)), //nolint: gosec
		IncludeH1: len(renderedHeader) > 0,
	})
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

// renderDocumentElements renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func renderDocumentElements(ctx renderer.Context, source []interface{}) ([]byte, error) {
	elements := []interface{}{}
	for i, e := range source {
		switch e := e.(type) {
		case types.Preamble:
			if !e.HasContent() { // why !HasContent ???
				// retain the preamble
				elements = append(elements, e)
				continue
			}
			// retain everything "as-is"
			elements = source
		case types.Section:
			if e.Level == 0 {
				// retain the section's elements...
				elements = append(elements, e.Elements)
				// ... and add the other elements (in case there's another section 0...)
				elements = append(elements, source[i+1:]...)
				continue
			}
			// retain everything "as-is"
			elements = source
		default:
			// retain everything "as-is"
			elements = source
		}
	}
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

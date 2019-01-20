package html5

import (
	"bytes"
	"html"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var paragraphTmpl texttemplate.Template
var admonitionParagraphTmpl texttemplate.Template
var listParagraphTmpl texttemplate.Template
var sourceParagraphTmpl texttemplate.Template
var verseParagraphTmpl texttemplate.Template
var quoteParagraphTmpl texttemplate.Template

// initializes the templates
func init() {
	paragraphTmpl = newTextTemplate("paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines .HardBreak }}{{ if ne $renderedLines "" }}<div {{ if ne .ID "" }}id="{{ .ID }}" {{ end }}class="paragraph">{{ if ne .Title "" }}
<div class="doctitle">{{ escape .Title }}</div>{{ end }}
<p>{{ $renderedLines }}</p>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLinesAsString,
			"escape":      html.EscapeString,
		})

	admonitionParagraphTmpl = newTextTemplate("admonition paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines false }}{{ if ne $renderedLines "" }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
{{ if .IconClass }}<i class="fa icon-{{ .IconClass }}" title="{{ .IconTitle }}"></i>{{ else }}<div class="title">{{ .IconTitle }}</div>{{ end }}
</td>
<td class="content">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
{{ $renderedLines }}
</td>
</tr>
</table>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLinesAsString,
			"escape":      html.EscapeString,
		})

	listParagraphTmpl = newTextTemplate("list paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<p>{{ .CheckStyle }}{{ renderLines $ctx .Lines false }}</p>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLinesAsString,
		})

	sourceParagraphTmpl = newTextTemplate("source paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<div class="listingblock">
<div class="content">
<pre class="highlight">{{ if .Language }}<code class="language-{{ .Language }}" data-lang="{{ .Language }}">{{ else }}<code>{{ end }}{{ renderLines $ctx .Lines | printf "%s" }}</code></pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderPlainString,
			"escape":      html.EscapeString,
		})

	verseParagraphTmpl = newTextTemplate("verse block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ renderElements $ctx .Lines | printf "%s" }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderPlainString,
			"escape":         html.EscapeString,
		})
	quoteParagraphTmpl = newTextTemplate("quote paragraph", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderElements $ctx .Lines false | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderLinesAsString,
			"escape":         html.EscapeString,
		})
}

func renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	if len(p.Lines) == 0 {
		return make([]byte, 0), nil
	}
	result := bytes.NewBuffer(nil)
	var id string
	if i, ok := p.Attributes[types.AttrID].(string); ok {
		id = i
	}
	var err error
	if _, ok := p.Attributes[types.AttrAdmonitionKind]; ok {
		log.Debug("rendering admonition paragraph...")
		k, ok := p.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
		if !ok {
			return nil, errors.Errorf("failed to render admonition with unknown kind: %T", p.Attributes[types.AttrAdmonitionKind])
		}
		err = admonitionParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID        string
				Title     string
				Class     string
				IconTitle string
				IconClass string
				Lines     []types.InlineElements
			}{
				ID:        id,
				Title:     getTitle(p.Attributes),
				Class:     getClass(k),
				IconTitle: getIconTitle(k),
				IconClass: getIconClass(ctx, k),
				Lines:     p.Lines,
			},
		})
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Source {
		log.Debug("rendering source paragraph...")
		err = sourceParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID       string
				Title    string
				Language string
				Lines    []types.InlineElements
			}{
				ID:       id,
				Title:    getTitle(p.Attributes),
				Language: p.Attributes.GetAsString(types.AttrLanguage),
				Lines:    p.Lines,
			},
		})
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Verse {
		log.Debug("rendering verse paragraph...")
		var attribution struct {
			First  string
			Second string
		}
		// TODO; duplicated code from `delimited_block`.
		if author := p.Attributes.GetAsString(types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := p.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := p.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
			attribution.First = title
		}
		err = verseParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID          string
				Title       string
				Attribution struct {
					First  string
					Second string
				}
				Lines []types.InlineElements
			}{
				ID:          id,
				Title:       getTitle(p.Attributes),
				Attribution: attribution,
				Lines:       p.Lines,
			},
		})
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Quote {
		log.Debug("rendering quote paragraph...")
		var attribution struct {
			First  string
			Second string
		}
		// TODO; duplicated code from `delimited_block`.
		if author := p.Attributes.GetAsString(types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := p.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := p.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
			attribution.First = title
		}
		err = quoteParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID          string
				Title       string
				Attribution struct {
					First  string
					Second string
				}
				Lines []types.InlineElements
			}{
				ID:          id,
				Title:       getTitle(p.Attributes),
				Attribution: attribution,
				Lines:       p.Lines,
			},
		})
	} else if ctx.WithinDelimitedBlock() || ctx.WithinList() {
		log.Debugf("rendering paragraph with %d lines within a delimited block or a list", len(p.Lines))

		err = listParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID         string
				Title      string
				CheckStyle string
				Lines      []types.InlineElements
			}{
				ID:         id,
				Title:      getTitle(p.Attributes),
				CheckStyle: renderCheckStyle(p.Attributes[types.AttrCheckStyle]),
				Lines:      p.Lines,
			},
		})
	} else {
		log.Debug("rendering a standalone paragraph")
		err = paragraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID        string
				Title     string
				Lines     []types.InlineElements
				HardBreak bool
			}{
				ID:        id,
				Title:     getTitle(p.Attributes),
				Lines:     p.Lines,
				HardBreak: p.Attributes.Has(types.AttrHardBreaks) || ctx.Document.Attributes.Has(types.DocumentAttrHardBreaks),
			},
		})
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	log.Debugf("rendered paragraph: '%s'", result.String())
	return result.Bytes(), nil
}

func renderCheckStyle(style interface{}) string {
	switch style {
	case types.Unchecked:
		return "&#10063; "
	case types.Checked:
		return "&#10003; "
	default:
		return ""
	}
}

func renderLineBreak() ([]byte, error) {
	return []byte("<br>"), nil
}

func getClass(kind types.AdmonitionKind) string {
	switch kind {
	case types.Tip:
		return "tip"
	case types.Note:
		return "note"
	case types.Important:
		return "important"
	case types.Warning:
		return "warning"
	case types.Caution:
		return "caution"
	default:
		log.Errorf("unexpected kind of admonition: %v", kind)
		return ""
	}
}

func getIconClass(ctx *renderer.Context, kind types.AdmonitionKind) string {
	if ctx.Document.Attributes.GetAsString("icons") == "font" {
		return getClass(kind)
	}
	return ""
}

func getIconTitle(kind types.AdmonitionKind) string {
	switch kind {
	case types.Tip:
		return "Tip"
	case types.Note:
		return "Note"
	case types.Important:
		return "Important"
	case types.Warning:
		return "Warning"
	case types.Caution:
		return "Caution"
	default:
		log.Errorf("unexpected kind of admonition: %v", kind)
		return ""
	}
}

func getTitle(attrs types.ElementAttributes) string {
	if attrs.Has(types.AttrTitle) {
		return strings.TrimSpace(attrs.GetAsString(types.AttrTitle))
	}
	return ""
}

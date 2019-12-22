package html5

import (
	"bytes"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var paragraphTmpl texttemplate.Template
var admonitionParagraphTmpl texttemplate.Template
var delimitedBlockParagraphTmpl texttemplate.Template
var sourceParagraphTmpl texttemplate.Template
var verseParagraphTmpl texttemplate.Template
var quoteParagraphTmpl texttemplate.Template

// initializes the templates
func init() {
	paragraphTmpl = newTextTemplate("paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines .HardBreaks | printf "%s" }}{{ if ne $renderedLines "" }}<div {{ if ne .ID "" }}id="{{ .ID }}" {{ end }}class="paragraph">{{ if ne .Title "" }}
<div class="doctitle">{{ escape .Title }}</div>{{ end }}
<p>{{ $renderedLines }}</p>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLines,
			"escape":      EscapeString,
		})

	admonitionParagraphTmpl = newTextTemplate("admonition paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines | printf "%s" }}{{ if ne $renderedLines "" }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
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
			"renderLines": renderLines,
			"plainText":   PlainText,
			"escape":      EscapeString,
		})

	delimitedBlockParagraphTmpl = newTextTemplate("delimited block paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<p>{{ .CheckStyle }}{{ renderLines $ctx .Lines | printf "%s" }}</p>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLines,
		})

	sourceParagraphTmpl = newTextTemplate("source paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<div class="listingblock">
<div class="content">
<pre class="highlight">{{ if .Language }}<code class="language-{{ .Language }}" data-lang="{{ .Language }}">{{ else }}<code>{{ end }}{{ renderLines $ctx .Lines | printf "%s" }}</code></pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderPlainText,
			"escape":      EscapeString,
		})

	verseParagraphTmpl = newTextTemplate("verse paragraph", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ renderLines $ctx .Lines plainText | printf "%s" }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLines,
			"plainText":   PlainText,
			"escape":      EscapeString,
		})
	quoteParagraphTmpl = newTextTemplate("quote paragraph", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderLines $ctx .Lines | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLines,
			"plainText":   PlainText,
			"escape":      EscapeString,
		})
}

func renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	if len(p.Lines) == 0 {
		return make([]byte, 0), nil
	}
	result := bytes.NewBuffer(nil)
	id := renderElementID(p.Attributes)
	var err error
	if _, ok := p.Attributes[types.AttrAdmonitionKind]; ok {
		return renderAdmonitionParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Source {
		return renderSourceParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Verse {
		return renderVerseParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Quote {
		return renderQuoteParagraph(ctx, p)
	} else if ctx.WithinDelimitedBlock() || ctx.WithinList() {
		return renderDelimitedBlockParagraph(ctx, p)
	} else {
		log.Debug("rendering a standalone paragraph")
		err = paragraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID         string
				Title      string
				Lines      [][]interface{}
				HardBreaks RenderLinesOption
			}{
				ID:         id,
				Title:      renderTitle(p.Attributes),
				Lines:      p.Lines,
				HardBreaks: WithHardBreaks(p.Attributes.Has(types.AttrHardBreaks) || ctx.Document.Attributes.Has(types.DocumentAttrHardBreaks)),
			},
		})
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	// log.Debugf("rendered paragraph: '%s'", result.String())
	return result.Bytes(), nil
}

func renderAdmonitionParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering admonition paragraph...")
	result := bytes.NewBuffer(nil)
	k, ok := p.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
	if !ok {
		return nil, errors.Errorf("failed to render admonition with unknown kind: %T", p.Attributes[types.AttrAdmonitionKind])
	}
	err := admonitionParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID        string
			Title     string
			Class     string
			IconTitle string
			IconClass string
			Lines     [][]interface{}
		}{
			ID:        renderElementID(p.Attributes),
			Title:     renderTitle(p.Attributes),
			Class:     renderClass(k),
			IconTitle: renderIconTitle(k),
			IconClass: renderIconClass(ctx, k),
			Lines:     p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderSourceParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering source paragraph...")
	result := bytes.NewBuffer(nil)
	err := sourceParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Language string
			Lines    [][]interface{}
		}{
			ID:       renderElementID(p.Attributes),
			Title:    renderTitle(p.Attributes),
			Language: p.Attributes.GetAsString(types.AttrLanguage),
			Lines:    p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderVerseParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering verse paragraph...")
	result := bytes.NewBuffer(nil)
	err := verseParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Lines       [][]interface{}
		}{
			ID:          renderElementID(p.Attributes),
			Title:       renderTitle(p.Attributes),
			Attribution: NewParagraphAttribution(p),
			Lines:       p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderQuoteParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering quote paragraph...")
	result := bytes.NewBuffer(nil)
	err := quoteParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Lines       [][]interface{}
		}{
			ID:          renderElementID(p.Attributes),
			Title:       renderTitle(p.Attributes),
			Attribution: NewParagraphAttribution(p),
			Lines:       p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderDelimitedBlockParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debugf("rendering paragraph with %d line(s) within a delimited block or a list", len(p.Lines))
	result := bytes.NewBuffer(nil)
	err := delimitedBlockParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID         string
			Title      string
			CheckStyle string
			Lines      [][]interface{}
		}{
			ID:         renderElementID(p.Attributes),
			Title:      renderTitle(p.Attributes),
			CheckStyle: renderCheckStyle(p.Attributes[types.AttrCheckStyle]),
			Lines:      p.Lines,
		},
	})
	return result.Bytes(), err
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

func renderIconClass(ctx *renderer.Context, kind types.AdmonitionKind) string {
	if icons, _ := ctx.Document.Attributes.GetAsString("icons"); icons == "font" {
		return renderClass(kind)
	}
	return ""
}

func renderClass(kind types.AdmonitionKind) string {
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

func renderIconTitle(kind types.AdmonitionKind) string {
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

func renderTitle(attrs types.ElementAttributes) string {
	if attrs.Has(types.AttrTitle) {
		return strings.TrimSpace(attrs.GetAsString(types.AttrTitle))
	}
	return ""
}

// RenderLinesConfig the config to use when rendering paragraph lines
type RenderLinesConfig struct {
	render     renderFunc
	hardbreaks bool
}

// RenderLinesOption an option to configure the rendering
type RenderLinesOption func(c *RenderLinesConfig)

// WithHardBreaks sets the hard break option
func WithHardBreaks(hardbreaks bool) RenderLinesOption {
	return func(c *RenderLinesConfig) {
		c.hardbreaks = hardbreaks
	}
}

// PlainText sets the render func to PlainText instead of HTML
func PlainText() RenderLinesOption {
	return func(c *RenderLinesConfig) {
		c.render = renderPlainText
	}
}

// renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// and includes an `\n` character in-between, until the last one.
// Trailing spaces are removed for each line.
func renderLines(ctx *renderer.Context, lines [][]interface{}, options ...RenderLinesOption) ([]byte, error) { // renderLineFunc renderFunc, hardbreak bool
	config := RenderLinesConfig{
		render:     renderLine,
		hardbreaks: false,
	}
	for _, apply := range options {
		apply(&config)
	}
	buf := bytes.NewBuffer(nil)
	for i, e := range lines {
		renderedElement, err := config.render(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render lines")
		}
		if len(renderedElement) > 0 {
			_, err := buf.Write(renderedElement)
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}

		if i < len(lines)-1 && (len(renderedElement) > 0 || ctx.WithinDelimitedBlock()) {
			// log.Debugf("rendered line is not the last one in the slice")
			var err error
			if config.hardbreaks {
				_, err = buf.WriteString("<br>\n")
			} else {
				_, err = buf.WriteString("\n")
			}
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}
	}
	// log.Debugf("rendered lines: '%s'", buf.String())
	return buf.Bytes(), nil
}

func renderLine(ctx *renderer.Context, element interface{}) ([]byte, error) {
	if elements, ok := element.([]interface{}); ok {
		return renderInlineElements(ctx, elements)
	}

	return nil, errors.Errorf("invalid type of element for a line: %T", element)
}

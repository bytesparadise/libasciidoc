package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var paragraphTmpl texttemplate.Template
var admonitionParagraphTmpl texttemplate.Template
var admonitionParagraphContentTmpl texttemplate.Template
var listParagraphTmpl texttemplate.Template

// initializes the templates
func init() {
	paragraphTmpl = newTextTemplate("paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedElements := renderElements $ctx .Lines | printf "%s"  }}{{ if ne $renderedElements "" }}<div {{ if ne .ID "" }}id="{{ .ID }}" {{ end }}class="paragraph">{{ if ne .Title "" }}
<div class="doctitle">{{ .Title }}</div>{{ end }}
<p>{{ $renderedElements }}</p>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderAllInlineElements,
			"includeNewline": includeNewline,
		})

	admonitionParagraphTmpl = newTextTemplate("admonition paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedElements := renderElements $ctx .Lines | printf "%s"  }}{{ if ne $renderedElements "" }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
<div class="title">{{ .Icon }}</div>
</td>
<td class="content">{{ if .Title }}
<div class="title">{{ .Title }}</div>{{ end }}
{{ $renderedElements }}
</td>
</tr>
</table>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderAllInlineElements,
			"includeNewline": includeNewline,
		})

	admonitionParagraphContentTmpl = newTextTemplate("admonition paragraph content",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if includeNewline $ctx $index $lines }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})

	listParagraphTmpl = newTextTemplate("list paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<p>{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if includeNewline $ctx $index $lines }}{{ print "\n" }}{{ end }}{{ end }}</p>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
}

func renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	if len(p.Lines) == 0 {
		return make([]byte, 0), nil
	}
	result := bytes.NewBuffer(nil)
	var id, title string
	if i, ok := p.Attributes[types.AttrID].(string); ok {
		id = i
	}
	if t, ok := p.Attributes[types.AttrTitle].(string); ok {
		title = t
	}
	var err error
	if _, ok := p.Attributes[types.AttrAdmonitionKind]; ok {
		k, ok := p.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
		if !ok {
			return nil, errors.Errorf("failed to render admonition with unknown kind: %T", p.Attributes[types.AttrAdmonitionKind])
		}
		err = admonitionParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID    string
				Class string
				Icon  string
				Title string
				Lines []types.InlineElements
			}{
				ID:    id,
				Class: getClass(k),
				Icon:  getIcon(k),
				Title: title,
				Lines: p.Lines,
			},
		})
	} else if ctx.WithinDelimitedBlock() || ctx.WithinList() {
		log.Debug("rendering paragraph within a delimited block or a list")
		err = listParagraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID    string
				Title string
				Lines []types.InlineElements
			}{
				ID:    id,
				Title: title,
				Lines: p.Lines,
			},
		})
	} else {
		log.Debug("rendering a standalone paragraph")
		err = paragraphTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID    string
				Title string
				Lines []types.InlineElements
			}{
				ID:    id,
				Title: title,
				Lines: p.Lines,
			},
		})

	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	return result.Bytes(), nil
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
		log.Error("unexpected kind of admonition: %v", kind)
		return ""
	}
}

func getIcon(kind types.AdmonitionKind) string {
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
		log.Error("unexpected kind of admonition: %v", kind)
		return ""
	}
}

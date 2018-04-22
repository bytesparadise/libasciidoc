package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var admonitionParagraphTmpl texttemplate.Template
var admonitionParagraphContentTmpl texttemplate.Template

// initializes the templates
func init() {
	admonitionParagraphTmpl = newTextTemplate("admonition paragraph", `<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
<div class="title">{{ .Icon }}</div>
</td>
<td class="content">{{ if .Title }}
<div class="title">{{ .Title }}</div>{{ end }}
{{ .Content }}
</td>
</tr>
</table>
</div>`)

	admonitionParagraphContentTmpl = newTextTemplate("admonition paragraph content",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if notLastItem $index $lines }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"notLastItem":   notLastItem,
		})
}

func renderAdmonitionParagraph(ctx *renderer.Context, a types.AdmonitionParagraph) ([]byte, error) {
	log.Debugf("rendering admonition")
	result := bytes.NewBuffer(nil)
	var id, title string
	if i, ok := a.Attributes[types.AttrID].(string); ok {
		id = i
	}
	if t, ok := a.Attributes[types.AttrTitle].(string); ok {
		title = t
	}
	renderedContent, err := renderElement(ctx, a.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to render admonition")
	}
	err = admonitionParagraphTmpl.Execute(result, struct {
		ID      string
		Class   string
		Icon    string
		Title   string
		Content string
	}{
		ID:      id,
		Class:   getClass(a.Kind),
		Icon:    getIcon(a.Kind),
		Title:   title,
		Content: string(renderedContent),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render admonition")
	}
	return result.Bytes(), nil
}

func renderAdmonitionParagraphContent(ctx *renderer.Context, p types.AdmonitionParagraphContent) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := admonitionParagraphContentTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    p,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render admonition paragraph")
	}

	log.Debugf("rendered admonition paragraph: %s", result.Bytes())
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

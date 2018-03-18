package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var admonitionTmpl texttemplate.Template
var admonitionParagraphTmpl texttemplate.Template

// initializes the templates
func init() {
	admonitionTmpl = newTextTemplate("admonition", `<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
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

	admonitionParagraphTmpl = newTextTemplate("admonition paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if notLastItem $index $lines }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"notLastItem":   notLastItem,
		})
}

func renderAdmonition(ctx *renderer.Context, a types.Admonition) ([]byte, error) {
	log.Debugf("rendering admonition")
	result := bytes.NewBuffer(nil)
	var id, title *string
	if a.ID != (types.ElementID{}) {
		id = &a.ID.Value
	}
	if a.Title != (types.ElementTitle{}) {
		title = &a.Title.Value
	}
	renderedContent, err := renderElement(ctx, a.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to render admonition")
	}
	err = admonitionTmpl.Execute(result, struct {
		ID      *string
		Class   string
		Icon    string
		Title   *string
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

func renderAdmonitionParagraph(ctx *renderer.Context, p types.AdmonitionParagraph) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := admonitionParagraphTmpl.Execute(result, ContextualPipeline{
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

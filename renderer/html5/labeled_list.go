package html5

import (
	"bytes"
	"fmt"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// <div class="dlist">
// <dl>
// <dt class="hdlist1">item 1</dt>
// <dd>
// <p>description 1.</p>
// </dd>

// initializes the templates

const defaultLabeledListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="dlist">
<dl>
{{ range .Items }}{{ template "items" wrap $ctx . }}{{ end }}</dl>
</div>{{ end }}`

const defaultLabeledListItemTmpl = `{{ define "items" }}{{ $ctx := .Context }}{{ with .Data }}<dt class="hdlist1">{{ .Term }}</dt>{{ if .Elements }}
<dd>
{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $index $elements }}{{ print "\n" }}{{ end }}{{ end }}
</dd>{{ end }}{{ end }}
{{ end }}`

const horizontalLabeledListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="hdlist">
<table>
{{ range .Items }}<tr>
{{ template "items" wrap $ctx . }}
</tr>
{{ end }}</table>
</div>{{ end }}`

const horizontalLabeledListItemTmpl = `{{ define "items" }}{{ $ctx := .Context }}{{ with .Data }}<td class="hdlist1">
{{ .Term }}
</td>{{ if .Elements }}
<td class="hdlist2">
{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $index $elements }}{{ print "\n" }}{{ end }}{{ end }}
</td>{{ end }}{{ end }}{{ end }}`

func renderLabeledList(ctx *renderer.Context, l *types.LabeledList) ([]byte, error) {
	// TODO: move this to init
	t := texttemplate.New("labeled list")
	t.Funcs(texttemplate.FuncMap{
		"renderElement": renderElement,
		"wrap":          wrap,
		"notLastItem":   notLastItem,
	})
	var err error
	if layout, ok := l.Attributes["layout"]; ok {
		fmt.Printf("Layout: %s\n", layout)
		switch layout {
		case "horizontal":
			t, err = t.Parse(horizontalLabeledListItemTmpl)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to parse labeled list item template")
			}
			t, err = t.Parse(horizontalLabeledListTmpl)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to parse labeled list template")
			}

		default:
			return nil, errors.Wrapf(err, "unsupported labeled list layout: %s", layout)
		}

	} else {
		t, err = t.Parse(defaultLabeledListItemTmpl)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse labeled list item template")
		}
		t, err = t.Parse(defaultLabeledListTmpl)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse labeled list template")
		}
	}

	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err = t.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    l,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}
	log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.Bytes(), nil
}

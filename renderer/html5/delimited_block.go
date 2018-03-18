package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var listingBlockTmpl texttemplate.Template
var exampleBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	listingBlockTmpl = newTextTemplate("listing block", `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>{{ $ctx := .Context }}{{ with .Data }}{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $index $elements }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}</code></pre>
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"notLastItem":   notLastItem,
		})
	exampleBlockTmpl = newTextTemplate("example block", `<div class="exampleblock">
<div class="content">
{{ $ctx := .Context }}{{ with .Data }}{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $index $elements }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"notLastItem":   notLastItem,
		})
}

func renderDelimitedBlock(ctx *renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block")
	result := bytes.NewBuffer(nil)
	tmpl, err := selectDelimitedBlockTemplate(b)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}

	err = tmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    b,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}

func selectDelimitedBlockTemplate(b types.DelimitedBlock) (texttemplate.Template, error) {
	switch b.Kind {
	case types.FencedBlock, types.ListingBlock:
		return listingBlockTmpl, nil
	case types.ExampleBlock:
		return exampleBlockTmpl, nil
	default:
		return texttemplate.Template{}, errors.Errorf("no template for block of kind %v", b.Kind)
	}
}

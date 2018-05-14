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
var admonitionBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	listingBlockTmpl = newTextTemplate("listing block", `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>{{ $ctx := .Context }}{{ with .Data }}{{ .Element }}{{ end }}</code></pre>
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
	exampleBlockTmpl = newTextTemplate("example block", `<div class="exampleblock">
<div class="content">
{{ $ctx := .Context }}{{ with .Data }}{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $index $elements }}{{ print "\n" }}{{ end }}{{ end }}{{ end }}
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
	admonitionBlockTmpl = newTextTemplate("example block", `{{ $ctx := .Context }}{{ with .Data }}<div class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
<div class="title">{{ .Icon }}</div>
</td>
<td class="content">
{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $index $elements }}{{ print "\n" }}{{ end }}{{ end }}
</td>
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
}

func renderDelimitedBlock(ctx *renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block")
	result := bytes.NewBuffer(nil)
	var err error
	elements := discardTrailingBlankLines(b.Elements)
	switch b.Kind {
	case types.FencedBlock, types.ListingBlock:
		content := make([]byte, 0)
		for _, e := range elements {
			s, err := renderPlainString(ctx, e)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to initialize a new delimited block")
			}
			content = append(content, s...)
		}
		ctx.SetIncludeBlankLine(true)
		ctx.SetWithinDelimitedBlock(true)
		defer func() {
			ctx.SetIncludeBlankLine(false)
			ctx.SetIncludeBlankLine(false)
		}()

		err = listingBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				Element string
			}{
				Element: string(content),
			},
		})
	case types.ExampleBlock:
		if k, ok := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind); ok {
			err = admonitionBlockTmpl.Execute(result, ContextualPipeline{
				Context: ctx,
				Data: struct {
					ID       string
					Class    string
					Icon     string
					Title    string
					Elements []interface{}
				}{
					Class:    getClass(k),
					Icon:     getIcon(k),
					Elements: elements,
				},
			})
		} else {
			// default, example block
			err = exampleBlockTmpl.Execute(result, ContextualPipeline{
				Context: ctx,
				Data: struct {
					Elements []interface{}
				}{
					Elements: elements,
				},
			})
		}
	default:
		err = errors.Errorf("no template for block of kind %v", b.Kind)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}

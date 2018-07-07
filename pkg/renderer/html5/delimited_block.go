package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var listingBlockTmpl texttemplate.Template
var exampleBlockTmpl texttemplate.Template
var admonitionBlockTmpl texttemplate.Template
var verseBlockTmpl texttemplate.Template

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
	verseBlockTmpl = newTextTemplate("verse block", `<div class="verseblock">
{{ $ctx := .Context }}{{ with .Data }}<pre class="content">{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $index $elements }}{{ print "\n" }}{{ end }}{{ end }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
	admonitionBlockTmpl = newTextTemplate("admonition block", `{{ $ctx := .Context }}{{ with .Data }}<div class="admonitionblock {{ .Class }}">
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
	kind := b.Attributes[types.AttrBlockKind]
	switch kind {
	case types.Fenced, types.Listing:
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
	case types.Example:
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
	case types.Verse:
		var elements []types.InlineElements
		if len(b.Elements) > 0 {
			if p, ok := b.Elements[0].(types.Paragraph); ok {
				elements = p.Lines
			}
		} else {
			elements = make([]types.InlineElements, 0)
		}
		var attribution struct {
			First  string
			Second string
		}
		if b.Attributes[types.AttrVerseAuthor].(string) != "" {
			attribution.First = b.Attributes[types.AttrVerseAuthor].(string)
			if b.Attributes[types.AttrVerseTitle].(string) != "" {
				attribution.Second = b.Attributes[types.AttrVerseTitle].(string)
			}
		} else if b.Attributes[types.AttrVerseTitle].(string) != "" {
			attribution.First = b.Attributes[types.AttrVerseTitle].(string)
		}
		err = verseBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				Attribution struct {
					First  string
					Second string
				}
				Title    string
				Elements []types.InlineElements
			}{
				Attribution: attribution,
				Elements:    elements,
			},
		})
	case types.Comment:
		// nothing to do
	default:
		err = errors.Errorf("no template for block of kind %v", kind)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}

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

var listingBlockTmpl texttemplate.Template
var exampleBlockTmpl texttemplate.Template
var admonitionBlockTmpl texttemplate.Template
var quoteBlockTmpl texttemplate.Template
var verseBlockTmpl texttemplate.Template
var sidebarBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	listingBlockTmpl = newTextTemplate("listing block", `<div class="listingblock">
<div class="content">
<pre>{{ $ctx := .Context }}{{ with .Data }}{{ .Content }}{{ end }}</pre>
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
		})

	exampleBlockTmpl = newTextTemplate("example block", `<div class="exampleblock">
<div class="content">
{{ $ctx := .Context }}{{ with .Data }}{{ $elements := .Elements }}{{ renderElements $ctx $elements }}{{ end }}
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderElementsAsString,
		})

	quoteBlockTmpl = newTextTemplate("quote block", `<div class="quoteblock">
{{ $ctx := .Context }}{{ with .Data }}<blockquote>
{{ renderElements $ctx .Elements }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderElementsAsString,
		})

	verseBlockTmpl = newTextTemplate("verse block", `<div class="verseblock">
{{ $ctx := .Context }}{{ with .Data }}<pre class="content">{{ renderElements $ctx .Elements }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderInlineElementsAsString,
		})

	admonitionBlockTmpl = newTextTemplate("admonition block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID}}" {{ end}}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
<div class="title">{{ .Icon }}</div>
</td>
<td class="content">
{{ if .Title }}<div class="title">{{ .Title }}</div>
{{ end }}{{ renderElements $ctx .Elements }}
</td>
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElementsAsString,
		})

	sidebarBlockTmpl = newTextTemplate("sidebar block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="sidebarblock">
<div class="content">{{ if .Title }}
<div class="title">{{ .Title }}</div>{{ end }}
{{ renderElements $ctx .Elements }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElementsAsString,
		})
}

func renderDelimitedBlock(ctx *renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block of kind '%v'", b.Attributes[types.AttrKind])
	result := bytes.NewBuffer(nil)
	elements := discardTrailingBlankLines(b.Elements)
	var id, title string
	if i, ok := b.Attributes[types.AttrID].(string); ok {
		id = strings.TrimSpace(i)
	}
	if t, ok := b.Attributes[types.AttrTitle].(string); ok {
		title = strings.TrimSpace(t)
	}
	var err error
	kind := b.Attributes[types.AttrKind]
	switch kind {
	case types.Fenced, types.Listing:
		ctx.SetWithinDelimitedBlock(true)
		ctx.SetIncludeBlankLine(true)
		defer func() {
			ctx.SetWithinDelimitedBlock(false)
			ctx.SetIncludeBlankLine(false)
		}()
		content := make([]byte, 0)
		for _, e := range elements {
			s, err := renderPlainString(ctx, e)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to initialize a new delimited block")
			}
			content = append(content, s...)
		}
		err = listingBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				Content string
			}{
				Content: string(content),
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
					ID:       id,
					Class:    getClass(k),
					Icon:     getIcon(k),
					Title:    title,
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
	case types.Quote:
		var attribution struct {
			First  string
			Second string
		}
		if author := attributeAsString(b.Attributes, types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := attributeAsString(b.Attributes, types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := attributeAsString(b.Attributes, types.AttrQuoteTitle); title != "" {
			attribution.First = title
		}
		err = quoteBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				Attribution struct {
					First  string
					Second string
				}
				Title    string
				Elements []interface{}
			}{
				Attribution: attribution,
				Elements:    b.Elements,
			},
		})
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
		if author := attributeAsString(b.Attributes, types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := attributeAsString(b.Attributes, types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := attributeAsString(b.Attributes, types.AttrQuoteTitle); title != "" {
			attribution.First = title
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
	case types.Sidebar:
		err = sidebarBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID       string
				Title    string
				Elements []interface{}
			}{
				ID:       id,
				Title:    title,
				Elements: elements,
			},
		})
	default:
		err = errors.Errorf("no template for block of kind %v", kind)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}

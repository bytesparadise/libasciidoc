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
<pre>{{ $ctx := .Context }}{{ with .Data }}{{ range $index, $element := .Elements }}{{ renderPlainString $ctx $element | printf "%s" }}{{ end }}{{ end }}</pre>
</div>
</div>`,
		texttemplate.FuncMap{
			"renderPlainString": renderPlainString,
		})

	exampleBlockTmpl = newTextTemplate("example block", `<div class="exampleblock">
<div class="content">
{{ $ctx := .Context }}{{ with .Data }}{{ $elements := .Elements }}{{ renderElements $ctx $elements | printf "%s" }}{{ end }}
</div>
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})

	quoteBlockTmpl = newTextTemplate("quote block", `<div class="quoteblock">
{{ $ctx := .Context }}{{ with .Data }}<blockquote>
{{ renderElements $ctx .Elements | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})

	verseBlockTmpl = newTextTemplate("verse block", `<div class="verseblock">
{{ $ctx := .Context }}{{ with .Data }}<pre class="content">{{ renderElements $ctx .Elements | printf "%s" }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})

	admonitionBlockTmpl = newTextTemplate("admonition block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID}}" {{ end}}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
<div class="title">{{ .Icon }}</div>
</td>
<td class="content">
{{ if .Title }}<div class="title">{{ .Title }}</div>
{{ end }}{{ renderElements $ctx .Elements | printf "%s" }}
</td>
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})

	sidebarBlockTmpl = newTextTemplate("sidebar block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="sidebarblock">
<div class="content">{{ if .Title }}
<div class="title">{{ .Title }}</div>{{ end }}
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
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
		oldWithin := ctx.SetWithinDelimitedBlock(true)
		oldInclude := ctx.SetIncludeBlankLine(true)
		defer func() {
			ctx.SetWithinDelimitedBlock(oldWithin)
			ctx.SetIncludeBlankLine(oldInclude)
		}()
		err = listingBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				Elements []interface{}
			}{
				Elements: elements,
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
		if author := b.Attributes.GetAsString(types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := b.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := b.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
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
		var elements = make([]interface{}, 0)
		if len(b.Elements) > 0 {
			if p, ok := b.Elements[0].(types.Paragraph); ok {
				for _, e := range p.Lines {
					elements = append(elements, e)
				}
			}
		}
		var attribution struct {
			First  string
			Second string
		}
		if author := b.Attributes.GetAsString(types.AttrQuoteAuthor); author != "" {
			attribution.First = author
			if title := b.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
				attribution.Second = title
			}
		} else if title := b.Attributes.GetAsString(types.AttrQuoteTitle); title != "" {
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
				Elements []interface{}
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

func discardTrailingBlankLines(elements []interface{}) []interface{} {
	// discard blank lines at the end
	filteredElements := make([]interface{}, len(elements))
	copy(filteredElements, elements)
	for {
		if len(filteredElements) == 0 {
			break
		}
		if _, ok := filteredElements[len(filteredElements)-1].(types.BlankLine); ok {
			log.Debugf("element of type '%T' at position %d is a blank line, discarding it", filteredElements[len(filteredElements)-1], len(filteredElements)-1)
			// remove last element of the slice since it's a blankline
			filteredElements = filteredElements[:len(filteredElements)-1]
		} else {
			break
		}
	}
	return filteredElements
}

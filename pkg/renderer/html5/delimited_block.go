package html5

import (
	"bytes"
	"fmt"
	"html"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var fencedBlockTmpl texttemplate.Template
var listingBlockTmpl texttemplate.Template
var sourceBlockTmpl texttemplate.Template
var exampleBlockTmpl texttemplate.Template
var admonitionBlockTmpl texttemplate.Template
var quoteBlockTmpl texttemplate.Template
var verseBlockTmpl texttemplate.Template
var sidebarBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	fencedBlockTmpl = newTextTemplate("listing block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="highlight"><code>{{ range $index, $element := .Elements }}{{ renderPlainString $ctx $element | printf "%s" }}{{ end }}</code></pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderPlainString": renderPlainString,
			"escape":            html.EscapeString,
		})

	listingBlockTmpl = newTextTemplate("listing block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre>{{ range $index, $element := .Elements }}{{ renderPlainString $ctx $element | printf "%s" }}{{ end }}</pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderPlainString": renderPlainString,
			"escape":            html.EscapeString,
		})

	sourceBlockTmpl = newTextTemplate("source block",
		`{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="highlight"><code{{ if .Language}} class="language-{{ .Language}}" data-lang="{{ .Language}}"{{ end }}>{{ range $index, $element := .Elements }}{{ renderPlainString $ctx $element | printf "%s" }}{{ end }}</code></pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderPlainString": renderPlainString,
			"escape":            html.EscapeString,
		})

	exampleBlockTmpl = newTextTemplate("example block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="exampleblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
{{ $elements := .Elements }}{{ renderElements $ctx $elements | printf "%s" }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         html.EscapeString,
		})

	quoteBlockTmpl = newTextTemplate("quote block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderElements $ctx .Elements | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         html.EscapeString,
		})

	verseBlockTmpl = newTextTemplate("verse block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ renderElements $ctx .Elements | printf "%s" }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         html.EscapeString,
		})

	admonitionBlockTmpl = newTextTemplate("admonition block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID}}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
{{ if .IconClass }}<i class="fa icon-{{ .IconClass }}" title="{{ .IconTitle }}"></i>{{ else }}<div class="title">{{ .IconTitle }}</div>{{ end }}
</td>
<td class="content">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}{{ renderElements $ctx .Elements | printf "%s" }}
</td>
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         html.EscapeString,
		})

	sidebarBlockTmpl = newTextTemplate("sidebar block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="sidebarblock">
<div class="content">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         html.EscapeString,
		})
}

func renderDelimitedBlock(ctx *renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block of kind '%v'", b.Attributes[types.AttrKind])
	result := bytes.NewBuffer(nil)
	elements := discardTrailingBlankLines(b.Elements)
	var id string
	if i, ok := b.Attributes[types.AttrID].(string); ok { // TODO: replace with b.Attributes.GetAsString?
		id = strings.TrimSpace(i)
	}
	var err error
	kind := b.Kind
	switch kind {
	case types.Fenced:
		previouslyWithin := ctx.SetWithinDelimitedBlock(true)
		previouslyInclude := ctx.SetIncludeBlankLine(true)
		defer func() {
			ctx.SetWithinDelimitedBlock(previouslyWithin)
			ctx.SetIncludeBlankLine(previouslyInclude)
		}()
		err = fencedBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID       string
				Title    string
				Elements []interface{}
			}{
				ID:       id,
				Title:    getTitle(b.Attributes),
				Elements: elements,
			},
		})
	case types.Listing:
		previouslyWithin := ctx.SetWithinDelimitedBlock(true)
		previouslyInclude := ctx.SetIncludeBlankLine(true)
		defer func() {
			ctx.SetWithinDelimitedBlock(previouslyWithin)
			ctx.SetIncludeBlankLine(previouslyInclude)
		}()
		err = listingBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID       string
				Title    string
				Elements []interface{}
			}{
				ID:       id,
				Title:    getTitle(b.Attributes),
				Elements: elements,
			},
		})
	case types.Source:
		previouslyWithin := ctx.SetWithinDelimitedBlock(true)
		previouslyInclude := ctx.SetIncludeBlankLine(true)
		defer func() {
			ctx.SetWithinDelimitedBlock(previouslyWithin)
			ctx.SetIncludeBlankLine(previouslyInclude)
		}()
		language := b.Attributes.GetAsString(types.AttrLanguage)
		err = sourceBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID       string
				Title    string
				Language string
				Elements []interface{}
			}{
				ID:       id,
				Title:    getTitle(b.Attributes),
				Language: language,
				Elements: elements,
			},
		})
	case types.Example:
		if k, ok := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind); ok {
			err = admonitionBlockTmpl.Execute(result, ContextualPipeline{
				Context: ctx,
				Data: struct {
					ID        string
					Class     string
					IconClass string
					IconTitle string
					Title     string
					Elements  []interface{}
				}{
					ID:        id,
					Class:     getClass(k),
					IconClass: getIconClass(ctx, k),
					IconTitle: getIconTitle(k),
					Title:     getTitle(b.Attributes),
					Elements:  elements,
				},
			})
		} else {
			// default, example block
			var title string
			if b.Attributes.Has(types.AttrTitle) {
				title = fmt.Sprintf("Example %d. %s", ctx.GetAndIncrementExampleBlockCounter(), getTitle(b.Attributes))
			}
			err = exampleBlockTmpl.Execute(result, ContextualPipeline{
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
				ID          string
				Title       string
				Attribution struct {
					First  string
					Second string
				}
				Elements []interface{}
			}{
				ID:          id,
				Title:       getTitle(b.Attributes),
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
				ID          string
				Title       string
				Attribution struct {
					First  string
					Second string
				}
				Elements []interface{}
			}{
				ID:          id,
				Title:       getTitle(b.Attributes),
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
				Title:    getTitle(b.Attributes),
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

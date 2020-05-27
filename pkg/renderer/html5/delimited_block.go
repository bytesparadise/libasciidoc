package html5

import (
	"bytes"
	"strconv"
	texttemplate "text/template"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var fencedBlockTmpl texttemplate.Template
var listingBlockTmpl texttemplate.Template
var sourceBlockTmpl texttemplate.Template
var sourceBlockContentTmpl texttemplate.Template
var exampleBlockTmpl texttemplate.Template
var admonitionBlockTmpl texttemplate.Template
var quoteBlockTmpl texttemplate.Template
var verseBlockTmpl texttemplate.Template
var verseBlockParagraphTmpl texttemplate.Template
var sidebarBlockTmpl texttemplate.Template
var passthroughBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	fencedBlockTmpl = newTextTemplate("listing block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="highlight"><code>{{ render $ctx .Elements | printf "%s" }}</code></pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"render": renderElements,
			"escape": EscapeString,
		})

	listingBlockTmpl = newTextTemplate("listing block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre>{{ render $ctx .Elements | printf "%s" }}</pre>
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"render": renderElements,
			"escape": EscapeString,
		})

	sourceBlockTmpl = newTextTemplate("source block",
		`<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="{{ if .SyntaxHighlighter }}{{ .SyntaxHighlighter }} {{ end }}highlight"><code{{ if .Language }}{{ if not .SyntaxHighlighter }} class="language-{{ .Language}}"{{ end }} data-lang="{{ .Language}}"{{ end }}>{{ .Content }}</code></pre>
</div>
</div>`,
		texttemplate.FuncMap{
			"escape": EscapeString,
		})

	sourceBlockContentTmpl = newTextTemplate("source block content",
		`{{ $ctx := .Context }}{{ with .Data }}{{ render $ctx .Elements | printf "%s" }}{{ end }}`,
		texttemplate.FuncMap{
			"render": renderElements,
		})

	exampleBlockTmpl = newTextTemplate("example block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="exampleblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         EscapeString,
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
			"escape":         EscapeString,
		})

	verseBlockTmpl = newTextTemplate("verse block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ range $index, $element := .Elements }}{{ renderElement $ctx $element | printf "%s" }}{{ end }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement": renderVerseBlockElement,
			"escape":        EscapeString,
		})

	verseBlockParagraphTmpl = newTextTemplate("verse block paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ renderLines $ctx .Lines | printf "%s" }}{{ end }}`,
		texttemplate.FuncMap{
			"renderLines": renderLines,
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
			"escape":         EscapeString,
		})

	sidebarBlockTmpl = newTextTemplate("sidebar block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="sidebarblock">
<div class="content">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
			"escape":         EscapeString,
		})

	passthroughBlockTmpl = newTextTemplate("passthrough block", `{{ $ctx := .Context }}{{ with .Data }}{{ render $ctx .Elements | printf "%s" }}{{ end }}`,
		texttemplate.FuncMap{
			"render": renderElements,
		})
}

func renderDelimitedBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block of kind '%v'", b.Attributes[types.AttrKind])
	var err error
	kind := b.Kind
	switch kind {
	case types.Fenced:
		return renderFencedBlock(ctx, b)
	case types.Listing:
		return renderListingBlock(ctx, b)
	case types.Source:
		return renderSourceBlock(ctx, b)
	case types.Example:
		return renderExampleBlock(ctx, b)
	case types.Quote, types.MarkdownQuote:
		return renderQuoteBlock(ctx, b)
	case types.Verse:
		return renderVerseBlock(ctx, b)
	case types.Sidebar:
		return renderSidebarBlock(ctx, b)
	case types.Passthrough:
		return renderPassthrough(ctx, b)
	default:
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
}

func renderFencedBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankline := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankline
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	result := bytes.NewBuffer(nil)
	err := fencedBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Elements []interface{}
		}{
			ID:       renderElementID(b.Attributes),
			Title:    renderElementTitle(b.Attributes),
			Elements: discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func renderListingBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankline := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankline
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	result := bytes.NewBuffer(nil)
	err := listingBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Elements []interface{}
		}{
			ID:       renderElementID(b.Attributes),
			Title:    renderElementTitle(b.Attributes),
			Elements: discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func renderSourceBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankline := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankline
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	// first, render the content
	contentBuf := bytes.NewBuffer(nil)
	err := sourceBlockContentTmpl.Execute(contentBuf, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Elements []interface{}
		}{
			Elements: discardTrailingBlankLines(b.Elements),
		}})
	if err != nil {
		return []byte{}, err
	}
	content := contentBuf.String()
	language := b.Attributes.GetAsString(types.AttrLanguage)

	hightligher, _ := ctx.Attributes.GetAsString(types.AttrSyntaxHighlighter)
	if language != "" && hightligher == "pygments" {
		// using github.com/alecthomas/chroma to highlight the content
		contentBuf = bytes.NewBuffer(nil)
		lexer := lexers.Get(language)
		lexer = chroma.Coalesce(lexer)
		style := styles.Fallback
		if s, exists := ctx.Attributes.GetAsString("pygments-style"); exists {
			style = styles.Get(s)
		}
		iterator, err := lexer.Tokenise(nil, content)
		if err != nil {
			return []byte{}, err
		}
		options := []html.Option{
			html.ClassPrefix("tok-"),
			html.PreventSurroundingPre(true),
		}
		// extra option: inline CSS instead of classes
		if ctx.Attributes.GetAsStringWithDefault("pygments-css", "classes") == "style" {
			options = append(options, html.WithClasses(false))
		} else {
			options = append(options, html.WithClasses(true))
		}
		// extra option: line numbers
		if b.Attributes.Has(types.AttrLineNums) {
			options = append(options, html.WithLineNumbers(true))
		}
		err = html.New(options...).Format(contentBuf, style, iterator)
		if err != nil {
			return []byte{}, err
		}
		content = contentBuf.String()
	}

	result := bytes.NewBuffer(nil)
	err = sourceBlockTmpl.Execute(result, struct {
		ID                string
		Title             string
		Language          string
		SyntaxHighlighter string
		Content           string
	}{
		ID:                renderElementID(b.Attributes),
		Title:             renderElementTitle(b.Attributes),
		SyntaxHighlighter: hightligher,
		Language:          language,
		Content:           content,
	})
	return result.Bytes(), err
}

func renderExampleBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if k, ok := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind); ok {
		err := admonitionBlockTmpl.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID        string
				Class     string
				IconClass string
				IconTitle string
				Title     string
				Elements  []interface{}
			}{
				ID:        renderElementID(b.Attributes),
				Class:     renderClass(k),
				IconClass: renderIconClass(ctx, k),
				IconTitle: renderIconTitle(k),
				Title:     renderElementTitle(b.Attributes),
				Elements:  discardTrailingBlankLines(b.Elements),
			},
		})
		return result.Bytes(), err
	}
	// default, example block
	var title string
	if b.Attributes.Has(types.AttrTitle) {
		title = "Example " + strconv.Itoa(ctx.GetAndIncrementExampleBlockCounter()) + ". " + renderElementTitle(b.Attributes)
	}
	err := exampleBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Elements []interface{}
		}{
			ID:       renderElementID(b.Attributes),
			Title:    title,
			Elements: discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func renderQuoteBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := quoteBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Elements    []interface{}
		}{
			ID:          renderElementID(b.Attributes),
			Title:       renderElementTitle(b.Attributes),
			Attribution: NewDelimitedBlockAttribution(b),
			Elements:    b.Elements,
		},
	})
	return result.Bytes(), err
}

func renderVerseBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := verseBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Elements    []interface{}
		}{
			ID:          renderElementID(b.Attributes),
			Title:       renderElementTitle(b.Attributes),
			Attribution: NewDelimitedBlockAttribution(b),
			Elements:    discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func renderVerseBlockElement(ctx renderer.Context, element interface{}) ([]byte, error) {
	previousIncludeBlankline := ctx.IncludeBlankLine
	defer func() {
		ctx.IncludeBlankLine = previousIncludeBlankline
	}()
	ctx.IncludeBlankLine = true
	switch e := element.(type) {
	case types.Paragraph:
		return renderVerseBlockParagraph(ctx, e)
	case types.BlankLine:
		return renderBlankLine(ctx, e)
	default:
		return nil, errors.Errorf("unexpected type of element to include in verse block: %T", element)
	}
}

func renderVerseBlockParagraph(ctx renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debugf("rendering paragraph with %d line(s) within a delimited block or a list", len(p.Lines))
	result := bytes.NewBuffer(nil)
	err := verseBlockParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Lines [][]interface{}
		}{
			Lines: p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderSidebarBlock(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := sidebarBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Elements []interface{}
		}{
			ID:       renderElementID(b.Attributes),
			Title:    renderElementTitle(b.Attributes),
			Elements: discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func renderPassthrough(ctx renderer.Context, b types.DelimitedBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := passthroughBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Elements []interface{}
		}{
			ID:       renderElementID(b.Attributes),
			Elements: discardTrailingBlankLines(b.Elements),
		},
	})
	return result.Bytes(), err
}

func discardTrailingBlankLines(elements []interface{}) []interface{} {
	// discard blank elements at the end
	log.Debugf("discarding trailing blank lines on %d elements...", len(elements))
	filteredElements := make([]interface{}, len(elements))
	copy(filteredElements, elements)

	for {
		if len(filteredElements) == 0 {
			break
		}
		if l, ok := filteredElements[len(filteredElements)-1].(types.VerbatimLine); ok && l.IsEmpty() {
			log.Debugf("element of type '%T' at position %d is a blank line, discarding it", filteredElements[len(filteredElements)-1], len(filteredElements)-1)
			// remove last element of the slice since it's a blankline
			filteredElements = filteredElements[:len(filteredElements)-1]
		} else if _, ok := filteredElements[len(filteredElements)-1].(types.BlankLine); ok {
			log.Debugf("element of type '%T' at position %d is a blank line, discarding it", filteredElements[len(filteredElements)-1], len(filteredElements)-1)
			// remove last element of the slice since it's a blankline
			filteredElements = filteredElements[:len(filteredElements)-1]
		} else {
			break
		}
	}
	log.Debugf("returning %d elements", len(filteredElements))
	return filteredElements
}

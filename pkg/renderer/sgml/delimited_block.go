package sgml

import (
	"bytes"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderDelimitedBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	log.Debugf("rendering delimited block of kind '%v'", b.Attributes[types.AttrKind])
	var err error
	kind := b.Kind
	switch kind {
	case types.Fenced:
		return r.renderFencedBlock(ctx, b)
	case types.Listing:
		return r.renderListingBlock(ctx, b)
	case types.Source:
		return r.renderSourceBlock(ctx, b)
	case types.Example:
		return r.renderExampleBlock(ctx, b)
	case types.Quote, types.MarkdownQuote:
		return r.renderQuoteBlock(ctx, b)
	case types.Verse:
		return r.renderVerseBlock(ctx, b)
	case types.Sidebar:
		return r.renderSidebarBlock(ctx, b)
	case types.Passthrough:
		return r.renderPassthrough(ctx, b)
	default:
		return "", errors.Wrap(err, "unable to render delimited block")
	}
}

func (r *sgmlRenderer) renderFencedBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElement(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.fencedBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       sanitized
		Title    sanitized
		Roles    sanitized
		Content  sanitized
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Roles:    r.renderElementRoles(b.Attributes),
		Content:  sanitized(content),
		Elements: elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderListingBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}

	err = r.listingBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       sanitized
		Title    sanitized
		Roles    sanitized
		Content  sanitized
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Roles:    r.renderElementRoles(b.Attributes),
		Content:  sanitized(content),
		Elements: discardTrailingBlankLines(b.Elements),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderSourceBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	// first, render the content

	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)

	if err != nil {
		return "", err
	}

	highlighter, _ := ctx.Attributes.GetAsString(types.AttrSyntaxHighlighter)
	language, found := b.Attributes.GetAsString(types.AttrLanguage)
	if found && highlighter == "pygments" {
		// using github.com/alecthomas/chroma to highlight the content
		contentBuf := &strings.Builder{}
		lexer := lexers.Get(language)
		lexer = chroma.Coalesce(lexer)
		style := styles.Fallback
		if s, found := ctx.Attributes.GetAsString("pygments-style"); found {
			style = styles.Get(s)
		}
		iterator, err := lexer.Tokenise(nil, content)
		if err != nil {
			return "", err
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
			return "", err
		}
		content = contentBuf.String()
	}

	result := &bytes.Buffer{}
	err = r.sourceBlock.Execute(result, struct {
		ID                sanitized
		Title             sanitized
		Roles             sanitized
		Language          string
		SyntaxHighlighter string
		Content           string
	}{
		ID:                r.renderElementID(b.Attributes),
		Title:             r.renderElementTitle(b.Attributes),
		SyntaxHighlighter: highlighter,
		Roles:             r.renderElementRoles(b.Attributes),
		Language:          language,
		Content:           content,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderAdmonitionBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	kind, _ := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: string(kind)}, true)
	if err != nil {
		return "", err
	}
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block content")
	}
	err = r.admonitionBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       sanitized
		Title    sanitized
		Kind     types.AdmonitionKind
		Roles    sanitized
		Icon     sanitized
		Content  sanitized
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Kind:     kind,
		Roles:    r.renderElementRoles(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Icon:     icon,
		Content:  sanitized(content),
		Elements: elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderExampleBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	if b.Attributes.Has(types.AttrAdmonitionKind) {
		return r.renderAdmonitionBlock(ctx, b)
	}
	result := &strings.Builder{}

	// default, example block
	number := ctx.GetAndIncrementExampleBlockCounter()
	elements := b.Elements
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	err = r.exampleBlock.Execute(result, struct {
		Context       *renderer.Context
		ID            sanitized
		Title         sanitized
		Roles         sanitized
		ExampleNumber int
		Content       sanitized
		Elements      []interface{}
	}{
		Context:       ctx,
		ID:            r.renderElementID(b.Attributes),
		Title:         r.renderElementTitle(b.Attributes),
		Roles:         r.renderElementRoles(b.Attributes),
		ExampleNumber: number,
		Content:       sanitized(content),
		Elements:      elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderQuoteBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}

	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}

	err = r.quoteBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          sanitized
		Title       sanitized
		Roles       sanitized
		Attribution Attribution
		Content     sanitized
		Elements    []interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       r.renderElementRoles(b.Attributes),
		Attribution: newDelimitedBlockAttribution(b),
		Content:     sanitized(content),
		Elements:    b.Elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderVerseBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content := &strings.Builder{}

	for _, item := range elements {
		s, err := r.renderVerseBlockElement(ctx, item)
		if err != nil {
			return "", errors.Wrap(err, "unable to render verse block element")
		}
		content.WriteString(s)
	}
	err := r.verseBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          sanitized
		Title       sanitized
		Roles       sanitized
		Attribution Attribution
		Content     sanitized
		Elements    []interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       r.renderElementRoles(b.Attributes),
		Attribution: newDelimitedBlockAttribution(b),
		Content:     sanitized(content.String()),
		Elements:    elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderVerseBlockElement(ctx *renderer.Context, element interface{}) (string, error) {
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.IncludeBlankLine = true
	switch e := element.(type) {
	case types.Paragraph:
		return r.renderLines(ctx, e.Lines)
	case types.BlankLine:
		return r.renderBlankLine(ctx, e)
	default:
		return "", errors.Errorf("unexpected type of element to include in verse block: %T", element)
	}
}

func (r *sgmlRenderer) renderSidebarBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}

	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render sidebar block content")
	}

	err = r.sidebarBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       sanitized
		Title    sanitized
		Roles    sanitized
		Content  sanitized
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Roles:    r.renderElementRoles(b.Attributes),
		Content:  sanitized(content),
		Elements: discardTrailingBlankLines(b.Elements),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderPassthrough(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElement(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render passthrough")
	}
	err = r.passthroughBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       sanitized
		Roles    sanitized
		Content  string
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Roles:    r.renderElementRoles(b.Attributes),
		Content:  content,
		Elements: elements,
	})
	return result.String(), err
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
			// remove last element of the slice since it's a blank line
			filteredElements = filteredElements[:len(filteredElements)-1]
		} else if _, ok := filteredElements[len(filteredElements)-1].(types.BlankLine); ok {
			log.Debugf("element of type '%T' at position %d is a blank line, discarding it", filteredElements[len(filteredElements)-1], len(filteredElements)-1)
			// remove last element of the slice since it's a blank line
			filteredElements = filteredElements[:len(filteredElements)-1]
		} else {
			break
		}
	}
	log.Debugf("returning %d elements", len(filteredElements))
	return filteredElements
}

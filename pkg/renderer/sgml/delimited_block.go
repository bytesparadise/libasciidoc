package sgml

import (
	"bytes"
	"fmt"
	"strconv"
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
	log.Debugf("rendering delimited block of kind '%v'", b.Kind)
	switch b.Kind {
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
		return "", fmt.Errorf("unable to render delimited block of kind '%v'", b.Kind)
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
	lines := discardTrailingBlankLines(b.Elements)
	content, err := r.renderLines(ctx, lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block roles")
	}

	err = r.fencedBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       string
		Title    string
		Roles    string
		Content  string
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Roles:    roles,
		Content:  content,
		Elements: lines,
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
	content, err := r.renderLines(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block roles")
	}

	err = r.listingBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Title:   r.renderElementTitle(b.Attributes),
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderSourceBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	// first, render the content
	content, highlighter, language, err := r.renderSourceLines(ctx, b)
	if err != nil {
		return "", errors.Wrap(err, "unable to render source block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render source block roles")
	}
	result := &bytes.Buffer{}
	err = r.sourceBlock.Execute(result, struct {
		ID                string
		Title             string
		Roles             string
		Language          string
		SyntaxHighlighter string
		Content           string
	}{
		ID:                r.renderElementID(b.Attributes),
		Title:             r.renderElementTitle(b.Attributes),
		SyntaxHighlighter: highlighter,
		Roles:             roles,
		Language:          language,
		Content:           content,
	})

	return result.String(), err
}

func (r *sgmlRenderer) renderSourceLines(ctx *renderer.Context, b types.DelimitedBlock) (string, string, string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true

	lines := discardTrailingBlankLines(b.Elements)
	highlighter, _ := ctx.Attributes.GetAsString(types.AttrSyntaxHighlighter)
	language, found := b.Attributes.GetAsString(types.AttrLanguage)
	if found && (highlighter == "chroma" || highlighter == "pygments") {
		ctx.EncodeSpecialChars = false
		defer func() {
			ctx.EncodeSpecialChars = true
		}()
		// using github.com/alecthomas/chroma to highlight the content
		lexer := lexers.Get(language)
		if lexer == nil {
			lexer = lexers.Fallback
		}
		lexer = chroma.Coalesce(lexer)
		style := styles.Fallback
		if s, found := ctx.Attributes.GetAsString(highlighter + "-style"); found {
			style = styles.Get(s)
		}
		options := []html.Option{
			html.ClassPrefix("tok-"),
			html.PreventSurroundingPre(true),
		}
		// extra option: inline CSS instead of classes
		if ctx.Attributes.GetAsStringWithDefault(highlighter+"-css", "classes") == "style" {
			options = append(options, html.WithClasses(false))
		} else {
			options = append(options, html.WithClasses(true))
		}
		result := &strings.Builder{}
		for i, line := range lines {
			// extra option: line numbers
			if b.Attributes.Has(types.AttrLineNums) {
				options = append(options, html.WithLineNumbers(true), html.BaseLineNumber(i+1))
			}

			renderedLine, callouts, err := r.renderSourceLine(ctx, line)
			if err != nil {
				return "", "", "", err
			}
			highlightedLineBuf := &strings.Builder{}
			// iterator, err := lexer.Tokenise(nil, content)
			iterator, err := lexer.Tokenise(nil, renderedLine)
			if err != nil {
				return "", "", "", err
			}
			if err = html.New(options...).Format(highlightedLineBuf, style, iterator); err != nil {
				return "", "", "", err
			}
			result.WriteString(highlightedLineBuf.String())
			// append callouts at the end of the highlighted line
			for _, callout := range callouts {
				renderedCallout, err := r.renderCalloutRef(callout)
				if err != nil {
					return "", "", "", err
				}
				result.WriteString(renderedCallout)
			}
			// content, err = r.renderSourceCallouts(content)
			// if err != nil {
			// 	return "", "", "", err
			// }
			if i < len(lines)-1 {
				result.WriteRune('\n')
			}

		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debug("source block content:")
			fmt.Fprintln(log.StandardLogger().Out, result.String())
		}
		return result.String(), highlighter, language, nil
	}
	result, err := r.renderLines(ctx, lines)
	if err != nil {
		if err != nil {
			return "", "", "", err
		}
	}
	return result, highlighter, language, nil
}

func (r *sgmlRenderer) renderSourceLine(ctx *renderer.Context, line interface{}) (string, []types.Callout, error) {
	elements, ok := line.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid type of line: '%T'", line)
	}
	result := strings.Builder{}
	callouts := make([]types.Callout, 0, len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case types.StringElement, types.SpecialCharacter:
			s, err := r.renderElement(ctx, e)
			if err != nil {
				return "", nil, err
			}
			result.WriteString(s)
		case types.Callout:
			callouts = append(callouts, e)
		default:
			return "", nil, fmt.Errorf("unexpected type of element: '%T'", line)
		}
	}
	return result.String(), callouts, nil
}

func (r *sgmlRenderer) renderCalloutRef(co types.Callout) (string, error) {
	result := &strings.Builder{}
	err := r.calloutRef.Execute(result, co)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout number")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderAdmonitionBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	kind, _ := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: string(kind), Attributes: b.Attributes}, true)
	if err != nil {
		return "", err
	}
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.admonitionBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       string
		Title    string
		Kind     types.AdmonitionKind
		Roles    string
		Icon     string
		Content  string
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Kind:     kind,
		Roles:    roles,
		Title:    r.renderElementTitle(b.Attributes),
		Icon:     icon,
		Content:  content,
		Elements: elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderExampleBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	if b.Attributes.Has(types.AttrAdmonitionKind) {
		return r.renderAdmonitionBlock(ctx, b)
	}
	result := &strings.Builder{}
	caption := &strings.Builder{}

	// default, example block
	number := 0
	elements := b.Elements
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	c, ok := b.Attributes.GetAsString(types.AttrCaption)
	if !ok {
		c = ctx.Attributes.GetAsStringWithDefault(types.AttrExampleCaption, "Example")
		if c != "" {
			c += " {counter:example-number}. "
		}
	}
	// TODO: Replace this hack when we have attribute substitution fully working
	if strings.Contains(c, "{counter:example-number}") {
		number = ctx.GetAndIncrementExampleBlockCounter()
		c = strings.ReplaceAll(c, "{counter:example-number}", strconv.Itoa(number))
	}
	caption.WriteString(c)
	err = r.exampleBlock.Execute(result, struct {
		Context       *renderer.Context
		ID            string
		Title         string
		Caption       string
		Roles         string
		ExampleNumber int
		Content       string
		Elements      []interface{}
	}{
		Context:       ctx,
		ID:            r.renderElementID(b.Attributes),
		Title:         r.renderElementTitle(b.Attributes),
		Caption:       caption.String(),
		Roles:         roles,
		ExampleNumber: number,
		Content:       content,
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
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.quoteBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Roles       string
		Attribution Attribution
		Content     string
		Elements    []interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       roles,
		Attribution: newDelimitedBlockAttribution(b),
		Content:     content,
		Elements:    b.Elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderVerseBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verser block")
	}
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	content, err := r.renderLines(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse block")
	}
	// for _, item := range elements {
	// 	s, err := r.renderVerseBlockElement(ctx, item)
	// 	if err != nil {
	// 		return "", errors.Wrap(err, "unable to render verse block element")
	// 	}
	// 	content.WriteString(s)
	// }

	err = r.verseBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Roles       string
		Attribution Attribution
		Content     string
		Elements    []interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       roles,
		Attribution: newDelimitedBlockAttribution(b),
		Content:     string(content),
		Elements:    elements,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderSidebarBlock(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}

	elements := discardTrailingBlankLines(b.Elements)
	content, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render sidebar block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.sidebarBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       string
		Title    string
		Roles    string
		Content  string
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Title:    r.renderElementTitle(b.Attributes),
		Roles:    roles,
		Content:  content,
		Elements: discardTrailingBlankLines(b.Elements),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderPassthrough(ctx *renderer.Context, b types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	elements := discardTrailingBlankLines(b.Elements)
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	content, err := r.renderLines(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render passthrough")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.passthroughBlock.Execute(result, struct {
		Context  *renderer.Context
		ID       string
		Roles    string
		Content  string
		Elements []interface{}
	}{
		Context:  ctx,
		ID:       r.renderElementID(b.Attributes),
		Roles:    roles,
		Content:  content,
		Elements: elements,
	})
	return result.String(), err
}

func discardTrailingBlankLines(lines []interface{}) []interface{} {
	// discard blank elements at the end
	log.Debugf("discarding trailing blank lines on %d elements...", len(lines))
	filteredLines := make([]interface{}, len(lines))
	copy(filteredLines, lines)
	// heading empty lines
	for {
		if len(filteredLines) == 0 {
			break
		}
		if l, ok := filteredLines[0].([]interface{}); ok && len(l) == 0 {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[1:]
		} else if _, ok := filteredLines[0].(types.BlankLine); ok {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[:len(filteredLines)-1]
		} else {
			break
		}
	}
	// trailing empty lines
	for {
		if len(filteredLines) == 0 {
			break
		}
		if l, ok := filteredLines[len(filteredLines)-1].([]interface{}); ok && len(l) == 0 {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[:len(filteredLines)-1]
		} else if _, ok := filteredLines[len(filteredLines)-1].(types.BlankLine); ok {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[:len(filteredLines)-1]
		} else {
			break
		}
	}
	return filteredLines
}

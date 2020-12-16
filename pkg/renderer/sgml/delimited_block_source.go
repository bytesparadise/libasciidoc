package sgml

import (
	"bytes"
	"fmt"
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

func (r *sgmlRenderer) renderSourceBlock(ctx *renderer.Context, b types.ListingBlock) (string, error) {
	// first, render the content
	content, highlighter, language, err := r.renderSourceLines(ctx, b)
	if err != nil {
		return "", errors.Wrap(err, "unable to render source block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render source block roles")
	}
	var nowrap bool
	if options, ok := b.Attributes[types.AttrOptions].([]interface{}); ok {
		for _, opt := range options {
			if opt == "nowrap" {
				nowrap = true
				break
			}
		}
	}
	result := &bytes.Buffer{}
	err = r.sourceBlock.Execute(result, struct {
		ID                string
		Title             string
		Roles             string
		Language          string
		Nowrap            bool
		SyntaxHighlighter string
		Content           string
	}{
		ID:                r.renderElementID(b.Attributes),
		Title:             r.renderElementTitle(b.Attributes),
		SyntaxHighlighter: highlighter,
		Roles:             roles,
		Language:          language,
		Nowrap:            nowrap,
		Content:           content,
	})

	return result.String(), err
}

func (r *sgmlRenderer) renderSourceParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	lines := make([][]interface{}, len(p.Lines))
	copy(lines, p.Lines)
	attributes := p.Attributes
	attributes[types.AttrStyle] = types.Source
	return r.renderSourceBlock(ctx, types.ListingBlock{
		Attributes: attributes,
		Lines:      lines,
	})
}

func (r *sgmlRenderer) renderSourceLines(ctx *renderer.Context, b types.ListingBlock) (string, string, string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true

	lines := discardEmptyLines(b.Lines)
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

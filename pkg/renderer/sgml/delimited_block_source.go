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
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderSourceBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	// first, render the content
	content, highlighter, language, err := r.renderSourceBlockElements(ctx, b)
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
	title, err := r.renderElementTitle(b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:             title,
		SyntaxHighlighter: highlighter,
		Roles:             roles,
		Language:          language,
		Nowrap:            nowrap,
		Content:           content,
	})

	return result.String(), err
}

func (r *sgmlRenderer) renderSourceParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	attributes := p.Attributes
	attributes[types.AttrStyle] = types.Source
	return r.renderSourceBlock(ctx, &types.DelimitedBlock{
		Attributes: attributes,
		Elements:   p.Elements,
	})
}

func (r *sgmlRenderer) renderSourceBlockElements(ctx *renderer.Context, b *types.DelimitedBlock) (string, string, string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	highlighter := ctx.Attributes.GetAsStringWithDefault(types.AttrSyntaxHighlighter, "")
	language := b.Attributes.GetAsStringWithDefault(types.AttrLanguage, "")

	// render without syntax highlight
	if language == "" || (highlighter != "chroma" && highlighter != "pygments") {
		log.Debug("rendering souce block without syntax highlighting")
		content, err := r.renderElements(ctx, b.Elements)
		return content, highlighter, language, err
	}

	log.Debug("rendering souce block with syntax highlighting")
	// render with syntax highlight
	lines := types.SplitElementsPerLine(b.Elements)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("splitted lines:\n%s", spew.Sdump(lines))
	}
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
	if s, found, err := ctx.Attributes.GetAsString(highlighter + "-style"); err != nil {
		return "", "", "", err
	} else if found {
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
		log.Debugf("source block content:\n%s", result.String())
	}
	return result.String(), highlighter, language, nil
}

func (r *sgmlRenderer) renderSourceLine(ctx *renderer.Context, line interface{}) (string, []*types.Callout, error) {
	elements, ok := line.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid type of line: '%T'", line)
	}
	result := strings.Builder{}
	callouts := make([]*types.Callout, 0, len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case *types.StringElement, *types.SpecialCharacter:
			s, err := r.renderElement(ctx, e)
			if err != nil {
				return "", nil, err
			}
			result.WriteString(s)
		case *types.Callout:
			callouts = append(callouts, e)
		default:
			return "", nil, fmt.Errorf("unexpected type of element: '%T'", line)
		}
	}
	return result.String(), callouts, nil
}

func (r *sgmlRenderer) renderCalloutRef(co *types.Callout) (string, error) {
	result := &strings.Builder{}
	err := r.calloutRef.Execute(result, co)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout number")
	}
	return result.String(), nil
}

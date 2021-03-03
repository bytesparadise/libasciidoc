package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	result := &strings.Builder{}
	hardbreaks := p.Attributes.HasOption(types.AttrHardBreaks) ||
		ctx.Attributes.HasOption(types.DocumentAttrHardBreaks)
	content, err := r.renderLines(ctx, p.Lines, r.withHardBreaks(hardbreaks))
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph content")
	}
	if k, ok := p.Attributes[types.AttrStyle].(string); ok {
		switch k {
		case types.Example:
			return r.renderExampleParagraph(ctx, p)
		case types.Listing:
			return r.renderListingParagraph(ctx, p)
		case types.Source:
			return r.renderSourceParagraph(ctx, p)
		case types.Verse:
			return r.renderVerseParagraph(ctx, p)
		case types.Quote:
			return r.renderQuoteParagraph(ctx, p)
		case types.Tip, types.Note, types.Important, types.Warning, types.Caution:
			return r.renderAdmonitionParagraph(ctx, p)
		case "manpage":
			return r.renderManpageNameParagraph(ctx, p)
		default:
			// do nothing, will move to default below
		}
	} else if ctx.WithinDelimitedBlock || ctx.WithinList > 0 {
		return r.renderParagraphWithinDelimitedBlock(ctx, p)
	}
	// default case
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph roles")
	}
	title, err := r.renderElementTitle(p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph roles")
	}
	log.Debug("rendering a standalone paragraph")
	err = r.paragraph.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Roles   string
		Title   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(p.Attributes),
		Title:   title,
		Roles:   roles,
		Content: string(content),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderManpageNameParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	log.Debug("rendering name section paragraph in manpage...")
	result := &strings.Builder{}

	content, err := r.renderLines(ctx, p.Lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}

	err = r.manpageNameParagraph.Execute(result, struct {
		Context *renderer.Context
		Content string
		Lines   [][]interface{}
	}{
		Context: ctx,
		Content: string(content),
		Lines:   p.Lines,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderParagraphWithinDelimitedBlock(ctx *renderer.Context, p types.Paragraph) (string, error) {
	// log.Debugf("rendering paragraph with %d line(s) within a delimited block or a list", len(p.Lines))
	result := &strings.Builder{}

	content, err := r.renderLines(ctx, p.Lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render delimited block paragraph content")
	}
	title, err := r.renderElementTitle(p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render delimited block paragraph content")
	}

	err = r.delimitedBlockParagraph.Execute(result, struct {
		Context    *renderer.Context
		ID         string
		Title      string
		CheckStyle string
		Content    string
		Lines      [][]interface{}
	}{
		Context:    ctx,
		ID:         r.renderElementID(p.Attributes),
		Title:      title,
		CheckStyle: renderCheckStyle(p.Attributes[types.AttrCheckStyle]),
		Content:    string(content),
		Lines:      p.Lines,
	})
	return result.String(), err
}

func renderCheckStyle(style interface{}) string {
	// default checkboxes
	switch style {
	case types.Checked:
		return "&#10003; "
	case types.CheckedInteractive:
		return `<input type="checkbox" data-item-complete="1" checked> `
	case types.Unchecked:
		return "&#10063; "
	case types.UncheckedInteractive:
		return `<input type="checkbox" data-item-complete="0"> `
	default:
		return ""
	}
}

func (r *sgmlRenderer) renderElementTitle(attrs types.Attributes) (string, error) {
	if title, found, err := attrs.GetAsString(types.AttrTitle); err != nil {
		return "", err
	} else if found {
		result := EscapeString(strings.TrimSpace(title))
		// log.Debugf("rendered title: '%s'", result)
		return result, nil
	}
	log.Debug("no title to render")
	return "", nil
}

// RenderLinesConfig the config to use when rendering paragraph lines
type RenderLinesConfig struct {
	render     renderFunc
	hardBreaks bool
}

// RenderLinesOption an option to configure the rendering
type RenderLinesOption func(c *RenderLinesConfig)

// WithHardBreaks sets the hard break option
func (r *sgmlRenderer) withHardBreaks(hardBreaks bool) RenderLinesOption {
	return func(c *RenderLinesConfig) {
		c.hardBreaks = hardBreaks
	}
}

// PlainText sets the render func to PlainText instead of SGML
func (r *sgmlRenderer) withPlainText() RenderLinesOption {
	return func(c *RenderLinesConfig) {
		c.render = r.renderPlainText
	}
}

// renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// and includes an `\n` character in-between, until the last one.
// Trailing spaces are removed for each line.
func (r *sgmlRenderer) renderLines(ctx *renderer.Context, lines [][]interface{}, options ...RenderLinesOption) (string, error) { // renderLineFunc renderFunc, hardbreak bool
	linesRenderer := RenderLinesConfig{
		render:     r.renderLine,
		hardBreaks: false,
	}
	for _, apply := range options {
		apply(&linesRenderer)
	}
	buf := &strings.Builder{}
	for i, e := range lines {
		renderedLine, err := linesRenderer.render(ctx, e)
		if err != nil {
			return "", errors.Wrap(err, "unable to render lines")
		}
		if len(renderedLine) > 0 {
			if _, err := buf.WriteString(renderedLine); err != nil {
				return "", errors.Wrap(err, "unable to render lines")
			}
		}

		if i < len(lines)-1 && (len(renderedLine) > 0 || ctx.WithinDelimitedBlock) {
			// log.Debugf("rendered line is not the last one in the slice")
			var err error
			if linesRenderer.hardBreaks {
				if br, err := r.renderLineBreak(); err != nil {
					return "", errors.Wrap(err, "unable to render hardbreak")
				} else if _, err = buf.WriteString(br); err != nil {
					return "", errors.Wrap(err, "unable to write hardbreak")
				}
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				return "", errors.Wrap(err, "unable to render lines")
			}
		}
	}
	// log.Debugf("rendered lines: '%s'", buf.String())
	return buf.String(), nil
}

func (r *sgmlRenderer) renderLine(ctx *renderer.Context, element interface{}) (string, error) {
	if elements, ok := element.([]interface{}); ok {
		return r.renderInlineElements(ctx, elements)
	}

	return "", errors.Errorf("invalid type of element for a line: %T", element)
}

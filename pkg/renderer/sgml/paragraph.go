package sgml

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	// if len(p.Lines) == 0 {
	// 	return make([]byte, 0), nil
	// }
	result := &bytes.Buffer{}
	id := r.renderElementID(p.Attributes)
	var err error
	if _, ok := p.Attributes[types.AttrAdmonitionKind]; ok {
		return r.renderAdmonitionParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Source {
		return r.renderSourceParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Verse {
		return r.renderVerseParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == types.Quote {
		return r.renderQuoteParagraph(ctx, p)
	} else if kind, ok := p.Attributes[types.AttrKind]; ok && kind == "manpage" {
		return r.renderManpageNameParagraph(ctx, p)
	} else if ctx.WithinDelimitedBlock || ctx.WithinList > 0 {
		return r.renderDelimitedBlockParagraph(ctx, p)
	} else {
		log.Debug("rendering a standalone paragraph")
		err = r.paragraph.Execute(result, ContextualPipeline{
			Context: ctx,
			Data: struct {
				ID         string
				Class      string
				Title      string
				Lines      [][]interface{}
				HardBreaks RenderLinesOption
			}{
				ID:         id,
				Class:      getParagraphClass(p),
				Title:      r.renderElementTitle(p.Attributes),
				Lines:      p.Lines,
				HardBreaks: r.withHardBreaks(p.Attributes.Has(types.AttrHardBreaks) || ctx.Attributes.Has(types.DocumentAttrHardBreaks)),
			},
		})
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	// log.Debugf("rendered paragraph: '%s'", result.String())
	return result.Bytes(), nil
}

func (r *sgmlRenderer) renderAdmonitionParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering admonition paragraph...")
	result := &bytes.Buffer{}
	k, ok := p.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
	if !ok {
		return nil, errors.Errorf("failed to render admonition with unknown kind: %T", p.Attributes[types.AttrAdmonitionKind])
	}
	err := r.admonitionParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID        string
			Title     string
			Class     string
			IconTitle string
			IconClass string
			Lines     [][]interface{}
		}{
			ID:        r.renderElementID(p.Attributes),
			Title:     r.renderElementTitle(p.Attributes),
			Class:     renderClass(k),
			IconTitle: renderIconTitle(k),
			IconClass: renderIconClass(ctx, k),
			Lines:     p.Lines,
		},
	})
	return result.Bytes(), err
}

func (r *sgmlRenderer) renderSourceParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering source paragraph...")
	result := &bytes.Buffer{}
	err := r.sourceParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID       string
			Title    string
			Language string
			Lines    [][]interface{}
		}{
			ID:       r.renderElementID(p.Attributes),
			Title:    r.renderElementTitle(p.Attributes),
			Language: p.Attributes.GetAsStringWithDefault(types.AttrLanguage, ""),
			Lines:    p.Lines,
		},
	})
	return result.Bytes(), err
}

func (r *sgmlRenderer) renderVerseParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering verse paragraph...")
	result := &bytes.Buffer{}
	err := r.verseParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Lines       [][]interface{}
		}{
			ID:          r.renderElementID(p.Attributes),
			Title:       r.renderElementTitle(p.Attributes),
			Attribution: newParagraphAttribution(p),
			Lines:       p.Lines,
		},
	})
	return result.Bytes(), err
}

func (r *sgmlRenderer) renderQuoteParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering quote paragraph...")
	result := &bytes.Buffer{}
	err := r.quoteParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID          string
			Title       string
			Attribution Attribution
			Lines       [][]interface{}
		}{
			ID:          r.renderElementID(p.Attributes),
			Title:       r.renderElementTitle(p.Attributes),
			Attribution: newParagraphAttribution(p),
			Lines:       p.Lines,
		},
	})
	return result.Bytes(), err
}

func (r *sgmlRenderer) renderManpageNameParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debug("rendering name section paragraph in manpage...")
	result := &bytes.Buffer{}
	err := r.manpageNameParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Lines [][]interface{}
		}{
			Lines: p.Lines,
		},
	})
	return result.Bytes(), err
}

func (r *sgmlRenderer) renderDelimitedBlockParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	log.Debugf("rendering paragraph with %d line(s) within a delimited block or a list", len(p.Lines))
	result := &bytes.Buffer{}
	err := r.delimitedBlockParagraph.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID         string
			Title      string
			CheckStyle string
			Lines      [][]interface{}
		}{
			ID:         r.renderElementID(p.Attributes),
			Title:      r.renderElementTitle(p.Attributes),
			CheckStyle: renderCheckStyle(p.Attributes[types.AttrCheckStyle]),
			Lines:      p.Lines,
		},
	})
	return result.Bytes(), err
}

func renderCheckStyle(style interface{}) string {
	switch style {
	case types.Unchecked:
		return "&#10063; "
	case types.Checked:
		return "&#10003; "
	default:
		return ""
	}
}

func renderIconClass(ctx *renderer.Context, kind types.AdmonitionKind) string {
	if icons, _ := ctx.Attributes.GetAsString("icons"); icons == "font" {
		return renderClass(kind)
	}
	return ""
}

func renderClass(kind types.AdmonitionKind) string {
	switch kind {
	case types.Tip:
		return "tip"
	case types.Note:
		return "note"
	case types.Important:
		return "important"
	case types.Warning:
		return "warning"
	case types.Caution:
		return "caution"
	default:
		log.Errorf("unexpected kind of admonition: %v", kind)
		return ""
	}
}

func renderIconTitle(kind types.AdmonitionKind) string {
	switch kind {
	case types.Tip:
		return "Tip"
	case types.Note:
		return "Note"
	case types.Important:
		return "Important"
	case types.Warning:
		return "Warning"
	case types.Caution:
		return "Caution"
	default:
		log.Errorf("unexpected kind of admonition: %v", kind)
		return ""
	}
}

func getParagraphClass(p types.Paragraph) string {
	result := "paragraph"
	if role, found := p.Attributes.GetAsString(types.AttrRole); found {
		result = result + " " + role
	}
	return result
}

func (r *sgmlRenderer) renderElementTitle(attrs types.Attributes) string {
	if title, found := attrs.GetAsString(types.AttrTitle); found {
		return strings.TrimSpace(title)
	}
	return ""
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
func (r *sgmlRenderer) renderLines(ctx *renderer.Context, lines [][]interface{}, options ...RenderLinesOption) ([]byte, error) { // renderLineFunc renderFunc, hardbreak bool
	linesRenderer := RenderLinesConfig{
		render:     r.renderLine,
		hardBreaks: false,
	}
	for _, apply := range options {
		apply(&linesRenderer)
	}
	buf := &bytes.Buffer{}
	for i, e := range lines {
		renderedElement, err := linesRenderer.render(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render lines")
		}
		if len(renderedElement) > 0 {
			_, err := buf.Write(renderedElement)
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}

		if i < len(lines)-1 && (len(renderedElement) > 0 || ctx.WithinDelimitedBlock) {
			// log.Debugf("rendered line is not the last one in the slice")
			var err error
			if linesRenderer.hardBreaks {
				_, err = buf.WriteString("<br>\n") // TODO: linebreak template
			} else {
				_, err = buf.WriteString("\n")
			}
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}
	}
	// log.Debugf("rendered lines: '%s'", buf.String())
	return buf.Bytes(), nil
}

func (r *sgmlRenderer) renderLine(ctx *renderer.Context, element interface{}) ([]byte, error) {
	if elements, ok := element.([]interface{}); ok {
		return r.renderInlineElements(ctx, elements)
	}

	return nil, errors.Errorf("invalid type of element for a line: %T", element)
}

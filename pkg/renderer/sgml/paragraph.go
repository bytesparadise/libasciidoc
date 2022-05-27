package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debugf("rendering a regular paragraph with style '%v' and embedded: %t", p.Attributes[types.AttrStyle], (ctx.WithinDelimitedBlock || ctx.WithinList > 0))
	if ctx.WithinDelimitedBlock || ctx.WithinList > 0 {
		return r.renderEmbeddedParagraph(ctx, p, "")
	}
	switch p.Attributes[types.AttrStyle] {
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
	case types.Passthrough:
		return r.renderPassthroughParagraph(ctx, p)
	case "manpage":
		return r.renderManpageNameParagraph(ctx, p)
	case types.Tip, types.Note, types.Important, types.Warning, types.Caution:
		return r.renderAdmonitionParagraph(ctx, p)
	case types.LiteralParagraph, types.Literal:
		return r.renderLiteralParagraph(ctx, p)
	default:
		// default case
		return r.renderRegularParagraph(ctx, p)
	}
}

func (r *sgmlRenderer) renderRegularParagraph(ctx *renderer.Context, p *types.Paragraph, opts ...lineRendererOption) (string, error) {
	log.Debug("rendering a regular paragraph")
	content, err := r.renderParagraphElements(ctx, p, opts...)
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph content")
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph roles")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph roles")
	}
	result := &strings.Builder{}
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
		Content: strings.Trim(string(content), "\n"),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render paragraph")
	}
	return result.String(), nil

}

func (r *sgmlRenderer) renderManpageNameParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debug("rendering name section paragraph in manpage...")
	result := &strings.Builder{}

	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}

	err = r.manpageNameParagraph.Execute(result, struct {
		Context *renderer.Context
		Content string
	}{
		Context: ctx,
		Content: string(content),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderEmbeddedParagraph(ctx *renderer.Context, p *types.Paragraph, class string) (string, error) {
	log.Debug("rendering paragraph within a delimited block or a list")
	result := &strings.Builder{}

	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render delimited block paragraph content")
	}

	err = r.embeddedParagraph.Execute(result, struct {
		Context    *renderer.Context
		CheckStyle string
		Class      string
		Content    string
	}{
		Context:    ctx,
		Class:      class,
		CheckStyle: renderCheckStyle(p.Attributes[types.AttrCheckStyle]),
		Content:    trimSpaces(content),
	})
	return result.String(), err
}

// trimSpaces removes heading and trailing spaces on each line of the given content
func trimSpaces(content string) string {
	// trim spaces
	contentBuf := &strings.Builder{}
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		contentBuf.WriteString(strings.TrimSpace(line))
		if i < len(lines)-1 {
			contentBuf.WriteString("\n")
		}
	}
	return contentBuf.String()
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

func (r *sgmlRenderer) renderElementTitle(ctx *renderer.Context, attrs types.Attributes) (string, error) {
	title, found := attrs[types.AttrTitle]
	if !found {
		log.Debug("no title to render")
		return "", nil
	}
	switch title := title.(type) {
	case string:
		return title, nil
	case []interface{}:
		return r.renderElements(ctx, title)
	default:
		return "", errors.Errorf("unable to render title of type '%T'", title)
	}
}

type lineRenderer struct {
	render     renderFunc
	hardBreaks bool
}

func (r *sgmlRenderer) newLineRenderer(opts ...lineRendererOption) *lineRenderer {
	lr := &lineRenderer{
		render: r.renderElement,
	}
	for _, apply := range opts {
		apply(lr)
	}
	return lr
}

// RenderLinesOption an option to configure the rendering
type lineRendererOption func(c *lineRenderer)

// WithHardBreaks sets the hard break option
func withHardBreaks(hardBreaks bool) lineRendererOption {
	return func(lr *lineRenderer) {
		lr.hardBreaks = hardBreaks
	}
}

// withRenderer sets the render func
func withRenderer(f renderFunc) lineRendererOption {
	return func(c *lineRenderer) {
		c.render = f
	}
}

func (r *sgmlRenderer) renderParagraphElements(ctx *renderer.Context, p *types.Paragraph, opts ...lineRendererOption) (string, error) {
	hardbreaks := p.Attributes.HasOption(types.AttrHardBreaks) || ctx.Attributes.HasOption(types.AttrHardBreaks)
	lr := r.newLineRenderer(append(opts, withHardBreaks(hardbreaks))...)
	buf := &strings.Builder{}
	for _, e := range p.Elements {
		renderedElement, err := lr.render(ctx, e)
		if err != nil {
			return "", errors.Wrap(err, "unable to render paragraph elements")
		}
		if _, err := buf.WriteString(renderedElement); err != nil {
			return "", errors.Wrap(err, "unable to render paragraph elements")
		}
	}
	result := buf.String()
	if lr.hardBreaks { // TODO: move within the call to `render`?
		linebreak := &strings.Builder{}
		if err := r.lineBreak.Execute(linebreak, nil); err != nil {
			return "", err
		}
		result = strings.ReplaceAll(result, "\n", linebreak.String()+"\n")
	}
	return result, nil
}

package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

type linesRenderer struct {
	render renderFunc
}

type renderLinesOption func(config *linesRenderer)

func (r *sgmlRenderer) withVerbatim() renderLinesOption {
	return func(config *linesRenderer) {
		config.render = r.renderPlainText
	}
}

func (r *sgmlRenderer) renderInlineElements(ctx *renderer.Context, elements []interface{}, options ...renderLinesOption) (string, error) {
	if len(elements) == 0 {
		return "", nil
	}
	log.Debugf("rendering line with %d element(s)...", len(elements))
	lr := linesRenderer{
		render: r.renderElement,
	}
	for _, apply := range options {
		apply(&lr)
	}
	// first pass or rendering, using the provided `renderElementFunc`:
	buf := &strings.Builder{}
	for i, element := range elements {
		renderedElement, err := lr.render(ctx, element)
		if err != nil {
			return "", err
		}
		if i == len(elements)-1 {
			if _, ok := element.(types.StringElement); ok { // TODO: only for StringElement? or for any kind of element?
				// trim trailing spaces before returning the line
				buf.WriteString(strings.TrimRight(string(renderedElement), " "))
				// log.Debugf("trimmed spaces on '%v'", string(renderedElement))
			} else {
				buf.WriteString(renderedElement)
			}
		} else {
			buf.WriteString(renderedElement)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendered inline elements: '%s'", buf.String())
	}
	return buf.String(), nil
}

type renderFunc func(*renderer.Context, interface{}) (string, error)

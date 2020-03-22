package html5

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type linesRenderer struct {
	render renderFunc
}

type renderLinesOption func(config *linesRenderer)

func verbatim() renderLinesOption {
	return func(config *linesRenderer) {
		config.render = renderPlainText
	}
}

func renderInlineElements(ctx renderer.Context, elements []interface{}, options ...renderLinesOption) ([]byte, error) {
	log.Debugf("rendering line with %d element(s)...", len(elements))
	r := linesRenderer{
		render: renderElement,
	}
	for _, apply := range options {
		apply(&r)
	}
	// first pass or rendering, using the provided `renderElementFunc`:
	buf := bytes.NewBuffer(nil)
	for i, element := range elements {
		renderedElement, err := r.render(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render line")
		}
		if i == len(elements)-1 {
			if _, ok := element.(types.StringElement); ok { // TODO: only for StringElement? or for any kind of element?
				// trim trailing spaces before returning the line
				buf.WriteString(strings.TrimRight(string(renderedElement), " "))
				log.Debugf("trimmed spaces on '%v'", string(renderedElement))
			} else {
				buf.Write(renderedElement)
			}
		} else {
			buf.Write(renderedElement)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendered inlines elements: '%s'", buf.String())
	}
	return buf.Bytes(), nil
}

type renderFunc func(renderer.Context, interface{}) ([]byte, error)

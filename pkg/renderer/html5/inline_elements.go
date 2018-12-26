package html5

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// // renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// // and includes an `\n` character in-between, until the last one.
// // Trailing spaces are removed for each line.
// func renderLinesWithHardbreak(ctx *renderer.Context, elements []types.InlineElements, hardbreak bool) (string, error) {
// 	r, err := renderLines(ctx, elements)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(r), nil
// }

func renderLinesAsString(ctx *renderer.Context, elements []types.InlineElements, hardbreak bool) (string, error) {
	result, err := renderLines(ctx, elements, renderElement, hardbreak)
	return string(result), err
}

// renderLines renders all lines (i.e, all `InlineElements`` - each `InlineElements` being a slice of elements to generate a line)
// and includes an `\n` character in-between, until the last one.
// Trailing spaces are removed for each line.
func renderLines(ctx *renderer.Context, elements []types.InlineElements, renderElementFunc rendererFunc, hardbreak bool) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for i, e := range elements {
		renderedElement, err := renderElementFunc(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render lines")
		}
		if len(renderedElement) > 0 {
			_, err := buff.Write(renderedElement)
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}

		if i < len(elements)-1 && (len(renderedElement) > 0 || ctx.WithinDelimitedBlock()) {
			log.Debugf("rendered line is not the last one in the slice")
			var err error
			if hardbreak {
				_, err = buff.WriteString("<br>\n")
			} else {
				_, err = buff.WriteString("\n")
			}
			if err != nil {
				return nil, errors.Wrap(err, "unable to render lines")
			}
		}
	}
	log.Debugf("rendered lines: '%s'", buff.String())
	return buff.Bytes(), nil
}

func renderLine(ctx *renderer.Context, elements types.InlineElements, renderElementFunc rendererFunc) ([]byte, error) {
	log.Debugf("rendering line with %d element(s)...", len(elements))
	elements, err := applySubstitutions(ctx, elements)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render line")
	}

	buff := bytes.NewBuffer(nil)
	for i, element := range elements {
		renderedElement, err := renderElementFunc(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render line")
		}
		if i == len(elements)-1 {
			if _, ok := element.(types.StringElement); ok { // TODO: only for StringElement? or for any kind of element?
				// trim trailing spaces before returning the line
				buff.WriteString(strings.TrimRight(string(renderedElement), " "))
				log.Debugf("trimmed spaces on '%v'", string(renderedElement))
			} else {
				buff.Write(renderedElement)
			}
		} else {
			buff.Write(renderedElement)
		}
	}
	log.Debugf("rendered elements: '%s'", buff.String())
	return buff.Bytes(), nil
}

// check if there's at least on substitution before doing the whole process
func hasSubstitutions(e types.InlineElements) bool {
	for _, element := range e {
		if _, ok := element.(types.DocumentAttributeSubstitution); ok {
			return true
		}
	}
	return false
}

func applySubstitutions(ctx *renderer.Context, e types.InlineElements) (types.InlineElements, error) {
	if !hasSubstitutions(e) {
		log.Debug("no subsitution in the line of elements")
		return e, nil
	}
	log.Debugf("applying substitutions on %v (%d)", e, len(e))
	// apply substitution...
	s := make([]interface{}, 0)
	for _, element := range e {
		switch element := element.(type) {
		case types.DocumentAttributeSubstitution:
			r, err := renderAttributeSubstitution(ctx, element)
			if err != nil {
				return types.InlineElements{}, errors.Wrap(err, "failed to apply substitution")
			}
			s = append(s, types.NewStringElement(string(r)))
		default:
			s = append(s, element)
		}
	}
	// ... and then see with surrounding elements
	// if anything can be parsed again
	s, err := types.NewInlineElements(s...)
	if err != nil {
		return types.InlineElements{}, errors.Wrap(err, "failed to apply substitution")
	}
	log.Debugf("substitution(s) result: %v (%d)", s, len(s))
	// now parse the StringElements
	result := make([]interface{}, 0)
	for _, element := range s {
		log.Debugf("now processing element of type %T", element)
		switch element := element.(type) {
		case types.StringElement:
			r, err := parser.Parse("",
				[]byte(element.Content),
				parser.Entrypoint("InlineElementsWithoutSubtitution")) // parse a single InlineElements
			if err != nil {
				return types.InlineElements{}, errors.Wrap(err, "failed process elements after substitution")
			}
			if r, ok := r.(types.InlineElements); ok {
				// here the doc should have directly an InlineElements since we specified a specific entrypoint for the parser
				result = append(result, r...)
			} else {
				return types.InlineElements{}, errors.Errorf("expected an groupg of elements, but got a result of type %T", r)
			}
		default:
			result = append(result, element)
		}
	}
	log.Debugf("parsing after substitution(s): %v (%d)", result, len(result))
	return result, nil
}

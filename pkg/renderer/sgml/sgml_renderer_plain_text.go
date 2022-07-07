package sgml

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func RenderPlainText(element interface{}) (string, error) {
	r := &plaintextRenderer{}
	return r.render(element)
}

func RenderParagraphElements(p *types.Paragraph) (string, error) {
	r := &plaintextRenderer{}
	buf := &strings.Builder{}
	for _, e := range p.Elements {
		renderedElement, err := r.render(e)
		if err != nil {
			return "", errors.Wrap(err, "unable to render paragraph elements")
		}
		if _, err := buf.WriteString(renderedElement); err != nil {
			return "", errors.Wrap(err, "unable to render paragraph elements")
		}
	}
	result := buf.String()
	return result, nil
}

type plaintextRenderer struct {
}

func (r *plaintextRenderer) render(element interface{}) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering plain string for element of type %T", element)
	}
	switch e := element.(type) {
	case []interface{}:
		return r.renderInlineElements(e)
	case *types.QuotedText:
		return r.render(e.Elements)
	case *types.Icon:
		return e.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""), nil
	case *types.InlineImage:
		return e.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""), nil
	case *types.InlineLink:
		return r.renderInlineLink(e)
	case *types.BlankLine, types.ThematicBreak:
		return "\n\n", nil
	case *types.SpecialCharacter:
		return e.Name, nil
	case *types.Symbol:
		return r.renderSymbol(e)
	case *types.StringElement:
		return e.Content, nil
	case *types.FootnoteReference:
		return r.renderFootnoteReference(e)
	default:
		return "", errors.Errorf("unable to render plain string for element of type '%T'", e)
	}
}

func (r *plaintextRenderer) renderInlineElements(elements []interface{}) (string, error) {
	// log.Debugf("rendering line with %d element(s)...", len(elements))
	buf := &strings.Builder{}
	for i, element := range elements {
		renderedElement, err := r.render(element)
		if err != nil {
			return "", err
		}
		if i == len(elements)-1 {
			if _, ok := element.(*types.StringElement); ok { // TODO: only for StringElement? or for any kind of element?
				// trim trailing spaces before returning the line
				buf.WriteString(strings.TrimRight(string(renderedElement), " "))
			} else {
				buf.WriteString(renderedElement)
			}
		} else {
			buf.WriteString(renderedElement)
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendered inline elements: '%s'", buf.String())
	// }
	return buf.String(), nil
}

func (r *plaintextRenderer) renderSymbol(s *types.Symbol) (string, error) {
	if v, found := symbols[s.Name]; found {
		return s.Prefix + v, nil
	}
	return s.Prefix + s.Name, nil
}

func (r *plaintextRenderer) renderInlineLink(l *types.InlineLink) (string, error) {
	switch alt := l.Attributes[types.AttrInlineLinkText].(type) {
	case []interface{}:
		return r.render(alt)
	case string:
		return alt, nil
	default:
		return l.Location.ToDisplayString(), nil
	}
}

func (r *plaintextRenderer) renderFootnoteReference(note *types.FootnoteReference) (string, error) {
	if note.ID != types.InvalidFootnoteReference {
		// valid case for a footnote with content, with our without an explicit reference
		return `<sup class="footnote">[` + strconv.Itoa(note.ID) + `]</sup>`, nil
	}
	return "", fmt.Errorf("unable to render missing footnote")
}

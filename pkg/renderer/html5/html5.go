package html5

import (
	"bytes"
	"io"
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx *renderer.Context, output io.Writer) (map[string]interface{}, error) {
	return renderDocument(ctx, output)
}

func renderElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	log.Debugf("rendering %d elements(s)...", len(elements))
	buff := bytes.NewBuffer(nil)
	hasContent := false
	if !ctx.IncludeHeaderFooter() && len(elements) > 0 {
		if s, ok := elements[0].(types.Section); ok && s.Level == 0 {
			// don't render the top-level section, but only its elements (plus the rest if there's anything)
			if len(elements) > 1 {
				elements = append(s.Elements, elements[1:])
			} else {
				elements = s.Elements
			}
		}
	}
	for _, element := range elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render an element")
		}
		// insert new line if there's already some content (except for BlankLine)
		_, isBlankline := element.(types.BlankLine)
		if !isBlankline && hasContent && len(renderedElement) > 0 {
			buff.WriteString("\n")
		}
		buff.Write(renderedElement)
		if len(renderedElement) > 0 {
			hasContent = true
		}
	}
	// log.Debugf("rendered elements: '%s'", buff.String())
	return buff.Bytes(), nil
}

// renderListElements is similar to the `renderElements` func above,
// but it sets the `withinList` context flag to true for the first element only
func renderListElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	log.Debugf("rendering list with %d element(s)...", len(elements))
	buff := bytes.NewBuffer(nil)
	hasContent := false
	for i, element := range elements {
		if i == 0 {
			ctx.SetWithinList(true)
		}
		renderedElement, err := renderElement(ctx, element)
		if i == 0 {
			ctx.SetWithinList(false)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render a list block")
		}
		// insert new line if there's already some content
		if hasContent && len(renderedElement) > 0 {
			buff.WriteString("\n")
		}
		buff.Write(renderedElement)
		if len(renderedElement) > 0 {
			hasContent = true
		}
	}
	// log.Debugf("rendered elements: '%s'", buff.String())
	return buff.Bytes(), nil
}

// nolint: gocyclo
func renderElement(ctx *renderer.Context, element interface{}) ([]byte, error) {
	// log.Debugf("rendering element of type `%T`", element)
	switch e := element.(type) {
	case []interface{}:
		return renderElements(ctx, e)
	case types.TableOfContentsMacro:
		return renderTableOfContents(ctx, e)
	case types.Section:
		return renderSection(ctx, e)
	case types.Preamble:
		return renderPreamble(ctx, e)
	case types.BlankLine:
		return renderBlankLine(ctx, e)
	case types.LabeledList:
		return renderLabeledList(ctx, e)
	case types.OrderedList:
		return renderOrderedList(ctx, e)
	case types.UnorderedList:
		return renderUnorderedList(ctx, e)
	case types.Paragraph:
		return renderParagraph(ctx, e)
	case types.CrossReference:
		return renderCrossReference(ctx, e)
	case types.QuotedText:
		return renderQuotedText(ctx, e)
	case types.Passthrough:
		return renderPassthrough(ctx, e)
	case types.ImageBlock:
		return renderImageBlock(ctx, e)
	case types.InlineImage:
		return renderInlineImage(ctx, e)
	case types.DelimitedBlock:
		return renderDelimitedBlock(ctx, e)
	case types.Table:
		return renderTable(ctx, e)
	case types.LiteralBlock:
		return renderLiteralBlock(ctx, e)
	case types.InlineLink:
		return renderLink(ctx, e)
	case types.StringElement:
		return renderStringElement(ctx, e)
	case types.Footnote:
		return renderFootnote(ctx, e)
	case types.LineBreak:
		return renderLineBreak()
	case types.UserMacro:
		return renderUserMacro(ctx, e)
	case types.SingleLineComment:
		return nil, nil // nothing to do
	default:
		return nil, errors.Errorf("unsupported type of element: %T", element)
	}
}

// nolint: gocyclo
func renderPlainText(ctx *renderer.Context, element interface{}) ([]byte, error) {
	log.Debugf("rendering plain string for element of type %T", element)
	switch element := element.(type) {
	case []interface{}:
		return renderInlineElements(ctx, element, verbatim())
	case [][]interface{}:
		return renderLines(ctx, element, PlainText())
	case types.QuotedText:
		return renderPlainText(ctx, element.Elements)
	case types.InlineImage:
		return []byte(element.Attributes.GetAsString(types.AttrImageAlt)), nil
	case types.InlineLink:
		if alt, ok := element.Attributes[types.AttrInlineLinkText].([]interface{}); ok {
			return renderPlainText(ctx, alt)
		}
		return []byte(element.Location.String()), nil
	case types.BlankLine:
		return []byte("\n\n"), nil
	case types.StringElement:
		return []byte(element.Content), nil
	case types.Paragraph:
		return renderLines(ctx, element.Lines, PlainText())
	default:
		return nil, errors.Errorf("unable to render plain string for element of type '%T'", element)
	}
}

// includeNewline returns an "\n" sequence if the given index is NOT the last entry in the given description lines, empty string otherwise.
// also, it ignores the element if it is a blankline, depending on the context
func includeNewline(ctx renderer.Context, index int, content interface{}) string {
	switch reflect.TypeOf(content).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(content)
		if _, match := s.Index(index).Interface().(types.BlankLine); match {
			if ctx.IncludeBlankLine() {
				return "\n"
			}
			return ""
		}
		if index < s.Len()-1 {
			return "\n"
		}
	default:
		log.Warnf("content of type '%T' is not an array or a slice", content)
	}
	return ""
}

package sgml

import (
	"bytes"
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	log.Debugf("rendering %d elements(s)...", len(elements))
	buff := &bytes.Buffer{}
	hasContent := false
	if !ctx.Config.IncludeHeaderFooter && len(elements) > 0 {
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
		renderedElement, err := r.renderElement(ctx, element)
		if err != nil {
			return nil, err // no need to wrap the error here
		}
		// insert new line if there's already some content (except for BlankLine)
		_, isBlankline := element.(types.BlankLine)
		_, isVerbatimLine := element.(types.VerbatimLine)
		if hasContent && (isVerbatimLine || (!isBlankline && len(renderedElement) > 0)) {
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
func (r *sgmlRenderer) renderListElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	log.Debugf("rendering list with %d element(s)...", len(elements))
	buff := &bytes.Buffer{}
	hasContent := false
	for i, element := range elements {
		if i == 0 {
			ctx.WithinList++
		}
		renderedElement, err := r.renderElement(ctx, element)
		if i == 0 {
			ctx.WithinList--
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
func (r *sgmlRenderer) renderElement(ctx *renderer.Context, element interface{}) ([]byte, error) {
	log.Debugf("rendering element of type `%T`", element)
	switch e := element.(type) {
	case []interface{}:
		return r.renderElements(ctx, e)
	case types.TableOfContentsPlaceHolder:
		return r.renderTableOfContents(ctx, ctx.TableOfContents)
	case types.Section:
		return r.renderSection(ctx, e)
	case types.Preamble:
		return r.renderPreamble(ctx, e)
	case types.BlankLine:
		return r.renderBlankLine(ctx, e)
	case types.LabeledList:
		return r.renderLabeledList(ctx, e)
	case types.OrderedList:
		return r.renderOrderedList(ctx, e)
	case types.UnorderedList:
		return r.renderUnorderedList(ctx, e)
	case types.CalloutList:
		return r.renderCalloutList(ctx, e)
	case types.Paragraph:
		return r.renderParagraph(ctx, e)
	case types.InternalCrossReference:
		return r.renderInternalCrossReference(ctx, e)
	case types.ExternalCrossReference:
		return r.renderExternalCrossReference(ctx, e)
	case types.QuotedText:
		return r.renderQuotedText(ctx, e)
	case types.InlinePassthrough:
		return r.renderInlinePassthrough(ctx, e)
	case types.ImageBlock:
		return r.renderImageBlock(ctx, e)
	case types.InlineImage:
		return r.renderInlineImage(e)
	case types.DelimitedBlock:
		return r.renderDelimitedBlock(ctx, e)
	case types.Table:
		return r.renderTable(ctx, e)
	case types.LiteralBlock:
		return r.renderLiteralBlock(ctx, e)
	case types.InlineLink:
		return r.renderLink(ctx, e)
	case types.StringElement:
		return r.renderStringElement(ctx, e)
	case types.FootnoteReference:
		return r.renderFootnoteReference(e)
	case types.LineBreak:
		return r.renderLineBreak()
	case types.UserMacro:
		return r.renderUserMacro(ctx, e)
	case types.IndexTerm:
		return r.renderIndexTerm(ctx, e)
	case types.ConcealedIndexTerm:
		return r.renderConcealedIndexTerm(e)
	case types.VerbatimLine:
		return r.renderVerbatimLine(e)
	default:
		return nil, errors.Errorf("unsupported type of element: %T", element)
	}
}

// nolint: gocyclo
func (r *sgmlRenderer) renderPlainText(ctx *renderer.Context, element interface{}) ([]byte, error) {
	log.Debugf("rendering plain string for element of type %T", element)
	switch element := element.(type) {
	case []interface{}:
		return r.renderInlineElements(ctx, element, r.withVerbatim())
	case [][]interface{}:
		return r.renderLines(ctx, element, r.withPlainText())
	case types.QuotedText:
		return r.renderPlainText(ctx, element.Elements)
	case types.InlineImage:
		return []byte(element.Attributes.GetAsStringWithDefault(types.AttrImageAlt, "")), nil
	case types.InlineLink:
		if alt, ok := element.Attributes[types.AttrInlineLinkText].([]interface{}); ok {
			return r.renderPlainText(ctx, alt)
		}
		return []byte(element.Location.String()), nil
	case types.BlankLine:
		return []byte("\n\n"), nil
	case types.StringElement:
		return []byte(element.Content), nil
	case types.Paragraph:
		return r.renderLines(ctx, element.Lines, r.withPlainText())
	case types.FootnoteReference:
		// footnotes are rendered in HTML so they can appear as such in the table of contents
		return r.renderFootnoteReferencePlainText(element)
	default:
		return nil, errors.Errorf("unable to render plain string for element of type '%T'", element)
	}
}

// includeNewline returns an "\n" sequence if the given index is NOT the last entry in the given description lines, empty string otherwise.
// also, it ignores the element if it is a blank line, depending on the context
func (r *sgmlRenderer) includeNewline(ctx renderer.Context, index int, content interface{}) string {
	switch reflect.TypeOf(content).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(content)
		if _, match := s.Index(index).Interface().(types.BlankLine); match {
			if ctx.IncludeBlankLine {
				return "\n" // TODO: parameterize this?
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

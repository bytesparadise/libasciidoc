package html5

import (
	"bytes"
	"io"
	"reflect"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx *renderer.Context, output io.Writer) (map[string]interface{}, error) {
	return renderDocument(ctx, output)
}

func renderElement(ctx *renderer.Context, element types.DocElement) ([]byte, error) {
	log.Debugf("rendering element of type `%T`", element)
	switch e := element.(type) {
	case *types.TableOfContentsMacro:
		return renderTableOfContent(ctx, e)
	case *types.Section:
		return renderSection(ctx, e)
	case *types.Preamble:
		return renderPreamble(ctx, e)
	case *types.LabeledList:
		return renderLabeledList(ctx, e)
	case *types.UnorderedList:
		return renderUnorderedList(ctx, e)
	case *types.Paragraph:
		return renderParagraph(ctx, e)
	case *types.ListParagraph:
		return renderListParagraph(ctx, e)
	case *types.CrossReference:
		return renderCrossReference(ctx, e)
	case *types.QuotedText:
		return renderQuotedText(ctx, e)
	case *types.Passthrough:
		return renderPassthrough(ctx, e)
	case *types.BlockImage:
		return renderBlockImage(ctx, e)
	case *types.InlineImage:
		return renderInlineImage(ctx, e)
	case *types.DelimitedBlock:
		return renderDelimitedBlock(ctx, e)
	case *types.LiteralBlock:
		return renderLiteralBlock(ctx, e)
	case *types.InlineContent:
		return renderInlineContent(ctx, e)
	case *types.ExternalLink:
		return renderExternalLink(ctx, e)
	case *types.StringElement:
		return renderStringElement(ctx, e)
	case *types.DocumentAttributeDeclaration:
		// 'process' function do not return any rendered content, but may return an error
		return nil, processAttributeDeclaration(ctx, e)
	case *types.DocumentAttributeReset:
		// 'process' function do not return any rendered content, but may return an error
		return nil, processAttributeReset(ctx, e)
	case *types.DocumentAttributeSubstitution:
		return renderAttributeSubstitution(ctx, e)
	default:
		log.Errorf("unsupported type of element: %T", element)
		return nil, errors.Errorf("unsupported type of element: %T", element)
	}
}

func renderPlainString(ctx *renderer.Context, element types.DocElement) ([]byte, error) {
	switch element := element.(type) {
	case *types.SectionTitle:
		return renderPlainStringForInlineElements(ctx, element.Content.Elements)
	case *types.QuotedText:
		return renderPlainStringForInlineElements(ctx, element.Elements)
	case *types.InlineImage:
		return []byte(element.Macro.Alt), nil
	case *types.ExternalLink:
		return []byte(element.Text), nil
	case *types.StringElement:
		return []byte(element.Content), nil
	default:
		return nil, errors.Errorf("unexpectedResult type of element to process: %T", element)
	}
}

func renderPlainStringForInlineElements(ctx *renderer.Context, elements []types.InlineElement) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	for _, e := range elements {
		plainStringElement, err := renderPlainString(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render plain string value")
		}
		buff.Write(plainStringElement)
	}
	return buff.Bytes(), nil
}

// notLastItem returns true if the given index is NOT the last entry in the given description lines, false otherwise.
func notLastItem(index int, content interface{}) bool {
	switch reflect.TypeOf(content).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(content)
		return index < s.Len()-1
	default:
		log.Warnf("content of type %T is not an array or a slice")
		return false
	}
}

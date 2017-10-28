package html5

import (
	"bytes"
	"io"

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
	log.Debugf("rendering element of type %T", element)
	switch element.(type) {
	case *types.Section:
		return renderSection(ctx, *element.(*types.Section))
	case *types.Preamble:
		return renderPreamble(ctx, *element.(*types.Preamble))
	case *types.List:
		return renderList(ctx, *element.(*types.List))
	case *types.Paragraph:
		return renderParagraph(ctx, *element.(*types.Paragraph))
	case *types.QuotedText:
		return renderQuotedText(ctx, *element.(*types.QuotedText))
	case *types.BlockImage:
		return renderBlockImage(ctx, *element.(*types.BlockImage))
	case *types.InlineImage:
		return renderInlineImage(ctx, *element.(*types.InlineImage))
	case *types.DelimitedBlock:
		return renderDelimitedBlock(ctx, *element.(*types.DelimitedBlock))
	case *types.LiteralBlock:
		return renderLiteralBlock(ctx, *element.(*types.LiteralBlock))
	case *types.InlineContent:
		return renderInlineContent(ctx, *element.(*types.InlineContent))
	case *types.StringElement:
		return renderStringElement(ctx, *element.(*types.StringElement))
	case *types.DocumentAttributeDeclaration:
		// 'process' function do not return any rendered content, but may return an error
		return nil, processAttributeDeclaration(ctx, element.(*types.DocumentAttributeDeclaration))
	case *types.DocumentAttributeReset:
		// 'process' function do not return any rendered content, but may return an error
		return nil, processAttributeReset(ctx, *element.(*types.DocumentAttributeReset))
	case *types.DocumentAttributeSubstitution:
		return renderAttributeSubstitution(ctx, *element.(*types.DocumentAttributeSubstitution))
	default:
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
		return nil, errors.Errorf("unexpected type of element to process: %T", element)
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

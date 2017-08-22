package html5

import (
	"context"
	"io"

	"reflect"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx context.Context, document types.Document, output io.Writer) error {
	for i, element := range document.Elements {
		content, err := renderElement(ctx, element)
		if err != nil {
			return errors.Wrapf(err, "failed to render document")
		}
		output.Write(content)
		if _, ok := element.(*types.BlankLine); !ok && i < len(document.Elements)-1 {
			output.Write([]byte("\n"))
		}
	}
	return nil
}

func renderElement(ctx context.Context, element types.DocElement) ([]byte, error) {
	switch element.(type) {
	case *types.Section:
		return renderSection(ctx, *element.(*types.Section))
	case *types.List:
		return renderList(ctx, *element.(*types.List))
	case *types.Paragraph:
		return renderParagraph(ctx, *element.(*types.Paragraph))
	case *types.QuotedText:
		return renderQuotedText(ctx, *element.(*types.QuotedText))
	case *types.BlockImage:
		return renderBlockImage(ctx, *element.(*types.BlockImage))
	case *types.DelimitedBlock:
		return renderDelimitedBlock(ctx, *element.(*types.DelimitedBlock))
	case *types.InlineContent:
		return renderInlineContent(ctx, *element.(*types.InlineContent))
	case *types.StringElement:
		return renderStringElement(ctx, *element.(*types.StringElement))
	case *types.BlankLine, *types.ElementID, *types.ElementLink, *types.ElementTitle:
		// ignored in the output
		return make([]byte, 0), nil
	default:
		return nil, errors.Errorf("unsupported type of element: %v", reflect.TypeOf(element))
	}

}

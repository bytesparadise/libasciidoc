package html5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Render renders the given document in HTML and writes the result in the given `writer`
func Render(ctx *renderer.Context, output io.Writer) error {
	if ctx.IncludeHeaderFooter() {
		return renderFullDocument(ctx, output)
	}
	return renderElements(ctx, output)

}

func renderElements(ctx *renderer.Context, output io.Writer) error {
	hasContent := false
	for _, element := range ctx.Document.Elements {
		content, err := renderElement(ctx, element)
		if err != nil {
			return errors.Wrapf(err, "failed to render the document")
		}
		// if there's already some content, we need to insert a `\n` before writing
		// the rendering output of the current element (if application, ie, not empty)
		if hasContent && len(content) > 0 {
			output.Write([]byte("\n"))
		}
		// if the element was rendering into 'something' (ie, not enpty result)
		if len(content) > 0 {
			output.Write(content)
			hasContent = true
		}
	}
	return nil
}

func renderElement(ctx *renderer.Context, element types.DocElement) ([]byte, error) {
	log.Debugf("Rendering element of type %T", element)
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

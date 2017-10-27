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
	metadata := make(map[string]interface{})
	var doctitle *string
	if ctx.IncludeHeaderFooter() {
		title, err := renderFullDocument(ctx, output)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		doctitle = title
	} else {
		err := renderElements(ctx, output)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render full document")
		}
		doctitle, err = RenderDocumentTitle(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error while rendering the HTML document")
		}
	}
	// copy all document attributes, and override the title with its rendered value instead
	for k, v := range ctx.Document.Attributes {
		switch k {
		case "doctitle":
			metadata[k] = *doctitle
		default:
			metadata[k] = v
		}
	}
	return metadata, nil

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

func renderPlainString(ctx *renderer.Context, element types.DocElement) (*string, error) {
	switch element := element.(type) {
	case *types.SectionTitle:
		return renderPlainStringForInlineElements(ctx, element.Content.Elements)
	case *types.QuotedText:
		return renderPlainStringForInlineElements(ctx, element.Elements)
	case *types.InlineImage:
		return &element.Macro.Alt, nil
	case *types.ExternalLink:
		return &element.Text, nil
	case *types.StringElement:
		return &element.Content, nil
	default:
		return nil, errors.Errorf("unexpected type of element to process: %T", element)
	}
}

func renderPlainStringForInlineElements(ctx *renderer.Context, elements []types.InlineElement) (*string, error) {
	buff := bytes.NewBuffer(nil)
	for i, e := range elements {
		plainStringElement, err := renderPlainString(ctx, e)
		if err != nil {
			return nil, errors.Wrap(err, "unable to render plain string value")
		}
		buff.WriteString(*plainStringElement)
		if i < len(elements)-1 {
			buff.WriteString(" ")
		}
	}
	result := buff.String()
	return &result, nil
}

package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderElements(ctx *context, elements []interface{}) (string, error) {
	// log.Debugf("rendering %d elements(s)...", len(elements))
	buff := &strings.Builder{}
	for _, element := range elements {
		renderedElement, err := r.renderElement(ctx, element)
		if err != nil {
			return "", err // no need to wrap the error here
		}
		buff.WriteString(renderedElement)
	}
	return buff.String(), nil
}

// renderListElements is similar to the `renderElements` func above,
// but it sets the `withinList` context flag to true for the first element only
func (r *sgmlRenderer) renderListElements(ctx *context, elements []interface{}) (string, error) {
	// log.Debugf("rendering list with %d element(s)...", len(elements))
	buff := &strings.Builder{}
	for i, element := range elements {
		if i == 0 {
			ctx.withinList++
		}
		renderedElement, err := r.renderElement(ctx, element)
		if err != nil {
			return "", errors.Wrap(err, "unable to render a list element")
		}
		if i == 0 {
			ctx.withinList--
		}
		buff.WriteString(renderedElement)
	}
	return buff.String(), nil
}

//nolint:gocyclo
func (r *sgmlRenderer) renderElement(ctx *context, element interface{}) (string, error) {
	// log.Debugf("rendering element of type `%T`", element)
	switch e := element.(type) {
	case *types.TableOfContents:
		return r.renderTableOfContents(ctx, e)
	case *types.Section:
		return r.renderSection(ctx, e)
	case *types.Preamble:
		return r.renderPreamble(ctx, e)
	case *types.List:
		return r.renderList(ctx, e)
	case *types.Callout:
		return r.renderCalloutRef(e)
	case *types.Paragraph:
		return r.renderParagraph(ctx, e)
	case *types.InternalCrossReference:
		return r.renderInternalCrossReference(ctx, e)
	case *types.ExternalCrossReference:
		return r.renderExternalCrossReference(ctx, e)
	case *types.QuotedText:
		return r.renderQuotedText(ctx, e)
	case *types.InlinePassthrough:
		return r.renderInlinePassthrough(ctx, e)
	case *types.ImageBlock:
		return r.renderImageBlock(ctx, e)
	case *types.InlineButton:
		return r.renderInlineButton(e)
	case *types.InlineImage:
		return r.renderInlineImage(ctx, e)
	case *types.InlineMenu:
		return r.renderInlineMenu(e)
	case *types.Icon:
		return r.renderInlineIcon(ctx, e)
	case *types.DelimitedBlock:
		return r.renderDelimitedBlock(ctx, e)
	case *types.Table:
		return r.renderTable(ctx, e)
	case *types.InlineLink:
		return r.renderLink(ctx, e)
	case *types.StringElement:
		return r.renderStringElement(ctx, e)
	case *types.FootnoteReference:
		return r.renderFootnoteReference(e)
	case *types.LineBreak:
		return r.renderLineBreak()
	case *types.UserMacro:
		return r.renderUserMacro(ctx, e)
	case *types.IndexTerm:
		return r.renderIndexTerm(ctx, e)
	case *types.ConcealedIndexTerm:
		return r.renderConcealedIndexTerm(e)
	case *types.ThematicBreak:
		return r.renderThematicBreak()
	case *types.SpecialCharacter:
		return r.renderSpecialCharacter(e)
	case *types.Symbol:
		return r.renderSymbol(e)
	case *types.PredefinedAttribute:
		return r.renderPredefinedAttribute(e)
	case *types.AttributeDeclaration:
		ctx.attributes[e.Name] = e.Value
		return "", nil
	case *types.AttributeReset:
		delete(ctx.attributes, e.Name)
		return "", nil
	case *types.FrontMatter:
		ctx.attributes.AddAll(e.Attributes)
		return "", nil
	default:
		return "", errors.Errorf("unsupported type of element: %T", element)
	}
}

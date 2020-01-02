package parser

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// applyDocumentAttributeSubstitutions(elements applies the document attribute substitutions
// and re-parse the paragraphs that were affected
func applyDocumentAttributeSubstitutions(element interface{}, attrs types.DocumentAttributes) (interface{}, error) {
	// the document attributes, as they are resolved while processing the blocks
	log.Debugf("applying document substitutions on block of type %T", element)
	switch e := element.(type) {
	case []interface{}:
		elements := make([]interface{}, 0, len(e)) // maximum capacity cannot exceed initial input
		for _, element := range e {
			r, err := applyDocumentAttributeSubstitutions(element, attrs)
			if err != nil {
				return []interface{}{}, err
			}
			elements = append(elements, r)
		}
		// elements = filter(elements, DocumentAttributeMatcher)
		return parseInlineLinks(types.MergeStringElements(elements))
	case types.DocumentAttributeDeclaration:
		attrs[e.Name] = e.Value
		return e, nil
	case types.DocumentAttributeReset:
		delete(attrs, e.Name)
		return e, nil
	case types.DocumentAttributeSubstitution:
		if value, ok := attrs[e.Name].(string); ok {
			return types.StringElement{
				Content: value,
			}, nil
		}
		return types.StringElement{
			Content: "{" + e.Name + "}",
		}, nil
	case types.ImageBlock:
		return e.ResolveLocation(attrs), nil
	case types.InlineImage:
		return e.ResolveLocation(attrs), nil
	case types.Section:
		title, err := applyDocumentAttributeSubstitutions(e.Title, attrs)
		if err != nil {
			return struct{}{}, err
		}
		if title, ok := title.([]interface{}); ok {
			e.Title = title
		}
		return e.ResolveID(attrs)
	case types.OrderedListItem:
		elements, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Elements = elements.([]interface{})
		return e, nil
	case types.UnorderedListItem:
		elements, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Elements = elements.([]interface{})
		return e, nil
	case types.LabeledListItem:
		elements, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Elements = elements.([]interface{})
		return e, nil
	case types.QuotedText:
		elements, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Elements = elements.([]interface{})
		return e, nil
	case types.ContinuedListItemElement:
		element, err := applyDocumentAttributeSubstitutions(e.Element, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Element = element
		return e, nil
	case types.DelimitedBlock:
		elements, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, err
		}
		e.Elements = elements.([]interface{})
		return e, nil
	case types.Paragraph:
		for i, line := range e.Lines {
			line, err := applyDocumentAttributeSubstitutions(line, attrs)
			if err != nil {
				return struct{}{}, err
			}
			e.Lines[i] = line.([]interface{})
		}
		return e, nil
	default:
		return e, nil
	}
}

// if a document attribute substitution happened, we need to parse the string element in search
// for a potentially new link. Eg `{url}` giving `https://foo.com`
func parseInlineLinks(elements []interface{}) ([]interface{}, error) {
	result := []interface{}{}
	for _, element := range elements {
		switch element := element.(type) {
		case types.StringElement:
			log.Debugf("looking for links in line element of type %[1]T (%[1]v)", element)
			elements, err := ParseReader("", strings.NewReader(element.Content), Entrypoint("InlineLinks"))
			if err != nil {
				return []interface{}{}, errors.Wrap(err, "error while parsing content for inline links")
			}
			log.Debugf("  found: %+v", elements)
			result = append(result, elements.([]interface{})...)
		default:
			result = append(result, element)
		}
	}
	return result, nil
}

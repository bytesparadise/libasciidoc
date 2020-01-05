package parser

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// applyDocumentAttributeSubstitutions(elements applies the document attribute substitutions
// and re-parse the paragraphs that were affected
// nolint: gocyclo
func applyDocumentAttributeSubstitutions(element interface{}, attrs types.DocumentAttributes) (interface{}, bool, error) {
	// the document attributes, as they are resolved while processing the blocks
	log.Debugf("applying document substitutions on block of type %T", element)
	switch e := element.(type) {
	case []interface{}:
		elements := make([]interface{}, 0, len(e)) // maximum capacity cannot exceed initial input
		applied := false
		for _, element := range e {
			r, a, err := applyDocumentAttributeSubstitutions(element, attrs)
			if err != nil {
				return []interface{}{}, false, err
			}
			elements = append(elements, r)
			applied = applied || a
		}
		elements = types.MergeStringElements(elements)
		if applied {
			elements, err := parseInlineLinks(elements)
			return elements, true, err
		}
		return elements, false, nil
	case types.DocumentAttributeDeclaration:
		attrs[e.Name] = e.Value
		return e, false, nil
	case types.DocumentAttributeReset:
		delete(attrs, e.Name)
		return e, false, nil
	case types.DocumentAttributeSubstitution:
		if value, ok := attrs[e.Name].(string); ok {
			return types.StringElement{
				Content: value,
			}, true, nil
		}
		return types.StringElement{
			Content: "{" + e.Name + "}",
		}, false, nil
	case types.ImageBlock:
		return e.ResolveLocation(attrs), false, nil
	case types.InlineImage:
		return e.ResolveLocation(attrs), false, nil
	case types.ExternalCrossReference:
		return e.ResolveLocation(attrs), false, nil
	case types.Section:
		title, applied, err := applyDocumentAttributeSubstitutions(e.Title, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		if title, ok := title.([]interface{}); ok {
			e.Title = title
		}
		e, err = e.ResolveID(attrs)
		return e, applied, err
	case types.OrderedListItem:
		elements, applied, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Elements = elements.([]interface{})
		return e, applied, nil
	case types.UnorderedListItem:
		elements, applied, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Elements = elements.([]interface{})
		return e, applied, nil
	case types.LabeledListItem:
		elements, applied, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Elements = elements.([]interface{})
		return e, applied, nil
	case types.QuotedText:
		elements, applied, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Elements = elements.([]interface{})
		return e, applied, nil
	case types.ContinuedListItemElement:
		element, applied, err := applyDocumentAttributeSubstitutions(e.Element, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Element = element
		return e, applied, nil
	case types.DelimitedBlock:
		elements, applied, err := applyDocumentAttributeSubstitutions(e.Elements, attrs)
		if err != nil {
			return struct{}{}, false, err
		}
		e.Elements = elements.([]interface{})
		return e, applied, nil
	case types.Paragraph:
		applied := false
		for i, line := range e.Lines {
			line, a, err := applyDocumentAttributeSubstitutions(line, attrs)
			if err != nil {
				return struct{}{}, false, err
			}
			e.Lines[i] = line.([]interface{})
			applied = applied || a
		}
		return e, applied, nil
	default:
		return e, false, nil
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

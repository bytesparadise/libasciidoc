package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocument parses the content of the reader identitied by the filename
func ParseDocument(filename string, r io.Reader, opts ...Option) (types.Document, error) {
	draftDoc, err := ParseDraftDocument(filename, r, opts...)
	if err != nil {
		return types.Document{}, err
	}
	attrs := types.DocumentAttributes{}
	// add all predefined attributes
	for k, v := range Predefined {
		if v, ok := v.(string); ok {
			attrs[k] = v
		}
	}

	// also, add all front-matter key/values
	for k, v := range draftDoc.FrontMatter.Content {
		if v, ok := v.(string); ok {
			attrs[k] = v
		}
	}

	// also, add all DocumentAttributeDeclaration at the top of the document
	documentAttributes := draftDoc.DocumentAttributes()
	for k, v := range documentAttributes {
		attrs[k] = v
	}

	// apply document attribute substitutions and re-parse paragraphs that were affected
	blocks, _, err := applyDocumentAttributeSubstitutions(draftDoc.Blocks, attrs)
	if err != nil {
		return types.Document{}, err
	}

	// now, merge list items into proper lists
	blocks, err = rearrangeListItems(blocks.([]interface{}), false)
	if err != nil {
		return types.Document{}, err
	}
	// apply document attribute substitutions and re-parse paragraphs that were affected
	blocks = filter(blocks.([]interface{}), allMatchers...)

	// now, rearrange elements in a hierarchical manner
	doc, err := rearrangeSections(blocks.([]interface{}))
	if err != nil {
		return types.Document{}, err
	}
	// now, add front-matter attributes
	for k, v := range draftDoc.FrontMatter.Content {
		doc.Attributes[k] = v
	}
	// and add all remaining attributes, too
	for k, v := range documentAttributes {
		doc.Attributes[k] = v
	}
	return doc, nil
}

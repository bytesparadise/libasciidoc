package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocument parses the content of the reader identitied by the filename
func ParseDocument(r io.Reader, config configuration.Configuration) (types.Document, error) {
	draftDoc, err := ParseDraftDocument(r, config)
	if err != nil {
		return types.Document{}, err
	}
	attrs := types.DocumentAttributesWithOverrides{
		Content:   types.DocumentAttributes{},
		Overrides: config.AttributeOverrides,
	}
	// add all predefined attributes
	attrs.AddAll(Predefined)
	// also, add all front-matter key/values
	attrs.AddAll(draftDoc.FrontMatter.Content)
	// also, add all DocumentAttributeDeclaration at the top of the document
	attrs.AddAll(draftDoc.DocumentAttributes())

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
	// filter out blocks not needed in the final doc
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
	doc.Attributes.AddAll(draftDoc.DocumentAttributes())
	return doc, nil
}

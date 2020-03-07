package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"

	log "github.com/sirupsen/logrus"
)

// ContextKey a non-built-in type for keys in the context
type ContextKey string

// LevelOffset the key for the level offset of the file to include
const LevelOffset ContextKey = "leveloffset"

// ParseDraftDocument parses a document's content and applies the preprocessing directives (file inclusions)
func ParseDraftDocument(r io.Reader, config configuration.Configuration, options ...Option) (types.DraftDocument, error) {
	options = append(options, Entrypoint("AsciidocDocument"))
	return parseDraftDocument(r, []levelOffset{}, config, options...)
}

func parseDraftDocument(r io.Reader, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) (types.DraftDocument, error) {
	d, err := ParseReader(config.Filename, r, options...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	doc := d.(types.DraftDocument)
	attrs := types.DocumentAttributesWithOverrides{
		Content:   map[string]interface{}{},
		Overrides: map[string]string{},
	}
	blocks, err := parseElements(doc.Blocks, attrs, levelOffsets, config, options...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	doc.Blocks = blocks
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("draft document:")
		spew.Dump(doc)
	}
	return doc, nil
}

// parseElements resolves the file inclusions if any is found in the given elements
func parseElements(elements []interface{}, attrs types.DocumentAttributesWithOverrides, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) ([]interface{}, error) {
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case types.DocumentAttributeDeclaration:
			attrs.Add(e.Name, e.Value)
			result = append(result, e)
		case types.FileInclusion:
			// read the file and include its content
			embedded, err := parseFileToInclude(e, attrs, levelOffsets, config, options...)
			if err != nil {
				// do not fail, but instead report the error in the console
				log.Errorf("failed to include file '%s': %v", e.Location, err)
			}
			result = append(result, embedded.Blocks...)
		case types.DelimitedBlock:
			elmts, err := parseElements(e.Elements, attrs, levelOffsets, config,
				// use a new var to avoid overridding the current one which needs to stay as-is for the rest of the doc parsing
				append(options, Entrypoint("AsciidocDocumentWithinDelimitedBlock"))...)
			if err != nil {
				return nil, err
			}
			result = append(result, types.DelimitedBlock{
				Attributes: e.Attributes,
				Kind:       e.Kind,
				Elements:   elmts,
			})
		case types.Section:
			for _, offset := range levelOffsets {
				oldLevel := e.Level
				offset.apply(&e)
				// replace the absolute when the first section is processed with a relative offset
				// which is based on the actual level offset that resulted in the application of the absolute offset
				if offset.absolute {
					levelOffsets = []levelOffset{
						relativeOffset(e.Level - oldLevel),
					}
				}
			}
			result = append(result, e)
		default:
			result = append(result, e)
		}
	}
	return result, nil
}

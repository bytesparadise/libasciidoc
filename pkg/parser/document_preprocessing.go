package parser

import (
	"io"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ContextKey a non-built-in type for keys in the context
type ContextKey string

// LevelOffset the key for the level offset of the file to include
const LevelOffset ContextKey = "leveloffset"

// ParseDraftDocument parses a document's content and applies the preprocessing directives (file inclusions)
func ParseDraftDocument(filename string, r io.Reader, opts ...Option) (types.DraftDocument, error) {
	opts = append(opts, Entrypoint("DraftAsciidocDocument"))
	return parseDraftDocument(filename, r, "", opts...)
}

func parseDraftDocument(filename string, r io.Reader, levelOffset string, opts ...Option) (types.DraftDocument, error) {
	d, err := ParseReader(filename, r, opts...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	doc := d.(types.DraftDocument)
	attrs := types.DocumentAttributes{}
	blocks, err := parseElements(filename, doc.Blocks, attrs, levelOffset, opts...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	doc.Blocks = blocks
	return doc, nil
}

// parseElements resolves the file inclusions if any is found in the given elements
func parseElements(filename string, elements []interface{}, attrs types.DocumentAttributes, levelOffset string, opts ...Option) ([]interface{}, error) {
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case types.DocumentAttributeDeclaration:
			attrs[e.Name] = e.Value
			result = append(result, e)
		case types.FileInclusion:
			// read the file and include its content
			embedded, err := parseFileToInclude(filename, e, attrs, opts...)
			if err != nil {
				// do not fail, but instead report the error in the console
				log.Errorf("failed to include file '%s': %v", e.Location, err)
			}
			result = append(result, embedded.Blocks...)
		case types.DelimitedBlock:
			elmts, err := parseElements(filename, e.Elements, attrs, levelOffset,
				// use a new var to avoid overridding the current one which needs to stay as-is for the rest of the doc parsing
				append(opts, Entrypoint("DraftAsciidocDocumentWithinDelimitedBlock"))...)
			if err != nil {
				return nil, err
			}
			result = append(result, types.DelimitedBlock{
				Attributes: e.Attributes,
				Kind:       e.Kind,
				Elements:   elmts,
			})
		case types.Section:
			if levelOffset != "" {
				log.Debugf("applying level offset '%s'", levelOffset)
				offset, err := strconv.Atoi(levelOffset)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to preparse '%s'", filename)
				}
				e.Level += offset
			}
			result = append(result, e)
		default:
			result = append(result, e)
		}
	}
	return result, nil
}

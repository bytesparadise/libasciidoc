package parser

import (
	"bytes"
	"fmt"
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
	log.Debugf("parsing draft document '%s'", config.Filename)
	d, err := ParseReader(config.Filename, r, options...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	doc := d.(types.DraftDocument)
	attrs := types.DocumentAttributesWithOverrides{
		Content:   map[string]interface{}{},
		Overrides: map[string]string{},
	}
	doc.Blocks, err = processFileInclusions(doc.Blocks, attrs, levelOffsets, config, options...)
	if err != nil {
		return types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: err.Error(),
							},
						},
					},
				},
			},
		}, nil
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("draft document:")
		spew.Dump(doc)
	}
	return doc, nil
}

// processFileInclusions resolves the file inclusions if any is found in the given elements
// and applies level offset on sections when needed
func processFileInclusions(elements []interface{}, attrs types.DocumentAttributesWithOverrides, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) ([]interface{}, error) {
	result := []interface{}{}
	log.Debugf("processing file inclusions found in %d element(s)", len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case types.DocumentAttributeDeclaration: // may be needed if there's an attribute substitution in the path of the file to include
			attrs.Add(e.Name, e.Value)
			result = append(result, e)
		case types.FileInclusion:
			// read the file and include its content
			embedded, err := parseFileToInclude(e, attrs, levelOffsets, config, options...)
			if errr, ok := err.(FileInclusionError); ok {
				log.Errorf("failed to include content of '%s' in '%s'", e.Location, errr.Filename)
				return nil, err
			} else if err != nil {
				return nil, err
			}
			result = append(result, embedded.Blocks...)
		case types.DelimitedBlock:
			elmts, err := processFileInclusions(e.Elements, attrs, levelOffsets, config,
				// use a new var to avoid overridding the current one which needs to stay as-is for the rest of the doc parsing
				append(options, Entrypoint("VerbatimDocument"))...)
			if err != nil {
				// do not fail but retain the error message
				elmts = []interface{}{
					types.VerbatimLine{
						Content: err.Error(),
					},
				}
			}
			// next, parse the elements with the grammar rule that corresponds to the delimited block substitutions (based on its type)
			extraAttrs, elmts, err := parseDelimitedBlockContent(config.Filename, e.Kind, elmts, options...)
			if err != nil {
				return nil, err
			}
			e.Attributes.AddAll(extraAttrs)
			result = append(result, types.DelimitedBlock{
				Attributes: e.Attributes,
				Kind:       e.Kind,
				Elements:   types.NilSafe(elmts),
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

// parseDelimitedBlockContent parses the given verbatim elements, depending on the given delimited block kind.
// May return the elements unchanged, or convert the elements to a source doc and parse with a custom entrypoint
func parseDelimitedBlockContent(filename string, kind types.BlockKind, elements []interface{}, options ...Option) (types.ElementAttributes, []interface{}, error) {
	switch kind {
	case types.Fenced, types.Listing, types.Literal, types.Source, types.Comment:
		// return the verbatim elements
		return types.ElementAttributes{}, elements, nil
	case types.Example, types.Quote, types.Sidebar:
		return parseDelimitedBlockElements(filename, elements, append(options, Entrypoint("NormalBlockContent"))...)
	case types.MarkdownQuote:
		return parseMarkdownQuoteBlockElements(filename, elements, append(options, Entrypoint("NormalBlockContent"))...)
	case types.Verse:
		return parseDelimitedBlockElements(filename, elements, append(options, Entrypoint("VerseBlockContent"))...)
	default:
		return nil, nil, fmt.Errorf("unexpected kind of delimited block: '%s'", kind)
	}
}

func parseDelimitedBlockElements(filename string, elements []interface{}, options ...Option) (types.ElementAttributes, []interface{}, error) {
	verbatim, err := serialize(elements)
	if err != nil {
		return nil, nil, err
	}
	e, err := ParseReader(filename, verbatim, options...)
	if err != nil {
		return nil, nil, err
	}
	if result, ok := e.([]interface{}); ok {
		return types.ElementAttributes{}, result, nil
	}
	return nil, nil, fmt.Errorf("unexpected type of element after parsing the content of a delimited block: '%T'", e)
}

func parseMarkdownQuoteBlockElements(filename string, elements []interface{}, options ...Option) (types.ElementAttributes, []interface{}, error) {
	author := parseMarkdownQuoteBlockAttribution(filename, elements)
	if author != "" {
		elements = elements[:len(elements)-1]
	}
	attrs, lines, err := parseDelimitedBlockElements(filename, elements, options...)
	attrs.AddNonEmpty(types.AttrQuoteAuthor, author)
	return attrs, lines, err
}

func parseMarkdownQuoteBlockAttribution(filename string, elements []interface{}) string {
	// first, check if last line is an attribution (author)
	if lastLine, ok := elements[len(elements)-1].(types.VerbatimLine); ok {
		buf := bytes.NewBuffer(nil)
		buf.WriteString(lastLine.Content)
		a, err := ParseReader(filename, buf, Entrypoint("MarkdownQuoteBlockAttribution"))
		// assume that the last line is not an author attribution if an error occurred
		if err != nil {
			return ""
		}
		if a, ok := a.(string); ok {
			return a
		}
	}
	return ""
}

func serialize(elements []interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	for _, e := range elements {
		if r, ok := e.(types.VerbatimLine); ok {
			if _, err := buf.WriteString(r.Content); err != nil {
				return nil, err
			}
			if _, err := buf.WriteString("\n"); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unexpected type of element while marshalling the content of a delimited block: '%T'", e)
		}
	}
	log.Debugf("verbatim content: '%s'", buf.String())
	return buf, nil
}

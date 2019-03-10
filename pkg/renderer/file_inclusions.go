package renderer

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ProcessFileInclusions inspects the DOM and replaces all `FileInclusions`
func ProcessFileInclusions(doc types.Document) (types.Document, error) {
	var err error
	doc.Elements, err = processFileInclusions(doc.Elements, false)
	return doc, err
}

// ProcessFileInclusions inspects the DOM and replaces all `FileInclusions`
func processFileInclusions(elements []interface{}, withinDelimitedBlock bool) ([]interface{}, error) {
	result := []interface{}{}
	for _, element := range elements {
		switch e := element.(type) {
		case types.Section:
			elements, err := processFileInclusions(e.Elements, false)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to process file inclusions")
			}
			e.Elements = elements
			result = append(result, e)
		case types.DelimitedBlock:
			elements, err := processFileInclusions(e.Elements, true)
			if err != nil {
				return nil, errors.Wrapf(err, "fail to process file inclusions")
			}
			e.Elements = elements
			result = append(result, e)
		case types.FileInclusion:
			docElements, err := getElementsToInclude(e, withinDelimitedBlock)
			if err != nil {
				return result, errors.Wrapf(err, "fail to process file inclusions")
			}
			result = append(result, docElements...)
		default:
			result = append(result, e)
		}
	}
	return result, nil
}

func getElementsToInclude(file types.FileInclusion, withinDelimitedBlock bool) ([]interface{}, error) {
	if withinDelimitedBlock {
		return getRawLinesToInclude(file)
	}
	// parse the file if it is an Asciidoc only (.asciidoc, .adoc, .ad, .asc, or .txt)
	if ext := filepath.Ext(file.Path); ext == ".asciidoc" || ext == ".adoc" || ext == ".ad" || ext == ".asc" || ext == ".txt" {
		doc, err := parser.ParseFile(file.Path)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse file to include")
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debug("document to include:")
			spew.Dump(doc)
		}
		doc, err = ApplyLevelOffset(doc.(types.Document), file.Attributes.GetAsString(types.AttrLevelOffset))
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse file to include")
		}
		return doc.(types.Document).Elements, nil
	}
	return getRawLinesToInclude(file)
}

func getRawLinesToInclude(file types.FileInclusion) ([]interface{}, error) {
	// otherwise, read the files and wrap the lines in a paragraph
	// f, err := os.Open(file.Path)
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "unable to read file to include")
	// }
	// p := types.Paragraph{
	// 	Attributes: types.ElementAttributes{},
	// 	Lines:      []types.InlineElements{},
	// }
	// defer f.Close()
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	p.Lines = append(p.Lines, types.InlineElements{
	// 		types.StringElement{
	// 			Content: scanner.Text(),
	// 		},
	// 	})
	// }
	// otherwise, parse the rendered line, in case some new elements (links, etc.) "appeared" after document attribute substitutions

	f, err := os.Open(file.Path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file to include")
	}
	r, err := parser.ParseReader("", bufio.NewReader(f),
		parser.Entrypoint("DocumentToInclude"))
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse file to include")
	}
	elements, ok := r.([]interface{})
	if !ok {
		return nil, errors.Wrap(err, "failed to include file")
	}
	return elements, nil
}

// ApplyLevelOffset returns a document in which all section levels have been offset
func ApplyLevelOffset(doc types.Document, levelOffset string) (types.Document, error) {
	elements := doc.Elements
	// before returning the the doc elements, we need to check if there's a 'section 0', i.e., a title
	if title, ok := doc.Attributes[types.AttrTitle]; ok {
		elements = []interface{}{
			types.Section{
				Level:    0,
				Title:    title.(types.SectionTitle),
				Elements: doc.Elements,
			},
		}
	}
	if levelOffset != "" {
		offset, err := strconv.Atoi(levelOffset)
		if err != nil {
			return types.Document{}, errors.Wrapf(err, "fail to apply level offset '%s' to document to include", levelOffset)
		}
		// traverse the document and apply the offset on all sections
		elements = doApplyLevelOffset(elements, offset)
	}
	doc.Elements = elements
	return doc, nil
}

func doApplyLevelOffset(elements []interface{}, offset int) []interface{} {
	result := make([]interface{}, len(elements))
	for i, element := range elements {
		switch e := element.(type) {
		case types.Section:
			e.Level = e.Level + offset // TODO: need to support "absolute offset" as well
			// recursive call on child elements of this section
			e.Elements = doApplyLevelOffset(e.Elements, offset)
			result[i] = e
		default:
			result[i] = element
		}
	}
	return result
}

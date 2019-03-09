package renderer

import (
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
	doc.Elements, err = processFileInclusions(doc.Elements)
	return doc, err
}

// ProcessFileInclusions inspects the DOM and replaces all `FileInclusions`
func processFileInclusions(elements []interface{}) ([]interface{}, error) {
	result := []interface{}{}
	for _, element := range elements {
		switch e := element.(type) {
		case types.Section:
			var err error
			e.Elements, err = processFileInclusions(e.Elements)
			if err != nil {
				return result, errors.Wrapf(err, "fail to process file inclusions")
			}
			result = append(result, e)
		case types.FileInclusion:
			docElements, err := getElementsToInclude(e)
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

func getElementsToInclude(file types.FileInclusion) ([]interface{}, error) {
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

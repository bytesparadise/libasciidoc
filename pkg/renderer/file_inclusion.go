package renderer

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ProcessFileInclusions inspects the DOM and replaces all `FileInclusions`
func ProcessFileInclusions(ctx *Context) error {
	var err error
	ctx.Document.Elements, err = processFileInclusions(ctx.Document.Elements, false)
	return err
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
			docElements, err := parseFileToInclude(e, withinDelimitedBlock)
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

func parseFileToInclude(file types.FileInclusion, withinDelimitedBlock bool) ([]interface{}, error) {
	if withinDelimitedBlock || !file.IsAsciidoc() {
		return parseNestedFileToInclude(file)
	}
	log.Debugf("parsing content from '%s'", file.Path)
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

func parseNestedFileToInclude(file types.FileInclusion) ([]interface{}, error) {
	log.Debug("including raw lines...")
	f, err := os.Open(file.Path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file to include")
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	buf := new(bytes.Buffer)
	if lr, exists := file.Attributes[types.AttrLineRanges]; exists {
		if lineRanges, ok := lr.(types.LineRanges); ok { // could be a string if the input was invalid
			scanner := bufio.NewScanner(bufio.NewReader(f))
			line := 1
			for scanner.Scan() {
				log.Debugf("line %d: '%s'", line, scanner.Text())
				// TODO: stop reading if current line above highest range
				if lineRanges.Match(line) {
					_, err := buf.WriteString(scanner.Text())
					if err != nil {
						return nil, err
					}
					_, err = buf.WriteString("\n")
					if err != nil {
						return nil, err
					}
					log.Debugf("wrote line %d in buffer: '%s'", line, buf.String())
				}
				line++
			}
			log.Debugf("document lines to include: \n%s", buf.String())
		}
	} else {
		// just read all the doc
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}

	r, err := parser.Parse(file.Path, buf.Bytes(),
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

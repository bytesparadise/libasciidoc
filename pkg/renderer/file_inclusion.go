package renderer

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

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
	log.Debugf("processing file inclusions in %d element(s)", len(elements))
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
			elements, err := parseFileToInclude(e, withinDelimitedBlock)
			if err != nil {
				return result, errors.Wrapf(err, "fail to process file inclusions")
			}
			result = append(result, elements...)
		default:
			result = append(result, e)
		}
	}
	return result, nil
}

func parseFileToInclude(file types.FileInclusion, withinDelimitedBlock bool) ([]interface{}, error) {
	log.Debugf("parsing file '%s' to include...", file.Path)
	options := []parser.Option{}
	if withinDelimitedBlock {
		options = append(options, parser.Entrypoint("VerbatimBlock"))
	}
	content, restoreWD, err := readFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file to include: %s", file.Path)
	}
	defer restoreWD()
	c, err := parser.Parse(file.Path, content, options...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse file to include")
	}
	doc, err := ApplyLevelOffset(c, file.Attributes.GetAsString(types.AttrLevelOffset))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse file to include")
	}
	// recursively process the elements returned by the parsing of the file to include
	return processFileInclusions(doc.Elements, withinDelimitedBlock)
}

// readFile returns the content of the file, taking into account the optional ranges to limit the content
func readFile(file types.FileInclusion) ([]byte, func(), error) {
	log.Debugf("reading '%s'...", file.Path)
	// manage new working directory based on the file's location
	// so that if this file also includes other files with relative path,
	// then the it can work ;)
	wd, err := os.Getwd()
	if err != nil {
		return nil, func() {}, err
	}
	absPath, err := filepath.Abs(file.Path)
	if err != nil {
		return nil, func() {}, err
	}
	dir := filepath.Dir(absPath)
	err = os.Chdir(dir)
	if err != nil {
		return nil, func() {}, err
	}
	restoreWDFunc := func() {
		err = os.Chdir(wd) // restore the previous working directory
		if err != nil {
			log.WithError(err).Error("failed to restore previous working directory")
		}
	}
	// read the file per-se
	f, err := os.Open(absPath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unable to read file to include")
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	if lr, exists := file.Attributes[types.AttrLineRanges]; exists {
		buf := new(bytes.Buffer)
		if lineRanges, ok := lr.(types.LineRanges); ok { // could be a string if the input was invalid
			scanner := bufio.NewScanner(bufio.NewReader(f))
			line := 1
			for scanner.Scan() {
				log.Debugf("line %d: '%s'", line, scanner.Text())
				// TODO: stop reading if current line above highest range
				if lineRanges.Match(line) {
					_, err := buf.WriteString(scanner.Text())
					if err != nil {
						return nil, nil, errors.Wrap(err, "unable to parse file to include")
					}
					_, err = buf.WriteString("\n")
					if err != nil {
						return nil, nil, errors.Wrap(err, "unable to parse file to include")
					}
				}
				line++
			}
		}
		return buf.Bytes(), restoreWDFunc, nil
	}
	// just read all the doc
	content, err := ioutil.ReadAll(f)
	return content, restoreWDFunc, err
}

// ApplyLevelOffset returns a document in which all section levels have been offset
func ApplyLevelOffset(c interface{}, levelOffset string) (types.Document, error) {
	var doc types.Document
	switch c := c.(type) {
	case types.Document:
		doc = c
	case []interface{}:
		doc = types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements:           c,
		}
	default:
		return types.Document{}, errors.Errorf("fail to apply level offset '%s' to document to include: unexpected type of content: %T", levelOffset, c)
	}
	// before returning the the doc elements, we need to check if there's a 'section 0', i.e., a title
	if title, ok := doc.Attributes[types.AttrTitle]; ok {
		doc.Elements = []interface{}{
			types.Section{
				Level:      0,
				Title:      title.(types.SectionTitle),
				Attributes: types.ElementAttributes{},
				Elements:   doc.Elements,
			},
		}
	}
	if levelOffset != "" {
		offset, err := strconv.Atoi(levelOffset)
		if err != nil {
			return types.Document{}, errors.Wrapf(err, "fail to apply level offset '%s' to document to include", levelOffset)
		}
		// traverse the document and apply the offset on all sections
		doc.Elements = doApplyLevelOffset(doc.Elements, offset)
	}
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

package parser

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ContextKey a non-built-in type for keys in the context
type ContextKey string

// LevelOffset the key for the level offset of the file to include
const LevelOffset ContextKey = "leveloffset"

// ParsePreflightDocument parses a document's content and applies the preprocessing directives (file inclusions)
func ParsePreflightDocument(filename string, r io.Reader, opts ...Option) (types.PreflightDocument, error) {
	opts = append(opts, Entrypoint("PreflightDocument"))
	return parsePreflightDocument(filename, r, "", opts...)
}

func parsePreflightDocument(filename string, r io.Reader, levelOffset string, opts ...Option) (types.PreflightDocument, error) {
	d, err := ParseReader(filename, r, opts...)
	if err != nil {
		return types.PreflightDocument{}, err
	}
	doc := d.(types.PreflightDocument)
	attrs := types.DocumentAttributes{}
	blocks, err := parseElements(filename, doc.Blocks, attrs, levelOffset, opts...)
	if err != nil {
		return types.PreflightDocument{}, err
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
				append(opts, Entrypoint("PreflightDocumentWithinDelimitedBlock"))...)
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

var invalidFileTmpl *template.Template

func init() {
	var err error
	invalidFileTmpl, err = template.New("invalid file to include").Parse(`Unresolved directive in {{ .Filename }} - {{ .Error }}`)
	if err != nil {
		log.Fatalf("failed to initialize template: %v", err)
	}
}

type invalidFileData struct {
	Filename string
	Error    string
}

func invalidFileErrMsg(filename, path, rawText string, err error) (types.PreflightDocument, error) {
	log.WithError(err).Errorf("failed to include '%s'", path)
	buf := bytes.NewBuffer(nil)
	err = invalidFileTmpl.Execute(buf, invalidFileData{
		Filename: filename,
		Error:    rawText,
	})
	if err != nil {
		return types.PreflightDocument{}, err
	}
	return types.PreflightDocument{
		Blocks: []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: buf.String(),
						},
					},
				},
			},
		},
	}, nil
}

func parseFileToInclude(filename string, incl types.FileInclusion, attrs types.DocumentAttributes, opts ...Option) (types.PreflightDocument, error) {
	path := incl.Location.Resolve(attrs)
	log.Debugf("parsing '%s'...", path)
	f, absPath, done, err := open(path)
	defer done()
	if err != nil {
		return invalidFileErrMsg(filename, path, incl.RawText, err)
	}
	content := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	if lineRanges, ok := incl.LineRanges(); ok {
		if err := readWithinLines(scanner, content, lineRanges); err != nil {
			return invalidFileErrMsg(filename, path, incl.RawText, err)
		}
	} else if tagRanges, ok := incl.TagRanges(); ok {
		if err := readWithinTags(scanner, content, tagRanges); err != nil {
			return invalidFileErrMsg(filename, path, incl.RawText, err)
		}
	} else {
		if err := readAll(scanner, content); err != nil {
			return invalidFileErrMsg(filename, path, incl.RawText, err)
		}
	}
	if err := scanner.Err(); err != nil {
		msg, err2 := invalidFileErrMsg(filename, path, incl.RawText, err)
		if err2 != nil {
			return types.PreflightDocument{}, err2
		}
		return msg, errors.Wrap(err, "unable to read file to include")
	}
	// parse the content, and returns the corresponding elements
	levelOffset := incl.Attributes.GetAsString(types.AttrLevelOffset)
	return parsePreflightDocument(absPath, content, levelOffset, opts...)
}

func readWithinLines(scanner *bufio.Scanner, content *bytes.Buffer, lineRanges types.LineRanges) error {
	log.Debugf("limiting to line ranges: %v", lineRanges)
	line := 0
	for scanner.Scan() {
		line++
		log.Debugf("line %d: '%s' (matching range: %t)", line, scanner.Text(), lineRanges.Match(line))
		// parse the line in search for the `tag::<tag>[]` or `end:<tag>[]` macros
		l, err := Parse("", scanner.Bytes(), Entrypoint("IncludedFileLine"))
		if err != nil {
			return err
		}
		fl, ok := l.(types.IncludedFileLine)
		if !ok {
			return errors.Errorf("unexpected type of parsed line in file to include: %T", l)
		}
		// skip if the line has tags
		if fl.HasTag() {
			continue
		}
		// TODO: stop reading if current line above highest range
		if lineRanges.Match(line) {
			_, err := content.Write(scanner.Bytes())
			if err != nil {
				return err
			}
			_, err = content.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readWithinTags(scanner *bufio.Scanner, content *bytes.Buffer, tagRanges types.TagRanges) error {
	log.Debugf("limiting to tag ranges: %v", tagRanges)
	ranges := make(map[string]bool, len(tagRanges)) // ensure capacity
	for scanner.Scan() {
		line := scanner.Bytes()
		// parse the line in search for the `tag::<tag>[]` or `end:<tag>[]` macros
		l, err := Parse("", line, Entrypoint("IncludedFileLine"))
		if err != nil {
			return err
		}
		fl, ok := l.(types.IncludedFileLine)
		if !ok {
			return errors.Errorf("unexpected type of parsed line in file to include: %T", l)
		}
		// check if a start or end tag was found in the line
		if startTag, ok := fl.GetStartTag(); ok {
			ranges[startTag.Value] = true
		}
		if endTag, ok := fl.GetEndTag(); ok {
			ranges[endTag.Value] = false
		}

		if tagRanges.Match(ranges) && !fl.HasTag() {
			_, err := content.Write(scanner.Bytes())
			if err != nil {
				return err
			}
			_, err = content.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readAll(scanner *bufio.Scanner, content *bytes.Buffer) error {
	for scanner.Scan() {
		// parse the line in search for the `tag::<tag>[]` or `end:<tag>[]` macros
		l, err := Parse("", scanner.Bytes(), Entrypoint("IncludedFileLine"))
		if err != nil {
			return err
		}
		fl, ok := l.(types.IncludedFileLine)
		if !ok {
			return errors.Errorf("unexpected type of parsed line in file to include: %T", l)
		}
		// skip if the line has tags
		if fl.HasTag() {
			continue
		}
		_, err = content.Write(scanner.Bytes())
		if err != nil {
			return err
		}
		_, err = content.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func open(path string) (*os.File, string, func(), error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, "", func() {}, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, "", func() {
			log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	dir := filepath.Dir(absPath)
	err = os.Chdir(dir)
	if err != nil {
		return nil, "", func() {
			log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	// read the file per-se
	f, err := os.Open(absPath)
	if err != nil {
		return nil, absPath, func() {
			log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	return f, absPath, func() {
		log.Debugf("restoring current working dir to: %s", wd)
		if err := os.Chdir(wd); err != nil { // restore the previous working directory
			log.WithError(err).Error("failed to restore previous working directory")
		}
		if err := f.Close(); err != nil {
			log.WithError(err).Errorf("failed to close file '%s'", absPath)
		}
	}, nil
}

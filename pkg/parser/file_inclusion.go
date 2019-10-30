package parser

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var invalidFileTmpl *template.Template

func init() {
	var err error
	invalidFileTmpl, err = template.New("invalid file to include").Parse(`Unresolved directive in {{ .Filename }} - {{ .Error }}`)
	if err != nil {
		log.Fatalf("failed to initialize template: %v", err)
	}
}

func parseFileToInclude(filename string, incl types.FileInclusion, attrs types.DocumentAttributes, opts ...Option) (types.PreflightDocument, error) {
	path := incl.Location.Resolve(attrs)
	currentDir := filepath.Dir(filename)
	log.Debugf("parsing '%s' from '%s' (%s)", path, currentDir, filename)
	log.Debugf("file inclusion attributes: %s", spew.Sdump(incl.Attributes))
	f, absPath, done, err := open(filepath.Join(currentDir, path))
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
		if err := readWithinTags(path, scanner, content, tagRanges); err != nil {
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
	if !IsAsciidoc(absPath) {
		opts = append(opts, Entrypoint("PreflightTextDocument"))
	}

	return parsePreflightDocument(absPath, content, levelOffset, opts...)
}

func invalidFileErrMsg(filename, path, rawText string, err error) (types.PreflightDocument, error) {
	log.WithError(err).Errorf("failed to include '%s'", path)
	buf := bytes.NewBuffer(nil)
	err = invalidFileTmpl.Execute(buf, struct {
		Filename string
		Error    string
	}{
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

func readWithinTags(path string, scanner *bufio.Scanner, content *bytes.Buffer, expectedRanges types.TagRanges) error {
	log.Debugf("limiting to tag ranges: %v", expectedRanges)
	currentRanges := make(map[string]*types.CurrentTagRange, len(expectedRanges)) // ensure capacity
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
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
			currentRanges[startTag.Value] = &types.CurrentTagRange{
				StartLine: lineNumber,
				EndLine:   -1,
			}
		}
		if endTag, ok := fl.GetEndTag(); ok {
			currentRanges[endTag.Value].EndLine = lineNumber
		}
		if expectedRanges.Match(lineNumber, currentRanges) && !fl.HasTag() {
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
	// after the file has been processed, let's check if all tags were "found"
	for _, tag := range expectedRanges {
		log.Debugf("checking if tag '%s' was found...", tag.Name)
		switch tag.Name {
		case "*", "**":
			continue
		default:
			tr, found := currentRanges[tag.Name]
			if !found {
				log.Errorf("tag '%s' not found in include file: %s", tag.Name, path)
			} else if tr.EndLine == -1 {
				log.Errorf("detected unclosed tag '%s' starting at line %d of include file: %s", tag.Name, tr.StartLine, path)
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
	log.Debugf("file path: %s", absPath)
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

// IsAsciidoc returns true if the file to include is an asciidoc file (based on the file location extension)
func IsAsciidoc(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".asciidoc" || ext == ".adoc" || ext == ".ad" || ext == ".asc" || ext == ".txt"
}

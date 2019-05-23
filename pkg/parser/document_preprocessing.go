package parser

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ContextKey a non-built-in type for keys in the context
type ContextKey string

// LevelOffset the key for the level offset of the file to include
const LevelOffset ContextKey = "leveloffset"

// PreparseDocument reads a document raw content and applies the preprocessing directives (file inclusions)
// the result is a `[]byte` which can then be fully parsed
func PreparseDocument(filename string, r io.Reader, opts ...Option) ([]byte, error) {
	return preparseDocument(filename, r, "", opts...)
}

func preparseDocument(filename string, r io.Reader, levelOffset string, opts ...Option) ([]byte, error) {
	d, err := ParseReader(filename, r, Entrypoint("PreparsedDocument"))
	if err != nil {
		return nil, err
	}
	doc := d.(types.PreparsedDocument)
	result := bytes.NewBuffer(nil)
	attrs := types.DocumentAttributes{}
	for i, e := range doc.Elements {
		switch e := e.(type) {
		case types.FileInclusion:
			// read the file and include its content
			content, err := parseFileToInclude(e, attrs, opts...)
			if err != nil {
				// do not fail, but instead report the error in the console
				log.Errorf("failed to include file '%s': %v", e.Location, err)
			}
			result.Write(content)
		case types.DocumentAttributeDeclaration:
			attrs[e.Name] = e.Value
			result.WriteRune(':')
			result.WriteString(e.Name)
			result.WriteRune(':')
			if e.Value != "" {
				result.WriteRune(' ')
				result.WriteString(e.Value)
			}
		case types.BlankLine:
			// nothing to do in this case. We can ignore the spaces/tabs that were on this line anyways
		case types.RawSectionTitle:
			s, err := e.Bytes(levelOffset)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to preparse '%s'", filename)
			}
			result.Write(s)
		case types.RawText:
			// just write the content
			result.Write(e.Bytes())
		}
		if i < len(doc.Elements)-1 {
			result.WriteString("\n")
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("preparsed document '%s':", filename)
		log.Debugf("%s", result.String())
	}
	return result.Bytes(), nil
}

var invalidFileTmpl *template.Template

func init() {
	var err error
	invalidFileTmpl, err = template.New("invalid file to include").Parse(`Unresolved directive in test.adoc - {{ . }}`)
	if err != nil {
		log.Fatalf("failed to initialize template: %v", err)
	}
}

func invalidFileErrMsg(incl types.FileInclusion) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := invalidFileTmpl.Execute(buf, incl.RawText)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseFileToInclude(incl types.FileInclusion, attrs types.DocumentAttributes, opts ...Option) ([]byte, error) {
	path := incl.Location.Resolve(attrs)
	log.Debugf("parsing '%s'...", path)
	// manage new working directory based on the file's location
	// so that if this file also includes other files with relative path,
	// then the it can work ;)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return invalidFileErrMsg(incl)
	}
	dir := filepath.Dir(absPath)
	err = os.Chdir(dir)
	if err != nil {
		return invalidFileErrMsg(incl)
	}
	defer func() {
		err = os.Chdir(wd) // restore the previous working directory
		if err != nil {
			log.WithError(err).Error("failed to restore previous working directory")
		}
	}()
	// read the file per-se
	f, err := os.Open(absPath)
	if err != nil {
		return invalidFileErrMsg(incl)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	var lineRanges types.LineRanges
	if lr, ok := incl.Attributes[types.AttrLineRanges].(types.LineRanges); ok {
		lineRanges = lr
	} else {
		lineRanges = types.LineRanges{ // default line ranges: include all content
			{
				Start: 1,
				End:   -1,
			},
		}
	}
	content := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	line := 1
	for scanner.Scan() {
		log.Debugf("line %d: '%s' (matching range: %t)", line, scanner.Text(), lineRanges.Match(line))
		// TODO: stop reading if current line above highest range
		if lineRanges.Match(line) {
			_, err := content.WriteString(scanner.Text())
			if err != nil {
				return invalidFileErrMsg(incl)
			}
			_, err = content.WriteString("\n")
			if err != nil {
				return invalidFileErrMsg(incl)
			}
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		msg, err2 := invalidFileErrMsg(incl)
		if err2 != nil {
			return nil, err2
		}
		return msg, errors.Wrap(err, "unable to read file to include")
	}
	if levelOffset, ok := incl.Attributes[types.AttrLevelOffset].(string); ok {
		return preparseDocument(path, content, levelOffset, opts...)
	}
	return preparseDocument(path, content, "", opts...)
}

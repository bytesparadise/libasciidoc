package parser

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	errs "github.com/pkg/errors"
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
	doc, ok := d.(types.PreparsedDocument)
	if !ok {
		return nil, errs.Errorf("invalid type of result: %T (expected a PreparsedDocument)", d)
	}
	result := bytes.NewBuffer(nil)
	for i, e := range doc.Elements {
		switch e := e.(type) {
		case types.FileInclusion:
			// read the file and include its content
			content, err := parseFileToInclude(e, opts...)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to include file '%s", e.Path)
			}
			result.Write(content)
		case types.DocumentAttributeDeclaration:
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

func parseFileToInclude(incl types.FileInclusion, opts ...Option) ([]byte, error) {
	log.Debugf("parsing '%s'...", incl.Path)
	// manage new working directory based on the file's location
	// so that if this file also includes other files with relative path,
	// then the it can work ;)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(incl.Path)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(absPath)
	err = os.Chdir(dir)
	if err != nil {
		return nil, err
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
		return nil, errors.Wrapf(err, "unable to read file to include")
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
				return nil, errors.Wrap(err, "unable to read file to include")
			}
			_, err = content.WriteString("\n")
			if err != nil {
				return nil, errors.Wrap(err, "unable to read file to include")
			}
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "unable to read file to include")
	}
	if levelOffset, ok := incl.Attributes[types.AttrLevelOffset].(string); ok {
		return preparseDocument(incl.Path, content, levelOffset, opts...)
	}
	return preparseDocument(incl.Path, content, "", opts...)
}

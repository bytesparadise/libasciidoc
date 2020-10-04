package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// ParseRawSource parses a document's content and applies the preprocessing directives (file inclusions)
func ParseRawSource(r io.Reader, config configuration.Configuration, options ...Option) ([]byte, error) {
	attrs := types.AttributesWithOverrides{
		Content:   map[string]interface{}{},
		Overrides: map[string]string{},
		Counters:  map[string]interface{}{},
	}
	return parseRawSource(r, attrs, []levelOffset{}, config, append(options, Entrypoint("RawSource"))...)
}

func parseRawSource(r io.Reader, attrs types.AttributesWithOverrides, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) ([]byte, error) {
	// log.Debugf("parsing raw document '%s'", config.Filename)
	lines, err := ParseReader(config.Filename, r, options...)
	if err != nil {
		log.Errorf("failed to parse raw document: %s", err)
		return nil, err
	}
	l, ok := lines.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type of raw lines: '%T'", lines)
	}
	return processFileInclusions(l, attrs, levelOffsets, config, options...)
}

// processFileInclusions processes the file inclusions in the given lines and returns a serialized content which can be parsed again
func processFileInclusions(lines []interface{}, globalAttrs types.AttributesWithOverrides, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	for _, line := range lines {
		switch l := line.(type) {
		case types.RawLine:
			result.WriteString(l.Stringify())
			// append linefeed
			result.WriteString("\n")
		case types.RawSection:
			for _, offset := range levelOffsets {
				oldLevel := l.Level
				offset.apply(&l)
				// replace the absolute when the first section is processed with a relative offset
				// which is based on the actual level offset that resulted in the application of the absolute offset
				if offset.absolute {
					levelOffsets = []levelOffset{
						relativeOffset(l.Level - oldLevel),
					}
				}
			}
			result.WriteString(l.Stringify())
			// append linefeed
			result.WriteString("\n")
		case types.AttributeDeclaration:
			globalAttrs.Set(l.Name, l.Value)
			result.WriteString(l.Stringify())
			// append linefeed
			result.WriteString("\n")
		case types.FileInclusion:
			includedLines, err := parseFileToInclude(l, globalAttrs, levelOffsets, config, options...)
			if err != nil {
				return nil, err
			}
			result.Write(includedLines)
		default:
			return nil, fmt.Errorf("unexpected type of line: '%T'", line)
		}
	}
	return result.Bytes(), nil

}

// levelOffset a func that applies a given offset to the sections of a child document to include in a parent doc (the caller)
type levelOffset struct {
	absolute bool
	value    int
	apply    func(*types.RawSection)
}

func relativeOffset(offset int) levelOffset {
	return levelOffset{
		absolute: false,
		value:    offset,
		apply: func(s *types.RawSection) {
			s.Level += offset
		},
	}
}

func absoluteOffset(offset int) levelOffset {
	return levelOffset{
		absolute: true,
		value:    offset,
		apply: func(s *types.RawSection) {
			s.Level = offset
		},
	}
}

// applies the elements and attributes substitutions on the given image block.
func applySubstitutionsOnFileInclusion(f types.FileInclusion, attrs types.AttributesWithOverrides, options ...Option) (types.FileInclusion, error) {
	elements := f.Location.Path
	subs := []elementsSubstitution{substituteAttributes} // TODO: no need for an array here
	var err error
	for _, sub := range subs {
		if elements, err = sub(elements, attrs, options...); err != nil {
			return types.FileInclusion{}, err
		}
	}
	f.Location.Path = elements
	return f, nil
}

func parseFileToInclude(incl types.FileInclusion, attrs types.AttributesWithOverrides, levelOffsets []levelOffset, config configuration.Configuration, options ...Option) ([]byte, error) {
	incl, err := applySubstitutionsOnFileInclusion(incl, attrs)
	if err != nil {
		return nil, err
	}
	path := incl.Location.Stringify()
	currentDir := filepath.Dir(config.Filename)
	// log.Debugf("parsing '%s' from current dir '%s' (%s)", path, currentDir, config.Filename)
	f, absPath, done, err := open(filepath.Join(currentDir, path))
	defer done()
	if err != nil {
		return nil, fmt.Errorf("Unresolved directive in %s - %s", config.Filename, incl.RawText)
	}
	content := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	if lineRanges, ok := incl.LineRanges(); ok {
		if err := readWithinLines(scanner, content, lineRanges); err != nil {
			return nil, fmt.Errorf("Unresolved directive in %s - %s", config.Filename, incl.RawText)
		}
	} else if tagRanges, ok := incl.TagRanges(); ok {
		if err := readWithinTags(path, scanner, content, tagRanges); err != nil {
			return nil, err
		}
	} else {
		if err := readAll(scanner, content); err != nil {
			return nil, fmt.Errorf("Unresolved directive in %s - %s", config.Filename, incl.RawText)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Unresolved directive in %s - %s", config.Filename, incl.RawText)
	}
	// just include the file content if the file to include is not an Asciidoc document.
	if !IsAsciidoc(absPath) {
		return content.Bytes(), nil
	}
	// parse the content, and returns the corresponding elements
	if l, found := incl.Attributes.GetAsString(types.AttrLevelOffset); found {
		offset, err := strconv.Atoi(l)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read file to include")
		}
		if strings.HasPrefix(l, "+") || strings.HasPrefix(l, "-") {
			levelOffsets = append(levelOffsets, relativeOffset(offset))
		} else {
			levelOffsets = []levelOffset{absoluteOffset(offset)}

		}
	}

	inclConfig := config.Clone()
	inclConfig.Filename = absPath
	// now, let's parse this content and process nested file inclusions
	return parseRawSource(content, attrs, levelOffsets, inclConfig, options...)
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
			if tr, found := currentRanges[tag.Name]; !found {
				return fmt.Errorf("tag '%s' not found in include file: %s", tag.Name, path)
			} else if tr.EndLine == -1 {
				log.Warnf("detected unclosed tag '%s' starting at line %d of include file: %s", tag.Name, tr.StartLine, path)
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
	log.Debugf("opening '%s'", absPath)
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

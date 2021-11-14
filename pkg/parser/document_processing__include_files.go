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

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func IncludeFiles(ctx *ParseContext, done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) <-chan types.DocumentFragment {
	resultStream := make(chan types.DocumentFragment, 1)
	go func() {
		defer close(resultStream)
		for fragment := range fragmentStream {
			select {
			case resultStream <- includeFiles(ctx, fragment.Elements, done):
			case <-done:
				log.WithField("pipeline_task", "include_files").Debug("received 'done' signal")
				return
			}
		}
		log.WithField("pipeline_task", "include_files").Debug("done")
	}()
	return resultStream
}

func includeFiles(ctx *ParseContext, elements []interface{}, done <-chan interface{}) types.DocumentFragment {
	result := make([]interface{}, 0, len(elements))
	for _, element := range elements {
		switch e := element.(type) {
		case *types.AttributeDeclaration:
			ctx.attributes.set(e.Name, e.Value)
			result = append(result, element)
		case *types.AttributeReset:
			ctx.attributes.unset(e.Name)
			result = append(result, element)
		case *types.FileInclusion:
			// use an Entrypoint based on the Delimited block kind
			f := doIncludeFile(ctx.Clone(), e, done)
			if f.Error != nil {
				return f
			}
			result = append(result, f.Elements...)
		case *types.DelimitedBlock:
			f := includeFiles(ctx.WithinDelimitedBlock(e), e.Elements, done)
			if f.Error != nil {
				return f
			}
			e.Elements = f.Elements
			result = append(result, e)
		default:
			result = append(result, element)
		}
	}
	return types.NewDocumentFragment(result...)
}

// replace the content of this FileInclusion element the content of the target file
// note: there is a trade-off here: we include the whole content of the file in the current
// fragment, making it potentially big, but at the same time we ensure that the context
// of the inclusion (for example, within a delimited block) is not lost.
func doIncludeFile(ctx *ParseContext, e *types.FileInclusion, done <-chan interface{}) types.DocumentFragment {
	// ctx.Opts = append(ctx.Opts, GlobalStore(documentHeaderKey, true))
	fileContent, err := contentOf(ctx, e)
	if err != nil {
		return types.NewErrorFragment(err)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("including content of '%s' with offsets %v", e.Location.Stringify(), ctx.levelOffsets)
	}
	elements := []interface{}{}
	for f := range ParseFragments(ctx, fileContent, done) {
		if f.Error != nil {
			return f
		}
		// apply level offset on sections
		for i, e := range f.Elements {
			switch e := e.(type) {
			case *types.DocumentHeader:
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf("applying offsets to header section: %v", ctx.levelOffsets)
				}
				// header becomes a "regular" section
				s := &types.Section{
					Title:    e.Title,
					Elements: e.Elements,
				}
				ctx.levelOffsets.apply(s)
				if s.Level == 0 { // no level change: keep as the header
					f.Elements[i] = e
				} else { // level changed: becomes a section with some elements
					f.Elements[i] = s
				}
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf("applied offsets to header/section: level is now %d", s.Level)
				}
			case *types.Section:
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf("applying offsets to section of level %d: %v", e.Level, ctx.levelOffsets)
				}
				ctx.levelOffsets.apply(e)
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf("applied offsets to section: level is now %d", e.Level)
				}
			}
		}
		elements = append(elements, f.Elements...)
	}
	// and recursively...
	return includeFiles(ctx, elements, done)
}

func contentOf(ctx *ParseContext, incl *types.FileInclusion) (io.Reader, error) {
	if err := applySubstitutionsOnBlockWithLocation(ctx, incl); err != nil {
		log.Error(err)
		return nil, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	}
	path := incl.Location.Stringify()
	currentDir := filepath.Dir(ctx.filename)
	f, absPath, closeFile, err := open(filepath.Join(currentDir, path))
	if err != nil {
		return nil, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	}
	defer closeFile()
	content := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsing file to %s", incl.RawText)
	}
	if lr, ok, err := lineRanges(incl); err != nil {
		log.Error(err)
		return nil, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if ok {
		if err := readWithinLines(scanner, content, lr); err != nil {
			log.Error(err)
			return nil, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	} else if tr, ok, err := tagRanges(incl); err != nil {
		log.Error(err)
		return nil, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if ok {
		if err := readWithinTags(path, scanner, content, tr); err != nil {
			log.Error(err)
			return nil, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	} else {
		if err := readAll(scanner, content); err != nil {
			log.Error(err)
			return nil, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Error(err)
		return nil, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	}
	// cloning the context to avoid altering the original as we process recursively embedded file inclusions
	ctx.filename = absPath
	// if the file to include is not an Asciidoc document, just return the content as "raw lines"
	if !IsAsciidoc(absPath) {
		log.Debugf("file '%s' is not an Asciidoc file. Scanning content without looking for nested file inclusions", absPath)
		// send(types.NewDocumentFragment(lineOffset, []interface{}{types.RawLine(content.String())}), done, resultStream)
		// return
		ctx.Opts = append(ctx.Opts, DisableRule(FileInclusion))
	}

	// level offset
	// parse the content, and returns the corresponding elements
	if lvl, found, err := incl.Attributes.GetAsString(types.AttrLevelOffset); err != nil {
		log.Error(err)
		return nil, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if found {
		offset, err := strconv.Atoi(lvl)
		if err != nil {
			log.Error(err)
			return nil, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
		if strings.HasPrefix(lvl, "+") || strings.HasPrefix(lvl, "-") {
			ctx.levelOffsets = append(ctx.levelOffsets, relativeOffset(offset))
		} else {
			ctx.levelOffsets = []*levelOffset{absoluteOffset(offset)}
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("content of '%s':\n%s", absPath, content.String())
	}
	return content, nil
}

type levelOffsets []*levelOffset

func (l levelOffsets) apply(s *types.Section) {
	for _, offset := range l {
		offset.apply(s)
	}
}

func (l levelOffsets) clone() levelOffsets {
	result := make([]*levelOffset, len(l))
	copy(result, l)
	return result
}

// levelOffset a func that applies a given offset to the sections of a child document to include in a parent doc (the caller)
type levelOffset struct {
	absolute bool
	value    int
}

func (l *levelOffset) apply(s *types.Section) {
	// also, absolute offset becomes relative offset after processing the first section,
	// so that the hierarchy of subsequent sections of the doc to include is preserved
	if l.absolute {
		l.absolute = false
		l.value = l.value - s.Level
	}
	s.Level += l.value
}

func relativeOffset(offset int) *levelOffset {
	return &levelOffset{
		absolute: false,
		value:    offset,
	}
}

func absoluteOffset(offset int) *levelOffset {
	return &levelOffset{
		absolute: true,
		value:    offset,
	}
}

func (l levelOffset) String() string {
	if l.absolute {
		return strconv.Itoa(l.value)
	}
	return fmt.Sprintf("+%d", l.value)
}

// lineRanges parses the `lines` attribute if it exists in the given FileInclusion, and returns
// a corresponding `LineRanges` (or `false` if parsing failed to invalid input)
func lineRanges(incl *types.FileInclusion) (types.LineRanges, bool, error) {
	lineRanges, exists, err := incl.Attributes.GetAsString(types.AttrLineRanges)
	if err != nil {
		return types.LineRanges{}, false, err
	}
	if exists {
		lr, err := Parse("", []byte(lineRanges), Entrypoint("LineRanges"))
		if err != nil {
			return types.LineRanges{}, false, err
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("line ranges to include: %s", spew.Sdump(lr))
		}
		return types.NewLineRanges(lr), true, nil
	}
	return types.LineRanges{}, false, nil
}

// tagRanges parses the `tags` attribute if it exists in the given FileInclusion, and returns
// a corresponding `TagRanges` (or `false` if parsing failed to invalid input)
func tagRanges(incl *types.FileInclusion) (types.TagRanges, bool, error) {
	tagRanges, exists, err := incl.Attributes.GetAsString(types.AttrTagRanges)
	if err != nil {
		return types.TagRanges{}, false, err
	}
	if exists {
		log.Debugf("tag ranges to include: %v", spew.Sdump(tagRanges))
		tr, err := Parse("", []byte(tagRanges), Entrypoint("TagRanges"))
		if err != nil {
			return types.TagRanges{}, false, err
		}
		return types.NewTagRanges(tr), true, nil
	}
	return types.TagRanges{}, false, nil
}

// TODO: instead of reading and parsing afterwards, simply parse the lines immediately? ie: `readWithinLines` -> `parseWithinLines`
// (also, use a specific entrypoint if the doc is not a .adoc)
func readWithinLines(scanner *bufio.Scanner, content *bytes.Buffer, lineRanges types.LineRanges) error {
	line := 0
	for scanner.Scan() {
		line++
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
	// log.Debugf("limiting to tag ranges: %v", expectedRanges)
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
		// log.Debugf("checking if tag '%s' was found...", tag.Name)
		switch tag.Name {
		case "*", "**":
			continue
		default:
			if tr, found := currentRanges[tag.Name]; !found {
				return fmt.Errorf("tag '%s' not found in file to include", tag.Name)
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
			// log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	dir := filepath.Dir(absPath)
	if err = os.Chdir(dir); err != nil {
		return nil, "", func() {
			// log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	// read the file per-se
	// log.Debugf("opening '%s'", absPath)
	f, err := os.Open(absPath)
	if err != nil {
		return nil, absPath, func() {
			// log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	return f, absPath, func() {
		// log.Debugf("restoring current working dir to: %s", wd)
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

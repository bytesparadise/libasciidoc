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
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Preprocess reads line by line to look-up and process file inclusions and conditionals (`ifdef`, `ifndef` and `ifeval`)
func Preprocess(source io.Reader, config *configuration.Configuration, opts ...Option) (string, error) {
	ctx := NewParseContext(config, opts...) // each pipeline step will have its own clone of `ctx`
	return preprocess(ctx, source)
}

func preprocess(ctx *ParseContext, source io.Reader) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("processing file inclusions in %s with leveloffset=%s", ctx.filename, spew.Sdump(ctx.levelOffsets))
	}
	b := &builder{
		enabled: true,
	}
	c := conditions{}
	scanner := bufio.NewScanner(source)
	t := newBlockDelimiterTracker()
	for scanner.Scan() {
		line := scanner.Bytes()
		element, err := Parse("", line, append(ctx.opts, Entrypoint("DocumentRawLine"))...)
		if err != nil {
			// log.Error(err)
			// content of line was not relevant in the context of preparsing (ie, it's a regular line), so let's keep it as-is
			b.Write(line)
		} else {
			if log.IsLevelEnabled(log.DebugLevel) {
				log.Debugf("checking element of type '%T'", element)
			}
			switch e := element.(type) {
			case *types.AttributeDeclaration:
				ctx.attributes.set(e.Name, e.Value)
				b.WriteString(e.RawText())
			case *types.AttributeReset:
				ctx.attributes.unset(e.Name)
				b.WriteString(e.RawText())
			case *types.RawSection:
				b.WriteString(ctx.levelOffsets.apply(e))
			case *types.FileInclusion:
				f, err := includeFile(ctx.Clone(), e)
				if err != nil {
					return "", err
				}
				b.WriteString(f)
			case *types.BlockDelimiter:
				t.track(e.Kind, e.Length)
				ctx.opts = append(ctx.opts, t.withinDelimitedBlock())
				b.WriteString(e.RawText())
			case types.ConditionalInclusion:
				if content, ok := e.SingleLineContent(); ok {
					if e.Eval(ctx.attributes.allAttributes()) {
						b.WriteString(content)
					}
				} else {
					b.enabled = c.push(ctx, e)
				}
			case *types.EndOfCondition:
				b.enabled = c.pop()
			default:
				return "", fmt.Errorf("unexpected type of element while preprocessinh document: '%T'", e)
			}
		}
	}
	return b.String(), nil
}

type blockDelimiterTracker struct {
	stack []blockDelimiter
}

type blockDelimiter struct {
	kind   string
	length int
}

func newBlockDelimiterTracker() *blockDelimiterTracker {
	return &blockDelimiterTracker{
		stack: []blockDelimiter{},
	}
}

func (t *blockDelimiterTracker) track(kind string, length int) {
	switch {
	case len(t.stack) > 0 && t.stack[len(t.stack)-1].kind == kind && t.stack[len(t.stack)-1].length == length:
		// pop
		t.stack = t.stack[:len(t.stack)-1]
	default:
		// push
		t.stack = append(t.stack, blockDelimiter{
			kind:   kind,
			length: length,
		})
	}
}

func (t *blockDelimiterTracker) withinDelimitedBlock() Option {
	return GlobalStore(withinDelimitedBlockKey, len(t.stack) > 0)
}

// replace the content of this FileInclusion element the content of the target file
// note: there is a trade-off here: we include the whole content of the file in the current
// fragment, making it potentially big, but at the same time we ensure that the context
// of the inclusion (for example, within a delimited block) is not lost.
func includeFile(ctx *ParseContext, incl *types.FileInclusion) (string, error) {
	ctx.opts = append(ctx.opts, GlobalStore(documentHeaderKey, false))
	if l, ok := incl.GetLocation().Path.([]interface{}); ok {
		l, _, err := replaceAttributeRefsInSlice(ctx, l, noneSubstitutions())
		if err != nil {
			return "", errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
		incl.GetLocation().SetPath(l)
	}
	content, adoc, err := contentOf(ctx, incl)
	if err != nil {
		return "", err
	}
	if !adoc {
		return string(content), nil
	}
	ctx.opts = append(ctx.opts, sectionEnabled())
	return preprocess(ctx, bytes.NewReader(content))
}

type builder struct {
	strings.Builder
	insertLF bool
	enabled  bool
}

func (b *builder) WriteString(s string) {
	if !b.enabled {
		return
	}
	b.doInsertLF()
	b.Builder.WriteString(s)
}

func (b *builder) Write(p []byte) {
	if !b.enabled {
		return
	}
	b.doInsertLF()
	b.Builder.Write(p)
}

func (b *builder) doInsertLF() {
	if b.insertLF {
		b.Builder.WriteString("\n")
	}
	// from now on, we will insert a `\n` before each new line (but there will be no extra `\n` at the end of the content)
	b.insertLF = true
}

// a stack of conditions
type conditions struct {
	elements []condition
}

func newConditions() *conditions {
	return &conditions{
		elements: []condition{},
	}
}

type condition struct {
	element types.ConditionalInclusion
	eval    bool
}

func (c *conditions) push(ctx *ParseContext, element types.ConditionalInclusion) bool {
	c.elements = append(c.elements, condition{
		element: element,
		eval:    element.Eval(ctx.attributes.allAttributes()),
	})
	return c.eval()
}

func (c *conditions) pop() bool {
	// TODO: report an error if stack is empty (ie, unbalanced )
	if len(c.elements) > 0 {
		c.elements = c.elements[:len(c.elements)-1]
	}
	return c.eval()
}

func (c *conditions) eval() bool {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("evaluating %s", spew.Sdump(c.elements))
	}
	for _, e := range c.elements {
		// assume all elements are `ConditionalInclusion`
		if !e.eval {
			return false
		}
	}
	return true
}

func contentOf(ctx *ParseContext, incl *types.FileInclusion) ([]byte, bool, error) {
	path := incl.Location.Stringify()
	currentDir := filepath.Dir(ctx.filename)
	filename := filepath.Join(currentDir, path)

	f, absPath, closeFile, err := open(filename)
	defer closeFile()
	if err != nil {
		return nil, false, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	}
	result := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("reading %s", filename)
	}
	if lr, ok, err := lineRanges(incl); err != nil {
		return nil, false, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if ok {
		if err := readWithinLines(scanner, result, lr); err != nil {
			return nil, false, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	} else if tr, ok, err := tagRanges(incl); err != nil {
		return nil, false, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if ok {
		if err := readWithinTags(path, scanner, result, tr); err != nil {
			return nil, false, errors.Wrapf(err, "Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	} else {
		if err := readAll(scanner, result); err != nil {
			log.Error(err)
			return nil, false, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Error(err)
		return nil, false, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	}
	// cloning the context to avoid altering the original as we process recursively embedded file inclusions
	ctx.filename = absPath
	// if the file to include is not an Asciidoc document, just return the content as "raw lines"

	// level offset
	// parse the content, and returns the corresponding elements
	if lvl, found, err := incl.Attributes.GetAsString(types.AttrLevelOffset); err != nil {
		log.Error(err)
		return nil, false, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
	} else if found {
		offset, err := strconv.Atoi(lvl)
		if err != nil {
			log.Error(err)
			return nil, false, errors.Errorf("Unresolved directive in %s - %s", ctx.filename, incl.RawText)
		}
		if strings.HasPrefix(lvl, "+") || strings.HasPrefix(lvl, "-") {
			ctx.levelOffsets = append(ctx.levelOffsets, relativeOffset(offset))
		} else {
			ctx.levelOffsets = []*levelOffset{absoluteOffset(offset)}
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("content of '%s':\n%s", absPath, result.String())
	// }
	return result.Bytes(), IsAsciidoc(absPath), nil
}

type levelOffsets []*levelOffset

func (l levelOffsets) apply(s *types.RawSection) string {
	for _, offset := range l {
		offset.apply(s)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applied offsets to section: level is now %d", s.Level)
	}
	return s.Stringify()
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

func (l *levelOffset) apply(s *types.RawSection) {
	// also, absolute offset becomes relative offset after processing the first section,
	// so that the hierarchy of subsequent sections of the doc to include is preserved
	if l.absolute {
		l.absolute = false
		l.value = l.value - s.Level
	}
	s.OffsetLevel(l.value)
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
			if _, err := content.Write(scanner.Bytes()); err != nil {
				return err
			}
			if _, err := content.WriteString("\n"); err != nil {
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
			log.Debugf("restoring current working dir to: %s", wd)
			if err := os.Chdir(wd); err != nil { // restore the previous working directory
				log.WithError(err).Error("failed to restore previous working directory")
			}
		}, err
	}
	dir := filepath.Dir(absPath)
	// TODO: we could skip the Chdir part if we retain the absPath in the context,
	// and use `filepath.Join` to compute the abspath of the file to include
	if err = os.Chdir(dir); err != nil {
		return nil, "", func() {
			log.Debugf("restoring current working dir to: %s", wd)
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

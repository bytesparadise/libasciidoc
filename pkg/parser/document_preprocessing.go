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
func ParsePreflightDocument(filename string, r io.Reader, opts ...Option) (*types.PreflightDocument, error) {
	// opts = append(opts, Entrypoint("PreflightDocument"), Memoize(true), Recover(false))
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, Debug(true))
	// }
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	preparsingStats := Stats{}
	// 	opts = append(opts, Statistics(&preparsingStats, "no match"))
	// 	start := time.Now()
	// 	defer func() {
	// 		preparseDuration := time.Since(start)
	// 		log.Infof("preparsing stats:")
	// 		log.Infof("- duration:                  %v", preparseDuration)
	// 		log.Infof("- expressions:               %v", preparsingStats.ExprCnt)
	// 	}()
	// }
	opts = append(opts, Entrypoint("PreflightDocument"))
	return parsePreflightDocument(filename, r, "", opts...)
}

func parsePreflightDocument(filename string, r io.Reader, levelOffset string, opts ...Option) (*types.PreflightDocument, error) {
	d, err := ParseReader(filename, r, opts...)
	if err != nil {
		return nil, err
	}
	doc := d.(*types.PreflightDocument)
	attrs := types.DocumentAttributes{}
	blocks, err := parseElements(filename, doc.Blocks, attrs, levelOffset, opts...)
	if err != nil {
		return nil, err
	}
	doc.Blocks = blocks
	return doc, nil
}

// parseElements resolves the file inclusions if any is found in the given elements
func parseElements(filename string, elements []interface{}, attrs types.DocumentAttributes, levelOffset string, opts ...Option) ([]interface{}, error) {
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case *types.DocumentAttributeDeclaration:
			attrs[e.Name] = e.Value
			result = append(result, e)
		case *types.FileInclusion:
			// read the file and include its content
			embedded, err := parseFileToInclude(e, attrs, opts...)
			if err != nil {
				// do not fail, but instead report the error in the console
				log.Errorf("failed to include file '%s': %v", e.Location, err)
			}
			result = append(result, embedded.Blocks...)
		case *types.DelimitedBlock:
			elmts, err := parseElements(filename, e.Elements, attrs, levelOffset,
				// use a new var to avoid overridding the current one which needs to stay as-is for the rest of the doc parsing
				append(opts, Entrypoint("PreflightDocumentWithinDelimitedBlock"))...)
			if err != nil {
				return nil, err
			}
			result = append(result, &types.DelimitedBlock{
				Attributes: e.Attributes,
				Kind:       e.Kind,
				Elements:   elmts,
			})
		case *types.Section:
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
	invalidFileTmpl, err = template.New("invalid file to include").Parse(`Unresolved directive in test.adoc - {{ . }}`)
	if err != nil {
		log.Fatalf("failed to initialize template: %v", err)
	}
}

func invalidFileErrMsg(incl *types.FileInclusion) (*types.PreflightDocument, error) {
	buf := bytes.NewBuffer(nil)
	err := invalidFileTmpl.Execute(buf, incl.RawText)
	if err != nil {
		return nil, err
	}
	return &types.PreflightDocument{
		Blocks: []interface{}{
			&types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: buf.String(),
						},
					},
				},
			},
		},
	}, nil
}

func parseFileToInclude(incl *types.FileInclusion, attrs types.DocumentAttributes, opts ...Option) (*types.PreflightDocument, error) {
	path := incl.Location.Resolve(attrs)
	log.Debugf("parsing '%s'...", path)
	// manage new working directory based on the file's location
	// so that if this file also includes other files with relative path,
	// then the it can work ;)
	return parseAsciidocFile(incl, path, opts...)
}

func parseAsciidocFile(incl *types.FileInclusion, path string, opts ...Option) (*types.PreflightDocument, error) {
	f, absPath, done, err := open(path)
	defer done()
	if err != nil {
		return invalidFileErrMsg(incl)
	}
	content := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	lineRanges := incl.LineRanges()
	log.Debugf("limiting to line range: %v", lineRanges)
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
	// parse the content, and returns the corresponding elements
	if levelOffset, ok := incl.Attributes[types.AttrLevelOffset].(string); ok {
		return parsePreflightDocument(absPath, content, levelOffset, opts...)
	}
	return parsePreflightDocument(absPath, content, "", opts...)
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
		log.Errorf("error while opening '%s': %v", path, err)
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

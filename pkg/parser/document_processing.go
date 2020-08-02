package parser

import (
	"fmt"
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ParseDocument parses the content of the reader identitied by the filename
func ParseDocument(r io.Reader, config configuration.Configuration, options ...Option) (types.Document, error) {
	rawDoc, err := ParseRawDocument(r, config, options...)
	if err != nil {
		return types.Document{}, err
	}

	draftDoc, err := ApplySubstitutions(rawDoc, config, options...)
	if err != nil {
		return types.Document{}, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("draft doc:")
		spew.Fdump(log.StandardLogger().Out, draftDoc)
	}

	// now, merge list items into proper lists
	blocks, err := rearrangeListItems(draftDoc.Blocks, false)
	if err != nil {
		return types.Document{}, err
	}
	// filter out blocks not needed in the final doc
	blocks = filter(blocks, allMatchers...)

	blocks, footnotes := processFootnotes(blocks)
	// now, rearrange elements in a hierarchical manner
	doc := rearrangeSections(blocks)
	// also, set the footnotes
	doc.Footnotes = footnotes
	// insert the preamble at the right location
	doc = includePreamble(doc)
	doc.Attributes = doc.Attributes.Add(draftDoc.Attributes)
	// also insert the table of contents
	doc = includeTableOfContentsPlaceHolder(doc)
	// finally
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("final document:")
		spew.Fdump(log.StandardLogger().Out, doc)
	}
	return doc, nil
}

// ContextKey a non-built-in type for keys in the context
type ContextKey string

// LevelOffset the key for the level offset of the file to include
const LevelOffset ContextKey = "leveloffset"

// ParseRawDocument parses a document's content and applies the preprocessing directives (file inclusions)
func ParseRawDocument(r io.Reader, config configuration.Configuration, options ...Option) (types.RawDocument, error) {
	doc, err := parseRawDocument(r, config, options...)
	if err != nil {
		return types.RawDocument{}, err
	}
	attrs := types.AttributesWithOverrides{
		Content:   map[string]interface{}{},
		Overrides: map[string]string{},
		Counters:  map[string]interface{}{},
	}
	if doc.Blocks, err = processFileInclusions(doc.Blocks, attrs, []levelOffset{}, config, options...); err != nil {
		return types.RawDocument{}, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("raw document:")
		spew.Fdump(log.StandardLogger().Out, doc)
	}
	return doc, nil
}

func parseRawDocument(r io.Reader, config configuration.Configuration, options ...Option) (types.RawDocument, error) {
	log.Debugf("parsing raw document '%s'", config.Filename)
	if d, err := ParseReader(config.Filename, r, options...); err != nil {
		log.Errorf("failed to parse raw document: %s", err)
		return types.RawDocument{}, err
	} else if doc, ok := d.(types.RawDocument); ok {
		return doc, nil
	} else {
		return types.RawDocument{}, fmt.Errorf("unexpected type of raw document: '%T'", d)
	}
}

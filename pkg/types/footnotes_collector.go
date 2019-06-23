package types

import (
	log "github.com/sirupsen/logrus"
)

// Footnotes the footnotes of a document. Footnotes are "collected"
// during the parsing phase and displayed at the bottom of the document
// during the rendering.
type Footnotes []*Footnote

// FootnotesContainer interface for all types which may contain footnotes
type FootnotesContainer interface {
	Footnotes() (Footnotes, FootnoteReferences, error)
}

// IndexOf returns the index of the given note in the footnotes.
func (f Footnotes) IndexOf(note *Footnote) (int, bool) {
	for i, n := range f {
		if n.ID == note.ID {
			return i, true
		}
	}
	return -1, false
}

// FootnoteReferences some footnote have a ref to be re-used in the document
type FootnoteReferences map[string]*Footnote

// FootnotesCollector the visitor that traverses the whole document structure in search for footnotes
type FootnotesCollector struct {
	Footnotes          Footnotes
	FootnoteReferences FootnoteReferences
}

var _ Visitor = &FootnotesCollector{}

// NewFootnotesCollector initializes a new FootnotesCollector
func NewFootnotesCollector() *FootnotesCollector {
	return &FootnotesCollector{
		Footnotes:          make([]*Footnote, 0),
		FootnoteReferences: make(map[string]*Footnote),
	}
}

// Visit Implements Visitable#Visit()
func (c *FootnotesCollector) Visit(element Visitable) error {
	if note, ok := element.(*Footnote); ok {
		if len(note.Elements) > 0 { // a foot note with some content
			c.Footnotes = append(c.Footnotes, note)
			log.Debugf("indexed footnote to %d", len(c.Footnotes))
			if note.Ref != "" { // foot note has a reference for further usage
				c.FootnoteReferences[note.Ref] = note
			}
		}
	}
	return nil
}

package types

// Footnotes the footnotes of a document. Footnotes are "collected"
// during the parsing phase and displayed at the bottom of the document
// during the rendering.
type Footnotes []Footnote

// FootnotesContainer interface for all types which may contain footnotes
type FootnotesContainer interface {
	Footnotes() (Footnotes, FootnoteReferences, error)
}

// IndexOf returns the index of the given note in the footnotes.
func (f Footnotes) IndexOf(note Footnote) (int, bool) {
	for i, n := range f {
		if n.ID == note.ID {
			return i, true
		}
	}
	return -1, false
}

// FootnoteReferences some footnote have a ref to be re-used in the document
type FootnoteReferences map[string]Footnote

// FindFootnotes finds the footnotes in the given content. Also returns footnote references
// if applicable
func FindFootnotes(content interface{}) (Footnotes, FootnoteReferences, error) {
	footnotes := []Footnote{}
	footnoteRefs := map[string]Footnote{}
	switch c := content.(type) {
	case [][]interface{}:
		for _, e := range c {
			notes, refs, err := FindFootnotes(e)
			if err != nil {
				return footnotes, footnoteRefs, err
			}
			footnotes = append(footnotes, notes...)
			for ref, note := range refs {
				footnoteRefs[ref] = note
			}
		}
	case []interface{}:
		for _, e := range c {
			notes, refs, err := FindFootnotes(e)
			if err != nil {
				return footnotes, footnoteRefs, err
			}
			footnotes = append(footnotes, notes...)
			for ref, note := range refs {
				footnoteRefs[ref] = note
			}
		}
	case Footnote:
		if len(c.Elements) > 0 { // a foot note with some content
			footnotes = append(footnotes, c)
			if c.Ref != "" { // foot note has a reference for further usage
				footnoteRefs[c.Ref] = c
			}
		}
	}
	return footnotes, footnoteRefs, nil
}

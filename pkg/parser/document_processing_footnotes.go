package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/sirupsen/logrus"
)

// processes the footnotes in the blocks, replaces them with `FootnoteReference`
// and keep them in a separate `Footnotes`
func processFootnotes(blocks []interface{}) ([]interface{}, []types.Footnote) {
	logrus.Debug("processing footnotes...")
	footnotes := types.NewFootnotes()
	for i, block := range blocks {
		if c, ok := block.(types.FootnotesContainer); ok {
			blocks[i] = c.ReplaceFootnotes(footnotes)
		}
	}
	return blocks, footnotes.Notes()
}

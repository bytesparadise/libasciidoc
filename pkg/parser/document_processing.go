package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

const bufferSize = 10

// ParseDocument parses the content of the reader identitied by the filename and applies all the substitutions and arrangements
func ParseDocument(r io.Reader, config *configuration.Configuration, opts ...Option) (*types.Document, error) {
	done := make(chan interface{})
	defer close(done)

	ctx := NewParseContext(config, opts...) // each pipeline step will have its own clone of `ctx`
	footnotes := types.NewFootnotes()
	doc, _, err := Aggregate(ctx.Clone(),
		// SplitHeader(done,
		FilterOut(done,
			ArrangeLists(done,
				CollectFootnotes(footnotes, done,
					ApplySubstitutions(ctx.Clone(), done, // needs to be before 'ArrangeLists'
						ParseFragments(ctx.Clone(), r, done),
					),
				),
			),
		),
		// ),
	)
	if err != nil {
		return nil, err
	}
	if len(footnotes.Notes) > 0 {
		doc.Footnotes = footnotes.Notes
	}
	// if log.IsLevelEnabled(log.InfoLevel) {
	// 	log.Infof("parsed document:\n%s", spew.Sdump(doc))
	// }
	return doc, nil
}

package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

const bufferSize = 10

// ParseDocument parses the content of the reader identitied by the filename and applies all the substitutions and arrangements
func ParseDocument(r io.Reader, config *configuration.Configuration, opts ...Option) (*types.Document, error) {
	done := make(chan interface{})
	defer close(done)

	footnotes := types.NewFootnotes()
	doc, err := Aggregate(NewParseContext(config, opts...),
		// SplitHeader(done,
		FilterOut(done,
			ArrangeLists(done,
				CollectFootnotes(footnotes, done,
					ApplySubstitutions(NewParseContext(config, opts...), done, // needs to be before 'ArrangeLists'
						RefineFragments(NewParseContext(config, opts...), r, done,
							ParseFragments(NewParseContext(config, opts...), r, done),
						),
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
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsed document:\n%s", spew.Sdump(doc))
	}
	return doc, nil
}

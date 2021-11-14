package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

func CollectFootnotes(n *types.Footnotes, done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) chan types.DocumentFragment {
	processedFragmentStream := make(chan types.DocumentFragment, 10)
	go func() {
		defer close(processedFragmentStream)
		for f := range fragmentStream {
			select {
			case <-done:
				log.WithField("pipeline_stage", "collect_footnotes").Debug("received 'done' signal")
				return
			case processedFragmentStream <- collectFootnotes(n, f):
			}
		}
		log.WithField("pipeline_stage", "collect_footnotes").Debug("done")
	}()
	return processedFragmentStream
}

func collectFootnotes(n *types.Footnotes, f types.DocumentFragment) types.DocumentFragment {
	for _, e := range f.Elements {
		if e, ok := e.(types.WithFootnotes); ok {
			e.SubstituteFootnotes(n)
		}
	}
	return f
}

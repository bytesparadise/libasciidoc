package sgml

import "github.com/bytesparadise/libasciidoc/pkg/renderer"

// ContextualPipeline carries the renderer context along with
// the pipeline data to process in a template or in a nested template
type ContextualPipeline struct {
	Context *renderer.Context
	// The actual pipeline
	Data interface{}
}

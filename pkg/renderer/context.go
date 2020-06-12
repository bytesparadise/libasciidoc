package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Context is a custom implementation of the standard golang context.Context interface,
// which carries the types.Document which is being processed
type Context struct {
	Config configuration.Configuration
	// TableOfContents exists even if the document did not specify the `:toc:` attribute.
	// It will take into account the configured `:toclevels:` attribute value.
	TableOfContents      types.TableOfContents
	IncludeBlankLine     bool
	WithinDelimitedBlock bool
	WithinList           int
	counters             map[string]int
	Attributes           types.Attributes
	Footnotes            []types.Footnote
	ElementReferences    types.ElementReferences
	HasHeader            bool
}

// NewContext returns a new rendering context for the given document.
func NewContext(doc types.Document, config configuration.Configuration) *Context {
	_, hasHeader := doc.Header()
	return &Context{
		Config:            config,
		counters:          make(map[string]int),
		Attributes:        doc.Attributes,
		ElementReferences: doc.ElementReferences,
		Footnotes:         doc.Footnotes,
		HasHeader:         hasHeader,
	}
}

const tableCounter = "tableCounter"

// GetAndIncrementTableCounter returns the current value for the table counter after internally incrementing it.
func (ctx *Context) GetAndIncrementTableCounter() int {
	return ctx.getAndIncrementCounter(tableCounter)
}

const imageCounter = "imageCounter"

// GetAndIncrementImageCounter returns the current value for the image counter after internally incrementing it.
func (ctx *Context) GetAndIncrementImageCounter() int {
	return ctx.getAndIncrementCounter(imageCounter)
}

const exampleBlockCounter = "exampleBlockCounter"

// GetAndIncrementExampleBlockCounter returns the current value for the example block counter after internally incrementing it.
func (ctx *Context) GetAndIncrementExampleBlockCounter() int {
	return ctx.getAndIncrementCounter(exampleBlockCounter)
}

// getAndIncrementCounter returns the current value for the  counter after internally incrementing it.
func (ctx *Context) getAndIncrementCounter(name string) int {
	if _, found := ctx.counters[name]; !found {
		ctx.counters[name] = 1
		return 1
	}
	ctx.counters[name]++
	return ctx.counters[name]
}

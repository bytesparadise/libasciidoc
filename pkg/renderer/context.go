package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Context is a custom implementation of the standard golang context.Context interface,
// which carries the types.Document which is being processed
type Context struct {
	Config               *configuration.Configuration // TODO: use composition (remove the `Config` field)
	WithinDelimitedBlock bool
	WithinList           int
	counters             map[string]int
	Attributes           types.Attributes
	ElementReferences    types.ElementReferences
	HasHeader            bool
	SectionNumbering     types.SectionNumbers
}

// NewContext returns a new rendering context for the given document.
func NewContext(doc *types.Document, config *configuration.Configuration) *Context {
	header := doc.Header()
	ctx := &Context{
		Config:            config,
		counters:          make(map[string]int),
		Attributes:        config.Attributes,
		ElementReferences: doc.ElementReferences,
		HasHeader:         header != nil,
	}
	// TODO: add other attributes from https://docs.asciidoctor.org/asciidoc/latest/attributes/document-attributes-ref/#builtin-attributes-i18n
	ctx.Attributes[types.AttrFigureCaption] = "Figure"
	ctx.Attributes[types.AttrExampleCaption] = "Example"
	ctx.Attributes[types.AttrTableCaption] = "Table"
	ctx.Attributes[types.AttrVersionLabel] = "version"
	// also, expand authors and revision
	if header != nil {
		if authors := header.Authors(); authors != nil {
			ctx.Attributes.AddAll(authors.Expand())
		}

		if revision := header.Revision(); revision != nil {
			ctx.Attributes.AddAll(revision.Expand())

		}
	}
	return ctx
}

func (ctx *Context) UseUnicode() bool {
	return ctx.Attributes.GetAsBoolWithDefault(types.AttrUnicode, true)
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

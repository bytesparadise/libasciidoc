package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// context is a custom implementation of the standard golang context.context interface,
// which carries the types.Document which is being processed
type context struct {
	config               *configuration.Configuration // TODO: use composition (remove the `Config` field)
	withinDelimitedBlock bool
	withinList           int
	counters             map[string]int
	attributes           types.Attributes
	elementReferences    types.ElementReferences
	hasHeader            bool
	sectionNumbering     types.SectionNumbers
}

// newContext returns a new rendering context for the given document.
func newContext(doc *types.Document, config *configuration.Configuration) *context {
	header, _ := doc.Header()
	ctx := &context{
		config:            config,
		counters:          make(map[string]int),
		attributes:        config.Attributes,
		elementReferences: doc.ElementReferences,
		hasHeader:         header != nil,
	}
	// TODO: add other attributes from https://docs.asciidoctor.org/asciidoc/latest/attributes/document-attributes-ref/#builtin-attributes-i18n
	ctx.attributes[types.AttrFigureCaption] = "Figure"
	ctx.attributes[types.AttrExampleCaption] = "Example"
	ctx.attributes[types.AttrTableCaption] = "Table"
	ctx.attributes[types.AttrVersionLabel] = "version"
	// also, expand authors and revision
	if header != nil {
		if authors := header.Authors(); authors != nil {
			ctx.attributes.AddAll(authors.Expand())
		}

		if revision := header.Revision(); revision != nil {
			ctx.attributes.AddAll(revision.Expand())

		}
	}
	return ctx
}

func (ctx *context) UseUnicode() bool {
	return ctx.attributes.GetAsBoolWithDefault(types.AttrUnicode, true)
}

const tableCounter = "tableCounter"

// GetAndIncrementTableCounter returns the current value for the table counter after internally incrementing it.
func (ctx *context) GetAndIncrementTableCounter() int {
	return ctx.getAndIncrementCounter(tableCounter)
}

const imageCounter = "imageCounter"

// GetAndIncrementImageCounter returns the current value for the image counter after internally incrementing it.
func (ctx *context) GetAndIncrementImageCounter() int {
	return ctx.getAndIncrementCounter(imageCounter)
}

const exampleBlockCounter = "exampleBlockCounter"

// GetAndIncrementExampleBlockCounter returns the current value for the example block counter after internally incrementing it.
func (ctx *context) GetAndIncrementExampleBlockCounter() int {
	return ctx.getAndIncrementCounter(exampleBlockCounter)
}

// getAndIncrementCounter returns the current value for the  counter after internally incrementing it.
func (ctx *context) getAndIncrementCounter(name string) int {
	if _, found := ctx.counters[name]; !found {
		ctx.counters[name] = 1
		return 1
	}
	ctx.counters[name]++
	return ctx.counters[name]
}

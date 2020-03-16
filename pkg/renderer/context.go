package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	log "github.com/sirupsen/logrus"
)

// Context is a custom implementation of the standard golang context.Context interface,
// which carries the types.Document which is being processed
type Context struct {
	Document types.Document
	Config   configuration.Configuration
	// TableOfContents exists even if the document did not specify the `:toc:` attribute.
	// It will take into account the configured `:toclevels:` attribute value.
	TableOfContents types.TableOfContents
	// macros               map[string]MacroTemplate
	includeBlankLine     bool
	withinDelimitedBlock bool
	withinList           int
	counters             map[string]int
}

// NewContext returns a new rendering context for the given document.
func NewContext(document types.Document, config configuration.Configuration) Context {
	return Context{
		Document: document,
		Config:   config,
		counters: make(map[string]int),
		// macros:   make(map[string]MacroTemplate),
	}
}

// SetIncludeBlankLine sets the rendering context to include (or not) the blank lines
func (ctx *Context) SetIncludeBlankLine(b bool) bool {
	oldvalue := ctx.includeBlankLine
	ctx.includeBlankLine = b
	log.Debugf("set 'includeBlankLine' context param to '%t' (was '%t' before)", b, oldvalue)
	return oldvalue
}

// IncludeBlankLine indicates if blank lines should be rendered (default false)
func (ctx *Context) IncludeBlankLine() bool {
	return ctx.includeBlankLine
}

const withinDelimitedBlock string = "withinDelimitedBlock"

// SetWithinDelimitedBlock sets the rendering context to be within a delimited block
func (ctx *Context) SetWithinDelimitedBlock(b bool) bool {
	log.Debugf("set rendering elements within a delimited block to `%t`", b)
	oldvalue := ctx.withinDelimitedBlock
	log.Debugf("set '%s' context param to '%t' (was '%t' before)", withinDelimitedBlock, b, oldvalue)
	ctx.withinDelimitedBlock = b
	return oldvalue
}

// WithinDelimitedBlock indicates if the current element to render is within a delimited block or not
func (ctx *Context) WithinDelimitedBlock() bool {
	return ctx.withinDelimitedBlock
}

const withinList string = "withinList"

// SetWithinList sets the rendering context to be within a list or a nest list
func (ctx *Context) SetWithinList(w bool) {
	if w {
		ctx.withinList++
	} else {
		ctx.withinList--
	}
}

// WithinList indicates if the current element to render is within a list or not
func (ctx *Context) WithinList() bool {
	return ctx.withinList > 0
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

// //Option the options when rendering a document
// type Option func(ctx *Context)

// const (
// 	// keyLastUpdated the key to specify the last update of the document to render.
// 	// Can be a string or a time, which will be formatted using the 2006/01/02 15:04:05 MST` pattern
// 	keyLastUpdated string = types.AttrLastUpdated
// 	// keyIncludeHeaderFooter key to a bool value to indicate if the header and footer should be rendered
// 	keyIncludeHeaderFooter string = "IncludeHeaderFooter"
// 	// keyCSS key to the options CSS to add in the document head. Default is empty ("")
// 	keyCSS string = "CSS"
// 	// keyEntrypoint key to the entrypoint to start with when parsing the document
// 	keyEntrypoint string = "Entrypoint"
// 	// LastUpdatedFormat key to the time format for the `last updated` document attribute
// 	LastUpdatedFormat string = "2006-01-02 15:04:05 -0700"
// )

// // LastUpdated function to set the `last updated` option in the renderer context (default is `time.Now()`)
// func WithLastUpdated(value time.Time) Option {
// 	return func(ctx *Context) {
// 		ctx.options[keyLastUpdated] = value
// 	}
// }

// // IncludeHeaderFooter function to set the `include header/footer` option in the renderer context
// func WithHeaderFooter(value bool) Option {
// 	return func(ctx *Context) {
// 		ctx.options[keyIncludeHeaderFooter] = value
// 	}
// }

// // IncludeCSS function to set the `css` option in the renderer context
// func WithCSS(href string) Option {
// 	return func(ctx *Context) {
// 		ctx.options[keyCSS] = href
// 	}
// }

// // Entrypoint function to set the `entrypoint` option in the renderer context
// func Entrypoint(entrypoint string) Option {
// 	return func(ctx *Context) {
// 		ctx.options[keyEntrypoint] = entrypoint
// 	}
// }

// // DefineMacro defines the given template to a user macro with the given name
// func WithMacroTemplate(name string, t MacroTemplate) Option {
// 	return func(ctx *Context) {
// 		ctx.macros[name] = t
// 	}
// }

// // LastUpdated returns the value of the 'LastUpdated' Option if it was present,
// // otherwise it returns the current time using the `2006/01/02 15:04:05 MST` format
// func (ctx *Context) WithLastUpdated() string {
// 	if lastUpdated, ok := ctx.options[keyLastUpdated].(time.Time); ok {
// 		return lastUpdated.Format(LastUpdatedFormat)
// 	}
// 	return time.Now().Format(LastUpdatedFormat)
// }

// IncludeHeaderFooter returns the value of the 'IncludeHeaderFooter' Option if it was present,
// otherwise it returns `false`
// func (ctx *Context) WithHeaderFooter() bool {
// 	return ctx.IncludeHeaderFooter
// }

// // CSS returns the value of the 'CSS' Option if it was present,
// // otherwise it returns an empty string
// func (ctx *Context) CSS() string {
// 	if css, ok := ctx.options[keyCSS].(string); ok {
// 		return css
// 	}
// 	return ""
// }

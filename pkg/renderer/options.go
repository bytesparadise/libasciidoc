package renderer

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

//Option the options when rendering a document
type Option func(ctx *Context)

const (
	// keyLastUpdated the key to specify the last update of the document to render.
	// Can be a string or a time, which will be formatted using the 2006/01/02 15:04:05 MST` pattern
	keyLastUpdated string = types.AttrLastUpdated
	// keyIncludeHeaderFooter key to a bool value to indicate if the header and footer should be rendered
	keyIncludeHeaderFooter string = "IncludeHeaderFooter"
	// keyCSS key to the options CSS to add in the document head. Default is empty ("")
	keyCSS string = "CSS"
	// keyEntrypoint key to the entrypoint to start with when parsing the document
	keyEntrypoint string = "Entrypoint"
	// LastUpdatedFormat key to the time format for the `last updated` document attribute
	LastUpdatedFormat string = "2006-01-02 15:04:05 -0700"
)

// LastUpdated function to set the `last updated` option in the renderer context (default is `time.Now()`)
func LastUpdated(value time.Time) Option {
	return func(ctx *Context) {
		ctx.options[keyLastUpdated] = value
	}
}

// IncludeHeaderFooter function to set the `include header/footer` option in the renderer context
func IncludeHeaderFooter(value bool) Option {
	return func(ctx *Context) {
		ctx.options[keyIncludeHeaderFooter] = value
	}
}

// IncludeCSS function to set the `css` option in the renderer context
func IncludeCSS(href string) Option {
	return func(ctx *Context) {
		ctx.options[keyCSS] = href
	}
}

// Entrypoint function to set the `entrypoint` option in the renderer context
func Entrypoint(entrypoint string) Option {
	return func(ctx *Context) {
		ctx.options[keyEntrypoint] = entrypoint
	}
}

// DefineMacro defines the given template to a user macro with the given name
func DefineMacro(name string, t MacroTemplate) Option {
	return func(ctx *Context) {
		ctx.macros[name] = t
	}
}

// LastUpdated returns the value of the 'LastUpdated' Option if it was present,
// otherwise it returns the current time using the `2006/01/02 15:04:05 MST` format
func (ctx *Context) LastUpdated() string {
	if lastUpdated, ok := ctx.options[keyLastUpdated].(time.Time); ok {
		return lastUpdated.Format(LastUpdatedFormat)
	}
	return time.Now().Format(LastUpdatedFormat)
}

// IncludeHeaderFooter returns the value of the 'IncludeHeaderFooter' Option if it was present,
// otherwise it returns `false`
func (ctx *Context) IncludeHeaderFooter() bool {
	if includeHeaderFooter, ok := ctx.options[keyIncludeHeaderFooter].(bool); ok {
		return includeHeaderFooter
	}
	return false
}

// CSS returns the value of the 'CSS' Option if it was present,
// otherwise it returns an empty string
func (ctx *Context) CSS() string {
	if css, ok := ctx.options[keyCSS].(string); ok {
		return css
	}
	return ""
}

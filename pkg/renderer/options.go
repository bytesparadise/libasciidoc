package renderer

import "time"

//Option the options when rendering a document
type Option func(ctx *Context)

const (
	//keyLastUpdated the key to specify the last update of the document to render.
	// Can be a string or a time, which will be formatted using the 2006/01/02 15:04:05 MST` pattern
	keyLastUpdated string = "LastUpdated"
	//keyIncludeHeaderFooter a bool value to indicate if the header and footer should be rendered
	keyIncludeHeaderFooter string = "IncludeHeaderFooter"
	//keyEntrypoint a bool value to indicate if the entrypoint to start with when parsing the document
	keyEntrypoint string = "Entrypoint"
	// LastUpdatedFormat the time format for the `last updated` document attribute
	LastUpdatedFormat string = "2006/01/02 15:04:05 MST"
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

// Entrypoint function to set the `entrypoint` option in the renderer context
func Entrypoint(entrypoint string) Option {
	return func(ctx *Context) {
		ctx.options[keyEntrypoint] = entrypoint
	}
}

// LastUpdated returns the value of the 'LastUpdated' Option if it was present,
// otherwise it returns the current time using the `2006/01/02 15:04:05 MST` format
func (ctx *Context) LastUpdated() string {
	if lastUpdated, found := ctx.options[keyLastUpdated]; found {
		if lastUpdated, typeMatch := lastUpdated.(time.Time); typeMatch {
			return lastUpdated.Format(LastUpdatedFormat)
		}
	}
	return time.Now().Format(LastUpdatedFormat)
}

// IncludeHeaderFooter returns the value of the 'LastUpdated' Option if it was present,
// otherwise it returns `false`
func (ctx *Context) IncludeHeaderFooter() bool {
	if includeHeaderFooter, found := ctx.options[keyIncludeHeaderFooter]; found {
		if includeHeaderFooter, typeMatch := includeHeaderFooter.(bool); typeMatch {
			return includeHeaderFooter
		}
	}
	return false
}

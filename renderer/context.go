package renderer

import (
	"context"
	"time"

	"github.com/bytesparadise/libasciidoc/types"
	log "github.com/sirupsen/logrus"
)

// Context is a custom implementation of the standard golang context.Context interface,
// which carries the types.Document which is being processed
type Context struct {
	context  context.Context
	Document types.Document
	options  map[string]interface{}
}

// Wrap wraps the given `ctx` context into a new context which will contain the given `document` document.
func Wrap(ctx context.Context, document types.Document, options ...Option) *Context {
	result := &Context{
		context:  ctx,
		Document: document,
		options:  make(map[string]interface{}),
	}
	for _, option := range options {
		option(result)
	}
	return result
}

const includeBlankLine string = "includeBlankLine"

// SetIncludeBlankLine sets the rendering context to include (or not) the blank lines
func (ctx *Context) SetIncludeBlankLine(b bool) {
	ctx.options[includeBlankLine] = b
}

// IncludeBlankLine indicates if blank lines should be rendered (default false)
func (ctx *Context) IncludeBlankLine() bool {
	if b, found := ctx.options[includeBlankLine].(bool); found {
		return b
	}
	// by default, ignore blank lines
	return false
}

const withinDelimitedBlock string = "withinDelimitedBlock"

// SetWithinDelimitedBlock sets the rendering context to be within a delimited block
func (ctx *Context) SetWithinDelimitedBlock(b bool) {
	log.Debugf("set rendering elements within a delimited block to `%t`", b)
	ctx.options[withinDelimitedBlock] = b
}

// WithinDelimitedBlock indicates if the current element to render is within a delimited block or not
func (ctx *Context) WithinDelimitedBlock() bool {
	if b, found := ctx.options[withinDelimitedBlock].(bool); found {
		log.Debugf("rendering elements within a delimited block? %t", b)
		return b
	}
	// by default, ignore blank lines
	return false
}

const withinList string = "withinList"

// SetWithinList sets the rendering context to be within a list or a nest list
func (ctx *Context) SetWithinList(w bool) {
	log.Debugf("set rendering elements within a list to `%t`", w)
	var counter int
	var ok bool
	if counter, ok = ctx.options[withinList].(int); ok {
		// keep track of the depth of the list
		if w {
			counter++
		} else {
			counter--
		}
	} else {
		if w {
			counter = 1
		} else {
			counter = 0
		}
	}
	// update the counter in the context
	ctx.options[withinList] = counter
}

// WithinList indicates if the current element to render is within a list or not
func (ctx *Context) WithinList() bool {
	if counter, found := ctx.options[withinList].(int); found {
		log.Debugf("rendering elements within a list? %t (%d)", (counter > 0), counter)
		return counter > 0
	}
	// by default, ignore blank lines
	return false
}

// Deadline wrapper implementation of context.Context.Deadline()
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.context.Deadline()
}

// Done wrapper implementation of context.Context.Done()
func (ctx *Context) Done() <-chan struct{} {
	return ctx.context.Done()
}

// Err wrapper implementation of context.Context.Err()
func (ctx *Context) Err() error {
	return ctx.context.Err()
}

// Value wrapper implementation of context.Context.Value(interface{})
func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.context.Value(key)
}

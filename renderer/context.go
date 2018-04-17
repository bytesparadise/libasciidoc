package renderer

import (
	"context"
	"time"

	"github.com/bytesparadise/libasciidoc/types"
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

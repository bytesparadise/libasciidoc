package renderer

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// MacroTemplate an interface of template for user macro.
type MacroTemplate interface {
	Execute(wr io.Writer, data interface{}) error
}

// Context is a custom implementation of the standard golang context.Context interface,
// which carries the types.Document which is being processed
type Context struct {
	context  context.Context
	Document types.Document
	options  map[string]interface{}
	macros   map[string]MacroTemplate
}

// Wrap wraps the given `ctx` context into a new context which will contain the given `document` document.
func Wrap(ctx context.Context, document types.Document, options ...Option) *Context {
	result := &Context{
		context:  ctx,
		Document: document,
		options:  make(map[string]interface{}),
		macros:   make(map[string]MacroTemplate),
	}
	for _, option := range options {
		option(result)
	}
	return result
}

const includeBlankLine string = "includeBlankLine"

// SetIncludeBlankLine sets the rendering context to include (or not) the blank lines
func (ctx *Context) SetIncludeBlankLine(b bool) bool {
	var oldvalue bool
	if v, ok := ctx.options[includeBlankLine].(bool); ok {
		oldvalue = v
	}
	ctx.options[includeBlankLine] = b
	log.Debugf("set '%s' context param to '%t' (was '%t' before)", includeBlankLine, b, oldvalue)
	return oldvalue
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
func (ctx *Context) SetWithinDelimitedBlock(b bool) bool {
	log.Debugf("set rendering elements within a delimited block to `%t`", b)
	var oldvalue bool
	if v, ok := ctx.options[withinDelimitedBlock].(bool); ok {
		oldvalue = v
	}
	log.Debugf("set '%s' context param to '%t' (was '%t' before)", withinDelimitedBlock, b, oldvalue)
	ctx.options[withinDelimitedBlock] = b
	return oldvalue
}

// WithinDelimitedBlock indicates if the current element to render is within a delimited block or not
func (ctx *Context) WithinDelimitedBlock() bool {
	if b, found := ctx.options[withinDelimitedBlock].(bool); found {
		log.Debugf("rendering elements within a delimited block? %t", b)
		return b
	}
	// by default, consider not within a block
	return false
}

const withinList string = "withinList"

// SetWithinList sets the rendering context to be within a list or a nest list
func (ctx *Context) SetWithinList(w bool) {
	// log.Debugf("set rendering elements within a list to `%t`", w)
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
		// log.Debugf("rendering elements within a list? %t (%d)", (counter > 0), counter)
		return counter > 0
	}
	// by default, ignore blank lines
	return false
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
func (ctx *Context) getAndIncrementCounter(counter string) int {
	if _, found := ctx.options[counter]; !found {
		ctx.options[counter] = 1
	}
	if c, ok := ctx.options[counter].(int); ok {
		ctx.options[counter] = c + 1
		return c
	}
	ctx.options[counter] = 1
	log.Warnf("'%s' counter was set to a non-int value", counter)
	return 1
}

const imagesdir = "imagesdir"

// GetImagesDir returns the value of the `imagesdir` attribute if it was set (as a string), empty string otherwise
func (ctx *Context) GetImagesDir() string {
	if imagesdir, found := ctx.Document.Attributes.GetAsString(imagesdir); found {
		return imagesdir
	}
	return ""
}

// MacroTemplate finds and returns a user macro function by specified name.
func (ctx *Context) MacroTemplate(name string) (MacroTemplate, error) {
	macro, ok := ctx.macros[name]
	if ok {
		return macro, nil
	}
	return nil, errors.New("unknown user macro: " + name)
}

// -----------------------
// context.Context methods
// -----------------------

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

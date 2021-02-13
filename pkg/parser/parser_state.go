package parser

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// initPositionalIndex sets the `types.AttrPositionalIndex` value to 0 in the current state
func initPositionalIndex(c *current) error {
	// TODO: use current state instead of global state?
	c.globalStore[types.AttrPositionalIndex] = 0
	return nil
}

// incrementPositionalIndex increments the value of `types.AttrPositionalIndex`
// returns the current index (after increment) or an error if the value is not an `int`
func incrementPositionalIndex(c *current) (int, error) {
	// TODO: use current state instead of global state?
	p, ok := c.globalStore[types.AttrPositionalIndex].(int)
	if !ok {
		return 0, fmt.Errorf("unexpected kind attribute positional index: '%T'", c.globalStore[types.AttrPositionalPrefix])
	}
	p = p + 1
	c.globalStore[types.AttrPositionalIndex] = p
	return p, nil
}

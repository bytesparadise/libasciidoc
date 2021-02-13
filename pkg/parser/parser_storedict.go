package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
)

// extra methods on the generated parser's `storeDict` type

const attributesKey = "attributes"

const usermacrosKey = "user_macros"

func (c storeDict) pushAttributes(value interface{}) {
	if s, ok := c[attributesKey].(*stack); ok {
		s.push(value)
		return
	}
	s := newStack()
	s.push(value)
	c[attributesKey] = s
}

func (c storeDict) discardAttributes() {
	if s, ok := c[attributesKey].(*stack); ok {
		s.pop()
	}
}

func (c storeDict) getAttributes() interface{} {
	if s, ok := c[attributesKey].(*stack); ok {
		return s.get()
	}
	return nil
}

func (c storeDict) hasUserMacro(name string) bool {
	if macros, exists := c[usermacrosKey].(map[string]configuration.MacroTemplate); exists {
		_, exists := macros[name]
		return exists
	}
	return false
}

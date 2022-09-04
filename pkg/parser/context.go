package parser

import (
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

type ParseContext struct {
	filename     string
	opts         []Option
	levelOffsets levelOffsets
	attributes   *contextAttributes
	userMacros   map[string]configuration.MacroTemplate
	counters     map[string]interface{}
}

func NewParseContext(config *configuration.Configuration, options ...Option) *ParseContext {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new parser context with attributes: %s", spew.Sdump(config.Attributes))
	}
	opts := []Option{
		Entrypoint("DocumentFragment"),
		GlobalStore(frontMatterKey, true),
		GlobalStore(documentHeaderKey, true),
		GlobalStore(usermacrosKey, config.Macros),
		GlobalStore(enabledSubstitutionsKey, attributeDeclarations()),
	}
	opts = append(opts, options...)
	return &ParseContext{
		filename:     config.Filename,
		opts:         opts,
		levelOffsets: []*levelOffset{},
		attributes:   newContextAttributes(config.Attributes),
		userMacros:   config.Macros,
		counters:     map[string]interface{}{},
	}
}

func (c *ParseContext) Clone() *ParseContext {
	return &ParseContext{
		filename:     c.filename,
		opts:         options(c.opts).clone(),
		levelOffsets: c.levelOffsets.clone(),
		attributes:   c.attributes.clone(),
		userMacros:   c.userMacros,
		counters:     c.counters,
	}
}

type options []Option

func (o options) clone() []Option {
	result := make([]Option, len(o))
	copy(result, o)
	return result
}

type contextAttributes struct {
	immutableAttributes types.Attributes
	attributes          types.Attributes
	mutex               *sync.RWMutex
}

func newContextAttributes(attrs types.Attributes) *contextAttributes {
	return &contextAttributes{
		immutableAttributes: attrs,
		attributes:          types.Attributes{},
		mutex:               &sync.RWMutex{},
	}
}

func (a *contextAttributes) clone() *contextAttributes {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return &contextAttributes{
		mutex:               &sync.RWMutex{},
		immutableAttributes: a.immutableAttributes.Clone(),
		attributes:          a.attributes.Clone(),
	}
}

func (a *contextAttributes) allAttributes() map[string]interface{} {
	result := make(map[string]interface{}, len(a.attributes)+len(a.immutableAttributes))
	for k, v := range a.attributes {
		result[k] = v
	}
	// imautables attributes should not be overridden, hence adding them after
	for k, v := range a.immutableAttributes {
		result[k] = v
	}
	return result
}

func (a *contextAttributes) get(k string) (interface{}, bool) {
	if v, found := a.immutableAttributes[k]; found {
		return v, true
	}
	v, found := a.attributes[k]
	return v, found
}

func (a *contextAttributes) getAsIntWithDefault(k string, defaultValue int) int {
	if a.immutableAttributes.Has(k) {
		return a.immutableAttributes.GetAsIntWithDefault(k, defaultValue)
	}
	return a.attributes.GetAsIntWithDefault(k, defaultValue)
}

func (a *contextAttributes) set(k string, v interface{}) {
	a.mutex.RLock() // TODO: needed? each go routine has its own context
	defer a.mutex.RUnlock()
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("setting context attribute: %s -> %s", k, spew.Sdump(v))
	}
	a.attributes[k] = v
}

func (a *contextAttributes) unset(k string) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	delete(a.attributes, k)
}

func (a *contextAttributes) setAll(attrs map[string]interface{}) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	for k, v := range attrs {
		a.attributes[k] = v
	}
}

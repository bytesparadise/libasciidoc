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
	Opts         []Option
	levelOffsets levelOffsets
	attributes   *contextAttributes
	userMacros   map[string]configuration.MacroTemplate
	counters     map[string]interface{}
}

func NewParseContext(config *configuration.Configuration, opts ...Option) *ParseContext {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new parser context with attributes: %s", spew.Sdump(config.Attributes))
	}
	opts = append(opts, Entrypoint("DocumentFragment"))
	opts = append(opts, GlobalStore(frontMatterKey, true))
	opts = append(opts, GlobalStore(documentHeaderKey, true))
	opts = append(opts, GlobalStore(substitutionKey, newNormalSubstitution()))
	opts = append(opts, GlobalStore(usermacrosKey, config.Macros))
	return &ParseContext{
		filename:     config.Filename,
		Opts:         opts,
		levelOffsets: []*levelOffset{},
		attributes:   newContextAttributes(config.Attributes),
		userMacros:   config.Macros,
		counters:     map[string]interface{}{},
	}
}

func (c *ParseContext) Clone() *ParseContext {
	opts := make([]Option, len(c.Opts))
	copy(opts, c.Opts)
	attributes := c.attributes.clone()
	return &ParseContext{
		filename:     c.filename,
		Opts:         opts,
		levelOffsets: c.levelOffsets.clone(),
		attributes:   attributes,
		userMacros:   c.userMacros,
		counters:     c.counters,
	}
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

func (a *contextAttributes) has(k string) bool {
	_, found := a.immutableAttributes[k]
	if found {
		return true
	}
	_, found = a.attributes[k]
	return found
}

func (a *contextAttributes) get(k string) interface{} {
	value, found := a.immutableAttributes[k]
	if found {
		return value
	}
	return a.attributes[k]
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

func (a *contextAttributes) getAsString(k string) (string, bool, error) {
	if a.immutableAttributes.Has(k) {
		return a.immutableAttributes.GetAsString(k)
	}
	return a.attributes.GetAsString(k)
}

func (a *contextAttributes) getAsStringWithDefault(k string, defaultValue string) string {
	if a.immutableAttributes.Has(k) {
		return a.immutableAttributes.GetAsStringWithDefault(k, defaultValue)
	}
	return a.attributes.GetAsStringWithDefault(k, defaultValue)
}

func (a *contextAttributes) getAsIntWithDefault(k string, defaultValue int) int {
	if a.immutableAttributes.Has(k) {
		return a.immutableAttributes.GetAsIntWithDefault(k, defaultValue)
	}
	return a.attributes.GetAsIntWithDefault(k, defaultValue)
}

func (a *contextAttributes) set(k string, v interface{}) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
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

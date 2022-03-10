package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"plugin"
)

// TODO: Add other hooks? PreProcess, PostRender, etc.
type PreRenderFunc func(*types.Document) (*types.Document, error)

type Plugin struct {
	Path      string
	PreRender PreRenderFunc
}

// takes a list of paths, opens the plugins at those paths, checks the sigs of
// the supported exported functions, and returns a list of plugins that can be
// stored in the configuration
func LoadPlugins(pluginPaths []string) ([]Plugin, error) {
	loadedPlugins := []Plugin{}
	for _, path := range pluginPaths {
		log.Debugf("plugins: loading %s\n", path)
		// make sure we can open it
		p, err := plugin.Open(path)
		if err != nil {
			return nil, err
		}
		newPlugin := Plugin{}
		newPlugin.Path = path
		// right now only PreRender is supported
		for _, name := range []string{"PreRender"} {
			log.Debugf("plugins: checking %s for %s function\n", path, name)
			// check if it has a supported function
			symbol, err := p.Lookup(name)
			if err != nil {
				log.Debugf("plugins: %s function not found\n", name)
				continue
			}
			// if the function signature is what we expect, add it
			switch name {
			case "PreRender":
				// NOTE: Lookup will give use a _pointer_ to the function
				addFunc, ok := symbol.(*PreRenderFunc)
				if !ok {
					return nil, errors.New("Invalid function signature for PreRender")
				}
				newPlugin.PreRender = *addFunc
				log.Debugf("plugins: added %s for %s function\n", path, name)
			}
		}
		loadedPlugins = append(loadedPlugins, newPlugin)
	}
	return loadedPlugins, nil
}

// runs the PreRender plugins
func RunPreRender(doc *types.Document, plugins []Plugin) (*types.Document, error) {
	for _, curPlugin := range plugins {
		log.Debugf("plugins: running %s PreRender\n", curPlugin.Path)
		var err error
		doc, err = curPlugin.PreRender(doc)
		if err != nil {
			return nil, err
		}
	}
	return doc, nil
}

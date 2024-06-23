package configuration

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// NewConfiguration returns a new configuration
func NewConfiguration(settings ...Setting) *Configuration {
	config := &Configuration{
		Attributes: map[string]interface{}{
			"basebackend-html": true, // along with default backend, to support `ifdef::basebackend-html` conditionals out-of-the-box
		},
		BackEnd: "html5", // default backend
		Macros:  map[string]MacroTemplate{},
	}
	// default backed
	WithBackEnd("html5")(config)
	// custom settings
	for _, set := range settings {
		set(config)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new configuration: %s", spew.Sdump(config))
	}
	return config
}

// Configuration the configuration used when rendering a document
type Configuration struct {
	Filename              string // TODO: move out of Configuration?
	Attributes            types.Attributes
	LastUpdated           time.Time
	WrapInHTMLBodyElement bool // flag to include the content in an html>body element
	CSS                   []string
	BackEnd               string
	Macros                map[string]MacroTemplate
}

const (
	// LastUpdatedFormat key to the time format for the `last updated` document attribute
	LastUpdatedFormat string = "2006-01-02 15:04:05 -0700"
)

// Setting a setting to customize the configuration used during parsing and rendering of a document
type Setting func(config *Configuration)

// WithFigureCaption function to set the `figure-caption` attribute
func WithFigureCaption(caption string) Setting {
	return func(config *Configuration) {
		config.Attributes[types.AttrFigureCaption] = caption
	}
}

// WithLastUpdated function to set the `last updated` option in the renderer context (default is `time.Now()`)
func WithLastUpdated(value time.Time) Setting {
	return func(config *Configuration) {
		config.LastUpdated = value
	}
}

// WithAttributes function to set the `attribute overrides`
func WithAttributes(attrs map[string]interface{}) Setting {
	return func(config *Configuration) {
		config.Attributes = attrs
	}
}

// WithAttribute function to set an attribute as if it was passed as an argument in the CLI
func WithAttribute(key string, value interface{}) Setting {
	return func(config *Configuration) {
		config.Attributes[key] = value
	}
}

// WithHeaderFooter function to set the `include header/footer` setting in the config
func WithHeaderFooter(value bool) Setting {
	return func(config *Configuration) {
		config.WrapInHTMLBodyElement = value
	}
}

// WithCSS function to set the `css` setting in the config
func WithCSS(hrefs []string) Setting {
	return func(config *Configuration) {
		config.CSS = hrefs
	}
}

// WithBackEnd sets the backend format, valid values are "html", "html5", "xhtml", "xhtml5", and "" (defaults to html5)
func WithBackEnd(backend string) Setting {
	return func(config *Configuration) {
		config.Attributes.Set("backend", backend)
		switch backend {
		case "html", "html5", "xhtml", "xhtml5":
			config.Attributes.Set("basebackend-html", true)
		default:
			config.Attributes.Unset("basebackend-html")
		}
		config.BackEnd = backend
		switch backend {
		case "html", "html5", "xhtml", "xhtml5":
			config.Attributes.Set("basebackend-html", true)
		default:
			config.Attributes.Unset("basebackend-html")
		}
	}
}

// WithFilename function to set the `filename` setting in the config
func WithFilename(filename string) Setting {
	return func(config *Configuration) {
		config.Filename = filename
	}
}

// WithMacroTemplate defines the given template to a user macro with the given name
func WithMacroTemplate(name string, t MacroTemplate) Setting {
	return func(config *Configuration) {
		log.Debugf("registering user macro '%s'", name)
		config.Macros[name] = t
	}
}

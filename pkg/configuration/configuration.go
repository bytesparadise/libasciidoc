package configuration

import (
	"time"
)

// NewConfiguration returns a new configuration
func NewConfiguration(settings ...Setting) Configuration {
	config := Configuration{
		AttributeOverrides: map[string]string{},
		Macros:             map[string]MacroTemplate{},
	}
	for _, set := range settings {
		set(&config)
	}
	return config
}

// Configuration the configuration used when rendering a document
type Configuration struct {
	Filename           string
	AttributeOverrides map[string]string
	LastUpdated        time.Time
	// WrapInHTMLBodyElement flag to include the content in an html>body element
	WrapInHTMLBodyElement bool
	CSS                   string
	BackEnd               string
	Macros                map[string]MacroTemplate
}

const (
	// LastUpdatedFormat key to the time format for the `last updated` document attribute
	LastUpdatedFormat string = "2006-01-02 15:04:05 -0700"
)

// Setting a setting to customize the configuration used during parsing and rendering of a document
type Setting func(config *Configuration)

// WithLastUpdated function to set the `last updated` option in the renderer context (default is `time.Now()`)
func WithLastUpdated(value time.Time) Setting {
	return func(config *Configuration) {
		config.LastUpdated = value
	}
}

// WithAttributes function to set the `attribute overrides`
func WithAttributes(attrs map[string]string) Setting {
	return func(config *Configuration) {
		config.AttributeOverrides = attrs
	}
}

// WithAttribute function to set an attribute as if it was passed as an argument in the CLI
func WithAttribute(key, value string) Setting {
	return func(config *Configuration) {
		config.AttributeOverrides[key] = value
	}
}

// WithHeaderFooter function to set the `include header/footer` setting in the config
func WithHeaderFooter(value bool) Setting {
	return func(config *Configuration) {
		config.WrapInHTMLBodyElement = value
	}
}

// WithCSS function to set the `css` setting in the config
func WithCSS(href string) Setting {
	return func(config *Configuration) {
		config.CSS = href
	}
}

// WithBackEnd sets the backend format, valid values are "html", "html5", "xhtml", "xhtml5", and "" (defaults to html5)
func WithBackEnd(backend string) Setting {
	return func(config *Configuration) {
		config.BackEnd = backend
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
		config.Macros[name] = t
	}
}

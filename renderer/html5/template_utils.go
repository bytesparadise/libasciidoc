package html5

import (
	htmltemplate "html/template"
	texttemplate "text/template"

	log "github.com/sirupsen/logrus"
)

func newHTMLTemplate(name, src string) *htmltemplate.Template {
	t, err := htmltemplate.New(name).Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

func newTextTemplate(name, src string) *texttemplate.Template {
	t, err := texttemplate.New(name).Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

package html5

import (
	htmltemplate "html/template"
	texttemplate "text/template"

	log "github.com/sirupsen/logrus"
)

func newHTMLTemplate(name, src string, funcs ...htmltemplate.FuncMap) *htmltemplate.Template {
	t := htmltemplate.New(name)
	for _, f := range funcs {
		t.Funcs(f)
	}
	t, err := t.Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

func newTextTemplate(name, src string, funcs ...texttemplate.FuncMap) *texttemplate.Template {
	t := texttemplate.New(name)
	for _, f := range funcs {
		t.Funcs(f)
	}
	t, err := t.Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

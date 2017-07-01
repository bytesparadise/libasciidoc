package html5

import (
	"html/template"

	log "github.com/sirupsen/logrus"
)

func newTemplate(name, src string) *template.Template {
	t, err := template.New(name).Parse(src)
	if err != nil {
		log.Fatalf("failed to initialize '%s' template: %s", name, err.Error())
	}
	return t
}

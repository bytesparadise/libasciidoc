package sgml

import (
	"strings"
)

func (r *sgmlRenderer) renderLineBreak() (string, error) {
	buf := &strings.Builder{}
	tmpl, err := r.lineBreak()
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *sgmlRenderer) renderThematicBreak() (string, error) {
	buf := &strings.Builder{}
	tmpl, err := r.thematicBreak()
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

package sgml

import (
	"strings"
)

func (r *sgmlRenderer) renderLineBreak() (string, error) {
	buf := &strings.Builder{}
	if err := r.lineBreak.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *sgmlRenderer) renderThematicBreak() (string, error) {
	buf := &strings.Builder{}
	if err := r.thematicBreak.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLink(ctx *renderer.Context, l types.InlineLink) (string, error) { //nolint: unparam
	result := &strings.Builder{}
	location := l.Location.String()
	text := ""
	class := ""
	var err error
	// TODO; support `mailto:` positional attributes
	positionals := l.Attributes.Positionals()
	if len(positionals) > 0 {
		buf := &strings.Builder{}
		for i, arg := range positionals {
			t, err := r.renderInlineElements(ctx, arg)
			if err != nil {
				return "", errors.Wrap(err, "unable to render external link")
			}
			buf.WriteString(t)
			if i < len(positionals)-1 {
				buf.WriteString(",")
			}
		}
		text = buf.String()
	} else {
		class = "bare"
		text = location
	}
	err = r.link.Execute(result, struct {
		URL   string
		Text  string
		Class string
	}{
		URL:   location,
		Text:  text,
		Class: class,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render external link")
	}
	log.Debugf("rendered external link: %s", result.String())
	return result.String(), nil
}

package sgml

import (
	"html"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLink(ctx *renderer.Context, l types.InlineLink) (string, error) { //nolint: unparam
	result := &strings.Builder{}
	location := l.Location.Stringify()
	text := ""
	class := ""
	// TODO; support `mailto:` positional attributes
	if t, exists := l.Attributes[types.AttrInlineLinkText]; exists {
		switch t := t.(type) {
		case string:
			text = t
		case []interface{}:
			var err error
			if text, err = r.renderInlineElements(ctx, t); err != nil {
				return "", errors.Wrap(err, "unable to render link")
			}
		}
	} else {
		class = "bare"
		text = html.EscapeString(location)
	}

	err := r.link.Execute(result, struct {
		URL   string
		Text  string
		Class string
	}{
		URL:   location,
		Text:  text,
		Class: class,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render link")
	}
	log.Debugf("rendered link: %s", result.String())
	return result.String(), nil
}

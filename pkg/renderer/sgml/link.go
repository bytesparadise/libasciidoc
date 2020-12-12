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
	// positionals := l.Attributes.Positionals()
	// if len(positionals) > 0 {
	// 	buf := &strings.Builder{}
	// 	for i, arg := range positionals {
	// 		t, err := r.renderInlineElements(ctx, arg)
	// 		if err != nil {
	// 			return "", errors.Wrap(err, "unable to render link")
	// 		}
	// 		buf.WriteString(t)
	// 		if i < len(positionals)-1 {
	// 			buf.WriteString(",")
	// 		}
	// 	}
	// 	text = buf.String()
	// } else {
	// 	class = "bare"
	// 	text = html.EscapeString(location)
	// }
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

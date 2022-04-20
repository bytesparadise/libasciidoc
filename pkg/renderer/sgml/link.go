package sgml

import (
	"html"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLink(ctx *renderer.Context, l *types.InlineLink) (string, error) {
	result := &strings.Builder{}
	location := l.Location.Stringify()
	text := ""
	class := ""
	id := l.Attributes.GetAsStringWithDefault(types.AttrID, "")
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render link")
	}
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
		class = roles // can be empty (and it's fine)
	} else {
		text = html.EscapeString(location)
		if len(roles) > 0 {
			class = "bare " + roles
		} else if len(text) > 0 { // keep empty 'class' when there's no location
			class = "bare"
		}
	}
	target := l.Attributes.GetAsStringWithDefault(types.AttrInlineLinkTarget, "")
	noopener := target == "_blank" || l.Attributes.HasOption("noopener")
	err = r.link.Execute(result, struct {
		ID       string
		URL      string
		Text     string
		Class    string
		Target   string
		NoOpener bool
	}{
		ID:       id,
		URL:      location,
		Text:     text,
		Class:    class,
		Target:   target,
		NoOpener: noopener,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render link")
	}
	log.Debugf("rendered link: %s", result.String())
	return result.String(), nil
}

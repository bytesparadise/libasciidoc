package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLink(ctx *context, l *types.InlineLink) (string, error) {
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
			text = htmlEscaper.Replace(t)
		case []interface{}:
			var err error
			if text, err = r.renderInlineElements(ctx, t); err != nil {
				return "", errors.Wrap(err, "unable to render link")
			}
		}
		class = roles // can be empty (and it's fine)
	} else {
		text = htmlEscaper.Replace(l.Location.ToDisplayString())
		if l.Location != nil && l.Location.Scheme != "mailto:" {
			class = "bare"
		}
		if len(roles) > 0 {
			class = strings.Join([]string{class, roles}, " ") // support case where class == "" (for email addresses)
		}
	}
	target := l.Attributes.GetAsStringWithDefault(types.AttrInlineLinkTarget, "")
	noopener := target == "_blank" || l.Attributes.HasOption("noopener")
	return r.execute(r.link, struct {
		ID       string
		URL      string
		Text     string
		Class    string
		Target   string
		NoOpener bool
	}{
		ID:       id,
		URL:      htmlEscaper.Replace(l.Location.ToString()),
		Text:     text,
		Class:    class,
		Target:   target,
		NoOpener: noopener,
	})
}

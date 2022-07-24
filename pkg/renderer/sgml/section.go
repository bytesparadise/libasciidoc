package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderSection(ctx *context, s *types.Section) (string, error) {
	// log.Debugf("rendering section level %d", s.Level)
	title, err := r.renderSectionTitle(ctx, s)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section title")
	}

	content, err := r.renderElements(ctx, s.Elements)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section content")
	}
	roles, err := r.renderElementRoles(ctx, s.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render section roles")
	}
	return r.execute(r.sectionContent, struct {
		Context  *context
		Header   string
		Content  string
		Elements []interface{}
		ID       string
		Roles    string
		Level    int
	}{
		Context:  ctx,
		Header:   title,
		Level:    s.Level,
		Elements: s.Elements,
		ID:       r.renderElementID(s.Attributes),
		Roles:    roles,
		Content:  string(content),
	})
}

func (r *sgmlRenderer) renderSectionTitle(ctx *context, s *types.Section) (string, error) {
	renderedContent, err := r.renderInlineElements(ctx, s.Title)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section title")
	}
	renderedContentStr := strings.TrimSpace(renderedContent)
	var number string
	if ctx.sectionNumbering != nil {
		id := s.GetID()
		log.Debugf("number for section '%s': '%s'", id, number)
		number = ctx.sectionNumbering[id]
	}
	return r.execute(r.sectionTitle, struct {
		Level        int
		LevelPlusOne int
		ID           string
		Number       string
		Content      string
	}{
		Level:        s.Level,
		LevelPlusOne: s.Level + 1, // Level 1 is <h2>.
		ID:           r.renderElementID(s.Attributes),
		Number:       number,
		Content:      renderedContentStr,
	})
}

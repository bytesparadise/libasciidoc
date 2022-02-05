package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderSection(ctx *renderer.Context, s *types.Section) (string, error) {
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

	result := &strings.Builder{}
	err = r.sectionContent.Execute(result, struct {
		Context  *renderer.Context
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
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section")
	}
	// log.Debugf("rendered section: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderSectionTitle(ctx *renderer.Context, s *types.Section) (string, error) {
	result := &strings.Builder{}
	renderedContent, err := r.renderInlineElements(ctx, s.Title)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering sectionTitle content")
	}
	renderedContentStr := strings.TrimSpace(renderedContent)
	var number string
	if ctx.SectionNumbering != nil {
		id, err := s.GetID()
		if err != nil {
			return "", errors.Wrap(err, "error while rendering sectionTitle content")
		}
		log.Debugf("number for section '%s': '%s'", id, number)
		number = ctx.SectionNumbering[id]
	}
	err = r.sectionTitle.Execute(result, struct {
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
	if err != nil {
		return "", errors.Wrapf(err, "error while rendering sectionTitle")
	}
	// log.Debugf("rendered sectionTitle: %s", result.Bytes())
	return string(result.String()), nil
}

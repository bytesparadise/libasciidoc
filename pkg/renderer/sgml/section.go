package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderPreamble(ctx *renderer.Context, p types.Preamble) (string, error) {
	log.Debugf("rendering preamble...")
	result := &strings.Builder{}
	// the <div id="preamble"> wrapper is only necessary
	// if the document has a section 0

	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "error rendering preamble elements")
	}
	err = r.preamble.Execute(result, struct {
		Context *renderer.Context
		Wrapper bool
		Content string
	}{
		Context: ctx,
		Wrapper: ctx.HasHeader,
		Content: string(content),
	})
	if err != nil {
		return "", errors.Wrap(err, "error while rendering preamble")
	}
	// log.Debugf("rendered preamble: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderSection(ctx *renderer.Context, s types.Section) (string, error) {
	log.Debugf("rendering section level %d", s.Level)
	title, err := r.renderSectionTitle(ctx, s)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section title")
	}

	content, err := r.renderElements(ctx, s.Elements)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section content")
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
		Roles:    r.renderElementRoles(s.Attributes),
		Content:  string(content),
	})
	if err != nil {
		return "", errors.Wrap(err, "error while rendering section")
	}
	// log.Debugf("rendered section: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderSectionTitle(ctx *renderer.Context, s types.Section) (string, error) {
	result := &strings.Builder{}
	renderedContent, err := r.renderInlineElements(ctx, s.Title)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering sectionTitle content")
	}
	renderedContentStr := strings.TrimSpace(renderedContent)
	err = r.sectionHeader.Execute(result, struct {
		Level        int
		LevelPlusOne int
		ID           string
		Roles        string
		Content      string
	}{
		Level:        s.Level,
		LevelPlusOne: s.Level + 1, // Level 1 is <h2>.
		ID:           r.renderElementID(s.Attributes),
		Roles:        r.renderElementRoles(s.Attributes),
		Content:      renderedContentStr,
	})
	if err != nil {
		return "", errors.Wrapf(err, "error while rendering sectionTitle")
	}
	// log.Debugf("rendered sectionTitle: %s", result.Bytes())
	return string(result.String()), nil
}

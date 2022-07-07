package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderFootnoteReference(note *types.FootnoteReference) (string, error) {
	result := &strings.Builder{}
	if note.ID != types.InvalidFootnoteReference && !note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		tmpl, err := r.footnote()
		if err != nil {
			return "", errors.Wrap(err, "unable to load footnote template")
		}
		if err := tmpl.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		}); err != nil {
			return "", errors.Wrap(err, "unable to render footnote")
		}
	} else if note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		tmpl, err := r.footnoteRef()
		if err != nil {
			return "", errors.Wrap(err, "unable to load footnote template")
		}
		if err := tmpl.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		}); err != nil {
			return "", errors.Wrap(err, "unable to render footnote")
		}
	} else {
		// invalid footnote
		tmpl, err := r.invalidFootnote()
		if err != nil {
			return "", errors.Wrap(err, "unable to load missing footnote template")
		}
		if err := tmpl.Execute(result, struct {
			Ref string
		}{
			Ref: note.Ref,
		}); err != nil {
			return "", errors.Wrap(err, "unable to render missing footnote")
		}
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderFootnotes(ctx *context, notes []*types.Footnote) (string, error) {
	// skip if there's no foot note in the doc
	if len(notes) == 0 {
		return "", nil
	}
	content := &strings.Builder{}
	for _, note := range notes {
		renderedNote, err := r.renderFootnoteElement(ctx, note)
		if err != nil {
			return "", errors.Wrap(err, "failed to render footnote element")
		}
		content.WriteString(renderedNote)
	}
	return r.execute(r.footnotes, struct {
		Context   *context
		Content   string
		Footnotes []*types.Footnote
	}{
		Context:   ctx,
		Content:   content.String(),
		Footnotes: notes,
	})
}

func (r *sgmlRenderer) renderFootnoteElement(ctx *context, note *types.Footnote) (string, error) {
	content, err := r.renderInlineElements(ctx, note.Elements)
	if err != nil {
		return "", errors.Wrapf(err, "unable to render foot note content")
	}
	content = strings.TrimSpace(content)
	// Note: Asciidoctor will render the footnote content on a single line
	content = strings.ReplaceAll(content, "\n", " ")
	return r.execute(r.footnoteElement, struct {
		Context *context
		ID      int
		Ref     string
		Content string
	}{
		Context: ctx,
		ID:      note.ID,
		Ref:     note.Ref,
		Content: string(content),
	})
}

package sgml

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderFootnote(ctx *renderer.Context, elements []interface{}) (string, error) {
	result, err := r.renderInlineElements(ctx, elements)
	if err != nil {
		return "", errors.Wrapf(err, "unable to render foot note content")
	}
	return strings.TrimSpace(string(result)), nil
}

func (r *sgmlRenderer) renderFootnoteReference(note types.FootnoteReference) (string, error) {
	result := &strings.Builder{}
	if note.ID != types.InvalidFootnoteReference && !note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		err := r.footnote.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		})
		if err != nil {
			return "", errors.Wrap(err, "unable to render footnote")
		}
	} else if note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		err := r.footnoteRef.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		})
		if err != nil {
			return "", errors.Wrap(err, "unable to render footnote")
		}
	} else {
		// invalid footnote
		err := r.invalidFootnote.Execute(result, struct {
			Ref string
		}{
			Ref: note.Ref,
		})
		if err != nil {
			return "", errors.Wrap(err, "unable to render missing footnote")
		}
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderFootnoteReferencePlainText(note types.FootnoteReference) (string, error) {
	result := &strings.Builder{}
	if note.ID != types.InvalidFootnoteReference {
		// valid case for a footnote with content, with our without an explicit reference
		err := r.footnoteRefPlain.Execute(result, struct {
			ID    int
			Class string
		}{
			ID:    note.ID,
			Class: "footnote",
		})
		if err != nil {
			return "", errors.Wrap(err, "unable to render footnote")
		}
	} else {
		return "", fmt.Errorf("unable to render missing footnote")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderFootnotes(ctx *renderer.Context, notes []types.Footnote) (string, error) {
	// skip if there's no foot note in the doc
	if len(notes) == 0 {
		return "", nil
	}
	result := &strings.Builder{}
	err := r.footnotes.Execute(result,
		ContextualPipeline{
			Context: ctx,
			Data: struct {
				Footnotes []types.Footnote
			}{
				Footnotes: notes,
			},
		})
	if err != nil {
		return "", errors.Wrap(err, "failed to render footnotes")
	}
	return result.String(), nil
}

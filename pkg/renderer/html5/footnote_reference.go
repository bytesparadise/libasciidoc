package html5

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var footnoteTmpl texttemplate.Template
var footnoterefTmpl texttemplate.Template
var footnoterefPlainTextTmpl texttemplate.Template
var invalidFootnoteTmpl texttemplate.Template
var footnotesTmpl texttemplate.Template

// initializes the templates
func init() {
	footnoteTmpl = newTextTemplate("footnote", `<sup class="footnote"{{ if .Ref }} id="_footnote_{{ .Ref }}"{{ end }}>[<a id="_footnoteref_{{ .ID }}" class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`,
		texttemplate.FuncMap{
			"renderIndex": renderFootnoteIndex,
		})
	footnoterefTmpl = newTextTemplate("footnote ref", `<sup class="footnoteref">[<a class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`,
		texttemplate.FuncMap{
			"renderIndex": renderFootnoteIndex,
		})
	footnoterefPlainTextTmpl = newTextTemplate("footnote ref", `<sup class="{{ .Class }}">[{{ .ID }}]</sup>`,
		texttemplate.FuncMap{
			"renderIndex": renderFootnoteIndex,
		})

	invalidFootnoteTmpl = newTextTemplate("invalid footnote", `<sup class="footnoteref red" title="Unresolved footnote reference.">[{{ .Ref }}]</sup>`)
	footnotesTmpl = newTextTemplate("footnotes", `
<div id="footnotes">
<hr>{{ $ctx := .Context }}{{ with .Data }}{{ $footnotes := .Footnotes }}{{ range $index, $footnote := $footnotes }}
<div class="footnote" id="_footnotedef_{{ renderIndex $index }}">
<a href="#_footnoteref_{{ renderIndex $index }}">{{ renderIndex $index }}</a>. {{ renderFootnoteContent $ctx $footnote.Elements }}
</div>{{ end }}{{ end }}
</div>`,
		texttemplate.FuncMap{
			"renderFootnoteContent": func(ctx renderer.Context, elements []interface{}) (string, error) {
				result, err := renderInlineElements(ctx, elements)
				if err != nil {
					return "", errors.Wrapf(err, "unable to render foot note content")
				}
				return strings.TrimSpace(string(result)), nil
			},
			"renderIndex": renderFootnoteIndex,
		})
}

func renderFootnoteIndex(idx int) string {
	return strconv.Itoa(idx + 1)
}

func renderFootnoteReference(ctx renderer.Context, note types.FootnoteReference) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if note.ID != types.InvalidFootnoteReference && !note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		err := footnoteTmpl.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render footnote")
		}
	} else if note.Duplicate {
		// valid case for a footnote with content, with our without an explicit reference
		err := footnoterefTmpl.Execute(result, struct {
			ID  int
			Ref string
		}{
			ID:  note.ID,
			Ref: note.Ref,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render footnote")
		}
	} else {
		// invalid footnote
		err := invalidFootnoteTmpl.Execute(result, struct {
			Ref string
		}{
			Ref: note.Ref,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render missing footnote")
		}
	}
	return result.Bytes(), nil
}

func renderFootnoteReferencePlainText(ctx renderer.Context, note types.FootnoteReference) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if note.ID != types.InvalidFootnoteReference {
		// valid case for a footnte with content, with our without an explicit reference
		err := footnoterefPlainTextTmpl.Execute(result, struct {
			ID    int
			Class string
		}{
			ID:    note.ID,
			Class: "footnote",
		})
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render footnote")
		}
	} else {
		return nil, fmt.Errorf("unable to render missing footnote")
	}
	return result.Bytes(), nil
}

func renderFootnotes(ctx renderer.Context, notes []types.Footnote) ([]byte, error) {
	// skip if there's no foot note in the doc
	if len(notes) == 0 {
		return []byte{}, nil
	}
	result := bytes.NewBuffer(nil)
	err := footnotesTmpl.Execute(result,
		ContextualPipeline{
			Context: ctx,
			Data: struct {
				Footnotes []types.Footnote
			}{
				Footnotes: notes,
			},
		})
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to render footnotes")
	}
	return result.Bytes(), nil
}

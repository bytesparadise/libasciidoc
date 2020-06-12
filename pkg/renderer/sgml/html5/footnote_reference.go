package html5

const (
	footnoteTmpl         = `<sup class="footnote"{{ if .Ref }} id="_footnote_{{ .Ref }}"{{ end }}>[<a id="_footnoteref_{{ .ID }}" class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`
	footnoteRefTmpl      = `<sup class="footnoteref">[<a class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`
	footnoteRefPlainTmpl = `<sup class="{{ .Class }}">[{{ .ID }}]</sup>`
	invalidFootnoteTmpl  = `<sup class="footnoteref red" title="Unresolved footnote reference.">[{{ .Ref }}]</sup>`
	footnotesTmpl        = `
<div id="footnotes">
<hr>{{ $ctx := .Context }}{{ with .Data }}{{ $footnotes := .Footnotes }}{{ range $index, $footnote := $footnotes }}
<div class="footnote" id="_footnotedef_{{ $footnote.ID }}">
<a href="#_footnoteref_{{ $footnote.ID }}">{{ $footnote.ID }}</a>. {{ renderFootnote $ctx $footnote.Elements }}
</div>{{ end }}{{ end }}
</div>`
)

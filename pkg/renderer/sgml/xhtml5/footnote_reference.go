package xhtml5

const (
	footnotesTmpl = `
<div id="footnotes">
<hr/>{{ $ctx := .Context }}{{ with .Data }}{{ $footnotes := .Footnotes }}{{ range $index, $footnote := $footnotes }}
<div class="footnote" id="_footnotedef_{{ $footnote.ID }}">
<a href="#_footnoteref_{{ $footnote.ID }}">{{ $footnote.ID }}</a>. {{ renderFootnote $ctx $footnote.Elements }}
</div>{{ end }}{{ end }}
</div>`
)

package html5

const (
	footnoteTmpl        = `<sup class="footnote"{{ if .Ref }} id="_footnote_{{ .Ref }}"{{ end }}>[<a id="_footnoteref_{{ .ID }}" class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`
	footnoteRefTmpl     = `<sup class="footnoteref">[<a class="footnote" href="#_footnotedef_{{ .ID }}" title="View footnote.">{{ .ID }}</a>]</sup>`
	invalidFootnoteTmpl = `<sup class="footnoteref red" title="Unresolved footnote reference.">[{{ .Ref }}]</sup>`
	footnotesTmpl       = "<div id=\"footnotes\">\n<hr>\n{{ .Content }}</div>\n"

	// arguably this should instead be an ordered list.
	footnoteElementTmpl = "<div class=\"footnote\" id=\"_footnotedef_{{ .ID }}\">\n" +
		"<a href=\"#_footnoteref_{{ .ID }}\">{{ .ID }}</a>. {{ .Content }}\n" +
		"</div>\n"
)

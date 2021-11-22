package xhtml5

const (
	tableTmpl = "<table {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"tableblock" +
		" frame-{{ .Frame }} grid-{{ .Grid }}" +
		"{{ if .Stripes }} stripes-{{ .Stripes }}{{ end }}" +
		"{{ if .Fit }} {{ .Fit }}{{ end }}" +
		"{{ if .Float }} {{ .Float }}{{ end }}" +
		"{{ if .Roles }} {{ .Roles }}{{ end }}\"" +
		"{{ if .Width }} style=\"width: {{ .Width }}%;\"{{ end }}" +
		">\n" +
		"{{ if .Title }}<caption class=\"title\">{{ .Caption }}{{ .Title }}</caption>\n{{ end }}" +
		"{{ if .Body }}" +
		"<colgroup>\n" +
		"{{ range $i, $w := .Columns }}<col" +
		"{{ if $w.Width }} style=\"width: {{ $w.Width }}%;\"{{ end }}" +
		"/>\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ end }}" +
		"</table>\n"
)

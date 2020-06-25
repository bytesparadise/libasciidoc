package xhtml5

const (
	tableTmpl = "<table class=\"tableblock frame-all grid-all stretch{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<caption class=\"title\">Table {{ .TableNumber }}. {{ .Title }}</caption>\n{{ end }}" +
		"{{ if .Body }}" +
		"<colgroup>\n" +
		"{{ range $i, $w := .CellWidths }}<col style=\"width: {{ $w }}%;\"/>\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ end }}" +
		"</table>"
)

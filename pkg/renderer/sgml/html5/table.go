package html5

const (
	// TODO: These class settings need to be overridable via attributes
	tableTmpl = "<table class=\"tableblock frame-all grid-all stretch{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<caption class=\"title\">Table {{ .TableNumber }}. {{ .Title }}</caption>\n{{ end }}" +
		"{{ if .Body }}" +
		"<colgroup>\n" +
		"{{ range $i, $w := .CellWidths }}<col style=\"width: {{ $w }}%;\">\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ end }}" +
		"</table>"

	tableBodyTmpl = "{{ if .Content }}<tbody>\n{{ .Content }}</tbody>\n{{ end }}"

	tableHeaderTmpl = "{{ if .Content }}<thead>\n<tr>\n{{ .Content }}</tr>\n</thead>\n{{ end }}"

	tableRowTmpl = "<tr>\n{{ .Content }}</tr>\n"

	// TODO: review these alignment choices ... should they be overrideable?

	tableHeaderCellTmpl = "<th class=\"tableblock halign-left valign-top\">{{ .Content }}</th>\n"

	tableCellTmpl = "<td class=\"tableblock halign-left valign-top\"><p class=\"tableblock\">{{ .Content }}</p></td>\n"
)

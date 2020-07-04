package html5

const (
	tableTmpl = "<table class=\"tableblock" +
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
		"{{ range $i, $w := .CellWidths }}<col style=\"width: {{ $w }}%;\">\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ end }}" +
		"</table>"

	tableBodyTmpl = "{{ if .Content }}<tbody>\n{{ .Content }}</tbody>\n{{ end }}"

	tableHeaderTmpl = "{{ if .Content }}<thead>\n<tr>\n{{ .Content }}</tr>\n</thead>\n{{ end }}"

	tableRowTmpl = "<tr>\n{{ .Content }}</tr>\n"

	tableCaptionTmpl = "Table {{ .TableNumber }}. "

	// TODO: cell styling via attributes

	tableHeaderCellTmpl = "<th class=\"tableblock halign-left valign-top\">{{ .Content }}</th>\n"

	tableCellTmpl = "<td class=\"tableblock halign-left valign-top\"><p class=\"tableblock\">{{ .Content }}</p></td>\n"
)

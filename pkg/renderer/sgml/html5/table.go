package html5

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
		">\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ .Footer }}" +
		"{{ end }}" +
		"</table>\n"

	tableBodyTmpl = "{{ if .Content }}<tbody>\n{{ .Content }}</tbody>\n{{ end }}"

	tableHeaderTmpl = "{{ if .Content }}<thead>\n<tr>\n{{ .Content }}</tr>\n</thead>\n{{ end }}"

	tableHeaderCellTmpl = "<th class=\"tableblock {{ halign .HAlign }} {{ valign .VAlign }}\">{{ .Content }}</th>\n"

	tableFooterTmpl = "{{ if .Content }}<tfoot>\n<tr>\n{{ .Content }}</tr>\n</tfoot>\n{{ end }}"

	tableFooterCellTmpl = "<td class=\"tableblock {{ halign .HAlign }} {{ valign .VAlign }}\"><p class=\"tableblock\">{{ .Content }}</p></td>\n"

	tableRowTmpl = "<tr>\n{{ .Content }}</tr>\n"

	tableCellTmpl = "<td class=\"tableblock {{ halign .HAlign }} {{ valign .VAlign }}\"><p class=\"tableblock\">{{ .Content }}</p></td>\n"
)

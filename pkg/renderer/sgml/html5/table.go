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
		"{{ range $i, $w := .Columns }}<col" +
		"{{ if $w.Width }} style=\"width: {{ $w.Width }}%;\"{{ end }}" +
		">\n{{ end}}" +
		"</colgroup>\n" +
		"{{ .Header }}" +
		"{{ .Body }}" +
		"{{ end }}" +
		"</table>\n"

	tableBodyTmpl = "{{ if .Content }}<tbody>\n{{ .Content }}</tbody>\n{{ end }}"

	tableHeaderTmpl = "{{ if .Content }}<thead>\n<tr>\n{{ .Content }}</tr>\n</thead>\n{{ end }}"

	tableRowTmpl = "<tr>\n{{ .Content }}</tr>\n"

	tableHeaderCellTmpl = "<th class=\"tableblock {{ halign .HAlign }} {{ valign .VAlign }}\">{{ .Content }}</th>\n"

	tableCellTmpl = "<td class=\"tableblock {{ halign .HAlign }} {{ valign .VAlign }}\"><p class=\"tableblock\">{{ .Content }}</p></td>\n"
)

package html5

const (
	labeledListTmpl = `<div` +
		`{{ if .ID }} id="{{ .ID }}"{{ end }}` +
		" class=\"dlist{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<dl>\n{{ .Content }}</dl>\n</div>"

	labeledListItemTmpl = "<dt class=\"hdlist1\">{{ .Term }}</dt>\n" +
		"{{ if .Content }}<dd>\n{{ .Content }}\n</dd>\n{{ end }}"

	labeledListHorizontalTmpl = `<div` +
		`{{ if .ID }} id="{{ .ID }}"{{ end }} ` +
		"class=\"hdlist{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<table>\n{{ .Content }}</table>\n</div>"

	// Continuation items (multiple terms sharing a single definition) make this a bit more complex.
	labeledListHorizontalItemTmpl = "{{ if not .Continuation }}<tr>\n" +
		"<td class=\"hdlist1\">\n{{ else }}<br>\n{{ end }}" +
		"{{ .Term }}\n" +
		"{{ if .Content }}</td>\n<td class=\"hdlist2\">\n{{ .Content }}\n</td>\n</tr>\n{{ end }}"

	qAndAListTmpl = "<div{{ if .ID }} id=\"{{ .ID }}\"{{ end }} " +
		"class=\"qlist qanda{{ if .Roles }} {{ .Roles }}{{ end }}\"" +
		">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<ol>\n{{ .Content }}</ol>\n</div>"

	qAndAListItemTmpl = "<li>\n<p><em>{{ .Term }}</em></p>\n{{ .Content }}\n</li>\n"
)

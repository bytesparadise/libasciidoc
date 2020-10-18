package html5

const (
	admonitionBlockTmpl = `<div {{ if .ID }}id="{{ .ID}}" {{ end }}` +
		"class=\"admonitionblock {{ .Kind }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<table>\n" +
		"<tr>\n" +
		"<td class=\"icon\">\n{{ .Icon }}\n</td>\n" +
		"<td class=\"content\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"{{ .Content }}" +
		"</td>\n</tr>\n</table>\n</div>\n"
)

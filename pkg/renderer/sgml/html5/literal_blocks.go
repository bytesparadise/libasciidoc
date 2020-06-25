package html5

const (
	literalBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"literalblock\"{{ if .Roles }} {{ .Roles }}{{ end }}>\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre>{{ .Content }}</pre>\n" +
		"</div>\n" +
		"</div>"
)

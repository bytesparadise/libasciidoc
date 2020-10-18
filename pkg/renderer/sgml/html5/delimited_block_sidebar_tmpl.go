package html5

const (
	sidebarBlockTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}" +
		"class=\"sidebarblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<div class=\"content\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"{{ .Content }}" +
		"</div>\n" +
		"</div>\n"
)

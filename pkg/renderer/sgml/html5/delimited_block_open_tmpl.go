package html5

const (
	openBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"openblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"</div>\n"
)

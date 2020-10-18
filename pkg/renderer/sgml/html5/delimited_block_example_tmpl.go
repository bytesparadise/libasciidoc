package html5

const (
	exampleBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"exampleblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Caption }}{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"</div>\n"
)

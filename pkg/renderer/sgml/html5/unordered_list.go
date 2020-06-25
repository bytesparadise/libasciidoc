package html5

const (
	unorderedListTmpl = `<div{{ if .ID }} id="{{ .ID }}"{{ end }}` +
		` class="ulist{{ if .Checklist }} checklist{{ end }}` +
		`{{ if .Roles }} {{ .Roles }}{{ end }}"` +
		">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<ul{{ if .Checklist }} class=\"checklist\"{{ end }}>\n" +
		"{{ .Content }}</ul>\n</div>"

	unorderedListItemTmpl = "<li>\n{{ .Content }}\n</li>\n"
)

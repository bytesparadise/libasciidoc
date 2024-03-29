package html5

const (
	unorderedListTmpl = `<div{{ if .ID }} id="{{ .ID }}"{{ end }}` +
		` class="ulist{{ if .Checklist }} checklist{{ end }}` +
		`{{ if .Style }} {{ .Style }}{{ end }}` +
		`{{ if .Roles }} {{ .Roles }}{{ end }}"` +
		">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<ul{{ if .Checklist }} class=\"checklist\"{{ else if .Style }} class=\"{{ .Style }}\"{{ end }}>\n" +
		"{{ .Content }}</ul>\n</div>\n"

	unorderedListElementTmpl = "<li>\n{{ .Content }}</li>\n"
)

package html5

const (
	orderedListTmpl = `<div{{ if .ID }} id="{{ .ID }}"{{ end }}` +
		` class="olist {{ .NumberingStyle }}` +
		`{{ if .Roles }} {{ .Roles }}{{ end }}"` +
		">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		`<ol` +
		` class="{{ .NumberingStyle }}"` +
		`{{ if .ListStyle }} type="{{ .ListStyle }}"{{ end }}` +
		`{{ if .Start }} start="{{ .Start }}"{{ end }}` +
		">\n{{ .Content }}</ol>\n</div>"

	orderedListItemTmpl = "<li>\n{{ .Content }}\n</li>\n"
)

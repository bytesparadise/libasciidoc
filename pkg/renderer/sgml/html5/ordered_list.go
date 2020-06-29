package html5

const (
	orderedListTmpl = `<div{{ if .ID }} id="{{ .ID }}"{{ end }}` +
		` class="olist {{ .Style }}` +
		`{{ if .Roles }} {{ .Roles }}{{ end }}"` +
		">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		`<ol` +
		` class="{{ .Style }}"` +
		`{{ if .ListStyle }} type="{{ .ListStyle }}"{{ end }}` +
		`{{ if .Start }} start="{{ .Start }}"{{ end }}` +
		`{{ if .Reversed }} reversed{{ end }}` +
		">\n{{ .Content }}</ol>\n</div>"

	orderedListItemTmpl = "<li>\n{{ .Content }}\n</li>\n"
)

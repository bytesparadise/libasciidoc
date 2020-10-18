package html5

const (
	quoteBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"quoteblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<blockquote>\n" +
		"{{ .Content }}" +
		"</blockquote>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n" +
		"&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br>\n<cite>{{ .Attribution.Second }}</cite>{{ end }}\n" +
		"</div>\n{{ end }}" +
		"</div>\n"
)

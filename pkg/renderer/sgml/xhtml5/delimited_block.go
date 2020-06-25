package xhtml5

const (
	quoteBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"quoteblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<blockquote>\n" +
		"{{ .Content }}\n" +
		"</blockquote>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n" +
		"&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br/>\n<cite>{{ .Attribution.Second }}</cite>{{ end }}\n" +
		"</div>\n{{ end }}" +
		"</div>"

	verseBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"verseblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<pre class=\"content\">{{ .Content }}</pre>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br/>\n<cite>{{ .Attribution.Second }}</cite>{{ end }}\n" +
		"</div>\n{{ end }}" +
		"</div>"
)

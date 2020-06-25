package xhtml5

const (
	verseParagraphTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"verseblock\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<pre class=\"content\">{{ .Content }}</pre>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br/>\n<cite>{{ .Attribution.Second }}</cite>\n{{ else }}\n{{ end }}" +
		"</div>\n{{ end }}</div>"

	quoteParagraphTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"quoteblock\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<blockquote>\n" +
		"{{ .Content }}\n" +
		"</blockquote>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br/>\n<cite>{{ .Attribution.Second }}</cite>\n{{ else }}\n{{ end }}" +
		"</div>\n{{ end }}</div>"

	thematicBreakTmpl = "<hr/>"
)

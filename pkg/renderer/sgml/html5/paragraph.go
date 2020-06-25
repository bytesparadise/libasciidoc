package html5

const (
	paragraphTmpl = "<div{{ if .ID }} id=\"{{ .ID }}\"{{ end }}" +
		" class=\"paragraph{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title  }}<div class=\"doctitle\">{{ .Title }}</div>\n{{ end }}" +
		"<p>{{ .Content }}</p>\n" +
		"</div>"

	admonitionParagraphTmpl = `{{ if .Content }}` +
		"<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"admonitionblock {{ .Kind }}" +
		"{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<table>\n<tr>\n<td class=\"icon\">\n" +
		"{{ if .Icon }}{{ .Icon }}{{ end }}\n" +
		"</td>\n" +
		"<td class=\"content\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"{{ .Content }}\n" +
		"</td>\n</tr>\n</table>\n</div>{{ end }}"

	delimitedBlockParagraphTmpl = `<p>{{ .CheckStyle }}{{ .Content }}</p>`

	sourceParagraphTmpl = "<div class=\"listingblock\">\n" +
		"<div class=\"content\">\n" +
		"<pre class=\"highlight\">" +
		`<code{{ if .Language }} class="language-{{ .Language }}" data-lang="{{ .Language }}"{{ end }}>` +
		"{{ .Content }}" +
		"</code></pre>\n" +
		"</div>\n" +
		"</div>"

	verseParagraphTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"verseblock\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<pre class=\"content\">{{ .Content }}</pre>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br>\n<cite>{{ .Attribution.Second }}</cite>\n{{ else }}\n{{ end }}" +
		"</div>\n{{ end }}</div>"

	quoteParagraphTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}class=\"quoteblock\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<blockquote>\n" +
		"{{ .Content }}\n" +
		"</blockquote>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br>\n<cite>{{ .Attribution.Second }}</cite>\n{{ else }}\n{{ end }}" +
		"</div>\n{{ end }}</div>"

	manpageNameParagraphTmpl = `<p>{{ .Content }}</p>`

	thematicBreakTmpl = "<hr>"
)

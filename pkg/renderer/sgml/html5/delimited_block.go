package html5

const (
	fencedBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"listingblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre class=\"highlight\"><code>{{ .Content }}</code></pre>\n" +
		"</div>\n" +
		"</div>\n"

	listingBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"listingblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre>{{ .Content }}</pre>\n" +
		"</div>\n" +
		"</div>\n"

	sourceBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"listingblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre class=\"" +
		`{{ if .SyntaxHighlighter }}{{ .SyntaxHighlighter }} {{ end }}` +
		`highlight">` +
		`<code{{ if .Language }}{{ if not .SyntaxHighlighter }} class="language-{{ .Language}}"{{ end }} ` +
		`data-lang="{{ .Language}}"{{ end }}>` +
		"{{ .Content }}</code></pre>\n" +
		"</div>\n" +
		"</div>\n"

	exampleBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"exampleblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Caption }}{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"</div>\n"

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

	verseBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"verseblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<pre class=\"content\">{{ .Content }}</pre>\n" +
		"{{ if .Attribution.First }}<div class=\"attribution\">\n&#8212; {{ .Attribution.First }}" +
		"{{ if .Attribution.Second }}<br>\n<cite>{{ .Attribution.Second }}</cite>{{ end }}\n" +
		"</div>\n{{ end }}" +
		"</div>\n"

	admonitionBlockTmpl = `<div {{ if .ID }}id="{{ .ID}}" {{ end }}` +
		"class=\"admonitionblock {{ .Kind }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<table>\n" +
		"<tr>\n" +
		"<td class=\"icon\">\n{{ .Icon }}\n</td>\n" +
		"<td class=\"content\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"{{ .Content }}" +
		"</td>\n</tr>\n</table>\n</div>\n"

	sidebarBlockTmpl = "<div {{ if .ID }}id=\"{{ .ID }}\" {{ end }}" +
		"class=\"sidebarblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<div class=\"content\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"{{ .Content }}" +
		"</div>\n" +
		"</div>\n"

	// the name here is weird because "pass" as a prefix triggers a false security warning
	pssThroughBlock = "{{ .Content }}\n"
)

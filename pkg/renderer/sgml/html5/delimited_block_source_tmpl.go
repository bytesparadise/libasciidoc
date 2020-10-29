package html5

const (
	sourceBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"listingblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre class=\"" +
		`{{ if .SyntaxHighlighter }}{{ .SyntaxHighlighter }} {{ end }}` +
		`highlight` +
		`{{ if .Option }} {{ .Option }}{{ end }}` + // space before the option as it's the last value in the 'class' attribute
		`">` +
		`<code{{ if .Language }}{{ if not .SyntaxHighlighter }} class="language-{{ .Language}}"{{ end }} ` +
		`data-lang="{{ .Language}}"{{ end }}>` +
		"{{ .Content }}</code></pre>\n" +
		"</div>\n" +
		"</div>\n"
)

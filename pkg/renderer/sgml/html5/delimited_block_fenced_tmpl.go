package html5

const (
	fencedBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}` +
		"class=\"listingblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<div class=\"content\">\n" +
		"<pre class=\"highlight\"><code>{{ .Content }}</code></pre>\n" +
		"</div>\n" +
		"</div>\n"
)

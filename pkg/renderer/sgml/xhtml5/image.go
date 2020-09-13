package xhtml5

const (
	blockImageTmpl = "<div" +
		"{{ if .ID }} id=\"{{ .ID }}\"{{ end }}" +
		" class=\"imageblock{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"<div class=\"content\">\n" +
		`{{ if .Href }}<a class="image" href="{{ .Href }}">{{ end }}` +
		`<img src="{{ .Path }}" alt="{{ .Alt }}"` +
		`{{ if .Width }} width="{{ .Width }}"{{ end }}` +
		`{{ if .Height }} height="{{ .Height }}"{{ end }}` +
		"/>{{ if .Href }}</a>{{ end }}\n" +
		"</div>\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Caption }}{{ .Title }}</div>\n{{ end }}" +
		"</div>\n"

	inlineImageTmpl = `<span class="image{{ if .Roles }} {{ .Roles }}{{ end }}">` +
		`{{ if .Href }}<a class="image" href="{{ .Href }}">{{ end }}` +
		`<img src="{{ .Path }}" alt="{{ .Alt }}"` +
		`{{ if .Width }} width="{{ .Width }}"{{ end }}` +
		`{{ if .Height }} height="{{ .Height }}"{{ end }}` +
		`{{ if .Title }} title="{{ .Title }}"{{ end }}` +
		`/>{{ if .Href }}</a>{{ end }}</span>`
)

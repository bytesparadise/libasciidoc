package html5

// initializes the sgml
const (
	preambleTmpl = "{{ if .Wrapper }}<div id=\"preamble\">\n" +
		"<div class=\"sectionbody\">\n{{ end }}" +
		`{{ .Content }}` +
		"{{ if .Wrapper }}</div>\n</div>\n{{ end }}"

	sectionContentTmpl = "<div class=\"sect{{ .Level }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ .Header }}" +
		"{{ if eq .Level 1 }}<div class=\"sectionbody\">\n{{ end }}" +
		"{{ .Content }}" +
		"{{ if eq .Level 1 }}</div>\n{{ end }}" +
		"</div>\n"

	sectionHeaderTmpl = "<h{{ .LevelPlusOne }} id=\"{{ .ID }}\">{{ .Content }}</h{{ .LevelPlusOne }}>\n"
)

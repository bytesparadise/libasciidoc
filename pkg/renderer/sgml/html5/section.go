package html5

// initializes the sgml
const (
	preambleTmpl = "{{ if .Wrapper }}<div id=\"preamble\">\n" +
		"<div class=\"sectionbody\">\n{{ end }}" +
		`{{ .Content }}` +
		"{{ if .Wrapper }}\n</div>\n</div>{{ end }}"

	sectionContentTmpl = "<div class=\"sect{{ .Level }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ .Header }}" +
		"{{ if eq .Level 1 }}\n<div class=\"sectionbody\">{{ end }}" +
		"{{ if .Content }}\n{{ .Content }}{{ end }}\n" +
		"{{ if eq .Level 1 }}</div>\n{{ end }}" +
		"</div>"

	sectionHeaderTmpl = `<h{{ .LevelPlusOne }} id="{{ .ID }}">{{ .Content }}</h{{ .LevelPlusOne }}>`
)

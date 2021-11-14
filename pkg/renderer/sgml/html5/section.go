package html5

// initializes the sgml
const (
	// sectionContentTmpl = "<div class=\"sect{{ .Level }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
	// 	"{{ .Header }}" +
	// 	"{{ if eq .Level 1 }}<div class=\"sectionbody\">\n{{ end }}" +
	// 	"{{ .Content }}" +
	// 	"{{ if eq .Level 1 }}</div>\n{{ end }}" +
	// 	"</div>\n"
	sectionContentTmpl = `<div class="sect{{ .Level }}{{ if .Roles }} {{ .Roles }}{{ end }}">
{{ .Header }}{{ if eq .Level 1 }}<div class="sectionbody">
{{ end }}{{ .Content }}{{ if eq .Level 1 }}</div>
{{ end }}</div>
`
	// sectionHeaderTmpl = "<h{{ .LevelPlusOne }} id=\"{{ .ID }}\">{{ .Content }}</h{{ .LevelPlusOne }}>\n"
	sectionHeaderTmpl = `<h{{ .LevelPlusOne }} id="{{ .ID }}">{{ .Content }}</h{{ .LevelPlusOne }}>
`
)

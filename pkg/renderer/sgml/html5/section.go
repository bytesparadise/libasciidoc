package html5

// initializes the sgml
const (
	sectionContentTmpl = `<div class="sect{{ .Level }}{{ if .Roles }} {{ .Roles }}{{ end }}">
{{ .Header }}{{ if eq .Level 1 }}<div class="sectionbody">
{{ end }}{{ .Content }}{{ if eq .Level 1 }}</div>
{{ end }}</div>
`
	sectionTitleTmpl = `<h{{ .LevelPlusOne }} id="{{ .ID }}">{{ if .Number }}{{ .Number }}. {{ end }}{{ .Content }}</h{{ .LevelPlusOne }}>
`
)

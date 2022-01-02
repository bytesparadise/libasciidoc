package html5

const (
	linkTmpl = `<a{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .URL }} href="{{ .URL }}"{{ end }}{{if .Class}} class="{{ .Class }}"{{ end }}{{if .Target}} target="{{ .Target }}"{{ end }}>{{ .Text }}</a>`
)

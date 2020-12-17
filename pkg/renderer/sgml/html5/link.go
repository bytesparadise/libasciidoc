package html5

const (
	linkTmpl = `<a href="{{ .URL }}"{{if .Class}} class="{{ .Class }}"{{ end }}{{if .Target}} target="{{ .Target }}"{{ end }}>{{ .Text }}</a>`
)

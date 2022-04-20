package html5

const (
	linkTmpl = `<a{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .URL }} href="{{ escape .URL }}"{{ end }}{{if .Class}} class="{{ .Class }}"{{ end }}{{if .Target}} target="{{ .Target }}"{{ end }}{{ if .NoOpener}} rel="noopener"{{ end }}>{{ .Text }}</a>`
)

package html5

const (
	linkTmpl = `<a href="{{ .URL }}"{{if .Class}} class="{{ .Class }}"{{ end }}>{{ .Text }}</a>`
)

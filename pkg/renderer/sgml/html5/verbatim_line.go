package html5

const (
	verbatimLineTmpl = `{{ if .Callouts}}{{ escape .Content }}{{ else }}{{ .Content | escape | trimRight }}{{ end }}{{ range $i, $c := .Callouts }}<b class="conum">({{ $c.Ref }})</b>{{ end }}`
)

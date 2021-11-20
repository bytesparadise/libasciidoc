package html5

const (
	preambleTmpl = `{{ if .Wrapper }}<div id="preamble">
<div class="sectionbody">
{{ end }}{{ .Content }}{{ if .Wrapper }}</div>
{{ if .ToC }}{{ .ToC }}{{ end }}</div>
{{ end }}`
)

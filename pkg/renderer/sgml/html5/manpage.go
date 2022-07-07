package html5

const (
	manpageHeaderTmpl = `{{ if.IncludeH1 }}<div id="header">
<h1>{{ .Header }} Manual Page</h1>
{{ end }}<h2 id="_name">{{ .Name }}</h2>
<div class="sectionbody">
{{ .Content }}</div>
{{ if .IncludeH1 }}</div>
{{ end }}`
)

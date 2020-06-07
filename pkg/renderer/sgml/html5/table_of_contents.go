package html5

const (
	tocRootTmpl = `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
{{ . }}
</div>`

	tocSectionTmpl = `{{ $ctx := .Context }}{{ with .Data }}<ul class="sectlevel{{ .Level }}">
{{ range .Sections }}<li><a href="#{{ .ID }}">{{ .Title }}</a>{{ if .Children }}
{{ renderToC $ctx .Children }}
</li>{{else}}</li>{{end}}
{{end}}{{end}}</ul>`
)

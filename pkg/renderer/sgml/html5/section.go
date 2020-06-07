package html5

// initializes the sgml
const (
	preambleTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ if .Wrapper }}<div id="preamble">
<div class="sectionbody">
{{ end }}{{ renderElements $ctx .Elements | printf "%s" }}{{ if .Wrapper }}
</div>
</div>{{ end }}{{ end }}`

	sectionOneTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div class="{{ .Class }}">
{{ .SectionTitle }}
<div class="sectionbody">{{ $elements := renderElements $ctx .Elements | printf "%s" }}{{ if $elements }}
{{ $elements }}{{ end }}
</div>
</div>{{ end }}`

	sectionContentTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div class="{{ .Class }}">
{{ .SectionTitle }}{{ $elements := renderElements $ctx .Elements | printf "%s" }}{{ if $elements }}
{{ $elements }}{{ end }}
</div>{{ end }}`

	sectionHeaderTmpl = `<h{{ .Level }} id="{{ .ID }}">{{ .Content }}</h{{ .Level }}>`
)

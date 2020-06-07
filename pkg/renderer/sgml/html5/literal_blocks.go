package html5

const (
	literalBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="literalblock">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<div class="content">
<pre>{{ $lines := .Lines }}{{ range $index, $line := $lines}}{{ $line }}{{ includeNewline $ctx $index $lines }}{{ end }}</pre>
</div>
</div>{{ end }}`
)

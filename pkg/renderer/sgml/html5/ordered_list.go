package html5

const (
	orderedListTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="olist {{ .NumberingStyle }}{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol class="{{ .NumberingStyle }}"{{ .ListStyle }}{{ if .Start }} start="{{ .Start }}"{{ end }}>
{{ range $itemIndex, $item := $items }}<li>
{{ renderList $ctx $item.Elements | printf "%s" }}
</li>
{{ end }}</ol>
</div>{{ end }}`
)

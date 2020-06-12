package html5

// initializes the sgml
const (
	calloutListTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="colist arabic{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol>
{{ range $itemIndex, $item := $items }}<li>
{{ renderList $ctx $item.Elements | printf "%s" }}
</li>
{{ end }}</ol>
</div>{{ end }}`
)

package html5

const (
	unorderedListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="ulist{{ if .Checklist }} checklist{{ end }}{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ul{{ if .Checklist }} class="checklist"{{ end }}>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ renderList $ctx $elements | printf "%s" }}
</li>
{{ end }}</ul>
</div>{{ end }}`
)

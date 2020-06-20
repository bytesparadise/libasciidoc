package xhtml5

const (
	labeledListHorizontalTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="hdlist{{ if .Role }} {{ .Role }}{{ end }}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<table>
<tr>
<td class="hdlist1">{{ $items := .Items }}{{ range $itemIndex, $item := $items }}
{{ renderInline $ctx $item.Term }}
{{ if $item.Elements }}</td>
<td class="hdlist2">
{{ renderList $ctx $item.Elements }}
{{ if includeNewline $ctx $itemIndex $items }}</td>
</tr>
<tr>
<td class="hdlist1">{{ else }}</td>{{ end }}{{ else }}<br/>{{ end }}{{ end }}
</tr>
</table>
</div>{{ end }}`
)

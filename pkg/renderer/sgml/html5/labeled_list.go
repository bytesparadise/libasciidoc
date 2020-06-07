package html5

const (
	labeledListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="dlist{{ if .Role }} {{ .Role }}{{ end }}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<dl>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<dt class="hdlist1">{{ renderInline $ctx $item.Term | printf "%s" }}</dt>{{ if $item.Elements }}
<dd>
{{ renderList $ctx $item.Elements | printf "%s" }}
</dd>{{ end }}
{{ end }}</dl>
</div>{{ end }}`

	labeledListHorizontalTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="hdlist{{ if .Role }} {{ .Role }}{{ end }}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<table>
<tr>
<td class="hdlist1">{{ $items := .Items }}{{ range $itemIndex, $item := $items }}
{{ renderInline $ctx $item.Term | printf "%s" }}
{{ if $item.Elements }}</td>
<td class="hdlist2">
{{ renderList $ctx $item.Elements | printf "%s" }}
{{ if includeNewline $ctx $itemIndex $items }}</td>
</tr>
<tr>
<td class="hdlist1">{{ else }}</td>{{ end }}{{ else }}<br>{{ end }}{{ end }}
</tr>
</table>
</div>{{ end }}`

	qAndAListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="qlist qanda">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
<p><em>{{ renderInline $ctx $item.Term | printf "%s" }}</em></p>
{{ if $item.Elements }}{{ renderList $ctx $item.Elements | printf "%s" }}{{ end }}
</li>
{{ end }}</ol>
</div>{{ end }}`
)

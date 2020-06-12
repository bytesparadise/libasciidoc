package html5

const (
	tableTmpl = `{{ $ctx := .Context }}{{ with .Data }}<table class="tableblock frame-all grid-all stretch">{{ if .Lines }}
{{ if .Title }}<caption class="title">{{ escape .Title }}</caption>
{{ end }}<colgroup>
{{ $cellWidths := .CellWidths }}{{ range $index, $width := $cellWidths }}<col style="width: {{ $width }}%;">{{ includeNewline $ctx $index $cellWidths }}{{ end }}
</colgroup>
{{ if .Header }}{{ if .Header.Cells }}<thead>
<tr>
{{ $headerCells := .Header.Cells }}{{ range $index, $cell := $headerCells }}<th class="tableblock halign-left valign-top">{{ renderInline $ctx $cell | printf "%s" }}</th>{{ includeNewline $ctx $index $headerCells }}{{ end }}
</tr>
</thead>
{{ end }}{{ end }}<tbody>
{{ range $indexLine, $line := .Lines }}<tr>
{{ range $indexCells, $cell := $line.Cells }}<td class="tableblock halign-left valign-top"><p class="tableblock">{{ renderInline $ctx $cell | printf "%s" }}</p></td>{{ includeNewline $ctx $indexCells $line.Cells }}{{ end }}
</tr>
{{ end }}</tbody>{{ end }}
</table>{{ end }}`
)

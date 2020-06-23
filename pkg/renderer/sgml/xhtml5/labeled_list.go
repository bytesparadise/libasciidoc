package xhtml5

const (
	labeledListHorizontalItemTmpl = "{{ if not .Continuation }}<tr>\n" +
		"<td class=\"hdlist1\">\n{{ else }}<br/>\n{{ end }}" +
		"{{ .Term }}\n" +
		"{{ if .Content }}</td>\n<td class=\"hdlist2\">\n{{ .Content }}\n</td>\n</tr>\n{{ end }}"
)

package html5

// initializes the sgml
const (
	calloutListTmpl = `<div` +
		`{{ if .ID }} id="{{ .ID }}"{{ end }} ` +
		"class=\"colist arabic{{ if .Roles }} {{ .Roles }}{{ end}}\">\n" +
		"{{ if .Title }}<div class=\"title\">{{ .Title }}</div>\n{{ end }}" +
		"<ol>\n" +
		"{{ .Content }}" +
		"</ol>\n</div>\n"

	// NB: The items are numbered sequentially.
	calloutListItemTmpl = "<li>\n{{ .Content }}</li>\n"
)

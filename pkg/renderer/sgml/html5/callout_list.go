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
	calloutListElementTmpl = "<li>\n{{ .Content }}</li>\n"

	// This should probably have been a <span>, but for compatibility we use <b>
	calloutRefTmpl = "<b class=\"conum\">({{ .Ref }})</b>"
)

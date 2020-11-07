package xhtml5

const (
	articleTmpl = "<!DOCTYPE html>\n" +
		"<html xmlns=\"http://www.w3.org/1999/xhtml\" lang=\"en\">\n" +
		"<head>\n" +
		"<meta charset=\"UTF-8\"/>\n" +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"/>\n" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>\n" +
		"{{ if .Generator }}<meta name=\"generator\" content=\"{{ .Generator }}\"/>\n{{ end }}" +
		"{{ if .Authors }}<meta name=\"author\" content=\"{{ .Authors }}\"/>\n{{ end }}" +
		"{{ if .CSS}}<link type=\"text/css\" rel=\"stylesheet\" href=\"{{ .CSS }}\"/>\n{{ end }}" +
		"<title>{{ .Title }}</title>\n" +
		"</head>\n" +
		"<body" +
		"{{ if .ID }} id=\"{{ .ID }}\"{{ end }}" +
		" class=\"{{ .Doctype }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .IncludeHTMLBodyHeader }}{{ .Header }}{{ end }}" +
		"<div id=\"content\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"{{ if .IncludeHTMLBodyFooter }}<div id=\"footer\">\n" +
		"<div id=\"footer-text\">\n" +
		"{{ if .RevNumber }}Version {{ .RevNumber }}<br/>\n{{ end }}" +
		"Last updated {{ .LastUpdated }}\n" +
		"</div>\n" +
		"</div>\n{{ end }}" +
		"</body>\n" +
		"</html>\n"
)

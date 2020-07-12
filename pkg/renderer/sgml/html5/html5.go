package html5

const (
	articleTmpl = "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n" +
		"<head>\n" +
		"<meta charset=\"UTF-8\">\n" +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n" +
		"{{ if .Generator }}<meta name=\"generator\" content=\"{{ .Generator }}\">\n{{ end }}" +
		"{{ if .Authors }}<meta name=\"author\" content=\"{{ .Authors }}\">\n{{ end }}" +
		"{{ if .CSS}}<link type=\"text/css\" rel=\"stylesheet\" href=\"{{ .CSS }}\">\n{{ end }}" +
		"<title>{{ .Title }}</title>\n" +
		"</head>\n" +
		"<body" +
		"{{ if .ID }} id=\"{{ .ID }}\"{{ end }}" +
		" class=\"{{ .Doctype }}{{ if .Roles }} {{ .Roles }}{{ end }}\">\n" +
		"{{ if .IncludeHeader }}{{ .Header }}{{ end }}" +
		"<div id=\"content\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"{{ if .IncludeFooter }}<div id=\"footer\">\n" +
		"<div id=\"footer-text\">\n" +
		"{{ if .RevNumber }}Version {{ .RevNumber }}<br>\n{{ end }}" +
		"Last updated {{ .LastUpdated }}\n" +
		"</div>\n" +
		"</div>\n{{ end }}" +
		"</body>\n" +
		"</html>\n"

	articleHeaderTmpl = "<div id=\"header\">\n" +
		"<h1>{{ .Header }}</h1>\n" +
		"{{ if.Details }}{{ .Details }}{{ end }}" +
		"</div>\n"

	manpageHeaderTmpl = "{{ if.IncludeH1 }}<div id=\"header\">\n" +
		"<h1>{{ .Header }} Manual Page</h1>\n{{ end }}" +
		"<h2 id=\"_name\">{{ .Name }}</h2>\n" +
		"<div class=\"sectionbody\">\n" +
		"{{ .Content }}" +
		"</div>\n" +
		"{{ if .IncludeH1 }}</div>\n{{ end }}"
)

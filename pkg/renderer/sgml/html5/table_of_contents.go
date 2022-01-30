package html5

const (
	tocRootTmpl = "<div id=\"toc\" class=\"toc\">\n" +
		"<div id=\"toctitle\">{{ .Title }}</div>\n" +
		"{{ .Sections }}" +
		"</div>\n"

	tocSectionTmpl = "<ul class=\"sectlevel{{ .Level }}\">\n{{ .Content }}</ul>\n"

	tocEntryTmpl = "<li><a href=\"#{{ .ID }}\">{{ .Title }}</a>" +
		"{{ if .Content }}\n{{ .Content }}{{ end }}</li>\n"
)

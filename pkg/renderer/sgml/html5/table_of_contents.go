package html5

const (
	tocRootTmpl = "<div id=\"toc\" class=\"toc\">\n" +
		"<div id=\"toctitle\">Table of Contents</div>\n" +
		"{{ . }}\n" +
		"</div>"

	tocSectionTmpl = "<ul class=\"sectlevel{{ .Level }}\">\n{{ .Content }}</ul>"

	tocEntryTmpl = "<li><a href=\"#{{ .ID }}\">{{ .Title }}</a>" +
		"{{ if .Content }}\n{{ .Content }}\n{{ end }}</li>\n"
)

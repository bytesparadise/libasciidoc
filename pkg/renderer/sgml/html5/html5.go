package html5

const (
	articleTmpl = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">{{ if .Generator }}
<meta name="generator" content="{{ .Generator }}">{{ end }}{{ if .Authors }}
<meta name="author" content="{{ .Authors }}">{{ end }}{{ if .CSS}}
<link type="text/css" rel="stylesheet" href="{{ .CSS }}">{{ end }}
<title>{{ .Title }}</title>
</head>
<body class="{{ .Doctype }}{{ if .Role }} {{ .Role }}{{ end }}">{{ if .IncludeHeader }}
{{ .Header }}{{ end }}
<div id="content">
{{ .Content }}
</div>{{ if .IncludeFooter }}
<div id="footer">
<div id="footer-text">{{ if .RevNumber }}
Version {{ .RevNumber }}<br>{{ end }}
Last updated {{ .LastUpdated }}
</div>
</div>{{ end }}
</body>
</html>`

	articleHeaderTmpl = `<div id="header">
<h1>{{ .Header }}</h1>{{ if .Details }}
{{ .Details }}{{ end }}
</div>`

	manpageHeaderTmpl = `{{ if .IncludeH1 }}<div id="header">
<h1>{{ .Header }} Manual Page</h1>
{{ end }}<h2 id="_name">{{ .Name }}</h2>
<div class="sectionbody">
{{ .Content }}
</div>{{ if .IncludeH1 }}
</div>{{ end }}`
)

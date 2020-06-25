package xhtml5

const (
	articleTmpl = `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>{{ if .Generator }}
<meta name="generator" content="{{ .Generator }}"/>{{ end }}{{ if .Authors }}
<meta name="author" content="{{ .Authors }}"/>{{ end }}{{ if .CSS}}
<link type="text/css" rel="stylesheet" href="{{ .CSS }}"/>{{ end }}
<title>{{ .Title }}</title>
</head>
<body class="{{ .Doctype }}{{ if .Role }} {{ .Role }}{{ end }}">{{ if .IncludeHeader }}
{{ .Header }}{{ end }}
<div id="content">
{{ .Content }}
</div>{{ if .IncludeFooter }}
<div id="footer">
<div id="footer-text">{{ if .RevNumber }}
Version {{ .RevNumber }}<br/>{{ end }}
Last updated {{ .LastUpdated }}
</div>
</div>{{ end }}
</body>
</html>`
)

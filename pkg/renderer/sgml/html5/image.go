package html5

const (
	blockImageTmpl = `<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="imageblock{{ if .Roles }} {{ .Roles }}{{ end }}">
<div class="content">
{{ if ne .Href "" }}<a class="image" href="{{ .Href }}">{{ end }}<img src="{{ .Src }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}>{{ if ne .Href "" }}</a>{{ end }}
</div>{{ if .Title }}
<div class="title">{{ .Caption }}{{ .Title }}</div>
{{ else }}
{{ end }}</div>
`
	inlineImageTmpl = `<span class="image{{ if .Roles }} {{ .Roles }}{{ end }}">{{ if ne .Href "" }}<a class="image" href="{{ .Href }}">{{ end }}<img src="{{ .Src }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}{{ if .Title }} title="{{ .Title }}"{{ end }}>{{ if ne .Href "" }}</a>{{ end }}</span>`
)

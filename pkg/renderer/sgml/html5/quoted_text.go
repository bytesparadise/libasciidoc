package html5

const (
	boldTextTmpl        = `<strong{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .Roles }} class="{{ .Roles }}"{{ end }}>{{ .Content }}</strong>`
	italicTextTmpl      = `<em{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .Roles }} class="{{ .Roles }}"{{ end }}>{{ .Content }}</em>`
	monospaceTextTmpl   = `<code{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .Roles }} class="{{ .Roles }}"{{ end }}>{{ .Content }}</code>`
	subscriptTextTmpl   = `<sub{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .Roles }} class="{{ .Roles }}"{{ end }}>{{ .Content }}</sub>`
	superscriptTextTmpl = `<sup{{ if .ID }} id="{{ .ID }}"{{ end }}{{ if .Roles }} class="{{ .Roles }}"{{ end }}>{{ .Content }}</sup>`
)

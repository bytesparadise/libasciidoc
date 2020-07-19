package html5

const (
	specialCharacterTmpl = `{{ if eq .Content ">" }}&gt;{{ else if eq .Content "<" }}&lt;{{ else if eq .Content "&" }}&amp;{{ else if eq .Content "+" }}&#43;{{ end }}`
)

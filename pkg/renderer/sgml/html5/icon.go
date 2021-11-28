package html5

const (
	// Inline icons are presented as the icon as an image, or font icon, or the text "[class]" where class
	// is the Alt text, and defaults to the icon type.  Unlike asciidoctor, we do place class settings on
	// the image to match rotate, flip, and size, allowing images to be manipulated with css style just
	// like the icons can.  We also allow for an ID to be placed on the icon, as another enhancement.
	inlineIconTmpl = `<span {{ if .ID }}id="{{ .ID }}" {{ end}}class="icon{{ if .Role }} {{ .Role }}{{ end }}">` +
		`{{ if .Link }}<a class="image" href="{{ .Link }}"{{ if .Window }} target="{{ .Window }}"{{ end }}>{{ end }}` +
		`{{ .Icon }}` +
		`{{ if .Link }}</a>{{ end }}` +
		`</span>`

	iconImageTmpl = `<img src="{{ .Src }}"` +
		`{{ if .Alt }} alt="{{ .Alt }}{{ end }}"` +
		`{{ if .Width }} width="{{ .Width }}"{{ end }}` +
		`{{ if .Height }} height="{{ .Height }}"{{ end }}` +
		`{{ if .Title }} title="{{ .Title }}"{{ end }}` +
		`{{ if or .Size .Rotate .Flip }} class="` +
		`{{ if .Size }}fa-{{ .Size }}{{ end }}` +
		`{{ if .Rotate }}{{ if .Size }} {{ end }}fa-rotate-{{ .Rotate }}{{ end }}` +
		`{{ if .Flip }}{{ if or .Size .Rotate }} {{ end }}fa-flip-{{ .Flip }}{{ end }}` +
		`"{{ end }}>`

	iconFontTmpl = `<i class="fa` +
		`{{ if .Admonition }} icon-{{ .Class }}{{ else }} fa-{{ .Class }}{{ end }}` +
		`{{ if .Size }} fa-{{ .Size }}{{ end }}` +
		`{{ if .Rotate }} fa-rotate-{{ .Rotate }}{{ end }}` +
		`{{ if .Flip }} fa-flip-{{ .Flip }}{{ end }}"` +
		`{{ if .Title }} title="{{ .Title }}"{{ end }}></i>`

	iconTextTmpl = `{{ if .Admonition }}<div class="title">{{ .Alt }}</div>{{ else }}[{{ .Alt }}]{{ end }}`
)

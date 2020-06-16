package xhtml5

const (
	// Inline icons are presented as the icon as an image, or font icon, or the text "[class]" where class
	// is the Alt text, and defaults to the icon type.  Unlike asciidoctor, we do place class settings on
	// the image to match rotate, flip, and size, allowing images to be manipulated with css style just
	// like the icons can.  We also allow for an ID to be placed on the icon, as another enhancement.
	//
	// Only the img tag needs to be made XHTML safe.

	iconImageTmpl = `<img src="{{ .Path }}"` +
		`{{ if .Alt }} alt="{{ .Alt }}{{ end }}"` +
		`{{ if .Width }} width="{{ .Width }}"{{ end }}` +
		`{{ if .Height }} height="{{ .Height }}"{{ end }}` +
		`{{ if .Title }} title="{{ .Title }}"{{ end }}` +
		`{{ if or .Size .Rotate .Flip }} class="` +
		`{{ if .Size }}fa-{{ .Size }}{{ end }}` +
		`{{ if .Rotate }}{{ if .Size }} {{ end }}fa-rotate-{{ .Rotate }}{{ end }}` +
		`{{ if .Flip }}{{ if or .Size .Rotate }} {{ end }}fa-flip-{{ .Flip }}{{ end }}` +
		`"{{ end }}/>`
)

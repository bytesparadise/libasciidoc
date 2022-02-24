package html5

// initializes the sgml
const (
	inlineMenuTmpl =
	// eg: `<b class="menuref">File</b>`
	`{{ if len .Path | eq 1 }}<b class="menuref">{{ index .Path 0 }}</b>` +
		// eg: `<span class="menuseq"><b class="menu">File</b>&#160;<b class="caret">&#8250;</b> <b class="submenu">Zoom</b>&#160;<b class="caret">&#8250;</b> <b class="menuitem">Reset</b></span>`
		`{{ else }}` +
		`<span class="menuseq">` +
		`{{ with $path := .Path }}` +
		`{{ range $index, $element := $path }}` +
		`{{ if eq $index 0 }}<b class="menu">{{ $element }}</b>` +
		`{{ else if lastInStrings $path $index }}&#160;<b class="caret">&#8250;</b> <b class="menuitem">{{ $element }}</b>` +
		`{{ else }}&#160;<b class="caret">&#8250;</b> <b class="submenu">{{ $element }}</b>` +
		`{{ end }}` +
		`{{ end }}` +
		`</span>` +
		`{{ end }}` +
		`{{ end }}`
)

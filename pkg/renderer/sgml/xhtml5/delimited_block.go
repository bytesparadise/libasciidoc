package xhtml5

const (
	quoteBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderElements $ctx .Elements | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br/>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	verseBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ range $index, $element := .Elements }}{{ renderVerse $ctx $element | printf "%s" }}{{ end }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br/>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`
)

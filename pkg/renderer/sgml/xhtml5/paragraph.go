package xhtml5

const (
	verseParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ renderLines $ctx .Lines plainText }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br/>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	quoteParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderLines $ctx .Lines }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br/>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`
)

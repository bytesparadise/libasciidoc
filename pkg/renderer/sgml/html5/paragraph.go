package html5

const (
	paragraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines .HardBreaks | printf "%s" }}<div {{ if ne .ID "" }}id="{{ .ID }}" {{ end }}class="{{ .Class }}">{{ if ne .Title "" }}
<div class="doctitle">{{ escape .Title }}</div>{{ end }}
<p>{{ $renderedLines }}</p>
</div>{{ end }}`

	admonitionParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines | printf "%s" }}{{ if ne $renderedLines "" }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
{{ if .IconClass }}<i class="fa icon-{{ .IconClass }}" title="{{ .IconTitle }}"></i>{{ else }}<div class="title">{{ .IconTitle }}</div>{{ end }}
</td>
<td class="content">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
{{ $renderedLines }}
</td>
</tr>
</table>
</div>{{ end }}{{ end }}`

	delimitedBlockParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<p>{{ .CheckStyle }}{{ renderLines $ctx .Lines | printf "%s" }}</p>{{ end }}`

	sourceParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div class="listingblock">
<div class="content">
<pre class="highlight">{{ if .Language }}<code class="language-{{ .Language }}" data-lang="{{ .Language }}">{{ else }}<code>{{ end }}{{ renderLines $ctx .Lines | printf "%s" }}</code></pre>
</div>
</div>{{ end }}`

	verseParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ renderLines $ctx .Lines plainText | printf "%s" }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	quoteParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderLines $ctx .Lines | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	manpageNameParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ $renderedLines := renderLines $ctx .Lines | printf "%s" }}<p>{{ $renderedLines }}</p>{{ end }}`
)

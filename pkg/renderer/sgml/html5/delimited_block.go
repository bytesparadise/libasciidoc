package html5

const (
	fencedBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="highlight"><code>{{ render $ctx .Elements | printf "%s" }}</code></pre>
</div>
</div>{{ end }}`

	listingBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre>{{ renderElements $ctx .Elements | printf "%s" }}</pre>
</div>
</div>{{ end }}`

	sourceBlockTmpl = `<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="listingblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
<pre class="{{ if .SyntaxHighlighter }}{{ .SyntaxHighlighter }} {{ end }}highlight"><code{{ if .Language }}{{ if not .SyntaxHighlighter }} class="language-{{ .Language}}"{{ end }} data-lang="{{ .Language}}"{{ end }}>{{ .Content }}</code></pre>
</div>
</div>`

	sourceBlockContentTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ render $ctx .Elements | printf "%s" }}{{ end }}`

	exampleBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="exampleblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<div class="content">
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`

	quoteBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="quoteblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<blockquote>
{{ renderElements $ctx .Elements | printf "%s" }}
</blockquote>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	verseBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="verseblock">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
<pre class="content">{{ range $index, $element := .Elements }}{{ renderVerse $ctx $element | printf "%s" }}{{ end }}</pre>{{ if .Attribution.First }}
<div class="attribution">
&#8212; {{ .Attribution.First }}{{ if .Attribution.Second }}<br>
<cite>{{ .Attribution.Second }}</cite>{{ end }}
</div>{{ end }}
</div>{{ end }}`

	verseBlockParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}{{ renderLines $ctx .Lines | printf "%s" }}{{ end }}`

	admonitionBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID}}" {{ end }}class="admonitionblock {{ .Class }}">
<table>
<tr>
<td class="icon">
{{ if .IconClass }}<i class="fa icon-{{ .IconClass }}" title="{{ .IconTitle }}"></i>{{ else }}<div class="title">{{ .IconTitle }}</div>{{ end }}
</td>
<td class="content">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}{{ renderElements $ctx .Elements | printf "%s" }}
</td>
</tr>
</table>
</div>{{ end }}`

	sidebarBlockTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="sidebarblock">
<div class="content">{{ if .Title }}
<div class="title">{{ escape .Title }}</div>{{ end }}
{{ renderElements $ctx .Elements | printf "%s" }}
</div>
</div>{{ end }}`

	// the name here is weird because "pass" as a prefix triggers a false security warning
	pssThroughBlock = `{{ $ctx := .Context }}{{ with .Data }}{{ render $ctx .Elements | printf "%s" }}{{ end }}`
)

package html5

const (
	documentDetailsTmpl = `<div class="details">{{ if .Authors }}
{{ .Authors }}{{ end }}{{ if .RevNumber }}
<span id="revnumber">version {{ .RevNumber }}{{ if .RevDate }},{{ end }}</span>{{ end }}{{ if .RevDate }}
<span id="revdate">{{ .RevDate }}</span>{{ end }}{{ if .RevRemark }}
<br><span id="revremark">{{ .RevRemark }}</span>{{ end }}
</div>`

	documentAuthorDetailsTmpl = `{{ if .Name }}<span id="author{{ .Index }}" class="author">{{ .Name }}</span><br>{{ end }}{{ if .Email }}
<span id="email{{ .Index }}" class="email"><a href="mailto:{{ .Email }}">{{ .Email }}</a></span><br>{{ end }}`
)

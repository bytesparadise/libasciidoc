package html5

const (
	internalCrossReferenceTmpl = `<a href="#{{ .Href }}">{{ .Label }}</a>`
	externalCrossReferenceTmpl = `<a href="{{ .Href }}">{{ .Label }}</a>`
)

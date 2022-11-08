package html5

const (
	internalCrossReferenceTmpl = `<a href="#{{ toLower .Href }}">{{ .Label }}</a>`
	externalCrossReferenceTmpl = `<a href="{{ .Href }}">{{ .Label }}</a>`
)

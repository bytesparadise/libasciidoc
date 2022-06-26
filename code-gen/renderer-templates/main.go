package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"text/template"
	"unicode"
)

const (
	sgmlRenderer = `package sgml

import (
	"sync"
	text "text/template"
)

type sgmlRenderer struct {
	    templates   Templates
		functions   text.FuncMap
{{ range  $i, $tmpl := . }}
		{{ once $tmpl}} sync.Once
		{{ tmpl $tmpl}} *text.Template
{{ end }}
}

{{ range  $i, $tmpl := . }}
func (r *sgmlRenderer) {{ func $tmpl }} (*text.Template, error) {
	var err error
	r.{{ once $tmpl }}.Do(func() {
		r.{{ tmpl $tmpl }}, err = r.newTemplate("{{ $tmpl }}", r.templates.{{ $tmpl }}, err)
	})
	return r.{{ tmpl $tmpl }}, err
}

{{ end }}
`
)

func main() {
	// read the content of pkg/renderer/sgml/templates.go
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "pkg/renderer/sgml/templates.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	v := &visitor{
		fields: []string{},
	}
	ast.Walk(v, f)
	tmpl, err := template.New("templates").Funcs(template.FuncMap{
		"once": func(s string) string {
			return string(unicode.ToLower(rune(s[0]))) + s[1:] + "Once"
		},
		"func": func(s string) string {
			return string(unicode.ToLower(rune(s[0]))) + s[1:] + "()"
		},
		"tmpl": func(s string) string {
			return string(unicode.ToLower(rune(s[0]))) + s[1:] + "Tmpl"
		},
	}).Parse(sgmlRenderer)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("templates: %v\n", v.fields)
	result := &bytes.Buffer{}
	if err := tmpl.Execute(result, v.fields); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result.String())
}

type visitor struct {
	fields []string
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	if f, ok := n.(*ast.Field); ok {
		v.fields = append(v.fields, f.Names[0].Name)
	}
	return v
}

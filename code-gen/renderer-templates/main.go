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
	"fmt"
	"sync"
	"strings"
	texttemplate "text/template"

	log "github.com/sirupsen/logrus"
)

type sgmlRenderer struct {
	    templates   Templates
		functions   texttemplate.FuncMap
{{ range  $i, $tmpl := . }}
		{{ once $tmpl}} sync.Once
		{{ tmpl $tmpl}} *texttemplate.Template
{{ end }}
}

type template func() (*texttemplate.Template, error)

func (r *sgmlRenderer) execute(tmpl template, data interface{}) (string, error) {
	result := &strings.Builder{}
	t, err := tmpl()
	if err != nil {
		return "", err
	}
	if err := t.Execute(result, data); err != nil {
		return "", err
	}
	return result.String(), nil
}

func (r *sgmlRenderer) newTemplate(name string, tmpl string, err error) (*texttemplate.Template, error) {
	// NB: if the data is missing below, it will be an empty string.
	if err != nil {
		return nil, err
	}
	if len(tmpl) == 0 {
		return nil, fmt.Errorf("empty template for '%s'", name)
	}
	t := texttemplate.New(name)
	t.Funcs(r.functions)
	if t, err = t.Parse(tmpl); err != nil {
		log.Errorf("failed to initialize the '%s' template: %v", name, err)
		return nil, err
	}
	return t, nil
}

{{ range  $i, $tmpl := . }}
func (r *sgmlRenderer) {{ func $tmpl }} (*texttemplate.Template, error) {
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

package html5_test

import (
	"html"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	. "github.com/onsi/ginkgo"
)

var helloMacroTmpl *texttemplate.Template

var _ = Describe("user macros", func() {

	Context("user macros", func() {
		It("undefined macro block", func() {

			actualContent := "hello::[]"
			expectedResult := `<div class="paragraph">
<p>hello::[]</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("user macro block", func() {

			actualContent := "hello::[]"
			expectedResult := `<div class="helloblock">
<div class="content">
<span>hello world</span>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("user macro block with attribute", func() {

			actualContent := `hello::[suffix="!!!!"]`
			expectedResult := `<div class="helloblock">
<div class="content">
<span>hello world!!!!</span>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("user macro block with value", func() {

			actualContent := `hello::John Doe[]`
			expectedResult := `<div class="helloblock">
<div class="content">
<span>hello John Doe</span>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("user macro block with value and attributes", func() {

			actualContent := `hello::John Doe[prefix="Hi ",suffix="!!"]`
			expectedResult := `<div class="helloblock">
<div class="content">
<span>Hi John Doe!!</span>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("undefined inline macro", func() {

			actualContent := "hello:[]"
			expectedResult := `<div class="paragraph">
<p>hello:[]</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("inline macro", func() {

			actualContent := "AAA hello:[]"
			expectedResult := `<div class="paragraph">
<p>AAA <span>hello world</span></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("inline macro with attribute", func() {

			actualContent := `AAA hello:[suffix="!!!!!"]`
			expectedResult := `<div class="paragraph">
<p>AAA <span>hello world!!!!!</span></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("inline macro with value", func() {

			actualContent := `AAA hello:John Doe[]`
			expectedResult := `<div class="paragraph">
<p>AAA <span>hello John Doe</span></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

		It("inline macro with value and attributes", func() {

			actualContent := `AAA hello:John Doe[prefix="Hi ",suffix="!!"]`
			expectedResult := `<div class="paragraph">
<p>AAA <span>Hi John Doe!!</span></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl))
		})

	})
})

func init() {
	t := texttemplate.New("hello")
	t.Funcs(texttemplate.FuncMap{
		"escape": html.EscapeString,
	})
	helloMacroTmpl = texttemplate.Must(t.Parse(`{{- if eq .Kind "block" -}}
<div class="helloblock">
<div class="content">
{{end -}}
<span>
{{- if .Attributes.Has "prefix"}}{{escape (.Attributes.GetAsString "prefix")}} {{else}}hello {{end -}}
{{- if ne .Value ""}}{{escape .Value}}{{else}}world{{- end -}}
{{- escape (.Attributes.GetAsString "suffix") -}}
</span>
{{- if eq .Kind "block"}}
</div>
</div>
{{- end -}}`))
}

package html5_test

import (
	"html"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var helloMacroTmpl *texttemplate.Template

var _ = Describe("user macros", func() {

	Context("user macros", func() {

		It("undefined macro block", func() {

			source := "hello::[]"
			expected := `<div class="paragraph">
<p>hello::[]</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("user macro block", func() {

			source := "hello::[]"
			expected := `<div class="helloblock">
<div class="content">
<span>hello world</span>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("user macro block with attribute", func() {

			source := `hello::[suffix="!!!!"]`
			expected := `<div class="helloblock">
<div class="content">
<span>hello world!!!!</span>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("user macro block with value", func() {

			source := `hello::JohnDoe[]`
			expected := `<div class="helloblock">
<div class="content">
<span>hello JohnDoe</span>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("user macro block with value and attributes", func() {

			source := `hello::JohnDoe[prefix="Hi ",suffix="!!"]`
			expected := `<div class="helloblock">
<div class="content">
<span>Hi JohnDoe!!</span>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("undefined inline macro", func() {

			source := "hello:[]"
			expected := `<div class="paragraph">
<p>hello:[]</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("inline macro", func() {

			source := "AAA hello:[]"
			expected := `<div class="paragraph">
<p>AAA <span>hello world</span></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("inline macro with attribute", func() {

			source := `AAA hello:[suffix="!!!!!"]`
			expected := `<div class="paragraph">
<p>AAA <span>hello world!!!!!</span></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("inline macro with value", func() {

			source := `AAA hello:JohnDoe[]`
			expected := `<div class="paragraph">
<p>AAA <span>hello JohnDoe</span></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
		})

		It("inline macro with value and attributes", func() {

			source := `AAA hello:JohnDoe[prefix="Hi ",suffix="!!"]`
			expected := `<div class="paragraph">
<p>AAA <span>Hi JohnDoe!!</span></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected, renderer.DefineMacro(helloMacroTmpl.Name(), helloMacroTmpl)))
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

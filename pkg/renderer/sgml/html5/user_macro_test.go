package html5_test

import (
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var helloMacroTmpl *texttemplate.Template

var _ = Describe("user macros", func() {

	Context("user macros", func() {

		It("undefined macro block", func() {

			source := "hello::[]"
			expected := `<div class="paragraph">
<p>hello::[]</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("user macro block", func() {

			source := "hello::[]"
			expected := `<div class="helloblock">
<div class="content">
<span>hello world</span>
</div>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("user macro block with attribute", func() {

			source := `hello::[suffix="!!!!"]`
			expected := `<div class="helloblock">
<div class="content">
<span>hello world!!!!</span>
</div>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("user macro block with value", func() {

			source := `hello::JohnDoe[]`
			expected := `<div class="helloblock">
<div class="content">
<span>hello JohnDoe</span>
</div>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("user macro block with value and attributes", func() {

			source := `hello::JohnDoe[prefix="Hi ",suffix="!!"]`
			expected := `<div class="helloblock">
<div class="content">
<span>Hi JohnDoe!!</span>
</div>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("undefined inline macro", func() {

			source := "hello:[]"
			expected := `<div class="paragraph">
<p>hello:[]</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("inline macro", func() {

			source := "AAA hello:[]"
			expected := `<div class="paragraph">
<p>AAA <span>hello world</span></p>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("inline macro with attribute", func() {

			source := `AAA hello:[suffix="!!!!!"]`
			expected := `<div class="paragraph">
<p>AAA <span>hello world!!!!!</span></p>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("inline macro with value", func() {

			source := `AAA hello:JohnDoe[]`
			expected := `<div class="paragraph">
<p>AAA <span>hello JohnDoe</span></p>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

		It("inline macro with value and attributes", func() {

			source := `AAA hello:JohnDoe[prefix="Hi ",suffix="!!"]`
			expected := `<div class="paragraph">
<p>AAA <span>Hi JohnDoe!!</span></p>
</div>`
			Expect(RenderHTML(source, configuration.WithMacroTemplate(helloMacroTmpl.Name(), helloMacroTmpl))).To(Equal(expected))
		})

	})
})

func init() {
	t := texttemplate.New("hello")
	t.Funcs(texttemplate.FuncMap{
		"escape": sgml.EscapeString,
	})
	helloMacroTmpl = texttemplate.Must(t.Parse(`{{- if eq .Kind "block" -}}
<div class="helloblock">
<div class="content">
{{end -}}
<span>
{{- $prefix := index .Attributes "prefix" }}{{ $suffix := index .Attributes "suffix" }}{{ if $prefix }}{{- escape ($prefix) }} {{ else }}hello {{end -}}
{{- if ne .Value "" }}{{ escape .Value }}{{ else }}world{{- end -}}
{{- if $suffix }}{{ escape $suffix -}}{{ end -}}
</span>
{{- if eq .Kind "block"}}
</div>
</div>
{{- end -}}`))
}

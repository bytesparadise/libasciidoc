package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("passthroughs", func() {

	Context("tripleplus passthrough", func() {

		It("an empty standalone tripleplus passthrough", func() {
			source := `++++++`
			expected := ``
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("an empty tripleplus passthrough in a paragraph", func() {
			source := `++++++ with more content afterwards...`
			expected := `<div class="paragraph">
<p> with more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("a standalone tripleplus passthrough", func() {
			source := `+++*bold content*+++`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("tripleplus passthrough in paragraph", func() {
			source := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := `<div class="paragraph">
<p>The text <u>underline & me</u> is underlined.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("singleplus Passthrough", func() {

		It("an empty standalone singleplus passthrough", func() {
			source := `++`
			expected := `<div class="paragraph">
<p>&#43;&#43;</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("an empty singleplus passthrough in a paragraph", func() {
			source := `++ with more content afterwards...`
			expected := `<div class="paragraph">
<p>&#43;&#43; with more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("a singleplus passthrough", func() {
			source := `+*bold content*+`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("singleplus passthrough in paragraph", func() {
			source := `The text +<u>underline me</u>+ is not underlined.`
			expected := `<div class="paragraph">
<p>The text &lt;u&gt;underline me&lt;/u&gt; is not underlined.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("invalid singleplus passthrough in paragraph", func() {
			source := `The text + *hello*, world + is not passed through.`
			expected := `<div class="paragraph">
<p>The text &#43; <strong>hello</strong>, world &#43; is not passed through.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("passthrough Macro", func() {

		It("passthrough macro with single word", func() {
			source := `pass:[hello]`
			expected := `<div class="paragraph">
<p>hello</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("passthrough macro with words", func() {
			source := `pass:[hello, world]`
			expected := `<div class="paragraph">
<p>hello, world</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("empty passthrough macro", func() {
			source := `pass:[]`
			expected := ``
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("passthrough macro with spaces", func() {
			source := `pass:[ *hello*, world ]`
			expected := `<div class="paragraph">
<p> *hello*, world </p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("passthrough macro with line break", func() {
			source := "pass:[hello,\nworld]"
			expected := `<div class="paragraph">
<p>hello,
world</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("passthrough Macro with Quoted Text", func() {

		It("passthrough macro with single quoted word", func() {
			source := `pass:q[*hello*]`
			expected := `<div class="paragraph">
<p><strong>hello</strong></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("passthrough macro with quoted word in sentence and trailing spaces", func() {
			source := `pass:q[ a *hello*, world ]   `
			expected := `<div class="paragraph">
<p> a <strong>hello</strong>, world </p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("passthrough macro within paragraph", func() {
			source := `an pass:q[ *hello*, world ] mention`
			expected := `<div class="paragraph">
<p>an  <strong>hello</strong>, world  mention</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})
})

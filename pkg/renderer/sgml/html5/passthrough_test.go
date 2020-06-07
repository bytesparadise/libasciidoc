package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("passthroughs", func() {

	Context("tripleplus passthrough", func() {

		It("an empty standalone tripleplus passthrough", func() {
			// here it differs from Asciidoctor which returns no content but reports an error ("unterminated pass block")
			source := `++++++`
			expected := `<div class="paragraph">
<p></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("an empty tripleplus passthrough in a paragraph", func() {
			source := `++++++ with more content afterwards...`
			expected := `<div class="paragraph">
<p> with more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a standalone tripleplus passthrough", func() {
			source := `+++*bold content*+++`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in paragraph", func() {
			source := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := `<div class="paragraph">
<p>The text <u>underline & me</u> is underlined.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("singleplus InlinePassthrough", func() {

		It("an empty standalone singleplus passthrough", func() {
			source := `++`
			expected := `<div class="paragraph">
<p>++</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("an empty singleplus passthrough in a paragraph", func() {
			source := `++ with more content afterwards...`
			expected := `<div class="paragraph">
<p>++ with more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a singleplus passthrough", func() {
			source := `+*bold content*+`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in paragraph", func() {
			source := `The text +<u>underline me</u>+ is not underlined.`
			expected := `<div class="paragraph">
<p>The text &lt;u&gt;underline me&lt;/u&gt; is not underlined.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid singleplus passthrough in paragraph", func() {
			source := `The text + *hello*, world + is not passed through.`
			expected := `<div class="paragraph">
<p>The text + <strong>hello</strong>, world + is not passed through.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("passthrough Macro", func() {

		It("passthrough macro with single word", func() {
			source := `pass:[hello]`
			expected := `<div class="paragraph">
<p>hello</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("passthrough macro with words", func() {
			source := `pass:[hello, world]`
			expected := `<div class="paragraph">
<p>hello, world</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty passthrough macro", func() {
			source := `pass:[]`
			expected := `<div class="paragraph">
<p></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("passthrough macro with spaces", func() {
			source := `pass:[ *hello*, world ]`
			expected := `<div class="paragraph">
<p> *hello*, world </p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("passthrough macro with line break", func() {
			source := "pass:[hello,\nworld]"
			expected := `<div class="paragraph">
<p>hello,
world</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("passthrough Macro with Quoted Text", func() {

		It("passthrough macro with single quoted word", func() {
			source := `pass:q[*hello*]`
			expected := `<div class="paragraph">
<p><strong>hello</strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("passthrough macro with quoted word in sentence and trailing spaces", func() {
			source := `pass:q[ a *hello*, world ]   `
			expected := `<div class="paragraph">
<p> a <strong>hello</strong>, world </p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("passthrough macro within paragraph", func() {
			source := `an pass:q[ *hello*, world ] mention`
			expected := `<div class="paragraph">
<p>an  <strong>hello</strong>, world  mention</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

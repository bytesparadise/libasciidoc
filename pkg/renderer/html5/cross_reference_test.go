package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cross-references", func() {

	Context("section reference", func() {

		It("cross-reference with custom id", func() {

			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thetitle">a title</a>!</p>
</div>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("cross-reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thetitle">a label to the title</a>!</p>
</div>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("invalid section reference", func() {

			source := `[[thetitle]]
== a title

with some content linked to <<thewrongtitle>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thewrongtitle">[thewrongtitle]</a>!</p>
</div>
</div>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})
})

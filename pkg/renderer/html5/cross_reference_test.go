package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("cross references", func() {

	Context("internal references", func() {

		It("cross reference with custom id", func() {

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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("cross reference with custom id and label", func() {
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("external references", func() {

		It("external cross reference to other doc with plain text location and rich label", func() {
			source := `some content linked to xref:another-doc.adoc[*another doc*]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="another-doc.html"><strong>another doc</strong></a>!</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("external cross reference to other doc with document attribute in location and label with special chars", func() {
			source := `:foo: foo-doc
some content linked to xref:{foo}.adoc[another_doc()]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="foo-doc.html">another_doc()</a>!</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

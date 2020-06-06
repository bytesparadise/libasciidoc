package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("strings", func() {

	It("text with ellipsis", func() {
		source := `some text...`
		expected := `<div class="paragraph">
<p>some text&#8230;&#8203;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with copyright", func() {
		source := `Copyright (C)`
		expected := `<div class="paragraph">
<p>Copyright &#169;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with trademark", func() {
		source := `TheRightThing(TM)`
		expected := `<div class="paragraph">
<p>TheRightThing&#153;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with registered", func() {
		source := `TheRightThing(R)`
		expected := `<div class="paragraph">
<p>TheRightThing&#174;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with registered", func() {
		// We will often want to use these symbols in headers.
		source := `== Registered(R)`
		expected := `<div class="sect1">
<h2 id="_registered_r">Registered&#174;</h2>
<div class="sectionbody">
</div>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

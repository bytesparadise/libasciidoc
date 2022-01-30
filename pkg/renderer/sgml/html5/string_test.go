package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("strings", func() {

	It("text with ellipsis", func() {
		source := `some text...`
		expected := `<div class="paragraph">
<p>some text&#8230;&#8203;</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with copyright", func() {
		source := `Copyright (C)`
		expected := `<div class="paragraph">
<p>Copyright &#169;</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with trademark", func() {
		source := `TheRightThing(TM)`
		expected := `<div class="paragraph">
<p>TheRightThing&#8482;</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with registered", func() {
		source := `TheRightThing(R)`
		expected := `<div class="paragraph">
<p>TheRightThing&#174;</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with implicit apostrophe", func() {
		source := `Mother's Day`
		expected := `<div class="paragraph">
<p>Mother&#8217;s Day</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with implicit apostrophe no match", func() {
		source := `Mothers' Day`
		expected := `<div class="paragraph">
<p>Mothers' Day</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with explicit apostrophe no match", func() {
		source := "Mothers`' Day"
		expected := `<div class="paragraph">
<p>Mothers&#8217; Day</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with registered", func() {
		// We will often want to use these symbols in headers.
		source := `== Registered(R)`
		expected := `<div class="sect1">
<h2 id="_registered">Registered&#174;</h2>
<div class="sectionbody">
</div>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with registered unicode", func() {
		// We will often want to use these symbols in headers.
		source := ":unicode:\n\n== Registered(R)"
		expected := "<div class=\"sect1\">\n" +
			"<h2 id=\"_registered\">Registered\u00ae</h2>\n" +
			"<div class=\"sectionbody\">\n" +
			"</div>\n" +
			"</div>\n"
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with explicit apostrophe", func() {
		source := "== It`'s A Wonderful Life"
		expected := "<div class=\"sect1\">\n" +
			"<h2 id=\"_it_s_a_wonderful_life\">It&#8217;s A Wonderful Life</h2>\n" +
			"<div class=\"sectionbody\">\n" +
			"</div>\n" +
			"</div>\n"
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with explicit apostrophe (unicode)", func() {
		source := ":unicode:\n\n== It`'s A Wonderful Life"
		expected := "<div class=\"sect1\">\n" +
			"<h2 id=\"_it_s_a_wonderful_life\">It\u2019s A Wonderful Life</h2>\n" +
			"<div class=\"sectionbody\">\n" +
			"</div>\n" +
			"</div>\n"
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with implicit apostrophe", func() {
		source := "== It's A Wonderful Life"
		expected := "<div class=\"sect1\">\n" +
			"<h2 id=\"_it_s_a_wonderful_life\">It&#8217;s A Wonderful Life</h2>\n" +
			"<div class=\"sectionbody\">\n" +
			"</div>\n" +
			"</div>\n"
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("title with implicit apostrophe (unicode)", func() {
		source := ":unicode:\n\n== It's A Wonderful Life"
		expected := "<div class=\"sect1\">\n" +
			"<h2 id=\"_it_s_a_wonderful_life\">It\u2019s A Wonderful Life</h2>\n" +
			"<div class=\"sectionbody\">\n" +
			"</div>\n" +
			"</div>\n"
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

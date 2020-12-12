package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quoted texts", func() {

	Context("bold content", func() {

		It("bold content alone", func() {
			source := "*bold content*"
			expected := `<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("bold content in sentence", func() {
			source := "some *bold content*."
			expected := `<div class="paragraph">
<p>some <strong>bold content</strong>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("bold content across newline", func() {
			source := "some *bold\ncontent*."
			expected := `<div class="paragraph">
<p>some <strong>bold
content</strong>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("bold content across paragraph", func() {
			source := "some *bold\n\ncontent*."
			expected := `<div class="paragraph">
<p>some *bold</p>
</div>
<div class="paragraph">
<p>content*.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double bold content across newline", func() {
			source := "some **bold\ncontent**."
			expected := `<div class="paragraph">
<p>some <strong>bold
content</strong>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double bold content across paragraph", func() {
			source := "some **bold\n\ncontent**."
			expected := `<div class="paragraph">
<p>some **bold</p>
</div>
<div class="paragraph">
<p>content**.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("italic content", func() {

		It("italic content alone", func() {
			source := "_italic content_"
			expected := `<div class="paragraph">
<p><em>italic content</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content in sentence", func() {

			source := "some _italic content_."
			expected := `<div class="paragraph">
<p>some <em>italic content</em>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content across newline", func() {
			source := "some _italic\ncontent_."
			expected := `<div class="paragraph">
<p>some <em>italic
content</em>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content across paragraph", func() {
			source := "some _italic\n\ncontent_."
			expected := `<div class="paragraph">
<p>some _italic</p>
</div>
<div class="paragraph">
<p>content_.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double italic content across newline", func() {
			source := "some __italic\ncontent__."
			expected := `<div class="paragraph">
<p>some <em>italic
content</em>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double italic content across paragraph", func() {
			source := "some __italic\n\ncontent__."
			expected := `<div class="paragraph">
<p>some __italic</p>
</div>
<div class="paragraph">
<p>content__.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("monospace content", func() {

		It("monospace content alone", func() {
			source := "`monospace content`"
			expected := `<div class="paragraph">
<p><code>monospace content</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("monospace content in sentence", func() {

			source := "some `monospace content`."
			expected := `<div class="paragraph">
<p>some <code>monospace content</code>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("monospace content across newline", func() {
			source := "some `monospace\ncontent`."
			expected := `<div class="paragraph">
<p>some <code>monospace
content</code>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("monospace content across paragraph", func() {
			source := "some `monospace\n\ncontent`."
			expected := "<div class=\"paragraph\">\n" +
				"<p>some `monospace</p>\n" +
				"</div>\n" +
				"<div class=\"paragraph\">\n" +
				"<p>content`.</p>\n</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double monospace content across newline", func() {
			source := "some ``monospace\ncontent``."
			expected := `<div class="paragraph">
<p>some <code>monospace
content</code>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double monospace content across paragraph", func() {
			source := "some ``monospace\n\ncontent``."
			expected := "<div class=\"paragraph\">\n" +
				"<p>some ``monospace</p>\n" +
				"</div>\n" +
				"<div class=\"paragraph\">\n" +
				"<p>content``.</p>\n</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("marked content", func() {

		It("marked content alone", func() {
			source := "#marked content#"
			expected := `<div class="paragraph">
<p><mark>marked content</mark></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked content in sentence", func() {

			source := "some #marked content#."
			expected := `<div class="paragraph">
<p>some <mark>marked content</mark>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked content across newline", func() {
			source := "some #marked\ncontent#."
			expected := `<div class="paragraph">
<p>some <mark>marked
content</mark>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked content across paragraph", func() {
			source := "some #marked\n\ncontent#."
			expected := `<div class="paragraph">
<p>some #marked</p>
</div>
<div class="paragraph">
<p>content#.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double marked content across newline", func() {
			source := "some ##marked\ncontent##."
			expected := `<div class="paragraph">
<p>some <mark>marked
content</mark>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("double marked content across paragraph", func() {
			source := "some ##marked\n\ncontent##."
			expected := `<div class="paragraph">
<p>some ##marked</p>
</div>
<div class="paragraph">
<p>content##.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("subscript content", func() {

		It("subscript content alone", func() {
			source := "~subscriptcontent~"
			expected := `<div class="paragraph">
<p><sub>subscriptcontent</sub></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("subscript content in sentence", func() {

			source := "some ~subscriptcontent~."
			expected := `<div class="paragraph">
<p>some <sub>subscriptcontent</sub>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("superscript content", func() {

		It("superscript content alone", func() {
			source := "^superscriptcontent^"
			expected := `<div class="paragraph">
<p><sup>superscriptcontent</sup></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("superscript content in sentence", func() {

			source := "some ^superscriptcontent^."
			expected := `<div class="paragraph">
<p>some <sup>superscriptcontent</sup>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("attributes", func() {

		It("simple role italics", func() {
			source := "[myrole]_italics_"
			expected := `<div class="paragraph">
<p><em class="myrole">italics</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("simple role italics unconstrained", func() {
			source := "it[uncle]__al__ic"
			expected := `<div class="paragraph">
<p>it<em class="uncle">al</em>ic</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("simple role bold", func() {
			source := "[myrole]*bold*"
			expected := `<div class="paragraph">
<p><strong class="myrole">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("simple role bold unconstrained", func() {
			source := "it[uncle]**al**ic"
			expected := `<div class="paragraph">
<p>it<strong class="uncle">al</strong>ic</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("simple role mono", func() {
			source := "[.myrole]`true`"
			expected := `<div class="paragraph">
<p><code class="myrole">true</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("short-hand role with special characters", func() {
			source := `[."a <role>"]**bold**`
			// wrapping quotes are not preserved
			expected := `<div class="paragraph">
<p><strong class="a &lt;role&gt;">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("simple role mono unconstrained", func() {
			source := "int[.uncle]``eg``rate"
			expected := `<div class="paragraph">
<p>int<code class="uncle">eg</code>rate</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("role with comma truncates", func() {
			source := "[.myrole,and=nothing_else]_italics_"
			expected := `<div class="paragraph">
<p><em class="myrole">italics</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("short-hand ID only", func() {
			source := "[#here]*bold*"
			expected := `<div class="paragraph">
<p><strong id="here">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("short-hand role only", func() {
			source := "[.bob]**bold**"
			expected := `<div class="paragraph">
<p><strong class="bob">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked role (span) only", func() {
			source := "[.bob]##bold##"
			expected := `<div class="paragraph">
<p><span class="bob">bold</span></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked role id only", func() {
			source := "[#link]##content##"
			expected := `<div class="paragraph">
<p><mark id="link">content</mark></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty role", func() {
			source := "[]**bold**"
			expected := `<div class="paragraph">
<p><strong>bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("short-hand multiple roles and id", func() {
			source := "[.r1#anchor.r2.r3]**bold**[#here.second.class]_text_"
			expected := `<div class="paragraph">
<p><strong id="anchor" class="r1 r2 r3">bold</strong><em id="here" class="second class">text</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		// This is a departure from asciidoctor, as we support quoting the role in first position.
		// (Asciidoctor passes it, but does not sanitize it, leading to invalid HTML.)
		It("quoted role", func() {
			source := "[.\"something <wicked>\"]**bold**"
			expected := `<div class="paragraph">
<p><strong class="something &lt;wicked&gt;">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("quoted short-hand role", func() {
			source := "[.'something \"wicked\"']**bold**"
			expected := `<div class="paragraph">
<p><strong class="something "wicked"">bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		// This demonstrates that we cannot inject malicious data in these attributes.
		// Note that this is a divergence from asciidoctor, which lets you put whatever you want here.
		It("bad syntax", func() {
			source := "[.<something \"wicked>]**bold**"
			expected := `<div class="paragraph">
<p>[.&lt;something "wicked&gt;]<strong>bold</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("nested content", func() {

		It("nested bold quote within bold quote with same punctuation", func() {

			source := "*some *nested bold* content*."
			expected := `<div class="paragraph">
<p><strong>some *nested bold</strong> content*.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content within bold quote in sentence", func() {
			source := "some *bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some <strong>bold and <em>italic content</em></strong> together.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("marked content within bold quote in sentence", func() {
			source := "some *bold and #marked content#* together."
			expected := `<div class="paragraph">
<p>some <strong>bold and <mark>marked content</mark></strong> together.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("span content within italic quote in sentence", func() {
			source := "some *bold and [.strikeout]#span content#* together."
			expected := `<div class="paragraph">
<p>some <strong>bold and <span class="strikeout">span content</span></strong> together.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("invalid  content", func() {

		It("italic content within invalid bold quote in sentence", func() {
			source := "some *bold and _italic content_ * together."
			expected := `<div class="paragraph">
<p>some *bold and <em>italic content</em> * together.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid italic content within bold quote in sentence", func() {

			source := "some *bold and _italic content _ together*."
			expected := `<div class="paragraph">
<p>some <strong>bold and _italic content _ together</strong>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("prevented substitution", func() {

		It("escaped bold content in sentence", func() {
			source := "some \\*bold content*."
			expected := `<div class="paragraph">
<p>some *bold content*.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content within escaped bold quote in sentence", func() {
			source := "some \\*bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some *bold and <em>italic content</em>* together.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("mixed content", func() {

		It("unbalanced bold in monospace - case 1", func() {
			source := "`*a`"
			expected := `<div class="paragraph">
<p><code>*a</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unbalanced bold in monospace - case 2", func() {
			source := "`a*b`"
			expected := `<div class="paragraph">
<p><code>a*b</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic in monospace", func() {
			source := "`_a_`"
			expected := `<div class="paragraph">
<p><code><em>a</em></code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unbalanced italic in monospace", func() {
			source := "`a_b`"
			expected := `<div class="paragraph">
<p><code>a_b</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unparsed bold in monospace", func() {
			source := "`a*b*`"
			expected := `<div class="paragraph">
<p><code>a*b*</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("parsed subscript in monospace", func() {
			source := "`a~b~`"
			expected := `<div class="paragraph">
<p><code>a<sub>b</sub></code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline in monospace - case 1", func() {
			source := "`a\nb`"
			expected := `<div class="paragraph">
<p><code>a
b</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline in monospace - case 2", func() {
			source := "`a\n*b*`"
			expected := `<div class="paragraph">
<p><code>a
<strong>b</strong></code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in bold", func() {
			source := "*a link:/[b]*"
			expected := `<div class="paragraph">
<p><strong>a <a href="/">b</a></strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in bold", func() {
			source := "*a image:foo.png[]*"
			expected := `<div class="paragraph">
<p><strong>a <span class="image"><img src="foo.png" alt="foo"></span></strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in bold", func() {
			source := "*a +image:foo.png[]+*"
			expected := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in bold", func() {
			source := "*a +++image:foo.png[]+++*"
			expected := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in italic", func() {
			source := "_a link:/[b]_"
			expected := `<div class="paragraph">
<p><em>a <a href="/">b</a></em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in italic", func() {
			source := "_a image:foo.png[]_"
			expected := `<div class="paragraph">
<p><em>a <span class="image"><img src="foo.png" alt="foo"></span></em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in italic", func() {
			source := "_a +image:foo.png[]+_"
			expected := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in italic", func() {
			source := "_a +++image:foo.png[]+++_"
			expected := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in monospace", func() {
			source := "`a link:/[b]`"
			expected := `<div class="paragraph">
<p><code>a <a href="/">b</a></code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in monospace", func() {
			source := "`a image:foo.png[]`"
			expected := `<div class="paragraph">
<p><code>a <span class="image"><img src="foo.png" alt="foo"></span></code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in monospace", func() {
			source := "`a +image:foo.png[]+`"
			expected := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in monospace", func() {
			source := "`a +++image:foo.png[]+++`"
			expected := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in single bold", func() {
			source := "this *mother's mothers' mothers`'*\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <strong>mother&#8217;s mothers' mothers&#8217;</strong></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in double bold", func() {
			source := "this **mother's mothers' mothers`'**\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <strong>mother&#8217;s mothers' mothers&#8217;</strong></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in single italic", func() {
			source := "this _mother's mothers' mothers`'_\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <em>mother&#8217;s mothers' mothers&#8217;</em></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in double italic", func() {
			source := "this __mother's mothers' mothers`'__\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <em>mother&#8217;s mothers' mothers&#8217;</em></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in single mono", func() {
			source := "this `mother's mothers' day`\n" // no typographic quotes here
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <code>mother&#8217;s mothers' day</code></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in double mono", func() {
			source := "this ``mother's mothers' mothers`' day``\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <code>mother&#8217;s mothers' mothers&#8217; day</code></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in single marked", func() {
			source := "this #mother's mothers' mothers`'#\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <mark>mother&#8217;s mothers' mothers&#8217;</mark></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("apostrophes in double marked", func() {
			source := "this ##mother's mothers' mothers`'##\n"
			expected := "<div class=\"paragraph\">\n" +
				"<p>this <mark>mother&#8217;s mothers' mothers&#8217;</mark></p>\n" +
				"</div>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

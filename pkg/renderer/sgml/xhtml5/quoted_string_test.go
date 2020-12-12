package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quoted strings", func() {

	Context("quoted strings", func() {

		It("bold content alone", func() {
			source := "*bold content*"
			expected := `<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("simple single quoted string", func() {
			source := "'`curly was single`'"
			expected := `<div class="paragraph">
<p>&#8216;curly was single&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("spaces with single quoted string", func() {
			source := "'` curly was single `' or so they say"
			expected := "<div class=\"paragraph\">\n" +
				"<p>'` curly was single &#8217; or so they say</p>\n" +
				"</div>\n"
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("bold in single quoted string", func() {
			source := "'`curly *was* single`'"
			expected := `<div class="paragraph">
<p>&#8216;curly <strong>was</strong> single&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("italics in single quoted string", func() {
			source := "'`curly _was_ single`'"
			expected := `<div class="paragraph">
<p>&#8216;curly <em>was</em> single&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("span in single quoted string", func() {
			source := "'`curly [.strikeout]#was#_is_ single`'"
			expected := `<div class="paragraph">
<p>&#8216;curly <span class="strikeout">was</span><em>is</em> single&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("curly in monospace  string", func() {
			source := "'`curly `is` single`'"
			expected := `<div class="paragraph">
<p>&#8216;curly <code>is</code> single&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("curly as monospace string", func() {
			source := "'``curly``'"
			expected := `<div class="paragraph">
<p>&#8216;<code>curly</code>&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly with nested double curly", func() {
			source := "'`single\"`double`\"`'"
			expected := `<div class="paragraph">
<p>&#8216;single&#8220;double&#8221;&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly in monospace string", func() {
			source := "`'`curly`'`"
			expected := `<div class="paragraph">
<p><code>&#8216;curly&#8217;</code></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("curly in italics", func() {
			source := "_'`curly`'_"
			expected := `<div class="paragraph">
<p><em>&#8216;curly&#8217;</em></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("curly in bold", func() {
			source := "*'`curly`'*"
			expected := `<div class="paragraph">
<p><strong>&#8216;curly&#8217;</strong></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly in title", func() {
			source := "== a '`curly`' episode"
			expected := `<div class="sect1">
<h2 id="_a_episode">a &#8216;curly&#8217; episode</h2>
<div class="sectionbody">
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly in list element", func() {
			source := "* a '`curly`' episode"
			expected := `<div class="ulist">
<ul>
<li>
<p>a &#8216;curly&#8217; episode</p>
</li>
</ul>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly in labeled list", func() {
			source := "'`term`':: something '`quoted`'"
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">&#8216;term&#8217;</dt>
<dd>
<p>something &#8216;quoted&#8217;</p>
</dd>
</dl>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("curly in link", func() {
			source := "https://www.example.com/a['`example`']"
			expected := `<div class="paragraph">
<p><a href="https://www.example.com/a">&#8216;example&#8217;</a></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("curly in quoted link", func() {
			source := "https://www.example.com/a[\"an '`example`'\"]"
			expected := `<div class="paragraph">
<p><a href="https://www.example.com/a">an &#8216;example&#8217;</a></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("image in curly", func() {
			source := "'`a image:foo.png[]`'"
			expected := `<div class="paragraph">
<p>&#8216;a <span class="image"><img src="foo.png" alt="foo"/></span>&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("icon in curly", func() {
			source := ":icons: font\n\n'`a icon:note[]`'"
			expected := `<div class="paragraph">
<p>&#8216;a <span class="icon"><i class="fa fa-note"></i></span>&#8217;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("simple single quoted string", func() {
			source := "\"`curly was single`\""
			expected := `<div class="paragraph">
<p>&#8220;curly was single&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})

		It("spaces with double quoted string", func() {
			source := "\"` curly was single `\""
			expected := "<div class=\"paragraph\">\n" +
				"<p>\"` curly was single `\"</p>\n" +
				"</div>\n"
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("bold in double quoted string", func() {
			source := "\"`curly *was* single`\""
			expected := `<div class="paragraph">
<p>&#8220;curly <strong>was</strong> single&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("italics in double quoted string", func() {
			source := "\"`curly _was_ single`\""
			expected := `<div class="paragraph">
<p>&#8220;curly <em>was</em> single&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("span in double quoted string", func() {
			source := "\"`curly [.strikeout]#was#_is_ single`\""
			expected := `<div class="paragraph">
<p>&#8220;curly <span class="strikeout">was</span><em>is</em> single&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("double curly in monospace string", func() {
			source := "\"`curly `is` single`\""
			expected := `<div class="paragraph">
<p>&#8220;curly <code>is</code> single&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("double curly as monospace string", func() {
			source := "\"``curly``\""
			expected := `<div class="paragraph">
<p>&#8220;<code>curly</code>&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("double curly with nested single curly", func() {
			source := "\"`double'`single`'`\""
			expected := `<div class="paragraph">
<p>&#8220;double&#8216;single&#8217;&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("double curly in monospace string", func() {
			source := "`\"`curly`\"`"
			expected := `<div class="paragraph">
<p><code>&#8220;curly&#8221;</code></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("double curly in italics", func() {
			source := "_\"`curly`\"_"
			expected := `<div class="paragraph">
<p><em>&#8220;curly&#8221;</em></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("double curly in bold", func() {
			source := "*\"`curly`\"*"
			expected := `<div class="paragraph">
<p><strong>&#8220;curly&#8221;</strong></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})

		It("double curly in title", func() {
			source := "== a \"`curly`\" episode"
			expected := `<div class="sect1">
<h2 id="_a_episode">a &#8220;curly&#8221; episode</h2>
<div class="sectionbody">
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("double in list element", func() {
			source := "* a \"`curly`\" episode"
			expected := `<div class="ulist">
<ul>
<li>
<p>a &#8220;curly&#8221; episode</p>
</li>
</ul>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
		It("double curly in labeled list", func() {
			source := "\"`term`\":: something \"`quoted`\""
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">&#8220;term&#8221;</dt>
<dd>
<p>something &#8220;quoted&#8221;</p>
</dd>
</dl>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		// In a link, the quotes are ambiguous, and we default to assuming they are for enclosing
		// the link text.  Nest them explicitly if this is needed.
		It("double curly in link (becomes mono)", func() {
			source := "https://www.example.com/a[\"`example`\"]"
			expected := `<div class="paragraph">
<p><a href="https://www.example.com/a">&#8220;example&#8221;</a></p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		// This is the unambiguous form.
		It("curly in quoted link", func() {
			source := "https://www.example.com/a[\"\"`example`\"\"]"
			expected := `<div class="paragraph">
<p><a href="https://www.example.com/a">&#8220;example&#8221;</a></p>
</div>
`

			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("image in double curly", func() {
			source := "\"`a image:foo.png[]`\""
			expected := `<div class="paragraph">
<p>&#8220;a <span class="image"><img src="foo.png" alt="foo"/></span>&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
		It("icon in double curly", func() {
			source := ":icons: font\n\n\"`a icon:note[]`\""
			expected := `<div class="paragraph">
<p>&#8220;a <span class="icon"><i class="fa fa-note"></i></span>&#8221;</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
	})
})

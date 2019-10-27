package html5_test

import (
	"fmt"

	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("document with attributes", func() {

	Context("plaintext substitutions", func() {

		It("some attributes then a paragraph", func() {
			source := `:toc:
:date: 2017-01-01
:author: Xavier
a paragraph`
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("a paragraph then some attributes", func() {
			source := `a paragraph

:toc:
:date: 2017-01-01
:author: Xavier`
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("a paragraph with substitution", func() {
			source := `:author: Xavier

a paragraph written by {author}`
			expected := `<div class="paragraph">
<p>a paragraph written by Xavier</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("paragraphs with definitions, substitutions and resets", func() {
			source := `author is {author}.
		
:author: me
author is now {author}.

:author: you
author is now {author}.

:author!:
author is now {author}.`
			expected := `<div class="paragraph">
<p>author is {author}.</p>
</div>
<div class="paragraph">
<p>author is now me.</p>
</div>
<div class="paragraph">
<p>author is now you.</p>
</div>
<div class="paragraph">
<p>author is now {author}.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("front-matter then paragraph with substitutions", func() {
			source := `---
author: Xavier
---
		
author is {author}.`
			expected := `<div class="paragraph">
<p>author is Xavier.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("substitutions to elements", func() {

		It("replace to inline link in paragraph", func() {
			source := `:quick-uri: https://foo.com/bar
{quick-uri}[foo]`
			expected := `<div class="paragraph">
<p><a href="https://foo.com/bar">foo</a></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("predefined attributes", func() {

		DescribeTable("predefined attributes in a paragraph",
			func(code, rendered string) {
				source := fmt.Sprintf(`the {%s} symbol`, code)
				expected := fmt.Sprintf(`<div class="paragraph">
<p>the %s symbol</p>
</div>`, rendered)
				Expect(source).To(RenderHTML5Body(expected))
			},
			Entry("sp symbol", "sp", " "),
			Entry("blank symbol", "blank", ""),
			Entry("empty symbol", "empty", ""),
			Entry("nbsp symbol", "nbsp", "&#160;"),
			Entry("zwsp symbol", "zwsp", "&#8203;"),
			Entry("wj symbol", "wj", "&#8288;"),
			Entry("apos symbol", "apos", "&#39;"),
			Entry("quot symbol", "quot", "&#34;"),
			Entry("lsquo symbol", "lsquo", "&#8216;"),
			Entry("rsquo symbol", "rsquo", "&#8217;"),
			Entry("ldquo symbol", "ldquo", "&#8220;"),
			Entry("rdquo symbol", "rdquo", "&#8221;"),
			Entry("deg symbol", "deg", "&#176;"),
			Entry("plus symbol", "plus", "&#43;"),
			Entry("brvbar symbol", "brvbar", "&#166;"),
			Entry("vbar symbol", "vbar", "|"),
			Entry("amp symbol", "amp", "&amp;"),
			Entry("lt symbol", "lt", "&lt;"),
			Entry("gt symbol", "gt", "&gt;"),
			Entry("startsb symbol", "startsb", "["),
			Entry("endsb symbol", "endsb", "]"),
			Entry("caret symbol", "caret", "^"),
			Entry("asterisk symbol", "asterisk", "*"),
			Entry("tilde symbol", "tilde", "~"),
			Entry("backslash symbol", "backslash", `\`),
			Entry("backtick symbol", "backtick", "`"),
			Entry("two-colons symbol", "two-colons", "::"),
			Entry("two-semicolons symbol", "two-semicolons", ";"),
			Entry("cpp symbol", "cpp", "C++"),
		)

		It("overriding predefined attribute", func() {
			source := `:blank: foo
			
a {blank} here.`
			expected := `<div class="paragraph">
<p>a foo here.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

})

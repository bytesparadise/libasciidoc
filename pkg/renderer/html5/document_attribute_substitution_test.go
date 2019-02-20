package html5_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
)

var _ = Describe("document with attributes", func() {

	Context("plaintext substitutions", func() {

		It("some attributes then a paragraph", func() {
			actualContent := `:toc:
:date: 2017-01-01
:author: Xavier
a paragraph`
			expectedResult := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a paragraph then some attributes", func() {
			actualContent := `a paragraph

:toc:
:date: 2017-01-01
:author: Xavier`
			expectedResult := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a paragraph with substitution", func() {
			actualContent := `:author: Xavier

a paragraph written by {author}`
			expectedResult := `<div class="paragraph">
<p>a paragraph written by Xavier</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraphs with definitions, substitutions and resets", func() {
			actualContent := `author is {author}.
		
:author: me
author is now {author}.

:author: you
author is now {author}.

:author!:
author is now {author}.`
			expectedResult := `<div class="paragraph">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("front-matter then paragraph with substitutions", func() {
			actualContent := `---
author: Xavier
---
		
author is {author}.`
			expectedResult := `<div class="paragraph">
<p>author is Xavier.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("substitutions to elements", func() {

		It("replace to inline link in paragraph", func() {
			actualContent := `:quick-uri: http://foo.com/bar
{quick-uri}[foo]`
			expectedResult := `<div class="paragraph">
<p><a href="http://foo.com/bar">foo</a></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("predefined attributes", func() {

		DescribeTable("predefined attributes in a paragraph",
			func(code, rendered string) {
				actualContent := fmt.Sprintf(`the {%s} symbol`, code)
				expectedResult := fmt.Sprintf(`<div class="paragraph">
<p>the %s symbol</p>
</div>`, rendered)
				verify(GinkgoT(), expectedResult, actualContent)
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
			actualContent := `:blank: foo
			
a {blank} here.`
			expectedResult := `<div class="paragraph">
<p>a foo here.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})

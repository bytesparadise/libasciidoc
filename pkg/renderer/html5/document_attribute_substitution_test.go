package html5_test

import (
	. "github.com/onsi/ginkgo"
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

	Context("predefined elements", func() {

		It("single space", func() {
			actualContent := `a {sp} here.`
			expectedResult := `<div class="paragraph">
<p>a   here.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("blank", func() {
			actualContent := `a {blank} here.`
			expectedResult := `<div class="paragraph">
<p>a  here.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

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

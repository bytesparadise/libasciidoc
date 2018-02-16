package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Document With Attributes", func() {
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

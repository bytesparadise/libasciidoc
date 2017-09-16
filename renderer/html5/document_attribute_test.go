package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering With Attributes", func() {
	It("some attributes then a paragraph", func() {
		content := `:toc:
:date: 2017-01-01
:author: Xavier
a paragraph`

		expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
})

var _ = Describe("Rendering With Attributes", func() {
	It("a paragraph then some attributes", func() {
		content := `a paragraph

:toc:
:date: 2017-01-01
:author: Xavier`

		expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})

	It("a paragraph with substitution", func() {
		content := `:author: Xavier

a paragraph written by {author}`

		expected := `<div class="paragraph">
<p>a paragraph written by Xavier</p>
</div>`
		verify(GinkgoT(), expected, content)
	})

	It("paragraphs with definitions, substitutions and resets", func() {
		content := `author is {author}.
		
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
		verify(GinkgoT(), expected, content)
	})

})

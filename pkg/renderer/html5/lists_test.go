package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("lists of items", func() {

	Context("distinct list blocks", func() {

		It("same list without attributes", func() {
			actualContent := `[lowerroman]
. Five
.. a
. Six`
			expectedResult := `<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>Five</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>a</p>
</li>
</ol>
</div>
</li>
<li>
<p>Six</p>
</li>
</ol>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("same list with attribute on middle item", func() {
			actualContent := `[lowerroman]
. Five
[loweralpha]
.. a
. Six`
			expectedResult := `<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>Five</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>a</p>
</li>
</ol>
</div>
</li>
<li>
<p>Six</p>
</li>
</ol>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("distinct lists - case 1", func() {
			actualContent := `[lowerroman]
. Five

[loweralpha]
.. a
. Six`
			expectedResult := `<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>Five</p>
</li>
</ol>
</div>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>a</p>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>Six</p>
</li>
</ol>
</div>
</li>
</ol>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("distinct lists - case 2", func() {

			actualContent := `.Checklist
- [*] checked
-     normal list item

.Ordered, basic
. Step 1
. Step 2`
			expectedResult := `<div class="ulist checklist">
<div class="title">Checklist</div>
<ul class="checklist">
<li>
<p>&#10003; checked</p>
</li>
<li>
<p>normal list item</p>
</li>
</ul>
</div>
<div class="olist arabic">
<div class="title">Ordered, basic</div>
<ol class="arabic">
<li>
<p>Step 1</p>
</li>
<li>
<p>Step 2</p>
</li>
</ol>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})

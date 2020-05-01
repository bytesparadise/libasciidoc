package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("lists of items", func() {

	Context("distinct list blocks", func() {

		It("same list without attributes", func() {
			source := `[lowerroman]
. Five
.. a
. Six`
			expected := `<div class="olist lowerroman">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("same list with attribute on middle item", func() {
			source := `[lowerroman]
. Five
[loweralpha]
.. a
. Six`
			expected := `<div class="olist lowerroman">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("distinct lists separated by blankline and item attribute - case 1", func() {
			source := `[lowerroman]
. Five

[loweralpha]
.. a
. Six`
			expected := `<div class="olist lowerroman">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("distinct lists separated by blankline and item attribute - case 2", func() {

			source := `.Checklist
- [*] checked
-     normal list item

.Ordered, basic
. Step 1
. Step 2`
			expected := `<div class="ulist checklist">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	It("same list with single comment line inside", func() {
		source := `. a
// -
. b`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>a</p>
</li>
<li>
<p>b</p>
</li>
</ol>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("same list with multiple comment lines inside", func() {
		source := `. a
// -
// -
// -
. b`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>a</p>
</li>
<li>
<p>b</p>
</li>
</ol>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("distinct lists separated by single comment line", func() {
		source := `. a

// -
. b`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>a</p>
</li>
</ol>
</div>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>b</p>
</li>
</ol>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("distinct lists separated by multiple comment lines", func() {
		source := `. a

// -
// -
// -
. b`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>a</p>
</li>
</ol>
</div>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>b</p>
</li>
</ol>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

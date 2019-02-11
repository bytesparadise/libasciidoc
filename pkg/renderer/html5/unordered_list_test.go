package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("unordered lists", func() {

	It("simple unordered list with no title", func() {
		actualContent := `* item 1
* item 2
* item 3`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
<li>
<p>item 3</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("simple unordered list with no title then a paragraph", func() {
		actualContent := `* item 1
* item 2
* item 3

and a standalone paragraph`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
<li>
<p>item 3</p>
</li>
</ul>
</div>
<div class="paragraph">
<p>and a standalone paragraph</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("simple unordered list with title and role", func() {
		actualContent := `.mytitle
[#foo]
[.myrole]
* item 1
* item 2`
		expectedResult := `<div id="foo" class="ulist myrole">
<div class="title">mytitle</div>
<ul>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("simple unordered list with continuation", func() {
		actualContent := `* item 1
+
foo

* item 2`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>item 1</p>
<div class="paragraph">
<p>foo</p>
</div>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("nested unordered lists without a title", func() {
		actualContent := `* item 1
** item 1.1
** item 1.2
* item 2`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>item 1</p>
<div class="ulist">
<ul>
<li>
<p>item 1.1</p>
</li>
<li>
<p>item 1.2</p>
</li>
</ul>
</div>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("nested unordered lists with a title", func() {
		actualContent := `[#listID]
* item 1
** item 1.1
** item 1.2
* item 2`
		expectedResult := `<div id="listID" class="ulist">
<ul>
<li>
<p>item 1</p>
<div class="ulist">
<ul>
<li>
<p>item 1.1</p>
</li>
<li>
<p>item 1.2</p>
</li>
</ul>
</div>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("unordered list with item continuation", func() {
		actualContent := `* foo
+
----
a delimited block
----
+
----
another delimited block
----
* bar
`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>foo</p>
<div class="listingblock">
<div class="content">
<pre>a delimited block</pre>
</div>
</div>
<div class="listingblock">
<div class="content">
<pre>another delimited block</pre>
</div>
</div>
</li>
<li>
<p>bar</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("unordered list without item continuation", func() {
		actualContent := `* foo
----
a delimited block
----
* bar
----
another delimited block
----`
		expectedResult := `<div class="ulist">
<ul>
<li>
<p>foo</p>
</li>
</ul>
</div>
<div class="listingblock">
<div class="content">
<pre>a delimited block</pre>
</div>
</div>
<div class="ulist">
<ul>
<li>
<p>bar</p>
</li>
</ul>
</div>
<div class="listingblock">
<div class="content">
<pre>another delimited block</pre>
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})
})

var _ = Describe("checklists", func() {

	It("checklist with title and dashes", func() {
		actualContent := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
		expectedResult := `<div class="ulist checklist">
<div class="title">Checklist</div>
<ul class="checklist">
<li>
<p>&#10003; checked</p>
</li>
<li>
<p>&#10003; also checked</p>
</li>
<li>
<p>&#10063; not checked</p>
</li>
<li>
<p>normal list item</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("parent checklist with title and nested checklist", func() {
		actualContent := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
		expectedResult := `<div class="ulist checklist">
<div class="title">Checklist</div>
<ul class="checklist">
<li>
<p>&#10063; parent not checked</p>
<div class="ulist checklist">
<ul class="checklist">
<li>
<p>&#10003; checked</p>
</li>
<li>
<p>&#10003; also checked</p>
</li>
<li>
<p>&#10063; not checked</p>
</li>
</ul>
</div>
</li>
<li>
<p>normal list item</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("parent checklist with role and nested normal list", func() {
		actualContent := `[.Checklist]
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
		expectedResult := `<div class="ulist checklist Checklist">
<ul class="checklist">
<li>
<p>&#10063; parent not checked</p>
<div class="ulist">
<ul>
<li>
<p>a normal list item</p>
</li>
<li>
<p>another normal list item</p>
</li>
</ul>
</div>
</li>
<li>
<p>normal list item</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	Context("attach to unordered list item ancestor", func() {

		It("attach to grandparent unordered list item", func() {
			actualContent := `* grandparent list item
** parent list item
*** child list item


+
paragraph attached to grandparent list item`
			expectedResult := `<div class="ulist">
<ul>
<li>
<p>grandparent list item</p>
<div class="ulist">
<ul>
<li>
<p>parent list item</p>
<div class="ulist">
<ul>
<li>
<p>child list item</p>
</li>
</ul>
</div>
</li>
</ul>
</div>
<div class="paragraph">
<p>paragraph attached to grandparent list item</p>
</div>
</li>
</ul>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("attach to parent unordered list item", func() {
			actualContent := `* grandparent list item
** parent list item
*** child list item

+
paragraph attached to parent list item`
			expectedResult := `<div class="ulist">
<ul>
<li>
<p>grandparent list item</p>
<div class="ulist">
<ul>
<li>
<p>parent list item</p>
<div class="ulist">
<ul>
<li>
<p>child list item</p>
</li>
</ul>
</div>
<div class="paragraph">
<p>paragraph attached to parent list item</p>
</div>
</li>
</ul>
</div>
</li>
</ul>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("attach to child unordered list item", func() {
			actualContent := `* grandparent list item
** parent list item
*** child list item
+
paragraph attached to child list item`
			expectedResult := `<div class="ulist">
<ul>
<li>
<p>grandparent list item</p>
<div class="ulist">
<ul>
<li>
<p>parent list item</p>
<div class="ulist">
<ul>
<li>
<p>child list item</p>
<div class="paragraph">
<p>paragraph attached to child list item</p>
</div>
</li>
</ul>
</div>
</li>
</ul>
</div>
</li>
</ul>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})

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

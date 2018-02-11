package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("unordered lists", func() {
	It("simple unordered list with no title", func() {
		actualContent := `* item 1
* item 2
* item 3`
		expected := `<div class="ulist">
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
		verify(GinkgoT(), expected, actualContent)
	})
	It("simple unordered list with a title", func() {
		actualContent := `[#foo]
	* item 1
	* item 2`
		expected := `<div id="foo" class="ulist">
<ul>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})
	It("nested unordered lists without a title", func() {
		actualContent := `* item 1
** item 1.1
** item 1.2
* item 2`
		expected := `<div class="ulist">
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
		verify(GinkgoT(), expected, actualContent)
	})
	It("nested unordered lists with a title", func() {
		actualContent := `[#listID]
* item 1
** item 1.1
** item 1.2
* item 2`
		expected := `<div id="listID" class="ulist">
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
		verify(GinkgoT(), expected, actualContent)
	})
})

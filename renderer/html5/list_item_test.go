package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering List of Items", func() {
	It("simple list", func() {
		content := `* item 1
* item 2`
		expected := `<div class="ulist">
<ul>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
</ul>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("simple list with a title", func() {
		content := `[#foo]
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
		verify(GinkgoT(), expected, content)
	})
	It("nested lists", func() {
		content := `* item 1
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
		verify(GinkgoT(), expected, content)
	})
	It("nested lists with a title", func() {
		content := `[#foo]
* item 1
** item 1.1
** item 1.2
* item 2`
		expected := `<div id="foo" class="ulist">
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
		verify(GinkgoT(), expected, content)
	})
})

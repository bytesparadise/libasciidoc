package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("labeled lists of items", func() {

	Context("simple items", func() {

		It("simple labeled list with a title and a default layout", func() {
			actualContent := `[#listID]
item 1:: description 1.
item 2:: description 2
on 2 lines.
item 3:: description 3
on 2 lines, too.`
			expected := `<div id="listID" class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
<dd>
<p>description 1.</p>
</dd>
<dt class="hdlist1">item 2</dt>
<dd>
<p>description 2
on 2 lines.</p>
</dd>
<dt class="hdlist1">item 3</dt>
<dd>
<p>description 3
on 2 lines, too.</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with an empty entry", func() {
			actualContent := `item 1::
item 2:: description 2.`
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
<dt class="hdlist1">item 2</dt>
<dd>
<p>description 2.</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with an image", func() {
			actualContent := `item 1:: image:foo.png[]
item 2:: description 2.`
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
<dd>
<p><span class="image"><img src="foo.png" alt="foo"></span></p>
</dd>
<dt class="hdlist1">item 2</dt>
<dd>
<p>description 2.</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with script injection", func() {
			actualContent := `item 1:: <script>alert("foo!")</script>`
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
<dd>
<p>&lt;script&gt;alert(&#34;foo!&#34;)&lt;/script&gt;</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with fenced block", func() {
			actualContent := "item 1::\n" +
				"```\n" +
				"a fenced block\n" +
				"```\n" +
				"item 2:: something simple"
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
</dl>
</div>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code>a fenced block</code></pre>
</div>
</div>
<div class="dlist">
<dl>
<dt class="hdlist1">item 2</dt>
<dd>
<p>something simple</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with nested lists", func() {
			actualContent := `item 1:: 
* foo
* bar
** baz
item 2:: something simple`
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1">item 1</dt>
<dd>
<div class="ulist">
<ul>
<li>
<p>foo</p>
</li>
<li>
<p>bar</p>
<div class="ulist">
<ul>
<li>
<p>baz</p>
</li>
</ul>
</div>
</li>
</ul>
</div>
</dd>
<dt class="hdlist1">item 2</dt>
<dd>
<p>something simple</p>
</dd>
</dl>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

	})

	Context("horizontal layout", func() {

		It("simple labeled list with horizontal layout", func() {
			actualContent := `[horizontal]
item 1::
item 2:: description 2 on 1 line.
item 3:: description 3
on 2 lines, too.`
			expected := `<div class="hdlist">
<table>
<tr>
<td class="hdlist1">
item 1
<br>
item 2
</td>
<td class="hdlist2">
<p>description 2 on 1 line.</p>
</td>
</tr>
<tr>
<td class="hdlist1">
item 3
</td>
<td class="hdlist2">
<p>description 3
on 2 lines, too.</p>
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("labeled list with nested lists and horizontal layout", func() {
			actualContent := `[horizontal]
item 1:: 
* foo
* bar
** baz
item 2:: something simple`
			expected := `<div class="hdlist">
<table>
<tr>
<td class="hdlist1">
item 1
</td>
<td class="hdlist2">
<div class="ulist">
<ul>
<li>
<p>foo</p>
</li>
<li>
<p>bar</p>
<div class="ulist">
<ul>
<li>
<p>baz</p>
</li>
</ul>
</div>
</li>
</ul>
</div>
</td>
</tr>
<tr>
<td class="hdlist1">
item 2
</td>
<td class="hdlist2">
<p>something simple</p>
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

	})
})

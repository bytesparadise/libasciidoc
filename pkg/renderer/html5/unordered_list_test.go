package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("unordered lists", func() {

	It("simple unordered list with no title", func() {
		source := `* item 1
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("simple unordered list with no title then a paragraph", func() {
		source := `* item 1
* item 2
* item 3

and a standalone paragraph`
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
</div>
<div class="paragraph">
<p>and a standalone paragraph</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("simple unordered list with id, title and role", func() {
		source := `.mytitle
[#foo]
[.myrole]
* item 1
* item 2`
		expected := `<div id="foo" class="ulist myrole">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("simple unordered list with id, title and role", func() {
		source := `.mytitle
[#foo]
[.myrole]
* item 1
* item 2`
		expected := `<div id="foo" class="ulist myrole">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("simple unordered list with continuation", func() {
		source := `* item 1
+
foo

* item 2`
		expected := `<div class="ulist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("nested unordered lists without a title", func() {
		source := `* item 1
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("nested unordered lists with a title", func() {
		source := `[#listID]
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("unordered list with item continuation", func() {
		source := `* foo
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
		expected := `<div class="ulist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("unordered list without item continuation", func() {
		source := `* foo
----
a delimited block
----
* bar
----
another delimited block
----`
		expected := `<div class="ulist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

var _ = Describe("checklists", func() {

	It("checklist with title and dashes", func() {
		source := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
		expected := `<div class="ulist checklist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("parent checklist with title and nested checklist", func() {
		source := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
		expected := `<div class="ulist checklist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("parent checklist with role and nested normal list", func() {
		source := `[.Checklist]
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
		expected := `<div class="ulist checklist Checklist">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	Context("attach to unordered list item ancestor", func() {

		It("attach to grandparent unordered list item", func() {
			source := `* grandparent list item
** parent list item
*** child list item


+
paragraph attached to grandparent list item`
			expected := `<div class="ulist">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("attach to parent unordered list item", func() {
			source := `* grandparent list item
** parent list item
*** child list item

+
paragraph attached to parent list item`
			expected := `<div class="ulist">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("attach to child unordered list item", func() {
			source := `* grandparent list item
** parent list item
*** child list item
+
paragraph attached to child list item`
			expected := `<div class="ulist">
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

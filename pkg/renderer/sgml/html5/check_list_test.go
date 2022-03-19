package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("checklists", func() {

	It("with title and dashes", func() {
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
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with style id, title and role", func() {
		// style is overridden to checklist on ul, but div keeps it (asciidoctor compat)
		source := `.mytitle
[#foo]
[disc.myrole]
* [x] item 1
* [x] item 2`
		expected := `<div id="foo" class="ulist checklist disc myrole">
<div class="title">mytitle</div>
<ul class="checklist">
<li>
<p>&#10003; item 1</p>
</li>
<li>
<p>&#10003; item 2</p>
</li>
</ul>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with title and nested checklist", func() {
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
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with role and nested normal list", func() {
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
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with interactive checkboxes", func() {
		source := `[%interactive]
* [*] checked
* [x] also checked
* [ ] not checked
*     normal list item`
		expected := `<div class="ulist checklist">
<ul class="checklist">
<li>
<p><input type="checkbox" data-item-complete="1" checked> checked</p>
</li>
<li>
<p><input type="checkbox" data-item-complete="1" checked> also checked</p>
</li>
<li>
<p><input type="checkbox" data-item-complete="0"> not checked</p>
</li>
<li>
<p>normal list item</p>
</li>
</ul>
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

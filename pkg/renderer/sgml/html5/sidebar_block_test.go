package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("sidebar blocks", func() {

	Context("delimited blocks", func() {

		It("sidebar block with paragraph", func() {
			source := `****
some *verse* content

****`
			expected := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("sidebar block with id, title, paragraph and sourcecode block", func() {
			source := `[#id-for-sidebar]
.title for sidebar
****
some *verse* content

----
foo
bar
----
****`
			expected := `<div id="id-for-sidebar" class="sidebarblock">
<div class="content">
<div class="title">title for sidebar</div>
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
<div class="listingblock">
<div class="content">
<pre>foo
bar</pre>
</div>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

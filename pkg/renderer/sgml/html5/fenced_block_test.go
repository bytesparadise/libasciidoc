package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("delimited blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with surrounding empty lines", func() {
			source := "```\n\nsome source code \n\nhere  \n\n\n\n```"
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("fenced block with empty lines", func() {
			source := "```\n\n\n\n```"
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code></code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("fenced block with id and title", func() {
			source := "[#id-for-fences]\n.fenced block title\n```\nsome source code\n\nhere\n\n\n\n```"
			expected := `<div id="id-for-fences" class="listingblock">
<div class="title">fenced block title</div>
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("fenced block with external link inside", func() {
			source := "```" + "\n" +
				"a http://website.com" + "\n" +
				"and more text on the" + "\n" +
				"next lines" + "\n\n" +
				"```"
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>a http://website.com
and more text on the
next lines</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

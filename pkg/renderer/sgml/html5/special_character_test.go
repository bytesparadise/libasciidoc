package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("special characters", func() {

	It("should parse in paragraph", func() {
		source := "<b>*</b> &apos; &amp;"
		expected := `<div class="paragraph">
<p>&lt;b&gt;*&lt;/b&gt; &amp;apos; &amp;amp;</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("should parse in delimited block", func() {
		source := "```" + "\n" +
			"<b>*</b> &apos; &amp;" + "\n" +
			"```"
		expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>&lt;b&gt;*&lt;/b&gt; &amp;apos; &amp;amp;</code></pre>
</div>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("should parse in paragraph and delimited block", func() {
		source := "<b>*</b> &apos; &amp;" + "\n\n" +
			"```" + "\n" +
			"<b>*</b> &apos; &amp;" + "\n" +
			"```"
		expected := `<div class="paragraph">
<p>&lt;b&gt;*&lt;/b&gt; &amp;apos; &amp;amp;</p>
</div>
<div class="listingblock">
<div class="content">
<pre class="highlight"><code>&lt;b&gt;*&lt;/b&gt; &amp;apos; &amp;amp;</code></pre>
</div>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

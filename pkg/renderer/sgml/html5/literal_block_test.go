package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("literal blocks", func() {

	Context("literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			source := ` some literal content`
			expected := `<div class="literalblock">
<div class="content">
<pre>some literal content</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with single space on first line", func() {
			source := ` some literal content
on 3
lines.`
			expected := `<div class="literalblock">
<div class="content">
<pre> some literal content
on 3
lines.</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with same spaces on each line", func() {
			source := `  some literal content
  on 3
  lines.`
			expected := `<div class="literalblock">
<div class="content">
<pre>some literal content
on 3
lines.</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with single spaces on each line", func() {
			source := ` literal content
   on many lines  
     has some heading spaces preserved.`
			expected := `<div class="literalblock">
<div class="content">
<pre>literal content
  on many lines  
    has some heading spaces preserved.</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("mixing literal block with attributes followed by a paragraph ", func() {
			source := `.title
[#ID]
  some literal content

a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("literal blocks with block delimiter", func() {

		It("literal block with delimited and attributes followed by 1-line paragraph", func() {
			source := `[#ID]
.title
....
 some literal content with space preserved
....
a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre> some literal content with space preserved</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("literal blocks with attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			source := `[literal]   
 literal content
 on many lines 
 has its heading spaces preserved.

a normal paragraph.`
			expected := `<div class="literalblock">
<div class="content">
<pre> literal content
 on many lines 
 has its heading spaces preserved.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from 2-lines paragraph with attribute", func() {
			source := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content
on two lines.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

})

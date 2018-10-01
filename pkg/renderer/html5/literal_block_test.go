package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("literal blocks", func() {

	Context("literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			actualContent := ` some literal content`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre>some literal content</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("literal block from paragraph with single space on first line", func() {
			actualContent := ` some literal content
on 3
lines.`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre> some literal content
on 3
lines.</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("literal block from paragraph with double spaces on each line", func() {
			actualContent := `  some literal content
  on 3
  lines.`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre>some literal content
on 3
lines.</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("literal block from paragraph with single spaces on each line", func() {
			actualContent := ` literal content
   on many lines  
     has some heading spaces preserved.`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre>literal content
  on many lines  
    has some heading spaces preserved.</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("mixing literal block with attributes followed by a paragraph ", func() {
			actualContent := `.title
[#ID]
  some literal content

a normal paragraph.`
			expectedResult := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("literal blocks with block delimiter", func() {

		It("literal block with delimited and attributes followed by 1-line paragraph", func() {
			actualContent := `[#ID]
.title
....
 some literal content with space preserved
....
a normal paragraph.`
			expectedResult := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre> some literal content with space preserved</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("literal blocks with attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			actualContent := `[literal]   
 literal content
 on many lines 
 has some heading spaces preserved.

a normal paragraph.`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre> literal content
 on many lines 
 has some heading spaces preserved.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("literal block from 2-lines paragraph with attribute", func() {
			actualContent := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
			expectedResult := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content
on two lines.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})

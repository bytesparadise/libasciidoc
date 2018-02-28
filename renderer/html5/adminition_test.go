package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("admonitions", func() {

	Context("admonition paragraphs", func() {
		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
this is a note.
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline warning admonition paragraph", func() {
			actualContent := `WARNING: this is a multiline
warning!`
			expectedResult := `<div class="admonitionblock warning">
<table>
<tr>
<td class="icon">
<div class="title">Warning</div>
</td>
<td class="content">
this is a multiline
warning!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("admonition note paragraph with id and title", func() {
			actualContent := `[[foo]]
.bar
NOTE: this is a note.`
			expectedResult := `<div id="foo" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a note.
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition blocks", func() {
		It("caution admonition block", func() {
			actualContent := `[CAUTION] 
this is a caution!`
			expectedResult := `<div class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
this is a caution!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline caution admonition block with title and id", func() {
			actualContent := `[[foo]]
[CAUTION] 
.bar
this is a
*caution*!`
			expectedResult := `<div id="foo" class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a
<strong>caution</strong>!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})

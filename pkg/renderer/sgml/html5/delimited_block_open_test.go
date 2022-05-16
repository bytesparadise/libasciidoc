package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("open blocks", func() {

	Context("without masquerade", func() {

		It("with basic content and attributes", func() {
			source := `[#block-id]
.Block Title
--
basic content
--`
			expected := `<div id="block-id" class="openblock">
<div class="title">Block Title</div>
<div class="content">
<div class="paragraph">
<p>basic content</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with table", func() {
			source := `[#block-id]
.Block Title
--
[cols="2*^"]
|===
a|
[#id]
.A title
image::image.png[]
a|
[#another-id]
.Another title
image::another-image.png[]
|===
--`
			expected := `<div id="block-id" class="openblock">
<div class="title">Block Title</div>
<div class="content">
<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-center valign-top"><div class="content"><div id="id" class="imageblock">
<div class="content">
<img src="image.png" alt="image">
</div>
<div class="title">Figure 1. A title</div>
</div></div></td>
<td class="tableblock halign-center valign-top"><div class="content"><div id="another-id" class="imageblock">
<div class="content">
<img src="another-image.png" alt="another-image">
</div>
<div class="title">Figure 2. Another title</div>
</div></div></td>
</tr>
</tbody>
</table>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})

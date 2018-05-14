package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("cross references", func() {

	Context("section reference", func() {

		It("valid section reference", func() {

			actualContent := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expectedResult := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thetitle">a title</a>!</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("invalid section reference", func() {

			actualContent := `[[thetitle]]
== a title

with some content linked to <<thewrongtitle>>!`
			expectedResult := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thewrongtitle">[thewrongtitle]</a>!</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})

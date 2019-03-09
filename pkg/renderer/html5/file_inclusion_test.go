package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("file inclusions", func() {

	It("include adoc file with leveloffset attribute", func() {
		actualContent := `= Master Document

preamble

include::includes/chapter-a.adoc[leveloffset=+1]`
		expectedResult := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>preamble</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_chapter_a">Chapter A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})
})

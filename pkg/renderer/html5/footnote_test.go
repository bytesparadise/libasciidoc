package html5_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("footnotes", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	It("basic footnote in a paragraph", func() {
		actualContent := `foo footnote:[a note for foo]`
		expectedResult := `<div class="paragraph">
<p>foo <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("rich footnote in a paragraph", func() {
		actualContent := `foo footnote:[some *rich* http://foo.com[content]]`
		expectedResult := `<div class="paragraph">
<p>foo <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. some <strong>rich</strong> <a href="http://foo.com">content</a>
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("footnoteref with valid ref in a paragraph", func() {
		actualContent := `a note here footnoteref:[foo, a note for foo] and an there footnoteref:[foo] too`
		expectedResult := `<div class="paragraph">
<p>a note here <sup class="footnote" id="_footnote_foo">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> and an there <sup class="footnoteref">[<a class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> too</p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("footnoteref with invalid ref in a paragraph", func() {
		actualContent := `a note here footnoteref:[foo, a note for foo] and an unknown there footnoteref:[bar]`
		expectedResult := `<div class="paragraph">
<p>a note here <sup class="footnote" id="_footnote_foo">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> and an unknown there <sup class="footnoteref red" title="Unresolved footnote reference.">[bar]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("footnotes everywhere", func() {

		actualContent := `= title
	
a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`

		// differs from asciidoc in the footnotes at the end of the doc, and the section id (numbering)
		expectedResult := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>a premable with a <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_1">section 1 <sup class="footnote">[<a id="_footnoteref_2" class="footnote" href="#_footnotedef_2" title="View footnote.">2</a>]</sup></h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with another <sup class="footnote">[<a id="_footnoteref_3" class="footnote" href="#_footnotedef_3" title="View footnote.">3</a>]</sup></p>
</div>
</div>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. foo
</div>
<div class="footnote" id="_footnotedef_2">
<a href="#_footnoteref_2">2</a>. bar
</div>
<div class="footnote" id="_footnotedef_3">
<a href="#_footnoteref_3">3</a>. baz
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})
})

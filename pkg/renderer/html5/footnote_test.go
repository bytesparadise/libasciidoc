package html5_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("footnotes", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	It("basic footnote in a paragraph", func() {
		source := `foo footnote:[a note for foo]`
		expected := `<div class="paragraph">
<p>foo <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		Expect(source).To(RenderHTML5Body(expected))
	})

	It("rich footnote in a paragraph", func() {
		source := `foo footnote:[some *rich* https://foo.com[content]]`
		expected := `<div class="paragraph">
<p>foo <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. some <strong>rich</strong> <a href="https://foo.com">content</a>
</div>
</div>`
		Expect(source).To(RenderHTML5Body(expected))
	})

	It("footnoteref with valid ref in a paragraph", func() {
		source := `a note here footnoteref:[foo, a note for foo] and an there footnoteref:[foo] too`
		expected := `<div class="paragraph">
<p>a note here <sup class="footnote" id="_footnote_foo">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> and an there <sup class="footnoteref">[<a class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> too</p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		Expect(source).To(RenderHTML5Body(expected))
	})

	It("footnoteref with invalid ref in a paragraph", func() {
		source := `a note here footnoteref:[foo, a note for foo] and an unknown there footnoteref:[bar]`
		expected := `<div class="paragraph">
<p>a note here <sup class="footnote" id="_footnote_foo">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup> and an unknown there <sup class="footnoteref red" title="Unresolved footnote reference.">[bar]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a note for foo
</div>
</div>`
		Expect(source).To(RenderHTML5Body(expected))
	})

	It("footnotes everywhere", func() {

		source := `= title
	
a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`

		// differs from asciidoc in the footnotes at the end of the doc, and the section id (numbering)
		expected := `<div id="preamble">
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
		Expect(source).To(RenderHTML5Body(expected))
	})
})

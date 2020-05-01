package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("footnotes", func() {

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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("multiple footnotes including a reference", func() {
		source := `A statement.footnote:[a regular footnote.]   
A bold statement!footnote:disclaimer[Opinions are my own.] 

Another outrageous statement.footnote:disclaimer[]`
		expected := `<div class="paragraph">
<p>A statement.<sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup>
A bold statement!<sup class="footnote" id="_footnote_disclaimer">[<a id="_footnoteref_2" class="footnote" href="#_footnotedef_2" title="View footnote.">2</a>]</sup></p>
</div>
<div class="paragraph">
<p>Another outrageous statement.<sup class="footnoteref">[<a class="footnote" href="#_footnotedef_2" title="View footnote.">2</a>]</sup></p>
</div>
<div id="footnotes">
<hr>
<div class="footnote" id="_footnotedef_1">
<a href="#_footnoteref_1">1</a>. a regular footnote.
</div>
<div class="footnote" id="_footnotedef_2">
<a href="#_footnoteref_2">2</a>. Opinions are my own.
</div>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("footnotes everywhere", func() {

		source := `= title
	
a preamble with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`

		// WARNING: differs from asciidoc in the order of footnotes in the doc and at the end of the doc, and the section id (numbering)
		expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>a preamble with a <sup class="footnote">[<a id="_footnoteref_1" class="footnote" href="#_footnotedef_1" title="View footnote.">1</a>]</sup></p>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

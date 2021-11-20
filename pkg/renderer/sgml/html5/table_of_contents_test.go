package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document toc", func() {

	Context("in document with header", func() {

		Context("with default placement", func() {

			It("should include with default level", func() {
				source := `= A title
:toc:

A preamble...

== Section A

=== Section A.a

=== Section A.b

==== Section that shall not be in ToC

== Section B

=== Section B.a

== Section C`

				expected := `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_section_a">Section A</a>
<ul class="sectlevel2">
<li><a href="#_section_a_a">Section A.a</a></li>
<li><a href="#_section_a_b">Section A.b</a></li>
</ul>
</li>
<li><a href="#_section_b">Section B</a>
<ul class="sectlevel2">
<li><a href="#_section_b_a">Section B.a</a></li>
</ul>
</li>
<li><a href="#_section_c">Section C</a></li>
</ul>
</div>
<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>A preamble&#8230;&#8203;</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
</div>
<div class="sect2">
<h3 id="_section_a_b">Section A.b</h3>
<div class="sect3">
<h4 id="_section_that_shall_not_be_in_toc">Section that shall not be in ToC</h4>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_b">Section B</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_b_a">Section B.a</h3>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_c">Section C</h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include with custom level", func() {
				source := `= A title
:toc:
:toclevels: 4

A preamble...

== Section A

=== Section A.a

=== Section A.b

==== Section A.b.a

===== Section A.b.a.a

== Section B

=== Section B.a

== Section C`

				expected := `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_section_a">Section A</a>
<ul class="sectlevel2">
<li><a href="#_section_a_a">Section A.a</a></li>
<li><a href="#_section_a_b">Section A.b</a>
<ul class="sectlevel3">
<li><a href="#_section_a_b_a">Section A.b.a</a>
<ul class="sectlevel4">
<li><a href="#_section_a_b_a_a">Section A.b.a.a</a></li>
</ul>
</li>
</ul>
</li>
</ul>
</li>
<li><a href="#_section_b">Section B</a>
<ul class="sectlevel2">
<li><a href="#_section_b_a">Section B.a</a></li>
</ul>
</li>
<li><a href="#_section_c">Section C</a></li>
</ul>
</div>
<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>A preamble&#8230;&#8203;</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
</div>
<div class="sect2">
<h3 id="_section_a_b">Section A.b</h3>
<div class="sect3">
<h4 id="_section_a_b_a">Section A.b.a</h4>
<div class="sect4">
<h5 id="_section_a_b_a_a">Section A.b.a.a</h5>
</div>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_b">Section B</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_b_a">Section B.a</h3>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_c">Section C</h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should not include when no sections", func() {
				source := `= sect0
:toc:

level 1 sections not exists.`

				expected := `<div class="paragraph">
<p>level 1 sections not exists.</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))

			})
		})

		Context("within preamble", func() {

			It("should include with default level", func() {
				source := `= A title
:toc: preamble

A preamble...

== Section A

=== Section A.a

=== Section A.b

==== Section that shall not be in ToC

== Section B

=== Section B.a

== Section C`

				expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>A preamble&#8230;&#8203;</p>
</div>
</div>
<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_section_a">Section A</a>
<ul class="sectlevel2">
<li><a href="#_section_a_a">Section A.a</a></li>
<li><a href="#_section_a_b">Section A.b</a></li>
</ul>
</li>
<li><a href="#_section_b">Section B</a>
<ul class="sectlevel2">
<li><a href="#_section_b_a">Section B.a</a></li>
</ul>
</li>
<li><a href="#_section_c">Section C</a></li>
</ul>
</div>
</div>
<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
</div>
<div class="sect2">
<h3 id="_section_a_b">Section A.b</h3>
<div class="sect3">
<h4 id="_section_that_shall_not_be_in_toc">Section that shall not be in ToC</h4>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_b">Section B</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_b_a">Section B.a</h3>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_c">Section C</h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include with custom level", func() {
				source := `= A title
:toc: preamble
:toclevels: 4

A preamble...

== Section A

=== Section A.a

=== Section A.b

==== Section A.b.a

===== Section A.b.a.a

== Section B

=== Section B.a

== Section C`

				expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>A preamble&#8230;&#8203;</p>
</div>
</div>
<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
<ul class="sectlevel1">
<li><a href="#_section_a">Section A</a>
<ul class="sectlevel2">
<li><a href="#_section_a_a">Section A.a</a></li>
<li><a href="#_section_a_b">Section A.b</a>
<ul class="sectlevel3">
<li><a href="#_section_a_b_a">Section A.b.a</a>
<ul class="sectlevel4">
<li><a href="#_section_a_b_a_a">Section A.b.a.a</a></li>
</ul>
</li>
</ul>
</li>
</ul>
</li>
<li><a href="#_section_b">Section B</a>
<ul class="sectlevel2">
<li><a href="#_section_b_a">Section B.a</a></li>
</ul>
</li>
<li><a href="#_section_c">Section C</a></li>
</ul>
</div>
</div>
<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
</div>
<div class="sect2">
<h3 id="_section_a_b">Section A.b</h3>
<div class="sect3">
<h4 id="_section_a_b_a">Section A.b.a</h4>
<div class="sect4">
<h5 id="_section_a_b_a_a">Section A.b.a.a</h5>
</div>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_b">Section B</h2>
<div class="sectionbody">
<div class="sect2">
<h3 id="_section_b_a">Section B.a</h3>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_c">Section C</h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should not include when no sections", func() {
				source := `= sect0
:toc: preamble

level 1 sections not exists.`

				expected := `<div class="paragraph">
<p>level 1 sections not exists.</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))

			})
		})
	})
})

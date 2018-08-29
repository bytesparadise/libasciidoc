package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("document toc", func() {

	Context("document with toc", func() {

		It("toc with default level", func() {
			actualContent := `= A title
:toc:

A preamble...

== Section A

=== Section A.a

=== Section A.b

==== Section that shall not be in TOC

== Section B

=== Section B.a

== Section C`

			expectedResult := `<div id="toc" class="toc">
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
<h4 id="_section_that_shall_not_be_in_toc">Section that shall not be in TOC</h4>
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("toc with custom level", func() {
			actualContent := `= A title
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

			expectedResult := `<div id="toc" class="toc">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})
})

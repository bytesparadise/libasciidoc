package html5_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document toc", func() {

	Context("document with toc", func() {

		It("toc with default level", func() {
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
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("toc with custom level", func() {
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
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("document with no section", func() {
			source := `= sect0
:toc:

level 1 sections not exists.`

			expected := `<div class="paragraph">
<p>level 1 sections not exists.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
	})
})

var _ = Describe("table of contents initialization", func() {

	Context("document without section", func() {

		It("should return empty table of contents when doc has no section", func() {
			actual := types.Document{
				Attributes:        types.Attributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.Attributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			expected := types.TableOfContents{
				Sections: []types.ToCSection{},
			}
			Expect(TableOfContents(actual)).To(Equal(expected))
		})
	})

	Context("document with sections", func() {

		doctitle := []interface{}{
			types.StringElement{Content: "a header"},
		}
		sectionATitle := []interface{}{
			types.StringElement{Content: "Section A with link to "},
			types.InlineLink{
				Location: types.Location{
					Scheme: "https://",
					Path: []interface{}{
						types.StringElement{
							Content: "redhat.com",
						},
					},
				},
			},
		}
		sectionAaTitle := []interface{}{
			types.StringElement{Content: "Section A.a "},
			types.FootnoteReference{
				ID:  1,
				Ref: "foo",
			},
		}
		sectionAa1Title := []interface{}{
			types.StringElement{Content: "Section A.a.1"},
		}
		sectionBTitle := []interface{}{
			types.StringElement{Content: "Section B"},
		}
		document := types.Document{
			Attributes: types.Attributes{},
			ElementReferences: types.ElementReferences{
				"_a_header":    doctitle,
				"_section_a":   sectionATitle,
				"_section_a_a": sectionAaTitle,
				"_section_b":   sectionBTitle,
			},
			Footnotes: []types.Footnote{
				{
					ID:  1,
					Ref: "foo",
				},
			},
			Elements: []interface{}{
				types.Section{
					Attributes: types.Attributes{
						types.AttrID: "_a_header",
					},
					Level: 0,
					Title: doctitle,
					Elements: []interface{}{
						types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_a",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
								types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.Attributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
										types.Section{
											Attributes: types.Attributes{
												types.AttrID: "_section_a_a_1",
											},
											Level: 3,
											Title: sectionAa1Title,
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.Attributes{},
													Lines: [][]interface{}{
														{
															types.StringElement{Content: "a paragraph"},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_b",
							},
							Level: 1,
							Title: sectionBTitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		It("should return table of contents with section level 1,2,3,2 with default level", func() {
			delete(document.Attributes, types.AttrTableOfContentsLevels)
			expected := types.TableOfContents{
				Sections: []types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Title: "Section A with link to https://redhat.com",
						Children: []types.ToCSection{
							{

								ID:       "_section_a_a",
								Level:    2,
								Title:    "Section A.a <sup class=\"footnote\">[1]</sup>",
								Children: []types.ToCSection{},
							},
						},
					},
					{
						ID:       "_section_b",
						Level:    1,
						Title:    "Section B",
						Children: []types.ToCSection{},
					},
				},
			}
			Expect(TableOfContents(document)).To(Equal(expected))
		})

		It("should return table of contents with section level 1,2,3,2 with custom level", func() {
			document.Attributes[types.AttrTableOfContentsLevels] = "4" // must be a string
			expected := types.TableOfContents{
				Sections: []types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Title: "Section A with link to https://redhat.com",
						Children: []types.ToCSection{
							{

								ID:    "_section_a_a",
								Level: 2,
								Title: "Section A.a <sup class=\"footnote\">[1]</sup>",
								Children: []types.ToCSection{
									{

										ID:       "_section_a_a_1",
										Level:    3,
										Title:    "Section A.a.1",
										Children: []types.ToCSection{},
									},
								},
							},
						},
					},
					{
						ID:       "_section_b",
						Level:    1,
						Title:    "Section B",
						Children: []types.ToCSection{},
					},
				},
			}
			Expect(TableOfContents(document)).To(Equal(expected))
		})
	})

})

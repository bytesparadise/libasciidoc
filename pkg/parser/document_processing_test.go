package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document processing", func() {

	It("should retain attributes passed in configuration", func() {
		source := `[source]
----
foo
----`
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrSyntaxHighlighter: "pygments",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.DelimitedBlock{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Source,
					},
					Kind: types.Source,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "foo",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDocument(source, configuration.WithAttributes(map[string]string{
			types.AttrSyntaxHighlighter: "pygments",
		}))).To(Equal(expected))
	})

	It("should include toc and preamble", func() {
		source := `= A title
:toc:

A preamble...

== Section A

=== Section A.a

== Section B

== Section C`
		titleSection0 := []interface{}{
			types.StringElement{
				Content: "A title",
			},
		}
		titleSectionA := []interface{}{
			types.StringElement{
				Content: "Section A",
			},
		}
		titleSectionAa := []interface{}{
			types.StringElement{
				Content: "Section A.a",
			},
		}
		titleSectionB := []interface{}{
			types.StringElement{
				Content: "Section B",
			},
		}
		titleSectionC := []interface{}{
			types.StringElement{
				Content: "Section C",
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Title: titleSection0,
					Elements: []interface{}{
						types.TableOfContentsPlaceHolder{},
						types.Preamble{
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "A preamble...",
											},
										},
									},
								},
							},
						},
						types.Section{
							Level: 1,
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_a",
							},
							Title: titleSectionA,
							Elements: []interface{}{
								types.Section{
									Level: 2,
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_a",
									},
									Title:    titleSectionAa,
									Elements: []interface{}{},
								},
							},
						},
						types.Section{
							Level: 1,
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_b",
							},
							Title:    titleSectionB,
							Elements: []interface{}{},
						},
						types.Section{
							Level: 1,
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_c",
							},
							Title:    titleSectionC,
							Elements: []interface{}{},
						},
					},
				},
			},
			ElementReferences: types.ElementReferences{
				"_a_title":     titleSection0,
				"_section_a":   titleSectionA,
				"_section_a_a": titleSectionAa,
				"_section_b":   titleSectionB,
				"_section_c":   titleSectionC,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
		}
		Expect(ParseDocument(source)).To(Equal(expected))
	})
})

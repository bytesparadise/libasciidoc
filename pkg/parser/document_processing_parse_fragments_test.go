package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document fragment parsing", func() {

	Context("paragraphs", func() {

		It("should parse 1 paragraph with single line", func() {
			source := `a line`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a line"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse 2 paragraphs with single line each", func() {
			source := `a line
		
another line`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a line"),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("another line"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse markdown quote block with single marker", func() {
			source := `> some text
on *multiple lines*`

			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.RawLine("some text"),
								types.RawLine("on *multiple lines*"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse markdown quote block with multiple markers", func() {
			source := `> some text
> on *multiple lines*`

			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.RawLine("some text"),
								types.RawLine("on *multiple lines*"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})
	})

	Context("as delimited blocks", func() {

		It("should parse 1 delimited block with single rawline", func() {
			source := `----
a line
----`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.RawLine("a line"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should collect 1 delimited block with multiple rawlines only", func() {
			source := `----
a line

****
not a sidebar block
****
----
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.RawLine("a line"),
								types.RawLine(""),
								types.RawLine("****"),
								types.RawLine("not a sidebar block"),
								types.RawLine("****"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should collect 1 delimited block with attributes, multiple rawlines and content afterwards", func() {
			source := `[source,text]
----
a line

another line
----


a paragraph
on
3 lines.

`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Attributes: types.Attributes{
								types.AttrStyle:    "source",
								types.AttrLanguage: "text",
							},
							Elements: []interface{}{
								types.RawLine("a line"),
								types.RawLine(""),
								types.RawLine("another line"),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a paragraph"),
								types.RawLine("on"),
								types.RawLine("3 lines."),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})
	})

	Context("sections", func() {

		It("should collect 1 section and content afterwards", func() {
			source := `== section title


a paragraph
on
3 lines.

`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Title: []interface{}{
								types.RawLine("section title"),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a paragraph"),
								types.RawLine("on"),
								types.RawLine("3 lines."),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})
	})

	Context("lists", func() {

		It("should parse callout list elements without blankline in-between", func() {
			source := `<1> first element
<2> second element
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.CalloutListElement{
									Ref: 1,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("first element"),
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 2,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("second element"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse callout list elements with blanklines in-between", func() {
			source := `<1> first element

		
<2> second element
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.CalloutListElement{
									Ref: 1,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("first element"),
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 2,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("second element"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse ordered list elements without blanklines in-between", func() {
			source := `. first element
. second element
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("first element"),
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("second element"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse ordered list elements with blanklines in-between", func() {
			source := `. first element


. second element
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("first element"),
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												types.RawLine("second element"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse single ordered list element with multiple lines", func() {
			source := `. element 
on 
multiple 
lines
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												// suffix spaces are trimmed
												types.RawLine("element"),
												types.RawLine("on"),
												types.RawLine("multiple"),
												types.RawLine("lines"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse multiple ordered list element with multiple lines", func() {
			source := `. first element 
on 
multiple 
lines
. second element 
on 
multiple 
lines
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												// suffix spaces are trimmed on each line
												types.RawLine("first element"),
												types.RawLine("on"),
												types.RawLine("multiple"),
												types.RawLine("lines"),
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												// suffix spaces are trimmed on each line
												types.RawLine("second element"),
												types.RawLine("on"),
												types.RawLine("multiple"),
												types.RawLine("lines"),
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})
	})

})

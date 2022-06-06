package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("document fragment parsing", func() {

	Context("paragraphs", func() {

		It("should parse 1 paragraph with single line", func() {
			source := `a line`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   6,
					},
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
					Position: types.Position{
						Start: 0,
						End:   7,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a line"),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 7,
						End:   10,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 10,
						End:   22,
					},
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
					Position: types.Position{
						Start: 0,
						End:   31,
					},
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.RawLine("some text\n"),
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
					Position: types.Position{
						Start: 0,
						End:   33,
					},
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.RawLine("some text\n"),
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
					Position: types.Position{
						Start: 0,
						End:   16,
					},
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
					Position: types.Position{
						Start: 0,
						End:   48,
					},
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.RawLine("a line\n"),
								types.RawLine("\n"),
								types.RawLine("****\n"),
								types.RawLine("not a sidebar block\n"),
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
					Position: types.Position{
						Start: 0,
						End:   45,
					},
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Attributes: types.Attributes{
								types.AttrStyle:    "source",
								types.AttrLanguage: "text",
							},
							Elements: []interface{}{
								types.RawLine("a line\n"),
								types.RawLine("\n"),
								types.RawLine("another line"),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 45,
						End:   46,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 46,
						End:   47,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 47,
						End:   71,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a paragraph\n"),
								types.RawLine("on\n"),
								types.RawLine("3 lines."),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 71,
						End:   72,
					},
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
					Position: types.Position{
						Start: 0,
						End:   17,
					},
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Title: []interface{}{
								&types.StringElement{
									Content: "section title",
								},
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 17,
						End:   18,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 18,
						End:   19,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 19,
						End:   43,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("a paragraph\n"),
								types.RawLine("on\n"),
								types.RawLine("3 lines."),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 43,
						End:   44,
					},
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
					Position: types.Position{
						Start: 0,
						End:   37,
					},
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
					Position: types.Position{
						Start: 0,
						End:   41,
					},
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
					Position: types.Position{
						Start: 0,
						End:   33,
					},
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
					Position: types.Position{
						Start: 0,
						End:   35,
					},
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
					Position: types.Position{
						Start: 0,
						End:   31,
					},
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												// suffix spaces are trimmed
												types.RawLine("element\n"),
												types.RawLine("on\n"),
												types.RawLine("multiple\n"),
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
					Position: types.Position{
						Start: 0,
						End:   75,
					},
					Elements: []interface{}{
						&types.ListElements{
							Elements: []interface{}{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												// suffix spaces are trimmed on each line
												types.RawLine("first element\n"),
												types.RawLine("on\n"),
												types.RawLine("multiple\n"),
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
												types.RawLine("second element\n"),
												types.RawLine("on\n"),
												types.RawLine("multiple\n"),
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

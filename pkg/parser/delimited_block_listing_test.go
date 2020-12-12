package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("listing blocks", func() {

	Context("draft documents", func() {

		Context("delimited blocks", func() {

			It("with single rich line", func() {
				source := `----
some *listing* code
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some *listing* code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.

----
some listing code
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{}, // blankline is required between paragraph and the next block
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single callout", func() {
				source := `----
<import> <1>
----
<1> an import`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "import",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " ",
									},
									types.Callout{
										Ref: 1,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a single import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 3,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a single basic import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 4,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

		})

		Context("paragraph blocks", func() {

			It("with single rich line", func() {
				source := `[listing]
some *listing* content`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Listing,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some *listing* content", // no quote substitution
									},
								},
							},
						},
					},
				}

				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("final documents", func() {

		Context("delimited blocks", func() {

			It("with single rich line", func() {
				source := `----
some *listing* code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some *listing* code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.
	
----
some listing code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single callout", func() {
				source := `----
import <1>
----
<1> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a func",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a single import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 3,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a single basic import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 4,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a func",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("paragraph blocks", func() {

			It("with single rich line", func() {
				source := `[listing]
some *listing* content`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Listing,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some *listing* content", // no quote substitution
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})

	Context("with custom substitutions", func() {

	})

})

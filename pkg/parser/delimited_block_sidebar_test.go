package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sidebar blocks", func() {

	Context("in final documents", func() {

		It("with rich content", func() {
			source := `****
some *verse* content
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "some ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "verse",
											},
										},
									},
									&types.StringElement{
										Content: " content",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with single line starting with a dot", func() {
			source := `
****
.standalone
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with paragraph", func() {
			source := `****
some *verse* content
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "some ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "verse",
											},
										},
									},
									&types.StringElement{
										Content: " content",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with title, paragraph and sourcecode block", func() {
			source := `.a title
****
some *verse* content

----
foo
bar
----
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
						Attributes: types.Attributes{
							types.AttrTitle: "a title",
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "some ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "verse",
											},
										},
									},
									&types.StringElement{
										Content: " content",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Listing,
								Elements: []interface{}{
									&types.StringElement{
										Content: "foo\nbar",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with single line starting with a dot", func() {
			source := `
****
.standalone
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with last line starting with a dot", func() {
			source := `
****
some content

.standalone
****`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Sidebar,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "some content",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		Context("with variable delimiter length", func() {

			It("with 5 chars", func() {
				source := `*****
some *sidebar* content
*****`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Sidebar,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some ",
										},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{
													Content: "sidebar",
												},
											},
										},
										&types.StringElement{
											Content: " content",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with 5 chars with nested with 4 chars", func() {
				source := `*****
****
some *sidebar* content
****
*****`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Sidebar,
							Elements: []interface{}{
								&types.DelimitedBlock{
									Kind: types.Sidebar,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "some ",
												},
												&types.QuotedText{
													Kind: types.SingleQuoteBold,
													Elements: []interface{}{
														&types.StringElement{
															Content: "sidebar",
														},
													},
												},
												&types.StringElement{
													Content: " content",
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
		})
	})
})

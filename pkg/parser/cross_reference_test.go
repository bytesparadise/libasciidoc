package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cross references - draft", func() {

	Context("internal references", func() {

		It("cross reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "with some content linked to ",
								},
								types.InternalCrossReference{
									ID:    "thetitle",
									Label: "",
								},
								types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("cross reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "with some content linked to ",
								},
								types.InternalCrossReference{
									ID:    "thetitle",
									Label: "a label to the title",
								},
								types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("external references", func() {

		It("external cross reference to other doc with plain text location and rich label", func() {
			source := `some content linked to xref:another-doc.adoc[*another doc*]!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "some content linked to ",
								},
								types.ExternalCrossReference{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "another-doc.adoc",
											},
										},
									},
									Label: []interface{}{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "another doc",
												},
											},
										},
									},
								},
								types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("external cross reference to other doc with document attribute in location", func() {
			source := `some content linked to xref:{foo}.adoc[another doc]!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "some content linked to ",
								},
								types.ExternalCrossReference{
									Location: types.Location{
										Elements: []interface{}{
											types.DocumentAttributeSubstitution{
												Name: "foo",
											},
											types.StringElement{
												Content: ".adoc",
											},
										},
									},
									Label: []interface{}{
										types.StringElement{
											Content: "another doc",
										},
									},
								},
								types.StringElement{
									Content: "!",
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

var _ = Describe("cross references - document", func() {

	Context("internal references", func() {

		It("cross reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": []interface{}{
						types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes: []types.Footnote{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "with some content linked to ",
										},
										types.InternalCrossReference{
											ID:    "thetitle",
											Label: "",
										},
										types.StringElement{
											Content: "!",
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

		It("cross reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": []interface{}{
						types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes: []types.Footnote{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "with some content linked to ",
										},
										types.InternalCrossReference{
											ID:    "thetitle",
											Label: "a label to the title",
										},
										types.StringElement{
											Content: "!",
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

	Context("external references", func() {

		It("external cross reference to other doc with plain text location and rich label", func() {
			source := `some content linked to xref:another-doc.adoc[*another doc*]!`
			expected := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "some content linked to ",
								},
								types.ExternalCrossReference{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "another-doc.adoc",
											},
										},
									},
									Label: []interface{}{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "another doc",
												},
											},
										},
									},
								},
								types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("external cross reference to other doc with document attribute in location and label with special chars", func() {
			source := `
:foo: another-doc.adoc

some content linked to xref:{foo}[another_doc()]!`
			expected := types.Document{
				Attributes: types.DocumentAttributes{
					"foo": "another-doc.adoc",
				},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "some content linked to ",
								},
								types.ExternalCrossReference{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "another-doc.adoc",
											},
										},
									},
									Label: []interface{}{
										types.StringElement{
											Content: "another_doc()",
										},
									},
								},
								types.StringElement{
									Content: "!",
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

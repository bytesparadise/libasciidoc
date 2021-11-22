package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("q and a lists", func() {

	It("with title", func() {
		source := `.Q&A
[qanda]
What is libasciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expected := &types.Document{
			Elements: []interface{}{
				&types.List{
					Kind: types.LabeledListKind,
					Attributes: types.Attributes{
						types.AttrTitle: "Q&A",
						types.AttrStyle: "qanda",
					},
					Elements: []types.ListElement{
						&types.LabeledListElement{
							Style: "::",
							Term: []interface{}{
								&types.StringElement{
									Content: "What is libasciidoc?",
								},
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											// heading spaces are trimmed
											Content: "An implementation of the AsciiDoc processor in Golang.",
										},
									},
								},
							},
						},
						&types.LabeledListElement{
							Style: "::",
							Term: []interface{}{
								&types.StringElement{
									Content: "What is the answer to the Ultimate Question?",
								},
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "42",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		result, err := ParseDocument(source)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(MatchDocument(expected))
	})

	It("with role and id", func() {
		source := `.Q&A
[qanda#quiz]
[.role1.role2]
What is libasciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expected := &types.Document{
			Elements: []interface{}{
				&types.List{
					Kind: types.LabeledListKind,
					Attributes: types.Attributes{
						types.AttrTitle: "Q&A",
						types.AttrStyle: "qanda",
						types.AttrID:    "quiz",
						types.AttrRoles: []interface{}{"role1", "role2"},
					},
					Elements: []types.ListElement{
						&types.LabeledListElement{
							Style: "::",
							Term: []interface{}{
								&types.StringElement{Content: "What is libasciidoc?"},
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "An implementation of the AsciiDoc processor in Golang.", // heading spaces are trimmed
										},
									},
								},
							},
						},
						&types.LabeledListElement{
							Style: "::",
							Term: []interface{}{
								&types.StringElement{
									Content: "What is the answer to the Ultimate Question?",
								},
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "42",
										},
									},
								},
							},
						},
					},
				},
			},
			ElementReferences: types.ElementReferences{
				"quiz": "Q&A",
			},
		}
		Expect(ParseDocument(source)).To(MatchDocument(expected))
	})

})

package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("q and a lists", func() {

	It("q and a with title", func() {
		source := `.Q&A
[qanda]
What is libasciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expected := types.Document{
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.Attributes{
						types.AttrTitle: "Q&A",
						types.AttrStyle: "qanda",
					},
					Items: []types.LabeledListItem{
						{
							Level: 1,
							Term: []interface{}{
								types.StringElement{
									Content: "What is libasciidoc?",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "An implementation of the AsciiDoc processor in Golang.",
											},
										},
									},
								},
							},
						},
						{
							Level: 1,
							Term: []interface{}{
								types.StringElement{
									Content: "What is the answer to the Ultimate Question?",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "42",
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
		result, err := ParseDocument(source)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(MatchDocument(expected))
	})

	It("q and a with role and id", func() {
		source := `.Q&A
[qanda#quiz]
[.role1.role2]
What is libasciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expected := types.Document{
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.Attributes{
						types.AttrTitle:    "Q&A",
						types.AttrStyle:    "qanda",
						types.AttrID:       "quiz",
						types.AttrCustomID: true,
						types.AttrRole:     []interface{}{types.ElementRole{"role1"}, types.ElementRole{"role2"}},
					},
					Items: []types.LabeledListItem{
						{
							Level: 1,
							Term: []interface{}{
								types.StringElement{Content: "What is libasciidoc?"},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "An implementation of the AsciiDoc processor in Golang.",
											},
										},
									},
								},
							},
						},
						{
							Level: 1,
							Term: []interface{}{
								types.StringElement{
									Content: "What is the answer to the Ultimate Question?",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "42",
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

})

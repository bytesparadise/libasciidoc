package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
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
			// do not show parse errors in the logs for this test
			_, reset := ConfigureLogger(log.FatalLevel)
			defer reset()
			source := `
****
.foo
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
			// do not show parse errors in the logs for this test
			_, reset := ConfigureLogger(log.FatalLevel)
			defer reset()
			source := `
****
.foo
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
	})
})

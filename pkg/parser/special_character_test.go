package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("special characters", func() {

	Context("in final documents", func() {

		It("should parse in paragraph", func() {
			source := "<b>*</b> &apos; &amp;"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.SpecialCharacter{
								Name: "<",
							},
							&types.StringElement{
								Content: "b",
							},
							&types.SpecialCharacter{
								Name: ">",
							},
							&types.StringElement{
								Content: "*",
							},
							&types.SpecialCharacter{
								Name: "<",
							},
							&types.StringElement{
								Content: "/b",
							},
							&types.SpecialCharacter{
								Name: ">",
							},
							&types.StringElement{
								Content: " ",
							},
							&types.SpecialCharacter{
								Name: "&",
							},
							&types.StringElement{
								Content: "apos; ",
							},
							&types.SpecialCharacter{
								Name: "&",
							},
							&types.StringElement{
								Content: "amp;",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse in delimited block", func() {
			source := "```" + "\n" +
				"<b>*</b> &apos; &amp;" + "\n" +
				"```"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.SpecialCharacter{
								Name: "<",
							},
							&types.StringElement{
								Content: "b",
							},
							&types.SpecialCharacter{
								Name: ">",
							},
							&types.StringElement{
								Content: "*",
							},
							&types.SpecialCharacter{
								Name: "<",
							},
							&types.StringElement{
								Content: "/b",
							},
							&types.SpecialCharacter{
								Name: ">",
							},
							&types.StringElement{
								Content: " ",
							},
							&types.SpecialCharacter{
								Name: "&",
							},
							&types.StringElement{
								Content: "apos; ",
							},
							&types.SpecialCharacter{
								Name: "&",
							},
							&types.StringElement{
								Content: "amp;",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})

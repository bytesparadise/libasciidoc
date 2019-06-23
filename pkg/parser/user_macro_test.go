package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("user macros", func() {

	Context("inline macros", func() {

		It("inline macro empty", func() {
			source := "AAA hello:[]"
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "AAA ",
						},
						&types.UserMacro{
							Kind:       types.InlineMacro,
							Name:       "hello",
							Value:      "",
							Attributes: types.ElementAttributes{},
							RawText:    "hello:[]",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("inline macro with attribute", func() {
			source := `AAA hello:[suffix="!!!!!"]`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "AAA ",
						},
						&types.UserMacro{
							Kind:  types.InlineMacro,
							Name:  "hello",
							Value: "",
							Attributes: types.ElementAttributes{
								"suffix": "!!!!!",
							},
							RawText: `hello:[suffix="!!!!!"]`,
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("inline macro with value", func() {
			source := `AAA hello:John Doe[]`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "AAA ",
						},
						&types.UserMacro{
							Kind:       types.InlineMacro,
							Name:       "hello",
							Value:      "John Doe",
							Attributes: types.ElementAttributes{},
							RawText:    "hello:John Doe[]",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("inline user macro with value and attributes", func() {
			source := "repository: git:some/url.git[key1=value1,key2=value2]"
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "repository: ",
						},
						&types.UserMacro{
							Kind:  types.InlineMacro,
							Name:  "git",
							Value: "some/url.git",
							Attributes: types.ElementAttributes{
								"key1": "value1",
								"key2": "value2",
							},
							RawText: "git:some/url.git[key1=value1,key2=value2]",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})
	})

	Context("user macros", func() {

		It("user macro block without value", func() {

			source := "git::[]"
			expected := &types.UserMacro{
				Kind:       types.BlockMacro,
				Name:       "git",
				Value:      "",
				Attributes: types.ElementAttributes{},
				RawText:    "git::[]",
			}
			verifyDocumentBlock(expected, source)
		})

		It("user block macro with value and attributes", func() {
			source := "git::some/url.git[key1=value1,key2=value2]"
			expected := &types.UserMacro{
				Kind:  types.BlockMacro,
				Name:  "git",
				Value: "some/url.git",
				Attributes: types.ElementAttributes{
					"key1": "value1",
					"key2": "value2",
				},
				RawText: "git::some/url.git[key1=value1,key2=value2]",
			}
			verifyDocumentBlock(expected, source)
		})

		It("user macro block with attribute", func() {
			source := `git::[key1="value1"]`
			expected := &types.UserMacro{
				Kind:  types.BlockMacro,
				Name:  "git",
				Value: "",
				Attributes: types.ElementAttributes{
					"key1": "value1",
				},
				RawText: `git::[key1="value1"]`,
			}
			verifyDocumentBlock(expected, source)
		})

		It("user macro block with value", func() {
			source := `git::some/url.git[]`
			expected := &types.UserMacro{
				Kind:       types.BlockMacro,
				Name:       "git",
				Value:      "some/url.git",
				Attributes: types.ElementAttributes{},
				RawText:    "git::some/url.git[]",
			}
			verifyDocumentBlock(expected, source)
		})

	})
})

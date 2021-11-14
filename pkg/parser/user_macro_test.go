package parser_test

import (
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("user macros", func() {

	Context("in final documents", func() {

		userTmpl := &texttemplate.Template{}

		Context("inline macros", func() {

			It("without attributes", func() {
				source := "AAA hello:[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "AAA ",
								},
								&types.UserMacro{
									Kind:    types.InlineMacro,
									Name:    "hello",
									Value:   "",
									RawText: "hello:[]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("hello", userTmpl))).To(MatchDocument(expected))
			})

			It("with double quoted attributes", func() {
				source := `AAA hello:[prefix="hello ",suffix="!!"]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "AAA ",
								},
								&types.UserMacro{
									Kind:  types.InlineMacro,
									Name:  "hello",
									Value: "",
									Attributes: types.Attributes{
										"prefix": "hello ",
										"suffix": "!!",
									},
									RawText: `hello:[prefix="hello ",suffix="!!"]`,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("hello", userTmpl))).To(MatchDocument(expected))
			})

			It("with value", func() {
				source := `AAA hello:JohnDoe[]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "AAA ",
								},
								&types.UserMacro{
									Kind:    types.InlineMacro,
									Name:    "hello",
									Value:   "JohnDoe",
									RawText: "hello:JohnDoe[]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("hello", userTmpl))).To(MatchDocument(expected))
			})

			It("with value and attributes", func() {
				source := "repository: git:some/url.git[key1=value1,key2=value2]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "repository: ",
								},
								&types.UserMacro{
									Kind:  types.InlineMacro,
									Name:  "git",
									Value: "some/url.git",
									Attributes: types.Attributes{
										"key1": "value1",
										"key2": "value2",
									},
									RawText: "git:some/url.git[key1=value1,key2=value2]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("git", userTmpl))).To(MatchDocument(expected))
			})

			It("unknown", func() {
				source := "AAA hello:[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "AAA hello:[]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("block macros", func() {

			It("without value", func() {

				source := "git::[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.UserMacro{
							Kind:    types.BlockMacro,
							Name:    "git",
							Value:   "",
							RawText: "git::[]",
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("git", userTmpl))).To(MatchDocument(expected))
			})

			It("with value and attributes", func() {
				source := "git::some/url.git[key1=value1,key2=value2]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.UserMacro{
							Kind:  types.BlockMacro,
							Name:  "git",
							Value: "some/url.git",
							Attributes: types.Attributes{
								"key1": "value1",
								"key2": "value2",
							},
							RawText: "git::some/url.git[key1=value1,key2=value2]",
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("git", userTmpl))).To(MatchDocument(expected))
			})

			It("with attribute", func() {
				source := `git::[key1="value1"]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.UserMacro{
							Kind:  types.BlockMacro,
							Name:  "git",
							Value: "",
							Attributes: types.Attributes{
								"key1": "value1",
							},
							RawText: `git::[key1="value1"]`,
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("git", userTmpl))).To(MatchDocument(expected))
			})

			It("with value only", func() {
				source := `git::some/url.git[]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.UserMacro{
							Kind:    types.BlockMacro,
							Name:    "git",
							Value:   "some/url.git",
							RawText: "git::some/url.git[]",
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithMacroTemplate("git", userTmpl))).To(MatchDocument(expected))
			})

			It("unknown", func() {

				source := "git::[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "git::[]",
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

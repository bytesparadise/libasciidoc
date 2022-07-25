package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("symbols", func() {

	Context("in final documents", func() {

		Context("m-dashes", func() {

			It("should detect between word characters", func() {
				source := "some text--idea apart--continues here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some text",
								},
								&types.Symbol{
									Name: "--",
								},
								&types.StringElement{
									Content: "idea apart",
								},
								&types.Symbol{
									Name: "--",
								},
								&types.StringElement{
									Content: "continues here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between word character and line boundary", func() {
				source := "some text--idea apart--"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some text",
								},
								&types.Symbol{
									Name: "--",
								},
								&types.StringElement{
									Content: "idea apart",
								},
								&types.Symbol{
									Name: "--",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between spaces", func() {
				source := "some text -- idea apart -- continues here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some text",
								},
								&types.Symbol{
									Name: " -- ",
								},
								&types.StringElement{
									Content: "idea apart",
								},
								&types.Symbol{
									Name: " -- ",
								},
								&types.StringElement{
									Content: "continues here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between space and line boudary", func() {
				source := "some text -- idea apart --"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some text",
								},
								&types.Symbol{
									Name: " -- ",
								},
								&types.StringElement{
									Content: "idea apart",
								},
								&types.Symbol{
									Name: " -- ",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("invalid", func() {

				It("should not detect when missing spaces", func() {
					source := "some text --idea apart-- continues here" // `--idea` and `apart--` are missing spaces between characters and dashes
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "some text --idea apart-- continues here",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("single right arrows", func() {

			It("should detect between spaces", func() {
				source := "go -> here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "->",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between character and space", func() {
				source := "go-> here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "->",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between space and character", func() {
				source := "go ->here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "->",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between characters", func() {
				source := "go->here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "->",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("single left arrows", func() {

			It("should detect between spaces", func() {
				source := "go <- here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "<-",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between character and space", func() {
				source := "go<- here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "<-",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between space and character", func() {
				source := "go <-here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "<-",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between characters", func() {
				source := "go<-here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "<-",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("double right arrows", func() {

			It("should detect between spaces", func() {
				source := "go => here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "=>",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between character and space", func() {
				source := "go=> here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "=>",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between space and character", func() {
				source := "go =>here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "=>",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between characters", func() {
				source := "go=>here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "=>",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("double left arrows", func() {

			It("should detect between spaces", func() {
				source := "go <= here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "<=",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between character and space", func() {
				source := "go<= here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "<=",
								},
								&types.StringElement{
									Content: " here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between space and character", func() {
				source := "go <=here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go ",
								},
								&types.Symbol{
									Name: "<=",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect between characters", func() {
				source := "go<=here"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "go",
								},
								&types.Symbol{
									Name: "<=",
								},
								&types.StringElement{
									Content: "here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("trademark symbol", func() {

			It("should detect after space", func() {
				source := "registered (TM)"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "registered ",
								},
								&types.Symbol{
									Name: "(TM)",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should detect within word", func() {
				source := "registered(TM)"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "registered",
								},
								&types.Symbol{
									Name: "(TM)",
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

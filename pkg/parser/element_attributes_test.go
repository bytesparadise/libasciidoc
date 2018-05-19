package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("element attributes", func() {

	Context("element link", func() {

		Context("valid syntax", func() {
			It("element link alone", func() {
				actualContent := "[link=http://foo.bar]"
				expectedResult := map[string]interface{}{"link": "http://foo.bar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
			It("spaces in link", func() {
				actualContent := "[link= http://foo.bar  ]"
				expectedResult := map[string]interface{}{"link": "http://foo.bar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
		})

		Context("invalid syntax", func() {
			It("spaces before keyword", func() {
				actualContent := "[ link=http://foo.bar]"
				expectedResult := types.InvalidElementAttribute{Value: "[ link=http://foo.bar]"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			Context("Unbalanced brackets", func() {
				actualContent := "[link=http://foo.bar"
				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := types.InlineElements{
						types.StringElement{
							Content: "[link=http://foo.bar",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
				})
			})
		})
	})

	Context("element id", func() {

		Context("valid syntax", func() {

			It("normal syntax", func() {
				actualContent := "[[img-foobar]]"
				expectedResult := types.ElementID{Value: "img-foobar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			It("short-hand syntax", func() {
				actualContent := "[#img-foobar]"
				expectedResult := types.ElementID{Value: "img-foobar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
		})

		Context("invalid syntax", func() {

			It("extra spaces", func() {
				actualContent := "[ #img-foobar ]"
				expectedResult := types.InvalidElementAttribute{Value: "[ #img-foobar ]"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			Context("Unbalanced brackets", func() {
				actualContent := "[#img-foobar"

				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := types.InlineElements{
						types.StringElement{
							Content: "[#img-foobar",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
				})
			})
		})
	})
	Context("element title", func() {

		Context("valid syntax", func() {

			It("element title", func() {
				actualContent := ".a title"
				expectedResult := types.ElementTitle{Value: "a title"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
		})

		Context("invalid syntax", func() {
			Context("extra space after dot", func() {

				actualContent := ". a title"
				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := types.InlineElements{
						types.StringElement{
							Content: ". a title",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
				})
			})

			Context("not a dot", func() {
				actualContent := "!a title"

				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := types.InlineElements{
						types.StringElement{
							Content: "!a title",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
				})
			})
		})
	})
})

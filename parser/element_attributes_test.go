package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Element Attributes", func() {

	Context("Element link", func() {

		Context("valid syntax", func() {
			It("element link alone", func() {
				actualContent := "[link=http://foo.bar]"
				expectedResult := &types.ElementLink{Path: "http://foo.bar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
			It("spaces in link", func() {
				actualContent := "[link= http://foo.bar  ]"
				expectedResult := &types.ElementLink{Path: "http://foo.bar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
		})

		Context("invalid syntax", func() {
			It("spaces before keywork", func() {
				actualContent := "[ link=http://foo.bar]"
				expectedResult := &types.InvalidElementAttribute{Value: "[ link=http://foo.bar]"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			Context("unbalanced brackets", func() {
				actualContent := "[link=http://foo.bar"
				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{
								Content: "[link=http://foo.bar",
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
				})
			})
		})
	})

	Context("element id", func() {

		Context("valid syntax", func() {

			It("normal syntax", func() {
				actualContent := "[[img-foobar]]"
				expectedResult := &types.ElementID{Value: "img-foobar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			It("short-hand syntax", func() {
				actualContent := "[#img-foobar]"
				expectedResult := &types.ElementID{Value: "img-foobar"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})
		})

		Context("invalid syntax", func() {

			It("extra spaces", func() {
				actualContent := "[ #img-foobar ]"
				expectedResult := &types.InvalidElementAttribute{Value: "[ #img-foobar ]"}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("ElementAttribute"))
			})

			Context("unbalanced brackets", func() {
				actualContent := "[#img-foobar"

				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{
								Content: "[#img-foobar",
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
				})
			})
		})
	})
	Context("element title", func() {

		Context("valid syntax", func() {

			It("element title", func() {
				actualContent := ".a title"
				expectedResult := &types.ElementTitle{Value: "a title"}
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
					expectedResult := &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{
								Content: ". a title",
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
				})
			})

			Context("not a dot", func() {
				actualContent := "!a title"

				It("cannot be an attribute", func() {
					expectError(GinkgoT(), actualContent, parser.Entrypoint("ElementAttribute"))
				})

				It("is an inline content", func() {
					expectedResult := &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{
								Content: "!a title",
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
				})
			})
		})
	})
})

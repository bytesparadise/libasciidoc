package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Meta Elements", func() {

	Context("Element link", func() {

		It("element link alone", func() {
			actualContent := "[link=http://foo.bar]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.ElementLink{Path: "http://foo.bar"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element link with spaces", func() {
			actualContent := "[ link = http://foo.bar ]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.ElementLink{Path: "http://foo.bar"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element link invalid", func() {
			actualContent := "[ link = http://foo.bar"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.DocElement{
									&types.StringElement{Content: "[ link = "},
									&types.ExternalLink{URL: "http://foo.bar"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})

	Context("element id", func() {

		It("element id", func() {
			actualContent := "[#img-foobar]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.ElementID{Value: "img-foobar"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element id with spaces", func() {
			actualContent := "[ #img-foobar ]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.ElementID{Value: "img-foobar"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element id invalid", func() {
			actualContent := "[#img-foobar"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "[#img-foobar"}}},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})
	Context("element title", func() {

		It("element title", func() {
			actualContent := ".a title"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.ElementTitle{Value: "a title"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element title invalid1", func() {
			actualContent := ". a title"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: ". a title"}}},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("element title invalid2", func() {
			actualContent := "!a title"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "!a title"}}},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})
})

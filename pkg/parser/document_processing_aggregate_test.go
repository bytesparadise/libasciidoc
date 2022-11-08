package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("aggregate fragments", func() {

	// reusable elements
	doctitle := []interface{}{
		&types.StringElement{Content: "A Title"},
	}
	paragraph := &types.Paragraph{
		Elements: []interface{}{
			&types.StringElement{Content: "A short preamble"},
		},
	}
	section1Title := []interface{}{
		&types.StringElement{Content: "section 1"},
	}
	section1 := &types.Section{
		Level: 1,
		Attributes: types.Attributes{
			types.AttrID: "_Section_1",
		},
		Title: section1Title,
	}

	It("with default placement and no header with section1", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name: types.AttrTableOfContents,
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				section1,
			},
		}
		close(c)
		expected := &types.Document{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name: types.AttrTableOfContents,
				},
				paragraph,
				section1,
			},
			ElementReferences: types.ElementReferences{
				"_Section_1": section1Title,
			},
			TableOfContents: &types.TableOfContents{
				MaxDepth: 2,
				Sections: []*types.ToCSection{
					{
						ID:    "_Section_1",
						Level: 1,
					},
				},
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})

	It("with default placement and a header with section1", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name: types.AttrTableOfContents,
						},
					},
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				section1,
			},
		}
		close(c)
		expected := &types.Document{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name: types.AttrTableOfContents,
						},
					},
				},
				&types.Preamble{
					Elements: []interface{}{
						paragraph,
					},
				},
				section1,
			},
			ElementReferences: types.ElementReferences{
				"_Section_1": section1Title,
			},
			TableOfContents: &types.TableOfContents{
				MaxDepth: 2,
				Sections: []*types.ToCSection{
					{
						ID:    "_Section_1",
						Level: 1,
					},
				},
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})

	It("with default placement and a header without section1", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name: types.AttrTableOfContents,
						},
					},
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		close(c)
		expected := &types.Document{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name: types.AttrTableOfContents,
						},
					},
				},
				paragraph, // not wrapped in a preamble since there is nothing afterwards
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})

	It("with preamble placement and no header with section1", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name:  types.AttrTableOfContents,
					Value: "preamble",
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				section1,
			},
		}
		close(c)
		expected := &types.Document{
			Elements: []interface{}{ // no preamble since no header, thus no ToC either
				&types.AttributeDeclaration{
					Name:  types.AttrTableOfContents,
					Value: "preamble",
				},
				paragraph,
				section1,
			},
			ElementReferences: types.ElementReferences{
				"_Section_1": section1Title,
			},
			TableOfContents: &types.TableOfContents{
				MaxDepth: 2,
				Sections: []*types.ToCSection{
					{
						ID:    "_Section_1",
						Level: 1,
					},
				},
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})

	It("with preamble placement and header with content", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  types.AttrTableOfContents,
							Value: "preamble",
						},
					},
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				section1,
			},
		}
		close(c)
		expected := &types.Document{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "toc",
							Value: "preamble",
						},
					},
				},
				&types.Preamble{
					Elements: []interface{}{
						paragraph,
					},
				},
				section1,
			},
			ElementReferences: types.ElementReferences{
				"_Section_1": section1Title,
			},
			TableOfContents: &types.TableOfContents{
				MaxDepth: 2,
				Sections: []*types.ToCSection{
					{
						ID:    "_Section_1",
						Level: 1,
					},
				},
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})

	It("with preamble placement and header without section1", func() {
		ctx := parser.NewParseContext(configuration.NewConfiguration())
		c := make(chan types.DocumentFragment, 4)
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name:  types.AttrTableOfContents,
					Value: "preamble",
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Title: doctitle,
					Attributes: types.Attributes{
						types.AttrID: "_a_title",
					},
				},
			},
		}
		c <- types.DocumentFragment{
			Elements: []interface{}{
				paragraph,
			},
		}
		close(c)
		expected := &types.Document{ // no ToC since no header, so no preamble either
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name:  "toc",
					Value: "preamble",
				},
				&types.DocumentHeader{
					Title: doctitle,
					Attributes: types.Attributes{
						types.AttrID: "_a_title",
					},
				},
				paragraph,
			},
		}
		doc, err := parser.Aggregate(ctx, c)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc).To(MatchDocument(expected))
	})
})

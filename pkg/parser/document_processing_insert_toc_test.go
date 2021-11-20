package parser

// . "github.com/onsi/ginkgo" //nolint golint
// . "github.com/onsi/gomega" //nolint golint

// var _ = Describe("insert table of contents", func() {

// var header, sectionA, sectionB, paragraph, anotherParagraph, blankline, preamble interface{}
// BeforeEach(func() {

// 	header = &types.DocumentHeader{
// 		Title: []interface{}{
// 			&types.StringElement{
// 				Content: "title",
// 			},
// 		},
// 		Elements: []interface{}{
// 			&types.AttributeDeclaration{
// 				Name:  "biscuits",
// 				Value: "cookies",
// 			},
// 		},
// 	}
// 	sectionA = &types.Section{
// 		Level: 1,
// 		Title: []interface{}{
// 			&types.StringElement{Content: "Section A"},
// 		},
// 		Attributes: types.Attributes{
// 			types.AttrID: "_section_a",
// 		},
// 	}
// 	sectionB = &types.Section{
// 		Level: 1,
// 		Title: []interface{}{
// 			&types.StringElement{Content: "Section B"},
// 		},
// 		Attributes: types.Attributes{
// 			types.AttrID: "_section_b",
// 		},
// 	}
// 	paragraph = &types.Paragraph{
// 		Elements: []interface{}{
// 			&types.StringElement{
// 				Content: "a short paragraph",
// 			},
// 		},
// 	}
// 	anotherParagraph = &types.Paragraph{
// 		Elements: []interface{}{
// 			&types.StringElement{
// 				Content: "another short paragraph",
// 			},
// 		},
// 	}
// 	blankline = &types.BlankLine{}
// 	preamble = &types.Preamble{
// 		Elements: []interface{}{
// 			paragraph,
// 			blankline,
// 		},
// 	}
// })
// toc := &types.TableOfContents{
// 	Sections: []*types.ToCSection{
// 		{
// 			ID:    "_section_a",
// 			Level: 1,
// 			Title: "Section A",
// 		},
// 		{
// 			ID:    "_section_b",
// 			Level: 1,
// 			Title: "Section B",
// 		},
// 	},
// }

// Context("no insertion", func() {

// 	It("should not insert when no placement attribute", func() {
// 		// given
// 		doc := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				sectionA,
// 				sectionB,
// 			},
// 		}
// 		expected := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				sectionA,
// 				sectionB,
// 			},
// 		}
// 		// when
// 		ctx := NewParseContext(configuration.NewConfiguration()) // no `:toc:` attribute declaration
// 		insertTableOfContents(ctx, doc, toc)

// 		// then
// 		Expect(doc).To(Equal(expected))
// 	})

// 	It("should not insert when no toc generated", func() {
// 		// given
// 		doc := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				paragraph,
// 				blankline,
// 				anotherParagraph,
// 			},
// 		}

// 		expected := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				paragraph,
// 				blankline,
// 				anotherParagraph,
// 			},
// 		}
// 		// when
// 		ctx := NewParseContext(configuration.NewConfiguration())
// 		ctx.attributes[types.AttrTableOfContents] = nil // default placement
// 		insertTableOfContents(ctx, doc, nil)            // no ToC since no sections in doc

// 		// then
// 		Expect(doc).To(Equal(expected))
// 	})

// 	It("should not insert when missing preamble", func() {
// 		// given
// 		doc := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				sectionA,
// 				sectionB,
// 			},
// 		}
// 		expected := &types.Document{
// 			Elements: []interface{}{
// 				header,
// 				sectionA,
// 				sectionB,
// 			},
// 		}
// 		// when
// 		ctx := NewParseContext(configuration.NewConfiguration())
// 		ctx.attributes[types.AttrTableOfContents] = "preamble"
// 		insertTableOfContents(ctx, doc, toc)

// 		// then
// 		Expect(doc).To(Equal(expected))
// 		// TODO: also check that there was a warning in the logs
// 	})
// })

// It("should insert within header", func() {
// 	// given
// 	doc := &types.Document{
// 		Elements: []interface{}{
// 			header,
// 			preamble,
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	expected := &types.Document{
// 		Elements: []interface{}{
// 			&types.DocumentHeader{
// 				Title: []interface{}{
// 					&types.StringElement{
// 						Content: "title",
// 					},
// 				},
// 				Elements: []interface{}{
// 					&types.AttributeDeclaration{
// 						Name:  "biscuits",
// 						Value: "cookies",
// 					},
// 					toc, // inserted here
// 				},
// 			},
// 			preamble,
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	// when
// 	ctx := NewParseContext(configuration.NewConfiguration())
// 	ctx.attributes[types.AttrTableOfContents] = nil // default placement (within header)
// 	insertTableOfContents(ctx, doc, toc)

// 	// then
// 	Expect(doc).To(Equal(expected))
// })

// It("should insert at first position when no header", func() {
// 	// given
// 	doc := &types.Document{
// 		Elements: []interface{}{
// 			paragraph,
// 			anotherParagraph,
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	expected := &types.Document{
// 		Elements: []interface{}{
// 			toc, // inserted here
// 			paragraph,
// 			anotherParagraph,
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	// when
// 	ctx := NewParseContext(configuration.NewConfiguration())
// 	ctx.attributes[types.AttrTableOfContents] = nil // default placement (within header)
// 	insertTableOfContents(ctx, doc, toc)

// 	// then
// 	Expect(doc).To(Equal(expected))
// })

// It("should insert within preamble", func() {
// 	// given
// 	doc := &types.Document{
// 		Elements: []interface{}{
// 			header,
// 			preamble,
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	expected := &types.Document{
// 		Elements: []interface{}{
// 			header,
// 			&types.Preamble{
// 				Elements: []interface{}{
// 					paragraph,
// 					blankline,
// 					toc, // appended here
// 				},
// 			},
// 			sectionA,
// 			sectionB,
// 		},
// 	}
// 	// when
// 	ctx := NewParseContext(configuration.NewConfiguration())
// 	ctx.attributes[types.AttrTableOfContents] = "preamble"
// 	insertTableOfContents(ctx, doc, toc)

// 	// then
// 	Expect(doc).To(Equal(expected))
// })

// })

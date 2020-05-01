package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("rearrange sections", func() {

	It("section levels 1, 2, 3, 2", func() {
		// = a header
		//
		// == Section A
		// a paragraph
		//
		// === Section A.a
		// a paragraph
		//
		// == Section B
		// a paragraph
		doctitle := []interface{}{
			types.StringElement{Content: "a header"},
		}
		sectionATitle := []interface{}{
			types.StringElement{Content: "Section A"},
		}
		sectionAaTitle := []interface{}{
			types.StringElement{Content: "Section A.a"},
		}
		sectionBTitle := []interface{}{
			types.StringElement{Content: "Section B"},
		}
		actual := []interface{}{
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Level:    0,
				Title:    doctitle,
				Elements: []interface{}{},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a",
				},
				Level: 1,
				Title: sectionATitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_a",
				},
				Level: 2,
				Title: sectionAaTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_b",
				},
				Level: 1,
				Title: sectionBTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"_a_header":    doctitle,
				"_section_a":   sectionATitle,
				"_section_a_a": sectionAaTitle,
				"_section_b":   sectionBTitle,
			},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_header",
					},
					Level: 0,
					Title: doctitle,
					Elements: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_a",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
							},
						},
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_b",
							},
							Level: 1,
							Title: sectionBTitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		result := rearrangeSections(actual)
		Expect(result).To(Equal(expected))
	})

	It("section levels 1, 2, 3, 3", func() {
		// = a header
		//
		// == Section A
		// a paragraph
		//
		// === Section A.a
		// a paragraph
		//
		// === Section A.b
		// a paragraph
		doctitle := []interface{}{
			types.StringElement{Content: "a header"},
		}
		sectionATitle := []interface{}{
			types.StringElement{Content: "Section A"},
		}
		sectionAaTitle := []interface{}{
			types.StringElement{Content: "Section A.a"},
		}
		sectionBTitle := []interface{}{
			types.StringElement{Content: "Section A.b"},
		}
		actual := []interface{}{
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Level:    0,
				Title:    doctitle,
				Elements: []interface{}{},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a",
				},
				Level: 1,
				Title: sectionATitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_a",
				},
				Level: 2,
				Title: sectionAaTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_b",
				},
				Level: 2,
				Title: sectionBTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"_a_header":    doctitle,
				"_section_a":   sectionATitle,
				"_section_a_a": sectionAaTitle,
				"_section_a_b": sectionBTitle,
			},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_header",
					},
					Level: 0,
					Title: doctitle,
					Elements: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_a",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_b",
									},
									Level: 2,
									Title: sectionBTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		result := rearrangeSections(actual)
		Expect(result).To(Equal(expected))
	})

	It("section levels 1, 3, 4, 4", func() {
		// = a header
		//
		// === Section A
		// a paragraph
		//
		// ==== Section A.a
		// a paragraph
		//
		// ==== Section A.b
		// a paragraph
		doctitle := []interface{}{
			types.StringElement{Content: "a header"},
		}
		sectionATitle := []interface{}{
			types.StringElement{Content: "Section A"},
		}
		sectionAaTitle := []interface{}{
			types.StringElement{Content: "Section A.a"},
		}
		sectionBTitle := []interface{}{
			types.StringElement{Content: "Section A.b"},
		}
		actual := []interface{}{
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Level:    0,
				Title:    doctitle,
				Elements: []interface{}{},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a",
				},
				Level: 2,
				Title: sectionATitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_a",
				},
				Level: 3,
				Title: sectionAaTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
			types.Section{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_b",
				},
				Level: 3,
				Title: sectionBTitle,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"_a_header":    doctitle,
				"_section_a":   sectionATitle,
				"_section_a_a": sectionAaTitle,
				"_section_a_b": sectionBTitle,
			},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_header",
					},
					Level: 0,
					Title: doctitle,
					Elements: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_section_a",
							},
							Level: 2,
							Title: sectionATitle,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_a",
									},
									Level: 3,
									Title: sectionAaTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_section_a_b",
									},
									Level: 3,
									Title: sectionBTitle,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		result := rearrangeSections(actual)
		Expect(result).To(Equal(expected))
	})

})

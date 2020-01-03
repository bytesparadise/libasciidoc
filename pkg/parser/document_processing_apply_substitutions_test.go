package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document attribute subsititutions", func() {

	It("should replace with new StringElement on first position", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.DocumentAttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
			"foo": "bar",
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeTrue())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "bar and more content.",
						},
					},
				},
			},
		}))
	})

	It("should replace with new StringElement on first position", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.DocumentAttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
			"foo": "bar",
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeTrue())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "bar and more content.",
						},
					},
				},
			},
		}))
	})

	It("should replace with new StringElement on middle position", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, ",
						},
						types.DocumentAttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
			"foo": "bar",
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeTrue())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, bar and more content.",
						},
					},
				},
			},
		}))
	})

	It("should replace with undefined attribute", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, ",
						},
						types.DocumentAttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeFalse())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, {foo} and more content.",
						},
					},
				},
			},
		}))
	})

	It("should merge without substitution", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, ",
						},
						types.StringElement{
							Content: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeFalse())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "baz, foo and more content.",
						},
					},
				},
			},
		}))
	})

	It("should replace with new link", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "a link to ",
						},
						types.DocumentAttributeSubstitution{
							Name: "scheme",
						},
						types.StringElement{
							Content: "://",
						},
						types.DocumentAttributeSubstitution{
							Name: "host",
						},
						types.StringElement{
							Content: "[].", // explicit use of `[]` to avoid grabbing the `.`
						},
					},
				},
			},
		}
		// when
		result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
			"scheme": "https",
			"host":   "foo.bar",
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(applied).To(BeTrue())
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "a link to ",
						},
						types.InlineLink{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "https://foo.bar",
									},
								},
							},
						},
						types.StringElement{
							Content: ".",
						},
					},
				},
			},
		}))
	})

	Context("list items", func() {

		It("should replace with new StringElement in ordered list item", func() {
			// given
			elements := []interface{}{
				types.OrderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.DocumentAttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.OrderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "bar and more content.",
									},
								},
							},
						},
					},
				},
			}))
		})

		It("should replace with new StringElement in unordered list item", func() {
			// given
			elements := []interface{}{
				types.UnorderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.DocumentAttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.UnorderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "bar and more content.",
									},
								},
							},
						},
					},
				},
			}))
		})

		It("should replace with new StringElement in labeled list item", func() {
			// given
			elements := []interface{}{
				types.LabeledListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.DocumentAttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.LabeledListItem{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "bar and more content.",
									},
								},
							},
						},
					},
				},
			}))
		})
	})

	Context("delimited blocks", func() {

		It("should replace with new StringElement in delimited block", func() {
			// given
			elements := []interface{}{
				types.DelimitedBlock{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.DocumentAttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.DelimitedBlock{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "bar and more content.",
									},
								},
							},
						},
					},
				},
			}))
		})
	})

	Context("quoted text", func() {

		It("should replace with new StringElement in quoted text", func() {
			// given
			elements := []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "hello ",
							},
							types.QuotedText{
								Elements: []interface{}{
									types.DocumentAttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
						{
							types.StringElement{
								Content: "and another line",
							},
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "hello ",
							},
							types.QuotedText{
								Elements: []interface{}{
									types.StringElement{
										Content: "bar and more content.",
									},
								},
							},
						},
						{
							types.StringElement{
								Content: "and another line",
							},
						},
					},
				},
			}))
		})
	})

	Context("tables", func() {

		It("should replace with new StringElement in table cell", func() {
			// given
			elements := []interface{}{
				types.DelimitedBlock{
					Elements: []interface{}{
						types.DocumentAttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			}
			// when
			result, applied, err := applyDocumentAttributeSubstitutions(elements, types.DocumentAttributes{
				"foo": "bar",
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(applied).To(BeTrue())
			Expect(result).To(Equal([]interface{}{
				types.DelimitedBlock{
					Elements: []interface{}{
						types.StringElement{
							Content: "bar and more content.",
						},
					},
				},
			}))
		})
	})
})

package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document attribute subsititutions", func() {

	It("should replace with new StringElement on first position", func() {
		// given
		elements := []interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
						types.AttributeSubstitution{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"foo": "bar",
			},
			Overrides: map[string]string{},
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
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
				Lines: []interface{}{
					[]interface{}{
						types.AttributeSubstitution{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"foo": "bar",
			},
			Overrides: map[string]string{},
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
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
				Lines: []interface{}{
					[]interface{}{
						types.StringElement{
							Content: "baz, ",
						},
						types.AttributeSubstitution{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"foo": "bar",
			},
			Overrides: map[string]string{},
		})
		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
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
				Lines: []interface{}{
					[]interface{}{
						types.StringElement{
							Content: "baz, ",
						},
						types.AttributeSubstitution{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content:   map[string]interface{}{},
			Overrides: map[string]string{},
		})

		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
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
				Lines: []interface{}{
					[]interface{}{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content:   map[string]interface{}{},
			Overrides: map[string]string{},
		})

		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
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
				Lines: []interface{}{
					[]interface{}{
						types.StringElement{
							Content: "a link to ",
						},
						types.AttributeSubstitution{
							Name: "scheme",
						},
						types.StringElement{
							Content: "://",
						},
						types.AttributeSubstitution{
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
		result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"foo":    "bar",
				"scheme": "https",
				"host":   "foo.bar",
			},
			Overrides: map[string]string{},
		})

		// then
		Expect(err).To(Not(HaveOccurred()))
		Expect(result).To(Equal([]interface{}{
			types.Paragraph{
				Lines: []interface{}{
					[]interface{}{
						types.StringElement{
							Content: "a link to ",
						},
						types.InlineLink{
							Location: types.Location{
								Scheme: "https://",
								Path: []interface{}{
									types.StringElement{
										Content: "foo.bar",
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
							Lines: []interface{}{
								[]interface{}{types.AttributeSubstitution{
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
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{
				types.OrderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{types.StringElement{
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
							Lines: []interface{}{
								[]interface{}{types.AttributeSubstitution{
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
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{
				types.UnorderedListItem{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{types.StringElement{
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
							Lines: []interface{}{
								[]interface{}{types.AttributeSubstitution{
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
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{
				types.LabeledListItem{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{types.StringElement{
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
							Lines: []interface{}{
								[]interface{}{types.AttributeSubstitution{
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
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{
				types.DelimitedBlock{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{types.StringElement{
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
					Lines: []interface{}{
						[]interface{}{
							types.StringElement{
								Content: "hello ",
							},
							types.QuotedText{
								Elements: []interface{}{
									types.AttributeSubstitution{
										Name: "foo",
									},
									types.StringElement{
										Content: " and more content.",
									},
								},
							},
						},
						[]interface{}{
							types.StringElement{
								Content: "and another line",
							},
						},
					},
				},
			}
			// when
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{
				types.Paragraph{
					Lines: []interface{}{
						[]interface{}{
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
						[]interface{}{
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
						types.AttributeSubstitution{
							Name: "foo",
						},
						types.StringElement{
							Content: " and more content.",
						},
					},
				},
			}
			// when
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
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

	Context("attribute overrides", func() {

		It("should replace with new StringElement on first position", func() {
			// given
			elements := []interface{}{
				types.AttributeDeclaration{
					Name:  "foo",
					Value: "foo",
				},
				types.AttributeReset{
					Name: "foo",
				},
				types.Paragraph{
					Lines: []interface{}{
						[]interface{}{
							types.AttributeSubstitution{
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
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content: map[string]interface{}{
					"foo": "bar",
				},
				Overrides: map[string]string{
					"foo": "BAR",
				},
			})
			// then
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
				types.AttributeDeclaration{
					Name:  "foo",
					Value: "foo",
				},
				types.AttributeReset{
					Name: "foo",
				},
				types.Paragraph{
					Lines: []interface{}{
						[]interface{}{
							types.StringElement{
								Content: "BAR and more content.",
							},
						},
					},
				},
			}))
		})
	})

	Context("counters", func() {

		It("should start at one", func() {
			// given
			elements := []interface{}{
				types.CounterSubstitution{
					Name: "foo",
				},
			}
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content:   map[string]interface{}{},
				Overrides: map[string]string{},
				Counters:  map[string]interface{}{},
			})
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
				types.StringElement{
					Content: "1",
				},
			}))
		})

		It("should increment correctly", func() {
			// given
			elements := []interface{}{
				types.CounterSubstitution{
					Name: "foo",
				},
				types.CounterSubstitution{
					Name: "bar",
				},
				types.CounterSubstitution{
					Name: "foo",
				},
				types.CounterSubstitution{
					Name:   "alpha",
					Value:  'a',
					Hidden: true,
				},
				types.CounterSubstitution{
					Name: "alpha",
				},
				types.CounterSubstitution{
					Name:   "set",
					Value:  33,
					Hidden: true,
				},
				types.CounterSubstitution{
					Name:   "set",
					Hidden: true,
				},
				types.CounterSubstitution{
					Name: "set",
				},
			}
			result, err := applyAttributeSubstitutions(elements, types.AttributesWithOverrides{
				Content:   map[string]interface{}{},
				Overrides: map[string]string{},
				Counters:  map[string]interface{}{},
			})
			Expect(err).To(Not(HaveOccurred()))
			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
				types.StringElement{
					Content: "112b35", // elements get concatenated
				},
			}))
		})
	})
})

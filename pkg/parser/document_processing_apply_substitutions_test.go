package parser

// . "github.com/onsi/ginkgo" //nolint golint
// . "github.com/onsi/gomega" //nolint golint

// var _ = Describe("document substitutions", func() {

// 	Context("paragraphs", func() {

// 		It("should replace with new StringElement on first position", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							types.AttributeSubstitution{
// 								Name: "foo",
// 							},
// 							&types.StringElement{
// 								Content: " and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "bar and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace with new StringElement on middle position", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, ",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "foo",
// 							},
// 							&types.StringElement{
// 								Content: " and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, bar and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace with undefined attribute", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, ",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "foo",
// 							},
// 							&types.StringElement{
// 								Content: " and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content:   map[string]interface{}{},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)

// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, {foo} and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should merge without substitution", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, ",
// 							},
// 							&types.StringElement{
// 								Content: "foo",
// 							},
// 							&types.StringElement{
// 								Content: " and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content:   map[string]interface{}{},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)

// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "baz, foo and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace with new link", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "a link to ",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "scheme",
// 							},
// 							&types.StringElement{
// 								Content: "://",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "host",
// 							},
// 							&types.StringElement{
// 								Content: "[].", // explicit use of `[]` to avoid grabbing the `.`
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo":    "bar",
// 						"scheme": "https",
// 						"host":   "example.com",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "a link to ",
// 							},
// 							&types.InlineLink{
// 								Location: &types.Location{
// 									Scheme: "https://",
// 									Path: []interface{}{
// 										&types.StringElement{
// 											Content: "example.com",
// 										},
// 									},
// 								},
// 							},
// 							&types.StringElement{
// 								Content: ".",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should substitute title attribute", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						types.AttrTitle: []interface{}{
// 							&types.StringElement{
// 								Content: "a ",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "title",
// 							},
// 						},
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "a paragraph with title '",
// 							},
// 							types.AttributeSubstitution{
// 								Name: "title",
// 							},
// 							&types.StringElement{
// 								Content: "'",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"title": "cookie",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						types.AttrTitle: "a cookie",
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "a paragraph with title 'cookie'",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})
// 	})

// 	Context("paragraph with attributes", func() {

// 		It("should replace title attribute", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						"title": []interface{}{
// 							types.AttributeSubstitution{
// 								Name: "title",
// 							},
// 						},
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "some content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"title": "TITLE",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						"title": "TITLE",
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "some content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace roles attribute", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						types.AttrRoles: []interface{}{
// 							[]interface{}{
// 								types.AttributeSubstitution{
// 									Name: "role1",
// 								},
// 							},
// 							[]interface{}{
// 								types.AttributeSubstitution{
// 									Name: "role2",
// 								},
// 							},
// 						},
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "some content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"role1": "ROLE1",
// 						"role2": "ROLE2",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Attributes: types.Attributes{
// 						types.AttrRoles: []interface{}{"ROLE1", "ROLE2"},
// 					},
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "some content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 	})

// 	Context("image blocks", func() {

// 		It("should substitute inline attribute", func() {
// 			elements := []interface{}{
// 				&types.ImageBlock{
// 					Attributes: types.Attributes{
// 						types.AttrImageAlt: []interface{}{
// 							types.AttributeSubstitution{
// 								Name: "alt",
// 							},
// 						},
// 					},
// 					Location: &types.Location{
// 						Path: []interface{}{
// 							&types.StringElement{Content: "foo.png"},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"alt": "cookie",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				&types.ImageBlock{
// 					Attributes: types.Attributes{
// 						types.AttrImageAlt: "cookie", // substituted
// 					},
// 					Location: &types.Location{
// 						Path: []interface{}{
// 							&types.StringElement{Content: "foo.png"},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should substitute inline attribute", func() {
// 			elements := []interface{}{
// 				&types.ImageBlock{
// 					Location: &types.Location{
// 						Path: []interface{}{
// 							types.AttributeSubstitution{
// 								Name: "path",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"path": "cookie.png",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				&types.ImageBlock{
// 					Location: &types.Location{
// 						Path: []interface{}{
// 							&types.StringElement{Content: "cookie.png"},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 	})

// 	Context("list items", func() {

// 		It("should replace with new StringElement in ordered list item", func() {
// 			// given
// 			elements := []interface{}{
// 				types.OrderedListElement{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									types.AttributeSubstitution{
// 										Name: "foo",
// 									},
// 									&types.StringElement{
// 										Content: " and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.OrderedListElement{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									&types.StringElement{
// 										Content: "bar and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace with new StringElement in unordered list item", func() {
// 			// given
// 			elements := []interface{}{
// 				types.UnorderedListItem{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									types.AttributeSubstitution{
// 										Name: "foo",
// 									},
// 									&types.StringElement{
// 										Content: " and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.UnorderedListItem{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									&types.StringElement{
// 										Content: "bar and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 		It("should replace with new StringElement in labeled list item", func() {
// 			// given
// 			elements := []interface{}{
// 				types.LabeledListItem{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									types.AttributeSubstitution{
// 										Name: "foo",
// 									},
// 									&types.StringElement{
// 										Content: " and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.LabeledListItem{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									&types.StringElement{
// 										Content: "bar and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})
// 	})

// 	Context("as delimited blocks", func() {

// 		It("should replace with new StringElement in delimited block", func() {
// 			// given
// 			elements := []interface{}{
// 				types.ExampleBlock{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									types.AttributeSubstitution{
// 										Name: "foo",
// 									},
// 									&types.StringElement{
// 										Content: " and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.ExampleBlock{
// 					Elements: []interface{}{
// 						types.Paragraph{
// 							Lines: [][]interface{}{
// 								{
// 									&types.StringElement{
// 										Content: "bar and more content.",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})
// 	})

// 	Context("quoted texts", func() {

// 		It("should replace with new StringElement in quoted text", func() {
// 			// given
// 			elements := []interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "hello ",
// 							},
// 							&types.QuotedText{
// 								Elements: []interface{}{
// 									types.AttributeSubstitution{
// 										Name: "foo",
// 									},
// 									&types.StringElement{
// 										Content: " and more content.",
// 									},
// 								},
// 							},
// 						},
// 						{
// 							&types.StringElement{
// 								Content: "and another line",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "hello ",
// 							},
// 							&types.QuotedText{
// 								Elements: []interface{}{
// 									&types.StringElement{
// 										Content: "bar and more content.",
// 									},
// 								},
// 							},
// 						},
// 						{
// 							&types.StringElement{
// 								Content: "and another line",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})
// 	})

// 	Context("attribute overrides", func() {

// 		It("should replace with new StringElement on first position", func() {
// 			// given
// 			elements := []interface{}{
// 				&types.AttributeDeclaration{
// 					Name:  "foo",
// 					Value: "foo",
// 				},
// 				&types.AttributeReset{
// 					Name: "foo",
// 				},
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							types.AttributeSubstitution{
// 								Name: "foo",
// 							},
// 							&types.StringElement{
// 								Content: " and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content: map[string]interface{}{
// 						"foo": "bar",
// 					},
// 					Overrides: map[string]string{
// 						"foo": "BAR",
// 					},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			// then
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
// 				&types.AttributeDeclaration{
// 					Name:  "foo",
// 					Value: "foo",
// 				},
// 				&types.AttributeReset{
// 					Name: "foo",
// 				},
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "BAR and more content.",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})
// 	})

// 	Context("counters", func() {

// 		It("should start at one", func() {
// 			// given
// 			elements := []interface{}{
// 				types.CounterSubstitution{
// 					Name: "foo",
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content:   map[string]interface{}{},
// 					Overrides: map[string]string{},
// 					Counters:  map[string]interface{}{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
// 				&types.StringElement{
// 					Content: "1",
// 				},
// 			}))
// 		})

// 		It("should increment correctly", func() {
// 			// given
// 			elements := []interface{}{
// 				types.CounterSubstitution{
// 					Name: "foo",
// 				},
// 				types.CounterSubstitution{
// 					Name: "bar",
// 				},
// 				types.CounterSubstitution{
// 					Name: "foo",
// 				},
// 				types.CounterSubstitution{
// 					Name:   "alpha",
// 					Value:  'a',
// 					Hidden: true,
// 				},
// 				types.CounterSubstitution{
// 					Name: "alpha",
// 				},
// 				types.CounterSubstitution{
// 					Name:   "set",
// 					Value:  33,
// 					Hidden: true,
// 				},
// 				types.CounterSubstitution{
// 					Name:   "set",
// 					Hidden: true,
// 				},
// 				types.CounterSubstitution{
// 					Name: "set",
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content:   map[string]interface{}{},
// 					Overrides: map[string]string{},
// 					Counters:  map[string]interface{}{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{
// 				// elements are not concatenated after calling `applyAttributeSubstitutionsOnElements`
// 				&types.StringElement{Content: "1"},
// 				&types.StringElement{Content: "1"},
// 				&types.StringElement{Content: "2"},
// 				&types.StringElement{Content: ""},
// 				&types.StringElement{Content: "b"},
// 				&types.StringElement{Content: ""},
// 				&types.StringElement{Content: ""},
// 				&types.StringElement{Content: "35"},
// 			}))
// 		})
// 	})

// 	Context("recursive attributes", func() {

// 		It("should substitute an attribute in another attribute", func() {
// 			elements := []interface{}{
// 				&types.AttributeDeclaration{
// 					Name:  "def",
// 					Value: "foo",
// 				},
// 				&types.AttributeDeclaration{
// 					Name: "abc",
// 					Value: []interface{}{
// 						types.AttributeSubstitution{
// 							Name: "def",
// 						},
// 						&types.StringElement{
// 							Content: "bar",
// 						},
// 					},
// 				},
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							types.AttributeSubstitution{
// 								Name: "abc",
// 							},
// 						},
// 					},
// 				},
// 			}
// 			ctx := substitutionContextDeprecated{
// 				attributes: types.AttributesWithOverrides{
// 					Content:   map[string]interface{}{},
// 					Overrides: map[string]string{},
// 					Counters:  map[string]interface{}{},
// 				},
// 				config: configuration.NewConfiguration(),
// 			}
// 			// when
// 			result, err := applySubstitutions(ctx, elements)
// 			Expect(err).To(Not(HaveOccurred()))
// 			Expect(result).To(Equal([]interface{}{ // at this stage, AttributeDeclaration and AttributeReset are still present
// 				&types.AttributeDeclaration{
// 					Name:  "def",
// 					Value: "foo",
// 				},
// 				&types.AttributeDeclaration{
// 					Name:  "abc",
// 					Value: "foobar",
// 				},
// 				types.Paragraph{
// 					Lines: [][]interface{}{
// 						{
// 							&types.StringElement{
// 								Content: "foobar",
// 							},
// 						},
// 					},
// 				},
// 			}))
// 		})

// 	})
// })

// var _ = Describe("substitution funcs", func() {

// 	It("should append sub", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes"}
// 		// when
// 		f = f.append("macros")
// 		// then
// 		Expect(f).To(Equal(funcs{"attributes", "quotes", "macros"}))
// 	})

// 	It("should append subs", func() {
// 		// given"
// 		f := funcs{"attributes"}
// 		// when
// 		f = f.append("quotes", "macros")
// 		// then
// 		Expect(f).To(Equal(funcs{"attributes", "quotes", "macros"}))
// 	})

// 	It("should prepend sub", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes"}
// 		// when
// 		f = f.prepend("macros")
// 		// then
// 		Expect(f).To(Equal(funcs{"macros", "attributes", "quotes"}))
// 	})

// 	It("should remove first sub", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes", "macros"}
// 		// when
// 		f = f.remove("attributes")
// 		// then
// 		Expect(f).To(Equal(funcs{"quotes", "macros"}))
// 	})

// 	It("should remove middle sub", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes", "macros"}
// 		// when
// 		f = f.remove("quotes")
// 		// then
// 		Expect(f).To(Equal(funcs{"attributes", "macros"}))
// 	})

// 	It("should remove last sub", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes", "macros"}
// 		// when
// 		f = f.remove("macros")
// 		// then
// 		Expect(f).To(Equal(funcs{"attributes", "quotes"}))
// 	})

// 	It("should remove non existinge", func() {
// 		// given"
// 		f := funcs{"attributes", "quotes", "macros"}
// 		// when
// 		f = f.remove("other")
// 		// then
// 		Expect(f).To(Equal(funcs{"attributes", "quotes", "macros"}))
// 	})
// })

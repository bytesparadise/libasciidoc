package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("lists", func() {

	Context("unordered list", func() {

		It("multi-level list", func() {
			// // given
			elements := []interface{}{
				types.UnorderedListItem{
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1",
						},
					},
				},
				types.UnorderedListItem{
					Level:       2,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1.1",
						},
					},
				},
				types.UnorderedListItem{
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 2",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements, nil)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.StringElement{
								Content: "item 1",
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.StringElement{
												Content: "item 1.1",
											},
										},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.StringElement{
								Content: "item 2",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

	})

	Context("labeled list", func() {
		It("labeled list with 3 items", func() {
			// // given
			elements := []interface{}{
				types.LabeledListItem{
					Term: "item 1",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1",
						},
					},
				},
				types.LabeledListItem{
					Term: "item 2",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 2",
						},
					},
				},
				types.LabeledListItem{
					Term: "item 3",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 3",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements, nil)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: map[string]interface{}{},
				Items: []types.LabeledListItem{
					{
						Term: "item 1",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 1",
							},
						},
					},
					{
						Term: "item 2",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 2",
							},
						},
					},
					{
						Term: "item 3",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 3",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})
	})

	Context("mixed lists", func() {
		It("labeled list with unordered sublist", func() {
			// // given
			elements := []interface{}{
				types.LabeledListItem{
					Term: "item A",
					Elements: []interface{}{
						types.StringElement{
							Content: "item A",
						},
					},
				},
				types.UnorderedListItem{
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1",
						},
					},
				},
				types.UnorderedListItem{
					Level:       2,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1.1",
						},
					},
				},
				types.UnorderedListItem{
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.2",
						},
					},
				},
				types.LabeledListItem{
					Term: "item B",
					Elements: []interface{}{
						types.StringElement{
							Content: "item B",
						},
					},
				},
				types.LabeledListItem{
					Term: "item C",
					Elements: []interface{}{
						types.StringElement{
							Content: "item C",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements, nil)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: map[string]interface{}{},
				Items: []types.LabeledListItem{
					{
						Term: "item A",
						Elements: []interface{}{
							types.StringElement{
								Content: "item A",
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       1,
										BulletStyle: types.Dash,
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.1",
											},
											types.UnorderedList{
												Attributes: map[string]interface{}{},
												Items: []types.UnorderedListItem{
													{
														Level:       2,
														BulletStyle: types.OneAsterisk,
														Elements: []interface{}{
															types.StringElement{
																Content: "item A.1.1",
															},
														},
													},
												},
											},
										},
									},
									{
										Level:       1,
										BulletStyle: types.Dash,
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.2",
											},
										},
									},
								},
							},
						},
					},
					{
						Term: "item B",
						Elements: []interface{}{
							types.StringElement{
								Content: "item B",
							},
						},
					},
					{
						Term: "item C",
						Elements: []interface{}{
							types.StringElement{
								Content: "item C",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})
	})
})

func verify(t GinkgoTInterface, expectation, actual interface{}) {
	t.Logf("actual document: `%s`", spew.Sdump(actual))
	t.Logf("expected document: `%s`", spew.Sdump(expectation))
	assert.EqualValues(t, expectation, actual)
}

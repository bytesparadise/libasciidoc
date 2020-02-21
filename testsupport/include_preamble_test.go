package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document preamble assertions", func() {

	expected := types.Document{
		Attributes:         types.DocumentAttributes{},
		ElementReferences:  types.ElementReferences{},
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			types.Preamble{
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a short paragraph"},
							},
						},
					},
					types.BlankLine{},
				},
			},
			types.Section{
				Level: 1,
				Title: []interface{}{
					types.StringElement{Content: "Section A"},
				},
				Attributes: types.ElementAttributes{},
				Elements:   []interface{}{},
			},
			types.Section{
				Level: 1,
				Title: []interface{}{
					types.StringElement{Content: "Section B"},
				},
				Attributes: types.ElementAttributes{},
				Elements:   []interface{}{},
			},
		},
	}

	It("should match", func() {
		// given
		actual := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level: 1,
					Title: []interface{}{
						types.StringElement{Content: "Section A"},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: []interface{}{
						types.StringElement{Content: "Section B"},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				},
			},
		}
		// when
		result := testsupport.IncludePreamble(actual)
		// then
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := types.Document{}
		// when
		result := testsupport.IncludePreamble(actual)
		// then
		Expect(result).NotTo(Equal(expected))
	})

})

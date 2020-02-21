package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document metadata assertions", func() {

	expected := types.DocumentAttributes{
		"foo": "bar",
	}

	It("should match", func() {
		// given
		actual := types.Document{
			Attributes: types.DocumentAttributes{
				"foo": "bar",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Attributes: types.ElementAttributes{},
					Title:      []interface{}{},
					Elements:   []interface{}{},
				},
			},
		}
		// when
		result := testsupport.DocumentAttributes(actual)
		// then
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := types.Document{}
		// when
		result := testsupport.DocumentAttributes(actual)
		// then
		Expect(result).NotTo(Equal(expected))
	})

})

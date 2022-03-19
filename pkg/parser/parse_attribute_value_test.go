package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("parse attribute value", func() {

	It("should parse content with special characters", func() {
		source := "<h3>Table of Contents</h3>"
		expected := []interface{}{
			&types.SpecialCharacter{
				Name: "<",
			},
			&types.StringElement{
				Content: "h3",
			},
			&types.SpecialCharacter{
				Name: ">",
			},
			&types.StringElement{
				Content: "Table of Contents",
			},
			&types.SpecialCharacter{
				Name: "<",
			},
			&types.StringElement{
				Content: "/h3",
			},
			&types.SpecialCharacter{
				Name: ">",
			},
		}
		Expect(parser.ParseAttributeValue(source)).To(Equal(expected))
	})

	It("should parse content within inline passthrough", func() {
		source := "pass:[<h3>Table of Contents</h3>]"
		expected := []interface{}{
			&types.InlinePassthrough{
				Kind: types.PassthroughMacro,
				Elements: []interface{}{
					&types.StringElement{
						Content: "<h3>Table of Contents</h3>",
					},
				},
			},
		}
		Expect(parser.ParseAttributeValue(source)).To(Equal(expected))
	})
})

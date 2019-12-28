package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("attribute subsititutions", func() {

	It("should replace with new StringElement on first position", func() {
		// given
		e := []interface{}{
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: " and more content.",
			},
		}
		// when
		result, err := parser.ApplyDocumentAttributeSubstitutions(e, types.DocumentAttributes{
			"foo": "bar",
		})
		// then
		Expect(result).To(Equal([]interface{}{
			types.StringElement{
				Content: "bar and more content.",
			},
		}))
		Expect(err).To(Not(HaveOccurred()))
	})

	It("should replace with new StringElement on middle position", func() {
		// given
		e := []interface{}{
			types.StringElement{
				Content: "baz, ",
			},
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: " and more content.",
			},
		}
		// when
		result, err := parser.ApplyDocumentAttributeSubstitutions(e, types.DocumentAttributes{
			"foo": "bar",
		})
		// then
		Expect(result).To(Equal([]interface{}{
			types.StringElement{
				Content: "baz, bar and more content.",
			},
		}))
		Expect(err).To(Not(HaveOccurred()))
	})

	It("should replace with undefined attribute", func() {
		// given
		e := []interface{}{
			types.StringElement{
				Content: "baz, ",
			},
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: " and more content.",
			},
		}
		// when
		result, err := parser.ApplyDocumentAttributeSubstitutions(e, types.DocumentAttributes{})
		// then
		Expect(result).To(Equal([]interface{}{
			types.StringElement{
				Content: "baz, {foo} and more content.",
			},
		}))
		Expect(err).To(Not(HaveOccurred()))
	})

	It("should merge without substitution", func() {
		// given
		e := []interface{}{
			types.StringElement{
				Content: "baz, ",
			},
			types.StringElement{
				Content: "foo",
			},
			types.StringElement{
				Content: " and more content.",
			},
		}
		// when
		result, err := parser.ApplyDocumentAttributeSubstitutions(e, types.DocumentAttributes{})
		// then
		Expect(result).To(Equal([]interface{}{
			types.StringElement{
				Content: "baz, foo and more content.",
			},
		}))
		Expect(err).To(Not(HaveOccurred()))
	})
})

package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("preprocessing condition stack", func() {

	It("should eval to false when pushing single disabled entry", func() {
		// given
		c := newConditions()
		ctx := NewParseContext(configuration.NewConfiguration())
		// when
		eval := c.push(ctx, &types.IfdefCondition{
			Name: "cookie",
		})
		// then
		Expect(eval).To(BeFalse())
	})

	It("should eval to true when pushing single enabled entry", func() {
		// given
		c := newConditions()
		ctx := NewParseContext(configuration.NewConfiguration(configuration.WithAttribute("cookie", "yummy")))
		// when
		eval := c.push(ctx, &types.IfdefCondition{
			Name: "cookie",
		})
		// then
		Expect(eval).To(BeTrue())
	})

	It("should update when pushing multiple entries", func() {
		// given
		c := newConditions()
		ctx := NewParseContext(configuration.NewConfiguration(
			configuration.WithAttribute("cookie", "yummy"),
			configuration.WithAttribute("chocolate", "dark"),
			configuration.WithAttribute("pasta", ""),
		))
		// when
		eval := c.push(ctx, &types.IfdefCondition{
			Name: "cookie",
		})
		// then
		Expect(eval).To(BeTrue())

		// when
		eval = c.push(ctx, &types.IfdefCondition{
			Name: "cookie",
		})
		// then
		Expect(eval).To(BeTrue())

		// when
		eval = c.push(ctx, &types.IfdefCondition{
			Name: "unknown",
		})
		// then switch to false when condition is evaled to `false`
		Expect(eval).To(BeFalse())

		// when
		eval = c.push(ctx, &types.IfdefCondition{
			Name: "cookie",
		})
		// then remains to `false` because of `unknown` condition
		Expect(eval).To(BeFalse())

		// when
		eval = c.pop()
		// then remains to `false` because of `unknown` condition
		Expect(eval).To(BeFalse())

		// when
		eval = c.pop()
		// then back to `true`
		Expect(eval).To(BeTrue())

	})
})

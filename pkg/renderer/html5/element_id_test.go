package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("element ID generation", func() {

	It("should generate ID with default prefix", func() {
		// given
		ctx := &renderer.Context{
			Document: types.Document{
				Attributes: types.DocumentAttributes{},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		Expect(result).To(Equal("_foo"))
	})

	It("should generate ID with custom prefix", func() {
		// given
		ctx := &renderer.Context{
			Document: types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		Expect(result).To(Equal("id#foo"))
	})

	It("should generate custom ID", func() {
		// given
		ctx := &renderer.Context{
			Document: types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "foo",
			types.AttrCustomID: true,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		Expect(result).To(Equal("foo"))
	})

	It("should generate empty ID from empty value", func() {
		// given
		ctx := &renderer.Context{
			Document: types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrID:       "",
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		Expect(result).To(Equal(""))
	})

	It("should generate empty ID from missing value", func() {
		// given
		ctx := &renderer.Context{
			Document: types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "id#",
				},
			},
		}
		attrs := types.ElementAttributes{
			types.AttrCustomID: false,
		}
		// when
		result := generateID(ctx, attrs)
		// then
		Expect(result).To(Equal(""))
	})
})

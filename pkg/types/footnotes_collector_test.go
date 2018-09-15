package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("footnotes collector", func() {

	It("index footnotes without reference", func() {
		// given
		footnote1 := types.Footnote{
			ID: 0,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "a note",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "another note",
				},
			},
		}
		content := types.InlineElements{
			types.StringElement{
				Content: "foo",
			},
			footnote1,
			types.StringElement{
				Content: "bar",
			},
			footnote2,
		}
		c := types.NewFootnotesCollector()
		// when
		content.Accept(c)
		// then
		require.Len(GinkgoT(), c.Footnotes, 2)
		assert.Equal(GinkgoT(), footnote1, c.Footnotes[0])
		assert.Equal(GinkgoT(), footnote2, c.Footnotes[1])
		require.Empty(GinkgoT(), c.FootnoteReferences)
	})

	It("index footnotes with reference", func() {
		// given
		footnote1 := types.Footnote{
			Ref: "ref",
			Elements: types.InlineElements{
				types.StringElement{
					Content: "a note",
				},
			},
		}
		footnote2 := types.Footnote{
			Ref: "ref",
		}
		footnote3 := types.Footnote{
			Ref: "ref",
		}
		content := types.InlineElements{
			types.StringElement{
				Content: "foo",
			},
			footnote1,
			types.StringElement{
				Content: "bar",
			},
			footnote2,
			footnote3,
		}
		c := types.NewFootnotesCollector()
		// when
		content.Accept(c)
		// then
		require.Len(GinkgoT(), c.Footnotes, 1) // a single, yet referenced twice elsewhere
		assert.Equal(GinkgoT(), footnote1, c.Footnotes[0])
		require.Len(GinkgoT(), c.FootnoteReferences, 1) // a single, yet referenced twice elsewhere
		assert.Equal(GinkgoT(), footnote1, c.FootnoteReferences["ref"])
	})

})

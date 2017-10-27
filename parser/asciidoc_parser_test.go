package parser_test

import (
	"strings"

	. "github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Parsing content", func() {

	It("header section inline with bold quote", func() {

		actualContent := "= a header\n" +
			"\n" +
			"== section 1\n" +
			"\n" +
			"a paragraph with *bold content*"
		expectedDocument := &types.Document{
			Attributes: map[string]interface{}{
				"doctitle": &types.SectionTitle{
					Content: &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a header"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_header",
					},
				},
			},
			Elements: []types.DocElement{
				&types.Section{
					Level: 1,
					SectionTitle: types.SectionTitle{
						Content: &types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "section 1"},
							},
						},
						ID: &types.ElementID{
							Value: "_section_1",
						},
					},
					Elements: []types.DocElement{
						&types.Paragraph{
							Lines: []*types.InlineContent{
								&types.InlineContent{
									Elements: []types.InlineElement{
										&types.StringElement{Content: "a paragraph with "},
										&types.QuotedText{Kind: types.Bold,
											Elements: []types.InlineElement{
												&types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

})

func verify(t GinkgoTInterface, expectedDocument interface{}, content string, options ...Option) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	result, err := ParseReader("", reader, options...) //, Debug(true))
	if err != nil {
		log.WithError(err).Error("Error found while parsing the document")
	}
	require.Nil(t, err)
	t.Logf("actual document: `%s`", spew.Sdump(result))
	t.Logf("expected document: `%s`", spew.Sdump(expectedDocument))
	assert.EqualValues(t, expectedDocument, result)
}

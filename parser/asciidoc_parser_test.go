package parser_test

import (
	"strings"

	. "github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Parsing content", func() {

	It("heading section inline with bold quote", func() {

		actualContent := "= a heading\n" +
			"\n" +
			"== section 1\n" +
			"\n" +
			"a paragraph with *bold content*"
		expectedDocument := &types.Document{
			Metadata: &types.DocumentMetadata{
				"title": "a heading",
			},
			Elements: []types.DocElement{
				&types.Section{
					Heading: types.Heading{
						Level: 1,
						Content: &types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "a heading"},
							},
						},
						ID: &types.ElementID{
							Value: "_a_heading",
						},
					},
					Elements: []types.DocElement{
						&types.Section{
							Heading: types.Heading{
								Level: 2,
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
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

})

func verify(t GinkgoTInterface, expectedDocument *types.Document, content string) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	result, err := ParseReader("", reader)
	if err != nil {
		log.WithError(err).Error("Error found while parsing the document")
	}
	require.Nil(t, err)
	actualDocument := result.(*types.Document)
	t.Logf("actual document structure: %+v", actualDocument.Elements)
	t.Logf("actual document: `%s`", actualDocument.String(0))
	t.Logf("expected document: `%s`", expectedDocument.String(0))
	assert.EqualValues(t, *expectedDocument, *actualDocument)
}

package parser_test

import (
	"reflect"
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
			Elements: []types.DocElement{
				&types.Heading{
					Level: 1,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a heading"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_heading",
					},
				},
				&types.BlankLine{},
				&types.Heading{
					Level: 2,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "section 1"},
						},
					},
					ID: &types.ElementID{
						Value: "_section_1",
					},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a paragraph with "},
								&types.QuotedText{Kind: types.Bold,
									Elements: []types.DocElement{
										&types.StringElement{Content: "bold content"},
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
	log.Debugf("processing:\n%s", content)
	reader := strings.NewReader(content)
	result, err := ParseReader("", reader)
	if err != nil {
		log.WithError(err).Error("Error found while parsing the document")
	}
	require.Nil(t, err)
	actualDocument := result.(*types.Document)
	t.Logf("actual document: %+v", reflect.TypeOf(actualDocument.Elements[0]))
	t.Logf("expected document: %+v", expectedDocument)
	assert.EqualValues(t, *expectedDocument, *actualDocument)
}

package parser_test

import (
	"fmt"
	"strings"

	. "github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("heading only", func() {

		actualContent := "= a heading"
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
			}}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("heading invalid1", func() {

		actualContent := "=a heading"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "=a heading"},
							},
						},
					},
				},
			}}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("heading invalid2", func() {

		actualContent := " = a heading"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: " = a heading"},
							},
						},
					},
				},
			}}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("section2", func() {

		actualContent := `== section 1`
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
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
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("heading with section2", func() {

		actualContent := "= a heading\n" +
			"\n" +
			"== section 1"
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
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("heading with invalid section2", func() {

		actualContent := "= a heading\n" +
			"\n" +
			" == section 1"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Heading{
					Level: 1, Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a heading"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_heading",
					},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: " == section 1"},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("inline1 word", func() {

		actualContent := "hello"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "hello"},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("inline simple", func() {

		actualContent := "a paragraph with some content"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a paragraph with some content"},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
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
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("single list item", func() {

		actualContent := "* a list item"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ListItem{
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a list item"},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("invalid list item", func() {

		actualContent := "*an invalid list item"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "*an invalid list item"},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("list items", func() {

		actualContent := "* a first item\n" +
			"* a second item with *bold content*"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ListItem{
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a first item"},
						},
					},
				},
				&types.ListItem{
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a second item with "},
							&types.QuotedText{Kind: types.Bold,
								Elements: []types.DocElement{
									&types.StringElement{Content: "bold content"},
								},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("external link", func() {

		actualContent := "a link to https://foo.bar"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL: "https://foo.bar",
								},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("external link with empty text", func() {

		actualContent := "a link to https://foo.bar[]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL:  "https://foo.bar",
									Text: "",
								},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("external link with text", func() {

		actualContent := "a link to mailto:foo@bar[the foo@bar email]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL:  "mailto:foo@bar",
									Text: "the foo@bar email",
								},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element link", func() {

		actualContent := "[link=http://foo.bar]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ElementLink{Path: "http://foo.bar"},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element link with spaces", func() {

		actualContent := "[ link = http://foo.bar ]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ElementLink{Path: "http://foo.bar"},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element link invalid", func() {

		actualContent := "[ link = http://foo.bar"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "[ link = "},
								&types.ExternalLink{URL: "http://foo.bar"},
							},
						},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element i d", func() {

		actualContent := "[#img-foobar]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ElementID{Value: "img-foobar"},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element i d with spaces", func() {

		actualContent := "[ #img-foobar ]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ElementID{Value: "img-foobar"},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element i d invalid", func() {

		actualContent := "[#img-foobar"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "[#img-foobar"}}},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element title", func() {

		actualContent := ".a title"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.ElementTitle{Content: "a title"},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element title invalid1", func() {

		actualContent := ". a title"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: ". a title"}}},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("element title invalid2", func() {

		actualContent := "!a title"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "!a title"}}},
					},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
})

func compare(t GinkgoTInterface, expectedDocument *types.Document, content string) {
	t.Log(fmt.Sprintf("processing:\n%s", content))
	reader := strings.NewReader(content)
	result, err := ParseReader("", reader)
	if err != nil {
		log.Errorf("Error found while parsing the document: %v", err.Error())
	}
	require.Nil(t, err)
	actualDocument := result.(*types.Document)
	t.Log(fmt.Sprintf("actual document: %s", actualDocument.String()))
	t.Log(fmt.Sprintf("expected document: %s", expectedDocument.String()))
	assert.EqualValues(t, expectedDocument, actualDocument)
}

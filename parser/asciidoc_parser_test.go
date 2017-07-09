package parser_test

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	. "github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	args := flag.Args()
	if len(args) > 0 {
		log.Warnf("Starting test(s) with args=%v", flag.Args())
	} else {
		log.Warn("Starting test(s) with no custom arguments")
	}
}

func compare(t *testing.T, expectedDocument *types.Document, content string) {
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
func TestHeadingOnly(t *testing.T) {
	// given a valid heading
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
	compare(t, expectedDocument, actualContent)
}

func TestHeadingInvalid1(t *testing.T) {
	// given an invalid heading (missing space after '=')
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
	compare(t, expectedDocument, actualContent)
}
func TestHeadingInvalid2(t *testing.T) {
	// given an invalid heading (extra space before '=')
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
	compare(t, expectedDocument, actualContent)
}

func TestSection2(t *testing.T) {
	// given a section 2
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
	compare(t, expectedDocument, actualContent)
}

func TestHeadingWithSection2(t *testing.T) {
	// given a document with a heading, an empty line and a section
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
	compare(t, expectedDocument, actualContent)
}
func TestHeadingWithInvalidSection2(t *testing.T) {
	// given a document with a heading, an empty line and an invalid section (extra space at beginning of line)
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
	compare(t, expectedDocument, actualContent)
}
func TestInline1Word(t *testing.T) {
	// given a simple string
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
	compare(t, expectedDocument, actualContent)
}
func TestInlineSimple(t *testing.T) {
	// given a simple sentence
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
	compare(t, expectedDocument, actualContent)
}
func TestHeadingSectionInlineWithBoldQuote(t *testing.T) {
	// given
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
	compare(t, expectedDocument, actualContent)
}

func TestSingleListItem(t *testing.T) {
	// given an inline with invalid bold (3)
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
	compare(t, expectedDocument, actualContent)
}

func TestInvalidListItem(t *testing.T) {
	// given an inline with invalid bold (3)
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
	compare(t, expectedDocument, actualContent)
}

func TestListItems(t *testing.T) {
	// given an inline with invalid bold (3)
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
	compare(t, expectedDocument, actualContent)
}

func TestExternalLink(t *testing.T) {
	// given an inline with an external lin
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
	compare(t, expectedDocument, actualContent)
}

func TestExternalLinkWithEmptyText(t *testing.T) {
	// given an inline with an external lin
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
	compare(t, expectedDocument, actualContent)
}

func TestExternalLinkWithText(t *testing.T) {
	// given an inline with an external lin
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
	compare(t, expectedDocument, actualContent)
}

func TestElementLink(t *testing.T) {
	// given an inline with an external lin
	actualContent := "[link=http://foo.bar]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.ElementLink{Path: "http://foo.bar"},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestElementLinkWithSpaces(t *testing.T) {
	// given an inline with an element link
	actualContent := "[ link = http://foo.bar ]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.ElementLink{Path: "http://foo.bar"},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestElementLinkInvalid(t *testing.T) {
	// given an inline with an element link with missing ']'
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
	compare(t, expectedDocument, actualContent)
}

func TestElementID(t *testing.T) {
	// given an inline with an element ID
	actualContent := "[#img-foobar]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.ElementID{Value: "img-foobar"},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestElementIDWithSpaces(t *testing.T) {
	// given an inline with an element ID
	actualContent := "[ #img-foobar ]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.ElementID{Value: "img-foobar"},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestElementIDInvalid(t *testing.T) {
	// given an inline with an element ID with missing ']'
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
	compare(t, expectedDocument, actualContent)
}

func TestElementTitle(t *testing.T) {
	// given an inline with an element title
	actualContent := ".a title"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.ElementTitle{Content: "a title"},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestElementTitleInvalid1(t *testing.T) {
	// given an inline with an element title with extra space after '.'
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
	compare(t, expectedDocument, actualContent)
}

func TestElementTitleInvalid2(t *testing.T) {
	// given an inline with an element ID with missing '.' as first character
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
	compare(t, expectedDocument, actualContent)
}

package parser_test

import (
	"flag"
	"fmt"
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
		log.Warn(fmt.Sprintf("Starting test(s) with args=%v", flag.Args()))
	} else {
		log.Warn("Starting test(s) with no custom arguments")
	}
}

func TestHeadingOnly(t *testing.T) {
	// given a valid heading
	actualDocument, errs := ParseString("= a heading")
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Heading{Level: 1, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a heading"},
				},
			}},
		}}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInvalidHeading1(t *testing.T) {
	// given an invalid heading (missing space after '=')
	actualDocument, errs := ParseString("=a heading")
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "=a heading"},
				},
			},
		}}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInvalidHeading2(t *testing.T) {
	// given an invalid heading (extra space before '=')
	actualDocument, errs := ParseString(" = a heading")
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: " = a heading"},
				},
			},
		}}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestSection2(t *testing.T) {
	// given a section 2
	actualDocument, errs := ParseString(`== section 1`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Heading{Level: 2, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "section 1"},
				},
			}},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestHeadingWithSection2(t *testing.T) {
	// given a document with a heading, an empty line and a section
	actualDocument, errs := ParseString(`= a heading

== section 1`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Heading{Level: 1, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a heading"},
				},
			}},
			&types.EmptyLine{},
			&types.Heading{Level: 2, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "section 1"},
				},
			}},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestHeadingWithInvalidSection2(t *testing.T) {
	// given a document with a heading, an empty line and an invalid section (extra space at beginning of line)
	actualDocument, errs := ParseString(`= a heading

 == section 1`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Heading{Level: 1, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a heading"},
				},
			}},
			&types.EmptyLine{},
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: " == section 1"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInline1Word(t *testing.T) {
	// given a simple string
	actualDocument, errs := ParseString(`hello`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "hello"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInlineSimple(t *testing.T) {
	// given a simple sentence
	actualDocument, errs := ParseString(`a paragraph with some content`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with some content"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBoldQuote1Word(t *testing.T) {
	// given a bold quote of 1 word
	actualDocument, errs := ParseString(`*hello*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.BoldQuote{
						Content: "hello",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBoldQuote2Words(t *testing.T) {
	// given a bold quote of 2 words
	actualDocument, errs := ParseString(`*bold    content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.BoldQuote{
						Content: "bold    content",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBoldQuote3Words(t *testing.T) {
	// given a bold quote of 3 words
	actualDocument, errs := ParseString(`*some bold content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.BoldQuote{
						Content: "some bold content",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInlineWithBoldQuote(t *testing.T) {
	// given a sentence with a bold quote
	actualDocument, errs := ParseString(`a paragraph with *some bold content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with "},
					&types.BoldQuote{
						Content: "some bold content",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote1(t *testing.T) {
	// given an inline with invalid bold (1)
	actualDocument, errs := ParseString(`a paragraph with *some bold content`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with *some bold content"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote2(t *testing.T) {
	// given an inline with invalid bold (2)
	actualDocument, errs := ParseString(`a paragraph with *some bold content *`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with *some bold content *"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote3(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument, errs := ParseString(`a paragraph with * some bold content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with * some bold content*"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestHeadingSectionInlineWithBoldQuote(t *testing.T) {
	// given
	actualDocument, errs := ParseString(`= a heading

== section 1

a paragraph with *bold content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then a document with a heading, an empty line, a section and an inline with a bold quote
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Heading{Level: 1, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a heading"},
				},
			}},
			&types.EmptyLine{},
			&types.Heading{Level: 2, Content: &types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "section 1"},
				},
			}},
			&types.EmptyLine{},
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with "},
					&types.BoldQuote{
						Content: "bold content",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestSingleListItem(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument, errs := ParseString(`* a list item`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
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
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInvalidListItem(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument, errs := ParseString(`*an invalid list item`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "*an invalid list item"},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestListItems(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument, errs := ParseString(`* a first item
* a second item with *bold content*`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
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
						&types.BoldQuote{
							Content: "bold content",
						},
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLink(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`a link to https://foo.bar`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a link to "},
					&types.ExternalLink{
						URL: "https://foo.bar",
					},
				},
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLinkWithEmptyText(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`a link to https://foo.bar[]`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
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
	}
	t.Log(fmt.Sprintf("Actual document: %v", actualDocument.Elements[0].(*types.InlineContent).Elements[1]))
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLinkWithText(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`a link to mailto:foo@bar[the foo@bar email]`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
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
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBlockImageWithEmptyAltText(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`image::images/foo.png[]`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.BlockImage{
				Path: "images/foo.png",
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBlockImageWithAltText(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`image::images/foo.png[the foo.png image]`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	altText := "the foo.png image"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.BlockImage{
				Path:    "images/foo.png",
				AltText: &altText,
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBlockImageWithIDAndTitleAndDimensions(t *testing.T) {
	// given an inline with an external lin
	actualDocument, errs := ParseString(`[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`)
	require.Nil(t, errs)
	log.Debugf("actual document: %s", actualDocument.String())
	// then
	altText := "the foo.png image"
	width := "600"
	height := "400"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "[#img-foobar]"}}},
			&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: ".A title to foobar"}}},
			&types.InlineContent{Elements: []types.DocElement{&types.StringElement{Content: "[link=http://foo.bar]"}}},
			&types.BlockImage{
				Path:    "images/foo.png",
				AltText: &altText,
				Width:   &width,
				Height:  &height,
			},
		},
	}
	log.Debugf("expected document: %s", expectedDocument.String())
	assert.EqualValues(t, expectedDocument, actualDocument)
}

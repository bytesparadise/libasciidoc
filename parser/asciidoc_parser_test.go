package parser

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"reflect"

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

func parseString(t *testing.T, content string) *Document {
	reader := strings.NewReader(content)
	result, err := ParseReader("", reader)
	if err != nil {
		errors := err.(errList)
		for i := range errors {
			t.Log(fmt.Sprintf("Error: %v\n", errors[i].Error()))
		}
	}
	require.Nil(t, err)
	require.NotNil(t, result)
	display(t, result)
	require.IsType(t, &Document{}, result)
	document := result.(*Document)
	return document
}

func display(t *testing.T, content interface{}) {
	switch content.(type) {
	case *Document:
		document := content.(*Document)
		for i := range document.Elements {
			element := document.Elements[i]
			t.Log(fmt.Sprintf("%v", element.String()))
		}
	default:
		assert.Fail(t, fmt.Sprintf("Unexpected type of 'result': %v", reflect.TypeOf(content)))
	}
}

func TestHeadingOnly(t *testing.T) {
	// given a valid heading
	actualDocument := parseString(t, "= a heading")
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&Heading{Level: 1, Content: &InlineContent{
				Elements: []interface{}{
					"a heading",
				},
			}},
		}}
	display(t, expectedDocument)
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInvalidHeading1(t *testing.T) {
	// given an invalid heading (missing space after '=')
	actualDocument := parseString(t, "=a heading")
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"=a heading",
				},
			},
		}}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInvalidHeading2(t *testing.T) {
	// given an invalid heading (extra space before '=')
	actualDocument := parseString(t, " = a heading")
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					" = a heading",
				},
			},
		}}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestSection2(t *testing.T) {
	// given a section 2
	actualDocument := parseString(t, `== section 1`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&Heading{Level: 2, Content: &InlineContent{
				Elements: []interface{}{
					"section 1",
				},
			}},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestHeadingWithSection2(t *testing.T) {
	// given a document with a heading, an empty line and a section
	actualDocument := parseString(t, `= a heading

== section 1`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&Heading{Level: 1, Content: &InlineContent{
				Elements: []interface{}{
					"a heading",
				},
			}},
			&EmptyLine{},
			&Heading{Level: 2, Content: &InlineContent{
				Elements: []interface{}{
					"section 1",
				},
			}},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestHeadingWithInvalidSection2(t *testing.T) {
	// given a document with a heading, an empty line and an invalid section (extra space at beginning of line)
	actualDocument := parseString(t, `= a heading

 == section 1`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&Heading{Level: 1, Content: &InlineContent{
				Elements: []interface{}{
					"a heading",
				},
			}},
			&EmptyLine{},
			&InlineContent{
				Elements: []interface{}{
					" == section 1",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInline1Word(t *testing.T) {
	// given a simple string
	actualDocument := parseString(t, `hello`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"hello",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInlineSimple(t *testing.T) {
	// given a simple sentence
	actualDocument := parseString(t, `a paragraph with some content`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with some content",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBoldQuote1Word(t *testing.T) {
	// given a bold quote of 1 word
	actualDocument := parseString(t, `*hello*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					&BoldQuote{
						Content: "hello",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBoldQuote2Words(t *testing.T) {
	// given a bold quote of 2 words
	actualDocument := parseString(t, `*bold    content*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					&BoldQuote{
						Content: "bold    content",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBoldQuote3Words(t *testing.T) {
	// given a bold quote of 3 words
	actualDocument := parseString(t, `*some bold content*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					&BoldQuote{
						Content: "some bold content",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestInlineWithBoldQuote(t *testing.T) {
	// given a sentence with a bold quote
	actualDocument := parseString(t, `a paragraph with *some bold content*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with ",
					&BoldQuote{
						Content: "some bold content",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote1(t *testing.T) {
	// given an inline with invalid bold (1)
	actualDocument := parseString(t, `a paragraph with *some bold content`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with *some bold content",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote2(t *testing.T) {
	// given an inline with invalid bold (2)
	actualDocument := parseString(t, `a paragraph with *some bold content *`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with *some bold content *",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInlineWithInvalidBoldQuote3(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument := parseString(t, `a paragraph with * some bold content*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with * some bold content*",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestHeadingSectionInlineWithBoldQuote(t *testing.T) {
	// given
	actualDocument := parseString(t, `= a heading

== section 1

a paragraph with *bold content*`)
	// then a document with a heading, an empty line, a section and an inline with a bold quote
	expectedDocument := &Document{
		Elements: []DocElement{
			&Heading{Level: 1, Content: &InlineContent{
				Elements: []interface{}{
					"a heading",
				},
			}},
			&EmptyLine{},
			&Heading{Level: 2, Content: &InlineContent{
				Elements: []interface{}{
					"section 1",
				},
			}},
			&EmptyLine{},
			&InlineContent{
				Elements: []interface{}{
					"a paragraph with ",
					&BoldQuote{
						Content: "bold content",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestSingleListItem(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument := parseString(t, `* a list item`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&ListItem{
				Content: &InlineContent{
					Elements: []interface{}{
						"a list item",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestInvalidListItem(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument := parseString(t, `*an invalid list item`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"*an invalid list item",
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestListItems(t *testing.T) {
	// given an inline with invalid bold (3)
	actualDocument := parseString(t, `* a first item
* a second item with *bold content*`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&ListItem{
				Content: &InlineContent{
					Elements: []interface{}{
						"a first item",
					},
				},
			},
			&ListItem{
				Content: &InlineContent{
					Elements: []interface{}{
						"a second item with ",
						&BoldQuote{
							Content: "bold content",
						},
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLink(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `a link to https://foo.bar`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a link to ",
					&ExternalLink{
						URL: "https://foo.bar",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLinkWithEmptyText(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `a link to https://foo.bar[]`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a link to ",
					&ExternalLink{
						URL:  "https://foo.bar",
						Text: "",
					},
				},
			},
		},
	}
	t.Log(fmt.Sprintf("Actual document: %v", actualDocument.Elements[0].(*InlineContent).Elements[1]))
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestExternalLinkWithText(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `a link to mailto:foo@bar[the foo@bar email]`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{
				Elements: []interface{}{
					"a link to ",
					&ExternalLink{
						URL:  "mailto:foo@bar",
						Text: "the foo@bar email",
					},
				},
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBlockImageWithEmptyAltText(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `image::images/foo.png[]`)
	// then
	expectedDocument := &Document{
		Elements: []DocElement{
			&BlockImage{
				Path: "images/foo.png",
			},
		},
	}
	t.Log(fmt.Sprintf("Actual document: %v", actualDocument.Elements[0].(*BlockImage)))
	assert.EqualValues(t, expectedDocument, actualDocument)
}
func TestBlockImageWithAltText(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `image::images/foo.png[the foo.png image]`)
	// then
	altText := "the foo.png image"
	expectedDocument := &Document{
		Elements: []DocElement{
			&BlockImage{
				Path:    "images/foo.png",
				AltText: &altText,
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

func TestBlockImageWithIDAndTitleAndDimensions(t *testing.T) {
	// given an inline with an external lin
	actualDocument := parseString(t, `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`)
	// then
	altText := "the foo.png image"
	width := "600"
	height := "400"
	expectedDocument := &Document{
		Elements: []DocElement{
			&InlineContent{Elements: []interface{}{"[#img-foobar]"}},
			&InlineContent{Elements: []interface{}{".A title to foobar"}},
			&InlineContent{Elements: []interface{}{"[link=http://foo.bar]"}},
			&BlockImage{
				Path:    "images/foo.png",
				AltText: &altText,
				Width:   &width,
				Height:  &height,
			},
		},
	}
	assert.EqualValues(t, expectedDocument, actualDocument)
}

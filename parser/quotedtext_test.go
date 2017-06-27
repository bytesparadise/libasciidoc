package parser_test

import (
	"testing"

	"github.com/bytesparadise/libasciidoc/types"
)

func TestBoldTextOf1Word(t *testing.T) {
	// given a bold quote of 1 word
	actualContent := "*hello*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Bold,
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

func TestBoldTextOf2Words(t *testing.T) {
	// given a bold quote of 2 words
	actualContent := "*bold    content*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.DocElement{
							&types.StringElement{Content: "bold    content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}
func TestBoldTextOf3Words(t *testing.T) {
	// given a bold quote of 3 words
	actualContent := "*some bold content*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some bold content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestInlineWithBoldText(t *testing.T) {
	// given a sentence with a bold quote
	actualContent := "a paragraph with *some bold content*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with "},
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some bold content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestInlineWithInvalidBoldText1(t *testing.T) {
	// given an inline with invalid bold (1)
	actualContent := "a paragraph with *some bold content"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with *some bold content"},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestInlineWithInvalidBoldText2(t *testing.T) {
	// given an inline with invalid bold (2)
	actualContent := "a paragraph with *some bold content *"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with *some bold content *"},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestInlineWithInvalidBoldText3(t *testing.T) {
	// given an inline with invalid bold (3)
	actualContent := "a paragraph with * some bold content*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with * some bold content*"},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}
func TestItalicTextWith3Words(t *testing.T) {
	// given an italic quote of 3 words
	actualContent := "_some italic content_"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some italic content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestMonospaceTextWith3Words(t *testing.T) {
	// given a monospace quote of 3 words
	actualContent := "`some monospace content`"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some monospace content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestItalicTextWithinBoldText(t *testing.T) {
	actualContent := "*some _italic_ content*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.DocElement{
									&types.StringElement{Content: "italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}
func TestBoldTextWithinItalicText(t *testing.T) {
	// given a bold quote of 3 words
	actualContent := "_some *bold* content_"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.DocElement{
									&types.StringElement{Content: "bold"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestMonospaceTextWithinBoldTextWithinItalicQuote(t *testing.T) {
	actualContent := "*some _italic and `monospaced content`_*"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.DocElement{
									&types.StringElement{Content: "italic and "},
									&types.QuotedText{
										Kind: types.Monospace,
										Elements: []types.DocElement{
											&types.StringElement{Content: "monospaced content"},
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
	compare(t, expectedDocument, actualContent)
}

func TestItalicTextWithinItalicText(t *testing.T) {
	// given a bold quote of 3 words
	actualContent := "_some _very italic_ content_"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.DocElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.DocElement{
									&types.StringElement{Content: "very italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func xTestAllQuotes(t *testing.T) {
	// given an inline with invalid bold (3)
	actualContent := "*bold phrase* & **char**acter**s**\n" +
		"_italic phrase_ & __char__acter__s__\n" +
		"*_bold italic phrase_* & **__char__**acter**__s__**\n" +
		"`monospace phrase` & ``char``acter``s``\n" +
		"`*monospace bold phrase*` & ``**char**``acter``**s**``\n" +
		"`_monospace italic phrase_` & ``__char__``acter``__s__``\n" +
		"`*_monospace bold italic phrase_*` & \n" +
		"``**__char__**``acter``**__s__**``"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.InlineContent{
				Elements: []types.DocElement{
					&types.StringElement{Content: "a paragraph with * some bold content*"},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)

}

package parser_test

import (
	"testing"

	"github.com/bytesparadise/libasciidoc/types"
)

func TestBlockImageWithEmptyAlt(t *testing.T) {
	// given an inline with an external lin
	actualContent := "image::images/foo.png[]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.BlockImage{
				Macro: types.BlockImageMacro{
					Path: "images/foo.png",
					Alt:  "foo",
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}
func TestBlockImageWithAlt(t *testing.T) {
	// given an inline with an external lin
	actualContent := "image::images/foo.png[the foo.png image]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.BlockImage{
				Macro: types.BlockImageMacro{
					Path: "images/foo.png",
					Alt:  "the foo.png image",
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestBlockImageWithDimensionsAndIDLinkTitleMeta(t *testing.T) {
	// given an inline with an external lin
	actualContent := "[#img-foobar]\n" +
		".A title to foobar\n" +
		"[link=http://foo.bar]\n" +
		"image::images/foo.png[the foo.png image, 600, 400]"
	width := "600"
	height := "400"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.BlockImage{
				Macro: types.BlockImageMacro{
					Path:   "images/foo.png",
					Alt:    "the foo.png image",
					Width:  &width,
					Height: &height,
				},
				ID:    &types.ElementID{Value: "img-foobar"},
				Title: &types.ElementTitle{Content: "A title to foobar"},
				Link:  &types.ElementLink{Path: "http://foo.bar"},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestBlockImageAppendingInlineContent(t *testing.T) {
	// given an inline with an external lin
	actualContent := "a paragraph\nimage::images/foo.png[]"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a paragraph"},
						},
					},
					&types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "image::images/foo.png[]"},
						},
					},
				},
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

package parser_test

import (
	"testing"

	"github.com/bytesparadise/libasciidoc/types"
)

func TestDelimitedSourceBlockWithSingleLine(t *testing.T) {
	// given a source block of 1 line
	content := "some source code"
	actualContent := "```\n" + content + "\n```"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.DelimitedBlock{
				Kind:    types.SourceBlock,
				Content: content,
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestDelimitedSourceBlockWithMultipleLines(t *testing.T) {
	// given a source block of multiple lines
	content := "some source code\nwith an empty line\n\nin the middle"
	actualContent := "```\n" + content + "\n```"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.DelimitedBlock{
				Kind:    types.SourceBlock,
				Content: content,
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

func TestDelimitedSourceBlockWithNoLine(t *testing.T) {
	// given an empty source block
	content := ""
	actualContent := "```\n" + content + "```"
	expectedDocument := &types.Document{
		Elements: []types.DocElement{
			&types.DelimitedBlock{
				Kind:    types.SourceBlock,
				Content: content,
			},
		},
	}
	compare(t, expectedDocument, actualContent)
}

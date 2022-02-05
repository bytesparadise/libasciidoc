package renderer

import (
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// SectionNumbers a registry of section numbers

type SectionNumbers map[string]string // assigned number by section id

// NewSectionNumbers initializes the registry with the content of the given doc
// (ie, it traverses the doc to look for sections, and assigns them numbers, incrementally)
func NewSectionNumbers(doc *types.Document) (SectionNumbers, error) {
	// traverse doc and its sections and assign numbers to the latters
	return traverseElements(doc.Elements, "")
}

func traverseElements(elements []interface{}, prefix string) (map[string]string, error) {
	result := map[string]string{}
	counter := 0
	for _, e := range elements {
		if s, ok := e.(*types.Section); ok {
			id, err := s.GetID()
			if err != nil {
				return nil, err
			}
			counter++
			n := prefix + strconv.Itoa(counter)
			result[id] = n
			numbers, err := traverseElements(s.Elements, n+".")
			if err != nil {
				return nil, err
			}
			for id, n := range numbers {
				result[id] = n
			}
		}
	}
	return result, nil
}

package types

import (
	"bytes"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// filterOption allows for filtering elements by type
type filterOption func(element interface{}) bool

// filterEmptyPreamble filters the element if it is an empty preamble
func filterEmptyPreamble() filterOption {
	return func(element interface{}) bool {
		result := false
		if p, match := element.(Preamble); match {
			result = p.Elements == nil || len(p.Elements) == 0
		}
		log.Debugf(" element of type '%T' is an empty preamble: %t", element, result)
		return result
	}
}

// filterBlankLine filters the element if it is a blank line
func filterBlankLine() filterOption {
	return func(element interface{}) bool {
		_, result := element.(BlankLine)
		// defer log.Debugf(" element of type '%T' is a blankline: %t", element, result)
		return result
	}
}

// filterEmptyElements excludes the unrelevant (empty) blocks
func filterEmptyElements(blocks []interface{}, filters ...filterOption) []interface{} {
	log.Debugf("Filtering %d blocks...", len(blocks))
	elements := make([]interface{}, 0)
blocks:
	for _, block := range blocks {
		// check if filter option applies to the element
		switch block := block.(type) {
		case []interface{}:
			result := filterEmptyElements(block, filters...)
			elements = append(elements, result...)
		default:
			if block != nil {
				// log.Debugf(" converting block of type '%T' into a interface{}...", block)
				for _, filter := range filters {
					if filter(block) {
						log.Debugf(" discarding block of type '%T'.", block)
						continue blocks
					}
				}
				log.Debugf(" keeping block of type '%T'.", block)
				elements = append(elements, block)
				continue
			}
		}
	}
	return elements
}

// NilSafe returns a new slice if the given elements is nil, otherwise it returns the given elements
func NilSafe(elements []interface{}) []interface{} {
	if elements != nil {
		return elements
	}
	return make([]interface{}, 0)
}

// MergeStringElements merge string elements together
func MergeStringElements(elements ...interface{}) InlineElements {
	result := make([]interface{}, 0)
	buff := bytes.NewBuffer(nil)
	for _, element := range elements {
		if element == nil {
			continue
		}
		switch element := element.(type) {
		case string:
			buff.WriteString(element)
		case []byte:
			for _, b := range element {
				buff.WriteByte(b)
			}
		case StringElement:
			content := element.Content
			buff.WriteString(content)
		case *StringElement:
			content := element.Content
			buff.WriteString(content)
		case []interface{}:
			if len(element) > 0 {
				f := MergeStringElements(element...)
				result, buff = appendBuffer(result, buff)
				result = MergeStringElements(append(result, f...)...)
			}
		default:
			// log.Debugf("Merging with 'default' case an element of type %[1]T", element)
			result, buff = appendBuffer(result, buff)
			result = append(result, element)
		}
	}
	// if buff was filled because some text was found
	result, _ = appendBuffer(result, buff)
	return result
}

// appendBuffer appends the content of the given buffer to the given array of elements,
// and returns a new buffer, or returns the given arguments if the buffer was empty
func appendBuffer(elements []interface{}, buff *bytes.Buffer) ([]interface{}, *bytes.Buffer) {
	if buff.Len() > 0 {
		s, _ := NewStringElement(buff.String())
		return append(elements, s), bytes.NewBuffer(nil)
	}
	return elements, buff
}

// applyFunc a function to apply on the result of the `apply` function below, before returning
type applyFunc func(s string) string

// Apply applies the given funcs to transform the given input
func Apply(source string, fs ...applyFunc) string {
	result := source
	for _, f := range fs {
		result = f(result)
	}
	// log.Debugf("applied '%s' -> '%s' (%v characters)", source, result, len(result))
	return result
}

func toString(lines []interface{}) ([]string, error) {
	result := make([]string, len(lines))
	for i, line := range lines {
		l, ok := line.(string)
		if !ok {
			return []string{}, errors.Errorf("expected a string, but got a %T", line)
		}
		result[i] = l
	}
	return result, nil
}

// SearchAttributeDeclaration returns the value of the DocumentAttributeDeclaration whose name is given
func SearchAttributeDeclaration(elements []interface{}, name string) (DocumentAttributeDeclaration, bool) {
	for _, e := range elements {
		switch e := e.(type) {
		case Section:
			if result, found := SearchAttributeDeclaration(e.Elements, name); found {
				return result, found
			}
		case Preamble:
			if result, found := SearchAttributeDeclaration(e.Elements, name); found {
				return result, found
			}
		case DocumentAttributeDeclaration:
			if e.Name == name {
				return e, true
			}
		}
	}
	return DocumentAttributeDeclaration{}, false
}

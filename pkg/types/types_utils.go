package types

import (
	"bytes"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// FilterCriterion returns true if the given element is to be filtered out
type FilterCriterion func(element interface{}) bool

// EmptyPreambleMatcher filters the element if it is an empty preamble
var EmptyPreambleMatcher FilterCriterion = func(element interface{}) bool {
	result := false
	if p, match := element.(Preamble); match {
		result = p.Elements == nil || len(p.Elements) == 0
	}
	log.Debugf(" element of type '%T' is an empty preamble: %t", element, result)
	return result
}

// BlankLineMatcher filters the element if it is a blank line
var BlankLineMatcher FilterCriterion = func(element interface{}) bool {
	_, result := element.(BlankLine)
	return result
}

// DocumentAttributeMatcher filters the element if it is a blank line
var DocumentAttributeMatcher FilterCriterion = func(element interface{}) bool {
	switch element.(type) {
	case DocumentAttributeDeclaration, DocumentAttributeSubstitution, DocumentAttributeReset:
		return true
	default:
		return false
	}
}

// FilterOut excludes the unrelevant (empty) elements
func FilterOut(elements []interface{}, filters ...FilterCriterion) []interface{} {
	log.Debugf("filtering %d blocks...", len(elements))
	result := make([]interface{}, 0)
blocks:
	for _, element := range elements {
		// check if filter option applies to the element
		switch element := element.(type) {
		case []interface{}:
			result = append(result, FilterOut(element, filters...)...)
		default:
			for _, filter := range filters {
				if filter(element) {
					log.Debugf("discarding block of type '%T'", element)
					continue blocks
				}
			}
			log.Debugf("keeping block of type '%T'", element)
			result = append(result, element)
			continue
		}
	}
	return result
}

// NilSafe returns a new slice if the given elements is nil, otherwise it returns the given elements
func NilSafe(elements []interface{}) []interface{} {
	if elements != nil {
		return elements
	}
	return make([]interface{}, 0)
}

// MergeStringElements merge string elements together
func MergeStringElements(elements ...interface{}) []interface{} {
	result := make([]interface{}, 0)
	buf := bytes.NewBuffer(nil)
	for _, element := range elements {
		if element == nil {
			continue
		}
		switch element := element.(type) {
		case string:
			buf.WriteString(element)
		case []byte:
			for _, b := range element {
				buf.WriteByte(b)
			}
		case StringElement:
			content := element.Content
			buf.WriteString(content)
		case *StringElement:
			content := element.Content
			buf.WriteString(content)
		case []interface{}:
			if len(element) > 0 {
				f := MergeStringElements(element...)
				result, buf = appendBuffer(result, buf)
				result = MergeStringElements(append(result, f...)...)
			}
		default:
			// log.Debugf("Merging with 'default' case an element of type %[1]T", element)
			result, buf = appendBuffer(result, buf)
			result = append(result, element)
		}
	}
	// if buf was filled because some text was found
	result, _ = appendBuffer(result, buf)
	return result
}

// appendBuffer appends the content of the given buffer to the given array of elements,
// and returns a new buffer, or returns the given arguments if the buffer was empty
func appendBuffer(elements []interface{}, buf *bytes.Buffer) ([]interface{}, *bytes.Buffer) {
	if buf.Len() > 0 {
		s, _ := NewStringElement(buf.String())
		return append(elements, s), bytes.NewBuffer(nil)
	}
	return elements, buf
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

package types

import (
	"bytes"

	"github.com/pkg/errors"
)

// NilSafe returns a new slice if the given elements is nil, otherwise it returns the given elements
func NilSafe(elements []interface{}) []interface{} {
	if elements != nil {
		return elements
	}
	return make([]interface{}, 0)
}

// Merge merge string elements together
func Merge(elements ...interface{}) []interface{} {
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
			buf.Write(element)
		case StringElement:
			content := element.Content
			buf.WriteString(content)
		case []interface{}:
			if len(element) > 0 {
				f := Merge(element...)
				result, buf = appendBuffer(result, buf)
				result = Merge(append(result, f...)...)
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

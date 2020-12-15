package types

import (
	"bytes"
)

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

// ReduceOption an option to apply on the reduced content when it is a `string`
type ReduceOption func(string) string

// Reduce merges and returns a string if the given elements only contain a single StringElement
// (ie, return its `Content`), otherwise return the given elements or empty string if the elements
// is `nil` or an empty `[]interface{}`
func Reduce(elements interface{}, opts ...ReduceOption) interface{} {
	switch e := elements.(type) {
	case []interface{}:
		e = Merge(e...)
		switch len(e) {
		case 0: // if empty, return nil
			elements = nil
		case 1:
			if e, ok := e[0].(StringElement); ok {
				c := e.Content
				for _, apply := range opts {
					c = apply(c)
				}
				elements = c
			}
		}
	case string:
		for _, apply := range opts {
			e = apply(e)
		}
	}
	return elements
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

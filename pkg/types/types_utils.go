package types

import (
	"fmt"
	"strings"
)

// merge merge string elements together, keeping
// other elements intact
func merge(elements ...interface{}) []interface{} {
	result := make([]interface{}, 0, len(elements))
	buf := &strings.Builder{}
	for _, element := range elements {
		if element == nil {
			continue
		}
		switch element := element.(type) {
		case string:
			buf.WriteString(element)
		case []byte:
			buf.Write(element)
		case *StringElement:
			buf.WriteString(element.Content)
		case []interface{}:
			if len(element) > 0 {
				f := merge(element...)
				if content := buf.String(); len(content) > 0 {
					result = append(result, &StringElement{
						Content: content,
					})
					buf = &strings.Builder{}
				}
				result = merge(append(result, f...)...)
			}
		default:
			// log.Debugf("Merging with 'default' case an element of type %[1]T", element)
			if content := buf.String(); len(content) > 0 {
				if symbol, ok := element.(*Symbol); ok && symbol.Name == " -- " && strings.HasSuffix(content, " ") {
					// trim 1 space from actual result
					content = content[:len(content)-1]
				}
				result = append(result, &StringElement{
					Content: content,
				})
				buf = &strings.Builder{}
			}
			result = append(result, element)
		}
	}
	// if buf was filled because some text was found
	if buf.Len() > 0 {
		result = append(result, &StringElement{
			Content: buf.String(),
		})
	}
	return result
}

// Flatten
func Flatten(elements []interface{}) []interface{} {
	result := make([]interface{}, 0, len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case []interface{}:
			result = append(result, e...)
		default:
			result = append(result, e)
		}
	}
	return result
}

// AllNilEntries returns true if all the entries in the given `elements` are `nil`
func AllNilEntries(elements []interface{}) bool {
	for _, e := range elements {
		switch e := e.(type) {
		case []interface{}: // empty slice if not `nil` since it has a type
			if !AllNilEntries(e) {
				return false
			}
		default:
			if e != nil {
				return false
			}
		}
	}
	return true
}

// ReduceOption an option to apply on the reduced content when it is a `string`
type ReduceOption func(string) string

// Reduce merges and returns a string if the given elements only contain a single StringElement
// (ie, return its `Content`), otherwise return the given elements or empty string if the elements
// is `nil` or an empty `[]interface{}`
func Reduce(elements interface{}, opts ...ReduceOption) interface{} {
	switch e := elements.(type) {
	case []interface{}:
		e = merge(e...)
		switch len(e) {
		case 0: // if empty, return nil
			return nil
		case 1:
			if e, ok := e[0].(*StringElement); ok {
				c := e.Content
				for _, apply := range opts {
					c = apply(c)
				}
				return c
			}
			return e
		default:
			return e
		}
	case *StringElement:
		return Reduce(e.Content)
	case string:
		for _, apply := range opts {
			e = apply(e)
		}
		switch len(e) {
		case 0:
			return nil
		default:
			return e
		}
	default:
		return elements
	}
}

// applyFunc a function to apply on the result of the `apply` function below, before returning
type applyFunc func(s string) string

// Apply applies the given funcs to transform the given input
func Apply(source string, fs ...applyFunc) string {
	result := source
	for _, f := range fs {
		result = f(result)
	}
	return result
}

func stringify(element interface{}) string {
	switch element := element.(type) {
	case []interface{}:
		result := strings.Builder{}
		for _, e := range element {
			result.WriteString(stringify(e))
		}
		return result.String()
	case string:
		return element
	case *StringElement:
		return element.Content
	case *SpecialCharacter:
		return element.Name
	case *AttributeReference: // TODO: should never happen?
		return "{" + element.Name + "}"
	default:
		return fmt.Sprintf("%v", element) // "best-effort" here
	}
}

// TrimLeft returns a slice of elements where the
// `strings.TrimLeft` func was applied on the content of the first entry
// if it is a `*StringElement`
func TrimLeft(elements []interface{}, cutset string) []interface{} {
	if len(elements) == 0 {
		return elements
	}
	if first, ok := elements[0].(*StringElement); ok {
		first.Content = strings.TrimLeft(first.Content, cutset)
	}
	return elements
}

// Append appends all given elements. If an element is an `[]interface{}`, then it appends its content
func Append(elements ...interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0, len(elements)) // best guess for initial capacity
	for _, e := range elements {
		switch e := e.(type) {
		case []interface{}:
			result = append(result, e...)
		case nil:
			continue
		default:
			result = append(result, e)
		}
	}
	return result, nil
}

func SplitElementsPerLine(elements []interface{}) [][]interface{} {
	lines := make([][]interface{}, 0, len(elements))
	line := make([]interface{}, 0, len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case *StringElement:
			// split
			s := strings.Split(e.Content, "\n")
			for i := range s {
				// only append if string is non-empty
				if len(s[i]) > 0 {
					line = append(line, &StringElement{
						Content: s[i],
					})
				}
				if i < len(s)-1 { // move to next line
					lines = append(lines, line)
					// reset
					line = make([]interface{}, 0, len(elements))
				}
			}
		default:
			line = append(line, e)
		}
	}
	// don't forget the last line
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}

// InsertAt inserts the given element in the target at the given index
// (thus moving all following elements by 1)
func InsertAt(elements []interface{}, element interface{}, index int) []interface{} {
	if element == nil {
		return elements
	}
	result := make([]interface{}, len(elements)+1)
	copy(result[0:index], elements[0:index])
	result[index] = element
	copy(result[index+1:], elements[index:])
	return result
}

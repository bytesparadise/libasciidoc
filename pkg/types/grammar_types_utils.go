package types

import (
	"bytes"
	"reflect"

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

// nilSafe returns a new slice if the given elements is nil, otherwise it returns the given elements
func nilSafe(elements []interface{}) []interface{} {
	if elements != nil {
		return elements
	}
	return make([]interface{}, 0)
}

// removeEmptyTrailingStringElement removes the last
// func removeEmptyTrailingStringElement([]interface{}) []interface{}

func mergeElements(elements ...interface{}) InlineElements {
	result := make([]interface{}, 0)
	// log.Debugf("merging %d element(s):", len(elements))
	buff := bytes.NewBuffer(nil)
	for _, element := range elements {
		if element == nil {
			continue
		}
		switch element := element.(type) {
		case string:
			buff.WriteString(element)
		case *string:
			buff.WriteString(*element)
		case []byte:
			for _, b := range element {
				buff.WriteByte(b)
			}
		case StringElement:
			content := element.Content
			buff.WriteString(content)
		case []interface{}:
			if len(element) > 0 {
				f := mergeElements(element...)
				result, buff = appendBuffer(result, buff)
				result = mergeElements(append(result, f...)...)
			}
		default:
			// log.Debugf("Merging with 'default' case an element of type %[1]T", element)
			result, buff = appendBuffer(result, buff)
			result = append(result, element)
		}
	}
	// if buff was filled because some text was found
	result, _ = appendBuffer(result, buff)
	// log.Debugf(" -> '%[1]v' (%[1]T)", result)
	return result
}

// DefaultAttribute a function to specify a default attribute
type DefaultAttribute func(map[string]interface{})

// WithNumberingStyle specifies the numbering type in an OrderedList
func WithNumberingStyle(t NumberingStyle) DefaultAttribute {
	return func(attributes map[string]interface{}) {
		attributes["numbering"] = t
	}
}

func mergeAttributes(attributes []interface{}, defaults ...DefaultAttribute) map[string]interface{} {
	result := make(map[string]interface{})
	// fill with the default values, that can be overridden afterwards
	for _, d := range defaults {
		d(result)
	}
	if attributes == nil {
		return result
	}
	log.Debugf("attributes with defaults: %v", result)
	for _, attr := range attributes {
		log.Debugf("processing attributes of %T", attr)
		if attr, ok := attr.(map[string]interface{}); ok {
			for k, v := range attr {
				result[k] = v
			}
		}
	}
	log.Debugf("merged attributes : %v", result)
	return result
}

// appendBuffer appends the content of the given buffer to the given array of elements,
// and returns a new buffer, or returns the given arguments if the buffer was empty
func appendBuffer(elements []interface{}, buff *bytes.Buffer) ([]interface{}, *bytes.Buffer) {
	if buff.Len() > 0 {
		return append(elements, NewStringElement(buff.String())), bytes.NewBuffer(nil)
	}
	return elements, buff
}

// stringifyOption a function to apply on the result of the `stringify` function below, before returning
type stringifyOption func(s string) string

// stringify convert the given elements into a string, then applies the optional `funcs` to convert the string before returning it.
// These stringifyFuncs can be used to trim the content, for example
func stringify(elements []interface{}, options ...stringifyOption) (string, error) {
	mergedElements := mergeElements(elements)
	b := make([]byte, 0)
	buff := bytes.NewBuffer(b)
	for _, element := range mergedElements {
		switch element := element.(type) {
		case StringElement:
			buff.WriteString(element.Content)
		default:
			return "", errors.Errorf("cannot convert element of type '%T' to string content", element)
		}

	}
	result := buff.String()
	for _, f := range options {
		result = f(result)
	}
	// log.Debugf("stringified %v -> '%s' (%v characters)", elements, result, len(result))
	return result, nil
}

func toPtr(element interface{}) interface{} {
	value := reflect.ValueOf(element)
	if value.Type().Kind() == reflect.Ptr {
		return element
	}
	ptr := reflect.New(reflect.TypeOf(element))
	temp := ptr.Elem()
	temp.Set(value)
	// log.Debugf("Returning pointer of type %T", ptr.Interface())
	return ptr.Interface()
}

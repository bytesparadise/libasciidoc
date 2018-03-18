package types

import (
	"bytes"
	"reflect"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func toDocElements(elements []interface{}) ([]DocElement, error) {
	result := make([]DocElement, len(elements))
	for i, element := range elements {
		element, ok := element.(DocElement)
		if !ok {
			return nil, errors.Errorf("unexpected element of type: %T (expected a DocElement instead)", element)
		}
		result[i] = element
	}
	return result, nil
}

func toInlineElements(elements []interface{}) ([]InlineElement, error) {
	mergedElements := mergeElements(elements)
	result := make([]InlineElement, len(mergedElements))
	for i, element := range mergedElements {
		element, ok := element.(InlineElement)
		if !ok {
			return nil, errors.Errorf("unexpected element of type: %T (expected a InlineElement instead)", element)
		}
		result[i] = element
	}

	return result, nil
}

// filterUnrelevantElements excludes the unrelevant (empty) blocks
func filterUnrelevantElements(blocks []interface{}) []DocElement {
	log.Debugf("Filtering %d blocks...", len(blocks))
	elements := make([]DocElement, 0)
	for _, block := range blocks {
		log.Debugf(" converting block of type '%T' into a DocElement...", block)
		switch block := block.(type) {
		case BlankLine:
			// exclude blank lines from here, we won't need them in the rendering anyways
		case Preamble:
			// exclude empty preambles
			if len(block.Elements) > 0 {
				// exclude empty preamble
				elements = append(elements, block)
			}
		case []interface{}:
			result := filterUnrelevantElements(block)
			elements = append(elements, result...)
		default:
			if block != nil {
				elements = append(elements, block)
			}
		}
	}
	return elements
}

func mergeElements(elements []interface{}, extraElements ...interface{}) []interface{} {
	result := make([]interface{}, 0)
	allElements := append(elements, extraElements...)
	// log.Debugf("Merging %d element(s):", len(allElements))
	buff := bytes.NewBuffer(nil)
	for _, element := range allElements {
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
		case InlineContent:
			inlineElements := make([]interface{}, len(element.Elements))
			for i, e := range element.Elements {
				inlineElements[i] = e
			}
			result = mergeElements(result, inlineElements...)
		case []interface{}:
			if len(element) > 0 {
				f := mergeElements(element)
				result, buff = appendBuffer(result, buff)
				result = mergeElements(result, f...)
			}
		default:
			log.Debugf("Merging with 'default' case an element of type %[1]T", element)
			result, buff = appendBuffer(result, buff)
			result = append(result, element)
		}
	}
	// if buff was filled because some text was found
	result, _ = appendBuffer(result, buff)

	return result
}

func mergeAttributes(attributes []interface{}) map[string]interface{} {
	if attributes == nil || len(attributes) == 0 {
		return nil
	}
	result := make(map[string]interface{})
	for _, attr := range attributes {
		if attr, ok := attr.(map[string]interface{}); ok {
			for k, v := range attr {
				result[k] = v
			}
		}
	}
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
type stringifyOption func(s string) (string, error)

// stringify convert the given elements into a string, then applies the optional `funcs` to convert the string before returning it.
// These StringifyFuncs can be used to trim the content, for example
func stringify(elements []interface{}, options ...stringifyOption) (string, error) {
	mergedElements := mergeElements(elements)
	b := make([]byte, 0)
	buff := bytes.NewBuffer(b)
	for _, element := range mergedElements {
		switch element := element.(type) {
		case StringElement:
			buff.WriteString(element.Content)
		case []interface{}:
			stringifiedElement, err := stringify(element)
			if err != nil {
				// no need to wrap the error again in the same function
				return "", err
			}
			buff.WriteString(stringifiedElement)
		default:
			return "", errors.Errorf("cannot convert element of type '%T' to string content", element)
		}

	}
	result := buff.String()
	for _, f := range options {
		var err error
		result, err = f(result)
		if err != nil {
			return "", errors.Wrapf(err, "Failed to postprocess the stringified content")
		}
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

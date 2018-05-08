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

// filterOption allows for filtering elements by type
type filterOption func(element interface{}) bool

// filterEmptyPremable filters the element if it is an empty preamble
func filterEmptyPremable() filterOption {
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
		defer log.Debugf(" element of type '%T' is a blankline: %t", element, result)
		return result
	}
}

// filterEmptyElements excludes the unrelevant (empty) blocks
func filterEmptyElements(blocks []interface{}, filters ...filterOption) []DocElement {
	log.Debugf("Filtering %d blocks...", len(blocks))
	elements := make([]DocElement, 0)
blocks:
	for _, block := range blocks {
		// check if filter option applies to the element
		switch block := block.(type) {
		// case BlankLine:
		// 	// exclude blank lines from here, we won't need them in the rendering anyways
		// case Preamble:
		// 	// exclude empty preambles
		// 	if len(block.Elements) > 0 {
		// 		// exclude empty preamble
		// 		elements = append(elements, block)
		// 	}
		case []interface{}:
			result := filterEmptyElements(block, filters...)
			elements = append(elements, result...)
		default:
			if block != nil {
				log.Debugf(" converting block of type '%T' into a DocElement...", block)
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
type stringifyOption func(s string) (string, error)

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
		case BlankLine:
			buff.WriteString("\n\n")
		case Paragraph:
			for _, line := range element.Lines {
				el := make([]interface{}, len(line.Elements))
				for i, e := range line.Elements {
					el[i] = e
				}
				s, err := stringify(el, options...)
				if err != nil {
					return "", errors.Errorf("cannot convert element of type '%T' to string content", element)
				}
				buff.WriteString(s)
			}
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

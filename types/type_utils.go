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
		switch element := element.(type) {
		case DocElement:
			result[i] = element
		default:
			return nil, errors.Errorf("unexpected element type: %v (expected a DocElement instead)", reflect.TypeOf(element))
		}
	}
	return result, nil
}

func merge(elements []interface{}, extraElements ...interface{}) []interface{} {
	result := make([]interface{}, 0)
	allElements := append(elements, extraElements...)
	// log.Debugf("Merging %d element(s):", len(allElements))
	buff := bytes.NewBuffer(make([]byte, 0))
	for _, v := range allElements {
		if v == nil {
			continue
		}
		switch v.(type) {
		case string:
			buff.WriteString(v.(string))
		case []byte:
			for _, b := range v.([]byte) {
				buff.WriteByte(b)
			}
		case StringElement:
			content := v.(StringElement).Content
			buff.WriteString(content)
		case *StringElement:
			content := v.(*StringElement).Content
			buff.WriteString(content)
		case []interface{}:
			w := v.([]interface{})
			if len(w) > 0 {
				f := merge(w)
				result, buff = appendBuffer(result, buff)
				result = merge(result, f...)
			}
		default:
			result, buff = appendBuffer(result, buff)
			result = append(result, v.(DocElement))
		}
	}
	// if buff was filled because some text was found
	result, buff = appendBuffer(result, buff)
	// if len(extraElements) > 0 {
	// 	log.Debugf("merged '%v' (len=%d) with '%v' (len=%d) -> '%v' (len=%d)", elements, len(elements), extraElements, len(extraElements), result, len(result))

	// } else {
	// 	log.Debugf("merged '%v' (len=%d) -> '%v' (len=%d)", elements, len(elements), result, len(result))
	// }
	return result
}

// appendBuffer appends the content of the given buffer to the given array of elements, and returns a new buffer, or returns
// the given arguments if the buffer was empty
func appendBuffer(elements []interface{}, buff *bytes.Buffer) ([]interface{}, *bytes.Buffer) {
	if buff.Len() > 0 {
		return append(elements, NewStringElement(buff.String())), bytes.NewBuffer(make([]byte, 0))
	}
	return elements, buff
}

func stringify(elements []interface{}) (*string, error) {
	mergedElements := merge(elements)
	b := make([]byte, 0)
	buff := bytes.NewBuffer(b)
	for _, element := range mergedElements {
		if element == nil {
			continue
		}
		log.Debugf("%v (%s) ", element, reflect.TypeOf(element))
		switch element := element.(type) {
		case string:
			buff.WriteString(element)
		case []byte:
			for _, b := range element {
				buff.WriteByte(b)
			}
		case StringElement:
			buff.WriteString(element.Content)
		case *StringElement:
			buff.WriteString(element.Content)
		case []interface{}:
			stringifiedElement, err := stringify(element)
			if err != nil {
				// no need to wrap the error again in the same function
				return nil, err
			}
			buff.WriteString(*stringifiedElement)
		default:
			return nil, errors.Errorf("cannot convert element of type '%v' to string content", reflect.TypeOf(element))
		}

	}
	result := buff.String()
	log.Debugf("stringified %v -> '%s' (%v)", elements, result, reflect.TypeOf(result))
	return &result, nil
}

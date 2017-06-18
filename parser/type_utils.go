package parser

import (
	"bytes"
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func flatten(content []interface{}) []interface{} {
	elements := make([]interface{}, 0)
	buff := bytes.NewBuffer(make([]byte, 0))
	for _, v := range content {
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
		case []interface{}:
			w := v.([]interface{})
			if len(w) > 0 {
				f := flatten(w)
				elements, buff = appendBuffer(elements, buff)
				elements = join(elements, f)
			}
		default:
			elements, buff = appendBuffer(elements, buff)
			elements = append(elements, v)
		}
	}
	// if buff was filled because some text was found
	elements, buff = appendBuffer(elements, buff)
	log.Debug(fmt.Sprintf("flattened '%v' (len=%d)-> '%v' (len=%d)", content, len(content), elements, len(elements)))
	return elements
}

// appendBuffer appends the content of the given buffer to the given array of elements, and returns a new buffer, or returns
// the given arguments if the buffer was empty
func appendBuffer(elements []interface{}, buff *bytes.Buffer) ([]interface{}, *bytes.Buffer) {
	if buff.Len() > 0 {
		//buffContent := buff.String()
		return append(elements, buff.String()), bytes.NewBuffer(make([]byte, 0))
	}
	return elements, buff
}

func join(elements []interface{}, otherElements []interface{}) []interface{} {
	log.Debug(fmt.Sprintf("Joining '%v' with '%v'", elements, otherElements))
	result := make([]interface{}, 0)
	allElements := append(elements, otherElements...)
	buff := bytes.NewBuffer(make([]byte, 0))
	for _, element := range allElements {
		log.Debug(fmt.Sprintf(" processing '%v' (%v)", element, reflect.TypeOf(element)))
		// if the element is not a string, then just add the current buffer then this element
		switch element.(type) {
		case string:
			buff.WriteString(element.(string))
		default:
			if buff.Len() > 0 {
				result = append(result, buff.String())
			}
			result = append(result, element)
			// re-init the buffer for the subsequent string element
			buff = bytes.NewBuffer(make([]byte, 0))
		}
	}
	// don't forget to append the pending buffer, too
	if buff.Len() > 0 {
		result = append(result, buff.String())
	}
	log.Debug(fmt.Sprintf("Join result: '%v'", result))
	return result
}

func stringify(values interface{}) string {
	valueArray := values.([]interface{})
	b := make([]byte, 0)
	buff := bytes.NewBuffer(b)
	for _, v := range valueArray {
		if v == nil {
			continue
		}
		log.Debug(fmt.Sprintf("%v (%s) ", v, reflect.TypeOf(v)))
		switch v.(type) {
		case string:
			buff.WriteString(v.(string))
		case []byte:
			for _, b := range v.([]byte) {
				buff.WriteByte(b)
			}
		case []interface{}:
			buff.WriteString(stringify(v.([]interface{})))
		default:
			log.Warn(fmt.Sprintf("unexpected type to stringify: '%v' type: %v", v, reflect.TypeOf(v)))
		}

	}
	log.Debug(fmt.Sprintf("stringified %v -> '%s' (%v)", values, buff.String(), reflect.TypeOf(buff.String())))
	return buff.String()
}

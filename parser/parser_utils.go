package parser

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/sirupsen/logrus"
)

//ParseString parses the given `content` string and returns the output document
//or the array or errors that were found while parsing the document.
func ParseString(content string) (*types.Document, []error) {
	reader := strings.NewReader(content)
	result, errs := ParseReader("", reader)
	if errs != nil {
		errors := make([]error, 0)
		for _, e := range errs.(errList) {
			logrus.Debug(fmt.Sprintf("Error found while parsing the document: %v", e.Error()))
			errors = append(errors, e)

		}
		return nil, errors
	}
	document := result.(*types.Document)
	return document, nil
}

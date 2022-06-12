package validator

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Validate validates the given document
// May also alter some attributes (eg: doctype from `manpage` to `article`)
func Validate(doc *types.Document, doctype string) ([]Problem, error) {
	if doctype == "manpage" {
		return validateManpage(doc), nil
	}
	return nil, nil
}

// Problem a problem detected during validation
// Must have a severity and an associated message
// TODO: include element position once available in the AST.
type Problem struct {
	Severity Severity
	Message  string
}

// Severity the problem severity
type Severity string

const (
	// Error the severity level for errors.
	Error Severity = "Error"
	// Warning the severity level for warning
	Warning Severity = "Warning"
)

// validateManpage checks that the document has the expected structure, ie:
// A document header
// a section named `Name` (case insensitive) with a single paragraph
// a section named `Synopsis`
//
// If the document is invalid, its doctype is set to `article` (ie, the default doctype)
func validateManpage(doc *types.Document) []Problem {
	problems := []Problem{}
	// checks the presence of a header
	if header, _ := doc.Header(); header == nil {
		problems = append(problems, Problem{
			Severity: Error,
			Message:  "manpage document is missing a header",
		})
	}
	elements := doc.BodyElements()
	if nameSection, ok := assertThatElement(elements[0]).isSection(withLevel(1), withTitle("name")); !ok {
		problems = append(problems, Problem{
			Severity: Error,
			Message:  "manpage document is missing the 'Name' section",
		})
	} else if ok := assertThatElements(nameSection.Elements).haveCount(1); !ok {
		problems = append(problems, Problem{
			Severity: Error,
			Message:  "'Name' section should contain a single paragraph",
		})
	} else if _, ok := assertThatElement(elements[1]).isSection(withLevel(1), withTitle("synopsis")); !ok {
		problems = append(problems, Problem{
			Severity: Error,
			Message:  "manpage document is missing the 'Synopsis' section",
		})
	}
	return problems
}

// assert performs a set of assertions on a given element
func assertThatElement(element interface{}) elementAssertion {
	return elementAssertion{
		element: element,
	}
}

type elementAssertion struct {
	element interface{}
}

func (e elementAssertion) isSection(assertions ...sectionAssertion) (*types.Section, bool) {
	s, ok := e.element.(*types.Section)
	if !ok {
		return nil, false
	}
	match := true
	for _, assert := range assertions {
		match = match && assert(s)
	}
	return s, match
}

type sectionAssertion func(s *types.Section) bool

func withTitle(title string) sectionAssertion {
	return func(s *types.Section) bool {
		if len(s.Title) != 1 {
			return false
		}
		str, ok := s.Title[0].(*types.StringElement)
		return ok && strings.ToLower(str.Content) == title
	}
}

func withLevel(level int) sectionAssertion {
	return func(s *types.Section) bool {
		return s.Level == level
	}
}

// assert performs a set of assertions on a given slice of elements
func assertThatElements(elements []interface{}) elementsAssertion {
	return elementsAssertion{
		elements: elements,
	}
}

type elementsAssertion struct {
	elements []interface{}
}

func (e elementsAssertion) haveCount(count int) bool {
	return len(e.elements) == count
}

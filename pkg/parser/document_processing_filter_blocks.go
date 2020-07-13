package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Filter removes all blocks that should not appear in the final document:
// - blank lines (except in delimited blocks)
// - all document attribute declaration/substitution/reset
// - empty preambles
// - single line comments and comment blocks
// - standalone attributes
func filter(elements []interface{}, matchers ...filterMatcher) []interface{} {
	result := make([]interface{}, 0, len(elements))
elements:
	for _, element := range elements {
		// check if filter option applies to the element
		for _, match := range matchers {
			if match(element) {
				log.Debugf("discarding element of type '%T'", element)
				continue elements
			}
		}
		log.Debugf("keeping element of type '%T'", element)

		// also, process the content if the element to retain
		switch e := element.(type) {
		case types.Preamble:
			e.Elements = filter(e.Elements, matchers...)
			result = append(result, e)
		case types.Paragraph:
			log.Debug("filtering on paragraph")
			lines := make([]interface{}, 0, len(e.Lines))
			for _, l := range e.Lines {
				log.Debugf("filtering on paragraph line of type '%T'", l)
				if l, ok := l.([]interface{}); ok {
					l = filter(l, matchers...)
					if len(l) > 0 {
						lines = append(lines, l)
					}
				}
			}
			e.Lines = lines
			result = append(result, e)
		case types.DelimitedBlock:
			e.Elements = filter(e.Elements, matchers...)
			result = append(result, e)
		case types.OrderedList:
			items := make([]types.OrderedListItem, 0, len(e.Items))
			for _, i := range e.Items {
				i.Elements = filter(i.Elements, matchers...)
				items = append(items, i)
			}
			e.Items = items
			result = append(result, e)
		case types.UnorderedList:
			items := make([]types.UnorderedListItem, 0, len(e.Items))
			for _, i := range e.Items {
				i.Elements = filter(i.Elements, matchers...)
				items = append(items, i)
			}
			e.Items = items
			result = append(result, e)
		case types.LabeledList:
			items := make([]types.LabeledListItem, 0, len(e.Items))
			for _, i := range e.Items {
				i.Elements = filter(i.Elements, matchers...)
				items = append(items, i)
			}
			e.Items = items
			result = append(result, e)
		default:
			result = append(result, e)
		}
	}
	return result
}

// AllMatchers all the matchers needed to remove the unneeded blocks/elements from the final document
var allMatchers = []filterMatcher{emptyPreambleMatcher, attributeMatcher, singleLineCommentMatcher, commentBlockMatcher}

// filterMatcher returns true if the given element is to be filtered out
type filterMatcher func(element interface{}) bool

// emptyPreambleMatcher filters the element if it is an empty preamble
var emptyPreambleMatcher filterMatcher = func(element interface{}) bool {
	result := false
	if p, match := element.(types.Preamble); match {
		result = p.Elements == nil || len(p.Elements) == 0
	}
	// log.Debugf(" element of type '%T' is an empty preamble: %t", element, result)
	return result
}

// attributeMatcher filters the element if it is a AttributeDeclaration,
// a AttributeSubstitution, a AttributeReset or a standalone Attribute
var attributeMatcher filterMatcher = func(element interface{}) bool {
	switch element.(type) {
	case types.AttributeDeclaration, types.AttributeSubstitution, types.AttributeReset, types.Attributes, types.CounterSubstitution:
		return true
	default:
		return false
	}
}

// singleLineCommentMatcher filters the element if it is a SingleLineComment
var singleLineCommentMatcher filterMatcher = func(element interface{}) bool {
	_, ok := element.(types.SingleLineComment)
	return ok
}

// commentBlockMatcher filters the element if it is a DelimitedBlock of kind 'Comment'
var commentBlockMatcher filterMatcher = func(element interface{}) bool {
	switch e := element.(type) {
	case types.DelimitedBlock:
		return e.Kind == types.Comment
	default:
		return false
	}
}

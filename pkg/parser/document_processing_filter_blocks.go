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
			lines := make([][]interface{}, 0, len(e.Lines))
			for _, l := range e.Lines {
				l = filter(l, matchers...)
				if len(l) > 0 {
					lines = append(lines, l)
				}
			}
			e.Lines = lines
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
var allMatchers = []filterMatcher{blankLineMatcher, emptyPreambleMatcher, documentAttributeMatcher, singleLineCommentMatcher, commentBlockMatcher, concealedIndexTermMatcher}

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

// blankLineMatcher filters the element if it is a blank line
var blankLineMatcher filterMatcher = func(element interface{}) bool {
	_, ok := element.(types.BlankLine)
	return ok
}

// documentAttributeMatcher filters the element if it is a DocumentAttributeDeclaration,
// a DocumentAttributeSubstitution or a DocumentAttributeReset
var documentAttributeMatcher filterMatcher = func(element interface{}) bool {
	switch element.(type) {
	case types.DocumentAttributeDeclaration, types.DocumentAttributeSubstitution, types.DocumentAttributeReset:
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

// concealedIndexTermMatcher filters the element if it is a ConcealedIndexTerm
var concealedIndexTermMatcher filterMatcher = func(element interface{}) bool {
	_, ok := element.(types.ConcealedIndexTerm)
	return ok
}

package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// Filter removes all blocks that should not appear in the final document:
// - blank lines (except in delimited blocks)
// - empty preambles
// - single line comments and comment blocks
func FilterOut(done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) <-chan types.DocumentFragment {
	resultStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(resultStream)
		for fragment := range fragmentStream {
			select {
			case resultStream <- filterOut(fragment, allMatchers...):
			case <-done:
				log.WithField("pipeline_task", "filter_out").Debug("received 'done' signal")
				return
			}
		}
		log.WithField("pipeline_task", "filter_out").Debug("done")
	}()
	return resultStream
}

func filterOut(f types.DocumentFragment, matchers ...filterMatcher) types.DocumentFragment {
	if err := f.Error; err != nil {
		log.Debugf("skipping filter because of fragment with error: %v", f.Error)
		return f
	}
	elements, err := doFilterOut(f.Elements, matchers...)
	if err != nil {
		return types.NewErrorFragment(f.Position, err)
	}
	return types.NewDocumentFragment(f.Position, elements...)
}

func doFilterOut(elements []interface{}, matchers ...filterMatcher) ([]interface{}, error) {
	result := make([]interface{}, 0, len(elements))
	// log.Debug("filtering elements out")
elements:
	for _, element := range elements {
		// check if filter option applies to the element
		for _, match := range matchers {
			if match(element) {
				// log.Debugf("discarding element of type '%T'", element)
				continue elements
			}
		}
		// also, process the content of the element to retain
		switch element := element.(type) {
		case *types.DelimitedBlock:
			elmts, err := doFilterOut(element.Elements, singleLineCommentMatcher, commentBlockMatcher, blanklineMatcher) // keep blanklines are not retained in delimited blocks
			if err != nil {
				return nil, err
			}
			element.Elements = elmts
		case *types.List:
			for _, elmt := range element.Elements {
				elmts, err := doFilterOut(elmt.GetElements(), matchers...)
				if err != nil {
					return nil, err
				}
				if err := elmt.SetElements(elmts); err != nil {
					return nil, err
				}
			}
		case types.WithElements:
			elmts, err := doFilterOut(element.GetElements(), matchers...)
			if err != nil {
				return nil, err
			}
			if err := element.SetElements(elmts); err != nil {
				return nil, err
			}
		}
		if e, ok := element.(types.Filterable); ok {
			if e.IsEmpty() {
				// skip element
				continue
			}
		}
		result = append(result, element)
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}

// AllMatchers all the matchers needed to remove the unneeded blocks/elements from the final document
var allMatchers = []filterMatcher{blanklineMatcher, singleLineCommentMatcher, commentBlockMatcher}

// filterMatcher returns true if the given element is to be filtered out
type filterMatcher func(element interface{}) bool

// singleLineCommentMatcher filters the element if it is a SingleLineComment
var singleLineCommentMatcher filterMatcher = func(element interface{}) bool {
	_, ok := element.(*types.SinglelineComment)
	return ok
}

// commentBlockMatcher filters the element if it is a NormalDelimitedBlock of kind 'Comment'
var commentBlockMatcher filterMatcher = func(element interface{}) bool {
	e, ok := element.(*types.DelimitedBlock)
	return ok && e.Kind == types.Comment
}

// blanklineMatcher filters the element if it is a NormalDelimitedBlock of kind 'Comment'
var blanklineMatcher filterMatcher = func(element interface{}) bool {
	_, ok := element.(*types.BlankLine)
	return ok
}

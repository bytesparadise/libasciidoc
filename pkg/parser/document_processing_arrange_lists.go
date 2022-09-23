package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ArrangeLists pipeline task which consists join list elements into lists
func ArrangeLists(done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) <-chan types.DocumentFragment {
	arrangedStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(arrangedStream)
		for fragment := range fragmentStream {
			select {
			case <-done:
				log.WithField("pipeline_task", "arrange_lists").Debug("received 'done' signal")
				return
			case arrangedStream <- arrangeLists(fragment):
			}
		}
		log.WithField("pipeline_task", "arrange_lists").Debug("done")
	}()
	return arrangedStream
}

func arrangeLists(f types.DocumentFragment) types.DocumentFragment {
	// if the fragment contains an error, then send it as-is downstream
	if err := f.Error; err != nil {
		log.Debugf("skipping list elements arrangement: %v", f.Error)
		return f
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.WithField("pipeline_task", "arrange_lists").Debugf("arranging list elements: %s", spew.Sdump(f.Elements...))
	// }
	elements, err := arrangeListElements(f.Elements)
	if err != nil {
		return types.NewErrorFragment(f.Position, err)
	}
	// result := types.NewDocumentFragment(f.Position, elements...)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.WithField("pipeline_task", "arrange_lists").Debugf("arranged lists: %s", spew.Sdump(result))
	// }
	f.Elements = elements
	return f
}

func arrangeListElements(elements []interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0, len(elements))

	for _, element := range elements {
		switch e := element.(type) {
		case *types.DelimitedBlock:
			log.Debug("checking elements in DelimitedBlock")
			var err error
			if e.Elements, err = arrangeListElements(e.Elements); err != nil {
				return nil, err
			}
			result = append(result, e)
		case *types.ListElements:
			log.Debug("arranging list elements in ListElements")
			l, err := doArrangeListElements(e.Elements)
			if err != nil {
				return nil, err
			}
			result = append(result, l)
		default:
			result = append(result, e)
		}
	}
	if len(result) == 0 {
		result = nil
	}
	return result, nil
}

func doArrangeListElements(elements []interface{}) (interface{}, error) {
	lists := newListStack() // so we can support delimited blocks in list elements, etc.

content:
	for _, element := range elements {
		// if log.IsLevelEnabled(log.DebugLevel) {
		// 	log.Debugf("arranging element of type '%T'", element)
		// }
		// lookup the parent block which can add the given element
		if parentBlock := lists.parentFor(element); parentBlock != nil {
			if err := parentBlock.AddElement(element); err != nil {
				return nil, errors.Wrap(err, "unable to assemble list elements")
			}
			continue content
		}

		switch e := element.(type) {
		case types.ListElement:
			// adjust style/level if needed
			e.AdjustStyle(lists.get())
			list, err := types.NewList(e)
			if err != nil {
				return nil, errors.Wrap(err, "unable to assemble list elements")
			}
			log.Debugf("adding a new list of kind '%s'", list.Kind)
			if err := lists.push(list); err != nil {
				return nil, errors.Wrap(err, "unable to assemble list elements")
			}
		default:
			return nil, errors.Errorf("unable to process element of type '%T' in the list", element)
		}
	}
	return lists.root(), nil
}

type listStack struct {
	stack []*types.List
}

func newListStack() *listStack {
	return &listStack{
		stack: []*types.List{},
	}
}

func (s *listStack) root() *types.List {
	if len(s.stack) > 0 {
		return s.stack[0]
	}
	return nil
}

func (s *listStack) push(l *types.List) error {
	// also append to last element following the path from `root`
	if len(s.stack) > 0 {
		if err := s.stack[len(s.stack)-1].LastElement().AddElement(l); err != nil {
			return err
		}
	}
	s.stack = append(s.stack, l)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("added list to stack (len=%d)", len(s.stack))
	// }
	return nil
}

// returns the top element of the stack without removing it
func (s *listStack) get() *types.List {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

// returns and removes the top element of the stack
func (s *listStack) pop() *types.List {
	if len(s.stack) == 0 {
		return nil
	}
	l := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return l
}

func (s *listStack) parentFor(element interface{}) *types.List {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("looking-up parent for %s", spew.Sdump(element))
	// }
	if c, ok := element.(*types.ListContinuation); ok {
		for i := 0; i < c.Offset; i++ {
			s.pop()
		}
	}
	for i := len(s.stack) - 1; i >= 0; i-- {
		// log.Debugf("checking stack at index %d", i)
		if l := s.stack[i]; l.CanAddElement(element) {
			// if match, then `pop` all
			// log.Debugf("found matching list of kind '%s' for element of type '%T' at index %d", l.Kind, element, i)
			// clears the following content off the stack
			s.stack = s.stack[:i+1]
			return l
		}
	}
	// log.Debugf("can't add element of type '%T' to any list in the stack", element)
	return nil
}

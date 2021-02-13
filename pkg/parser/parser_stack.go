package parser

type stack struct {
	index    int
	elements []interface{}
}

func newStack() *stack {
	return &stack{
		elements: make([]interface{}, 100),
		index:    -1,
	}
}

func (s *stack) size() int {
	return s.index + 1
}

func (s *stack) push(a interface{}) {
	s.index++
	s.elements[s.index] = a
}

func (s *stack) pop() interface{} {
	if s.index < 0 {
		return nil
	}
	a := s.elements[s.index]
	s.index--
	return a
}

func (s *stack) get() interface{} {
	if s.index < 0 {
		return nil
	}
	return s.elements[s.index]
}

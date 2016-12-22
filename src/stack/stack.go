package stack

type Element struct {
	next  *Element
	Value interface{}
}

type Stack struct {
	top *Element
}

func New() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Push(value interface{}) {
	e := new(Element)
	e.Value = value
	s.insert(e)
}

func (s *Stack) Pop() *Element {
	return s.remove()
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Peek() *Element {
	return s.top
}

func (s *Stack) insert(e *Element) {
	e.next = s.top
	s.top = e
}

func (s *Stack) remove() *Element {
	previously := s.top
	if s.top != nil {
		s.top = s.top.next
	}
	return previously
}

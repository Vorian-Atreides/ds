// Package stack implements a Last In First Out (LIFO):
//
//                                    ^
//                                    |
// +-------------+                    |
// |             |             +------+------+
// |      4      +--+          |             |
// |             |  |          |      4      |
// +-------------+  |          |             |
//                  |          +-------------+
//           +------v------+   +-------------+   +-------------+
//           |             |   |             |   |             |
//           |      3      |   |      3      |   |      3      |
//           |             |   |             |   |             |
//           +-------------+   +-------------+   +-------------+
//           +-------------+   +-------------+   +-------------+
//           |             |   |             |   |             |
//           |      2      |   |      2      |   |      2      |
//           |             |   |             |   |             |
//           +-------------+   +-------------+   +-------------+
//           +-------------+   +-------------+   +-------------+
//           |             |   |             |   |             |
//           |      1      |   |      1      |   |      1      |
//           |             |   |             |   |             |
//           +-------------+   +-------------+   +-------------+
package stack

// Element is an element in a single linked list
type Element struct {
	next *Element
	// The value stored in this Element
	Value interface{}
}

// Stack is an abstract data type, stacking the Element
type Stack struct {
	top *Element
}

// New allocate a new Stack structure
func New() *Stack {
	s := new(Stack)
	return s
}

// Push a new Element in the Stack
func (s *Stack) Push(value interface{}) {
	e := new(Element)
	e.Value = value
	s.insert(e)
}

// Pop retrieve and remove the first item in the top of the Stack
func (s *Stack) Pop() *Element {
	return s.remove()
}

// IsEmpty define if the Stack is Empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Peek retrieve the first item in the top of the Stack
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

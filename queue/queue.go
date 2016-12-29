// Package queue implements a First In First Out (FIFO):
// +-----+           +-----+
// |     |           |     |
// |  3  +---+       |  3  |
// |     |   |       |     |
// +-----+   |       +-----+
//        +--v--+    +-----+    +-----+
//        |     |    |     |    |     |
//        |  2  |    |  2  |    |  3  |
//        |     |    |     |    |     |
//        +-----+    +-----+    +-----+
//        +-----+    +-----+    +-----+
//        |     |    |     |    |     |
//        |  1  |    |  1  |    |  2  |
//        |     |    |     |    |     |
//        +-----+    +--+--+    +-----+
//                      |
//                      v
package queue

// Element is an element in a single linked list
type Element struct {
	prev *Element
	// The value stored in this Element
	Value interface{}
}

// Queue is an abstract datatype, retrieving the items in a chronological order
// from their insertion
type Queue struct {
	top    *Element
	bottom *Element
}

// New create a new Queue structure
func New() *Queue {
	q := new(Queue)
	return q
}

// Enqueue add an item in the top of the Queue
func (q *Queue) Enqueue(value interface{}) {
	e := new(Element)
	e.Value = value
	q.insert(e)
}

// Dequeue retrieve and remove the oldest item in the Queue
func (q *Queue) Dequeue() *Element {
	return q.remove()
}

// IsEmpty define if the Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.bottom == nil
}

// Peek retrieve the oldest item in the Queue
func (q *Queue) Peek() *Element {
	return q.bottom
}

func (q *Queue) insert(e *Element) {
	if q.top != nil {
		q.top.prev = e
	}
	q.top = e
	if q.bottom == nil {
		q.bottom = e
	}
}

func (q *Queue) remove() *Element {
	previously := q.bottom
	if q.bottom != nil {
		q.bottom = q.bottom.prev
	}
	if q.bottom == nil {
		q.top = nil
	}
	return previously
}

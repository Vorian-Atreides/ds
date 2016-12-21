package queue

type Element struct {
	prev  *Element
	Value interface{}
}

type Queue struct {
	top    *Element
	bottom *Element
}

func New() *Queue {
	q := new(Queue)
	return q
}

func (q *Queue) Enqueue(value interface{}) {
	e := new(Element)
	e.Value = value
	q.insert(e)
}

func (q *Queue) Dequeue() *Element {
	return q.remove()
}

func (q *Queue) IsEmpty() bool {
	return q.bottom == nil
}

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
	return previously
}

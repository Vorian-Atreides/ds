package dequeue

// Element is an element in a double linked list
type Element struct {
	prev *Element
	next *Element
	// The value stored in this Element
	Value interface{}
}

// Dequeue is an abstract datatype
type Dequeue struct {
	top    *Element
	bottom *Element
}

func New() *Dequeue {
	d := new(Dequeue)
	return d
}

func (d *Dequeue) IsEmpty() bool {
	return d.top == nil
}

func (d *Dequeue) Front() *Element {
	return d.top
}

func (d *Dequeue) Back() *Element {
	return d.bottom
}

func (d *Dequeue) PushBack(value interface{}) {
	e := new(Element)
	e.Value = value
	d.insertBack(e)
}

func (d *Dequeue) PushFront(value interface{}) {
	e := new(Element)
	e.Value = value
	d.insertFront(e)
}

func (d *Dequeue) PopFront() *Element {
	return d.remove(d.top)
}

func (d *Dequeue) PopBack() *Element {
	return d.remove(d.bottom)
}

func (d *Dequeue) insertFront(newElement *Element) {
	if d.top == nil {
		d.top = newElement
		d.bottom = newElement
	} else {
		d.top.prev = newElement
		newElement.next = d.top
		d.top = newElement
	}
}

func (d *Dequeue) insertBack(newElement *Element) {
	if d.top == nil {
		d.top = newElement
		d.bottom = newElement
	} else {
		d.bottom.next = newElement
		newElement.prev = d.bottom
		d.bottom = newElement
	}
}

func (d *Dequeue) remove(toRemove *Element) *Element {
	if toRemove == nil {
		return nil
	}
	if toRemove.prev == nil {
		d.top = toRemove.next
		if d.top != nil {
			d.top.prev = nil
		}
	}
	if toRemove.next == nil {
		d.bottom = toRemove.prev
		if d.bottom != nil {
			d.bottom.next = nil
		}
	}
	return toRemove
}

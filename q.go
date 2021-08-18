package q

//Q is a thread-safe (lock-free) Queue
type Q item

//Reset skip all items in queue
func (r *Q) Reset() {
	r.next.Set(nil)
}

//Push add value to tail of queue
func (r *Q) Push(value interface{}) {
	newItem := &item{
		value: value,
	}
	for {
		tail := r.getTail()
		if tail.setNext(newItem) {
			return
		}
	}
}

//Peek returns head item of queue without removing it.
//if no item found returns nil,false
func (r *Q) Peek() (value interface{}, found bool) {
	next := r.asItem().getNext()
	if next == nil {
		return nil, false
	}
	return next.value, true
}

//Skip removes head item of queue.
//returns false if no item found
func (r *Q) Skip() (found bool) {
	for {
		head := r.getHead()
		if head == nil {
			return false
		}
		if r.casHead(head, head.getNext()) {
			return true
		}
	}
}

//Pull returns head item of queue and  remove it.
//if no item found returns nil,false
func (r *Q) Pull() (value interface{}, found bool) {
	for {
		head := r.getHead()
		if head == nil {
			return nil, false
		}
		if r.casHead(head, head.getNext()) {
			return head.value, true
		}
	}
}

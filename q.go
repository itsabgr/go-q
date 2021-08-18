package q

import (
	"github.com/itsabgr/atomic2"
	. "github.com/itsabgr/go-handy"
	"unsafe"
)

type item struct {
	value Any
	next  atomic2.Ptr
}

func (r *item) setNext(another *item) bool {
	return r.next.CAS(nil, unsafe.Pointer(another))
}
func (r *item) getNext() *item {
	return (*item)(r.next.Get())
}
func (r *item) tail() *item {
	tail := r
	for {
		next := tail.getNext()
		if next == nil {
			return tail
		}
		tail = next
	}
}
func (r *item) isTail() bool {
	return r.next.Get() == nil
}

//Q is a thread-safe Queue
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
func (r *Q) getTail() *item {
	tail := r.asItem()
	for {
		if tail.isTail() {
			return tail
		}
		tail = tail.getNext()
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
func (r *Q) getHead() *item {
	return r.asItem().getNext()
}
func (r *Q) casHead(test *item, new *item) bool {
	return r.next.CAS(unsafe.Pointer(test), unsafe.Pointer(new))
}

func (r *Q) asItem() *item {
	return (*item)(r)
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

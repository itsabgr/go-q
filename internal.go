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

func (r *Q) getHead() *item {
	return r.asItem().getNext()
}
func (r *Q) casHead(test *item, new *item) bool {
	return r.next.CAS(unsafe.Pointer(test), unsafe.Pointer(new))
}

func (r *Q) asItem() *item {
	return (*item)(r)
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

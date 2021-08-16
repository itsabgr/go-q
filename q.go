package q

import (
	. "github.com/itsabgr/go-handy"
	"sync"
)

type item struct {
	value Any
	next  *item
}

type Q struct {
	mutex sync.Mutex
	head  *item
	tail  *item
}

func (r *Q) Push(value interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.push(value)
}

func (r *Q) push(value interface{}) {
	pre := r.tail
	r.tail = &item{
		value: value,
	}
	pre.next = r.tail
}

func (r *Q) Peek() (interface{}, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.Peek()
}
func (r *Q) peek() (interface{}, bool) {
	if r.head == nil {
		return nil, false
	}
	value := r.head.value
	return value, true
}
func (r *Q) Skip() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.skip()
}
func (r *Q) skip() bool {
	if r.head == nil {
		return false
	}
	r.head = r.head.next
	if r.head == nil {
		r.tail = nil
	}
	return true
}
func (r *Q) Pull() (interface{}, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.pull()
}
func (r *Q) pull() (interface{}, bool) {
	value, ok := r.peek()
	r.skip()
	return value, ok
}

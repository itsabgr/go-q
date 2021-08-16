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

//Push add value to tail of queue
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

//Peek returns head item of queue without removing it.
//if no item found returns nil,false
func (r *Q) Peek() (value interface{}, found bool) {
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

//Skip removes head item of queue.
//returns false if no item found
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

//Pull returns head item of queue and  remove it.
//if no item found returns nil,false
func (r *Q) Pull() (value interface{}, found bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.pull()
}
func (r *Q) pull() (interface{}, bool) {
	value, ok := r.peek()
	r.skip()
	return value, ok
}

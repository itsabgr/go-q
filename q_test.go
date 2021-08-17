package q

import (
	"github.com/itsabgr/go-handy"
	"testing"
)

func TestQ_Overall(t *testing.T) {
	q := Q{}
	for i := range handy.N(10) {
		q.Push(i)
		q.Skip()
	}
	for i := range handy.N(10) {
		q.Push(i)
	}
	for _ = range handy.N(10) {
		q.Skip()
	}
	c := uint(9999)
	for i := range handy.N(c) {
		q.Push(i)
	}

	for i := range handy.N(c) {
		q.Peek()
		n, ok := q.Pull()
		if !ok {
			t.Fatalf("%d(s) entries are missing", c)
		}
		if i != n {
			t.Fatalf("invalid order of queue entries; expected %d got %d", i, n)
		}
		c -= 1
	}
}

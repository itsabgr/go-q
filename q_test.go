package q

import (
	"github.com/itsabgr/atomic2"
	"github.com/itsabgr/go-handy"
	"runtime"
	"sync"
	"testing"
)

func TestQ_ThreadSafety(t *testing.T) {
	if runtime.NumCPU() <= 1 {
		t.Fatal("can not test thread-safety on this environments")
	}
	total := atomic2.Uintptr(0)
	q := Q{}
	wg := sync.WaitGroup{}
	count := uint(9999)
	goroutines := uint(runtime.NumCPU())
	for _ = range handy.N(goroutines) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range handy.N(count) {
				q.Push(total.Add(1))
			}
		}()
	}
	wg.Wait()
	n := total.Get()
	if uint(n) != count*goroutines {
		t.Fatalf("expected %d got %d", count*goroutines, n)
	}
	for _ = range handy.N(goroutines) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range handy.N(count) {
				_, ok := q.Pull()
				if !ok {
					break
				}
				total.Sub(1)
			}
		}()
	}
	wg.Wait()
	n = total.Get()
	if n != 0 {
		t.Fatalf("expcted 0 got %d", int(n))
	}
}
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

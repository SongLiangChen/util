package pqueue

import (
	"testing"
)

func TestMinPqueue(t *testing.T) {
	pq := NewPriorityQueue(false, 10)
	for i := 9; i >= 0; i-- {
		pq.Push(&Item{
			Priority: int64(i),
			Val:      i,
		})
	}

	var a []int64
	for i := 0; i < 10; i++ {
		a = append(a, pq.Pop().Priority)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != int64(i) {
			t.FailNow()
		}
	}
}

func TestMaxPqueue(t *testing.T) {
	pq := NewPriorityQueue(true, 10)
	for i := 9; i >= 0; i-- {
		pq.Push(&Item{
			Priority: int64(i),
			Val:      i,
		})
	}

	var a []int64
	for i := 0; i < 10; i++ {
		a = append(a, pq.Pop().Priority)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != int64(len(a)-i-1) {
			t.FailNow()
		}
	}
}

func TestPeekAndShift(t *testing.T) {
	pq := NewPriorityQueue(true, 10)
	for i := 9; i >= 0; i-- {
		pq.Push(&Item{
			Priority: int64(i),
			Val:      i,
		})
	}

	a, b := pq.PeekAndShift(10)
	if a != nil || b != 1 {
		t.FailNow()
	}

	a, b = pq.PeekAndShift(8)
	if a.Priority != 9 || b != 0 {
		t.FailNow()
	}

	pq = NewPriorityQueue(false, 10)
	for i := 9; i >= 0; i-- {
		pq.Push(&Item{
			Priority: int64(i),
			Val:      i,
		})
	}

	a, b = pq.PeekAndShift(10)
	if a.Priority != 0 || b != 0 {
		t.FailNow()
	}

	a, b = pq.PeekAndShift(-1)
	if a != nil || b != 2 {
		t.FailNow()
	}
}

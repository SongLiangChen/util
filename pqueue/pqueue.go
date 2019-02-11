package pqueue

type PriorityQueue struct {
	heap []*Item
	max  bool
	cap  int
}

type Item struct {
	Priority int64
	Val      interface{}
}

// NewPriorityQueue returns a PriorityQueue
// a Minimum priority queue return when max is false
// a Maximum priority queue return when max is true
func NewPriorityQueue(max bool, cap int) *PriorityQueue {
	return &PriorityQueue{
		heap: make([]*Item, 0, cap),
		max:  max,
		cap:  cap,
	}
}

func (pq *PriorityQueue) swap(i int, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *PriorityQueue) down(i int, n int) {
	for {
		index := (2 * i) + 1 // left child
		if index >= n || index < 0 {
			break
		}

		if r := index + 1; r < n && (pq.heap[r].Priority >= pq.heap[index].Priority == pq.max) { // right child
			index = r
		}

		if pq.heap[index].Priority <= pq.heap[i].Priority == pq.max {
			break
		}

		pq.swap(i, index)
		i = index
	}
}

func (pq *PriorityQueue) up(i int) {
	for {
		p := (i - 1) / 2
		if p == i {
			break
		}
		if pq.heap[p].Priority >= pq.heap[i].Priority == pq.max {
			break
		}
		pq.swap(i, p)
		i = p
	}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue) Push(i *Item) {
	n := len(pq.heap)
	c := cap(pq.heap)

	if n+1 > c {
		tmp := make([]*Item, n, 2*c)
		copy(tmp, pq.heap)
		pq.heap = tmp
	}

	pq.heap = pq.heap[:n+1]
	pq.heap[n] = i
	pq.up(n)
}

func (pq *PriorityQueue) Pop() *Item {
	n := len(pq.heap)
	if n == 0 {
		return nil
	}
	c := cap(pq.heap)

	pq.swap(0, n-1)
	pq.down(0, n-1)

	if n < c/2 && c > pq.cap {
		tmp := make([]*Item, n, c/2)
		copy(tmp, pq.heap)
		pq.heap = tmp
	}

	x := pq.heap[n-1]
	pq.heap = pq.heap[:n-1]
	return x
}

func (pq *PriorityQueue) PeekAndShift(pri int64) (*Item, int64) {
	if len(pq.heap) == 0 {
		return nil, 0
	}

	x := pq.heap[0]

	if pq.max {
		if x.Priority < pri {
			return nil, pri - x.Priority
		}
	} else {
		if x.Priority > pri {
			return nil, x.Priority - pri
		}
	}
	pq.Pop()

	return x, 0
}

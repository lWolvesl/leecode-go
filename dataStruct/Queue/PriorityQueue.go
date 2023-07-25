package Queue

type PriorityQueue[T int | float64 | float32] []T

func NewPriorityQueue[T int | float64 | float32](values []T) PriorityQueue[T] {
	return values
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x T) {
	*pq = append(*pq, x)
}

func (pq *PriorityQueue[T]) Pop() T {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

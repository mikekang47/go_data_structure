package main

type PriorityItem interface {
	Priority() int
}

type PriorityQueue[T PriorityItem] interface {
	Push(item T)
	Pop() T
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type PriorityQueueImpl[T PriorityItem] []T

func NewPriorityQueue[T PriorityItem]() *PriorityQueueImpl[T] {
	return &PriorityQueueImpl[T]{}
}
func (pq PriorityQueueImpl[PriorityItem]) Len() int {
	return len(pq)
}

func (pq PriorityQueueImpl[T]) Less(i, j int) bool {
	return pq[i].Priority() > pq[j].Priority()
}

func (pq PriorityQueueImpl[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueImpl[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueueImpl[T]) Push(item interface{}) {
	*pq = append(*pq, item.(T))
}

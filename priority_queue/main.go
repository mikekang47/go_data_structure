package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	Title    string
	priority int
}

func (t Task) Priority() int {
	return t.priority
}

func main() {
	pq := NewPriorityQueue[Task]()
	heap.Init(pq)

	t1 := Task{Title: "t1", priority: 10}
	t2 := Task{Title: "t2", priority: 15}
	t3 := Task{Title: "t3", priority: 17}

	heap.Push(pq, t1)
	heap.Push(pq, t2)
	heap.Push(pq, t3)

	for pq.Len() > 0 {
		task := heap.Pop(pq).(Task)
		fmt.Printf("Popped task: %s with priority %d\n", task.Title, task.priority)
	}
}

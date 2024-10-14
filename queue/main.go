package main

import (
	"fmt"
)

type Task struct {
	Title string
}

func main() {
	q := NewQueue[Task]()
	t := Task{"abcd"}
	q.Push(&t)

	fmt.Println(q.IsEmpty())

	a, _ := q.Peek()
	fmt.Println(*a)
}

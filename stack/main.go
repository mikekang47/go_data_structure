package main

import "fmt"

type Task struct {
	Title string
}

func main() {
	stack := NewStack[Task]()
	task := Task{Title: "abasd"}
	stack.Push(&task)

	top, _ := stack.Top()

	fmt.Println(top.Title)
	pop, _ := stack.Pop()

	fmt.Println(pop.Title)
}

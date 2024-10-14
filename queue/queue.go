package main

import (
	"container/list"
	"errors"
)

type Queue[T any] interface {
	Poll() (*T, error)
	Push(item *T)
	Peek() (*T, error)
	IsEmpty() bool
}

type QueueImpl[T any] struct {
	items *list.List
}

func NewQueue[T any]() Queue[T] {
	return &QueueImpl[T]{
		items: list.New(),
	}
}

func (q *QueueImpl[T]) Poll() (*T, error) {
	if q.items.Len() == 0 {
		return nil, nil
	}

	item := q.items.Front()
	q.items.Remove(item)

	return convert[T](item.Value)
}

func (q *QueueImpl[T]) Push(item *T) {
	q.items.PushBack(item)
}

func (q *QueueImpl[T]) Peek() (*T, error) {
	if q.items.Len() == 0 {
		return nil, errors.New("error")
	}

	return convert[T](q.items.Front().Value)
}

func (q *QueueImpl[T]) IsEmpty() bool {
	return q.items.Len() == 0
}

func convert[T any](item interface{}) (*T, error) {
	conv, ok := item.(*T) // Type assertion

	if !ok {
		return nil, errors.New("error")
	}
	return conv, nil
}

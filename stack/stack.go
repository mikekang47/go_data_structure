package main

import "errors"

type Stack[T any] interface {
	Push(item *T)
	Pop() (*T, error)
	IsEmpty() bool
	Top() (*T, error)
}

type StackImpl[T any] struct {
	items []*T
}

func NewStack[T any]() Stack[T] {
	return &StackImpl[T]{
		items: []*T{},
	}
}

func (s *StackImpl[T]) Push(item *T) {
	s.items = append(s.items, item)
}

func (s *StackImpl[T]) Pop() (*T, error) {
	if len(s.items) == 0 {
		return nil, errors.New("error")
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return convert[T](item)
}

func (s *StackImpl[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *StackImpl[T]) Top() (*T, error) {
	if len(s.items) == 0 {
		return nil, nil
	}
	return convert[T](s.items[len(s.items)-1])
}

func convert[T any](item interface{}) (*T, error) {
	conv, ok := item.(*T)
	if !ok {
		return nil, errors.New("converting failed")
	}
	return conv, nil
}

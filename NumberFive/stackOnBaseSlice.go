package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func NewEmptyStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func NewStack[T any](items []T) *Stack[T] {
	return &Stack[T]{items}
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) Pop() (T, bool) {
	if len(stack.items) == 0 {
		var zero T
		return zero, false
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return item, true
}

func (stack *Stack[T]) Peek() (T, bool) {
	if len(stack.items) == 0 {
		var zero T
		return zero, false
	}

	return stack.items[len(stack.items)-1], true
}

func PrintStack[T any](stack *Stack[T]) {
	fmt.Println(stack.items)
}

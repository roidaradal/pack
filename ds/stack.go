package ds

import "fmt"

type Stack[T any] struct {
	items List[T]
}

// NewStack creates a new empty stack
func NewStack[T any]() Stack[T] {
	return Stack[T]{items: make(List[T], 0)}
}

// NewStackFrom creates a new stack from list of items (last item = stack top)
func NewStackFrom[T any](items []T) Stack[T] {
	return Stack[T]{items: items}
}

// String returns the string representation of the stack
func (s Stack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// Len returns the number of items in the stack
func (s Stack[T]) Len() int {
	return s.items.Len()
}

// IsEmpty checks if stack is empty
func (s Stack[T]) IsEmpty() bool {
	return s.items.IsEmpty()
}

// NotEmpty checks if stack is not empty
func (s Stack[T]) NotEmpty() bool {
	return s.items.NotEmpty()
}

// Items returns the List of stack items
func (s Stack[T]) Items() List[T] {
	return s.items
}

// Push

// Pop

// MustPop

// Top

// MustTop

package genstack

import "fmt"

var ErrStackEmpty = fmt.Errorf("stack is empty")

// Stack is a stack containing any type of data.
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack with type T.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push adds an element with type T to the stack.
func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

// Pop removes and returns the top element from the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, ErrStackEmpty
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

// Peek returns the top element from the stack without removing it.
func (s *Stack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, ErrStackEmpty
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

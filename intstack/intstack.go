package intstack

import "fmt"

var ErrStackEmpty = fmt.Errorf("stack is empty")

// Stack is a stack containing uint64 data.
type Stack struct {
	data []int
}

// NewStack creates a new stack with type uint64.
func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

// Push adds an element with type uint64 to the stack.
func (s *Stack) Push(elem int) {
	s.data = append(s.data, elem)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		var zero int
		return zero, ErrStackEmpty
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

// Peek returns the top element from the stack without removing it.
func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		var zero int
		return zero, ErrStackEmpty
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

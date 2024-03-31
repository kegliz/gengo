package structstack

import "fmt"

// Stack is a stack containing a complex struct.
type (
	Item struct {
		Name string
		Age  int
	}

	Stack struct {
		data []Item
	}
)

var ErrStackEmpty = fmt.Errorf("stack is empty")

// NewStack creates a new stack with type Item.
func NewStack() *Stack {
	return &Stack{
		data: make([]Item, 0),
	}
}

// Push adds an element with type Item to the stack.
func (s *Stack) Push(elem Item) {
	s.data = append(s.data, elem)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() (Item, error) {
	if len(s.data) == 0 {
		var zero Item
		return zero, ErrStackEmpty
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

// Peek returns the top element from the stack without removing it.
func (s *Stack) Peek() (Item, error) {
	if len(s.data) == 0 {
		var zero Item
		return zero, ErrStackEmpty
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

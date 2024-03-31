package genstack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
	if _, err := s.Pop(); err == nil {
		t.Errorf("Pop() = _, nil; want _, ErrStackEmpty")
	}
	if _, err := s.Peek(); err == nil {
		t.Errorf("Peek() = _, nil; want _, ErrStackEmpty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}
	if elem, err := s.Peek(); elem != 1 || err != nil {
		t.Errorf("Peek() = %d, %v; want 1, nil", elem, err)
	}
	if elem, err := s.Pop(); elem != 1 || err != nil {
		t.Errorf("Pop() = %d, %v; want 1, nil", elem, err)
	}
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

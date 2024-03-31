package structstack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
	if _, err := s.Pop(); err == nil {
		t.Errorf("Pop() = _, nil; want _, ErrStackEmpty")
	}
	if _, err := s.Peek(); err == nil {
		t.Errorf("Peek() = _, nil; want _, ErrStackEmpty")
	}

	item := Item{"Alice", 25}
	s.Push(item)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}
	if elem, err := s.Peek(); elem != item || err != nil {
		t.Errorf("Peek() = %v, %v; want %v, nil", elem, err, item)
	}
	if elem, err := s.Pop(); elem != item || err != nil {
		t.Errorf("Pop() = %v, %v; want %v, nil", elem, err, item)
	}
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

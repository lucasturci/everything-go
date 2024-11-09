package stack

import "testing"

func TestCreate(t *testing.T) {
	s := Create[int]()
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := Create[int]()
	s.Push(10)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
	if s.Top() != 10 {
		t.Errorf("Expected top element 10, got %d", s.Top())
	}
}

func TestPop(t *testing.T) {
	s := Create[int]()
	s.Push(10)
	s.Push(20)

	val, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected popped value 20, got %d", val)
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", s.Size())
	}
}

func TestPopEmpty(t *testing.T) {
	s := Create[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
}

func TestIsEmpty(t *testing.T) {
	s := Create[int]()
	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	s.Push(10)
	if s.IsEmpty() {
		t.Error("Expected stack with element to not be empty")
	}
}

func TestTop(t *testing.T) {
	s := Create[int]()
	s.Push(10)
	s.Push(20)

	if s.Top() != 20 {
		t.Errorf("Expected top element 20, got %d", s.Top())
	}
	if s.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", s.Size())
	}
}

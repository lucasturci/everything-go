package queue

import "testing"

func TestNew(t *testing.T) {
	q := New[int]()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
	if !q.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}
}

func TestPush(t *testing.T) {
	q := New[int]()
	q.Push(10)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	if q.IsEmpty() {
		t.Error("Expected queue to not be empty after push")
	}
}

func TestPop(t *testing.T) {
	q := New[int]()
	q.Push(10)
	q.Push(20)

	val, err := q.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected popped value 10, got %d", val)
	}
	if q.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", q.Size())
	}
}

func TestPopEmpty(t *testing.T) {
	q := New[int]()
	_, err := q.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty queue")
	}
}

func TestFront(t *testing.T) {
	q := New[int]()
	q.Push(10)
	q.Push(20)

	val, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected front element 10, got %d", val)
	}
	if q.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", q.Size())
	}
}

func TestFrontEmpty(t *testing.T) {
	q := New[int]()
	_, err := q.Front()
	if err == nil {
		t.Error("Expected error when getting front of empty queue")
	}
}

func TestMultipleOperations(t *testing.T) {
	q := New[int]()

	// Push elements
	q.Push(10)
	q.Push(20)
	q.Push(30)

	// Check size
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	// Pop and verify order
	val1, _ := q.Pop()
	val2, _ := q.Pop()
	val3, _ := q.Pop()

	if val1 != 10 || val2 != 20 || val3 != 30 {
		t.Error("Queue did not maintain FIFO order")
	}

	// Should be empty now
	if !q.IsEmpty() {
		t.Error("Queue should be empty after popping all elements")
	}
}

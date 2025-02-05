package heap

import (
	"testing"
)

func TestHeapPushPop(t *testing.T) {
	h := MinHeap[int]{}

	// Test pushing elements
	h.Push(5)
	if h.Get(0) != 5 {
		t.Errorf("Expected root to be 5, got %d", h.Get(0))
	}

	h.Push(3)
	if h.Get(0) != 3 {
		t.Errorf("Expected root to be 3, got %d", h.Get(0))
	}

	h.Push(7)
	if h.Get(0) != 3 {
		t.Errorf("Expected root to be 3, got %d", h.Get(0))
	}

	// Test popping elements
	err := h.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop: %v", err)
	}
	if h.Get(0) != 5 {
		t.Errorf("Expected root to be 5 after pop, got %d", h.Get(0))
	}

	err = h.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop: %v", err)
	}
	if h.Get(0) != 7 {
		t.Errorf("Expected root to be 7 after pop, got %d", h.Get(0))
	}
}

func TestHeapEmpty(t *testing.T) {
	h := MinHeap[int]{}

	// Test pop on empty heap
	err := h.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty heap")
	}
}

func TestHeapSet(t *testing.T) {
	h := NewMinHeap[int]()

	h.Push(5)
	h.Push(3)
	h.Push(7)

	// Test setting root to larger value
	h.Set(0, 10)
	if h.Get(0) != 5 {
		t.Errorf("Expected root to be 5 after setting root to 10, got %d", h.Get(0))
	}

	// Test setting leaf to smaller value
	h.Set(2, 1)
	if h.Get(0) != 1 {
		t.Errorf("Expected root to be 1 after setting leaf to 1, got %d", h.Get(0))
	}
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap[int]()

	h.Push(5)
	h.Push(3)
	h.Push(7)

	if h.Get(0) != 7 {
		t.Errorf("Expected root of max heap to be 7, got %d", h.Get(0))
	}

	err := h.Pop()
	if err != nil {
		t.Errorf("Unexpected error on Pop: %v", err)
	}
	if h.Get(0) != 5 {
		t.Errorf("Expected root to be 5 after pop, got %d", h.Get(0))
	}
}

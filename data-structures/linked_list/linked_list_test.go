package linked_list

import (
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	ll := New[int]()
	if !ll.IsEmpty() {
		t.Error("new list should be empty")
	}
	if ll.head != nil || ll.tail != nil {
		t.Error("new list should have nil head and tail")
	}
}

func TestSize(t *testing.T) {
	ll := New[int]()
	if ll.Size() != 0 {
		t.Error("new list should have size 0")
	}

	ll.PushBack(1)
	ll.PushFront(2)

	if ll.Size() != 2 {
		t.Error("list should have size two after pushing two elements")
	}

	ll.PopBack()
	ll.PopFront()
	if ll.Size() != 0 {
		t.Error("list have size 0 after popping two elements")
	}
}

func TestIsEmpty(t *testing.T) {
	ll := New[int]()
	if !ll.IsEmpty() {
		t.Error("new list should be empty")
	}

	ll.PushBack(1)
	if ll.IsEmpty() {
		t.Error("list should not be empty after pushing element")
	}

	ll.PopBack()
	if !ll.IsEmpty() {
		t.Error("list should be empty after popping last element")
	}
}

func TestFront(t *testing.T) {
	ll := New[int]()

	// Test empty list
	_, err := ll.Front()
	if err != ErrEmpty {
		t.Errorf("expected ErrEmpty for empty list, got %v", err)
	}

	// Test single element
	ll.PushFront(1)
	val, err := ll.Front()
	if err != nil || val != 1 {
		t.Errorf("expected front value 1, got %v with error %v", val, err)
	}

	// Test multiple elements
	ll.PushFront(2)
	val, err = ll.Front()
	if err != nil || val != 2 {
		t.Errorf("expected front value 2, got %v with error %v", val, err)
	}
}

func TestBack(t *testing.T) {
	ll := New[int]()

	// Test empty list
	_, err := ll.Back()
	if err != ErrEmpty {
		t.Errorf("expected ErrEmpty for empty list, got %v", err)
	}

	// Test single element
	ll.PushBack(1)
	val, err := ll.Back()
	if err != nil || val != 1 {
		t.Errorf("expected back value 1, got %v with error %v", val, err)
	}

	// Test multiple elements
	ll.PushBack(2)
	val, err = ll.Back()
	if err != nil || val != 2 {
		t.Errorf("expected back value 2, got %v with error %v", val, err)
	}
}

func TestPushBack(t *testing.T) {
	ll := New[int]()

	// Test pushing to empty list
	ll.PushBack(1)
	front, _ := ll.Front()
	back, _ := ll.Back()
	if front != 1 || back != 1 {
		t.Errorf("expected front=1, back=1; got front=%v, back=%v", front, back)
	}

	// Test pushing to non-empty list
	ll.PushBack(2)
	front, _ = ll.Front()
	back, _ = ll.Back()
	if front != 1 || back != 2 {
		t.Errorf("expected front=1, back=2; got front=%v, back=%v", front, back)
	}
}

func TestPushFront(t *testing.T) {
	ll := New[int]()

	// Test pushing to empty list
	ll.PushFront(1)
	front, _ := ll.Front()
	back, _ := ll.Back()
	if front != 1 || back != 1 {
		t.Errorf("expected front=1, back=1; got front=%v, back=%v", front, back)
	}

	// Test pushing to non-empty list
	ll.PushFront(2)
	front, _ = ll.Front()
	back, _ = ll.Back()
	if front != 2 || back != 1 {
		t.Errorf("expected front=2, back=1; got front=%v, back=%v", front, back)
	}
}

func TestPopBack(t *testing.T) {
	ll := New[int]()

	// Test pop on empty list
	err := ll.PopBack()
	if err != ErrPopEmpty {
		t.Errorf("expected ErrPopEmpty, got %v", err)
	}

	// Test pop with single element
	ll.PushBack(1)
	err = ll.PopBack()
	if err != nil || !ll.IsEmpty() {
		t.Error("list should be empty after popping single element")
	}

	// Test pop with multiple elements
	ll.PushBack(1)
	ll.PushBack(2)
	err = ll.PopBack()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	back, _ := ll.Back()
	if back != 1 {
		t.Errorf("expected back to be 1, got %v", back)
	}
}

func TestPopFront(t *testing.T) {
	ll := New[int]()

	// Test pop on empty list
	err := ll.PopFront()
	if err != ErrPopEmpty {
		t.Errorf("expected ErrPopEmpty, got %v", err)
	}

	// Test pop with single element
	ll.PushFront(1)
	err = ll.PopFront()
	if err != nil || !ll.IsEmpty() {
		t.Error("list should be empty after popping single element")
	}

	// Test pop with multiple elements
	ll.PushFront(1)
	ll.PushFront(2)
	err = ll.PopFront()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	front, _ := ll.Front()
	if front != 1 {
		t.Errorf("expected front to be 1, got %v", front)
	}
}

func TestLinkedListIterations(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		forward  []int
		backward []int
	}{
		{
			name:     "empty list",
			elements: []int{},
			forward:  []int{},
			backward: []int{},
		},
		{
			name:     "single element",
			elements: []int{1},
			forward:  []int{1},
			backward: []int{1},
		},
		{
			name:     "multiple elements",
			elements: []int{1, 2, 3},
			forward:  []int{1, 2, 3},
			backward: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := Collect[int](slices.Values(tt.elements))

			// Test Values() iteration
			var forward []int
			for val := range ll.Values() {
				forward = append(forward, val)
			}
			if !slices.Equal(forward, tt.forward) {
				t.Errorf("Values() = %v, want %v", forward, tt.forward)
			}

			// Test Backward() iteration
			var backward []int
			for val := range ll.Backward() {
				backward = append(backward, val)
			}
			if !slices.Equal(backward, tt.backward) {
				t.Errorf("Backward() = %v, want %v", backward, tt.backward)
			}
		})
	}
}

func TestLinkedListAppendSeq(t *testing.T) {
	ll := New[int]()
	seq := []int{1, 2, 3}

	ll.AppendSeq(slices.Values(seq))

	if ll.Size() != len(seq) {
		t.Errorf("Expected size %d, got %d", len(seq), ll.Size())
	}

	i := 0
	for val := range ll.Values() {
		if val != seq[i] {
			t.Errorf("At index %d: got %d, want %d", i, val, seq[i])
		}
		i++
	}
}

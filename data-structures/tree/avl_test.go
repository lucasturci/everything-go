package tree

import (
	"maps"
	"math/rand"
	"testing"
)

func TestNewAvlTree(t *testing.T) {
	tree := NewAvlTree[int, string]()
	if tree == nil {
		t.Error("NewAvlTree returned nil")
	}
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
}

func TestNewFromSeq(t *testing.T) {
	seq := func(yield func(int, string) bool) {
		if !yield(1, "one") {
			return
		}
		if !yield(2, "two") {
			return
		}
		if !yield(3, "three") {
			return
		}
	}
	tree := NewFromSeq(seq)
	if tree.Size() != 3 {
		t.Errorf("Expected size 3, got %d", tree.Size())
	}
}

func TestAvlTreeBasicOperations(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Test Add
	err := tree.Add(5, "five")
	if err != nil {
		t.Errorf("Unexpected error in Add(): %v", err)
	}
	if tree.Size() != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size())
	}

	// Test Find
	val, err := tree.Find(5)
	if err != nil {
		t.Errorf("Unexpected error in Find(): %v", err)
	}
	if val != "five" {
		t.Errorf("Expected 'five', got %v", val)
	}

	// Test Find non-existent
	_, err = tree.Find(6)
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}

	// Test Remove
	err = tree.Remove(5)
	if err != nil {
		t.Errorf("Unexpected error in Remove(): %v", err)
	}
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after removing all elements")
	}
}

func TestAvlTreeMinMax(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Test empty tree
	_, _, err := tree.Min()
	if err != ErrEmpty {
		t.Errorf("Expected ErrEmpty for Min(), got %v", err)
	}
	_, _, err = tree.Max()
	if err != ErrEmpty {
		t.Errorf("Expected ErrEmpty for Max(), got %v", err)
	}

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test Min
	k, v, err := tree.Min()
	if err != nil {
		t.Errorf("Unexpected error in Min(): %v", err)
	}
	if k != 3 || v != "three" {
		t.Errorf("Expected (3, 'three'), got (%v, %v)", k, v)
	}

	// Test Max
	k, v, err = tree.Max()
	if err != nil {
		t.Errorf("Unexpected error in Max(): %v", err)
	}
	if k != 7 || v != "seven" {
		t.Errorf("Expected (7, 'seven'), got (%v, %v)", k, v)
	}
}

func TestAvlTreeCountOperations(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")
	tree.Add(4, "four")
	tree.Add(6, "six")

	// Test Count
	if tree.Count(5) != 1 {
		t.Errorf("Expected count 1 for key 5, got %d", tree.Count(5))
	}

	// Test CountLessThan
	if tree.CountLessThan(5) != 2 {
		t.Errorf("Expected 2 elements less than 5, got %d", tree.CountLessThan(5))
	}

	// Test CountMoreThan
	if tree.CountMoreThan(5) != 2 {
		t.Errorf("Expected 2 elements more than 5, got %d", tree.CountMoreThan(5))
	}
}

func TestAvlTreeFirstGreaterThan(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test FirstGreaterThan
	k, v, err := tree.FirstGreaterThan(4)
	if err != nil {
		t.Errorf("Unexpected error in FirstGreaterThan(): %v", err)
	}
	if k != 5 || v != "five" {
		t.Errorf("Expected (5, 'five'), got (%v, %v)", k, v)
	}

	// Test FirstGreaterOrEqualThan
	k, v, err = tree.FirstGreaterOrEqualThan(5)
	if err != nil {
		t.Errorf("Unexpected error in FirstGreaterOrEqualThan(): %v", err)
	}
	if k != 5 || v != "five" {
		t.Errorf("Expected (5, 'five'), got (%v, %v)", k, v)
	}
}

func TestAvlTreeAt(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test At
	k, v, err := tree.At(1)
	if err != nil {
		t.Errorf("Unexpected error in At(): %v", err)
	}
	if k != 5 || v != "five" {
		t.Errorf("Expected (5, 'five'), got (%v, %v)", k, v)
	}

	// Test out of bounds
	_, _, err = tree.At(3)
	if err != ErrOutOfBounds {
		t.Errorf("Expected ErrOutOfBounds, got %v", err)
	}
}

func TestAvlTreeValues(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test Values
	items := make(map[int]string)
	for k, v := range tree.Values() {
		items[k] = v
	}

	expected := map[int]string{
		3: "three",
		5: "five",
		7: "seven",
	}

	if !maps.Equal(items, expected) {
		t.Errorf("Values() = %v, want %v", items, expected)
	}
}

func TestAvlTreeBackward(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test Backward
	items := maps.Collect(tree.Backward())

	expected := map[int]string{
		7: "seven",
		5: "five",
		3: "three",
	}
	if !maps.Equal(items, expected) {
		t.Errorf("Backward() = %v, want %v", items, expected)
	}
}

func TestAvlTreeSet(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add initial value
	tree.Add(5, "five")

	// Test Set
	err := tree.Set(5, "FIVE")
	if err != nil {
		t.Errorf("Unexpected error in Set(): %v", err)
	}

	val, err := tree.Find(5)
	if err != nil {
		t.Errorf("Unexpected error in Find(): %v", err)
	}
	if val != "FIVE" {
		t.Errorf("Expected 'FIVE', got %v", val)
	}
}

func TestAvlTreeClear(t *testing.T) {
	tree := NewAvlTree[int, string]()

	// Add elements
	tree.Add(5, "five")
	tree.Add(3, "three")
	tree.Add(7, "seven")

	// Test Clear
	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}
}

func TestAvlTreeBalancing(t *testing.T) {
	tree := NewAvlTree[int, int]()

	// Test right rotation
	tree.Add(5, 5)
	tree.Add(3, 3)
	tree.Add(1, 1)

	// Test left rotation
	tree.Add(7, 7)
	tree.Add(9, 9)

	// Test left-right rotation
	tree.Add(2, 2)

	// Test right-left rotation
	tree.Add(8, 8)

	// Verify tree is balanced by checking all operations work
	items := maps.Collect(tree.Values())
	expected := map[int]int{
		1: 1, 2: 2, 3: 3, 5: 5, 7: 7, 8: 8, 9: 9,
	}
	if !maps.Equal(items, expected) {
		t.Errorf("Values() = %v, want %v", items, expected)
	}
}

func TestBasicAddAndRemove(t *testing.T) {
	tree := NewAvlTree[int, int]()
	tree.Add(10, 10)
	tree.Add(6, 6)
	tree.Add(18, 18)
	tree.Add(9, 9)
	tree.Remove(10)
	items := maps.Collect(tree.Values())
	expected := map[int]int{6: 6, 9: 9, 18: 18}
	if !maps.Equal(items, expected) {
		t.Errorf("Values() = %v, want %v", items, expected)
	}
}

func TestAvlSort(t *testing.T) {
	var iterations int = 1000
	var size int = 1000

	r := rand.New(rand.NewSource(345))
	for i := 0; i < iterations; i++ {
		tree := NewAvlTree[int, int]()
		for j := 0; j < size; j++ {
			x := r.Intn(1000000000)
			tree.Add(x, x)
		}
		items := []int{}
		for k, _ := range tree.Values() {
			items = append(items, k)
		}
		for i := 0; i+1 < len(items); i++ {
			if items[i] > items[i+1] {
				t.Errorf("Values are not sorted, items[%d] = %d, items[%d] = %d", i, items[i], i+1, items[i+1])
				break
			}
		}
	}
}

package vector

import (
	"testing"
)

func TestCreate(t *testing.T) {
	v := Create[int]()
	if v.Size() != 0 {
		t.Errorf("Expected size 0, got %d", v.Size())
	}
}

func TestCreateWithSize(t *testing.T) {
	v := CreateWithSize[int](5)
	if v.Size() != 5 {
		t.Errorf("Expected size 5, got %d", v.Size())
	}
}

func TestCreateWithCapacity(t *testing.T) {
	v := CreateWithCapacity[int](5)
	if v.Size() != 0 {
		t.Errorf("Expected size 0, got %d", v.Size())
	}
	if v.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %d", v.Capacity())
	}
}

func TestPushBack(t *testing.T) {
	v := Create[int]()
	v.PushBack(10)
	if v.Size() != 1 {
		t.Errorf("Expected size 1, got %d", v.Size())
	}
	if v[0] != 10 {
		t.Errorf("Expected element 10, got %d", v[0])
	}
}

func TestPopBack(t *testing.T) {
	v := Create[int]()
	v.PushBack(10)
	v.PopBack()
	if v.Size() != 0 {
		t.Errorf("Expected size 0, got %d", v.Size())
	}
}

func TestClear(t *testing.T) {
	v := Create[int]()
	v.PushBack(10)
	v.Clear()
	if v.Size() != 0 {
		t.Errorf("Expected size 0, got %d", v.Size())
	}
}

func TestReserve(t *testing.T) {
	tests := []struct {
		name          string
		initialCap    int
		reserveAmount int
		expectedCap   int
	}{
		{"reserve more", 0, 5, 5},
		{"reserve less", 10, 5, 10},
		{"reserve equal", 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CreateWithCapacity[int](tt.initialCap)
			v.Reserve(tt.reserveAmount)
			if v.Capacity() != tt.expectedCap {
				t.Errorf("Expected capacity %d, got %d", tt.expectedCap, v.Capacity())
			}
		})
	}
}

func TestRemove(t *testing.T) {
	v := Create[int]()
	// Add test elements
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)

	// Remove middle element
	v.Remove(1)

	if v.Size() != 2 {
		t.Errorf("Expected size 2, got %d", v.Size())
	}
	if v[0] != 1 || v[1] != 3 {
		t.Errorf("Expected elements [1,3], got [%d,%d]", v[0], v[1])
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name        string
		srcSize     int
		destSize    int
		expectError bool
	}{
		{"successful copy", 3, 3, false},
		{"destination too small", 3, 2, true},
		{"destination larger", 2, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := CreateWithSize[int](tt.srcSize)
			dest := CreateWithSize[int](tt.destSize)

			// Fill source with test data
			for i := 0; i < tt.srcSize; i++ {
				src[i] = i + 1
			}

			err := Copy(dest, src)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if err == nil {
				// Verify copied contents
				for i := 0; i < tt.srcSize; i++ {
					if dest[i] != src[i] {
						t.Errorf("Element mismatch at index %d: expected %d, got %d",
							i, src[i], dest[i])
					}
				}
			}
		})
	}
}

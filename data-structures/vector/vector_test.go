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
}

func TestPushBack(t *testing.T) {
	v := Create[int]()
	v.PushBack(10)
	if v.Size() != 1 {
		t.Errorf("Expected size 1, got %d", v.Size())
	}
	if v.Get(0) != 10 {
		t.Errorf("Expected element 10, got %d", v.Get(0))
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
	v := Create[int]()
	v.Reserve(5)
	if cap(v.elements) < 5 {
		t.Errorf("Expected capacity at least 5, got %d", cap(v.elements))
	}
}

func TestSetAndGet(t *testing.T) {
	v := CreateWithSize[int](1)
	v.Set(0, 20)
	if v.Get(0) != 20 {
		t.Errorf("Expected element 20, got %d", v.Get(0))
	}
}

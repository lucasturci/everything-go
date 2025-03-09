package tree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type BaseTree[Tk constraints.Ordered, Tv any] struct {
	lef, rig *BaseTree[Tk, Tv]
	cnt      int
	key      Tk
	val      Tv
}

type Tree[Tk constraints.Ordered, Tv any] interface {
	// Read only
	Find(key Tk) (Tv, error)
	Min() (Tk, Tv, error)
	Max() (Tk, Tv, error)
	IsEmpty() bool
	Size() int
	Traverse(func(Tk, Tv))
	Print()
	CountLessThan(Tk) int
	CountMoreThan(Tk) int
	FirstGreaterThan(Tk) (Tk, Tv, error)
	FirstGreaterOrEqualThan(Tk) (Tk, Tv, error)
	At(int) (Tk, Tv, error)

	// Write
	Set(Tk, Tv) error
	Remove(Tk) error
}

// errors
var (
	ErrNotFound    = errors.New("key not found in the tree")
	ErrEmpty       = errors.New("tree is empty")
	ErrOutOfBounds = errors.New("index is out of bounds")
)

func (t *BaseTree[Tk, Tv]) Find(key Tk) (ret Tv, err error) {
	if t == nil {
		return ret, ErrNotFound
	}

	if t.key == key {
		return t.val, nil
	} else if key > t.key {
		return t.rig.Find(key)
	}
	return t.lef.Find(key)
}

func (t *BaseTree[Tk, Tv]) Min() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.lef == nil {
		return t.key, t.val, nil
	}
	return t.lef.Min()
}

func (t *BaseTree[Tk, Tv]) Max() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.rig == nil {
		return t.key, t.val, nil
	}
	return t.rig.Max()
}

func (t *BaseTree[Tk, Tv]) IsEmpty() bool {
	return t == nil
}

func (t *BaseTree[Tk, Tv]) Size() int {
	if t == nil {
		return 0
	}
	return t.cnt
}

// Traverse traverses the tree in-order and invokes function f
func (t *BaseTree[Tk, Tv]) Traverse(f func(Tk, Tv)) {
	if t == nil {
		return
	}
	t.lef.Traverse(f)
	f(t.key, t.val)
	t.rig.Traverse(f)
}

func (t *BaseTree[Tk, Tv]) CountLessThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.key < key { // go right
		return t.lef.Size() + 1 + t.rig.CountLessThan(key)
	}
	return t.lef.CountLessThan(key)
}

func (t *BaseTree[Tk, Tv]) CountMoreThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.key > key { // go left
		return t.rig.Size() + 1 + t.lef.CountMoreThan(key)
	}
	return t.rig.CountMoreThan(key)
}

func (t *BaseTree[Tk, Tv]) printImpl() {
	// Print prints the Tree's key-value pairs using in-order traversal
	t.Traverse(func(key Tk, val Tv) {
		fmt.Printf("(%v, %v) ", key, val)
	})
}

// Print prints the Tree's key-value pairs using in-order traversal
func (t *BaseTree[Tk, Tv]) Print() {
	if t == nil {
		fmt.Println(t)
		return
	}
	t.printImpl()
	fmt.Println()
}

func (t *BaseTree[Tk, Tv]) firstGreaterThanImpl(key Tk, orEqual bool) (k Tk, v Tv, retErr error) {
	if t == nil {
		return k, v, ErrNotFound
	}
	shouldGoRight := t.key < key || (!orEqual && t.key == key)
	if shouldGoRight { // go right
		return t.rig.firstGreaterThanImpl(key, orEqual)
	}
	lk, lv, err := t.lef.firstGreaterThanImpl(key, orEqual)
	if err != nil { // that means I am the greater
		return t.key, t.val, nil
	}
	return lk, lv, nil
}

func (t *BaseTree[Tk, Tv]) FirstGreaterThan(key Tk) (Tk, Tv, error) {
	return t.firstGreaterThanImpl(key, false /*orEqual*/)
}

func (t *BaseTree[Tk, Tv]) FirstGreaterOrEqualThan(key Tk) (Tk, Tv, error) {
	return t.firstGreaterThanImpl(key, true /*orEqual*/)
}

func (t *BaseTree[Tk, Tv]) At(idx int) (k Tk, v Tv, err error) {
	if t == nil {
		return k, v, ErrOutOfBounds
	}
	sizeLeft := t.lef.Size()
	if idx == sizeLeft {
		return t.key, t.val, nil
	} else if idx > sizeLeft {
		return t.rig.At(idx - sizeLeft - 1)
	}
	return t.lef.At(idx)
}

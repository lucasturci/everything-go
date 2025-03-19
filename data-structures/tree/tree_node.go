package tree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type BaseTreeNode[Tk constraints.Ordered, Tv any] struct {
	Lef TreeNode[Tk, Tv]
	Rig TreeNode[Tk, Tv]
	Cnt int
	Key Tk
	Val Tv
}

type TreeNode[Tk constraints.Ordered, Tv any] interface {
	// Read only
	Find(key Tk) (Tv, error)
	Min() (Tk, Tv, error)
	Max() (Tk, Tv, error)
	IsEmpty() bool
	Size() int
	Traverse(func(Tk, Tv))
	Print()
	Count(Tk) int
	CountLessThan(Tk) int
	CountMoreThan(Tk) int
	firstGreaterThanImpl(Tk, bool) (Tk, Tv, error)
	FirstGreaterThan(Tk) (Tk, Tv, error)
	FirstGreaterOrEqualThan(Tk) (Tk, Tv, error)
	At(int) (Tk, Tv, error)

	// Write
	Add(Tk, Tv) (TreeNode[Tk, Tv], error)
	Set(Tk, Tv) (TreeNode[Tk, Tv], error)
	Remove(key Tk) (TreeNode[Tk, Tv], bool)
}

// errors
var (
	ErrNotFound    = errors.New("key not found in the tree")
	ErrEmpty       = errors.New("tree is empty")
	ErrOutOfBounds = errors.New("index is out of bounds")
)

// Read Only
func (t *BaseTreeNode[Tk, Tv]) Find(key Tk) (ret Tv, err error) {
	if t == nil {
		return ret, ErrNotFound
	}

	if t.Key == key {
		return t.Val, nil
	} else if key > t.Key {
		return t.Rig.Find(key)
	}
	return t.Lef.Find(key)
}

func (t *BaseTreeNode[Tk, Tv]) Min() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.Lef == nil {
		return t.Key, t.Val, nil
	}
	return t.Lef.Min()
}

func (t *BaseTreeNode[Tk, Tv]) Max() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.Rig == nil {
		return t.Key, t.Val, nil
	}
	return t.Rig.Max()
}

func (t *BaseTreeNode[Tk, Tv]) IsEmpty() bool {
	return t == nil
}

func (t *BaseTreeNode[Tk, Tv]) Size() int {
	if t == nil {
		return 0
	}
	return t.Cnt
}

// Traverse traverses the tree in-order and invokes function f
func (t *BaseTreeNode[Tk, Tv]) Traverse(f func(Tk, Tv)) {
	if t == nil {
		return
	}
	t.Lef.Traverse(f)
	f(t.Key, t.Val)
	t.Rig.Traverse(f)
}

func (t *BaseTreeNode[Tk, Tv]) Count(key Tk) int {
	return t.Size() - t.CountMoreThan(key) - t.CountLessThan(key)
}

func (t *BaseTreeNode[Tk, Tv]) CountLessThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.Key < key { // go right
		return t.Lef.Size() + 1 + t.Rig.CountLessThan(key)
	}
	return t.Lef.CountLessThan(key)
}

func (t *BaseTreeNode[Tk, Tv]) CountMoreThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.Key > key { // go left
		return t.Rig.Size() + 1 + t.Lef.CountMoreThan(key)
	}
	return t.Rig.CountMoreThan(key)
}

func (t *BaseTreeNode[Tk, Tv]) printImpl() {
	// Print prints the Tree's key-value pairs using in-order traversal
	t.Traverse(func(key Tk, val Tv) {
		fmt.Printf("(%v, %v) ", key, val)
	})
}

// Print prints the Tree's key-value pairs using in-order traversal
func (t *BaseTreeNode[Tk, Tv]) Print() {
	if t == nil {
		fmt.Println(t)
		return
	}
	t.printImpl()
	fmt.Println()
}

func (t *BaseTreeNode[Tk, Tv]) firstGreaterThanImpl(key Tk, orEqual bool) (k Tk, v Tv, retErr error) {
	if t == nil {
		return k, v, ErrNotFound
	}
	shouldGoRight := t.Key < key || (!orEqual && t.Key == key)
	if shouldGoRight { // go right
		return t.Rig.firstGreaterThanImpl(key, orEqual)
	}
	lk, lv, err := t.Lef.firstGreaterThanImpl(key, orEqual)
	if err != nil { // that means I am the greater
		return t.Key, t.Val, nil
	}
	return lk, lv, nil
}

func (t *BaseTreeNode[Tk, Tv]) FirstGreaterThan(key Tk) (Tk, Tv, error) {
	return t.firstGreaterThanImpl(key, false /*orEqual*/)
}

func (t *BaseTreeNode[Tk, Tv]) FirstGreaterOrEqualThan(key Tk) (Tk, Tv, error) {
	return t.firstGreaterThanImpl(key, true /*orEqual*/)
}

func (t *BaseTreeNode[Tk, Tv]) At(idx int) (k Tk, v Tv, err error) {
	if t == nil {
		return k, v, ErrOutOfBounds
	}
	sizeLeft := t.Lef.Size()
	if idx == sizeLeft {
		return t.Key, t.Val, nil
	} else if idx > sizeLeft {
		return t.Rig.At(idx - sizeLeft - 1)
	}
	return t.Lef.At(idx)
}

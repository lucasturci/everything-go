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

// errors
var (
	ErrNotFound = errors.New("Key not found in the tree")
	ErrEmpty    = errors.New("Tree is empty")
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

func (t *BaseTree[Tk, Tv]) MinKey() (ret Tv, err error) {
	if t == nil {
		return ret, ErrEmpty
	}
	if t.lef == nil {
		return t.val, nil
	}
	return t.lef.MinKey()
}

func (t *BaseTree[Tk, Tv]) MaxKey() (ret Tv, err error) {
	if t == nil {
		return ret, ErrEmpty
	}
	if t.rig == nil {
		return t.val, nil
	}
	return t.rig.MaxKey()
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

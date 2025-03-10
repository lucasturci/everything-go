package tree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type BaseTree[Tk constraints.Ordered, Tv any] struct {
	lef, rig Tree[Tk, Tv]
	cnt      int
	key      Tk
	val      Tv
}

type Tree[Tk constraints.Ordered, Tv any] interface {
	// Getters
	left() Tree[Tk, Tv]
	right() Tree[Tk, Tv]
	Key() Tk
	Val() Tv
	count() int

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
	Set(Tk, Tv) error
	Add(Tk, Tv) error
	Remove(Tk) error
}

// errors
var (
	ErrNotFound    = errors.New("key not found in the tree")
	ErrEmpty       = errors.New("tree is empty")
	ErrOutOfBounds = errors.New("index is out of bounds")
)

func (t *BaseTree[Tk, Tv]) left() Tree[Tk, Tv]  { return t.lef }
func (t *BaseTree[Tk, Tv]) right() Tree[Tk, Tv] { return t.rig }
func (t *BaseTree[Tk, Tv]) Key() Tk             { return t.key }
func (t *BaseTree[Tk, Tv]) Val() Tv             { return t.val }
func (t *BaseTree[Tk, Tv]) count() int          { return t.cnt }

func (t *BaseTree[Tk, Tv]) Find(key Tk) (ret Tv, err error) {
	if t == nil {
		return ret, ErrNotFound
	}

	if t.Key() == key {
		return t.Val(), nil
	} else if key > t.Key() {
		return t.right().Find(key)
	}
	return t.left().Find(key)
}

func (t *BaseTree[Tk, Tv]) Min() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.left() == nil {
		return t.Key(), t.Val(), nil
	}
	return t.left().Min()
}

func (t *BaseTree[Tk, Tv]) Max() (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.right() == nil {
		return t.Key(), t.Val(), nil
	}
	return t.right().Max()
}

func (t *BaseTree[Tk, Tv]) IsEmpty() bool {
	return t == nil
}

func (t *BaseTree[Tk, Tv]) Size() int {
	if t == nil {
		return 0
	}
	return t.count()
}

// Traverse traverses the tree in-order and invokes function f
func (t *BaseTree[Tk, Tv]) Traverse(f func(Tk, Tv)) {
	if t == nil {
		return
	}
	t.left().Traverse(f)
	f(t.Key(), t.Val())
	t.right().Traverse(f)
}

func (t *BaseTree[Tk, Tv]) Count(key Tk) int {
	return t.Size() - t.CountMoreThan(key) - t.CountLessThan(key)
}

func (t *BaseTree[Tk, Tv]) CountLessThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.Key() < key { // go right
		return t.left().Size() + 1 + t.right().CountLessThan(key)
	}
	return t.left().CountLessThan(key)
}

func (t *BaseTree[Tk, Tv]) CountMoreThan(key Tk) int {
	if t == nil {
		return 0
	}
	if t.Key() > key { // go left
		return t.right().Size() + 1 + t.left().CountMoreThan(key)
	}
	return t.right().CountMoreThan(key)
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
	shouldGoRight := t.Key() < key || (!orEqual && t.Key() == key)
	if shouldGoRight { // go right
		return t.right().firstGreaterThanImpl(key, orEqual)
	}
	lk, lv, err := t.left().firstGreaterThanImpl(key, orEqual)
	if err != nil { // that means I am the greater
		return t.Key(), t.Val(), nil
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
	sizeLeft := t.left().Size()
	if idx == sizeLeft {
		return t.Key(), t.Val(), nil
	} else if idx > sizeLeft {
		return t.right().At(idx - sizeLeft - 1)
	}
	return t.left().At(idx)
}

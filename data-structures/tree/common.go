package tree

import (
	"cmp"
	"errors"
	"fmt"
	"iter"
)

var (
	ErrNilTree     = errors.New("tree is nil")
	ErrNotFound    = errors.New("key not found in the tree")
	ErrEmpty       = errors.New("tree is empty")
	ErrOutOfBounds = errors.New("index is out of bounds")
)

type baseTree[Tk cmp.Ordered, Tv any] interface {
	// Internal
	getKey() Tk
	getVal() Tv
	getLef() baseTree[Tk, Tv]
	getRig() baseTree[Tk, Tv]
	getCnt() int
}

func find[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) (ret Tv, err error) {
	if t == nil {
		return ret, ErrNotFound
	}

	if t.getKey() == key {
		return t.getVal(), nil
	} else if key > t.getKey() {
		return find(t.getRig(), key)
	}
	return find(t.getLef(), key)
}

func minImpl[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.getLef() == nil {
		return t.getKey(), t.getVal(), nil
	}
	return minImpl(t.getLef())
}

func maxImpl[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) (key Tk, val Tv, err error) {
	if t == nil {
		return key, val, ErrEmpty
	}
	if t.getRig() == nil {
		return t.getKey(), t.getVal(), nil
	}
	return maxImpl(t.getRig())
}

func isEmpty[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) bool {
	return t == nil
}

func size[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) int {
	if t == nil {
		return 0
	}
	return t.getCnt()
}

func traverse[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], f func(Tk, Tv)) {
	if t == nil {
		return
	}
	traverse(t.getLef(), f)
	f(t.getKey(), t.getVal())
	traverse(t.getRig(), f)
}

func count[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) int {
	return size(t) - countMoreThan(t, key) - countLessThan(t, key)
}

func countLessThan[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) int {
	if t == nil {
		return 0
	}
	if t.getKey() < key { // go right
		return size(t.getLef()) + 1 + countLessThan(t.getRig(), key)
	}
	return countLessThan(t.getLef(), key)
}

func countMoreThan[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) int {
	if t == nil {
		return 0
	}
	if t.getKey() > key { // go left
		return size(t.getRig()) + 1 + countMoreThan(t.getLef(), key)
	}
	return countMoreThan(t.getRig(), key)
}

func print[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) {
	if t == nil {
		fmt.Println(t)
		return
	}
	traverse(t, func(key Tk, val Tv) {
		fmt.Printf("(%v, %v) ", key, val)
	})
	fmt.Println()
}

func firstGreaterThanImpl[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk, orEqual bool) (k Tk, v Tv, retErr error) {
	if t == nil {
		return k, v, ErrNotFound
	}
	shouldGoRight := t.getKey() < key || (!orEqual && t.getKey() == key)
	if shouldGoRight { // go right
		return firstGreaterThanImpl(t.getRig(), key, orEqual)
	}
	lk, lv, err := firstGreaterThanImpl(t.getLef(), key, orEqual)
	if err != nil { // that means I am the greater
		return t.getKey(), t.getVal(), nil
	}
	return lk, lv, nil
}

func firstGreaterThan[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) (k Tk, v Tv, retErr error) {
	return firstGreaterThanImpl(t, key, false /*orEqual*/)
}

func firstGreaterOrEqualThan[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], key Tk) (k Tk, v Tv, retErr error) {
	return firstGreaterThanImpl(t, key, true /*orEqual*/)
}

func at[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv], idx int) (k Tk, v Tv, err error) {
	if t == nil {
		return k, v, ErrOutOfBounds
	}
	sizeLeft := size(t.getLef())
	if idx == sizeLeft {
		return t.getKey(), t.getVal(), nil
	} else if idx > sizeLeft {
		return at(t.getRig(), idx-sizeLeft-1)
	}
	return at(t.getLef(), idx)
}

// Iterations
func values[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) func(yield func(Tk, Tv) bool) {
	var traverseAndYield func(t baseTree[Tk, Tv], yield func(Tk, Tv) bool)
	traverseAndYield = func(t baseTree[Tk, Tv], yield func(Tk, Tv) bool) {
		if t == nil {
			return
		}
		traverseAndYield(t.getLef(), yield)
		if !yield(t.getKey(), t.getVal()) {
			return
		}
		traverseAndYield(t.getRig(), yield)
	}
	return func(yield func(Tk, Tv) bool) {
		traverseAndYield(t, yield)
	}
}

func backward[Tk cmp.Ordered, Tv any](t baseTree[Tk, Tv]) func(yield func(Tk, Tv) bool) {
	var traverseAndYield func(t baseTree[Tk, Tv], yield func(Tk, Tv) bool)
	traverseAndYield = func(t baseTree[Tk, Tv], yield func(Tk, Tv) bool) {
		if t == nil {
			return
		}
		traverseAndYield(t.getRig(), yield)
		if !yield(t.getKey(), t.getVal()) {
			return
		}
		traverseAndYield(t.getLef(), yield)
	}
	return func(yield func(Tk, Tv) bool) {
		traverseAndYield(t, yield)
	}
}

func appendSeq[Tk cmp.Ordered, Tv any](t Tree[Tk, Tv], seq iter.Seq2[Tk, Tv]) {
	for k, v := range seq {
		t.Add(k, v)
	}
}

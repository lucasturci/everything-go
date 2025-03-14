package vector

import (
	"errors"
	"iter"
)

type Vector[T any] []T

// Constructors
func New[T any]() Vector[T] {
	return []T{}
}

func NewWithElements[T any](elements []T) Vector[T] {
	return elements
}

func NewWithSize[T any](n int) Vector[T] {
	return make([]T, n)
}

func NewWithCapacity[T any](n int) Vector[T] {
	return make([]T, 0, n)
}

// Methods
func (v Vector[T]) Size() int {
	return len(v)
}

func (v Vector[T]) IsEmpty() bool {
	return v.Size() == 0
}

func (v Vector[T]) Capacity() int {
	return cap(v)
}

func (v *Vector[T]) PushBack(element T) {
	*v = append(*v, element)
}

func (v *Vector[T]) PopBack() {
	*v = (*v)[:v.Size()-1]
}

func (v *Vector[T]) Clear() {
	*v = New[T]()
}

func (v *Vector[T]) Reserve(n int) {
	if n <= v.Capacity() {
		return
	}
	t := NewWithCapacity[T](n)
	copy(t, *v)
	*v = t
}

// Now let's implement some common snippets from this cheatsheet: https://ueokande.github.io/go-slice-tricks/

func (v *Vector[T]) Remove(i int) {
	(*v) = append((*v)[:i], (*v)[i+1:v.Size()]...)
}

func Copy[T any](dest Vector[T], src Vector[T]) error {
	if dest.Size() < src.Size() {
		return errors.New("destination vector does not have enough size")
	}
	copy(dest, src)
	return nil
}

// Iterations

func (v Vector[T]) Values() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for i := 0; i < v.Size(); i++ {
			if !yield(v[i]) {
				return
			}
		}
	}
}

func (v Vector[T]) Backward() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for i := v.Size() - 1; i >= 0; i-- {
			if !yield(v[i]) {
				return
			}
		}
	}
}

func (v *Vector[T]) AppendSeq(seq iter.Seq[T]) {
	for x := range seq {
		v.PushBack(x)
	}
}

func Collect[T any](seq iter.Seq[T]) Vector[T] {
	ans := New[T]()
	ans.AppendSeq(seq)
	return ans
}

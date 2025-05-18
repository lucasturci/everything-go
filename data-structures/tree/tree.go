package tree

import (
	"cmp"
	"iter"
)

type Tree[Tk cmp.Ordered, Tv any] interface {
	// Read only
	Find(key Tk) (Tv, error)
	Min() (Tk, Tv, error)
	Max() (Tk, Tv, error)
	IsEmpty() bool
	Size() int
	Traverse(func(Tk, Tv))
	Values() func(yield func(Tk, Tv) bool)
	Backward() func(yield func(Tk, Tv) bool)
	AppendSeq(iter.Seq2[Tk, Tv])
	Print()
	Count(Tk) int
	CountLessThan(Tk) int
	CountMoreThan(Tk) int
	FirstGreaterThan(Tk) (Tk, Tv, error)
	FirstGreaterOrEqualThan(Tk) (Tk, Tv, error)
	At(int) (Tk, Tv, error)

	// Write functions
	Add(Tk, Tv) error
	Set(Tk, Tv) error
	Remove(Tk) error
	// RemoveAll(Tk) error
	Clear()
}

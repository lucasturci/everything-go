package tree

import "golang.org/x/exp/constraints"

type Tree[Tk constraints.Ordered, Tv any] interface {
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
	FirstGreaterThan(Tk) (Tk, Tv, error)
	FirstGreaterOrEqualThan(Tk) (Tk, Tv, error)
	At(int) (Tk, Tv, error)

	// Write functions
	Add(Tk, Tv) error
	Set(Tk, Tv) error
	Remove(Tk) bool
	// RemoveAll(Tk) error
	Clear()
}

type BaseTree[Tk constraints.Ordered, Tv any] struct {
	Root TreeNode[Tk, Tv]
}

func (t BaseTree[Tk, Tv]) Find(key Tk) (Tv, error)  { return t.Root.Find(key) }
func (t BaseTree[Tk, Tv]) Min() (Tk, Tv, error)     { return t.Root.Min() }
func (t BaseTree[Tk, Tv]) Max() (Tk, Tv, error)     { return t.Root.Max() }
func (t BaseTree[Tk, Tv]) IsEmpty() bool            { return t.Root.IsEmpty() }
func (t BaseTree[Tk, Tv]) Size() int                { return t.Root.Size() }
func (t BaseTree[Tk, Tv]) Traverse(f func(Tk, Tv))  { t.Root.Traverse(f) }
func (t BaseTree[Tk, Tv]) Print()                   { t.Root.Print() }
func (t BaseTree[Tk, Tv]) Count(key Tk) int         { return t.Root.Count(key) }
func (t BaseTree[Tk, Tv]) CountLessThan(key Tk) int { return t.Root.CountLessThan(key) }
func (t BaseTree[Tk, Tv]) CountMoreThan(key Tk) int { return t.Root.CountMoreThan(key) }
func (t BaseTree[Tk, Tv]) FirstGreaterThan(key Tk) (Tk, Tv, error) {
	return t.Root.FirstGreaterThan(key)
}
func (t BaseTree[Tk, Tv]) FirstGreaterOrEqualThan(key Tk) (Tk, Tv, error) {
	return t.Root.FirstGreaterOrEqualThan(key)
}
func (t BaseTree[Tk, Tv]) At(idx int) (Tk, Tv, error) { return t.Root.At(idx) }

func (t *BaseTree[Tk, Tv]) Clear() {
	t.Root = nil
}

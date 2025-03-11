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
	firstGreaterThanImpl(Tk, bool) (Tk, Tv, error)
	FirstGreaterThan(Tk) (Tk, Tv, error)
	FirstGreaterOrEqualThan(Tk) (Tk, Tv, error)
	At(int) (Tk, Tv, error)
}

type BaseTree[Tk constraints.Ordered, Tv any] struct {
	root Tree[Tk, Tv]
}

func (t BaseTree[Tk, Tv]) Find(key Tk) (Tv, error)  { return t.root.Find(key) }
func (t BaseTree[Tk, Tv]) Min() (Tk, Tv, error)     { return t.root.Min() }
func (t BaseTree[Tk, Tv]) Max() (Tk, Tv, error)     { return t.root.Max() }
func (t BaseTree[Tk, Tv]) IsEmpty() bool            { return t.root.IsEmpty() }
func (t BaseTree[Tk, Tv]) Size() int                { return t.root.Size() }
func (t BaseTree[Tk, Tv]) Traverse(f func(Tk, Tv))  { t.root.Traverse(f) }
func (t BaseTree[Tk, Tv]) Print()                   { t.root.Print() }
func (t BaseTree[Tk, Tv]) Count(key Tk) int         { return t.root.Count(key) }
func (t BaseTree[Tk, Tv]) CountLessThan(key Tk) int { return t.root.CountLessThan(key) }
func (t BaseTree[Tk, Tv]) CountMoreThan(key Tk) int { return t.root.CountMoreThan(key) }
func (t BaseTree[Tk, Tv]) FirstGreaterThan(key Tk) (Tk, Tv, error) {
	return t.root.FirstGreaterThan(key)
}
func (t BaseTree[Tk, Tv]) FirstGreaterOrEqualThan(key Tk) (Tk, Tv, error) {
	return t.root.FirstGreaterOrEqualThan(key)
}
func (t BaseTree[Tk, Tv]) At(idx int) (Tk, Tv, error) { return t.root.At(idx) }

package tree

import (
	"cmp"
	"iter"
)

type AvlTree[Tk cmp.Ordered, Tv any] struct {
	lef, rig      *AvlTree[Tk, Tv]
	root          *AvlTree[Tk, Tv] //
	key           Tk
	val           Tv
	cnt, hei, avl int
}

var _ Tree[int, any] = &AvlTree[int, any]{}

func NewAvlTree[Tk cmp.Ordered, Tv any]() *AvlTree[Tk, Tv] {
	return &AvlTree[Tk, Tv]{}
}

func NewFromSeq[Tk cmp.Ordered, Tv any](seq iter.Seq2[Tk, Tv]) *AvlTree[Tk, Tv] {
	t := NewAvlTree[Tk, Tv]()
	t.AppendSeq(seq)
	return t
}

func (t *AvlTree[Tk, Tv]) Find(key Tk) (Tv, error)  { return find(t.getRoot(), key) }
func (t *AvlTree[Tk, Tv]) Min() (Tk, Tv, error)     { return minImpl(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) Max() (Tk, Tv, error)     { return maxImpl(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) IsEmpty() bool            { return isEmpty(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) Size() int                { return size(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) Traverse(f func(Tk, Tv))  { traverse(t.getRoot(), f) }
func (t *AvlTree[Tk, Tv]) Print()                   { print(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) Count(key Tk) int         { return count(t.getRoot(), key) }
func (t *AvlTree[Tk, Tv]) CountLessThan(key Tk) int { return countLessThan(t.getRoot(), key) }
func (t *AvlTree[Tk, Tv]) CountMoreThan(key Tk) int { return countMoreThan(t.getRoot(), key) }
func (t *AvlTree[Tk, Tv]) FirstGreaterThan(key Tk) (Tk, Tv, error) {
	return firstGreaterThan(t.getRoot(), key)
}
func (t *AvlTree[Tk, Tv]) FirstGreaterOrEqualThan(key Tk) (Tk, Tv, error) {
	return firstGreaterOrEqualThan(t.getRoot(), key)
}
func (t *AvlTree[Tk, Tv]) At(idx int) (Tk, Tv, error)              { return at(t.getRoot(), idx) }
func (t *AvlTree[Tk, Tv]) Values() func(yield func(Tk, Tv) bool)   { return values(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) Backward() func(yield func(Tk, Tv) bool) { return backward(t.getRoot()) }
func (t *AvlTree[Tk, Tv]) AppendSeq(seq iter.Seq2[Tk, Tv])         { appendSeq(t, seq) }

// Write functions
func (t *AvlTree[Tk, Tv]) Add(key Tk, val Tv) (err error) {
	if t == nil {
		return ErrNilTree
	}
	t.root, err = t.root.addImpl(key, val)
	return err
}

func (t *AvlTree[Tk, Tv]) addImpl(key Tk, val Tv) (*AvlTree[Tk, Tv], error) {
	if t == nil {
		return &AvlTree[Tk, Tv]{
			key: key,
			val: val,
			cnt: 1,
			hei: 1,
			avl: 0,
		}, nil
	}
	if key < t.key {
		n, err := t.lef.addImpl(key, val)
		if err != nil {
			return nil, err
		}
		t.lef = n
	} else {
		n, err := t.rig.addImpl(key, val)
		if err != nil {
			return nil, err
		}
		t.rig = n
	}
	return t.maybeRotate(), nil
}

// Write functions
func (t *AvlTree[Tk, Tv]) Remove(key Tk) (err error) {
	if t == nil {
		return ErrNilTree
	}
	t.root, err = t.root.removeImpl(key)
	return
}

func (t *AvlTree[Tk, Tv]) removeImpl(key Tk) (*AvlTree[Tk, Tv], error) {
	if t == nil {
		return nil, ErrNotFound
	}
	if key == t.key {
		if t.lef == nil && t.rig == nil { // leaf node, just remove it
			return nil, nil
		} else if t.lef != nil { // predecessor is leaf, so swap the values and remove predecessor
			// predecessor is the max of the left subtree
			t.lef.swapAndRemoveMax(t)
		} else { // no left child, hence only one right child exists, so swap the values and remove right child
			t.key, t.val = t.rig.key, t.rig.val
			t.rig = nil
		}
	} else if key < t.key {
		n, err := t.lef.removeImpl(key)
		if err != nil {
			return nil, err
		}
		t.lef = n
	} else {
		n, err := t.rig.removeImpl(key)
		if err != nil {
			return nil, err
		}
		t.rig = n
	}
	return t.maybeRotate(), nil
}

func (t *AvlTree[Tk, Tv]) swapAndRemoveMax(par *AvlTree[Tk, Tv]) *AvlTree[Tk, Tv] {
	if t.rig != nil {
		t.rig = t.rig.swapAndRemoveMax(par)
	} else { // no right child, hence only one left child exists, so return left child
		par.key, par.val = t.key, t.val
		t = t.lef
	}
	return t.maybeRotate()
}

func (t *AvlTree[Tk, Tv]) maybeRotate() *AvlTree[Tk, Tv] {
	var a, b, c *AvlTree[Tk, Tv]
	if t == nil {
		return nil
	}
	t.hei = max(t.lef.getHei(), t.rig.getHei()) + 1
	t.avl = t.rig.getHei() - t.lef.getHei()
	t.cnt = t.lef.getCnt() + t.rig.getCnt() + 1
	if t.avl == -2 {
		a = t
		b = t.lef
		c = t.lef.rig
		if b.avl == 1 {
			t = zigZag(a, b, c, true /*left*/)
		} else {
			t = zigZig(a, b, true /*left*/)
		}
		a.maybeRotate()
		b.maybeRotate()
		c.maybeRotate()
	} else if t.avl == 2 {
		a = t
		b = t.rig
		c = t.rig.lef
		if b.avl == -1 {
			t = zigZag(a, b, c, false /*left*/)
		} else {
			t = zigZig(a, b, false /*left*/)
		}
		a.maybeRotate()
		b.maybeRotate()
		c.maybeRotate()
	}
	return t
}

func zigZag[Tk cmp.Ordered, Tv any](a, b, c *AvlTree[Tk, Tv], left bool) *AvlTree[Tk, Tv] {
	zigZig(b, c, !left)
	if left {
		a.lef = c
	} else {
		a.rig = c
	}
	return zigZig(a, c, left)
}

func zigZig[Tk cmp.Ordered, Tv any](a, b *AvlTree[Tk, Tv], left bool) *AvlTree[Tk, Tv] {
	if left {
		a.lef = b.rig
		b.rig = a
	} else {
		a.rig = b.lef
		b.lef = a
	}
	return b
}

func (t *AvlTree[Tk, Tv]) Set(key Tk, val Tv) error {
	err := t.Remove(key)
	if err != nil {
		return err
	}
	return t.Add(key, val)
}

func (t *AvlTree[Tk, Tv]) Clear() {
	if t == nil {
		return
	}
	t.root = nil
}

func (t *AvlTree[Tk, Tv]) getKey() (k Tk) {
	if t == nil {
		return
	}
	return t.key
}
func (t *AvlTree[Tk, Tv]) getVal() (v Tv) {
	if t == nil {
		return
	}
	return t.val
}
func (t *AvlTree[Tk, Tv]) getLef() baseTree[Tk, Tv] {
	if t == nil {
		return nil
	}
	if t.lef == nil { // return nil interface, o/w we may return a non-nil interface with a nil concrete type
		return nil
	}
	return t.lef
}
func (t *AvlTree[Tk, Tv]) getRig() baseTree[Tk, Tv] {
	if t == nil {
		return nil
	}
	if t.rig == nil { // return nil interface, o/w we may return a non-nil interface with a nil concrete type
		return nil
	}
	return t.rig
}
func (t *AvlTree[Tk, Tv]) getRoot() baseTree[Tk, Tv] {
	if t == nil {
		return nil
	}
	if t.root == nil { // return nil interface, o/w we may return a non-nil interface with a nil concrete type
		return nil
	}
	return t.root
}

func (t *AvlTree[Tk, Tv]) getCnt() int {
	if t == nil {
		return 0
	}
	return t.cnt
}
func (t *AvlTree[Tk, Tv]) getHei() int {
	if t == nil {
		return 0
	}
	return t.hei
}

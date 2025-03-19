package avl

import (
	"github.com/lucasturci/everything-go/data-structures/tree"
	"golang.org/x/exp/constraints"
)

var _ tree.TreeNode[int, any] = &AvlTreeNode[int, any]{}

type AvlTreeNode[Tk constraints.Ordered, Tv any] struct {
	tree.BaseTreeNode[Tk, Tv]
	hei int
	avl int
}

func newNode[Tk constraints.Ordered, Tv any](key Tk, val Tv) *AvlTreeNode[Tk, Tv] {
	return &AvlTreeNode[Tk, Tv]{
		BaseTreeNode: tree.BaseTreeNode[Tk, Tv]{
			Key: key,
			Val: val,
			Cnt: 1,
		},
	}
}

// Write functions
func (t *AvlTreeNode[Tk, Tv]) Add(key Tk, val Tv) (tree.TreeNode[Tk, Tv], error) {
	if t == nil {
		return newNode[Tk, Tv](key, val), nil
	}
	if key < t.Key {
		n, err := t.Lef.Add(key, val)
		if err != nil {
			return nil, err
		}
		t.Lef = n
	} else {
		n, err := t.Rig.Add(key, val)
		if err != nil {
			return nil, err
		}
		t.Rig = n
	}
	return nil, nil // TBD
}

func (t *AvlTreeNode[Tk, Tv]) Set(key Tk, val Tv) (tree.TreeNode[Tk, Tv], error) {
	if t == nil {
		return newNode[Tk, Tv](key, val), nil
	}
	if key < t.Key {
		n, err := t.Lef.Set(key, val)
		if err != nil {
			return nil, err
		}
		t.Lef = n
	} else {
		n, err := t.Rig.Set(key, val)
		if err != nil {
			return nil, err
		}
		t.Rig = n
	}
	return nil, nil // TBD
}

func (t *AvlTreeNode[Tk, Tv]) Remove(key Tk) (tree.TreeNode[Tk, Tv], bool) {
	return nil, false // TBD

}

func (t *AvlTreeNode[Tk, Tv]) maybeRotate(key Tk) tree.TreeNode[Tk, Tv] {
	if t == nil {
		return nil
	}
	t.hei = max(t.Lef.getHeight())
}

func (t *AvlTreeNode[Tk, Tv]) zigZig() tree.TreeNode[Tk, Tv] {

}

func (t *AvlTreeNode[Tk, Tv]) zigZag() tree.TreeNode[Tk, Tv] {

}

func (t *AvlTreeNode[Tk, Tv]) getHeight() int {
	if t == nil {
		return 0
	}
	return t.hei
}

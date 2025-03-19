package avl

import (
	"github.com/lucasturci/everything-go/data-structures/tree"
	"golang.org/x/exp/constraints"
)

/*
	Implementation of AVL tree.
	The AVL tree is an implementation of a balanced binary search tree, where the heights of the
	two child subtrees of any node differ by at most one (https://en.wikipedia.org/wiki/AVL_tree)
*/

var _ tree.Tree[int, any] = &AvlTree[int, any]{}

type AvlTree[Tk constraints.Ordered, Tv any] struct {
	tree.BaseTree[Tk, Tv]
}

func NewAvlTree[Tk constraints.Ordered, Tv any]() *AvlTree[Tk, Tv] {
	return &AvlTree[Tk, Tv]{
		tree.BaseTree[Tk, Tv]{
			Root: nil,
		},
	}
}

func (t *AvlTree[Tk, Tv]) Add(key Tk, val Tv) error {
	newRoot, err := t.BaseTree.Root.Add(key, val)
	if err != nil {
		return err
	}
	t.Root = newRoot
	return nil
}

func (t *AvlTree[Tk, Tv]) Set(key Tk, val Tv) error {
	newRoot, err := t.BaseTree.Root.Set(key, val)
	if err != nil {
		return err
	}
	t.Root = newRoot
	return nil
}

func (t *AvlTree[Tk, Tv]) Remove(key Tk) bool {
	newRoot, removed := t.BaseTree.Root.Remove(key)
	t.BaseTree.Root = newRoot
	return removed
}

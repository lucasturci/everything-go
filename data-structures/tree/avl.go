package tree

import "golang.org/x/exp/constraints"

/*
	Implementation of AVL tree.
	The AVL tree is an implementation of a balanced binary search tree, where the heights of the
	two child subtrees of any node differ by at most one (https://en.wikipedia.org/wiki/AVL_tree)
*/

type AvlTree[Tk constraints.Ordered, Tv any] struct {
	BaseTree[Tk, Tv]
}

func NewAvlTree[Tk constraints.Ordered, Tv any]() *AvlTree[Tk, Tv] {
	return nil
}

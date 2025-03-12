package avl

import (
	"github.com/lucasturci/everything-go/data-structures/tree"
	"golang.org/x/exp/constraints"
)

var _ tree.TreeNode[int, any] = &AvlTreeNode[int, any]{}

type AvlTreeNode[Tk constraints.Ordered, Tv any] struct {
	tree.BaseTreeNode[Tk, Tv]
}

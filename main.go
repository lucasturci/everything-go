package main

import (
	"fmt"
	"iter"
	"slices"

	"github.com/lucasturci/everything-go/data-structures/linked_list"
	"github.com/lucasturci/everything-go/data-structures/types"
)

func All(ar []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < len(ar); i++ {
			if !yield(ar[i]) {
				return
			}
		}
	}
}

func main() {
	x := linked_list.New[int]()
	x.PushBack(1)
	x.PushBack(2)
	fmt.Println(types.Container(x).Size())

	a := []int{1, 3, 2}
	fmt.Println(slices.Sorted(All(a)))
}

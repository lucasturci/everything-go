package priority_queue

import (
	"cmp"

	"github.com/lucasturci/everything-go/data-structures/comparator"
	"github.com/lucasturci/everything-go/data-structures/heap"
)

type PriorityQueue[T cmp.Ordered] struct {
	heap.MaxHeap[T]
}

type PriorityQueueCustom[T any, C comparator.Comparator[T]] struct {
	heap.Heap[T, C]
}

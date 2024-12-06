package priority_queue

import (
	"cmp"
	"everything-go/data-structures/comparator"
	"everything-go/data-structures/heap"
)

type PriorityQueue[T cmp.Ordered] struct {
	heap.MaxHeap[T]
}

type PriorityQueueCustom[T any, C comparator.Comparator[T]] struct {
	heap.Heap[T, C]
}

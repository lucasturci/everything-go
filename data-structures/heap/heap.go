package heap

import (
	"errors"
	alg "everything-go/algorithms"
	"everything-go/data-structures/comparator"
	"everything-go/data-structures/vector"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered, C comparator.Comparator[T]] struct {
	vector.Vector[T]
	cmp C
}

type MinHeap[T constraints.Ordered] struct {
	Heap[T, comparator.Less[T]]
}
type MaxHeap[T constraints.Ordered] struct {
	Heap[T, comparator.Greater[T]]
}

func Create[T constraints.Ordered, C comparator.Comparator[T]]() Heap[T, C] {
	return Heap[T, C]{}
}

func CreateWithCapacity[T constraints.Ordered, C comparator.Comparator[T]](capacity int) Heap[T, C] {
	return Heap[T, C]{Vector: vector.CreateWithCapacity[T](capacity)}
}

// Heapify fixes the subtree rooted at i
func (h *Heap[T, C]) Heapify(i int) {
	for (i<<1)+1 < h.Size() {
		l := (i << 1) + 1
		r := l + 1
		if r < h.Size() && h.cmp.Less(h.Vector[r], h.Vector[l]) { // make l be the smallest
			alg.Swap(&l, &r)
		}

		if h.cmp.Less(h.Vector[l], h.Vector[i]) {
			alg.Swap(&h.Vector[i], &h.Vector[l])
			i = l
		} else {
			break
		}
	}
}

// BubbleUp bubbles the element at i up to its correct position
func (h *Heap[T, C]) BubbleUp(i int) {
	for ; i > 0; i = (i >> 1) {
		if h.cmp.Less(h.Vector[i], h.Vector[(i-1)>>1]) {
			alg.Swap(&h.Vector[i], &h.Vector[(i-1)>>1])
		} else {
			break
		}
	}
}

func (h *Heap[T, C]) Push(x T) {
	h.PushBack(x)
	h.BubbleUp(h.Size() - 1)
}

func (h *Heap[T, C]) Pop() error {
	if h.Size() == 0 {
		return errors.New("Empty heap")
	}
	alg.Swap(&h.Vector[0], &h.Vector[h.Size()-1])
	h.PopBack()
	if h.Size() > 0 {
		h.Heapify(0)
	}

	return nil
}

func (h *Heap[T, C]) Get(i int) T {
	return h.Vector[i]
}

func (h *Heap[T, C]) Set(i int, x T) {
	h.Vector[i] = x
	h.Heapify(i)
	h.BubbleUp(i)
}

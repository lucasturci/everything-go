package heap

import (
	"errors"

	alg "github.com/lucasturci/everything-go/algorithms"
	"github.com/lucasturci/everything-go/data-structures/comparator"
	"github.com/lucasturci/everything-go/data-structures/vector"

	"cmp"
)

var ErrEmptyHeap = errors.New("empty heap")

type Heap[T any, C comparator.Comparator[T]] struct {
	vector.Vector[T]
	cmp C
}

type MinHeap[T cmp.Ordered] struct {
	Heap[T, comparator.Less[T]]
}

func NewMinHeap[T cmp.Ordered]() MinHeap[T] {
	return MinHeap[T]{}
}

type MaxHeap[T cmp.Ordered] struct {
	Heap[T, comparator.Greater[T]]
}

func NewMaxHeap[T cmp.Ordered]() MaxHeap[T] {
	return MaxHeap[T]{}
}

func New[T any, C comparator.Comparator[T]]() Heap[T, C] {
	return Heap[T, C]{}
}

func NewWithCapacity[T any, C comparator.Comparator[T]](capacity int) Heap[T, C] {
	return Heap[T, C]{Vector: vector.NewWithCapacity[T](capacity)}
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
	for ; i > 0; i = (i - 1) >> 1 {
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

func (h Heap[T, C]) Top() (val T, err error) {
	if h.Size() == 0 {
		return val, ErrEmptyHeap
	}
	return h.Get(0), nil
}

func (h *Heap[T, C]) Pop() error {
	if h.Size() == 0 {
		return ErrEmptyHeap
	}
	alg.Swap(&h.Vector[0], &h.Vector[h.Size()-1])
	h.PopBack()
	if h.Size() > 0 {
		h.Heapify(0)
	}

	return nil
}

func (h Heap[T, C]) Get(i int) T {
	return h.Vector[i]
}

func (h *Heap[T, C]) Set(i int, x T) {
	h.Vector[i] = x
	h.Heapify(i)
	h.BubbleUp(i)
}

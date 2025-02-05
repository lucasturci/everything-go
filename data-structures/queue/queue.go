package queue

import (
	"errors"

	"github.com/lucasturci/everything-go/data-structures/stack"
)

// Write queue with two stacks.

type Queue[T any] struct {
	r stack.Stack[T]
	l stack.Stack[T]
}

func New[T any]() Queue[T] {
	return Queue[T]{
		r: stack.New[T](),
		l: stack.New[T](),
	}
}

func (q *Queue[T]) Push(x T) {
	q.r.Push(x)
}

func (q Queue[T]) IsEmpty() bool {
	return q.l.IsEmpty() && q.r.IsEmpty()
}

func (q Queue[T]) Size() int {
	return q.l.Size() + q.r.Size()
}

func (q *Queue[T]) flush() {
	q.l = stack.Reverse(q.r)
	q.r = stack.New[T]()
}

func (q *Queue[T]) Pop() (ret T, err error) {
	if q.IsEmpty() {
		return ret, errors.New("queue is empty")
	}
	if q.l.IsEmpty() {
		q.flush()
	}
	return q.l.Pop()
}

func (q Queue[T]) Front() (ret T, err error) {
	if q.IsEmpty() {
		return ret, errors.New("queue is empty")
	}
	if q.l.IsEmpty() {
		q.flush()
	}
	return q.l.Top()
}

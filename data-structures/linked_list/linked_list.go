package linked_list

import "errors"

var (
	ErrEmpty    = errors.New("linked list is empty")
	ErrPopEmpty = errors.New("trying to pop an empty linked list")
)

type node[T any] struct {
	left, right *node[T]
	val         T
}

func newNode[T any](val T) node[T] {
	return node[T]{
		val: val,
	}
}

type LinkedList[T any] struct {
	head, tail *node[T]
	size       int
}

func New[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (ll LinkedList[T]) Front() (ret T, err error) {
	if ll.IsEmpty() {
		return ret, ErrEmpty
	}

	return ll.head.val, nil
}

func (ll LinkedList[T]) Back() (ret T, err error) {
	if ll.IsEmpty() {
		return ret, ErrEmpty
	}

	return ll.tail.val, nil
}

func (ll LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

func (ll LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) PushBack(val T) {
	x := newNode(val)
	ll.size++
	if ll.IsEmpty() {
		ll.head = &x
		ll.tail = &x
		return
	}

	x.left = ll.tail
	ll.tail.right = &x
	ll.tail = &x
}

func (ll *LinkedList[T]) PushFront(val T) {
	x := newNode(val)
	ll.size++
	if ll.IsEmpty() {
		ll.head = &x
		ll.tail = &x
		return
	}

	x.right = ll.head
	ll.head.left = &x
	ll.head = &x
}

func (ll *LinkedList[T]) PopBack() error {
	if ll.IsEmpty() {
		return ErrPopEmpty
	}
	ll.size--

	if ll.tail.left != nil {
		ll.tail = ll.tail.left
		ll.tail.right = nil
	} else {
		ll.head = nil
		ll.tail = nil
	}

	return nil
}

func (ll *LinkedList[T]) PopFront() error {
	if ll.IsEmpty() {
		return ErrPopEmpty
	}
	ll.size--

	if ll.head.right != nil {
		ll.head = ll.head.right
		ll.head.left = nil
	} else {
		ll.head = nil
		ll.tail = nil
	}

	return nil
}

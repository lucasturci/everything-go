package stack

import "errors"

type stack[T any] []T

func Create[T any]() stack[T] {
	return stack[T]{}
}

func (s *stack[T]) Size() int {
	return len(*s)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *stack[T]) Pop() (ret T, err error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	ret = s.Top()
	*s = (*s)[:s.Size()-1]
	return
}

func (s *stack[T]) Top() T {
	return (*s)[s.Size()-1]
}

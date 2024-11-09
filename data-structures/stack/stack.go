package stack

import "errors"

type stack[T any] []T

// Constructors
func Create[T any]() stack[T] {
	return stack[T]{}
}

// Methods
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
	ret, err = s.Top()
	if err != nil {
		return *new(T), err
	}
	*s = (*s)[:s.Size()-1]
	return
}

func (s *stack[T]) Top() (ret T, err error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return (*s)[s.Size()-1], nil
}

// Methods
func Reverse[T any](s stack[T]) stack[T] {
	ret := Create[T]()
	for i := s.Size() - 1; i >= 0; i-- {
		ret.Push(s[i])
	}
	return ret
}

package stack

import "errors"

type Stack[T any] []T

// Constructors
func Create[T any]() Stack[T] {
	return Stack[T]{}
}

// Methods
func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Pop() (ret T, err error) {
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

func (s *Stack[T]) Top() (ret T, err error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	return (*s)[s.Size()-1], nil
}

// Functions
func Reverse[T any](s Stack[T]) Stack[T] {
	ret := Create[T]()
	for i := s.Size() - 1; i >= 0; i-- {
		ret.Push(s[i])
	}
	return ret
}

package comparator

import "cmp"

type Comparator[T any] interface {
	Less(a, b T) bool // returns true if a < b, false otherwise
}

// Predefined comparators
type Less[T cmp.Ordered] struct{}

func (l Less[T]) Less(a, b T) bool {
	return a < b
}

type Greater[T cmp.Ordered] struct{}

func (l Greater[T]) Less(a, b T) bool {
	return a > b
}

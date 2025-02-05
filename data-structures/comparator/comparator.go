package comparator

import (
	"cmp"
	"fmt"
)

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

type CustomComparator[T any] struct {
	customLess func(a, b T) bool
}

func (cc CustomComparator[T]) Less(a, b T) bool {
	if cc.customLess == nil {
		panic(fmt.Sprintf("Attempting to call customLess of CustomComparator %v but it is nil", cc))
	}
	return cc.customLess(a, b)
}

func Custom[T any](customLess func(a, b T) bool) CustomComparator[T] {
	return CustomComparator[T]{customLess}
}

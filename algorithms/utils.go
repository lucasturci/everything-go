package algorithms

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

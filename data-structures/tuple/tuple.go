package tuple

// Pair type

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// Triple type

type Triple[T1 any, T2 any, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

// Tuple types

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
	Fifth  T5
}

// That's enough

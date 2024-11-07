package vector

type Vector[T any] struct {
	elements []T
	size     int
}

// Constructors
func Create[T any]() *Vector[T] {
	return &Vector[T]{
		elements: []T{},
		size:     0,
	}
}

func CreateWithSize[T any](n int) *Vector[T] {
	return &Vector[T]{
		elements: make([]T, n),
		size:     n,
	}
}

func CreateWithCapacity[T any](n int) *Vector[T] {
	return &Vector[T]{
		elements: make([]T, 0, n),
		size:     0,
	}
}

// Accessors
func (v *Vector[T]) Get(index int) T {
	return v.elements[index]
}

func (v *Vector[T]) Set(index int, element T) {
	v.elements[index] = element
}

// Methods
func (v *Vector[T]) Size() int {
	return v.size
}

func (v *Vector[T]) PushBack(element T) {
	v.elements = append(v.elements, element)
	v.size++
}

func (v *Vector[T]) PopBack() {
	v.elements = v.elements[:len(v.elements)-1]
	v.size--
}

func (v *Vector[T]) Clear() {
	v.elements = []T{}
	v.size = 0
}

func (v *Vector[T]) Reserve(n int) {
	v.elements = append(v.elements, make([]T, n)...)
}

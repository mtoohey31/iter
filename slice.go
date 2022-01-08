package iter

type sliceInner[T any] struct {
	index int
	slice []T
	zero  T // TODO: find a less gross way to produce a zero value of a generic type
}

func FromSlice[T any](s []T) *Iter[T] {
	return WithInner[T](&sliceInner[T]{index: 0, slice: s})
}

func (i *sliceInner[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *sliceInner[T]) Next() (T, error) {
	if i.index >= len(i.slice) {
		return i.zero, IteratorExhaustedError
	}

	defer func() { i.index = i.index + 1 }()
	return i.slice[i.index], nil
}

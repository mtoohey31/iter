package iter

type sliceInner[T any] struct {
	index int
	slice []T
}

// Elems returns an iterator over the values of the provided slice.
func Elems[T any](s []T) *Iter[T] {
	return WithInner[T](&sliceInner[T]{index: 0, slice: s})
}

func (i *sliceInner[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *sliceInner[T]) Next() (T, error) {
	if i.index >= len(i.slice) {
		var z T
		return z, IteratorExhaustedError
	}

	res := i.slice[i.index]
	i.index = i.index + 1
	return res, nil
}

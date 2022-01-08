package iter

import "errors"

type innerIter[T any] interface {
	HasNext() bool
	Next() (T, error)
}

type Iter[T any] struct {
	inner innerIter[T]
}

func WithInner[T any](inner innerIter[T]) *Iter[T] {
	return &Iter[T]{inner: inner}
}

func (i *Iter[T]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *Iter[T]) Next() (T, error) {
	return i.inner.Next()
}

var IteratorExhaustedError = errors.New("iterator exhausted")

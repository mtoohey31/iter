package iter

import "errors"

var IteratorExhaustedError = errors.New("iterator exhausted")

type innerIter[T any] interface {
	HasNext() bool
	Next() (T, error)
}

type Iter[T any] struct {
	inner innerIter[T]
}

func (i *Iter[T]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *Iter[T]) Next() (T, error) {
	return i.inner.Next()
}

func WithInner[T any](inner innerIter[T]) *Iter[T] {
	return &Iter[T]{inner: inner}
}

func (i *Iter[T]) Collect() []T {
	var res []T
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		res = append(res, next)
	}
	return res
}

func (i *Iter[T]) All(f func(T) bool) bool {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if !f(next) {
			return false
		}
	}
	return true
}

func (i *Iter[T]) Any(f func(T) bool) bool {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if f(next) {
			return true
		}
	}
	return false
}

func (i *Iter[T]) Count() int {
	j := 0
	for {
		_, err := i.Next()

		if err != nil {
			break
		}

		j++
	}
	return j
}

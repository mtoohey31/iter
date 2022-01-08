package iter

import (
	"errors"
)

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

// TODO: find a nicer way to fetch zero values of a generic type
type z[T any] struct{ z T }

func (i Iter[T]) zeroVal() T {
	return z[T]{}.z
}

type emptyInner[T any] struct{}

func (i *emptyInner[T]) HasNext() bool {
	return false
}

func (i *emptyInner[T]) Next() (T, error) {
	return Iter[T]{}.zeroVal(), IteratorExhaustedError
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

func (i *Iter[T]) Find(f func(T) bool) (T, error) {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if f(next) {
			return next, nil
		}
	}

	return i.zeroVal(), errors.New("no element found")
}

func (i *Iter[T]) FoldSame(init T, f func(curr T, next T) T) T {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

func Fold[T any, U any](i *Iter[T], init U, f func(curr U, next T) U) U {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

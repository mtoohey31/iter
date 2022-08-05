package iter

import (
	"github.com/barweiss/go-tuple"
	"golang.org/x/exp/constraints"
)

// Min returns the minimum value in the provided iterator.
func Min[T constraints.Ordered](oi Iter[T]) (T, bool) {
	return oi.Reduce(func(curr T, next T) T {
		if curr < next {
			return curr
		}

		return next
	})
}

// MinByKey returns the value with the minimum result after the application of
// the provided function.
func MinByKey[T any, U constraints.Ordered](i Iter[T], key func(T) U) (T, bool) {
	init, ok := i()

	if !ok {
		var z T
		return z, false
	}

	return Fold(i, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 < keyNext {
			return curr
		}

		return tuple.New2(next, keyNext)
	}).V1, true
}

// Max returns the maximum value in the provided iterator.
func Max[T constraints.Ordered](oi Iter[T]) (T, bool) {
	return oi.Reduce(func(curr T, next T) T {
		if curr > next {
			return curr
		}

		return next
	})
}

// MaxByKey returns the value with the maximum result after the application of
// the provided function.
func MaxByKey[T any, U constraints.Ordered](i Iter[T], key func(T) U) (T, bool) {
	init, ok := i()

	if !ok {
		var z T
		return z, false
	}

	return Fold(i, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 > keyNext {
			return curr
		}

		return tuple.New2(next, keyNext)
	}).V1, true
}

// Sum returns the sum of all the values in the provided iterator.
func Sum[T constraints.Ordered](oi Iter[T]) T {
	var z T
	return oi.FoldEndo(z, func(curr, next T) T { return curr + next })
}

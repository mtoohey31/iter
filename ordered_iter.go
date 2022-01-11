package iter

import "github.com/barweiss/go-tuple"

// Min returns the minimum value in the provided iterator.
func Min[T ordered](oi *Iter[T]) (T, error) {
	return oi.Reduce(func(curr T, next T) T {
		if curr < next {
			return curr
		} else {
			return next
		}
	})
}

// MinByKey returns the value with the minimum result after the application of
// the provided function.
func MinByKey[T any, U ordered](oi *Iter[T], key func(T) U) (T, error) {
	init, err := oi.Next()

	if err != nil {
		var z T
		return z, IteratorExhaustedError
	}

	return Fold(oi, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 < keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

// Max returns the maximum value in the provided iterator.
func Max[T ordered](oi *Iter[T]) (T, error) {
	return oi.Reduce(func(curr T, next T) T {
		if curr > next {
			return curr
		} else {
			return next
		}
	})
}

// MaxByKey returns the value with the maximum result after the application of
// the provided function.
func MaxByKey[T any, U ordered](oi *Iter[T], key func(T) U) (T, error) {
	init, err := oi.Next()

	if err != nil {
		var z T
		return z, IteratorExhaustedError
	}

	return Fold(oi, tuple.New2(init, key(init)), func(curr tuple.T2[T, U], next T) tuple.T2[T, U] {
		keyNext := key(next)
		if curr.V2 > keyNext {
			return curr
		} else {
			return tuple.New2(next, keyNext)
		}
	}).V1, nil
}

// Sum returns the sum of all the values in the provided iterator.
func Sum[T ordered](oi *Iter[T]) T {
	var z T
	return oi.FoldEndo(z, func(curr, next T) T { return curr + next })
}

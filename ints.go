package iter

import "constraints"

// Ints returns an iterator that produces constraints.Integer values of the specified
// generic type, starting from 0 and increasing by 1.
func Ints[T constraints.Integer]() Iter[T] {
	var curr T
	return func() (T, bool) {
		res := curr
		curr++
		return res, true
	}
}

// IntsFrom returns an iterator that produces constraints.Integer values of the specified
// generic type, starting from the provided value and increasing by 1.
func IntsFrom[T constraints.Integer](start T) Iter[T] {
	curr := start
	return func() (T, bool) {
		res := curr
		curr++
		return res, true
	}
}

// IntsBy returns an iterator that produces constraints.Integer values of the specified
// generic type, starting from 0 and increasing by the provided value.
func IntsBy[T constraints.Integer](by T) Iter[T] {
	var curr T
	return func() (T, bool) {
		res := curr
		curr += by
		return res, true
	}
}

// IntsFromBy returns an iterator that produces constraints.Integer values of the specified
// generic type, starting from, and increasing by the provided values.
func IntsFromBy[T constraints.Integer](start T, by T) Iter[T] {
	curr := start
	return func() (T, bool) {
		res := curr
		curr += by
		return res, true
	}
}

package iter

// Ints returns an iterator that produces integer values of the specified
// generic type, starting from 0 and increasing by 1.
func Ints[T integer]() Iter[T] {
	var curr T
	return Iter[T](func() (T, bool) {
		res := curr
		curr++
		return res, true
	})
}

// IntsFrom returns an iterator that produces integer values of the specified
// generic type, starting from the provided value and increasing by 1.
func IntsFrom[T integer](start T) Iter[T] {
	curr := start
	return Iter[T](func() (T, bool) {
		res := curr
		curr++
		return res, true
	})
}

// IntsBy returns an iterator that produces integer values of the specified
// generic type, starting from 0 and increasing by the provided value.
func IntsBy[T integer](by T) Iter[T] {
	var curr T
	return Iter[T](func() (T, bool) {
		res := curr
		curr += by
		return res, true
	})
}

// IntsFromBy returns an iterator that produces integer values of the specified
// generic type, starting from, and increasing by the provided values.
func IntsFromBy[T integer](start T, by T) Iter[T] {
	curr := start
	return Iter[T](func() (T, bool) {
		res := curr
		curr += by
		return res, true
	})
}

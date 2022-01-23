package iter

// FlatMapEndo returns a new iterator that yields the values produced by
// iterators returned by the provided function when it is applied to values
// from the input iterator.
func (i Iter[T]) FlatMapEndo(f func(T) Iter[T]) Iter[T] {
	return FlatMap(i, f)
}

// FlatMap returns a new iterator that yields the values produced by iterators
// returned by the provided function when it is applied to values from the
// input iterator.
func FlatMap[T, U any](i Iter[T], f func(T) Iter[U]) Iter[U] {
	curr := func() (U, bool) {
		var z U
		return z, false
	}

	var self Iter[U]
	self = func() (U, bool) {
		next, ok := curr()

		if ok {
			return next, true
		} else {
			nextCurr, ok := i()
			if ok {
				curr = f(nextCurr)
				return self()
			} else {
				var z U
				return z, false
			}
		}
	}
	return self
}

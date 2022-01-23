package iter

// Inspect produces an iterator with identical values as the input iterator,
// but it applies the provided function to values of the iterator as they are
// requested. This methodh differs from ForEach in that it is lazy, whereas
// ForEach is not.
func (i Iter[T]) Inspect(f func(T)) Iter[T] {
	return Iter[T](func() (T, bool) {
		next, ok := i()
		if ok {
			f(next)
			return next, true
		} else {
			var z T
			return z, false
		}
	})
}

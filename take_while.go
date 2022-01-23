package iter

// TakeWhile produces a new iterator that yields values from the input iterator
// while those values satisfy the provided function. Once the first failure
// occurs, no more values are yielded. If this occurs, the previously yielded
// values, as well as the first failing value, are consumed from the input
// iterator.
func (i Iter[T]) TakeWhile(f func(T) bool) Iter[T] {
	failed := false

	return Iter[T](func() (T, bool) {
		if !failed {
			next, ok := i()
			if ok {
				if f(next) {
					return next, true
				}
			}

			failed = true
		}

		var z T
		return z, false
	})
}

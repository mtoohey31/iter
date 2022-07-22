package iter

// TODO: refactor n to type uint in the next breaking change to enforce its
// positivity at a type level.

// Take returns an iterator that limits that yields up to (but no more than) n
// values from the input iterator.
func (i Iter[T]) Take(n int) Iter[T] {
	curr := 0

	return func() (T, bool) {
		if curr < n {
			curr++
			return i()
		} else {
			var z T
			return z, false
		}
	}
}

// TakeWhile produces a new iterator that yields values from the input iterator
// while those values satisfy the provided function. Once the first failure
// occurs, no more values are yielded. If this occurs, the previously yielded
// values, as well as the first failing value, are consumed from the input
// iterator.
func (i Iter[T]) TakeWhile(f func(T) bool) Iter[T] {
	failed := false

	return func() (T, bool) {
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
	}
}

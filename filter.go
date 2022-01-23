package iter

// Filter returns a new iterator that only yields the values in the input
// iterator that satisfy the provided function.
func (i Iter[T]) Filter(f func(T) bool) Iter[T] {
	return func() (T, bool) {
		for {
			next, ok := i()

			if !ok {
				var z T
				return z, false
			}

			if f(next) {
				return next, true
			}
		}
	}
}

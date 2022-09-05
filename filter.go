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

// FilterMap returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func (i Iter[T]) FilterMap(f func(T) (T, error)) Iter[T] {
	return FilterMap(i, f)
}

// FilterMap returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func FilterMap[T, U any](i Iter[T], f func(T) (U, error)) Iter[U] {
	return func() (U, bool) {
		for {
			next, ok := i()

			if !ok {
				var z U
				return z, false
			}

			if mappedNext, err := f(next); err == nil {
				return mappedNext, true
			}
		}
	}
}

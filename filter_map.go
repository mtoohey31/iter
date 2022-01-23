package iter

// FilterMapEndo returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func (i Iter[T]) FilterMapEndo(f func(T) (T, error)) Iter[T] {
	return FilterMap(i, f)
}

// FilterMap returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func FilterMap[T, U any](i Iter[T], f func(T) (U, error)) Iter[U] {
	return Iter[U](func() (U, bool) {
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
	})
}

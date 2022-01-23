package iter

// MapWhileEndo returns a new iterator that yields the values produced by
// applying the provided function to the values of the input iterator, until
// the first error occurs. At that point, no further values are returned.
func (i Iter[T]) MapWhileEndo(f func(T) (T, error)) Iter[T] {
	return MapWhile(i, f)
}

// MapWhile returns a new iterator that yields the values produced by applying
// the provided function to the values of the input iterator, until the first
// error occurs. At that point, no further values are returned.
func MapWhile[T, U any](i Iter[T], f func(T) (U, error)) Iter[U] {
	failed := false
	return func() (U, bool) {
		if failed {
			var z U
			return z, false
		} else {
			next, ok := i()
			if ok {
				if mappedNext, err := f(next); err == nil {
					return mappedNext, true
				} else {
					failed = true
					var z U
					return z, false
				}
			} else {
				var z U
				return z, false
			}
		}
	}
}

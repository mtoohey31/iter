package iter

// Elems returns an iterator over the values of the provided slice.
func Elems[T any](s []T) Iter[T] {
	index := -1
	return func() (T, bool) {
		index++
		if index < len(s) {
			return s[index], true
		}

		var z T
		return z, false
	}
}

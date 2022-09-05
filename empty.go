package iter

// Empty returns an iterator that yields no values.
func Empty[T any]() Iter[T] {
	return func() (T, bool) {
		var z T
		return z, false
	}
}

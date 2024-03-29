package iter

// GenWhile returns an iterator that yields return values from the provided
// function while it does not produce errors. After the first error, no more
// values are yielded.
func GenWhile[T any](f func() (T, error)) Iter[T] {
	failed := false
	var self Iter[T]
	self = func() (T, bool) {
		if failed {
			var z T
			return z, false
		}

		next, err := f()
		if err == nil {
			return next, true
		}

		failed = true
		return self()
	}
	return self
}

package iter

// TODO: rework this in the next breaking change so that we can enforce the
// requirements at a type level instead of having to panic.

// Windows returns an iterator over slices of the provided length, containing,
// first items 0 through n-1 of the input iterator, then items 1 through n,
// etc. Panics for n < 1.
func Windows[T any](i Iter[T], n int) Iter[[]T] {
	if n < 1 {
		panic("n < 1")
	}

	window := make([]T, n)
	index := n - 1
	if i.CollectInto(window[0:index]) < n-1 {
		var z []T
		return func() ([]T, bool) {
			return z, false
		}
	}

	return func() ([]T, bool) {
		next, ok := i()
		if !ok {
			var z []T
			return z, false
		}
		window[index] = next
		index = (index + 1) % n
		res := make([]T, n)
		copy(res[0:n-index], window[index:])
		copy(res[n-index:], window[:index])
		return res, true
	}

}

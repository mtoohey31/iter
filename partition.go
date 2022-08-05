package iter

// Partition returns two iterators, one containing the values of the original
// iterator that satisfy the provided function, the other containing the values
// that do not.
func (i Iter[T]) Partition(f func(T) bool) (Iter[T], Iter[T]) {
	var aCache []T
	var bCache []T

	return func() (T, bool) {
			if len(aCache) > 0 {
				res := aCache[0]
				aCache = aCache[1:]
				return res, true
			}

			for {
				next, ok := i()

				if !ok {
					var z T
					return z, false
				}

				if f(next) {
					return next, true
				}

				bCache = append(bCache, next)
			}
		}, func() (T, bool) {
			if len(bCache) > 0 {
				res := bCache[0]
				bCache = bCache[1:]
				return res, true
			}

			for {
				next, ok := i()

				if !ok {
					var z T
					return z, false
				}

				if !f(next) {
					return next, true
				}

				aCache = append(aCache, next)
			}
		}
}

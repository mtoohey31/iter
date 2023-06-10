package iter

import "github.com/barweiss/go-tuple"

// GroupBy returns a new iterator which yields tuples whose first field is a key
// returned by f, and whose second field is a sub-iterator yielding a group of
// consecutive values from the input iterator for which f returned the key in
// the first field.
func GroupBy[K comparable, V any](i Iter[V], f func(value V) (key K)) Iter[tuple.T2[K, Iter[V]]] {
	next, ok := i()
	if !ok {
		return Empty[tuple.T2[K, Iter[V]]]()
	}

	currentKey := f(next)

	// The outer pointer of recentCached is nil when there was no previous
	// sub-iterator. The inner slice is nil when the previous sub-iterator is
	// still being lazily evaluated and non-nil (but possibly of length 0)
	// otherwise.
	var recentCached *[]V
	return func() (tuple.T2[K, Iter[V]], bool) {
		if recentCached != nil && *recentCached == nil {
			// In this case, we've returned a sub-iterator in the past, and the
			// most recent one of those is still being lazily evaluated. In
			// order to determine whether we can return another sub-iterator, we
			// have to evaluate the input iterator until we get a new key.

			// Initialize this to mark the sub-iterator as not requiring any
			// further evaluation of the input iterator in case we exit the loop
			// on the first iteration.
			*recentCached = []V{}
			for {
				var ok bool
				next, ok = i()
				if !ok {
					// The input iterator is no longer returning values so there
					// cannot be another sub-iterator, so we return that the
					// outer iterator is exhausted.
					return tuple.T2[K, Iter[V]]{}, false
				}

				nextKey := f(next)
				if nextKey != currentKey {
					currentKey = nextKey
					break
				}

				*recentCached = append(*recentCached, next)
			}
		}

		// If we make it here then there's a next sub-iterator which corresponds
		// to currentKey, and next is the value that should be returned first
		// from this sub-iterator.

		// currentCached is nil when the sub-iterator that we're about to return
		// is still being evaluated lazily, and non-nil (but possibly of length
		// 0) otherwise.
		var currentCached []V
		recentCached = &currentCached

		// first saves the current value of next, which is always a value that
		// hasn't been returned by any iterator yet. We can't just use next
		// directly because this iterator might return its first value after
		// next has already been re-assigned.
		first := next
		firstReturned := false
		return tuple.T2[K, Iter[V]]{
			V1: currentKey,
			V2: func() (V, bool) {
				// Make sure we've returned the first value.
				if !firstReturned {
					firstReturned = true
					return first, true
				}

				if currentCached != nil {
					// If we're no longer being lazily evaluated...

					if len(currentCached) == 0 {
						// ...then if there are no more cached items, return
						// that the iterator is exhausted.
						var z V
						return z, false
					}

					// ...then if there are more cached items, return the next
					// cached value and remove it from the cached list.
					res := currentCached[0]
					currentCached = currentCached[1:]
					return res, true
				}

				// Otherwise, evaluate the input iterator.
				var ok bool
				next, ok = i()
				if !ok {
					// Don't set currentCached in this case, because when
					// evaluating the outer iterator, we should take the slow
					// path then realize that the input iterator is exhausted.
					var z V
					return z, false
				}

				nextKey := f(next)
				if nextKey != currentKey {
					// Update currentKey, and do set currentCached in this case,
					// because taking the slow path in the other iterator would
					// be incorrect since we'd miss the current value of next.
					currentKey = nextKey
					currentCached = []V{}
					var z V
					return z, false
				}

				return next, true
			},
		}, true
	}
}

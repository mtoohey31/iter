package iter

import "github.com/barweiss/go-tuple"

// Zip returns an iterator that yields tuples of the two provided input
// iterators.
func Zip[T, U any](a Iter[T], b Iter[U]) Iter[tuple.T2[T, U]] {
	return Iter[tuple.T2[T, U]](func() (tuple.T2[T, U], bool) {
		nextA, okA := a()
		nextB, okB := b()
		if okA && okB {
			return tuple.New2(nextA, nextB), true
		} else {
			var z tuple.T2[T, U]
			return z, false
		}
	})
}

// Enumerate returns an iterator of tuples indices and values from the input
// iterator.
func Enumerate[T any](i Iter[T]) Iter[tuple.T2[int, T]] {
	return Zip(Ints[int](), i)
}

// Unzip returns two iterators, one yielding the left values of the tuples
// yielded by the input iterator, the other yielding the right values of the
// tuples. Note that, while the input iterator is evaluated lazily,
// exceptionally inequal consumption of the left vs the right iterator can lead
// to high memory consumption by values cached for the other iterator.
func Unzip[T, U any](i Iter[tuple.T2[T, U]]) (Iter[T], Iter[U]) {
	var aCache []T
	var bCache []U

	// PERF: does using an index instead of reassigning the slice improve things?

	return func() (T, bool) {
			if len(aCache) == 0 {
				next, ok := i()
				if !ok {
					var z T
					return z, false
				}

				bCache = append(bCache, next.V2)
				return next.V1, true
			} else {
				res := aCache[0]
				aCache = aCache[1:]
				return res, true
			}
		}, func() (U, bool) {
			if len(bCache) == 0 {
				next, ok := i()
				if !ok {
					var z U
					return z, false
				}

				aCache = append(aCache, next.V1)
				return next.V2, true
			} else {
				res := bCache[0]
				bCache = bCache[1:]
				return res, true
			}
		}
}

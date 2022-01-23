package iter

import "github.com/barweiss/go-tuple"

// KVZip returns an iterator that yields tuples of the input map's keys and
// values. While the value lookup occurs lazily, the keys must be accumulated
// immediately when the iterator is created, so this operation can be expensive
// if performance is important.
func KVZip[T comparable, U any](m map[T]U) Iter[tuple.T2[T, U]] {
	var keys []T

	for key := range m {
		// TODO: try refactoring to use a goroutine and channel here so we
		// don't have to take up so much memory
		keys = append(keys, key)
	}

	return Iter[tuple.T2[T, U]](func() (tuple.T2[T, U], bool) {
		if len(keys) > 0 {
			// PERF: is going from back to front here faster?
			next := keys[0]
			keys = keys[1:]
			return tuple.New2(next, m[next]), true
		} else {
			var z tuple.T2[T, U]
			return z, false
		}
	})
}

// MapEndo returns a new iterator that yields the results of applying the
// provided function to the input iterator.
func (i Iter[T]) MapEndo(f func(T) T) Iter[T] {
	return Map(i, f)
}

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func Map[T, U any](i Iter[T], f func(T) U) Iter[U] {
	return func() (U, bool) {
		next, ok := i()
		if ok {
			return f(next), true
		} else {
			var z U
			return z, false
		}
	}
}

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

// FlatMapEndo returns a new iterator that yields the values produced by
// iterators returned by the provided function when it is applied to values
// from the input iterator.
func (i Iter[T]) FlatMapEndo(f func(T) Iter[T]) Iter[T] {
	return FlatMap(i, f)
}

// FlatMap returns a new iterator that yields the values produced by iterators
// returned by the provided function when it is applied to values from the
// input iterator.
func FlatMap[T, U any](i Iter[T], f func(T) Iter[U]) Iter[U] {
	curr := func() (U, bool) {
		var z U
		return z, false
	}

	var self Iter[U]
	self = func() (U, bool) {
		next, ok := curr()

		if ok {
			return next, true
		} else {
			nextCurr, ok := i()
			if ok {
				curr = f(nextCurr)
				return self()
			} else {
				var z U
				return z, false
			}
		}
	}
	return self
}

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
	return Iter[U](func() (U, bool) {
		next, ok := i()
		if ok {
			return f(next), true
		} else {
			var z U
			return z, false
		}
	})
}

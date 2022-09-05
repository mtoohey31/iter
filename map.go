package iter

import "github.com/barweiss/go-tuple"

// KVZipStrict returns an iterator that yields tuples of the input map's
// keys and values. While the value lookup occurs lazily, the keys must be
// accumulated immediately when the iterator is created, so this operation can
// be expensive for large maps. If this is a problem, see KVZipLazy.
func KVZipStrict[T comparable, U any](m map[T]U) Iter[tuple.T2[T, U]] {
	var keys []T

	for key := range m {
		keys = append(keys, key)
	}

	return func() (tuple.T2[T, U], bool) {
		if len(keys) > 0 {
			next := keys[0]
			keys = keys[1:]
			return tuple.New2(next, m[next]), true
		}

		var z tuple.T2[T, U]
		return z, false
	}
}

// KVZipLazy returns an iterator that yields tuples of the input map's
// keys and values. This function uses a channel and a goroutine to manage key
// values lazily. Warning: if you don't consume the whole iterator, you will end
// up leaking a goroutine:
// https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html.
// Unless you're dealing with large maps where strict evaluation of keys may be
// problematic, KVZipStrict is recommended.
func KVZipLazy[T comparable, U any](m map[T]U) Iter[tuple.T2[T, U]] {
	keyChan := make(chan tuple.T2[T, U])

	go func() {
		for key, value := range m {
			keyChan <- tuple.New2(key, value)
		}
		close(keyChan)
	}()

	return func() (tuple.T2[T, U], bool) {
		next, ok := <-keyChan
		if ok {
			return next, true
		}

		var z tuple.T2[T, U]
		return z, false
	}
}

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func (i Iter[T]) Map(f func(T) T) Iter[T] {
	return Map(i, f)
}

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func Map[T, U any](i Iter[T], f func(T) U) Iter[U] {
	return func() (U, bool) {
		next, ok := i()
		if ok {
			return f(next), true
		}

		var z U
		return z, false
	}
}

// MapWhile returns a new iterator that yields the values produced by applying
// the provided function to the values of the input iterator, until the first
// error occurs. At that point, no further values are returned.
func (i Iter[T]) MapWhile(f func(T) (T, error)) Iter[T] {
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
		}

		next, ok := i()
		if ok {
			if mappedNext, err := f(next); err == nil {
				return mappedNext, true
			}

			failed = true
			var z U
			return z, false
		}

		var z U
		return z, false
	}
}

// FlatMap returns a new iterator that yields the values produced by iterators
// returned by the provided function when it is applied to values from the
// input iterator.
func (i Iter[T]) FlatMap(f func(T) Iter[T]) Iter[T] {
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
		}

		nextCurr, ok := i()
		if ok {
			curr = f(nextCurr)
			return self()
		}

		var z U
		return z, false
	}
	return self
}

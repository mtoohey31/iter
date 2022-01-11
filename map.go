package iter

import "github.com/barweiss/go-tuple"

type mapDataInner[T comparable, U any] struct {
	innerKeys *Iter[T]
	mapping   map[T]U
}

// KVZip returns an iterator that yields tuples of the input map's keys and
// values. While the value lookup occurs lazily, the keys must be accumulated
// immediately when the iterator is created, so this operation can be expensive
// if performance is important.
func KVZip[T comparable, U any](m map[T]U) *Iter[tuple.T2[T, U]] {
	var keys []T

	for key := range m {
		keys = append(keys, key)
	}

	return WithInner[tuple.T2[T, U]](&mapDataInner[T, U]{innerKeys: Elems(keys), mapping: m})
}

func (i *mapDataInner[T, U]) HasNext() bool {
	return i.innerKeys.HasNext()
}

func (i *mapDataInner[T, U]) Next() (tuple.T2[T, U], error) {
	key, err := i.innerKeys.Next()

	if err == nil {
		return tuple.New2(key, i.mapping[key]), nil
	} else {
		return tuple.T2[T, U]{}, IteratorExhaustedError
	}
}

type mapFuncInner[T, U any] struct {
	inner   *Iter[T]
	mapFunc func(T) U
}

// MapEndo returns a new iterator that yields the results of applying the
// provided function to the input iterator.
func (i *Iter[T]) MapEndo(f func(T) T) *Iter[T] {
	return WithInner[T](&mapFuncInner[T, T]{inner: i, mapFunc: f})
}

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func Map[T, U any](i *Iter[T], f func(T) U) *Iter[U] {
	return WithInner[U](&mapFuncInner[T, U]{inner: i, mapFunc: f})
}

func (i *mapFuncInner[T, U]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *mapFuncInner[T, U]) Next() (U, error) {
	next, err := i.inner.Next()

	if err == nil {
		return i.mapFunc(next), nil
	} else {
		var z U
		return z, err
	}
}

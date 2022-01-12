package iter

import "github.com/barweiss/go-tuple"

type zipInner[T, U any] struct {
	innerA *Iter[T]
	innerB *Iter[U]
}

// TODO: figure out why this is an issue, or file a bug report
// func (i *Iter[T]) ZipEndo(o *Iter[T]) *Iter[tuple.T2[T, T]] {
// 	return WithInner[tuple.T2[T, T]](&zipInner[T, T]{innerA: i, innerB: o})
// }

// Zip returns an iterator that yields tuples of the two provided input
// iterators.
func Zip[T, U any](a *Iter[T], b *Iter[U]) *Iter[tuple.T2[T, U]] {
	return Wrap[tuple.T2[T, U]](&zipInner[T, U]{innerA: a, innerB: b})
}

func (i *zipInner[T, U]) HasNext() bool {
	return i.innerA.HasNext() && i.innerB.HasNext()
}

func (i *zipInner[T, U]) Next() (tuple.T2[T, U], error) {
	nextA, errA := i.innerA.Next()
	nextB, errB := i.innerB.Next()

	if errA == nil && errB == nil {
		return tuple.New2(nextA, nextB), nil
	} else {
		var z tuple.T2[T, U]
		return z, IteratorExhaustedError
	}
}

// func (i *Iter[T]) Enumerate() *Iter[tuple.T2[int, T]] {
// 	return Zip(InfRange(0, 1), i)
// }

// Enumerate returns an iterator of tuples indices and values from the input
// iterator.
func Enumerate[T any](i *Iter[T]) *Iter[tuple.T2[int, T]] {
	return Zip(Ints[int](), i)
}

type unzipInner1[T, U any] struct {
	inner  *Iter[tuple.T2[T, U]]
	other  *unzipInner2[T, U]
	cached []T
	index  int
}

type unzipInner2[T, U any] struct {
	inner  *Iter[tuple.T2[T, U]]
	other  *unzipInner1[T, U]
	cached []U
	index  int
}

// Unzip returns two iterators, one yielding the left values of the tuples
// yielded by the input iterator, the other yielding the right values of the
// tuples. Note that, while the input iterator is evaluated lazily,
// exceptionally inequal consumption of the left vs the right iterator can lead
// to high memory consumption by values cached for the other iterator.
func Unzip[T, U any](i *Iter[tuple.T2[T, U]]) (*Iter[T], *Iter[U]) {
	inner1 := unzipInner1[T, U]{inner: i}
	inner2 := unzipInner2[T, U]{inner: i}
	inner1.other, inner2.other = &inner2, &inner1
	return Wrap[T](&inner1), Wrap[U](&inner2)
}

func (i *unzipInner1[T, U]) HasNext() bool {
	return i.index < len(i.cached) || i.inner.HasNext()
}

func (i *unzipInner2[T, U]) HasNext() bool {
	return i.index < len(i.cached) || i.inner.HasNext()
}

func (i *unzipInner1[T, U]) Next() (T, error) {
	if i.index < len(i.cached) {
		res := i.cached[i.index]
		i.index = i.index + 1
		return res, nil
	}

	tup, err := i.inner.Next()

	if err != nil {
		var z T
		return z, IteratorExhaustedError
	}

	i.other.cached = append(i.other.cached, tup.V2)
	return tup.V1, nil
}

func (i *unzipInner2[T, U]) Next() (U, error) {
	if i.index < len(i.cached) {
		res := i.cached[i.index]
		i.index = i.index + 1
		return res, nil
	}

	tup, err := i.inner.Next()

	if err != nil {
		var z U
		return z, IteratorExhaustedError
	}

	i.other.cached = append(i.other.cached, tup.V1)
	return tup.V2, nil
}

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

func Zip[T, U any](a *Iter[T], b *Iter[U]) *Iter[tuple.T2[T, U]] {
	return WithInner[tuple.T2[T, U]](&zipInner[T, U]{innerA: a, innerB: b})
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
		return Iter[tuple.T2[T, U]]{}.zeroVal(), IteratorExhaustedError
	}
}

// func (i *Iter[T]) Enumerate() *Iter[tuple.T2[int, T]] {
// 	return Zip(InfRange(0, 1), i)
// }

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

func Unzip[T, U any](i *Iter[tuple.T2[T, U]]) (*Iter[T], *Iter[U]) {
	inner1 := unzipInner1[T, U]{inner: i}
	inner2 := unzipInner2[T, U]{inner: i}
	inner1.other, inner2.other = &inner2, &inner1
	return WithInner[T](&inner1), WithInner[U](&inner2)
}

func (i *unzipInner1[T, U]) HasNext() bool {
	return i.index < len(i.cached) || i.inner.HasNext()
}

func (i *unzipInner2[T, U]) HasNext() bool {
	return i.index < len(i.cached) || i.inner.HasNext()
}

func (i *unzipInner1[T, U]) Next() (T, error) {
	if i.index < len(i.cached) {
		defer func() { i.index = i.index + 1 }()
		return i.cached[i.index], nil
	}

	tup, err := i.inner.Next()

	if err != nil {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}

	i.other.cached = append(i.other.cached, tup.V2)
	return tup.V1, nil
}

func (i *unzipInner2[T, U]) Next() (U, error) {
	if i.index < len(i.cached) {
		defer func() { i.index = i.index + 1 }()
		return i.cached[i.index], nil
	}

	tup, err := i.inner.Next()

	if err != nil {
		return Iter[U]{}.zeroVal(), IteratorExhaustedError
	}

	i.other.cached = append(i.other.cached, tup.V1)
	return tup.V2, nil
}

package iter

import "github.com/barweiss/go-tuple"

type zipInner[T, U any] struct {
	innerA *Iter[T]
	innerB *Iter[U]
}

// TODO: figure out why this is an issue, or file a bug report
// func (i *Iter[T]) ZipSame(o *Iter[T]) *Iter[tuple.T2[T, T]] {
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

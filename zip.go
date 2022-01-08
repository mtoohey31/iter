package iter

type zipInner[T any, U any] struct {
	innerA *Iter[T]
	innerB *Iter[U]
}

type Pair[T any, U any] struct {
	A T
	B U
}

// TODO: figure out why this is an issue, or file a bug report
// func (i *Iter[T]) ZipSame(o *Iter[T]) *Iter[Pair[T, T]] {
// 	return WithInner[Pair[T, T]](&zipInner[T, T]{innerA: i, innerB: o})
// }

func Zip[T any, U any](a *Iter[T], b *Iter[U]) *Iter[Pair[T, U]] {
	return WithInner[Pair[T, U]](&zipInner[T, U]{innerA: a, innerB: b})
}

func (i *zipInner[T, U]) HasNext() bool {
	return i.innerA.HasNext() && i.innerB.HasNext()
}

func (i *zipInner[T, U]) Next() (Pair[T, U], error) {
	nextA, errA := i.innerA.Next()
	nextB, errB := i.innerB.Next()

	if errA == nil && errB == nil {
		return Pair[T, U]{nextA, nextB}, nil
	} else {
		return Pair[T, U]{}, IteratorExhaustedError
	}
}

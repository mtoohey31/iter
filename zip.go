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

type unzipInner[T, U any] struct {
	inner       *Iter[tuple.T2[T, U]]
	leftCached  []*T
	leftIndex   int
	rightCached []*U
	rightIndex  int
}

func (ui *unzipInner[T, U]) HasNextLeft() bool {
	return ui.leftIndex < len(ui.leftCached) || ui.inner.HasNext()
}

func (ui *unzipInner[T, U]) HasNextRight() bool {
	return ui.rightIndex < len(ui.rightCached) || ui.inner.HasNext()
}

func (ui *unzipInner[T, U]) NextLeft() (T, error) {
	if ui.leftIndex < len(ui.leftCached) {
		res := *ui.leftCached[ui.leftIndex]
		ui.leftIndex++
		return res, nil
	}

	next, err := ui.inner.Next()

	if err == nil {
		ui.rightCached = append(ui.rightCached, &next.V2)
		return next.V1, nil
	} else {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}
}

func (ui *unzipInner[T, U]) NextRight() (U, error) {
	if ui.rightIndex < len(ui.rightCached) {
		res := *ui.rightCached[ui.rightIndex]
		ui.rightIndex++
		return res, nil
	}

	next, err := ui.inner.Next()

	if err == nil {
		ui.leftCached = append(ui.leftCached, &next.V1)
		return next.V2, nil
	} else {
		return Iter[U]{}.zeroVal(), IteratorExhaustedError
	}
}

type unzipInnerLeft[T, U any] struct {
	inner *unzipInner[T, U]
}

type unzipInnerRight[T, U any] struct {
	inner *unzipInner[T, U]
}

func (i *unzipInnerLeft[T, U]) HasNext() bool {
	return i.inner.HasNextLeft()
}

func (i *unzipInnerRight[T, U]) HasNext() bool {
	return i.inner.HasNextRight()
}

func (i *unzipInnerLeft[T, U]) Next() (T, error) {
	return i.inner.NextLeft()
}

func (i *unzipInnerRight[T, U]) Next() (U, error) {
	return i.inner.NextRight()
}

func Unzip[T, U any](i *Iter[tuple.T2[T, U]]) (*Iter[T], *Iter[U]) {
	inner := unzipInner[T, U]{inner: i}
	innerLeft := unzipInnerLeft[T, U]{inner: &inner}
	innerRight := unzipInnerRight[T, U]{inner: &inner}
	return WithInner[T](&innerLeft), WithInner[U](&innerRight)
}

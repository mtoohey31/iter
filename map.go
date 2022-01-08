package iter

type mapInner[T any, U any] struct {
	inner   *Iter[T]
	mapFunc func(T) U
}

func (i *Iter[T]) MapSame(f func(T) T) *Iter[T] {
	return WithInner[T](&mapInner[T, T]{inner: i, mapFunc: f})
}

func Map[T any, U any](i *Iter[T], f func(T) U) *Iter[U] {
	return WithInner[U](&mapInner[T, U]{inner: i, mapFunc: f})
}

func (i *mapInner[T, U]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *mapInner[T, U]) Next() (U, error) {
	next, err := i.inner.Next()

	if err == nil {
		return i.mapFunc(next), nil
	} else {
		return Iter[U]{}.zeroVal(), err
	}
}

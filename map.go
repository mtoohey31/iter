package iter

type mapIter[T any, U any] struct {
	inner   Iter[T]
	mapFunc func(T) U
	zero    U // TODO: find a less gross way to produce a zero value of a generic type
}

func Map[T any, U any](i Iter[T], f func(T) U) Iter[U] {
	return mapIter[T, U]{inner: i, mapFunc: f}
}

func (i mapIter[T, U]) HasNext() bool {
	return i.inner.HasNext()
}

func (i mapIter[T, U]) Next() (U, error) {
	next, err := i.inner.Next()

	if err == nil {
		return i.mapFunc(next), nil
	} else {
		return i.zero, err
	}
}

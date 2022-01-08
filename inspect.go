package iter

type inspectInner[T any] struct {
	inner       *Iter[T]
	inspectFunc func(T)
}

func (i *Iter[T]) Inspect(f func(T)) *Iter[T] {
	return WithInner[T](&inspectInner[T]{inner: i, inspectFunc: f})
}

func (i *inspectInner[T]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *inspectInner[T]) Next() (T, error) {
	next, err := i.inner.Next()

	if err == nil {
		i.inspectFunc(next)
		return next, nil
	} else {
		return Iter[T]{}.zeroVal(), err
	}
}

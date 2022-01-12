package iter

type inspectInner[T any] struct {
	inner       *Iter[T]
	inspectFunc func(T)
}

// Inspect produces an iterator with identical values as the input iterator,
// but it applies the provided function to values of the iterator as they are
// requested. This methodh differs from ForEach in that it is lazy, whereas
// ForEach is not.
func (i *Iter[T]) Inspect(f func(T)) *Iter[T] {
	return Wrap[T](&inspectInner[T]{inner: i, inspectFunc: f})
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
		var z T
		return z, err
	}
}

package iter

type filterIter[T any] struct {
	inner      Iter[T]
	filterFunc func(T) bool
	cachedNext *T
	zero       T // TODO: find a less gross way to produce a zero value of a generic type
}

func Filter[T any](i Iter[T], f func(T) bool) Iter[T] {
	return filterIter[T]{inner: i, filterFunc: f}
}

func (i filterIter[T]) findNext() (T, error) {
	for {
		next, err := i.inner.Next()

		if err != nil {
			break
		}

		if i.filterFunc(next) {
			return next, nil
		}
	}
	return i.zero, IteratorExhaustedError
}

func (i filterIter[T]) HasNext() bool {
	if i.cachedNext != nil {
		return true
	}

	next, err := i.findNext()
	i.cachedNext = &next
	return err == nil
}

func (i filterIter[T]) Next() (T, error) {
	if i.cachedNext != nil {
		defer func() { i.cachedNext = nil }()
		return *i.cachedNext, nil
	} else {
		return i.findNext()
	}
}

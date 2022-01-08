package iter

type filterMapInner[T any, U any] struct {
	inner         *Iter[T]
	filterMapFunc func(T) (U, error)
	cachedNext    *U
}

func (i *Iter[T]) FilterMapSame(f func(T) (T, error)) *Iter[T] {
	return WithInner[T](&filterMapInner[T, T]{inner: i, filterMapFunc: f})
}

func FilterMap[T any, U any](i *Iter[T], f func(T) (U, error)) *Iter[U] {
	return WithInner[U](&filterMapInner[T, U]{inner: i, filterMapFunc: f})
}

func (i *filterMapInner[T, U]) findNext() (U, error) {
	for {
		next, err := i.inner.Next()

		if err != nil {
			break
		}

		if mappedNext, err := i.filterMapFunc(next); err == nil {
			return mappedNext, nil
		}
	}
	return Iter[U]{}.zeroVal(), IteratorExhaustedError
}

func (i *filterMapInner[T, U]) HasNext() bool {
	if i.cachedNext != nil {
		return true
	}

	next, err := i.findNext()
	i.cachedNext = &next
	return err == nil
}

func (i *filterMapInner[T, U]) Next() (U, error) {
	if i.cachedNext != nil {
		defer func() { i.cachedNext = nil }()
		return *i.cachedNext, nil
	} else {
		return i.findNext()
	}
}

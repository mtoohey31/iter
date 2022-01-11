package iter

type filterMapInner[T, U any] struct {
	inner         *Iter[T]
	filterMapFunc func(T) (U, error)
	cachedNext    *U
}

// FilterMapEndo returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func (i *Iter[T]) FilterMapEndo(f func(T) (T, error)) *Iter[T] {
	return WithInner[T](&filterMapInner[T, T]{inner: i, filterMapFunc: f})
}

// FilterMap returns a new iterator that yields the mapped values which are
// produced without errors from the provided function.
func FilterMap[T, U any](i *Iter[T], f func(T) (U, error)) *Iter[U] {
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

	var z U
	return z, IteratorExhaustedError
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
		res := *i.cachedNext
		i.cachedNext = nil
		return res, nil
	} else {
		return i.findNext()
	}
}

package iter

type filterInner[T any] struct {
	inner      *Iter[T]
	filterFunc func(T) bool
	cachedNext *T
}

// Filter returns a new iterator that only yields the values in the input
// iterator that satisfy the provided function.
func (i *Iter[T]) Filter(f func(T) bool) *Iter[T] {
	return Wrap[T](&filterInner[T]{inner: i, filterFunc: f})
}

func (i *filterInner[T]) findNext() (T, error) {
	for {
		next, err := i.inner.Next()

		if err != nil {
			break
		}

		if i.filterFunc(next) {
			return next, nil
		}
	}

	var z T
	return z, IteratorExhaustedError
}

func (i *filterInner[T]) HasNext() bool {
	if i.cachedNext != nil {
		return true
	}

	next, err := i.findNext()
	i.cachedNext = &next
	return err == nil
}

func (i *filterInner[T]) Next() (T, error) {
	if i.cachedNext != nil {
		res := *i.cachedNext
		i.cachedNext = nil
		return res, nil
	} else {
		return i.findNext()
	}
}

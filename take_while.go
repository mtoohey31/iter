package iter

type takeWhileInner[T any] struct {
	inner         *Iter[T]
	takeWhileFunc func(T) bool
	cachedNext    *T
	failed        bool
}

func (i *Iter[T]) TakeWhile(f func(T) bool) *Iter[T] {
	return WithInner[T](&takeWhileInner[T]{inner: i, takeWhileFunc: f})
}

func (i *takeWhileInner[T]) findNext() (T, error) {
	for {
		next, err := i.inner.Next()

		if err != nil {
			break
		}

		if ok := i.takeWhileFunc(next); ok {
			return next, nil
		} else {
			break
		}
	}
	return Iter[T]{}.zeroVal(), IteratorExhaustedError
}

func (i *takeWhileInner[T]) HasNext() bool {
	if i.failed {
		return false
	}

	if i.cachedNext != nil {
		return true
	}

	next, err := i.findNext()
	i.cachedNext = &next
	return err == nil
}

func (i *takeWhileInner[T]) Next() (T, error) {
	if i.failed {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}

	if i.cachedNext != nil {
		res := *i.cachedNext
		i.cachedNext = nil
		return res, nil
	} else {
		return i.findNext()
	}
}

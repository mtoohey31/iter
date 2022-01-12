package iter

type takeWhileInner[T any] struct {
	inner         *Iter[T]
	takeWhileFunc func(T) bool
	cachedNext    *T
	failed        bool
}

// TakeWhile produces a new iterator that yields values from the input iterator
// while those values satisfy the provided function. Once the first failure
// occurs, no more values are yielded. If this occurs, the previously yielded
// values, as well as the first failing value, are consumed from the input
// iterator.
func (i *Iter[T]) TakeWhile(f func(T) bool) *Iter[T] {
	return Wrap[T](&takeWhileInner[T]{inner: i, takeWhileFunc: f})
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
	var z T
	return z, IteratorExhaustedError
}

func (i *takeWhileInner[T]) HasNext() bool {
	if i.failed {
		return false
	}

	if i.cachedNext != nil {
		return true
	}

	next, err := i.findNext()

	if err == nil {
		i.cachedNext = &next
		return true
	} else {
		i.failed = true
		return false
	}
}

func (i *takeWhileInner[T]) Next() (T, error) {
	if i.failed {
		var z T
		return z, IteratorExhaustedError
	}

	if i.cachedNext != nil {
		res := *i.cachedNext
		i.cachedNext = nil
		return res, nil
	}

	next, err := i.findNext()

	if err != nil {
		i.failed = true
		var z T
		return z, IteratorExhaustedError
	}

	return next, nil
}

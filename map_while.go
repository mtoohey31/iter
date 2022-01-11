package iter

type mapWhileInner[T, U any] struct {
	inner        *Iter[T]
	mapWhileFunc func(T) (U, error)
	cachedNext   *U
	failed       bool
}

// MapWhileEndo returns a new iterator that yields the values produced by
// applying the provided function to the values of the input iterator, until
// the first error occurs. At that point, no further values are returned.
func (i *Iter[T]) MapWhileEndo(f func(T) (T, error)) *Iter[T] {
	return WithInner[T](&mapWhileInner[T, T]{inner: i, mapWhileFunc: f})
}

// MapWhile returns a new iterator that yields the values produced by applying
// the provided function to the values of the input iterator, until the first
// error occurs. At that point, no further values are returned.
func MapWhile[T, U any](i *Iter[T], f func(T) (U, error)) *Iter[U] {
	return WithInner[U](&mapWhileInner[T, U]{inner: i, mapWhileFunc: f})
}

func (i *mapWhileInner[T, U]) findNext() (U, error) {
	for {
		next, err := i.inner.Next()

		if err != nil {
			break
		}

		if mappedNext, err := i.mapWhileFunc(next); err == nil {
			return mappedNext, nil
		} else {
			break
		}
	}
	var z U
	return z, IteratorExhaustedError
}

func (i *mapWhileInner[T, U]) HasNext() bool {
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

func (i *mapWhileInner[T, U]) Next() (U, error) {
	if i.failed {
		var z U
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
		var z U
		return z, IteratorExhaustedError
	}

	return next, err
}

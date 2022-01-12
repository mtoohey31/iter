package iter

type generatorInner[T any] struct {
	generatorFunc func() T
}

// Gen returns an iterator that yields return values from the provided
// function.
func Gen[T any](f func() T) *Iter[T] {
	return Wrap[T](&generatorInner[T]{generatorFunc: f})
}

func (i *generatorInner[T]) HasNext() bool {
	return true
}

func (i *generatorInner[T]) Next() (T, error) {
	return i.generatorFunc(), nil
}

type generatorWhileInner[T any] struct {
	generatorFunc func() (T, error)
	cachedNext    *T
	failed        bool
}

// GenWhile returns an iterator that yields return values from the provided
// function while it does not produce errors. After the first error, no more
// values are yielded.
func GenWhile[T any](f func() (T, error)) *Iter[T] {
	return Wrap[T](&generatorWhileInner[T]{generatorFunc: f})
}

func (i *generatorWhileInner[T]) findNext() (T, error) {
	next, err := i.generatorFunc()

	if err == nil {
		return next, nil
	} else {
		var z T
		return z, err
	}
}

func (i *generatorWhileInner[T]) HasNext() bool {
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

func (i *generatorWhileInner[T]) Next() (T, error) {
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

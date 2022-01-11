package iter

type flatMapInner[T, U any] struct {
	inner   *Iter[T]
	mapFunc func(T) *Iter[U]
	curr    *Iter[U]
}

// FlatMapEndo returns a new iterator that yields the values produced by
// iterators returned by the provided function when it is applied to values
// from the input iterator.
func (i *Iter[T]) FlatMapEndo(f func(T) *Iter[T]) *Iter[T] {
	return WithInner[T](&flatMapInner[T, T]{
		inner:   i,
		mapFunc: f,
		curr:    WithInner[T](&emptyInner[T]{}),
	})
}

// FlatMap returns a new iterator that yields the values produced by iterators
// returned by the provided function when it is applied to values from the
// input iterator.
func FlatMap[T, U any](i *Iter[T], f func(T) *Iter[U]) *Iter[U] {
	return WithInner[U](&flatMapInner[T, U]{
		inner:   i,
		mapFunc: f,
		curr:    WithInner[U](&emptyInner[U]{}),
	})
}

func (i *flatMapInner[T, U]) findNext() (*Iter[U], error) {
	next, err := i.inner.Next()

	if err == nil {
		return i.mapFunc(next), nil
	} else {
		return &Iter[U]{}, err
	}
}

func (i *flatMapInner[T, U]) HasNext() bool {
	if i.curr.HasNext() {
		return true
	} else {
		next, err := i.findNext()

		if err != nil {
			return false
		}

		i.curr = next
		return i.HasNext()
	}
}

func (i *flatMapInner[T, U]) Next() (U, error) {
	next, err := i.curr.Next()

	if err == nil {
		return next, nil
	} else {
		next, err := i.findNext()

		if err != nil {
			var z U
			return z, err
		}

		i.curr = next
		return i.Next()
	}
}

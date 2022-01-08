package iter

type flatMapInner[T any, U any] struct {
	inner   *Iter[T]
	mapFunc func(T) *Iter[U]
	curr    *Iter[U]
}

func (i *Iter[T]) FlatMapSame(f func(T) *Iter[T]) *Iter[T] {
	return WithInner[T](&flatMapInner[T, T]{
		inner:   i,
		mapFunc: f,
		curr:    WithInner[T](&emptyInner[T]{}),
	})
}

func FlatMap[T any, U any](i *Iter[T], f func(T) *Iter[U]) *Iter[U] {
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
			return Iter[U]{}.zeroVal(), err
		}

		i.curr = next
		return i.Next()
	}
}

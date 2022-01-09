package iter

type chanInner[T any] struct {
	ch         *chan T
	cachedNext *T
}

func Msgs[T any](ch *chan T) *Iter[T] {
	return WithInner[T](&chanInner[T]{ch: ch})
}

func (i *chanInner[T]) getNext() (T, error) {
	next, ok := <-*i.ch

	if ok {
		return next, nil
	} else {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}
}

func (i *chanInner[T]) HasNext() bool {
	if i.cachedNext != nil {
		return true
	} else {
		next, err := i.getNext()

		if err != nil {
			return false
		}

		i.cachedNext = &next
		return true
	}
}

func (i *chanInner[T]) Next() (T, error) {
	if i.cachedNext != nil {
		return *i.cachedNext, nil
	} else {
		next, err := i.getNext()

		if err != nil {
			return Iter[T]{}.zeroVal(), IteratorExhaustedError
		}

		return next, nil
	}
}

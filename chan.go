package iter

type chanInner[T any] struct {
	ch         *chan T
	cachedNext *T
}

// Msgs returns an iterator that reads values from the provided channel, and is
// exhausted when the channel is closed. Note that since this iterator reads
// from a channel, every time the next value is requested the program may end
// up deadlocking if values have not been written: the same rules apply as
// those for reading from a channel in the usual manner.
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
		err := *i.cachedNext
		i.cachedNext = nil
		return err, nil
	} else {
		next, err := i.getNext()

		if err != nil {
			return Iter[T]{}.zeroVal(), IteratorExhaustedError
		}

		return next, nil
	}
}

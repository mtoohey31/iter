package iter

type rangeInner[T integer] struct {
	curr T
	end  T
	step T
}

func Range[T integer](start T, end T, step T) *Iter[T] {
	return WithInner[T](&rangeInner[T]{curr: start, end: end, step: step})
}

func (i *rangeInner[T]) HasNext() bool {
	return i.curr < i.end
}

func (i *rangeInner[T]) Next() (T, error) {
	if !i.HasNext() {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}

	defer func() { i.curr = i.curr + i.step }()
	return i.curr, nil
}

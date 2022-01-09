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
	if i.step > 0 {
		return i.curr < i.end
	} else if i.step < 0 {
		return i.curr > i.end
	} else {
		return true
	}
}

func (i *rangeInner[T]) Next() (T, error) {
	if !i.HasNext() {
		return Iter[T]{}.zeroVal(), IteratorExhaustedError
	}

	defer func() { i.curr = i.curr + i.step }()
	return i.curr, nil
}

type infRangeInner[T integer] struct {
	curr T
	step T
}

func InfRange[T integer](start T, step T) *Iter[T] {
	return WithInner[T](&infRangeInner[T]{curr: start, step: step})
}

func (i *infRangeInner[T]) HasNext() bool { return true }

func (i *infRangeInner[T]) Next() (T, error) {
	defer func() { i.curr = i.curr + i.step }()
	return i.curr, nil
}

package iter

type rangeInt interface {
	int | int8 | int16 | int32 | int64
}

type rangeInner[T rangeInt] struct {
	curr T
	end  T
	step T
	zero T // TODO: find a less gross way to produce a zero value of a generic type
}

func Range[T rangeInt](start T, end T, step T) *Iter[T] {
	return WithInner[T](&rangeInner[T]{curr: start, end: end, step: step})
}

func (i *rangeInner[T]) HasNext() bool {
	return i.curr < i.end
}

func (i *rangeInner[T]) Next() (T, error) {
	if !i.HasNext() {
		return i.zero, IteratorExhaustedError
	}

	defer func() { i.curr = i.curr + i.step }()
	return i.curr, nil
}

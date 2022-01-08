package iter

type takeInner[T any] struct {
	inner *Iter[T]
	curr  int
	max   int
}

func (i *Iter[T]) Take(n int) *Iter[T] {
	return WithInner[T](&takeInner[T]{inner: i, max: n})
}

func (i *takeInner[T]) HasNext() bool {
	return i.curr < i.max && i.inner.HasNext()
}

func (i *takeInner[T]) Next() (T, error) {
	if i.curr < i.max {
		i.curr++
		return i.inner.Next()
	} else {
		return i.inner.zeroVal(), IteratorExhaustedError
	}
}

package iter

type stepByInner[T any] struct {
	inner *Iter[T]
	step  int
}

func (i *Iter[T]) StepBy(step int) *Iter[T] {
	return WithInner[T](&stepByInner[T]{inner: i, step: step})
}

func (i *stepByInner[T]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *stepByInner[T]) Next() (T, error) {
	defer func() { i.inner.Take(i.step - 1).Collect() }()
	return i.inner.Next()
}

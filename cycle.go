package iter

type cycleInner[T any] struct {
	inner  *Iter[T]
	memory []T
	index  int
}

func (i *Iter[T]) Cycle() *Iter[T] {
	return WithInner[T](&cycleInner[T]{inner: i, index: -1})
}

func (i *cycleInner[T]) HasNext() bool {
	return true
}

func (i *cycleInner[T]) Next() (T, error) {
	if i.index == -1 {
		next, err := i.inner.Next()

		if err == nil {
			i.memory = append(i.memory, next)
			return next, nil
		} else {
			i.index = 0
			return i.Next()
		}
	} else {
		defer func() { i.index = (i.index + 1) % len(i.memory) }()
		return i.memory[i.index], nil
	}
}

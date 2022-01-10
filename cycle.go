package iter

type cycleInner[T any] struct {
	inner  *Iter[T]
	memory []T
	index  int
}

// Cycle returns an iterator that first consumes the provided input iterator,
// then repeatedly returns the previous values. This method will panic if the
// provided iterator is empty, to ensure this doesn't happen, check if your
// iterator `.HasNext()` before passing it to Cycle.
func (i *Iter[T]) Cycle() *Iter[T] {
	if !i.HasNext() {
		panic("Cycle iterator contained no values")
	}

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
		res := i.memory[i.index]
		i.index = (i.index + 1) % len(i.memory)
		return res, nil
	}
}

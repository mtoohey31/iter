package iter

import "fmt"

type stepByInner[T any] struct {
	inner *Iter[T]
	step  int
}

// StepBy returns a new iterator that returns every nth value of the input
// iterator. The provided step value must be positive, otherwise the method
// will panic.
func (i *Iter[T]) StepBy(step int) *Iter[T] {
	if step < 1 {
		panic(fmt.Sprintf("invalid StepBy step: %d", step))
	}

	return Wrap[T](&stepByInner[T]{inner: i, step: step})
}

func (i *stepByInner[T]) HasNext() bool {
	return i.inner.HasNext()
}

func (i *stepByInner[T]) Next() (T, error) {
	res, err := i.inner.Next()
	i.inner.Take(i.step - 1).Consume()
	return res, err
}

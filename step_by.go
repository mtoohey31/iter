package iter

import "fmt"

// StepBy returns a new iterator that returns every nth value of the input
// iterator. The provided step value must be positive, otherwise the method
// will panic.
func (i *Iter[T]) StepBy(step int) *Iter[T] {
	if step < 1 {
		panic(fmt.Sprintf("invalid StepBy step: %d", step))
	}

	tmp := Iter[T](func() (T, bool) {
		res, err := i.Next()
		i.Take(step - 1).Consume()
		return res, err
	})
	return &tmp
}

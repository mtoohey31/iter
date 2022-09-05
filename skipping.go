package iter

// Skipping returns a new iterator that returns every n+1'th value of the input
// iterator.
func (i Iter[T]) Skipping(n uint) Iter[T] {
	return func() (T, bool) {
		res, err := i()
		i.Take(n).Consume()
		return res, err
	}
}

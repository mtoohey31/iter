package iter

// Chain returns a new iterator that first consumes the first iterator, then
// the second.
func (i Iter[T]) Chain(o Iter[T]) Iter[T] {
	return Iter[T](func() (T, bool) {
		next, ok := i()
		if ok {
			return next, true
		} else {
			return o()
		}
	})
}

package iter

// FlatMapEndo returns a new iterator that yields the values produced by
// iterators returned by the provided function when it is applied to values
// from the input iterator.
func (i *Iter[T]) FlatMapEndo(f func(T) *Iter[T]) *Iter[T] {
	return FlatMap(i, f)
}

// FlatMap returns a new iterator that yields the values produced by iterators
// returned by the provided function when it is applied to values from the
// input iterator.
func FlatMap[T, U any](i *Iter[T], f func(T) *Iter[U]) *Iter[U] {
	tmp := Iter[U](func() (U, bool) {
		var z U
		return z, false
	})
	curr := &tmp

	var self Iter[U]
	self = Iter[U](func() (U, bool) {
		next, ok := curr.Next()

		if ok {
			return next, true
		} else {
			nextCurr, ok := i.Next()
			if ok {
				curr = f(nextCurr)
				return self.Next()
			} else {
				var z U
				return z, false
			}
		}
	})
	return &self
}

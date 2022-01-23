package iter

// Take returns an iterator that limits that yields up to (but no more than) n
// values from the input iterator.
func (i *Iter[T]) Take(n int) *Iter[T] {
	curr := 0

	tmp := Iter[T](func() (T, bool) {
		if curr < n {
			curr++
			return i.Next()
		} else {
			var z T
			return z, false
		}
	})
	return &tmp
}

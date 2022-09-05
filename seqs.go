package iter

// Seqs returns an iterator over slices of length n+1 containing first items 0
// through n of the input iterator, then items 1 through n+1, etc.
func Seqs[T any](i Iter[T], n uint) Iter[[]T] {
	m := int(n) + 1
	buf := make([]T, m)
	index := int(n)

	if i.CollectInto(buf[0:index]) < m-1 {
		return Empty[[]T]()
	}

	return func() ([]T, bool) {
		next, ok := i()
		if !ok {
			var z []T
			return z, false
		}
		buf[index] = next
		index = (index + 1) % m
		res := make([]T, m)
		copy(res[0:m-index], buf[index:])
		copy(res[m-index:], buf[:index])
		return res, true
	}
}

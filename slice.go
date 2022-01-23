package iter

// Elems returns an iterator over the values of the provided slice.
func Elems[T any](s []T) *Iter[T] {
	index := -1
	tmp := Iter[T](func() (T, bool) {
		index++
		if len(s) > index {
			return s[index], true
		} else {
			var z T
			return z, false
		}
	})
	return &tmp
}

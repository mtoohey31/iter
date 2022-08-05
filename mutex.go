package iter

import "sync"

// Mutex returns a new iterator that uses a mutex to prevent multiple
// simultaneous evaluations of the wrapped iterator. This method should be
// the final method called on an iterator before it is used across multiple
// goroutines, as other methods are not guaranteed to be behave correctly when
// multiple evaluations occur simultaneously.
func (i Iter[T]) Mutex() Iter[T] {
	var m sync.Mutex

	return func() (T, bool) {
		m.Lock()
		r, ok := i()
		m.Unlock()
		return r, ok
	}
}

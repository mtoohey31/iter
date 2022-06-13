package iter

import "sync"

// Mutex returns a new iterator that ensures that only one evaluation of the
// prior steps in the iterator takes place at a time. This is helpful when
// using iterators that are not safe to use across goroutines in combination
// with operators prefixed with Go (that use goroutines). For example, using
// GoCollect on an Ints iterator may result in some values being collected twice
// if the iterator is not chained through Mutex first.
//
// Since a large part of the potential performance advantages of Go reduction
// operators is that they allow for concurrent and sometimes parallel evaluation
// of expensive intermediate steps, Mutex should be called as early in the chain
// as possible. To be specific, Mutex should be used immediately after the last
// goroutine-unsafe operator, since using it any earlier will result in unsafe
// behaviour, and using it any later may result in potential performance losses.
//
// If you want to avoid having to Mutex after an operation that is unsafe but
// expensive, check whether there is an M-prefixed variant of that operator.
// Those operators are explained in more detail in the package doc.
func (i Iter[T]) Mutex() Iter[T] {
	var m sync.Mutex

	return func() (T, bool) {
		m.Lock()
		defer m.Unlock()
		return i()
	}
}

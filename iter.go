/*
Package iter provides generic, lazy iterators, functions for producing them
from primitive types, as well as functions and methods for transforming
and consuming them.

When reading the documentation of the functions contained in this package, you
should assume that any function which accepts an iterator, but does not return
one, consumes it unless otherwise stated, meaning that the values contained
within cannot be used again.

Reducers with names prefixed by Go indicate that they perform some operations
using goroutines. This means that multiple evaluations of the prior steps
in the iterator may take place at the same time. This could result in race
conditions depending on the operators that have been used so far in the chain.
The Mutex method will ensure that only one evaluation of the prior steps occurs
at once. See its documentation for more information on how to use it. These
operators expose an n parameter to set the number of goroutines that should be
spawned since the optimal value will be different in each case. Benchmarking
should be used to determine what's best for your situation.

On a related note, operators with names prefixed by M are themselves safe to use
across goroutines. However, if prior operations are unsafe with respect to use
across goroutines, using an M-prefixed operator will not protect against those
issues; in that case, Mutex must be used. If an operator has an M-variant, you
can assume that the usual variant is unsafe for use across goroutines. If an
operator does not have an M-variant, no assumptions should be made. To
complicate things further, some operators goroutine-safety may depend on the
parameters they're passed (particularly those that accept functions). If you are
unsure, examine the implemenation.

Methods with names suffixed by Endo indicate that the method transforms
iterators of generic type T to some type in terms of T, such as T or Iter[T].
Transformation between types is possible, but only through the corresponding
function whose name is identical to the method, without the Endo prefix.
Functions are required for these operations because Go does not support the
definition of type parameters on methods. The nomenclature comes from the term
endomorphism, though it is a bit of a misuse of the term in that some Endo
methods take extra parameters or return types derived from T other than
Iter[T].
*/
package iter

import (
	"sync"

	"github.com/barweiss/go-tuple"
)

// Iter is a generic iterator function, the basis of this whole package. Note
// that the typical iter.Next() method is replaced with iter(), since Iter
// is simply defined as func() (T, bool).
type Iter[T any] func() (T, bool)

// Consume fetches the next value of the iterator until no more values are
// found. Note that it doesn't do anything with the values that are produced,
// but it can be useful in certain cases, such dropping a fixed number of items
// from an iterator by chaining Take and Consume.
func (i Iter[T]) Consume() {
	for {
		_, ok := i()

		if !ok {
			break
		}
	}
}

// GoConsume fetches the next value of the iterator until no more values are
// found. Note that it doesn't do anything with the values that are produced,
// but it can be useful in certain cases, such dropping a fixed number of items
// from an iterator by chaining Take and Consume. n is the number of goroutines
// that should be spawned.
func (i Iter[T]) GoConsume(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for j := 0; j < n; j++ {
		go func() {
			for {
				_, ok := i()

				if !ok {
					break
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

// Collect fetches the next value of the iterator until no more values are
// found, places these values in a slice, and returns it. Note that since
// Collect does not know how many values it will find, it must resize the slice
// multiple times during the collection process. This can result in poor
// performance, so CollectInto should be used when possible.
func (i Iter[T]) Collect() []T {
	var res []T
	for {
		next, ok := i()

		if !ok {
			break
		}

		res = append(res, next)
	}
	return res
}

// CollectInto inserts values yielded by the input iterator into the provided
// slice, and returns the number of values it was able to add before the
// iterator was exhausted or the slice was full.
func (i Iter[T]) CollectInto(buf []T) int {
	for j := 0; j < len(buf); j++ {
		next, ok := i()

		if !ok {
			return j
		}

		buf[j] = next
	}

	return len(buf)
}

// All returns whether all values of the iterator satisfy the provided
// predicate function.
func (i Iter[T]) All(f func(T) bool) bool {
	for {
		next, ok := i()

		if !ok {
			break
		}

		if !f(next) {
			return false
		}
	}
	return true
}

// Any returns whether any of the values of the iterator satisfy the provided
// predicate function.
func (i Iter[T]) Any(f func(T) bool) bool {
	for {
		next, ok := i()

		if !ok {
			break
		}

		if f(next) {
			return true
		}
	}
	return false
}

// Count returns the number of remaining values in the iterator.
func (i Iter[T]) Count() int {
	j := 0
	for {
		_, ok := i()

		if !ok {
			break
		}

		j++
	}
	return j
}

// Find returns the first value in the iterator that satisfies the provided
// predicate function, as well as a boolean indicating whether any value was
// found. It consumes all values up to the first satisfactory value, or the
// whole iterator if no values satisfy the predicate.
func (i Iter[T]) Find(f func(T) bool) (T, bool) {
	for {
		next, ok := i()

		if !ok {
			break
		}

		if f(next) {
			return next, true
		}
	}

	var z T
	return z, false
}

// FindMapEndo returns the first transformed value in the iterator for which
// the provided function does not return an error. As with Find, it consumes
// all values up to the first passing one.
func (i Iter[T]) FindMapEndo(f func(T) (T, error)) (T, bool) {
	for {
		next, ok := i()

		if !ok {
			break
		}

		if mappedNext, err := f(next); err == nil {
			return mappedNext, true
		}
	}

	var z T
	return z, false
}

// FindMap returns the first transformed value in the iterator for which the
// provided function does not return an error. As with Find, it consumes all
// values up to the first passing one.
func FindMap[T, U any](i Iter[T], f func(T) (U, error)) (U, bool) {
	for {
		next, ok := i()

		if !ok {
			break
		}

		if mappedNext, err := f(next); err == nil {
			return mappedNext, true
		}
	}

	var z U
	return z, false
}

// FoldEndo repeatedly applies the provided function to the current value
// (starting with init) and the next value of the iterator, until the whole
// iterator is consumed.
func (i Iter[T]) FoldEndo(init T, f func(curr T, next T) T) T {
	curr := init

	for {
		next, ok := i()

		if !ok {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

// Fold repeatedly applies the provided function to the current value (starting
// with init) and the next value of the iterator, until the whole iterator is
// consumed.
func Fold[T, U any](i Iter[T], init U, f func(curr U, next T) U) U {
	curr := init

	for {
		next, ok := i()

		if !ok {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

// ForEach applies the provided function to all remaining values in the current
// iterator.
func (i Iter[T]) ForEach(f func(T)) {
	for {
		next, ok := i()

		if !ok {
			break
		}

		f(next)
	}
}

// ForEachParallel applies the provided function to all remaining values in the
// current iterator. It differs from ForEach in that, where ForEach runs on a
// single thread and waits for each execution of the function to complete
// before fetching the next value and calling the function again,
// ForEachParallel performs executions of the function on different threads and
// only waits for all executions at the end. When the function to be executed
// is expensive and the order in which values of the iterator are operated upon
// does not matter, this method can result in better performance than ForEach.
func (i Iter[T]) ForEachParallel(f func(T)) {
	var wg sync.WaitGroup
	for {
		next, ok := i()

		if !ok {
			break
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			f(next)
		}()
	}
	wg.Wait()
}

// Last returns the final value of the iterator, along with a boolean
// indicating whether the operation was successful, in other words, whether the
// iterator was already empty.
func (i Iter[T]) Last() (T, bool) {
	curr, ok := i()

	if !ok {
		var z T
		return z, false
	}

	for {
		next, ok := i()

		if ok {
			curr = next
		} else {
			return curr, true
		}
	}
}

// Nth returns the nth value in the iterator, and a boolean indicating whether
// the iterator was too short. The provided value of n should be non-negative.
func (i Iter[T]) Nth(n int) (T, bool) {
	for j := 0; j < n-1; j++ {
		_, ok := i()

		if !ok {
			var z T
			return z, false
		}
	}

	res, ok := i()

	if ok {
		return res, true
	} else {
		var z T
		return z, false
	}
}

// TryFoldEndo applies the provided fallible function to the current value
// (starting with init) and the next value of the iterator, until the whole
// iterator is consumed. If at any point an error is returned, the operation
// stops and that error is returned.
func (i Iter[T]) TryFoldEndo(init T, f func(curr T, next T) (T, error)) (T, error) {
	curr := init

	for {
		next, ok := i()

		if !ok {
			break
		}

		var err error
		curr, err = f(curr, next)

		if err != nil {
			var z T
			return z, err
		}
	}

	return curr, nil
}

// TryFold applies the provided fallible function to the current value
// (starting with init) and the next value of the iterator, until the whole
// iterator is consumed. If at any point an error is returned, the operation
// stops and that error is returned.
func TryFold[T, U any](i Iter[T], init U, f func(curr U, next T) (U, error)) (U, error) {
	curr := init

	for {
		next, ok := i()

		if !ok {
			break
		}

		var err error
		curr, err = f(curr, next)

		if err != nil {
			var z U
			return z, err
		}
	}

	return curr, nil
}

// TryForEach applies the provided fallible function to all remaining values in
// the current iterator, but stops if an error is returned at any point.
func (i Iter[T]) TryForEach(f func(T) error) error {
	for {
		next, ok := i()

		if !ok {
			break
		}

		err := f(next)

		if err != nil {
			return err
		}
	}

	return nil
}

// Reduce repeatedly applies the provided function to the current value
// (starting with iterator's first value) and the next value of the iterator,
// until the whole iterator is consumed. The boolean indicates whether the
// iterator was already empty, meaning that it could not be reduced.
func (i Iter[T]) Reduce(f func(curr T, next T) T) (T, bool) {
	curr, ok := i()

	if !ok {
		var z T
		return z, false
	}

	return i.FoldEndo(curr, f), true
}

// Position returns the index of the first element satisfying the provided
// function, or -1 if one is not found. It consumes every element up to and
// including the element that satisfies the function, or the whole iterator if
// no no satisfactory element is found.
func Position[T any](i Iter[T], f func(T) bool) int {
	tup, ok := Enumerate(i).Find(func(tup tuple.T2[int, T]) bool { return f(tup.V2) })
	if ok {
		return tup.V1
	} else {
		return -1
	}
}

// Rev produces a new iterator whose values are reversed from those of the
// input iterator. Note that this method is not lazy, it must consume the whole
// input iterator immediately to produce the reversed result.
func (i Iter[T]) Rev() Iter[T] {
	collected := i.Collect()
	for j, k := 0, len(collected)-1; j < k; j, k = j+1, k-1 {
		collected[j], collected[k] = collected[k], collected[j]
	}
	return Elems(collected)
}

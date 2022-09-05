/*
Package iter provides generic, lazy iterators, functions for producing them
from primitive types, as well as functions and methods for transforming
and consuming them.

When reading the documentation of the functions contained in this package, you
should assume that any function which accepts an iterator, but does not return
one, consumes it unless otherwise stated, meaning that the values contained
within cannot be used again.

It might seem as though this package's API contains some duplicate
functionality, since there are many functions that appear to do the same thing
as methods of the same names defined on Iter[T]. However, there is an important
difference between these functions and their method counterparts. Since Go
does not support type parameters on methods, (Iter[T]).Map can only return
Iter[T] (an iterator of the same type as the input iterator). Map does not have
this limitation; it can map from an Iter[T] to an Iter[U] (an iterator with a
different type than the input iterator). The method versions are still provided
even though their functionality is a strict subset of that of the function
versions because using methods when possible results in more readable code.
*/
package iter

import (
	"sync"

	"github.com/barweiss/go-tuple"
)

// Iter is a generic iterator function, the basis of this whole package. Note
// that the typical iter.Next() method is replaced with iter(), since Iter is
// simply defined as func() (T, bool).
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

// Collect fetches the next value of the iterator until no more values are
// found, places these values in a slice, and returns it. Note that since
// Collect does not know how many values it will find, it must resize the
// slice multiple times during the collection process. This can result in poor
// performance, so CollectInto should be used when possible.
func (i Iter[T]) Collect() []T {
	res := []T{}
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

// FindMap returns the first transformed value in the iterator for which the
// provided function does not return an error. As with Find, it consumes all
// values up to the first passing one.
func (i Iter[T]) FindMap(f func(T) (T, error)) (T, bool) {
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

// Fold repeatedly applies the provided function to the current value (starting
// with init) and the next value of the iterator, until the whole iterator
// is consumed.
func (i Iter[T]) Fold(init T, f func(curr T, next T) T) T {
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
// with init) and the next value of the iterator, until the whole iterator
// is consumed.
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

// ForEachParallel applies the provided function to all remaining values
// in the current iterator. It differs from ForEach in that, where ForEach
// runs on a single thread and waits for each execution of the function to
// complete before fetching the next value and calling the function again,
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
	for j := 0; j < n; j++ {
		_, ok := i()

		if !ok {
			var z T
			return z, false
		}
	}

	res, ok := i()

	if ok {
		return res, true
	}

	var z T
	return z, false
}

// TryFold applies the provided fallible function to the current value
// (starting with init) and the next value of the iterator, until the whole
// iterator is consumed. If at any point an error is returned, the operation
// stops and that error is returned.
func (i Iter[T]) TryFold(init T, f func(curr T, next T) (T, error)) (T, error) {
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

	return i.Fold(curr, f), true
}

// Position returns the index of the first element satisfying the provided
// function, or -1 if one is not found. It consumes every element up to and
// including the element that satisfies the function, or the whole iterator if
// no no satisfactory element is found.
func Position[T any](i Iter[T], f func(T) bool) int {
	tup, ok := Enumerate(i).Find(func(tup tuple.T2[int, T]) bool { return f(tup.V2) })
	if ok {
		return tup.V1
	}

	return -1
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

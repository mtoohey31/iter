/*
Package iter provides generic, lazy iterators, functions for producing them
from primitive types, as well as functions and methods for transforming
and consuming them.

When reading the documentation of the functions contained in this package, you
should assume that any function which accepts an iterator, but does not return
one, consumes it unless otherwise stated, meaning that the values contained
within cannot be used again.

Methods with names suffixed by Endo indicate that the method transforms
iterators of generic type T to some type in terms of T, such as T or *Iter[T].
Transformation between types is possible, but only through the corresponding
function whose name is identical to the method, without the Endo prefix.
Functions are required for these operations because Go does not support the
definition of type parameters on methods. The nomenclature comes from the term
endomorphism, though it is a bit of a misuse of the term in that some *Endo
methods take extra parameters or return types derived from T other than
*Iter[T].
*/
package iter

import (
	"errors"
	"sync"
)

// Indicates an error resulting from an iterator with no more values.
var IteratorExhaustedError = errors.New("iterator exhausted")

type InnerIter[T any] interface {
	HasNext() bool
	Next() (T, error)
}

// Iter is a generic iterator struct, the basis of this whole package.
type Iter[T any] struct {
	inner InnerIter[T]
}

// HasNext returns whether the struct contains a next element.
func (i *Iter[T]) HasNext() bool {
	return i.inner.HasNext()
}

// Next returns the iterator's next element, if it has one. Otherwise, it
// produces an error.
func (i *Iter[T]) Next() (T, error) {
	return i.inner.Next()
}

// WithInner produces a result of type *Iter[T], given a struct implementing
// the InnerIter interface.
func WithInner[T any](inner InnerIter[T]) *Iter[T] {
	return &Iter[T]{inner: inner}
}

type emptyInner[T any] struct{}

func (i *emptyInner[T]) HasNext() bool {
	return false
}

func (i *emptyInner[T]) Next() (T, error) {
	var z T
	return z, IteratorExhaustedError
}

// Consume fetches the next value of the iterator until no more values are
// found. Note that it doesn't do anything with the values that are produced,
// but it can be useful in certain cases, such as benchmarking.
func (i *Iter[T]) Consume() {
	for {
		_, err := i.Next()

		if err != nil {
			break
		}
	}
}

// Collect fetches the next value of the iterator until no more values are
// found, places these values in a slice, and returns it. Note that since
// Collect does not know how many values it will find, it must resize the slice
// multiple times during the collection process. This can result in poor
// performance, so CollectInto should be used when possible.
func (i *Iter[T]) Collect() []T {
	var res []T
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		res = append(res, next)
	}
	return res
}

// CollectInto inserts values yielded by the input iterator into the provided
// slice, and returns the number of values it was able to add before the
// iterator was exhausted or the slice was full.
func (i *Iter[T]) CollectInto(buf []T) int {
	for j := 0; j < len(buf); j++ {
		next, err := i.Next()

		if err != nil {
			return j
		}

		buf[j] = next
	}

	return len(buf)
}

// All returns whether all values of the iterator satisfy the provided
// predicate function.
func (i *Iter[T]) All(f func(T) bool) bool {
	for {
		next, err := i.Next()

		if err != nil {
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
func (i *Iter[T]) Any(f func(T) bool) bool {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if f(next) {
			return true
		}
	}
	return false
}

// Count returns the number of remaining values in the iterator.
func (i *Iter[T]) Count() int {
	j := 0
	for {
		_, err := i.Next()

		if err != nil {
			break
		}

		j++
	}
	return j
}

// Find returns the first value in the iterator that satisfies the provided
// predicate function, or returns an error if no satisfactory values were
// found. It consumes all values up to the first satisfactory value, or the
// whole iterator if no values satisfy the predicate.
func (i *Iter[T]) Find(f func(T) bool) (T, error) {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if f(next) {
			return next, nil
		}
	}

	var z T
	return z, errors.New("no element found")
}

// FindMapEndo returns the first transformed value in the iterator for which
// the provided function does not return an error. As with Find, it consumes
// all values up to the first passing one.
func (i *Iter[T]) FindMapEndo(f func(T) (T, error)) (T, error) {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if mappedNext, err := f(next); err == nil {
			return mappedNext, nil
		}
	}

	var z T
	return z, errors.New("no element found")
}

// FindMap returns the first transformed value in the iterator for which the
// provided function does not return an error. As with Find, it consumes all
// values up to the first passing one.
func FindMap[T, U any](i *Iter[T], f func(T) (U, error)) (U, error) {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if mappedNext, err := f(next); err == nil {
			return mappedNext, nil
		}
	}

	var z U
	return z, errors.New("no element found")
}

// FoldEndo repeatedly applies the provided function to the current value
// (starting with `init`) and the next value of the iterator, until the whole
// iterator is consumed.
func (i *Iter[T]) FoldEndo(init T, f func(curr T, next T) T) T {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

// Fold repeatedly applies the provided function to the current value (starting
// with `init`) and the next value of the iterator, until the whole iterator is
// consumed.
func Fold[T, U any](i *Iter[T], init U, f func(curr U, next T) U) U {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		curr = f(curr, next)
	}

	return curr
}

// ForEach applies the provided function to all remaining values in the current
// iterator.
func (i *Iter[T]) ForEach(f func(T)) {
	for {
		next, err := i.Next()

		if err != nil {
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
func (i *Iter[T]) ForEachParallel(f func(T)) {
	var wg sync.WaitGroup
	for {
		next, err := i.Next()

		if err != nil {
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

// Last returns the final value of the iterator, or an error if the iterator is
// already empty.
func (i *Iter[T]) Last() (T, error) {
	curr, err := i.Next()

	if err != nil {
		var z T
		return z, IteratorExhaustedError
	}

	for {
		next, err := i.Next()

		if err == nil {
			curr = next
		} else {
			return curr, nil
		}
	}
}

// Nth returns the nth value in the iterator, or an error if the iterator is
// too short. The provided value of `n` should be non-negative.
func (i *Iter[T]) Nth(n int) (T, error) {
	for j := 0; j < n-1; j++ {
		_, err := i.Next()

		if err != nil {
			var z T
			return z, IteratorExhaustedError
		}
	}

	res, err := i.Next()

	if err == nil {
		return res, nil
	} else {
		var z T
		return z, IteratorExhaustedError
	}
}

// TODO: refactor to return two iterators

// Partition returns two slices, one containing the values of the iterator that
// satisfy the provided function, the other containing the values that do not.
func (i *Iter[T]) Partition(f func(T) bool) ([]T, []T) {
	var a []T
	var b []T
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		if f(next) {
			a = append(a, next)
		} else {
			b = append(b, next)
		}
	}
	return a, b
}

// TryFoldEndo applies the provided fallible function to the current value
// (starting with `init`) and the next value of the iterator, until the whole
// iterator is consumed. If at any point an error is returned, the operation
// stops and that error is returned.
func (i *Iter[T]) TryFoldEndo(init T, f func(curr T, next T) (T, error)) (T, error) {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		curr, err = f(curr, next)

		if err != nil {
			var z T
			return z, err
		}
	}

	return curr, nil
}

// TryFold applies the provided fallible function to the current value
// (starting with `init`) and the next value of the iterator, until the whole
// iterator is consumed. If at any point an error is returned, the operation
// stops and that error is returned.
func TryFold[T, U any](i *Iter[T], init U, f func(curr U, next T) (U, error)) (U, error) {
	curr := init

	for {
		next, err := i.Next()

		if err != nil {
			break
		}

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
func (i *Iter[T]) TryForEach(f func(T) error) error {
	for {
		next, err := i.Next()

		if err != nil {
			break
		}

		err = f(next)

		if err != nil {
			return err
		}
	}

	return nil
}

// Reduce repeatedly applies the provided function to the current value
// (starting with iterator's first value) and the next value of the iterator,
// until the whole iterator is consumed.
func (i *Iter[T]) Reduce(f func(curr T, next T) T) (T, error) {
	curr, err := i.Next()

	if err != nil {
		var z T
		return z, IteratorExhaustedError
	}

	return i.FoldEndo(curr, f), nil
}

// func (i *Iter[T]) Position(f func(T) bool) int {
// 	tup, err := Enumerate(i).Find(func(tup tuple.T2[int, T]) bool { return f(tup.V2) })
// 	if err == nil {
// 		return tup.V1
// 	} else {
// 		return -1
// 	}
// }

// Rev produces a new iterator whose values are reversed from those of the
// input iterator. Note that this method is not lazy, it must consume the whole
// input iterator immediately to produce the reversed result.
func (i *Iter[T]) Rev() *Iter[T] {
	collected := i.Collect()
	for j, k := 0, len(collected)-1; j < k; j, k = j+1, k-1 {
		collected[j], collected[k] = collected[k], collected[j]
	}
	return Elems(collected)
}

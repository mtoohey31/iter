// lazy-eratosthenes demonstrates how to construct an infinite, lazy, prime
// sieve using iterators.
package main

import (
	"fmt"

	"github.com/barweiss/go-tuple"
	"mtoohey.com/iter/v2"
)

// this differs from iter.Elems in that it accepts a pointer to the slice,
// meaning that we can not only mutate the elements of s after the iterator has
// been created, but we can also manipulate its length by re-assigning it
func ptrElems[T any](s *[]T) iter.Iter[T] {
	index := -1
	return func() (T, bool) {
		index++
		if index < len(*s) {
			return (*s)[index], true
		}

		var z T
		return z, false
	}
}

// this differs from iter.TakeWhile in that it caches the previous value which
// failed to pass the predicate so that it can be tested again when this
// iterator is re-evaluated
func takeWhileResuming[T any](i iter.Iter[T], f func(T) bool) iter.Iter[T] {
	var prev *T

	return func() (T, bool) {
		if prev != nil {
			if f(*prev) {
				res := *prev
				prev = nil
				return res, true
			}
		} else {
			next, ok := i()
			if ok {
				if f(next) {
					return next, true
				}

				prev = new(T)
				*prev = next
			}
		}

		var z T
		return z, false
	}
}

func id[T any](v T) T { return v }

func lazyEratosthenes() iter.Iter[int] {
	// spots for 0 and 1 are allocated but not read to improve readability,
	// performance could be improved at the cost of clarity by making index 0
	// refer to 2
	isntPrime := make([]bool, 128 /* start with 128, will be resized later */)

	// an slice of iterators that mark indicies in isntPrime when evaluated
	markers := []iter.Iter[int]{}

	initial := iter.Enumerate(ptrElems(&isntPrime))
	// drop 0 and 1
	initial.Take(2).Consume()
	primes := initial.Inspect(func(t tuple.T2[int, bool]) {
		if t.V1+2 != len(isntPrime) {
			return
		}

		// double the array size, again we could save memory here by forgetting
		// the half that we've already used, but for simplicity that is omitted
		oldIsntPrime := isntPrime
		isntPrime = make([]bool, len(isntPrime)*2)
		copy(isntPrime, oldIsntPrime)

		// mark things in the new section
		iter.FlatMap(iter.Elems(markers), id[iter.Iter[int]]).Consume()
	}).Filter(func(t tuple.T2[int, bool]) bool {
		return !t.V2
	})
	return iter.Map(primes, func(t tuple.T2[int, bool]) int {
		return t.V1
	}).Inspect(func(i int) {
		marker := takeWhileResuming(iter.IntsFromBy(i*i, i), func(j int) bool {
			// because the predicate references len(isntPrime), and isntPrime
			// gets resized as we go, consuming this iterator will mark all
			// indices within the length of the current slice, then stop, but
			// when we resize isntPrime, we will be able to consume this again
			// to mark the values for the next section of the slice
			return j < len(isntPrime)
		}).Inspect(func(j int) {
			isntPrime[j] = true
		})
		marker.Consume()
		markers = append(markers, marker)
	})
}

func main() {
	lazyEratosthenes().ForEach(func(i int) { fmt.Println(i) })
}

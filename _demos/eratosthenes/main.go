// eratosthenes provides an example of how the laziness of iter.Elems allows us
// to mutate the slice's values after creating the iterator.
//
// It doesn't provide performance benefits over the strict version of this
// algorithm though, since we're always computing a fixed number of primes. For
// an example of an infinite, lazy, prime sieve, see lazy-eratosthenes.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/barweiss/go-tuple"
	"mtoohey.com/iter/v2"
)

func eratosthenes(n int) iter.Iter[int] {
	// spots for 0 and 1 are allocated but not read to improve readability,
	// performance could be improved at the cost of clarity by making index 0
	// refer to 2
	isntPrime := make([]bool, n)

	initial := iter.Enumerate(iter.Elems(isntPrime))
	// drop 0 and 1
	initial.Take(2).Consume()
	primes := initial.Filter(func(t tuple.T2[int, bool]) bool {
		return !t.V2
	})
	return iter.Map(primes, func(t tuple.T2[int, bool]) int {
		return t.V1
	}).Inspect(func(i int) {
		iter.IntsFromBy(i*i, i).TakeWhile(func(j int) bool {
			return j < n
		}).ForEach(func(j int) {
			isntPrime[j] = true
		})
	})
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s n\n", os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	eratosthenes(n).ForEach(func(i int) { fmt.Println(i) })
}

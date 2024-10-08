> As of Go 1.23, the [`for ... := range ...` syntax](https://go.dev/ref/spec#For_range) can now be used with "push-style" iterator functions. Since this package implements a different "pull-style" of iterators which are incompatible with the `for ... := range ...` syntax, I've decided to archive it. If you're looking for pull-style equivalents of some of the functions in this package, I'd recommend checking out the [`x/exp/iter` package proposal](https://github.com/golang/go/issues/61898).

# iter

Package iter provides generic, lazy iterators, functions for producing them from primitive types, as well as functions and methods for transforming and consuming them.

## Usage

<!-- `$ cat _demos/hello-world/main.go` as go -->
```go
package main

import (
	"fmt"
	"strings"

	"mtoohey.com/iter/v2"
)

func main() {
	initial := iter.Elems([]string{"hello", "beautiful", "world"})
	result := initial.Filter(func(s string) bool {
		return len(s) < 6
	}).Map(strings.ToUpper).Collect()
	fmt.Println(result) // produces: [HELLO WORLD]
}
```

## Regarding Performance

There is some overhead to using the iterators in this package, since each evaluation requires a function call, so if the performance of your application is a top priority, this package _might_ not be the best choice. Don't guess about performance though: I would recommend benchmarking to determine the impact of using iterators, because in some cases, lazy iterators may be faster than the equivalent strict loop.

## Acknowledgements

- Everyone that commented on my [r/golang post](https://www.reddit.com/r/golang/comments/s13jlz/iter_generic_lazy_iterators_for_go_118/)
- The inspiration for my own attempt: <https://github.com/polyfloyd/go-iterator>

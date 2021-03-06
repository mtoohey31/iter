# iter

Package iter provides generic, lazy iterators, functions for producing them from primitive types, as well as functions and methods for transforming and consuming them.

## Usage

<!-- `$ cat _demos/hello-world/main.go` as go -->
```go
package main

import (
	"fmt"
	"strings"

	"mtoohey.com/iter"
)

func main() {
	initial := iter.Elems([]string{"hello", "beautiful", "world"})
	result := initial.Filter(func(s string) bool {
		return len(s) < 6
	}).MapEndo(strings.ToUpper).Collect()
	fmt.Println(result) // produces: [HELLO WORLD]
}
```

## Practical Examples

- [mtoohey31/godoc-coverage](https://github.com/mtoohey31/godoc-coverage) makes use of this package and is heavily commented.

## Notes

- Some rudimentary benchmarks comparing performance to that of loop based solutions to similar problems can be found [here](https://github.com/mtoohey31/iter-loop-benchmarks).

## Acknowledgements

- Everyone that commented on my [r/golang post](https://www.reddit.com/r/golang/comments/s13jlz/iter_generic_lazy_iterators_for_go_118/)
- The inspiration for my own attempt: <https://github.com/polyfloyd/go-iterator>

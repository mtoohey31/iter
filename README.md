> The experimental [range over func](https://github.com/golang/go/issues/61405) feature, along with the proposed [`x/exp/xiter` package](https://github.com/golang/go/issues/61898) will hopefully soon provide a more convenient alternative to this package. The one downside is that they don't appear to be making any effort to support chaining with method calls, meaning iterators will have be read from right to left, instead of left to right. Once range over func is stable, I will likely archive this package.

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

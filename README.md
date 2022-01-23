# iter

Package iter provides generic, lazy iterators, functions for producing them from primitive types, as well as functions and methods for transforming and consuming them.

## Usage

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

See `examples/` for more demonstrations.

## Notes

- It should go without saying that since `1.18` is still beta, so is this library.

## Acknowledgements

- Everyone that commented on my [r/golang post](https://www.reddit.com/r/golang/comments/s13jlz/iter_generic_lazy_iterators_for_go_118/)
- The inspiration for my own attempt: <https://github.com/polyfloyd/go-iterator>

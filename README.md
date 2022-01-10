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

## Similar Projects

- The inspiration for my own attempt, which is more mature but doesn't support chaining calls: <https://github.com/polyfloyd/go-iterator>

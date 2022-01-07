# iter

Iterators for Go using `1.18beta1` generics.

## Usage

```go
package main

import (
  "fmt"
  "strings"

  "github.com/mtoohey31/iter"
)

func main() {
  initial := iter.FromSlice([]string{"hello", "beautiful", "world"})
  short := iter.Filter(initial, func(s string) bool {
    return len(s) < 6
  })
  upper := iter.Map(short, func(s string) string {
    return strings.ToUpper(s)
  })
  fmt.Println(iter.Collect(upper)) // produces: [HELLO WORLD]
}
```

## Notes

- It should go without saying that since `1.18` is still beta, so is this library.
- Operations commonly seen as methods are defined as standalone functions, since Go does not not support the definition of methods for interface types, and `Iter[T]` is a method.

## Similar Projects

- The inspiration for my own attempt, which is more mature but doesn't support chaining calls: <https://github.com/polyfloyd/go-iterator>

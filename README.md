# iter

Iterators for Go using `1.18beta1` generics.

## Usage

```go
package main

import (
  "fmt"
  "strings"

  "mtoohey.com/iter"
)

func main() {
  initial := iter.FromSlice([]string{"hello", "beautiful", "world"})
  result := initial.Filter(func(s string) bool {
    return len(s) < 6
  }).MapSame(strings.ToUpper).Collect()
  fmt.Println(result) // produces: [HELLO WORLD]
}
```

See `examples/` for more demonstrations.

## Notes

- It should go without saying that since `1.18` is still beta, so is this library.
- From what I can tell, with my limited understanding of go memory management, `sliceInner` struct is rather inefficient because it duplicates values and leaves past items sitting inside itself.

## Similar Projects

- The inspiration for my own attempt, which is more mature but doesn't support chaining calls: <https://github.com/polyfloyd/go-iterator>

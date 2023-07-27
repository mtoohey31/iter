package main

import (
	"fmt"
	"strings"

	"mtoohey.com/iter/v3"
)

func main() {
	initial := iter.Elems([]string{"hello", "beautiful", "world"})
	result := initial.Filter(func(s string) bool {
		return len(s) < 6
	}).Map(strings.ToUpper).Collect()
	fmt.Println(result) // produces: [HELLO WORLD]
}

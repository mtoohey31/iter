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

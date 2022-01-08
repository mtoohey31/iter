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
	}).MapSame(func(s string) string {
		return strings.ToUpper(s)
	}).Collect()
	fmt.Println(result) // produces: [HELLO WORLD]
}

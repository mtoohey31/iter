package main

import (
	"fmt"

	"github.com/mtoohey31/iter"
)

type alphebetInner struct {
	curr int
}

func (a *alphebetInner) HasNext() bool {
	return a.curr < 123
}

func (a *alphebetInner) Next() (rune, error) {
	if !a.HasNext() {
		return ' ', iter.IteratorExhaustedError
	} else {
		defer func() { a.curr = a.curr + 1 }()
		return rune(a.curr), nil
	}
}

func New() *iter.Iter[rune] {
	return iter.WithInner[rune](&alphebetInner{curr: 97})
}

func main() {
	fmt.Println(string(New().Collect())) // produces: abcdefghijklmnopqrstuvwxyz
}

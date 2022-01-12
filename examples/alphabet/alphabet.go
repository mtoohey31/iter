package main

import (
	"fmt"

	"mtoohey.com/iter"
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
		res := rune(a.curr)
		a.curr = a.curr + 1
		return res, nil
	}
}

func New() *iter.Iter[rune] {
	return iter.Wrap[rune](&alphebetInner{curr: 97})
}

func main() {
	fmt.Println(string(New().Collect())) // produces: abcdefghijklmnopqrstuvwxyz
}

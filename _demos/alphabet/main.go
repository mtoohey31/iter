package main

import "mtoohey.com/iter/v2"

func IterAlphabet() iter.Iter[rune] {
	c := 'a'
	return func() (rune, bool) {
		if c <= 'z' {
			res := c
			c++
			return res, true
		} else {
			return rune(0), false
		}
	}
}

func main() {
	println(string(IterAlphabet().Collect())) // produces: abcdefghijklmnopqrstuvwxyz
	cycle, _ := IterAlphabet().Take(3).Cycle()
	cycledRunes := make([]rune, 21)
	cycle.CollectInto(cycledRunes)
	println(string(cycledRunes)) // produces: abcabcabcabcabcabcabc
}

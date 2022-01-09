package iter

import (
	"errors"
	"strings"
	"testing"

	"mtoohey.com/iter/test"
)

func TestMapWhileEndo(t *testing.T) {
	initialIter := Elems([]string{"good", "bad", "good", "good"})
	mappedWhileIter := initialIter.MapWhileEndo(func(s string) (string, error) {
		if s == "bad" {
			return "", errors.New("")
		} else {
			return strings.ToUpper(s), nil
		}
	})

	test.AssertDeepEq(mappedWhileIter.Collect(), []string{"GOOD"}, t)
	test.AssertDeepEq(initialIter.Collect(), []string{"good", "good"}, t)
}

func TestMapwhile(t *testing.T) {
	initialIter := Elems([]string{"long string", "longer string", "short", "long string again"})
	mappedWhileIter := MapWhile(initialIter, func(s string) (int, error) {
		l := len(s)
		if l < 10 {
			return 0, errors.New("")
		} else {
			return l, nil
		}
	})

	test.AssertDeepEq(mappedWhileIter.Collect(), []int{11, 13}, t)
	test.AssertDeepEq(initialIter.Collect(), []string{"long string again"}, t)
}

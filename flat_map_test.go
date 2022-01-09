package iter

import (
	"strings"
	"testing"

	"mtoohey.com/iter/test"
)

func TestFlatMapEndo(t *testing.T) {
	initial := []int{1, 2, 3}
	iter := Elems(initial).FlatMapEndo(func(i int) *Iter[int] {
		return Range(i, i+2, 1)
	})

	actual := iter.Collect()
	expected := []int{1, 2, 2, 3, 3, 4}

	test.AssertDeepEq(actual, expected, t)
}

func TestFlatMap(t *testing.T) {
	initial := []string{"alpha", "beta", "gamma"}
	iter := FlatMap(Elems(initial), func(s string) *Iter[rune] {
		return Runes(s)
	})

	actual := string(iter.Collect())
	expected := strings.Join(initial, "")

	test.AssertDeepEq(actual, expected, t)
}

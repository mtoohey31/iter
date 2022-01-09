package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestFilter(t *testing.T) {
	iter := Elems([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i%2 == 0 })

	actual := iter.Collect()
	expected := []int{2, 4}

	test.AssertDeepEq(actual, expected, t)
}

func BenchmarkFilter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(b.N).Consume()
}

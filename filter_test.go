package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestFilter(t *testing.T) {
	iter := Elems([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i%2 == 0 })

	actualFirst, _ := iter.Next()
	expected := []int{2, 4}

	test.Assert(iter.HasNext(), t)
	test.Assert(iter.HasNext(), t)
	test.AssertDeepEq(append([]int{actualFirst}, iter.Collect()...), expected, t)
	test.Assert(!iter.HasNext(), t)
}

func BenchmarkFilter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(b.N).Consume()
}

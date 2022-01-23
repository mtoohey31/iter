package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestChain(t *testing.T) {
	iter := Elems([]int{1, 2}).Chain(Elems([]int{3, 4}))
	actual := iter.Collect()
	test.AssertDeepEq(actual, []int{1, 2, 3, 4}, t)
	// test.Assert(!iter.HasNext(), t)
}

// operations should not take much longer than that of the range iterator
func BenchmarkChain(b *testing.B) {
	Ints[int]().Take(b.N / 2).Chain(IntsFrom(b.N / 2).Take(b.N / 2)).Consume()
}

package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestCycle(t *testing.T) {
	iter := Elems([]int{1, 2}).Cycle()

	test.AssertDeepEq(iter.Take(6).Collect(), []int{1, 2, 1, 2, 1, 2}, t)
}

func BenchmarkCycle1(b *testing.B) {
	InfRange(0, 1).Take(1).Cycle().Take(b.N).Consume()
}

func BenchmarkCycle100(b *testing.B) {
	InfRange(0, 1).Take(100).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleQuarter(b *testing.B) {
	InfRange(0, 1).Take(1 + b.N/4).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleHalf(b *testing.B) {
	InfRange(0, 1).Take(1 + b.N/2).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleFull(b *testing.B) {
	InfRange(0, 1).Take(1 + b.N).Cycle().Take(b.N).Consume()
}

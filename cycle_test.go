package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestCycle(t *testing.T) {
	iter, ok := Elems([]int{1, 2}).Cycle()

	// test.Assert(iter.HasNext(), t)
	test.Assert(ok, t)

	test.AssertDeepEq(iter.Take(6).Collect(), []int{1, 2, 1, 2, 1, 2}, t)
}

func TestCyclePanic(t *testing.T) {
	_, ok := Elems([]bool{}).Cycle()

	test.Assert(!ok, t)
}

func BenchmarkCycle1(b *testing.B) {
	iter, _ := Ints[int]().Take(1).Cycle()
	iter.Take(b.N).Consume()
}

func BenchmarkCycle100(b *testing.B) {
	iter, _ := Ints[int]().Take(100).Cycle()
	iter.Take(b.N).Consume()
}

func BenchmarkCycleQuarter(b *testing.B) {
	iter, _ := Ints[int]().Take(1 + (b.N / 4)).Cycle()
	iter.Take(b.N).Consume()
}

func BenchmarkCycleHalf(b *testing.B) {
	iter, _ := Ints[int]().Take(1 + (b.N / 2)).Cycle()
	iter.Take(b.N).Consume()
}

func BenchmarkCycleFull(b *testing.B) {
	iter, _ := Ints[int]().Take(1 + b.N).Cycle()
	iter.Take(b.N).Consume()
}

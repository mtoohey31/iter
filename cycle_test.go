package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestCycle(t *testing.T) {
	iter := Elems([]int{1, 2}).Cycle()

	test.Assert(iter.HasNext(), t)

	test.AssertDeepEq(iter.Take(6).Collect(), []int{1, 2, 1, 2, 1, 2}, t)
}

func TestCyclePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Cycle should've panicked")
		}
	}()

	Elems([]bool{}).Cycle()
}

func BenchmarkCycle1(b *testing.B) {
	Ints[int]().Take(1).Cycle().Take(b.N).Consume()
}

func BenchmarkCycle100(b *testing.B) {
	Ints[int]().Take(100).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleQuarter(b *testing.B) {
	Ints[int]().Take(b.N / 4).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleHalf(b *testing.B) {
	Ints[int]().Take(b.N / 2).Cycle().Take(b.N).Consume()
}

func BenchmarkCycleFull(b *testing.B) {
	Ints[int]().Take(b.N).Cycle().Take(b.N).Consume()
}

package iter

import (
	"testing"

	"mtoohey.com/iter/test"
)

func TestRangeIter(t *testing.T) {
	test.AssertDeepEq(Range(1, 7, 2).Collect(), []int{1, 3, 5}, t)
}

func TestInfRangeIter(t *testing.T) {
	test.AssertDeepEq(InfRange(7, -2).Take(7).Collect(),
		[]int{7, 5, 3, 1, -1, -3, -5}, t)
}

func BenchmarkRangeIter1(b *testing.B) {
	Range(0, b.N, 1).Consume()
}

func BenchmarkRangeIter100(b *testing.B) {
	Range(0, b.N*100, 100).Consume()
}

func BenchmarkRangeIterMin1(b *testing.B) {
	Range(0, -b.N, -1).Consume()
}

func BenchmarkRangeIterMin100(b *testing.B) {
	Range(0, -b.N*100, -100).Consume()
}

func BenchmarkInfRangeIter1(b *testing.B) {
	InfRange(0, 1).Take(b.N).Consume()
}

func BenchmarkInfRangeIter100(b *testing.B) {
	InfRange(0, 100).Take(b.N).Consume()
}

func BenchmarkInfRangeIterMin1(b *testing.B) {
	InfRange(0, -1).Take(b.N).Consume()
}

func BenchmarkInfRangeIterMin100(b *testing.B) {
	InfRange(0, -100).Take(b.N).Consume()
}

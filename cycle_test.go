package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	iter, ok := Elems([]int{1, 2}).Cycle()

	assert.True(t, ok)

	assert.Equal(t, []int{1, 2, 1, 2, 1, 2}, iter.Take(6).Collect())
}

func TestCyclePanic(t *testing.T) {
	_, ok := Elems([]bool{}).Cycle()

	assert.True(t, !ok)
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

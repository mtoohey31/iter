package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzCycle(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		i, ok := Elems(b).Cycle()

		// when the input iterator is already empty, Cycle fails
		if len(b) == 0 {
			assert.False(t, ok)
		} else {
			assert.True(t, ok)

			expected := b
			expected = append(expected, b...)
			expected = append(expected, b...)
			assert.Equal(t, expected, i.Take(len(b)*3).Collect())
		}
	})
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

package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/internal/testutils"
)

func FuzzIter_Cycle(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		i, ok := Elems(b).Cycle()

		if len(b) == 0 {
			// When the input iterator is already empty, Cycle fails.
			assert.False(t, ok)
		} else {
			if assert.True(t, ok) {
				expected := b
				expected = append(expected, b...)
				expected = append(expected, b...)
				assert.Equal(t, expected, i.Take(uint(len(b))*3).Collect())
			}
		}
	})
}

func BenchmarkIter_Cycle_1(b *testing.B) {
	iter, ok := Ints[int]().Take(1).Cycle()
	assert.True(b, ok)
	iter.Take(uint(b.N)).Consume()
}

func BenchmarkIter_Cycle_100(b *testing.B) {
	iter, ok := Ints[int]().Take(100).Cycle()
	assert.True(b, ok)
	iter.Take(uint(b.N)).Consume()
}

func BenchmarkIter_Cycle_quarter(b *testing.B) {
	iter, ok := Ints[int]().Take(1 + uint(b.N)/4).Cycle()
	assert.True(b, ok)
	iter.Take(uint(b.N)).Consume()
}

func BenchmarkIter_Cycle_half(b *testing.B) {
	iter, ok := Ints[int]().Take(1 + uint(b.N)/2).Cycle()
	assert.True(b, ok)
	iter.Take(uint(b.N)).Consume()
}

func BenchmarkIter_Cycle_full(b *testing.B) {
	iter, ok := Ints[int]().Take(1 + uint(b.N)).Cycle()
	assert.True(b, ok)
	iter.Take(uint(b.N)).Consume()
}

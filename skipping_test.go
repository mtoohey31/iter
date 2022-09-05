package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzIter_Skipping(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		expected := []byte{}
		for i := uint(0); i < uint(len(b)); i += n + 1 {
			expected = append(expected, b[i])
		}

		assert.Equal(t, expected, Elems(b).Skipping(n).Collect())
	})
}

func BenchmarkIter_Skipping_1(b *testing.B) {
	Ints[int]().Skipping(1).Take(uint(b.N)).Consume()
}

func BenchmarkIter_Skipping_100(b *testing.B) {
	Ints[int]().Skipping(100).Take(uint(b.N)).Consume()
}

package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzIter_Take(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		var expected []byte
		if int(n) < len(b) {
			expected = b[:n]
		} else {
			expected = b
		}
		assert.Equal(t, expected, Elems(b).Take(n).Collect())
	})
}

func BenchmarkIter_Take(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Consume()
}

func FuzzIter_TakeWhile(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		expected := make([]uint, n)
		for i := uint(0); i < n; i++ {
			expected[i] = i
		}
		assert.Equal(t, expected, Ints[uint]().TakeWhile(func(u uint) bool {
			return u < n
		}).Collect())
	})
}

func BenchmarkIter_TakeWhile(b *testing.B) {
	Ints[int]().TakeWhile(func(i int) bool {
		return i < b.N
	}).Consume()
}

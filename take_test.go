package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzTake(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		var expected []byte
		if int(n) < len(b) {
			expected = b[:n]
		} else {
			expected = b
		}
		assert.Equal(t, expected, Elems(b).Take(int(n)).Collect())
	})
}

func BenchmarkTake(b *testing.B) {
	Ints[int]().Take(b.N).Consume()
}

func FuzzTakeWhile(f *testing.F) {
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

func BenchmarkTakeWhile(b *testing.B) {
	Ints[int]().TakeWhile(func(i int) bool {
		return i < b.N
	}).Consume()
}

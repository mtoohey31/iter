package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzMin(f *testing.F) {
	f.Add([]byte{2, 9, 1, 5, 0})
	f.Add([]byte{})
	f.Add([]byte{9, 4, 1})
	f.Add([]byte{30})

	f.Fuzz(func(t *testing.T, b []byte) {
		foundMin := false
		var min byte
		for _, v := range b {
			if !foundMin || v < min {
				foundMin = true
				min = v
			}
		}

		actual, ok := Min(Elems(b))
		if assert.Equal(t, foundMin, ok) && foundMin {
			assert.Equal(t, min, actual)
		}
	})
}

func BenchmarkMin(b *testing.B) {
	Min(Ints[int]().Take(b.N))
}

func FuzzMinByKey(f *testing.F) {
	f.Add([]byte{2, 9, 1, 5, 0})
	f.Add([]byte{})
	f.Add([]byte{9, 4, 1})
	f.Add([]byte{30})

	f.Fuzz(func(t *testing.T, b []byte) {
		foundMax := false
		var max byte
		for _, v := range b {
			if !foundMax || v > max {
				foundMax = true
				max = v
			}
		}

		actual, ok := MinByKey(Elems(b), func(v byte) int {
			return -int(v)
		})
		if assert.Equal(t, foundMax, ok) && foundMax {
			assert.Equal(t, max, actual)
		}
	})
}

func BenchmarkMinByKey(b *testing.B) {
	MinByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func FuzzMax(f *testing.F) {
	f.Add([]byte{2, 9, 1, 5, 0})
	f.Add([]byte{})
	f.Add([]byte{9, 4, 1})
	f.Add([]byte{30})

	f.Fuzz(func(t *testing.T, b []byte) {
		foundMax := false
		var max byte
		for _, v := range b {
			if !foundMax || v > max {
				foundMax = true
				max = v
			}
		}

		actual, ok := Max(Elems(b))
		if assert.Equal(t, foundMax, ok) && foundMax {
			assert.Equal(t, max, actual)
		}
	})
}

func BenchmarkMax(b *testing.B) {
	Max(Ints[int]().Take(b.N))
}

func FuzzMaxByKey(f *testing.F) {
	f.Add([]byte{2, 9, 1, 5, 0})
	f.Add([]byte{})
	f.Add([]byte{9, 4, 1})
	f.Add([]byte{30})

	f.Fuzz(func(t *testing.T, b []byte) {
		foundMin := false
		var min byte
		for _, v := range b {
			if !foundMin || v < min {
				foundMin = true
				min = v
			}
		}

		actual, ok := MaxByKey(Elems(b), func(v byte) int {
			return -int(v)
		})
		if assert.Equal(t, foundMin, ok) && foundMin {
			assert.Equal(t, min, actual)
		}
	})
}

func BenchmarkMaxByKey(b *testing.B) {
	MaxByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func FuzzSum(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		expected := uint(0)
		for i := uint(0); i < n; i++ {
			expected += i
		}

		assert.Equal(t, expected, Sum(Ints[uint]().Take(int(n))))
	})
}

func BenchmarkSum(b *testing.B) {
	Sum(Ints[int]().Take(b.N))
}

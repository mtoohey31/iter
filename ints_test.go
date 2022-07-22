package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Ints[int]().Take(10).Collect())
}

func FuzzIntsFrom(f *testing.F) {
	f.Add(0)
	f.Add(27)
	f.Add(168)
	f.Add(41981)

	f.Fuzz(func(t *testing.T, n int) {
		expected := []int{}
		for i := n; i < n+10; i++ {
			expected = append(expected, i)
		}

		assert.Equal(t, expected, IntsFrom(n).Take(10).Collect())
	})
}

func FuzzIntsBy(f *testing.F) {
	f.Add(0)
	f.Add(27)
	f.Add(168)
	f.Add(41981)

	f.Fuzz(func(t *testing.T, n int) {
		expected := []int{}
		i := 0
		for j := 0; j < 10; i, j = i+n, j+1 {
			expected = append(expected, i)
		}

		assert.Equal(t, expected, IntsBy(n).Take(10).Collect())
	})
}

func FuzzIntsFromBy(f *testing.F) {
	f.Add(0, 0)
	f.Add(27, 18)
	f.Add(168, 354)
	f.Add(41981, 94876)

	f.Fuzz(func(t *testing.T, m, n int) {
		expected := []int{}
		i := m
		for j := 0; j < 10; i, j = i+n, j+1 {
			expected = append(expected, i)
		}

		assert.Equal(t, expected, IntsFromBy(m, n).Take(10).Collect())
	})
}

func BenchmarkInts(b *testing.B) {
	Ints[int]().Take(b.N).Consume()
}

func BenchmarkIntsFromNegative(b *testing.B) {
	IntsFrom(0).Take(b.N).Consume()
}

func BenchmarkIntsFromZero(b *testing.B) {
	IntsFrom(0).Take(b.N).Consume()
}

func BenchmarkIntsFromPositive(b *testing.B) {
	IntsFrom(10000000).Take(b.N).Consume()
}

func BenchmarkIntsByZero(b *testing.B) {
	IntsBy(0).Take(b.N).Consume()
}

func BenchmarkIntsBySmallIncreasing(b *testing.B) {
	IntsBy(1).Take(b.N).Consume()
}

func BenchmarkIntsFromBigIncreasing(b *testing.B) {
	IntsBy(10000000).Take(b.N).Consume()
}

func BenchmarkIntsBySmallDecreasing(b *testing.B) {
	IntsBy(-1).Take(b.N).Consume()
}

func BenchmarkIntsFromBigDecreasing(b *testing.B) {
	IntsBy(-10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromNegativeByZero(b *testing.B) {
	IntsFromBy(-10000000, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromNegativeBySmallIncreasing(b *testing.B) {
	IntsFromBy(-10000000, 1).Take(b.N).Consume()
}
func BenchmarkIntsFromNegativeByLargeIncreasing(b *testing.B) {
	IntsFromBy(-10000000, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromNegativeBySmallDecreasing(b *testing.B) {
	IntsFromBy(-10000000, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromNegativeByLargeDecreasing(b *testing.B) {
	IntsFromBy(-10000000, -10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromZeroByZero(b *testing.B) {
	IntsFromBy(0, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromZeroBySmallIncreasing(b *testing.B) {
	IntsFromBy(0, 1).Take(b.N).Consume()
}

func BenchmarkIntsFromZeroByLargeIncreasing(b *testing.B) {
	IntsFromBy(0, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromZeroBySmallDecreasing(b *testing.B) {
	IntsFromBy(0, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromZeroByLargeDecreasing(b *testing.B) {
	IntsFromBy(0, -10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromPositiveByZero(b *testing.B) {
	IntsFromBy(10000000, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromPositiveBySmallIncreasing(b *testing.B) {
	IntsFromBy(10000000, 1).Take(b.N).Consume()
}
func BenchmarkIntsFromPositiveByLargeIncreasing(b *testing.B) {
	IntsFromBy(10000000, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromPositiveBySmallDecreasing(b *testing.B) {
	IntsFromBy(10000000, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromPositiveByLargeDecreasing(b *testing.B) {
	IntsFromBy(10000000, -10000000).Take(b.N).Consume()
}

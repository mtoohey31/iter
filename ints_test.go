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

func BenchmarkIntsFrom_negative(b *testing.B) {
	IntsFrom(0).Take(b.N).Consume()
}

func BenchmarkIntsFrom_zero(b *testing.B) {
	IntsFrom(0).Take(b.N).Consume()
}

func BenchmarkIntsFrom_positive(b *testing.B) {
	IntsFrom(10000000).Take(b.N).Consume()
}

func BenchmarkIntsBy_zero(b *testing.B) {
	IntsBy(0).Take(b.N).Consume()
}

func BenchmarkIntsBy_smallIncreasing(b *testing.B) {
	IntsBy(1).Take(b.N).Consume()
}

func BenchmarkIntsBy_bigIncreasing(b *testing.B) {
	IntsBy(10000000).Take(b.N).Consume()
}

func BenchmarkIntsBy_smallDecreasing(b *testing.B) {
	IntsBy(-1).Take(b.N).Consume()
}

func BenchmarkIntsBy_bigDecreasing(b *testing.B) {
	IntsBy(-10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_negativeByZero(b *testing.B) {
	IntsFromBy(-10000000, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_negativeBySmallIncreasing(b *testing.B) {
	IntsFromBy(-10000000, 1).Take(b.N).Consume()
}
func BenchmarkIntsFromBy_negativeByLargeIncreasing(b *testing.B) {
	IntsFromBy(-10000000, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_negativeBySmallDecreasing(b *testing.B) {
	IntsFromBy(-10000000, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_negativeByLargeDecreasing(b *testing.B) {
	IntsFromBy(-10000000, -10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_zeroByZero(b *testing.B) {
	IntsFromBy(0, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_zeroBySmallIncreasing(b *testing.B) {
	IntsFromBy(0, 1).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_zeroByLargeIncreasing(b *testing.B) {
	IntsFromBy(0, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_zeroBySmallDecreasing(b *testing.B) {
	IntsFromBy(0, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_zeroByLargeDecreasing(b *testing.B) {
	IntsFromBy(0, -10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_positiveByZero(b *testing.B) {
	IntsFromBy(10000000, 0).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_positiveBySmallIncreasing(b *testing.B) {
	IntsFromBy(10000000, 1).Take(b.N).Consume()
}
func BenchmarkIntsFromBy_positiveByLargeIncreasing(b *testing.B) {
	IntsFromBy(10000000, 10000000).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_positiveBySmallDecreasing(b *testing.B) {
	IntsFromBy(10000000, -1).Take(b.N).Consume()
}

func BenchmarkIntsFromBy_positiveByLargeDecreasing(b *testing.B) {
	IntsFromBy(10000000, -10000000).Take(b.N).Consume()
}

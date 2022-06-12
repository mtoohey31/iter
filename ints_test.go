package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, Ints[int]().Take(3).Collect())
}

func TestIntsFrom(t *testing.T) {
	assert.Equal(t, []int{2, 3, 4, 5}, IntsFrom(2).Take(4).Collect())
}

func TestIntsByZero(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0, 0, 0}, IntsBy(0).Take(5).Collect())
}

func TestIntsByIncreasing(t *testing.T) {
	assert.Equal(t, []int{0, 2, 4, 6}, IntsBy(2).Take(4).Collect())
}

func TestIntsByDecreasing(t *testing.T) {
	assert.Equal(t, []int{0, -4, -8}, IntsBy(-4).Take(3).Collect())
}

func TestIntsFromByZero(t *testing.T) {
	assert.Equal(t, []int{100, 100}, IntsFromBy(100, 0).Take(2).Collect())
}

func TestIntsFromByIncreasing(t *testing.T) {
	assert.Equal(t, []int{-3, 0, 3}, IntsFromBy(-3, 3).Take(3).Collect())
}

func TestIntsFromByDecreasing(t *testing.T) {
	assert.Equal(t, []int{7, 4, 1, -2}, IntsFromBy(7, -3).Take(4).Collect())
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

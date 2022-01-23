package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	assert.Equal(t, Ints[int]().Take(3).Collect(), []int{0, 1, 2})
}

func TestIntsFrom(t *testing.T) {
	assert.Equal(t, IntsFrom(2).Take(4).Collect(), []int{2, 3, 4, 5})
}

func TestIntsByZero(t *testing.T) {
	assert.Equal(t, IntsBy(0).Take(5).Collect(), []int{0, 0, 0, 0, 0})
}

func TestIntsByIncreasing(t *testing.T) {
	assert.Equal(t, IntsBy(2).Take(4).Collect(), []int{0, 2, 4, 6})
}

func TestIntsByDecreasing(t *testing.T) {
	assert.Equal(t, IntsBy(-4).Take(3).Collect(), []int{0, -4, -8})
}

func TestIntsFromByZero(t *testing.T) {
	assert.Equal(t, IntsFromBy(100, 0).Take(2).Collect(), []int{100, 100})
}

func TestIntsFromByIncreasing(t *testing.T) {
	assert.Equal(t, IntsFromBy(-3, 3).Take(3).Collect(), []int{-3, 0, 3})
}

func TestIntsFromByDecreasing(t *testing.T) {
	assert.Equal(t, IntsFromBy(7, -3).Take(4).Collect(), []int{7, 4, 1, -2})
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

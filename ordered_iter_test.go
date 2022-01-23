package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := Min(ordered)

	assert.True(t, ok)
	assert.Equal(t, actual, 0)

	_, ok = Min(ordered)

	assert.False(t, ok)

	Min(IntsBy(-1).Take(2))
}

func BenchmarkMin(b *testing.B) {
	Min(Ints[int]().Take(b.N))
}

func TestMinByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := MinByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.True(t, ok)
	assert.Equal(t, actual, 9)

	_, ok = MinByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.False(t, ok)

	ordered = IntsBy(-1).Take(10)

	actual, ok = MinByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.True(t, ok)
	assert.Equal(t, actual, 0)
}

func BenchmarkMinByKey(b *testing.B) {
	MinByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func TestMax(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := Max(ordered)

	assert.True(t, ok)
	assert.Equal(t, actual, 9)

	_, ok = Max(ordered)

	assert.False(t, ok)

	Max(IntsBy(-1).Take(2))
}

func BenchmarkMax(b *testing.B) {
	Max(Ints[int]().Take(b.N))
}

func TestMaxByKey(t *testing.T) {
	ordered := Ints[int]().Take(10)

	actual, ok := MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.True(t, ok)
	assert.Equal(t, actual, 0)

	_, ok = MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.False(t, ok)

	ordered = IntsBy(-1).Take(10)

	actual, ok = MaxByKey(ordered, func(n int) int {
		return n * -1
	})

	assert.True(t, ok)
	assert.Equal(t, actual, -9)
}

func BenchmarkMaxByKey(b *testing.B) {
	MaxByKey(Ints[int]().Take(b.N), func(n int) int {
		return n
	})
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(Ints[int]().Take(10)), 45)
}

func BenchmarkSum(b *testing.B) {
	Sum(Ints[int]().Take(b.N))
}

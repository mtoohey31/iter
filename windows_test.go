package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindows(t *testing.T) {
	assert.Equal(t, Windows(Ints[int]().Take(10), 5).Collect(),
		[][]int{
			{0, 1, 2, 3, 4},
			{1, 2, 3, 4, 5},
			{2, 3, 4, 5, 6},
			{3, 4, 5, 6, 7},
			{4, 5, 6, 7, 8},
			{5, 6, 7, 8, 9},
		})
	_, ok := Windows(Ints[int]().Take(10), 11)()
	assert.False(t, ok)
	_, ok = Windows(Ints[int]().Take(10), 12)()
	assert.False(t, ok)
}

func BenchmarkWindows1(b *testing.B) {
	Windows(Ints[int](), 1).Take(b.N).Consume()
}

func BenchmarkWindows3(b *testing.B) {
	Windows(Ints[int](), 1).Take(b.N).Consume()
}

func BenchmarkWindows10(b *testing.B) {
	Windows(Ints[int](), 10).Take(b.N).Consume()
}

func BenchmarkWindows100(b *testing.B) {
	Windows(Ints[int](), 100).Take(b.N).Consume()
}

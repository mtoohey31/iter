package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	actualA, actualB := Ints[int]().Take(4).Partition(func(i int) bool { return i%2 == 0 })

	assert.Equal(t, actualA.Collect(), []int{0, 2})
	assert.Equal(t, actualB.Collect(), []int{1, 3})

	actualA, actualB = Ints[int]().Take(4).Partition(func(i int) bool { return i%2 == 0 })

	assert.Equal(t, actualB.Collect(), []int{1, 3})
	assert.Equal(t, actualA.Collect(), []int{0, 2})
}

func BenchmarkPartition(b *testing.B) {
	iterA, iterB := Ints[int]().Partition(func(i int) bool {
		return i%2 == 0
	})
	iterA.Take(b.N).Consume()
	iterB.Take(b.N).Consume()
}

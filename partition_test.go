package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	actualA, actualB := Ints[int]().Take(4).Partition(func(i int) bool { return i%2 == 0 })

	assert.Equal(t, []int{0, 2}, actualA.Collect())
	assert.Equal(t, []int{1, 3}, actualB.Collect())

	actualA, actualB = Ints[int]().Take(4).Partition(func(i int) bool { return i%2 == 0 })

	assert.Equal(t, []int{1, 3}, actualB.Collect())
	assert.Equal(t, []int{0, 2}, actualA.Collect())
}

func BenchmarkPartition(b *testing.B) {
	iterA, iterB := Ints[int]().Partition(func(i int) bool {
		return i%2 == 0
	})
	iterA.Take(b.N).Consume()
	iterB.Take(b.N).Consume()
}

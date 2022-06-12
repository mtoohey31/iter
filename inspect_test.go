package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	iter := IntsFrom(1).Take(10)

	actualNum := 0
	expectedNumBefore := 0
	expectedNumAfter := 55

	newIter := iter.Inspect(func(n int) {
		actualNum = actualNum + n
	})

	assert.Equal(t, expectedNumBefore, actualNum)

	actualSlice := newIter.Collect()
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expectedNumAfter, actualNum)
	assert.Equal(t, expectedSlice, actualSlice)
}

func BenchmarkInspect(b *testing.B) {
	Ints[int]().Inspect(func(i int) {}).Take(b.N).Consume()
}

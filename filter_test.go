package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	iter := Elems([]int{1, 2, 3, 4}).Filter(func(i int) bool { return i%2 == 0 })

	actualFirst, _ := iter()
	expected := []int{2, 4}

	assert.Equal(t, append([]int{actualFirst}, iter.Collect()...), expected)
}

func BenchmarkFilter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(b.N).Consume()
}

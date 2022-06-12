package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTake(t *testing.T) {
	iter := Ints[int]().Take(10)
	assert.Equal(t, 5, iter.Take(5).Count())
	assert.Equal(t, 5, iter.Count())
}

func BenchmarkTake(b *testing.B) {
	Ints[int]().Take(b.N).Consume()
}

func TestTakeWhile(t *testing.T) {
	iter := Ints[int]().TakeWhile(func(i int) bool { return i < 10 })

	assert.Equal(t, Ints[int]().Take(10).Collect(), iter.Collect())

	iter = Ints[int]().Take(0).TakeWhile(func(i int) bool { return i < 10 })

	iter.Collect()

	_, ok := iter()

	assert.False(t, ok)
}

func BenchmarkTakeWhile(b *testing.B) {
	Ints[int]().TakeWhile(func(i int) bool {
		return i < b.N
	}).Consume()
}

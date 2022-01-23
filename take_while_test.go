package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	iter := Ints[int]().TakeWhile(func(i int) bool { return i < 10 })

	assert.Equal(
		t, iter.Collect(),
		Ints[int]().Take(10).Collect())

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

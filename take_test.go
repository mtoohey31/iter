package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTake(t *testing.T) {
	iter := Ints[int]().Take(10)
	assert.Equal(t, iter.Take(5).Count(), 5)
	assert.Equal(t, iter.Count(), 5)
}

func BenchmarkTake(b *testing.B) {
	Ints[int]().Take(b.N).Consume()
}

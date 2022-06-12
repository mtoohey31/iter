package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElems(t *testing.T) {
	expected := []string{"item1", "item2"}
	assert.Equal(t, expected, Elems(expected).Collect())
}

func BenchmarkElems(b *testing.B) {
	slice := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = i
	}

	Elems(slice)
}

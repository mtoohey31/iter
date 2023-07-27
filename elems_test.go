package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/internal/testutils"
)

func FuzzElems(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		assert.Equal(t, b, Elems(b).Collect())
	})
}

func BenchmarkElems(b *testing.B) {
	slice := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = i
	}

	Elems(slice).Consume()
}

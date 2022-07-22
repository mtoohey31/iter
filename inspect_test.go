package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzInspect(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		actual := []byte{}
		Elems(b).Inspect(func(v byte) {
			actual = append(actual, v)
		}).Consume()
		assert.Equal(t, b, actual)
	})
}

func BenchmarkInspect(b *testing.B) {
	Ints[int]().Inspect(func(i int) {}).Take(b.N).Consume()
}

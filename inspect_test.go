package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/testutils"
)

func FuzzIter_Inspect(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		actual := []byte{}
		Elems(b).Inspect(func(v byte) {
			actual = append(actual, v)
		}).Consume()
		assert.Equal(t, b, actual)
	})
}

func BenchmarkIter_Inspect(b *testing.B) {
	Ints[int]().Inspect(func(i int) {}).Take(uint(b.N)).Consume()
}

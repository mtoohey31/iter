package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v2/testutils"
)

func FuzzIter_Chain(f *testing.F) {
	testutils.AddByteSlicePairs(f)

	f.Fuzz(func(t *testing.T, b1 []byte, b2 []byte) {
		iter := Elems(b1).Chain(Elems(b2))
		actual := iter.Collect()
		assert.Equal(t, append(b1, b2...), actual)
	})
}

func BenchmarkIter_Chain(b *testing.B) {
	Ints[int]().Take(uint(b.N) / 2).Chain(IntsFrom(b.N / 2).Take(uint(b.N) / 2)).Consume()
}

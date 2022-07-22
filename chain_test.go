package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzChain(f *testing.F) {
	testutils.AddByteSlicePairs(f)

	f.Fuzz(func(t *testing.T, b1 []byte, b2 []byte) {
		iter := Elems(b1).Chain(Elems(b2))
		actual := iter.Collect()
		assert.Equal(t, append(b1, b2...), actual)
	})
}

// operations should not take much longer than that of the range iterator
func BenchmarkChain(b *testing.B) {
	Ints[int]().Take(b.N / 2).Chain(IntsFrom(b.N / 2).Take(b.N / 2)).Consume()
}

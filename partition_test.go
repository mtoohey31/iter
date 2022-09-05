package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzIter_Partition(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		l, r := Elems(b).Partition(func(v byte) bool { return v%2 == 0 })

		even := []byte{}
		odd := []byte{}
		for _, v := range b {
			if v%2 == 0 {
				even = append(even, v)
			} else {
				odd = append(odd, v)
			}
		}

		assert.Equal(t, even, l.Collect())
		assert.Equal(t, odd, r.Collect())

		l, r = Elems(b).Partition(func(v byte) bool { return v%2 == 0 })

		assert.Equal(t, odd, r.Collect())
		assert.Equal(t, even, l.Collect())
	})
}

func BenchmarkIter_Partition(b *testing.B) {
	iterA, iterB := Ints[int]().Partition(func(i int) bool {
		return i%2 == 0
	})
	iterA.Take(uint(b.N)).Consume()
	iterB.Take(uint(b.N)).Consume()
}

package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/testutils"
)

func FuzzIter_Seqs(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, n uint) {
		nSeqs := len(b) - int(n)
		actual := Seqs(Elems(b), n).Collect()
		if nSeqs > 0 {
			seqs := make([][]byte, nSeqs)
			for i := 0; i < nSeqs; i++ {
				w := make([]byte, n+1)
				for j := 0; j < int(n+1); j++ {
					w[j] = b[i+j]
				}
				seqs[i] = w
			}
			assert.Equal(t, seqs, actual)
		} else {
			assert.Empty(t, actual)
		}
	})
}

func BenchmarkIter_Seqs_1(b *testing.B) {
	Seqs(Ints[int](), 1).Take(uint(b.N)).Consume()
}

func BenchmarkIter_Seqs_3(b *testing.B) {
	Seqs(Ints[int](), 1).Take(uint(b.N)).Consume()
}

func BenchmarkIter_Seqs_10(b *testing.B) {
	Seqs(Ints[int](), 10).Take(uint(b.N)).Consume()
}

func BenchmarkIter_Seqs_100(b *testing.B) {
	Seqs(Ints[int](), 100).Take(uint(b.N)).Consume()
}

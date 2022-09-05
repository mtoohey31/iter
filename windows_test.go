package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzIter_Windows(f *testing.F) {
	testutils.AddByteSliceUintPairs(f)

	f.Fuzz(func(t *testing.T, b []byte, m uint) {
		// ensures n is >= 1
		n := int(m + 1)
		nWindows := len(b) - n + 1
		actual := Windows(Elems(b), n).Collect()
		if nWindows > 0 {
			windows := make([][]byte, nWindows)
			for i := 0; i < nWindows; i++ {
				w := make([]byte, n)
				for j := 0; j < n; j++ {
					w[j] = b[i+j]
				}
				windows[i] = w
			}
			assert.Equal(t, windows, actual)
		} else {
			assert.Empty(t, actual)
		}
	})
}

func TestIter_Windows_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Windows should've panicked")
		}
	}()

	Windows(Elems([]bool{}), 0)
}

func BenchmarkIter_Windows_1(b *testing.B) {
	Windows(Ints[int](), 1).Take(b.N).Consume()
}

func BenchmarkIter_Windows_3(b *testing.B) {
	Windows(Ints[int](), 1).Take(b.N).Consume()
}

func BenchmarkIter_Windows_10(b *testing.B) {
	Windows(Ints[int](), 10).Take(b.N).Consume()
}

func BenchmarkIter_Windows_100(b *testing.B) {
	Windows(Ints[int](), 100).Take(b.N).Consume()
}

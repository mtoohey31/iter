package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v3/internal/testutils"
)

func FuzzIter_Mutex(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		actual := uint(0)
		iter := Ints[uint]().Inspect(func(v uint) {
			actual += v
		}).Take(n).Mutex()

		done := false
		for !done {
			go func() {
				_, ok := iter()
				if !ok {
					done = true
				}
			}()
		}

		assert.Equal(t, (n-1)*n/2, actual)
	})
}

func BenchmarkIter_Mutex(b *testing.B) {
	Ints[int]().Take(uint(b.N)).Mutex().Consume()
}

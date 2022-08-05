package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/testutils"
)

func FuzzMutex(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		actual := uint(0)
		iter := Ints[uint]().Inspect(func(v uint) {
			actual += v
		}).Take(int(n)).Mutex()

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

func BenchmarkMutex(b *testing.B) {
	Ints[int]().Take(b.N).Mutex().Consume()
}

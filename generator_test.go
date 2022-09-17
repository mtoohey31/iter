package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v2/testutils"
)

func FuzzGenWhile(f *testing.F) {
	testutils.AddUints(f)

	f.Fuzz(func(t *testing.T, n uint) {
		expected := make([]uint, n)
		for i := uint(0); i < n; i++ {
			expected[i] = i
		}
		u := uint(0)
		assert.Equal(t, expected, GenWhile(func() (uint, error) {
			u++
			if u-1 < n {
				return u - 1, nil
			}

			return 0, assert.AnError
		}).Collect())
	})
}

func BenchmarkGenWhile(b *testing.B) {
	GenWhile(func() (int, error) {
		return 0, nil
	}).Take(uint(b.N)).Consume()
}

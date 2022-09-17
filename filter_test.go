package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"mtoohey.com/iter/v2/testutils"
)

func FuzzIter_Filter(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := []byte{}
		for _, v := range b {
			if v%2 == 0 {
				expected = append(expected, v)
			}
		}

		assert.Equal(t, expected, Elems(b).Filter(func(v byte) bool {
			return v%2 == 0
		}).Collect())
	})
}

func BenchmarkIter_Filter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(uint(b.N)).Consume()
}

func FuzzFilterMap(f *testing.F) {
	testutils.AddByteSlices(f)

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := []byte{}
		for _, v := range b {
			if v%2 == 0 {
				expected = append(expected, v*2)
			}
		}

		predicate := func(v byte) (byte, error) {
			if v%2 != 0 {
				return 0, assert.AnError
			}

			return v * 2, nil
		}

		assert.Equal(t, expected, Elems(b).FilterMap(predicate).Collect())
		assert.Equal(t, expected, FilterMap(Elems(b), predicate).Collect())
	})
}

func BenchmarkIter_FilterMap(b *testing.B) {
	Ints[int]().FilterMap(func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		}

		return 0, assert.AnError
	}).Take(uint(b.N)).Consume()
}

func BenchmarkFilterMap(b *testing.B) {
	FilterMap(Ints[int](), func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		}

		return 0, assert.AnError
	}).Take(uint(b.N)).Consume()
}

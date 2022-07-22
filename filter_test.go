package iter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"mtoohey.com/iter/testutils"
)

func FuzzFilter(f *testing.F) {
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

func BenchmarkFilter(b *testing.B) {
	Ints[int]().Filter(func(i int) bool {
		return i%2 == 0
	}).Take(b.N).Consume()
}

func FuzzFilterMap(f *testing.F) {
	testutils.AddByteSlices(f)
	err := errors.New("")

	f.Fuzz(func(t *testing.T, b []byte) {
		expected := []byte{}
		for _, v := range b {
			if v%2 == 0 {
				expected = append(expected, v*2)
			}
		}

		predicate := func(v byte) (byte, error) {
			if v%2 != 0 {
				return 0, err
			} else {
				return v * 2, nil
			}
		}

		assert.Equal(t, expected, Elems(b).FilterMapEndo(predicate).Collect())
		assert.Equal(t, expected, FilterMap(Elems(b), predicate).Collect())
	})
}

func BenchmarkFilterMapEndo(b *testing.B) {
	err := errors.New("")

	Ints[int]().FilterMapEndo(func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, err
		}
	}).Take(b.N).Consume()
}

func BenchmarkFilterMap(b *testing.B) {
	err := errors.New("")

	FilterMap(Ints[int](), func(i int) (int, error) {
		if i%2 == 0 {
			return i * 2, nil
		} else {
			return 0, err
		}
	}).Take(b.N).Consume()
}
